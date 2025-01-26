package main

import (
	"fmt"
	"time"
)

func main() {
	mine := [10]string{"rock", "ore", "ore", "ore", "rock", "rock", "ore", "rock", "rock", "ore"}
	oreChannel := make(chan string)
	minedOreChannel := make(chan string)

	//finder
	go func(m [10]string) {
		for _, item := range m {
			if item == "ore" {
				oreChannel <- item
			}
		}
	}(mine)

	//miner
	go func() {
		for i := 0; i < len(mine); i++ {
			foundOre := <-oreChannel
			fmt.Println("from finder:", foundOre)
			minedOreChannel <- "minedOre"
		}
	}()

	//smelter
	go func() {
		for i := 0; i < len(mine); i++ {
			minedOre := <-minedOreChannel
			fmt.Println("from miner:", minedOre)
			fmt.Println("from smelter: ore is smelted")
		}
	}()

	<-time.After(5 * time.Second)
}
