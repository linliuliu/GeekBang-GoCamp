package main

import (

	"fmt"
)

func main() {
	conn := ConnectDb()
	fmt.Println(conn)
	CreateMoment(conn)
	DeleteMoment(conn)
	GetMoment(conn)
}
