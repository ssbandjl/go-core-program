package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type NewCustomerEvent struct {
	Name  string
	Phone string
	Email string
}

type NewCustomerEventMore struct {
	Name  string
	Phone string
	Email string
	QQ    int
}

func main() {
	newCustomer := NewCustomerEventMore{Name: "x", Phone: "082213909101", Email: "xyz@gmail.com", QQ: 491926119}
	convert(newCustomer)
}

func convert(event interface{}) {
	c := NewCustomerEvent{}
	mapstructure.Decode(event, &c)
	fmt.Printf("Event is: %v", c)
}
