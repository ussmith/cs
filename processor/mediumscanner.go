package processor

import (
	"context"
	"time"

	"github.com/ussmith/cs/data"
)

type MediumScanner struct {
	Name string
}

func (fs MediumScanner) Scan(ctx context.Context, location string, results chan<- data.ScanStatus) {
	//fmt.Println("Starting a medium scanner")
	c := time.After(time.Millisecond * 1200)
	<-c

	var foundVals []string
	foundVals = append(foundVals, "v1")
	foundVals = append(foundVals, "v2")
	foundVals = append(foundVals, "v3")
	r := data.ScanStatus{
		ScannerName: fs.Name,
		Found:       true,
		Viruses:     foundVals,
		Err:         nil,
	}

	results <- r
	//fmt.Println("Closing medium scanner")
	//close(results)
}
