package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// ==========================================================
// 1. ITarget Interface (Dono Cases isi rule ko follow karenge)
// ==========================================================
type ITarget interface {
	Request() (string, error)
}

// ==========================================================
// CASE 1: Normal JSON Service (Koi Adapter nahi chahiye)
// ==========================================================
type NormalJSONService struct{}

func (n *NormalJSONService) Request() (string, error) {
	// Yeh modern API hai, pehle se hi JSON return karti hai
	return `{
  "id": 101,
  "name": "Neha Gupta",
  "department": "Frontend Engineering"
}`, nil
}

// ==========================================================
// CASE 2: Purana XML API (Adaptee) + Adapter
// ==========================================================
type Adaptee struct{}

func (a *Adaptee) SpecificRequest() string {
	// Yeh purani API hai, sirf XML return karti hai
	return `
	<employee>
		<id>501</id>
		<name>Rahul Verma</name>
		<department>Cloud Architecture</department>
	</employee>`
}

type Employee struct {
	ID         int    `xml:"id" json:"id"`
	Name       string `xml:"name" json:"name"`
	Department string `xml:"department" json:"department"`
}

type Adapter struct {
	adaptee *Adaptee
}

func (ad *Adapter) Request() (string, error) {
	// Step 1: Internally XML API ko call kiya
	rawXML := ad.adaptee.SpecificRequest()

	// Step 2: XML ko JSON mein convert kiya
	var emp Employee
	if err := xml.Unmarshal([]byte(rawXML), &emp); err != nil {
		return "", err
	}
	jsonBytes, err := json.MarshalIndent(emp, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

// ==========================================================
// CLIENT (Ise sirf ITarget se matlab hai)
// ==========================================================
type Client struct{}

func (c *Client) Execute(target ITarget) {
	data, err := target.Request()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(data)
	fmt.Println("-----------------------------------")
}

// ==========================================================
// MAIN: Dekhiye Call Kaise Hota Hai Aur Kya Hota Hai
// ==========================================================
func main() {
	client := &Client{}

	// --- CASE 1: Jab data normal JSON API se aa raha ho ---
	fmt.Println("1. Normal JSON Service ko call kar rahe hain:")
	normalService := &NormalJSONService{}
	client.Execute(normalService) // Seedha pass kar diya

	// --- CASE 2: Jab data XML API se aa raha ho ---
	fmt.Println("2. XML API ko Adapter ke through call kar rahe hain:")
	xmlAPI := &Adaptee{}
	xmlAdapter := &Adapter{adaptee: xmlAPI}
	client.Execute(xmlAdapter) // Adapter mein wrap karke pass kiya
}

/*

Call Karne Par Kya Hota Hai? (Step-by-Step)

1. Jab client.Execute(normalService) call hota hai:
target.Request() seedha NormalJSONService ka method chalata hai aur turant JSON print ho jaata hai. Koi conversion nahi hota.

2. Jab client.Execute(xmlAdapter) call hota hai:
Client phir se wahi target.Request() chalata hai, lekin is baar Adapter ka method chalta hai. Adapter chupchap peeche se xmlAPI.SpecificRequest() ko call karke XML mangwata hai, use JSON mein badalta hai, aur Client ko bilkul waisa hi JSON pakda deta hai jaisa Case 1 mein mila tha.
*/
