package customer

type Customer struct {
	Name  string
	Sex   string
	Age   int
	Phone string
	Email string
}

func newCustomer() (cus Customer) {
	cus = Customer{
		Name:  "",
		Sex:   "",
		Age:   0,
		Phone: "",
		Email: "",
	}
	return
}
