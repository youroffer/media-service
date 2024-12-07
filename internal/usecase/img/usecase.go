package imgUC

import (
	"context"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/cache"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	log "github.com/youroffer/logger"
)

type (
	ImgUC struct {
		db    DBXT
		repo  ImgRepo
		cache Cache
	}

	DBXT interface {
		DB() repository.Querier
		InTransaction(ctx context.Context, fn repository.TransactionFunc) error
	}

	ImgRepo interface {
		GetAllLogos(ctx context.Context, qe repository.Querier) (entity.LogosResp, error)
		GetImageTypeByID(ctx context.Context, qe repository.Querier, id int) (entity.ImageType, error)
		Get(ctx context.Context, qe repository.Querier, params repository.PaginationParams) ([]*entity.Image, error)
		Create(ctx context.Context, qe repository.Querier, image *entity.Image) (*entity.Image, error)
		Update(ctx context.Context, qe repository.Querier, id int, image *entity.ImageUpdate) (*entity.Image, error)
		Delete(ctx context.Context, qe repository.Querier, id int) error
		Count(ctx context.Context, qe repository.Querier) (int, error)
	}

	Cache interface {
		Set(ctx context.Context, key string, value any) error
		Get(ctx context.Context, key string) (string, error)
		Delete(ctx context.Context, prefix string) error
	}
)

func New(db DBXT, repo ImgRepo, cache Cache) *ImgUC {
	return &ImgUC{db: db, repo: repo, cache: cache}
}

func (uc *ImgUC) DeleteImageCache(ctx context.Context, imageType entity.ImageType) {
	var err error

	switch imageType {
	case entity.ImageTypeLogo:
		err = uc.cache.Delete(context.Background(), cache.LogoPrefixKey)

	case entity.ImageTypeAdv:
		err = uc.cache.Delete(context.Background(), cache.AdvPrefixKey)
	}

	if err != nil {
		log.ErrFieldsMsg(err, "delete image cache", log.Fields{
			"image_type": imageType,
		})
	}
}
