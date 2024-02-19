package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/beego/beego/v2/client/orm"

	_ "github.com/lib/pq"
)

type DBInstance struct {
	Db  *sql.DB
	Orm orm.Ormer
	Qb  orm.QueryBuilder
}

var Database DBInstance

func ConnectDB() {
	var dbURL string
	if os.Getenv("environment") == "test" {
		dbURL = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
			os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	} else {
		dbURL = fmt.Sprintf("user=%s password=%s database=%s host=%s",
			os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_INSTANCE"))
	}

	fmt.Println(dbURL)

	orm.RegisterDriver("postgres", orm.DRPostgres)
	err := orm.RegisterDataBase("default", "postgres", dbURL)
	if err != nil {
		fmt.Println(err)
	}
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println(err)
	}

	orm.Debug = true
	o := orm.NewOrm()

	qb, _ := orm.NewQueryBuilder("postgres")

	// Error.
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println(err)
	}

	Database = DBInstance{Db: db, Orm: o, Qb: qb}
	// CheckRelationship()
}

