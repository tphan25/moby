package images

import (
	"github.com/docker/docker/api/types"
)

func (i *ImageService) ImageMakeCache() (*types.ImageMakeCacheResult, error) {
	return &types.ImageMakeCacheResult{
		Success: true,
	}, nil
}
