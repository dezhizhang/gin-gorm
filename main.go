package main

import "gin-gorm/router"

func main() {
	engine := router.Router()

	engine.Run()

}
