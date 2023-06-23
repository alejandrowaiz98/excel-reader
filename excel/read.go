package excel

import (
	"fmt"
	"log"
	"os"
	"sync"
)

var mu sync.Mutex

func (e *Excel) Read(wg *sync.WaitGroup, ch chan (map[string]string)) {

	mu.Lock()
	defer mu.Unlock()
	defer close(ch)

	defer func() {
		// Close the spreadsheet.
		if err := e.file.Close(); err != nil {
			fmt.Println(err)
		}

		log.Println("File closed")
	}()

	rows, err := e.file.GetRows(os.Getenv("sheet_name"))

	if err != nil {
		fmt.Println(err)
		return
	}

	var columnNames []string

	for i, columns := range rows {

		log.Printf("column number %v", i)

		m := make(map[string]string)

		if i == 0 {
			log.Println("skipping")
			columnNames = columns
			continue
		}

		for i, columnValue := range columns {

			m[columnNames[i]] = columnValue

		}

		ch <- m
		m = nil

	}

	wg.Done()

}
