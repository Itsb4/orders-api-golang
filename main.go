package main

import (
	"context"
	"fmt"

	"github.com/it-shubham/orders-api-golang/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start server: ", err)
	}
}
