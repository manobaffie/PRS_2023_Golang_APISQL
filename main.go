package main

import (
	testSql "APISQL/sql"
)

type MyTestTable struct {
	id testSql.MySQLInt
	name testSql.MySQLString
}

func my() {
	var mySQL testSql.MySQL

	mySQL.New()
	mySQL.User.Name = "PRS_2023_GOLANG_APISQL_USER"
	mySQL.User.Password = "toor"
	// mySQL.Endpoint.Path = "PRS_2023_GOLANG_APISQL"
	mySQL.Endpoint.Path = "PRS_2023_GOLANG_APISQL_2"

	mySQL.Open()

	if (!mySQL.IsConnected()) {
		mySQL.Close()
	}


	var testTable *testSql.MySQLTable

	testTable.Set()

	mySQL.CreateTable(testTable)
}

func main() {
	my()
}
