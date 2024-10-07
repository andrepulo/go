package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10) // Указываем количество горутин

	for i := 0; i < 10; i++ {
		i := i // Захватываем текущее значение i
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}

	wg.Wait() // Ожидаем завершения всех горутин
}
