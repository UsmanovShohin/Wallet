package main

import (
	"fmt"
	"log"
	"wallet/funcs"
)

var client = make(map[string]int)

func main() {
	for {
		fmt.Println("Выберите операцию ")
		fmt.Println("Регистрация = 1")
		fmt.Println("Проверка баланса= 2")
		fmt.Println("Список клиентов = 3")
		fmt.Println("Пополнение = 4")
		fmt.Println("Списание = 5")
		fmt.Println("Выйти = 6")
		var a string
		_, err := fmt.Scanln(&a)
		if err != nil {
			log.Println("Неправильная операция.Выберите заново")
			continue
		}
		switch a {
		case "1":
			funcs.RegUsers(client)
		case "2":
			funcs.FindBalance(client)
		case "3":
			funcs.ListClient(client)
		case "4":
			funcs.Replenishment(client)
		case "5":
			funcs.WriteOff(client)
		case "6":
			return
		default:
			fmt.Println("Выберите число от 1 до 6")
		}
	}

}
