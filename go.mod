module github.com/mzky/mysqlbakup

go 1.12

require (
	github.com/go-sql-driver/mysql v1.4.1
	github.com/mzky/common
)

replace github.com/mzky/common => ../common
