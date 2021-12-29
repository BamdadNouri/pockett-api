package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sandbox/pockett-api/config"
	"sandbox/pockett-api/internal/handlers"
	"sandbox/pockett-api/internal/repositories"
	"sandbox/pockett-api/internal/store"

	"github.com/fvbock/endless"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	cmd := &cobra.Command{
		Use:   "api",
		Short: "use api to run Pockett API",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().String("config", "", "config file path")

			cmd.Flags().String("api_base", "api", "API base path")

			cmd.Flags().String("port", "8009", "pockett API port")

			cmd.Flags().String("database_host", "127.0.0.1", "database host")
			// cmd.Flags().String("database_host", "db", "database host")
			cmd.Flags().String("database_user", "pockett", "database user")
			cmd.Flags().String("database_password", "password", "database password")
			cmd.Flags().String("database_name", "pockett", "database name")
			cmd.Flags().String("database_port", "3306", "database port")
			cmd.Flags().String("database_ssl", "disabled", "database ssl mode")

			cmd.Flags().String("jwt_secret", "", "jwt secret")

			cmd.Flags().Bool("with_migration", false, "run DB migration at startup")

			err := cmd.ParseFlags(args)
			if err != nil {
				return err
			}
			configFilePath := cmd.Flags().Lookup("config").Value.String()
			if configFilePath != "" {
				viper.SetConfigFile(configFilePath)
				err := viper.ReadInConfig()
				if err != nil {
					return err
				}
			}
			err = viper.BindPFlags(cmd.Flags())
			if err != nil {
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runAPI()
		},
	}
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runAPI() error {
	config := config.NewConfig()
	db, err := store.NewStore(config.Database.DSN(), config.Database.Migrate).
		Init().
		RunMigration().
		Result()
	if err != nil {
		return fmt.Errorf("error in initializing database", err)
	}

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	engine := gin.Default()
	engineConfig := cors.DefaultConfig()
	engineConfig.AllowOriginFunc = func(origin string) bool {
		return true
	}
	engineConfig.AllowCredentials = true
	engineConfig.AllowHeaders = []string{
		"Origin", "Content-Length", "Content-Type", "X-Screen-Height", "X-Screen-Width", "Authorization",
	}
	engine.Use(cors.New(engineConfig))

	userRepo, transactionRepo, walletRepo := repositories.InitRepositories(db)

	userHandler := handlers.NewUserHandler(userRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionRepo)
	walletHandler := handlers.NewWalletHandler(walletRepo)

	baseAPIGroup := engine.Group(config.BaseURL)
	baseAPIGroup.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, "healthy")
	})
	apiGroup := baseAPIGroup.Group("")

	transactionGroup := apiGroup.Group("transaction")
	{
		transactionGroup.POST("/", transactionHandler.Add)
		transactionGroup.PUT("/", transactionHandler.Update)
		transactionGroup.GET("/", transactionHandler.Get)
		transactionGroup.DELETE("/:id", transactionHandler.Delete)
	}

	// tagGroup := apiGroup.Group("tag")
	// {
	// 	tagGroup.POST("/", tagHandler.Add)
	// 	tagGroup.PUT("/", tagHandler.Update)
	// 	tagGroup.GET("/", tagHandler.GetBulk)
	// 	tagGroup.DELETE("/", tagHandler.Delete)
	// }
	walletGroup := apiGroup.Group("wallet")
	{
		walletGroup.POST("/", walletHandler.Add)
		walletGroup.PUT("/", walletHandler.Update)
		walletGroup.GET("/", walletHandler.Get)
		walletGroup.DELETE("/", walletHandler.Delete)
	}
	userGroup := apiGroup.Group("user")
	{
		userGroup.POST("/register", userHandler.Add)
		userGroup.POST("/login", userHandler.Login)
		// userGroup.POST("/oauth/cb", userHandler.Add)
		userGroup.PUT("/", userHandler.Update)
		userGroup.GET("/me", userHandler.Me)
	}

	err = endless.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", config.Port), engine)
	if err != nil {
		return err
	}
	fmt.Println("Pockett API listening on port ", config.Port)
	return nil
}
