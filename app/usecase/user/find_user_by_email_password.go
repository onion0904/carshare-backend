package user

import (
	"context"
	userDomain "github.com/onion0904/CarShareSystem/app/domain/user"
	"time"
)

type FindUserByEmailPasswordUseCase struct {
	userRepo userDomain.UserRepository
}

func NewFindUserByEmailPasswordUseCase(
	userRepo userDomain.UserRepository,
) *FindUserByEmailPasswordUseCase {
	return &FindUserByEmailPasswordUseCase{
		userRepo: userRepo,
	}
}

type FindUserByEmailPasswordUseCaseDto struct {
	ID          string
	LastName    string
	FirstName   string
	Email       string
	Password    string
	Icon        string
	GroupIDs    []string
	EventIDs    []string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (uc *FindUserByEmailPasswordUseCase) Run(ctx context.Context, email,password string) (*FindUserByEmailPasswordUseCaseDto, error) {
	user, err := uc.userRepo.FindUserByEmailPassword(ctx,email,password)
	if err != nil {
		return nil, err
	}
	return &FindUserByEmailPasswordUseCaseDto{
		ID:          user.ID(),
		LastName:    user.LastName(),
		FirstName:   user.FirstName(),
		Email:       user.Email(),
		Password:    user.Password(),
		Icon:        user.Icon(),
		GroupIDs:    user.GroupIDs(),
		EventIDs:    user.EventIDs(),
		CreatedAt:   user.CreatedAt(),
        UpdatedAt:   user.UpdatedAt(),
	}, nil
}