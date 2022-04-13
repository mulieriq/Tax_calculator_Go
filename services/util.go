package services

import (
	"bufio"
	"fmt"
	"os"
	"tax_calculator/models"
)

func GetUserInput() (user models.StringConv) {
	var name string
	var age int
	var student int
	var kraPin string
	var school string
	var stipend int
	var salary int

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
 Age: age,
		Name:    name,
		School:  school,
		Stipend: stipend,
	}
	adlt := models.Adult{
		Age: age,
		Name:   name,
		KraPin: kraPin,
		Salary: salary,
	}
	return usersFactory(&std, &adlt)
}

func usersFactory(s *models.Student, a *models.Adult) (userDetails models.StringConv) {

	switch {
	case len(s.School) > 0 && s.Stipend > 0 :
		return s
	default:
		return a
	}
}
