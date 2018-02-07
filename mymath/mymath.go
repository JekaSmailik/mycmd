package mymath

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Путь места нахождения программы
var path = "../mycmd/mymath/historyFile.txt"

func MyTime() string {
	timeNow := time.Now()
	mytime := timeNow.Format("Дата 02.01.2006\nВремя 15:04\n")

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

// Сохранение всех вводов командной строки в файле по ссылке
func MyWrite(x string) {
	// открытие файла с возможностью добавления информации
	historyFile, err := os.OpenFile(path, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0777)

	if err != nil {
		// здесь перехватывается ошибка
		// создание нового файла
		historyFile, err = os.Create(path)
		if err != nil {
			fmt.Println(err)

			return
		}
	}
	defer historyFile.Close()

	// Добавление к выводу команд нумерацию в порядке возростания
	var countLine int
	f := bufio.NewReader(historyFile)
	for {
		_, err := f.ReadString('\n')
		countLine++
		if err != nil {
			break
		}
	}

	// Запись информации в файл
	byteSlice := []byte(strconv.FormatInt(int64(countLine), 10) + " :" + x + "\n")
	_, err = historyFile.Write(byteSlice)
	if err != nil {
		fmt.Println(err)
	}

	//	n, err := historyFile.WriteString(x)
	//	fmt.Println(err, n)
}

// Выводит историю вводимых команд
func MyHistory() {

	fmt.Println("Ниже предоставленна история вводимых команд")
	//fmt.Println("Общее колличество введенных команд", len())

	historyFile, err := os.Open(path)
	if err != nil {
		// здесь перехватывается ошибка
		fmt.Println(err)

		return
	}
	defer historyFile.Close()

	// Получение размера файла
	stat, err := historyFile.Stat()
	if err != nil {
		fmt.Println(err)

		return
	}

	// Чтение файла
	readFile := make([]byte, stat.Size())
	_, err = historyFile.Read(readFile)
	if err != nil {
		fmt.Println(err)

		return
	}
	// Вывод данных файла
	speakFile := string(readFile)
	fmt.Println(speakFile)

	/*
		strings.NewReader(path)
			fmt.Printf("%v", historyFile)
	*/
}

// Очищает файл по ссылке path
func MyClear() {
	err := os.Remove(path)
	if err != nil {
		fmt.Println(err)
	}
}
