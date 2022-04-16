package services

import (
	"bufio"
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

func EmailResults(user models.StringConv, number int, wg *sync.WaitGroup) {
	//emailFrom := "email"
	//emailTo := []string{"email"}
	//psw := "pass"
	//smtpHost := "smtp.gmail.com"
	//smtpPort := "587"
	fmt.Println("Sending Email for", user.ToMap()["name"])
	time.Sleep(time.Second * 20)
	//auth := smtp.PlainAuth("", emailFrom, psw, smtpHost)
	//fmt.Println(auth)
	//err := smtp.SendMail(smtpHost+":"+smtpPort, auth, emailFrom, emailTo, []byte(user.ToString(number)))
	//if err != nil {
	//	fmt.Println(err)
	//}

	fmt.Println("Email Sent")
	defer wg.Done()
}
