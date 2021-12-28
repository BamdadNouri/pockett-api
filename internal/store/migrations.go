package store

const (
	CreateUsersTable = `CREATE TABLE users(
		id  INT PRIMARY KEY AUTO_INCREMENT,
		email    NVARCHAR(255),
		username VARCHAR(15),
		password VARCHAR(512),
		theme INT,
		default_wallet_id INT ,
		is_active BOOLEAN
	);`
	CreateTagsTable = `CREATE TABLE tags(
		id INT PRIMARY KEY AUTO_INCREMENT,
		title VARCHAR(15) ,
		color TEXT,
		owner_id INT,
		FOREIGN KEY(owner_id) REFERENCES users(id)
		);`
	CreateTransactionsTable = `CREATE TABLE transactions(
		id	INT PRIMARY KEY AUTO_INCREMENT,
		amount  FLOAT,
		trtransaction_type	 INT,
		description VARCHAR(16),
		tag_id INT,
		owner_id INT,
		FOREIGN KEY(owner_id) REFERENCES users(id),
		is_deleted BOOLEAN
	);`
	CreateWalletsTable = `CREATE TABLE wallets(
		id INT PRIMARY KEY AUTO_INCREMENT,
		title VARCHAR(15),
		owner_id	INT,
		FOREIGN KEY(owner_id) REFERENCES users(id),
		is_deleted 	BOOLEAN
		);`
	CreateUserRecord        = `INSERT INTO users VALUES(0,'admin@pockett.ir','admin','xxxx',1,0001,TRUE);`
	CreateWalletRecord      = `INSERT INTO	transactions VALUES(0,10000,1,'init',2,1,false);`
	CreateTransactionRecord = `INSERT INTO	wallets VALUES(0,'my wallet',1,false);`
)
