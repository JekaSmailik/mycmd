package mymath

import (
	"errors"
	"fmt"
	"time"
)

func MyTime() string {
	timeNow := time.Now()
	mytime := timeNow.Format("Дата 2.01.2006" +
		"\nВремя 15:04\n")

	var myday string

	switch timeNow.Weekday().String() {
	case "Monday":
		myday = "понедельник"
	case "Tuesday":
		myday = "вторник"
	case "Wednesday":
		myday = "среда"
	case "Thursday":
		myday = "четверг"
	case "Friday":
		myday = "пятница"
	case "Saturday":
		myday = "суббота"
	case "Sunday":
		myday = "воскресение"
	}

	fmt.Println("День недели", myday)

	return mytime
}

func MySumma(number1 float64, number2 float64) (float64, error) {
	var calculator float64
	var act string

	fmt.Println("Укажите действие для введенных чисел: '+', '-', '/', '*'")

	fmt.Scanf("%s", &act)

	switch act {
	case "+":
		calculator = number1 + number2
	case "-":
		calculator = number1 - number2
	case "/":
		calculator = number1 / number2
	case "*":
		calculator = number1 * number2
	case act:
		return 0, errors.New("Вы вернулись в начало программы \n" +
			"Нужно было выбрать действие из четырех вариантов: '+', '-', '/', '*' \n")
	}

	return calculator, nil
}

func MyMassiv(nam int) []int {
	var massiv []int = make([]int, 0, nam)
	var uslovie string = " "

	fmt.Println("Если Вы хотите заполнить массив в ручную введите 'hand' \n" +
		"если хотите чтоб массив заполнился автоматически то нажмите enter")

	fmt.Scanf("%s", &uslovie)

	switch uslovie {
	// Заполнение массива в ручную, без ограничения длины
	case "hand":
		fmt.Println("Введите переменные массива, \n " +
			"для завершения ввода введите '0' или любую букву")

		var mas = -1
		// Баг!!! при вводе буквы, повторяет предыдущий ввод
		for {
			fmt.Scanf("%v", &mas)

			if mas == 0 {
				break
			}

			massiv = append(massiv, mas)

			fmt.Println(massiv)
		}
	// Перебор массива и присвоение ему данных от 0 до длины массива.
	default:
		fmt.Println("В какой последоватальности вы хотите заполнить массив? \n" +
			"если по возростанию введите '+', если на убывание то '-' \n" +
			"при вводе любой другой команды будет выведен пустой массив")

		var upDown string = " "

		fmt.Scanf("%s", &upDown)

		switch upDown {
		case "+":
			for i := 0; i < nam; i++ {
				massiv = append(massiv, (nam-1)-i)
			}
		case "-":
			for i := nam; i > 0; i-- {
				massiv = append(massiv, nam-i)
			}
		case upDown:
			massiv = append(massiv)
		}
	}

	return massiv
}
