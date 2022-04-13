package models

import "strconv"

type Student struct {
	Name    string
	School  string
	Stipend int
	Tax     int
	Age     int
}

type Adult struct {
	Name   string
	KraPin string
	Salary int
	Tax    int
	Age    int
}

type StringConv interface {
	ToString(int) string
	Taxes() int
	ToMap() map[string]any
}

func (s *Student) ToString(number int) string {

	return "| " + strconv.Itoa(number) + " | " + s.Name + " | " + " | " + strconv.Itoa(s.Stipend) + " | " + strconv.Itoa(s.Age) + " | " + strconv.Itoa(s.Stipend) + " | " + strconv.Itoa(s.Tax)
}

func (a *Adult) ToString(number int) string {
	return "| " + strconv.Itoa(number) + " | " + a.Name + " | " + strconv.Itoa(a.Age) + " | " + a.KraPin + " | " + strconv.Itoa(a.Salary) + " | " + strconv.Itoa(a.Tax)
}

func (s *Student) Taxes() int {
	s.Tax = 45

	return s.Tax
}

func (a *Adult) ToMap() map[string]any {
	adult := make(map[string]any)
	adult["name"] = a.Name
	return adult
}

func (s *Student) ToMap() map[string]any {

	student := make(map[string]any)
	student["name"] = s.Name
	return student
}

func (a *Adult) Taxes() int {
	a.Tax = 90
	return a.Tax
}

//def calc_tax(self, tax) -> int:
//income_tax = self.salary * (tax / 100)
//net_pay = self.salary - income_tax
//return net_pay
//
//def total_calc_tax(self) -> int:
//if 0 < self.age <= 45:
//return self.calc_tax(9)
//elif 46 <= self.age <= 65:
//return self.calc_tax(5)
//else:
//return self.calc_tax(3)
