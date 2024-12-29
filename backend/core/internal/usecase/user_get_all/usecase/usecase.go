package usecase

import (
	"context"
	"fmt"
	"math"

	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_get_all/domain"
	"golang.org/x/sync/errgroup"
)

type Usecase struct {
	userRepo userRepo
}

func New(userRepo userRepo) *Usecase {
	return &Usecase{userRepo: userRepo}
}

func (u *Usecase) Handle(ctx context.Context, in domain.UserGetAllIn) (*domain.UserGetAllOut, error) {
	if err := in.Validate(); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	var out domain.UserGetAllOut

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		users, err := u.userRepo.GetAll(ctx, in)
		if err != nil {
			return fmt.Errorf("user get all: %w", err)
		}
		out.Data = users
		return nil
	})

	g.Go(func() error {
		totalCount, err := u.userRepo.GetTotalCount(ctx)
		if err != nil {
			return fmt.Errorf("user get total count: %w", err)
		}
		out.Meta = domain.Meta{
			Page:         in.Page,
			PerPage:      in.PerPage,
			TotalPages:   int(math.Ceil(float64(totalCount) / float64(in.PerPage))),
			TotalRecords: totalCount,
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return &out, nil
}
