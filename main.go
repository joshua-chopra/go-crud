package main

import (
	"github.com/joshua-chopra/go-crud/api"
	"github.com/joshua-chopra/go-crud/database"
	"github.com/joshua-chopra/go-crud/internal"
)

func main() {
	internal.Setup()
	database.InitializeDatabase()
	api.StartRouter()
}
