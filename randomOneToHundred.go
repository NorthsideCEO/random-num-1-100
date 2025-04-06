package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	userSGuess := ""

	for {
		randomNum := rand.Intn(100) + 1

		fmt.Println("Попробуйте угадать число от 1 до 100 включительно")

		for i := 10; i >= 0; i-- {
			for {
				userSGuess, _ = scanner.ReadString('\n')
				userSGuess = strings.TrimSpace(strings.ToLower(userSGuess))
				_, err := strconv.Atoi(userSGuess)
				if err != nil {
					fmt.Println("Введите натуральное целое число: ")
					continue
				}
				break
			}
			userSGuess, _ := strconv.Atoi(userSGuess)
			if userSGuess < randomNum {
				fmt.Println("Ой. Ваша догадка ниже загаданного числа! \nУ вас осталось", i, "попыток!")
				continue
			} else if userSGuess > randomNum {
				fmt.Println("Ой. Ваша догадка выше загаданного числа! \nУ вас осталось", i, "попыток!")
				continue
			} else {
				fmt.Println("Вы угадали!\nХотите продолжить?")
				for {
					userChoice, _ := scanner.ReadString('\n')
					userChoice = strings.ToLower(strings.TrimSpace(userChoice))
					if userChoice == "да" {
						break
					} else if userChoice == "нет" {
						fmt.Println("Спасибо за игру!\nПрограмма автоматически закроется через 5 секунд...")
						time.Sleep(4 * time.Second)
						log.Fatal()
					} else {
						fmt.Println("Введите да или нет")
						continue
					}
				}
			}
			break
		}

		userSGuess = strings.TrimSpace(strings.ToLower(userSGuess))
		userSGuess, _ := strconv.Atoi(userSGuess)
		if userSGuess != randomNum {
			fmt.Println("Вы ни разу не угадали, загаданным числом было:", randomNum, "\nХотите продолжить? да/нет")
			for {
				userChoice, _ := scanner.ReadString('\n')
				userChoice = strings.ToLower(strings.TrimSpace(userChoice))
				if userChoice == "да" {
					break
				} else if userChoice == "нет" {
					fmt.Println("Спасибо за игру!\nПрограмма автоматически закроется через 5 секунд...")
					time.Sleep(4 * time.Second)
					log.Fatal()
				} else {
					fmt.Println("Выберите да или нет: ")
					continue
				}
			}
			continue
		}
		continue
	}
}
