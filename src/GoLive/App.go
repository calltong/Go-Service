package main

import (
	"fmt"
	"log"
	"net/http"
	"db"
	"config"
	"GoLive/router"
	"GoLive/common/database"
)

func main() {
	name := "Go Live"
	fmt.Println(name, " Server Starting..")
	configuration, err := config.LoadConfiguration("./src/GoLive/config.json")
	if err == nil {
		path := fmt.Sprintf("mongodb://%s:%s@%s:%d",
								configuration.Database.User.Username,
								configuration.Database.User.Password,
								configuration.Database.IpAddress,
								configuration.Database.Port)

	  Initial()
		err = db.InitSession(path)
		if err == nil {
			fmt.Println(name, " Server Database on" , configuration.Database.IpAddress, ":", configuration.Database.Port)
			fmt.Println(name, " Server Ready on port" , configuration.Port)

			routers := router.CreateAllRouter()
    	if err := http.ListenAndServe(fmt.Sprintf(":%d", configuration.Port), routers); err != nil {
        log.Fatal("ListenAndServe:", err)
    	}
		} else {
			fmt.Println("Cannot connect Database: ", err)
		}
	} else {
		fmt.Println("Cannot load Configuration: ", err)
	}
}

func Initial() {
	database.Initial(config.GetDatabaseName())
}
