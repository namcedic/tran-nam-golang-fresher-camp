package uploadprovider

import (
	"context"
	"food_delivery_service/common"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
}
