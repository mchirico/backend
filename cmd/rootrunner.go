package cmd

import (
	"context"
	"github.com/mchirico/backend/api"
)

func EntryPt(ctx context.Context) {
	api := api.NewAPIfile("/dummy.csv")
	api.MainListen(ctx)
}
