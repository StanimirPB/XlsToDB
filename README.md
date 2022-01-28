# XlsToDB
This is script convert Excel(XLXS) file to Postgres table

for correct work this script need two external package:
1. "github.com/jack/pgx/v4/pgxpool"
2. "github.com/360EntSecGroup-Skylar/exelize/v2"

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
var schema = "yourSchema" //Database Schema

#it is Name of new Postgres Table
var table = "yourTable" //Database Table
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

