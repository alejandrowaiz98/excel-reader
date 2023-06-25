package reader

import (
	"log"
	"sync"

	config "github.com/alejandrowaiz98/excel-reader/config"
)

func (r *Reader) Read(wg *sync.WaitGroup, ch chan (map[string]string)) {

	logger := config.GetLogger()

	defer close(ch)

	defer func() {
		// Close the spreadsheet.
		if err := r.file.Close(); err != nil {
			logger.Error().Err(err).Msg("Err closing file")
		}

		log.Println("File closed")
	}()

	rows, err := r.file.GetRows("data")

	if err != nil {
		logger.Error().Err(err).Msg("Err getting rows")
		return
	}

	var columnNames []string

	for i, columns := range rows {

		m := make(map[string]string)

		if i == 0 {
			logger.Info().Msg("Getting columns names")
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
