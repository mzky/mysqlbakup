package main

import (
	"database/sql"
	"fmt"

	"github.com/mzky/mysqlbakup/backup"

	"github.com/mzky/mysqlbakup/common/log"

	_ "github.com/go-sql-driver/mysql"
)

type Mysqlcfg struct {
	DBUser   string
	DBPasswd string
	DBAddr   string
	DBName   string
}

func main() {
	var cfg Mysqlcfg
	cfg.DBUser = "alliance"
	cfg.DBPasswd = "111111"
	cfg.DBAddr = "192.168.136.23:9633"
	cfg.DBName = "alliance_admin"

	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", cfg.DBUser, cfg.DBPasswd, cfg.DBAddr, cfg.DBName)
	context, err := sql.Open("mysql", dataSource)
	//context, err := sql.Open("mysql", "root:111111@tcp(192.168.22.212:3306)/x_game__s1?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("初始化数据库", err)
	}
	col := backup.NewMySqlBackup(context)

	col.ExportToFile("./1.sql")

	fmt.Println("====1")
	_, err = sql.Open("mysql", dataSource)
	//newdb, err := sql.Open("mysql", "root:root@tcp(192.168.22.212:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("初始化数据库%v", err)
	}
	//备份表到文件
	err = col.ExportTableToFile("./2.sql", "table")
	if err != nil {
		log.Fatal("导出表:%v", err)
	}
	//new := backup.NewMySqlBackup(newdb)
	//new.ImportFromFile("./11.sql")
	//fmt.Println("====2")
	//log.Debug("bbb:%v", col)

}
