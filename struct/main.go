package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Client struct {
	Name  string
	Email string
	CPF   int
}

type InternacionalClient struct {
	Client
	Country string `json:"country"`
}

func (c Client) Show() {
	fmt.Println("Exibindo cliente pelo método Show:", c.Name, c.Email, c.CPF)
}

// func (c InternacionalClient) Show() string {
// 	return fmt.Sprintf("Name: %v \nEmail: %v \nCPF: %v \nCountry: %v \n", c.Name, c.Email, c.CPF, c.Country)
// }

func main() {
	client := Client{
		Name:  "John Doe1",
		Email: "j@j.com",
		CPF:   123456789,
	}
	client.Show()

	client2 := Client{"John Doe2", "j@j.com", 12121}

	client2.Show()

	client3 := InternacionalClient{
		Client: Client{
			Name:  "John Doe3",
			Email: "j@j.com",
			CPF:   123456789,
		},
		Country: "US",
	}

	client3.Show()

	clientJson, err := json.Marshal(client3)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(clientJson))

	jsonClient := `{"Name":"John Doe4","Email":"j@j.com","CPF":123456789,"Country":"US"}`
	client4 := InternacionalClient{}
	json.Unmarshal([]byte(jsonClient), &client4)
	fmt.Println(client4)
}
