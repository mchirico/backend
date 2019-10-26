package cmd

import (
	"context"
	"testing"
	"time"
)

func TestEntryPt(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(4*time.Second))
	defer cancel()

	EntryPt(ctx)
}
