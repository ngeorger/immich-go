package adapters

import (
	"context"

	"github.com/ngeorger/immich-go/internal/assets"
	"github.com/spf13/cobra"
)

type Reader interface {
	Browse(cxt context.Context) chan *assets.Group
}

type AssetWriter interface {
	WriteAsset(context.Context, *assets.Asset) error
	// WriteGroup(ctx context.Context, group *assets.Group) error
}

type Runner interface {
	Run(cmd *cobra.Command, adapter Reader) error
}
