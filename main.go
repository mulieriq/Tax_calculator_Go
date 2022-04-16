package main

import (
	"fmt"
	"sync"
	"tax_calculator/models"
	"tax_calculator/services"
)

var wg = &sync.WaitGroup{}

func main() {
	defer fmt.Println("Thank you")
	func() {
		fmt.Print("\nTaxify...Tulipe Ushuru\n")
	}()

	users := make([]models.StringConv, 0)
	for counter := 0; counter < 2; counter++ {
		userDetails := services.GetUserInput()
		users = append(users, userDetails)
		wg.Add(1)
		go services.EmailResults(userDetails, counter, wg)

	}
	println("| No. | Name | Age | Gross Pay  | Tax | Net Pay | Student | Inst | KRA  |")

	println("|-----|------|-----|------------|------|-----------|--------|------|------|")
	for index, persons := range users {
		fmt.Println(persons.ToString(index + 1))
	}
	wg.Wait()
}
