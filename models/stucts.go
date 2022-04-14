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
	Taxes(int) int
	ToMap() map[string]any
}

func (s *Student) ToString(number int) string {

	return "| " + strconv.Itoa(number) + " | " + s.Name +
		" | " + strconv.Itoa(s.Age) +
		" | " + strconv.Itoa(s.Stipend) + " | " + strconv.Itoa(s.Taxes(5)) +
		" | " + strconv.Itoa(s.Stipend-s.Tax) + " | " + "True" + " | " + s.School + " | " + "N/A"
}

func (a *Adult) ToString(number int) string {
	return "| " + strconv.Itoa(number) + " | " + a.Name +
		" | " + strconv.Itoa(a.Age) +
		" | " + strconv.Itoa(a.Salary) + " | " + strconv.Itoa(a.Taxes(9)) +
		" | " + strconv.Itoa(a.Salary-a.Tax) + " | " + "False" + " | " + "N/A" + " | " + a.KraPin

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

func (a *Adult) Taxes(tax int) int {
	calc := a.Salary * tax / 100
	a.Tax = calc
	return calc
}
func (s *Student) Taxes(tax int) int {
	calc := s.Stipend * tax / 100
	s.Tax = calc
	return calc
}
