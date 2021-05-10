package customer

import (
	"fmt"
)

type cusslice []Customer

var cusSlice cusslice

func update(id int, name string, sex string, age int, phone string, email string) bool {
	if id <= 0 || id > len(cusSlice) {
		return false
	}
	cusSlice[id-1].Name = name
	cusSlice[id-1].Sex = sex
	cusSlice[id-1].Age = age
	cusSlice[id-1].Phone = phone
	cusSlice[id-1].Email = email
	return true
}

func delCus(id int) bool {
	if id == -1 || id > len(cusSlice) {
		return false
	} else {
		cusSlice = append(cusSlice[:id-1], cusSlice[id:]...)
	}
	return true
}

func show() {
	for index, value := range cusSlice {
		fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t", index+1, value.Name, value.Sex, value.Age, value.Phone, value.Email)
		fmt.Println()
	}
}

func add(name string, sex string, age int, phone string, email string) bool {
	for _, value := range cusSlice {
		if name == value.Name {
			return false
		}
	}
	cus := newCustomer()
	cus.Name = name
	cus.Sex = sex
	cus.Age = age
	cus.Phone = phone
	cus.Email = email
	cusSlice = append(cusSlice, cus)
	return true
}
