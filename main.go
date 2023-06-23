package main

import (
	"log"
	"sync"

	"github.com/alejandrowaiz98/excel-reader/excel"
	"github.com/joho/godotenv"
)

var wg sync.WaitGroup

func main() {

	godotenv.Load(".env")

	wg.Add(1)

	ch := make(chan map[string]string)

	excel, err := excel.New()

	if err != nil {
		panic(err)
	}

	go excel.Read(&wg, ch)

	for msg := range ch {

		log.Printf("To store: %v", msg)

	}

	wg.Wait()

}
