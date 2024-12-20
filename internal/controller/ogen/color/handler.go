package color

import (
	"context"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/usecase"
)

type (
	Handler struct {
		uc ColorUsecase
	}

	ColorUsecase interface {
		Get(ctx context.Context, params usecase.PageParams) (*entity.ColorsResp, error)
		Create(ctx context.Context, color *entity.Color) (*entity.Color, error)
		Update(ctx context.Context, id int, color *entity.ColorUpdate) (*entity.Color, error)
		Delete(ctx context.Context, id int) error
	}
)

func New(uc ColorUsecase) *Handler {
	return &Handler{
		uc: uc,
	}
}
