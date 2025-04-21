package main

import "fmt"

// В цикле спрашиваем у пользователя ввод транзакций: -10, 10, 40.5
// Добавлять каждую в массив транзакций
// Выводить массив транзакций
// !Вывести сумму баланса в консоль

func main() {
	transactions := []float64{} // массив транзакций
	for {
		transaction := scanTransaction()
		if transaction == 0 {
			break
		}
		transactions = append(transactions, transaction)
	}
	printTransactions(transactions)
}

func scanTransaction() float64 {
	var transaction float64
	fmt.Print("Введите транзакцию (n - выход): ")
	fmt.Scanln(&transaction)
	return transaction

}

func calculateBalance(transactions []float64) float64 {
	balance := 0.0
	for _, value := range transactions {
		balance += value
	}
	return balance
}

func printTransactions(transactions []float64) {
	balance := calculateBalance(transactions)
	fmt.Printf("Сумма баланса: %.2f\n", balance)
}
