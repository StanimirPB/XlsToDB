package main
import (
  "fmt"
  "context"
  "github.com/jack/pgx/v4/pgxpool"
  "github.com/360EntSecGroup-Skylar/exelize/v2"
)
//==========================================================================
//DB settings
var DB *pgxpool.Pool
var server_ip = "127.0.0.1" //Server IP
var port = "5433" //Server port
var database = "YourDatabase" //Database
var login = "YourLogin" //Login
var pass = "YourPassword" //Password
var schema = "yourSchema" //Database Schema
var table = "yourTable" //Database Table
//--------------------------------------------------------------------------
//Excel settings
var fName = "yourxlsx"//file name
var ex_sheet_name := "yourSheet" //excel sheet name
var ex_header_row = 4 //row where located headler
var ex_coutent_start_row = 5 //row where content starts
//==========================================================================
//rows counter - it neead for print cont
var rowsCnt = 1

func main()  {
  OpenXls()
}

func OpenXls ()  {
  //init DB
  ctx := context.Background(); Init(ctx); defer Close()
  //open xlsx file
  fmt.Printf("Open file %v\n",fName)
  f,_ := excelize.OpenFile(fName)
  //count rows
  ex_rows,_ := f.GetRows(ex_sheet_name)// get rows from Your sheet
  fmt.Printf("Count rows - %v\n", len(ex_rows))
  //read file by rows
  var colname []string
  for k,row := range ex_rows {
    var []string
    //read rows by cols
    for _,collNum := range row {
      str := append(str,collNum)
    }
    //create table if it header
    if k == ex_header_row {
      colname = str //coll names it nead for create db with same columns
      CreateTB(ctx,str)
    }
    //isert to table if it content
    if k == ex_coutent_start_row {
      xlsToDB(ctx,colname,str) //use header(colnames) and string contains content
    }
  }
}
//function for initiate connection to database
func Init()  {
  DATABASE_URL := fmt.Sprint("postgres://%v:%v@%v:%v/%v",login,pass,server_ip,port,database)
  con,err := pgxpool.Connect(ctx,DATABASE_URL)
  if err != nil {
    fmt.Println("err: ", err)
  }
  DB = con
  fmt.Println("DB inicialize!!!")
}
//create new table
func CreateTB(ctx context.Context, collname []string)  {
  //create basic part of SQL string
  _1st_sql := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS "%v"."%v" (`,schema,table)
  //Create collname of SQL string
  for k,val := range collname {
    //add new column to SQL string
    _1st_sql = fmt.Sprintf(`%v "%v" varchar(100) NOT NULL default ''`,_1st_sql, val)
    if k != len(colname)-1 {
      _1st_sql = _1st_sql + `,`
    }
  }
  //Create final part of SQL string
  _1st_sql = _1st_sql + `);`
  DB.Excec(ctx,_1st_sql)
}
//insert content to table
func xlsToDB(ctx context.Context, colname []string, stroka []string)  {
  //Create SQL string
  sql := fmt.Sprintf(`INSERT INTO "%v"."%v" (`,schema,table)
  //Create column part of SQL
  for k,val := range colname {
    sql = fmt.Sprintf(`%v "%v"`,sql, val)
    if k != len(colname)-1 {
      sql = sql + `,`
    }
  }
  sql = sql + `) VALUES (`
  //Create values part of SQL
  for k,val := range stroka {
    sql = fmt.Sprintf(`%v "%v"`,sql, val)
    if k != len(colname)-1 {
      sql = sql + `,`
    }
  }
  //Create final part SQL
  sql = sql + `)`
  rowsCnt ++
  DB.Exec(ctx,sql)
  fmt.Printf("\r Insert row %v to DB",rowsCnt)
}
//close database connection
func close()  {
  DB.Close()
}
