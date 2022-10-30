package main

import (
	"github.com/ginblog/model"
	"github.com/ginblog/routes"
)

func main() {
	model.InitDb()
	routes.InitRouter()
}
