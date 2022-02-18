package processor

import (
	"context"
	"time"

	"github.com/ussmith/cs/data"
)

type FastScanner struct {
	Name string
}

func (fs FastScanner) Scan(ctx context.Context, location string, results chan<- data.ScanStatus) {
	//fmt.Println("Starting a fast scanner")
	c := time.After(time.Millisecond * 300)
	<-c

	r := data.ScanStatus{
		ScannerName: fs.Name,
		Found:       false,
		Err:         nil,
	}

	results <- r
	//fmt.Println("Closing fast scanner")
	//close(results)
}
