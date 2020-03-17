module github.com/mzky/mysqlbakup

go 1.12

require (
	github.com/go-sql-driver/mysql v1.4.1
	common v0.0.0-20190810091352-ac74c668c3bc
)

replace common => ../common
