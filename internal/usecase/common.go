package usecase

import (
	"context"

	"golps/internal/usecase/entity"
)

type (
	Repo interface{}

	WebAPI interface {
		Run(*UseCase)
		NewText(msg entity.Message, withMarkup bool) error
		EditTextAndMarkup(msg entity.Message) error
	}

	UseCase struct {
		Ctx    context.Context
		Repo   Repo
		WebAPI WebAPI
	}
)

func New(ctx context.Context, r Repo, w WebAPI) *UseCase {
	return &UseCase{
		Ctx:    ctx,
		Repo:   r,
		WebAPI: w,
	}
}

func (u *UseCase) Start() {
	u.WebAPI.Run(u)
}
