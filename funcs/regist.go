package funcs

import (
	"fmt"
	"log"
)

func RegUsers(client map[string]int) {
	var (
		name    string
		balance int
	)
	fmt.Println("Введите имя:")
	_, err := fmt.Scanln(&name)
	if err != nil {
		log.Println("Регистрация не удалась")
		return
	}
	client[name] = balance
	log.Println("Вы успешно заригестрированы")
}
