package main

import "fmt"

func updStorage(balance map[string]int, buys []string) string {
	for _, item := range buys {
		if inStock, ok := balance[item]; ok {
			if inStock > 0 {
				balance[item]--
			} else {
				return fmt.Sprintf("Товара '%s' нет в наличии", item)
			}
		} else {
			return fmt.Sprintf("Товар '%s' отсутствует на складе", item)
		}
	}

	return ""
}

func main() {
	balance := map[string]int{
		"Pos1":  2,
		"Pos2":  5,
		"Pos3":  0,
		"Pos4":  4,
		"Pos5":  8,
		"Pos6":  9,
		"Pos7":  1,
		"Pos8":  0,
		"Pos9":  0,
		"Pos10": 4,
		"Pos11": 6,
		"Pos12": 3,
	}
	buys := []string{"Pos1", "Pos4", "Pos5", "Pos10", "Pos11", "Pos12"}
	errs := updStorage(balance, buys)

	if errs != "" {
		fmt.Println(errs)
	} else {
		fmt.Println("Обновлено", balance)
	}
}
