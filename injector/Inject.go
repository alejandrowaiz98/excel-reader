package injector

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/alejandrowaiz98/excel-reader/config"
)

var mu sync.Mutex

func (i *Injector) Inject(wg *sync.WaitGroup, ch chan (map[string]string), errCh chan (error)) {

	logger := config.GetLogger()

	mu.Lock()
	defer mu.Unlock()

	defer close(ch)
	defer close(errCh)

	ctx := context.Background()

	for msg := range ch {

		_, _, err := i.client.Collection(os.Getenv("destiny_collection")).Add(ctx, msg)

		if err != nil {

			logger.Error().Err(err).Msg(msg["ID"])

			errCh <- fmt.Errorf("Err creating docRef with document %v", msg["ID"])

		}

	}

	wg.Done()

}
