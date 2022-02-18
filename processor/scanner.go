package processor

import (
	"context"

	"github.com/ussmith/cs/data"
)

type Scanner interface {
	Scan(ctx context.Context, location string, results chan<- data.ScanStatus)
}
