package main

import "fmt"

type DatabaseConnection struct {
	name string
}

func Connect() *DatabaseConnection {
	println("Connecting to database: mongoDb")
	return &DatabaseConnection{name: "mongoDb"}
}

func main1() {
	connection := Connect()
	connection2 := Connect()
	connection3 := Connect()
	println("Connection 1:", connection.name)
	println("Connection 2:", connection2.name)
	fmt.Println(connection == connection2)
	fmt.Println(connection2 == connection3)
}

/*
Ab samjho kya hua:
- NewDatabase() first time call hua → object create hua.
- Second time call hua → another object create hua.
- Third time call hua → another object create hua.
Humare paas 3 different objects hain.

Lekin hum chahte hain ki chahe NewDatabase() 100 baar call ho, object ek hi create ho.

Yahan Singleton ka use aata hai.

*/
