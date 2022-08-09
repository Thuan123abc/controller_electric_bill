package main

import (
	"controller_electric_bill/database"
	"fmt"
)

func main() {
	Cus1 := database.Customer{
		1,
		"Thuan",
		"Hiep Hoa, Bac Giang",
		12,
	}
	Bil1 := database.ElectricBill{
		Cus1,
		11,
		12,
		1223,
		1334,
	}
	db := database.NewDB()
	CustRepo1 := database.NewCustomerRepo(db)
	err := CustRepo1.CreateCustomer(Cus1)
	if err != nil {
		return
	}
	fmt.Println(Bil1.Amount())
}
