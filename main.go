package main

import (
	"bufio"
	"fmt"
	"github.com/JekaSmailik/mycmd/mymath"
	"os"
)

func main() {
	fmt.Println("Здравствуйте! выберите один из предложеных вариантов: \n" +
		"1. Для получения точного времения введите 'time' \n" +
		"2. Для входа в калькулятор введите 'cal' \n" +
		"3. Для вывода массива введите 'mas' \n" +
		"4. Для выхода введите 'exit' \n" +
		"5. Для спросмотра истории(вывода введенных команд) введите 'his'")

	defer fmt.Println("Программа завершенa")

	// Присвоение переменной для дальнейшего действия ввода команд
	var x string = " "

	// Присвоение переменной тип bufio (для обнуления данных, воизбежании не верных вводов)
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Scanf("%s\n", &x)

		// создание файла и сохранение вводимых команд
		mymath.MyWrite(x)

		switch x {

		// вывод времени, даты и дня недели в данный момент
		case "time":

			fmt.Println(mymath.MyTime(), "\n ")

			// Запуск программы калькулятор
		case "cal":
			fmt.Println("Введите два числа через пробел")

			var number1 float64 = 0
			var number2 float64 = 0

			for {
				fmt.Scanln(&number1, &number2)

				if number1 == 0 || number2 == 0 {
					fmt.Println("для корректного вввода, необходимо вводить исключительно числа!")
					fmt.Println(number1, number2)

					// Очистка Stdin (обнуляет прошлые вводимые данные, если они не соответствовали условиям)
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

			// Запуск программы рассчета данных массива
		case "mas":

			fmt.Println("Укажите длину желаемого массива")

			var nam int

			fmt.Scanf("%v", &nam)

			fmt.Println("Получившийся массив", mymath.MyMassiv(nam))

			// Вывод истории введенных команд
		case "his":
			mymath.MyHistory()

			fmt.Println("Для удаления истории введите 'clear'")

		// Завершение работы всей программы
		case "exit":
			fmt.Println("Good bay!")
			return

		// Удаление истории введенных команд
		case "clear":
			mymath.MyClear()

			fmt.Println("История очищена!")

		// Повторение цикла программы
		case x:
			r.ReadBytes('\n')
			fmt.Println("Введите одну из команд: time, cal, mas, his или exit")
		}
	}
}
