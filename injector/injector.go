package injector

import (
	"context"
	"sync"

	"cloud.google.com/go/firestore"
	"github.com/alejandrowaiz98/excel-reader/config"
	"google.golang.org/api/option"
)

type Injector struct {
	client *firestore.Client
}

type InjectorInterface interface {
	Inject(wg *sync.WaitGroup, ch chan (map[string]string), errCh chan (error))
}

func New() (InjectorInterface, error) {

	logger := config.GetLogger()

	ctx := context.Background()

	client, err := firestore.NewClient(ctx, "pia-playground", option.WithCredentialsFile("service-account.json"))

	if err != nil {
		logger.Error().Err(err).Msg("Err creating firestore client")
		return nil, err

	}

	return &Injector{client: client}, nil
}
