package funcs

import (
	"fmt"
	"log"
)

func WriteOff(client map[string]int) {
	if len(client) == 0 {
		fmt.Println("Нет зарегистрированных клиентов.")
		return
	}

	fmt.Println("Введите имя:")
	var ScanName string
	_, err := fmt.Scanln(&ScanName)
	if err != nil {
		log.Println("Ошибка ввода имени")
		return
	}

	balance, exists := client[ScanName]
	if !exists {
		fmt.Println("Клиент не найден")
		return
	}
	fmt.Println("Текущий баланс:", ScanName, "=", balance)
	fmt.Println("Введите сумму списания:")

	var Replen int
	_, err = fmt.Scanln(&Replen)
	if err != nil || Replen <= 0 {
		log.Println("Ошибка ввода суммы списания")
		return
	}
	if Replen > balance {
		fmt.Println("Недостаточно средств")
		return
	}

	client[ScanName] -= Replen
	fmt.Println("Списание было на сумму =", Replen)
	fmt.Println("Новый баланс:", ScanName, "=", client[ScanName])
}
