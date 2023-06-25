package reader

import (
	"fmt"
	"os"
	"sync"

	"github.com/xuri/excelize/v2"
)

type Reader struct {
	file *excelize.File
}

type ReaderInterface interface {
	Read(*sync.WaitGroup, chan (map[string]string))
}

func New() (ReaderInterface, error) {

	f, err := excelize.OpenFile(os.Getenv("excel_name") + ".xlsx")
	if err != nil {
		return nil, fmt.Errorf("[Excel | New] err opening file: %v", err)
	}

	return &Reader{file: f}, nil

}
