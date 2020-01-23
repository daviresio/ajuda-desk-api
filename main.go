package main

import (
	"fmt"
	"github.com/daviresio/ajuda-desk-api/controller"
)

func main() {
	err := controller.GetRouter().Run(":8000")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}