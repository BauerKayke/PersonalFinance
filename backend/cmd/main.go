package main

import (
	"backend/internal/application"
	"fmt"
	"net/http"
)

func main() {
	app := application.NewApp()
	if err := app.Run(); err != nil && err != http.ErrServerClosed {
		fmt.Println(err)
		return
	}
}
