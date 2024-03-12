package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upAddTableUser, downAddTableUser)
}

func upAddTableUser(tx *sql.Tx) error {
	_, err := tx.Exec(`CREATE TABLE user (
		id int(11) NOT NULL AUTO_INCREMENT,
		username varchar(200) NOT NULL DEFAULT '',
		age int(11) DEFAULT NULL,
		create_time datetime DEFAULT NULL,
		status tinyint(2) DEFAULT NULL,
		PRIMARY KEY (id)
	  ) ENGINE=InnoDB AUTO_INCREMENT=84 DEFAULT CHARSET=utf8mb4
	`)
	return err
}

func downAddTableUser(tx *sql.Tx) error {
	_, err := tx.Exec("drop table user if exists")
	return err
}
