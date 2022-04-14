package test

import (
	"tax_calculator/models"
	"testing"
)

func setUpAll() (student models.Student, adult models.Adult) {
	sdt := models.Student{
		Name:    "Eric Muli",
		Age:     24,
		Stipend: 2432,
		School:  "UoN",
	}

	adt := models.Adult{
		Name:   "Deno Muli",
		Age:    24,
		Salary: 2432,
		KraPin: "PESA",
	}
	return sdt, adt
}
func TestToStringConvAdult(t *testing.T) {
	_, adt := setUpAll()
	expectedAdult := "| 1 | Deno Muli | 24 | 2432 | 218 | 2214 | False | N/A | PESA"
	result := adt.ToString(1)
	if result != expectedAdult {
		t.Errorf("Test failed: Expected %s , received : %s", expectedAdult, result)
	}

}

func TestToStringConvStudent(t *testing.T) {
	sdt, _ := setUpAll()
	expectedAdult := "| 1 | Eric Muli | 24 | 2432 | 121 | 2311 | True | UoN | N/A"
	result := sdt.ToString(1)
	if result != expectedAdult {
		t.Errorf("Test failed: Expected %s , received : %s", expectedAdult, result)
	}
}

func TestToMapConvAdult(t *testing.T) {
	_, adt := setUpAll()
	expectedAdult := map[string]string{
		"name": "Deno Muli",
	}
	result := adt.ToMap()
	if result["name"] != expectedAdult["name"] {
		t.Errorf("Test failed: Expected %s , received : %s", expectedAdult["name"], result["name"])
	}

}

func TestToMapConvStudent(t *testing.T) {
	sdt, _ := setUpAll()
	expectedAdult := map[string]string{
		"name": "Eric Muli",
	}
	result := sdt.ToMap()
	if result["name"] != expectedAdult["name"] {
		t.Errorf("Test failed: Expected %s , received : %s", expectedAdult["name"], result["name"])
	}

}
func TestTaxesAdult(t *testing.T) {
	_, adt := setUpAll()
	expectedAdult := 218
	result := adt.Taxes(9)
	if result != expectedAdult {
		t.Errorf("Test failed: Expected %d , received : %d", expectedAdult, result)
	}

}

func TestTaxesStudent(t *testing.T) {
	sdt, _ := setUpAll()
	expectedAdult := 121
	result := sdt.Taxes(5)
	if result != expectedAdult {
		t.Errorf("Test failed: Expected %d , received : %d", expectedAdult, result)
	}
}
