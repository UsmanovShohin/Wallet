package funcs

import "fmt"

func ListClient(client map[string]int) {
	if len(client) == 0 {
		fmt.Println("Нет зарегистрированных клиентов.")
		return
	}
	fmt.Println("Список клиентов:")
	for name, balance := range client {
		fmt.Printf("%s: %d\n", name, balance)
	}
}
