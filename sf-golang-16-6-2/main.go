package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"sf-golang-16-6-2/bank"
	"strconv"
	"sync"
	"time"
)

const (
	cmdBalance    = "balance"
	cmdDeposit    = "deposit"
	cmdWithdrawal = "withdrawal"
	cmdExit       = "exit"
)

var stdin = bufio.NewScanner(os.Stdin)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var wg sync.WaitGroup
	c := bank.NewClient()

	// Горутины параллельно пополняют и списывают
	// средства со счета случайным образом
	wg.Add(2)
	go seedDeposit(c, &wg)
	go seedWithdrawal(c, &wg)

	// Перед тем как вызвать обработчик
	// пользовательских команд, дожидаемся
	// завершения горутин
	wg.Wait()

	// Обработчик пользовательских команд
	cmdLoop(c)
}

// seedDeposit - создает 10 горутин, которые
// с задержкой 0.5..1 сек вызывают Deposit со случайным значением 1..10
func seedDeposit(c *bank.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			randSleep(time.Second/2, time.Second)
			c.Deposit(1 + rand.Intn(10))
		}()
	}
}

// seedWithdrawal - создает 5 горутин, которые
// с задержкой 0.5..1 сек вызывают Withdrawal со случайным значением 1..5.
// Если списание невозможнно, выводится сообщение об ошибке
func seedWithdrawal(c *bank.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			randSleep(time.Second/2, time.Second)
			if err := c.Withdrawal(1 + rand.Intn(5)); err != nil {
				printError(err)
			}
		}()
	}
}

// cmdLoop - цикл обработки пользовательских команд
func cmdLoop(c *bank.Client) {
	for {
		fmt.Print("Command: ")
		stdin.Scan()
		cmd := stdin.Text()

		switch cmd {

		// Узнать баланс
		case cmdBalance:
			fmt.Println(c.Balance())

		// Пополнить
		case cmdDeposit:
			amount, err := readAmount()
			if err != nil {
				printError(err)
				continue
			}
			c.Deposit(amount)

		// Списать
		case cmdWithdrawal:
			amount, err := readAmount()
			if err != nil {
				printError(err)
				continue
			}
			if err = c.Withdrawal(amount); err != nil {
				printError(err)
			}

		// Завершить обработчик
		case cmdExit:
			return

		// Неизвестная команда
		default:
			fmt.Println("Unsupported command. You can use commands: balance, deposit, withdrawal, exit")
			break
		}
	}
}

// readAmount - запрашивает ввод amount для Withdrawal и Deposit.
// Допускается ввод целых неотрицательных чисел, иначе возвращает ошибку
func readAmount() (int, error) {
	fmt.Print("Amount: ")
	stdin.Scan()
	amount, err := strconv.Atoi(stdin.Text())
	if err != nil || amount < 0 {
		return 0, errors.New("invalid amount")
	}
	return amount, nil
}

// randSleep - приостанавливает вызвавшую горутину на случайное время
// в диапазоне от min до max. Вызывает панику, если max <= min
func randSleep(min, max time.Duration) {
	time.Sleep(min + time.Duration(rand.Int63n(int64(max-min))))
}

func printError(err error) {
	fmt.Println("Error:", err.Error())
}
