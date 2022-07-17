// Створити структури та реалізувати (імплементувати) методи:
// a) метод для типу Rectangle, який намалює в консолі прямокутник літерами “O” (як на прикладі в кутку),
//     метод повинен давати вибір - малювати вертикально, або горизонтально.+
//
// b) метод, що збільшить розмір Прямокутника в задану кількість разів+
//
// c) на Прямокутнику зробити метод, що порівнює його площу із площею іншого Прямокутника+
//
// d) створити тип Квадрат, зробити на Прямокутнику метод, що приймає Квадрат і відповідає
//     скільки таких Квадратів можна вписати у цей Прямокутник+
//
// e) Створити слайс Користувачів (структура як із лекції), вивести повне імʼя (fullName)
//     і нік користувачів зі слайса, які старші за 20 років, і імʼя (firstName) яких “John”.

package main

import "fmt"

type rectangle struct {
	name string

	x int
	y int
}

func (rt rectangle) compareWithSquere(sq square) {
	volRec := rt.x * rt.y
	volSq := sq.x * sq.x
	temp := 0
	for volRec > 0 {

		volRec -= volSq
		if volRec >= 0 {

			temp++
		}

	}
	fmt.Println("Squeres in rectangle =", temp)

}

func (rt rectangle) compareWithRec(rt2 rectangle) {
	var isBigger bool = false

	isBigger = (rt.x * rt.y) > (rt2.x * rt.y)

	if isBigger {
		fmt.Println(rt.name, "bigger than ", rt2.name)
	} else {
		fmt.Println(rt2.name, "bigger than ", rt.name)
	}

}

func (rt *rectangle) expandRectangle(step int) {
	rt.x = rt.x * step
	rt.y = rt.y * step
}

func (rt rectangle) writeRectangle(orient int) {
	x := rt.x
	y := rt.y

	fmt.Println(rt)

	if orient == 0 {

		for i := 0; i < x; i++ {
			for n := 0; n < y; n++ {
				print("0")

			}
			println("")
		}

	} else if orient == 1 {
		for i := 0; i < y; i++ {

			for n := 0; n < x; n++ {
				print("0")

			}
			println("")

		}

	} else {
		return
	}

}

type square struct {
	name string
	x    int
}

func main() {

	rec1 := rectangle{name: "rec1", x: 4, y: 2}
	rec2 := rectangle{name: "rec2", x: 8, y: 8}
	rec1.writeRectangle(0)
	rec2.writeRectangle(0)

	square1 := square{name: "square1", x: 2}

	//fmt.Println("Rectang name = ", rec1.name)

	//rec1.compareWithRec(rec2)
	//rec1.expandRectangle(10)
	//rec1.compareWithRec(rec2)

	rec1.compareWithSquere(square1)

}
