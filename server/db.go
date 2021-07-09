package server

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

// InitDB функция, инициирующая подключение к БД
func InitDB() (err error) {
	//строка, содержащая данные для подключения к БД в следующем формате:
	var dataSourceName = "golang01:infomaximum01@tcp(t1.infomaximum.com:9301)/golang01"
	//подключаемся к БД, используя нужный драйвер и данные для подключения
	Db, err = sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	err = Db.Ping()
	return
}
