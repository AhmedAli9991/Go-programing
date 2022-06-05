package main

import (
	"sessionauth/model"
	"sessionauth/routes"
)


func main() {
	model.Setup()
	routes.Setup()
}