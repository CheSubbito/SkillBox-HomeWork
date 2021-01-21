package main

import (
	"fmt"
	"os"
)

func main() {

}

/*
Вывод в терминал с клавиатуры
*/
func keyScan() {
	var a int
	fmt.Fscan(os.Stdin, &a)
	fmt.Println(a)
}

func homeWork1() {

	//Вариант обычный
	fmt.Println("Hello, I`m Andrey")
	fmt.Println("I`m 27 years old")
	fmt.Println("Looking forward to the lesson!")

	//Вариант через fmt.Printf()

	name := "Andrey"
	age := 30

	fmt.Printf("Hello, I`m %v\nI`m %v years old\nLooking forward to the lesson!", name, age)

	//Вариант через fmt.Sprintf()

	name = "Andrey"
	age = 30
	text := fmt.Sprintf("Hello, I`m %v\nI`m %v years old\nLooking forward to the lesson!", name, age)

	fmt.Println(text)
}

func lesson2_2() {

	var apartamentCost = 42678434687 //стоимость квартиры (можно не писать int)
	var bootsSize int = 41           //размер обуми
	var animalsZooNow int = 82       //Колличество животных в зоопарке в данный момент
	var newProductPrice int = 65     //цена для нового товара
	fmt.Println(apartamentCost)
	fmt.Println(bootsSize)
	fmt.Println(animalsZooNow)
	fmt.Println(newProductPrice)
}

func lesson2_3() {

	var circle int = 4
	var speed int = 358
	var engine int = 254
	var wheels int = 93
	var steringWheel int = 49
	var wind int = 21
	var rain int = 17

	fmt.Println("===============")
	fmt.Println("Супер гонки. Круг", circle)
	fmt.Println("===================")
	fmt.Print("Шумахер (", speed, ")\n")
	fmt.Println("===================")
	fmt.Println("Водитель: Шумахер")
	fmt.Println("Скорость:", speed)
	fmt.Println("-------------------")
	fmt.Println("Оснащение")
	fmt.Print("Двигатель: +", engine, "\n")
	fmt.Print("Колеса: +", wheels, "\n")
	fmt.Print("Руль: +", steringWheel, "\n")
	fmt.Println("-------------------")
	fmt.Println("Действия плохой погоды")
	fmt.Print("Ветер: -", wind, "\n")
	fmt.Print("Дождь: -", rain, "\n")
}

func lesson2_4() {

	a := 7
	b := 13
	c := 2
	d := 3
	result := (a + b) / (c + d)
	fmt.Println(result)
}

/*
Тут константа
*/
func homeWork2_1() {
	//const lap2 int = 4 так пишется константа
	lap := 4
	engine := 254
	wheels := 93
	steeringWheel := 49
	wind := 21
	rain := 17
	speed := engine + wheels + steeringWheel - rain - wind

	fmt.Println("===================")
	fmt.Print("Супер гонки. Круг ", lap, "\n")
	fmt.Println("===================")
	fmt.Print("Шумахер (", speed, ")\n")
	fmt.Println("===================")
	fmt.Println("Водитель: Шумахер")
	fmt.Print("Скорость: ", speed, "\n")
	fmt.Println("-------------------")
	fmt.Println("Оснащение")
	fmt.Print("Двигатель: +", engine, "\n")
	fmt.Print("Колеса: +", wheels, "\n")
	fmt.Print("Руль: +", steeringWheel, "\n")
	fmt.Println("-------------------")
	fmt.Println("Действия плохой погоды")
	fmt.Print("Ветер: -", wind, "\n")
	fmt.Print("Дождь: -", rain, "\n")
}

func homeWork2_2() {

	product := 6400
	delivery := 350
	discount := 700
	totalCost := product + delivery - discount

	fmt.Print("Полная стоимость товара с учётом доставки и скидки :", totalCost, ".")
}

func homeWork2_3() {

	workShiftMin := 480
	ordetTime := 2
	orderAssembly := 4
	ClientsServed := workShiftMin / (ordetTime + orderAssembly)

	fmt.Println("За смену длинной", workShiftMin, "минут кассир успеет обслужить", ClientsServed, "клиентов.")
}

/*
Ввод переменных с клавиатуры
Fscan
*/
func homeWork2_4() {
	fullCostRepair := 0       //4000000
	numberOfEntrances := 0    //10
	apartmentsInEntrance := 0 //40
	fmt.Print("Введите общую стоимость: ")
	fmt.Fscan(os.Stdin, &fullCostRepair)
	fmt.Print("Введите колличество подъездов: ")
	fmt.Fscan(os.Stdin, &numberOfEntrances)
	fmt.Print("Введите колличество квартир в подъезде: ")
	fmt.Fscan(os.Stdin, &apartmentsInEntrance)
	apartmentsInHouse := numberOfEntrances * apartmentsInEntrance
	priceForEach := fullCostRepair / apartmentsInHouse

	fmt.Println("Каждая квартира должна заплатить по", priceForEach, "руб.")
}
