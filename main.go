package main

import (
	"sync"

	"github.com/alejandrowaiz98/excel-reader/config"
	"github.com/alejandrowaiz98/excel-reader/injector"
	"github.com/alejandrowaiz98/excel-reader/reader"
	"github.com/joho/godotenv"
)

var wg sync.WaitGroup
var logger = config.GetLogger()

func main() {

	godotenv.Load(".env")

	wg.Add(2)

	ch := make(chan map[string]string)
	errCh := make(chan (error))

	reader, injector := setup()

	go reader.Read(&wg, ch)

	go injector.Inject(&wg, ch, errCh)

	for err := range errCh {

		logger.Err(err).Msg("")

	}

	wg.Wait()

}

func setup() (reader.ReaderInterface, injector.InjectorInterface) {

	reader, err := reader.New()

	if err != nil {
		logger.Fatal().Err(err).Msg("Reader error")
	}

	injector, err := injector.New()

	if err != nil {
		logger.Fatal().Err(err).Msg("Reader error")
	}

	return reader, injector

}
