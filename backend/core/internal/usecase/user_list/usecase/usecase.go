package usecase

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"

	"github.com/PrikolTech/alpha/backend/core/internal/usecase/user_list/domain"
)

type Usecase struct {
	userRepo userRepo
}

func New(userRepo userRepo) *Usecase {
	return &Usecase{userRepo: userRepo}
}

func (u *Usecase) Handle(ctx context.Context, in domain.UserListIn) (*domain.UserListOut, error) {
	if err := in.Validate(); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	var out domain.UserListOut

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		users, err := u.userRepo.Get(ctx, in)
		if err != nil {
			return fmt.Errorf("user repo get all: %w", err)
		}
		out.Data = users
		return nil
	})

	g.Go(func() error {
		totalCount, err := u.userRepo.GetTotalCount(ctx, in.Filters)
		if err != nil {
			return fmt.Errorf("user repo get total count: %w", err)
		}
		out.Meta = domain.Meta{
			Page:         in.Page,
			PerPage:      in.PerPage,
			TotalPages:   (totalCount + in.PerPage - 1) / in.PerPage,
			TotalRecords: totalCount,
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return &out, nil
}
