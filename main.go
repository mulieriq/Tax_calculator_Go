package main

import (
	"fmt"
	"sync"
	"tax_calculator/models"
	"tax_calculator/services"
	"time"
)

var wg = &sync.WaitGroup{}

///

func EmailResults(user models.StringConv, number int) {
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
		go EmailResults(userDetails, counter)

	}
	println("| No. | Name | Age | Gross Pay  | Tax | Net Pay | Student | Inst | KRA  |")

	println("|-----|------|-----|------------|------|-----------|--------|------|------|")
	for index, persons := range users {
		fmt.Println(persons.ToString(index + 1))
	}
	wg.Wait()

}

/*







 */
