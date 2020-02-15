package main

import (
	"fmt"
	"github.com/daviresio/ajuda-desk-api/database"
	"github.com/daviresio/ajuda-desk-api/routes"
)

func main() {
	//database.SeedAll()

	err := routes.GetRouter().Run(":8000")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer database.DB.Close()

}