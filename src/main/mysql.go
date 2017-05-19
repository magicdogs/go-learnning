package main

import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"pkg/assert"
)


func main() {
	db,err :=sql.Open("mysql", "root:123456@tcp(localhost:3306)/iam?charset=utf8")


	//db,err :=sql.Open("mysql", "devapp:DEvasdfyuio@/192.168.105.28:3306/meta_space")
	assert.NotError(err)
	defer db.Close()

	/*stmtIns, err := db.Prepare("INSERT INTO t_role_permission(ROLE_ID,PERMISSION_ID,CREATE_AT) VALUES(?,?,?)") // ? = placeholder
	assert.NotError(err)
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates
	result,err := stmtIns.Exec(3,4,"2016-01-01 11:01:02")
	assert.NotError(err)
	fmt.Println(result)*/

	stmtQuery,err :=db.Prepare("select * from  t_role_permission");
	assert.NotError(err)
	rows,err :=stmtQuery.Query()
	assert.NotError(err)
	var id ,rid,pid int
	var date string
	for  rows.Next() {
		err = rows.Scan(&id,&rid,&pid,&date)
		assert.NotError(err)
		fmt.Println("id: ",id,",date: ",date)
	}

	/*// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT squareNumber FROM squarenum WHERE number = ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	// Insert square numbers for 0-24 in the database
	for i := 0; i < 25; i++ {
		_, err = stmtIns.Exec(i, (i * i)) // Insert tuples (i, i^2)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}

	var squareNum int // we "scan" the result in here

	// Query the square-number of 13
	err = stmtOut.QueryRow(13).Scan(&squareNum) // WHERE number = 13
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("The square number of 13 is: %d", squareNum)

	// Query another number.. 1 maybe?
	err = stmtOut.QueryRow(1).Scan(&squareNum) // WHERE number = 1
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("The square number of 1 is: %d", squareNum)*/

}