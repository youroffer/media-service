package imgRepo

import (
	"context"
	"errors"

	"github.com/himmel520/media-service/internal/entity"
	"github.com/himmel520/media-service/internal/infrastructure/repository"
	"github.com/himmel520/media-service/internal/infrastructure/repository/repoerr"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *ImgRepo) Create(ctx context.Context, qe repository.Querier, logo *entity.Logo) (*entity.LogoResp, error) {
	query, args, err := squirrel.Insert("colors").
		Columns(
			"title",
			"hex").
		Values(
			color.Title,
			color.Hex).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("returning id, title, hex").
		ToSql()
	if err != nil {
		return nil, err
	}
	
	newLogo := &entity.LogoResp{}

	err := qe.QueryRow(ctx, `
	insert into logos 
		(url, title) 
	values 
		($1, $2) 
	returning *`, logo.Url, logo.Title).Scan(&newLogo.ID, &newLogo.Url, &newLogo.Title)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == repoerr.UniqueConstraint {
		return nil, repoerr.ErrLogoExist
	}

	return newLogo, err
}
