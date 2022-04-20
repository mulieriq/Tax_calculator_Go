package main

import (
	"context"
	"fmt"
	"sync"
	"tax_calculator/models"
	"tax_calculator/services"
	"time"
)

var wg = &sync.WaitGroup{}

var emailConfirmationChannel = make(chan string)

func main() {
	now := time.Now()

	defer close(emailConfirmationChannel)
	defer func() {
		fmt.Println("Thank you")
		fmt.Println("Time Used - ", time.Since(now))
	}()

	calcContext, cancel := context.WithTimeout(context.Background(), 35*time.Second)
	defer cancel()

	func() {
		fmt.Print("\nTaxify...Tulipe Ushuru\n")
	}()

	users := make([]models.StringConv, 0)
	for counter := 0; counter < 2; counter++ {
		userDetails := services.GetUserInput()

		users = append(users, userDetails)
		calcContext = context.WithValue(calcContext, "user", userDetails)
		wg.Add(2)
		go services.EmailResults(emailConfirmationChannel, calcContext, counter, wg)

		go func() {
			select {
			case response := <-emailConfirmationChannel:
				fmt.Println("Email Response Confirmation status : ", response)
			case <-calcContext.Done():
				fmt.Println("Deadline Reached --- Cant receive more requests")
				fmt.Println(calcContext.Err())

			}
			defer wg.Done()
		}()

	}
	println("| No. | Name | Age | Gross Pay  | Tax | Net Pay | Student | Inst | KRA  |")

	println("|-----|------|-----|------------|------|-----------|--------|------|------|")
	for index, persons := range users {
		fmt.Println(persons.ToString(index + 1))
	}

	wg.Wait()
}
