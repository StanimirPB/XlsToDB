# XlsToDB
This is script convert Excel(XLXS) file to Postgres table

for correct work this script need two external package:
1. "github.com/jack/pgx/v4/pgxpool"

pgx - PostgreSQL Driver and Toolkit
pgx is a pure Go driver and toolkit for PostgreSQL.
pgx aims to be low-level, fast, and performant, while also enabling PostgreSQL-specific features that the standard database/sql package does not allow for.
The driver component of pgx can be used alongside the standard database/sql package.
The toolkit component is a related set of packages that implement PostgreSQL functionality such as parsing the wire protocol and type mapping between PostgreSQL and Go. These underlying packages can be used to implement alternative drivers, proxies, load balancers, logical replication clients, etc.
The current release of pgx v4 requires Go modules. To use the previous version, checkout and vendor the v3 branch.

2. "github.com/360EntSecGroup-Skylar/exelize/v2"

Excelize is a library written in pure Go providing a set of functions that allow you to write to and read from XLAM / XLSM / XLSX / XLTM / XLTX files. Supports reading and writing spreadsheet documents generated by Microsoft Excel™ 2007 and later. Supports complex components by high compatibility, and provided streaming API for generating or reading data from a worksheet with huge amounts of data. This library needs Go version 1.15 or later. The full API docs can be seen using go's built-in documentation tool, or online at go.dev and docs reference.

Basic Usage

Installation
go get github.com/xuri/excelize

If your packages are managed using Go Modules, please install with following command.
go get github.com/xuri/excelize/v2


in the top of script you can set up parametrs for task:

//==========================================================================

//DB settings

#it is default option for BD 
var DB *pgxpool.Pool

#it is yours Postgres Database server IP. Actualy it "127.0.0.1" or "localhost"
var server_ip = "127.0.0.1"

#it is yours Postgres DataBase server Port. default "5432". In my case it was "5433"
var port = "5433"

#it is yours Postgres Database name
var database = "YourDatabase"

#it is Login for yours Postgres DataBase 
var login = "YourLogin"

#it is Password for yours Postgres DataBase
var pass = "YourPassword"

#it is Postgres DataBase schema
var schema = "yourSchema"

#it is Name of new Postgres Table
var table = "yourTable" 

//--------------------------------------------------------------------------

//Excel settings

#it is name of source xlsx file
var fName = "yourxlsx.xlsx"

#it is name of sheet in xlsx file
var ex_sheet_name := "yourSheet"

#it is num row where locate columns names (header)
var ex_header_row = 4 //row where located headler

#it is num row where are content start  
var ex_coutent_start_row = 5 

//==========================================================================


