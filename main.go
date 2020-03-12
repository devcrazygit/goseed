package main

import (
	"goseed/routers"
	"goseed/utils"
)

func main() {
	router := routers.InitRoute()
	port := utils.EnvVar("SERVER_PORT")
	// fmt.Println(port)
	router.Run(port)
}
