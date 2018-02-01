package main

import (
	"bufio"
	"fmt"
	"github.com/JekaSmailik/mycmd/mymath"
	"os"
)

func main() {
	fmt.Println("Здравствуйте! выберите один из предложеных вариантов: \n" +
		"1. Для получения точного времения введите time \n" +
		"2. Для входа в калькулятор введите cal \n" +
		"3. Для вывода массива введите mas \n" +
		"4. Для выхода введите exit")

	defer fmt.Println("Программа завершенa")

	// Ввод команды для дальнейшего действия
	var x string = " "

	// Присвоение переменной тип bufio
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Scanf("%s\n", &x)
		switch x {

		case "time":

			fmt.Println(mymath.MyTime(), "\n ")

		case "cal":
			fmt.Println("Введите два числа через пробел")

			var number1 float64 = 0
			var number2 float64 = 0

			for {
				fmt.Scanln(&number1, &number2)

				if number1 == 0 || number2 == 0 {
					fmt.Println("для корректного вввода, необходимо вводить исключительно числа!")
					fmt.Println(number1, number2)

					// Очистка Stdin
					r.ReadBytes('\n')

				} else {
					break
				}
			}

			calculator, err := mymath.MySumma(number1, number2)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Значение округлено до одного знака после запятой")
				fmt.Printf("Результат ваших действий равен %.1f \n", calculator)
			}

		case "mas":

			fmt.Println("Укажите длину желаемого массива")

			var nam int

			fmt.Scanf("%v", &nam)

			fmt.Println("Получившийся массив", mymath.MyMassiv(nam))

		case "exit":
			fmt.Println("Good bay!")
			return

		case x:
			r.ReadBytes('\n')
			fmt.Println("Введите одну из команд: time, sum, mas или exit")

		}
	}
}
