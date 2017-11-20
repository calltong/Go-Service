package main

import (
	"fmt"
	"log"
	"net/http"
	"db"
	"config"
  "router"
)

func main() {
	name := "Go Shopping"
	fmt.Println(name, " Server Starting..")
	configuration, err := config.LoadConfiguration("./src/HexStore/config.json")
	if err == nil {
		path := fmt.Sprintf("mongodb://%s:%s@%s:%d", 	configuration.Database.User.Username,
																						configuration.Database.User.Password,
																						configuration.Database.IpAddress,
																						configuration.Database.Port)
		err = db.InitSession(path)
		if err == nil {
			router := router.CreateAllRouter()
			fmt.Println(name, " Server Database on" , configuration.Database.IpAddress, ":", configuration.Database.Port)
			fmt.Println(name, " Server Ready on port" , configuration.Port)
			log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", configuration.Port), router))
		} else {
			fmt.Println("Cannot connect Database: ", err)
		}
	} else {
		fmt.Println("Cannot load Configuration: ", err)
	}
}
