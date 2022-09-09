package service

import (
	"context"
	"time"

	"unittest-with-mock/domain"
	"unittest-with-mock/model/reqres"
)

type userService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) domain.UserService {
	return &userService{repo: repo}
}

func (u *userService) Create(ctx context.Context, req *reqres.CreateUserReq) error {
	newUser := domain.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := u.repo.Save(ctx, &newUser)
	if err != nil {
		return err
	}

	return nil
}

func (u *userService) UpdateByID(ctx context.Context, id int64, req *reqres.UpdateUserReq) error {
	newUser := domain.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}

	return u.repo.UpdateByID(ctx, id, &newUser)
}

func (u *userService) DeleteByID(ctx context.Context, id int64) error {
	return u.repo.DeleteByID(ctx, id)
}

func (u *userService) FindAll(ctx context.Context) (*[]domain.User, error) {
	return u.repo.FindAll(ctx)
}

func (u *userService) FindByID(ctx context.Context, id int64) (*domain.User, error) {
	return u.repo.FindByID(ctx, id)
}
