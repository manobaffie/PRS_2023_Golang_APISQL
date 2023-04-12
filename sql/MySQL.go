package mySQL

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func init() { }

type MySQLInt struct { value int }
type MySQLUInt struct { value int }
type MySQLString struct { value string }

type MySQLTable struct {
	id MySQLUInt
	table interface{}
}

type MySQLUser struct {
	Name     string
	Password string
}

type MySQLEndpoint struct {
	Ip   string
	Port string
	Path string
}

type MySQL struct {
	User     MySQLUser
	Endpoint MySQLEndpoint
	db       *sql.DB
}

func (mySQLTable *MySQLTable)Set(table interface{}) {
		
}

func (mySQLInt *MySQLInt)Set(value int) {
	mySQLInt.value = value
}

func (mySQLInt *MySQLInt)Build() (string) {
	return ""
}

func (mySQLString *MySQLString)Set(value string) {
	mySQLString.value = value
}

func (mySQLString *MySQLString)Build() (string) {
	return ""
}

func (mySQLUser *MySQLUser) New() {
	mySQLUser.Name = "root"
	mySQLUser.Password = ""
}

func (mySQLConnection *MySQLEndpoint) New() {
	mySQLConnection.Ip = "127.0.0.1"
	mySQLConnection.Port = "3306"
	mySQLConnection.Path = ""
}

func (mysql *MySQL) New() {
	mysql.User.New()
	mysql.Endpoint.New()
}

func (mysql *MySQL) Close() {
	mysql.db.Close()
}

func (mysql *MySQL) ConnectionString() string {
	var dataSourceName = mysql.User.Name
	if mysql.User.Password != "" {
		dataSourceName += ":" + mysql.User.Password
	}
	dataSourceName += "@tcp(" + mysql.Endpoint.Ip + ":" + mysql.Endpoint.Port + ")/"
	return dataSourceName
}

func (mysql *MySQL) Open() {
	db, err := sql.Open("mysql", mysql.ConnectionString() + mysql.Endpoint.Path)
	if err != nil {
		log.Fatal(err.Error())
	}
	mysql.db = db
}

func (mysql *MySQL) Create() {
	db, err := sql.Open("mysql", mysql.ConnectionString())
	if err != nil {
		log.Fatal(err.Error())
	}
	db.Exec("CREATE DATABASE IF NOT EXISTS " + mysql.Endpoint.Path)
	if err != nil {
		log.Fatal(err.Error())
	}
	mysql.db = db
}

func (mySQL *MySQL) CreateTable(table *MySQLTable) {
	_ , err := mySQL.db.Exec("CREATE TABLE example ( id integer, data varchar(32) )")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (mysql *MySQL) Drop() {
	db, err := sql.Open("mysql", mysql.ConnectionString())
	if err != nil {
		log.Fatal(err.Error())
	}
	db.Exec("DROP DATABASE IF EXISTS " + mysql.Endpoint.Path)
	if err != nil {
		log.Fatal(err.Error())
	}
	mysql.db = db
}

func (mysql *MySQL) IsConnected() (bool) {
	pingErr := mysql.db.Ping()
    if pingErr != nil {
		log.Fatal(pingErr.Error())
		return false
    }
	return true
}
