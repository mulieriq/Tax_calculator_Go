package services

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sync"
	"tax_calculator/models"
	"time"
)

func GetUserInput() (user models.StringConv) {
	var (
		name    string
		age     int
		student int
		kraPin  string
		school  string
		stipend int
		salary  int
	)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter Your Name : ")
	scanner.Scan()
	name = scanner.Text()

	fmt.Print("Enter Your Age : ")
	fmt.Scan(&age)

	fmt.Print("Are You A Student (1-yes ,0-no): ")
	fmt.Scan(&student)

	switch student {
	case 0:
		fmt.Print("Enter Your Kra Pin : ")
		fmt.Scan(&kraPin)

		fmt.Print("Enter Your Salary : ")
		fmt.Scan(&salary)
		fmt.Print("\n\n")
	case 1:
		fmt.Print("Enter Your School Name: ")
		fmt.Scan(&school)

		fmt.Print("Enter Your Stipend Amount : ")
		fmt.Scan(&stipend)
		fmt.Print("\n\n")

	}

	std := models.Student{
		Age:     age,
		Name:    name,
		School:  school,
		Stipend: stipend,
	}
	adlt := models.Adult{
		Age:    age,
		Name:   name,
		KraPin: kraPin,
		Salary: salary,
	}
	return UsersFactory(&std, &adlt)
}

func UsersFactory(s *models.Student, a *models.Adult) (userDetails models.StringConv) {

	switch {
	case len(s.School) > 0 && s.Stipend > 0:
		return s
	default:
		return a
	}
}

func EmailResults(emailCC chan<- string, user context.Context, number int, wg *sync.WaitGroup) {

	userDetails := user.Value("user").(models.StringConv)

	fmt.Println("Sending Email for", userDetails.ToMap()["name"])
	time.Sleep(time.Second * 10)

	emailCC <- "Email Sent"
	select {
	case <-user.Done():
		fmt.Println("Timed Out")

	}

	defer wg.Done()
}
