module mysqlbakup/example

go 1.14

require (
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/mzky/mysqlbakup v0.0.0-20200313153357-e172216fce1a // indirect
	github.com/mzky/mysqlbakup/common v0.0.0-20200313153357-e172216fce1a // indirect
)

replace (
	github.com/mzky/mysqlbakup/common =>  ../common
	github.com/mzky/mysqlbakup => ../
)