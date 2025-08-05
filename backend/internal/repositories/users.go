package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/juanMaAV92/zenith-financial/backend/internal/entities"
)

const (
	FieldCode  = "code"
	FieldEmail = "email"
)

type Store interface {
	Create(ctx context.Context, destination interface{}) error
	FindOne(ctx context.Context, destination interface{}, conditions interface{}) (bool, error)
}

type UserRepository struct {
	store Store
}

func NewUserRepository(store Store) *UserRepository {
	return &UserRepository{store: store}
}

func (r *UserRepository) Create(ctx context.Context, user *entities.User) error {
	err := r.store.Create(ctx, &user)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetByCode(ctx context.Context, code uuid.UUID) (*entities.User, error) {
	var user entities.User
	condition := map[string]interface{}{FieldCode: code}
	exists, err := r.store.FindOne(ctx, &user, condition)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, nil
	}
	return &user, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	var user entities.User
	condition := map[string]interface{}{FieldEmail: email}
	exists, err := r.store.FindOne(ctx, &user, condition)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, nil
	}
	return &user, nil
}
