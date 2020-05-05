package database

import (
	"fmt"

	"github.com/go-pg/pg/v9"
)

type PSQLConfig struct {
	Host     string
	User     string
	Password string
	Port     string
	Db       string
}

// db.DropTable((*Animal)(nil), &orm.DropTableOptions{
// 	IfExists: true,
// })

func New(config PSQLConfig) pg.DB {

	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		User:     config.User,
		Password: config.Password,
		Database: config.Db,
	})

	return *db
}
