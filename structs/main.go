package structs

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email       string
	phoneNumber int
	address     string
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func Run() {
	joe := &person{
		firstName: "jim", lastName: "stevens", contactInfo: contactInfo{
			email: "email@email.com", phoneNumber: 8433213491, address: "This is my address",
		}}

	fmt.Printf("%+v\n", joe)
	joe.updateName("STEVE")

	fmt.Printf("%+v", joe)
}
