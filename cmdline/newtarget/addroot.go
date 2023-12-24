package main

import (

	"fmt"
	"flag"
	"database/sql"
	"log"
	"os"
	"github.com/HawkEye/logger"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

)

var (

	name string 
	platform string
	root_domains string
	excluded_domains string

)

type target struct {
	name string
	platform string
	root_domains string
	excluded_domains string
}

func main () {

	err := godotenv.Load()
	if err != nil {
		logger.Log("Error loading .env file")
	}

	flag.StringVar(&name,"name","","Specify the Target name")
	flag.StringVar(&platform, "platform", "Indepedent" , "The platform on which target bug bounty program is hosted, (default: Independent)")
	flag.StringVar(&root_domains,"inscope","", "A comma separated list of root domains included in the scope of the target.")
	flag.StringVar(&excluded_domains, "outscope","","A comma separated list of root domains related to the target to exclude.")

	flag.Parse()

	// fmt.Println("DBUSER: "+os.Getenv("DBUSER")+", DBPASS: "+os.Getenv("DBPASS"))
	conf := mysql.Config{

		User: os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net: "tcp",
		Addr: "127.0.0.1:3306",
		DBName: "bounty",
		AllowNativePasswords : true,
	}

	db,err := sql.Open("mysql",conf.FormatDSN())

	if err != nil {
		logger.Log("Could not connect to mysql")
	}
	rows,err := db.Query("SELECT NAME,PLATFORM,SCOPE,OUTSCOPE FROM targets")
	if err != nil {
		logger.Log("Query failed")
	}


	for rows.Next()  {
		var t target
		rows.Scan(&t.name,&t.platform,&t.root_domains,&t.excluded_domains)
		fmt.Println(t)
	}


}