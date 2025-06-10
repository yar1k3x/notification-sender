package main

import (
	"NotificationSender/db"
	"NotificationSender/server"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	err := db.InitDB("root", "vdySqAwCIwMHUfdUyqaQlBOBlCrZovdD", "centerbeam.proxy.rlwy.net", "railway")

	if err != nil {
		log.Println("Не удалось подключиться к базе данных")
	}

	log.Println("БД успешно подключена")
	server.StartGRPCServer()
}
