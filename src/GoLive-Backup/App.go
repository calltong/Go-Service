package main

import (
	"golang.org/x/net/websocket"
	"fmt"
	"log"
	"net/http"
	"db"
	"config"
	"GoLive/router"
	"GoLive/data"
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

		data.InitLiveData()
		err = db.InitSession(path)
		if err == nil {
			http.Handle("/user", websocket.Handler(router.HandleUser))
			http.Handle("/admin", websocket.Handler(router.HandleAdmin))
			http.HandleFunc("/test", Index)
			fmt.Println(name, " Server Database on" , configuration.Database.IpAddress, ":", configuration.Database.Port)
			fmt.Println(name, " Server Ready on port" , configuration.Port)

    	if err := http.ListenAndServe(fmt.Sprintf(":%d", configuration.Port), nil); err != nil {
        log.Fatal("ListenAndServe:", err)
    	}
		} else {
			fmt.Println("Cannot connect Database: ", err)
		}
	} else {
		fmt.Println("Cannot load Configuration: ", err)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Welcome!")
}
