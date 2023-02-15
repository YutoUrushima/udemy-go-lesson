package main

import (
	"fmt"
	"udemy-go-lesson/app/controllers"
	"udemy-go-lesson/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
}
