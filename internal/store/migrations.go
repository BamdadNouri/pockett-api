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
		tr_type	 INT,
		description VARCHAR(1024),
		tag_id INT,
		owner_id INT,
		FOREIGN KEY(owner_id) REFERENCES users(id),
		wallet_id INT,
		FOREIGN KEY(wallet_id) REFERENCES wallets(id),
		is_deleted BOOLEAN
	);`
	CreateWalletsTable = `CREATE TABLE wallets(
		id INT PRIMARY KEY AUTO_INCREMENT,
		title VARCHAR(15),
		owner_id	INT,
		FOREIGN KEY(owner_id) REFERENCES users(id),
		is_deleted 	BOOLEAN
		);`
	CreateUserRecord        = `INSERT INTO users VALUES(0,'admin@pockett.ir','admin','xxxx',1,0001,0,TRUE);`
	CreateWalletRecord      = `INSERT INTO wallets VALUES(0,'my wallet',1,false);`
	CreateTransactionRecord = `INSERT INTO transactions VALUES(0,10000,1,'init',2,1,1,false);`
)
