package main

func main() {
	/*
		// Попытка обратиться к несуществующему индексу
		data := [3]int{1, 2, 3}
		idx := 4
		fmt.Println(data[idx])	// panic
		fmt.Println(data[4])	// compilation error
	*/

	/*
		// Попытка обратиться к отрицательному индексу
		data := [3]int{1, 2, 3}
		idx := -1
		fmt.Println(data[idx])	// panic
		fmt.Println(data[-1])	// compilation error
	*/

	/*
		// Длина и емкость пустого массива
		data := [10]int{}
		fmt.Println(len(data))	// 10
		fmt.Println(cap(data))	// 10
	*/

	/*
		// Сравнение массивов
		first := [...]int{1, 2, 3}
		second := [...}int{1, 2, 3}
		fmt.Println(first == second)	// true
		fmt.Println(first != second)	// false
		// Операции больше, меньше, больше или равно, меньше или равно не поддерживаются
	*/

	/*
		// Пустой массив
		var data [10]byte
		fmt.Println(unsafe.Sizeof(data))	// 10
		var data1 [0]byte
		fmt.Println(unsafe.Sizeof(data1))	// 0
	*/

	/*
	   // Отрицательная длина невозможна
	   var data [-1]int	// compilation error
	*/

	/*
		   // Задание длины массива из переменной невозожно, только константа
			len1 := 100
			var data1 [len1]int	// compilation error
			const len2 = 100
			var data [len2]int	// ok
	*/

	// make и append с массивами не работают
}
