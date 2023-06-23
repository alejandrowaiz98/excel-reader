package excel

import (
	"fmt"
	"os"
	"sync"

	"github.com/xuri/excelize/v2"
)

type Excel struct {
	file *excelize.File
}

type ExcelInterface interface {
	Read(*sync.WaitGroup, chan (map[string]string))
}

func New() (ExcelInterface, error) {

	f, err := excelize.OpenFile(os.Getenv("excel_name") + ".xlsx")
	if err != nil {
		return nil, fmt.Errorf("[Excel | New] err opening file: %v", err)
	}

	return &Excel{file: f}, nil

}
