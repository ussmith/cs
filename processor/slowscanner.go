package processor

import (
	"context"
	"errors"
	"time"

	"github.com/ussmith/cs/data"
)

type SlowScanner struct {
	Name string
}

func (fs SlowScanner) Scan(ctx context.Context, location string, results chan<- data.ScanStatus) {
	process := true
	for process == true {
		select {
		case <-time.After(time.Millisecond * 3000):
			//fmt.Println("3 seconds gone")
			r := data.ScanStatus{
				ScannerName: "SlowScanner",
				Found:       false,
				Err:         nil,
			}

			results <- r
			process = false
			continue

		case <-ctx.Done():
			//fmt.Println("Timeout")
			//fmt.Println("Failed to complete scan, timeout")
			r := data.ScanStatus{
				ScannerName: "SlowScanner",
				Found:       false,
				Err:         errors.New("Timeout Processing"),
			}

			results <- r
			process = false
			continue

		default:
			//fmt.Println("Slow scanning")
			time.Sleep(time.Millisecond * 500)
		}
	}
	//fmt.Println("Finishing a slow scanner")
}
