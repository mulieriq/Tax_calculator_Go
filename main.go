package main

import (
	"fmt"
	"tax_calculator/models"
	"tax_calculator/services"
)

func main() {
	defer fmt.Println("Thank you")
	func() {
		fmt.Print("\nTaxify...Tulipe Ushuru\n")
	}()

	users := make([]models.StringConv, 0)
	for counter := 0; counter < 2; counter++ {
		userDetails := services.GetUserInput()
		users = append(users, userDetails)
	}
	println("|No.| Name | Age | Gross Pay | Net Pay | Student | Inst")
	println("|---|------|-----|-----------|---------|")
	for index, persons := range users {
		fmt.Println(persons.ToString(index + 1))
	}
	return
}
