package main

import (
	"context"
	"log"
	"os"
	"runtime"
	"sync"
	"tax_calculator/models"
	"tax_calculator/services"
	"time"
)

var wg = &sync.WaitGroup{}

var emailConfirmationChannel = make(chan string, 4) // 4 threads writing to the channel
var logger = log.New(os.Stdin, "-- main node --", log.Ldate)

func main() {
	now := time.Now()
	defer close(emailConfirmationChannel)
	defer func() {
		logger.Println("Thank you")
		logger.Println("Latency- ", time.Since(now))
		logger.Println(runtime.NumGoroutine(), "Threads still running")
	}()

	calcContext, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	func() {
		logger.Print(" Taxify...Tulipe Ushuru\n")
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
				logger.Println("Email Response Confirmation status : ", response)
			case <-calcContext.Done():
				logger.Println("Deadline Reached --- Cant receive more requests")
				logger.Println(calcContext.Err())

			}
			defer wg.Done()
		}()

	}
	println("| No. | Name | Age | Gross Pay  | Tax | Net Pay | Student | Inst | KRA  |")

	println("|-----|------|-----|------------|------|-----------|--------|------|------|")
	for index, persons := range users {
		logger.Println(persons.ToString(index + 1))
	}

	wg.Wait()
}
