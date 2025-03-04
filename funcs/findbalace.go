package funcs

import (
	"fmt"
	"log"
)

func FindBalance(client map[string]int) {
	if len(client) == 0 {
		fmt.Println("Нет зарегистрированных клиентов.")
		return
	}

	fmt.Println("Введите имя")
	var ScanName string
	_, err := fmt.Scanln(&ScanName)
	if err != nil {
		log.Println("Проверка баланса не удалась")
		return
	}
	
	for name, balance := range client {
		if name == ScanName {
			fmt.Println(name, "= ", balance)
		} else {
			fmt.Println("Клиент не найден")
		}
	}
}
