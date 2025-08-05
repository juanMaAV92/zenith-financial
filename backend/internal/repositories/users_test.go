package repositories

import (
	"context"
	"errors"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
	"github.com/juanMaAV92/zenith-financial/backend/internal/entities"
	"github.com/stretchr/testify/mock"
)

type MockStore struct {
	mock.Mock
}

func (m *MockStore) FindOne(ctx context.Context, destination interface{}, conditions interface{}) (bool, error) {
	args := m.Called(ctx, destination, conditions)
	return args.Bool(0), args.Error(1)
}

func (m *MockStore) Create(ctx context.Context, destination interface{}) error {
	args := m.Called(ctx, destination)
	return args.Error(0)
}

func Test_UserRepository_GetByEmail(t *testing.T) {
	ctx := context.Background()

	testCase := []struct {
		name         string
		email        string
		mockFunc     func(*MockStore, string)
		expectedUser *entities.User
		expectError  error
	}{
		{
			name:  "find one by email",
			email: "test@test.com",
			mockFunc: func(store *MockStore, email string) {
				store.On("FindOne",
					mock.Anything,
					&entities.User{},
					map[string]interface{}{FieldEmail: email},
				).Return(true, nil).Run(func(args mock.Arguments) {
					user := args.Get(1).(*entities.User)
					user.Email = email
				})
			},
			expectedUser: &entities.User{
				Email: "test@test.com",
			},
			expectError: nil,
		},
		{
			name:  "not found user",
			email: "test_not_found@test.com",
			mockFunc: func(store *MockStore, email string) {
				store.On("FindOne",
					mock.Anything,
					&entities.User{},
					map[string]interface{}{FieldEmail: email},
				).Return(false, nil)
			},
			expectedUser: nil,
			expectError:  nil,
		},
		{
			name:  "error finding user",
			email: "test_error@test.com",
			mockFunc: func(store *MockStore, email string) {
				store.On("FindOne", mock.Anything, &entities.User{}, map[string]interface{}{FieldEmail: email}).Return(false, errors.New("error finding user"))
			},
			expectedUser: nil,
			expectError:  errors.New("error finding user"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			store := &MockStore{}
			repo := NewUserRepository(store)

			tc.mockFunc(store, tc.email)
			user, err := repo.GetByEmail(ctx, tc.email)

			if tc.expectError != nil {
				assert.Equal(t, tc.expectError, err)
			} else {
				if err != nil {
					t.Fatalf("Error getting user by email: %v", err)
				}
				assert.Equal(t, tc.expectedUser, user)
			}
		})
	}
}

func Test_UserRepository_GetByCode(t *testing.T) {
	ctx := context.Background()

	testCase := []struct {
		name         string
		code         uuid.UUID
		mockFunc     func(*MockStore, uuid.UUID)
		expectedUser *entities.User
		expectError  error
	}{
		{
			name: "find one by code",
			code: uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"),
			mockFunc: func(store *MockStore, code uuid.UUID) {
				store.On("FindOne",
					mock.Anything,
					&entities.User{},
					map[string]interface{}{FieldCode: code},
				).Return(true, nil).Run(func(args mock.Arguments) {
					user := args.Get(1).(*entities.User)
					user.Code = code
				})
			},
			expectedUser: &entities.User{
				Code: uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"),
			},
			expectError: nil,
		},
		{
			name: "not found user",
			code: uuid.New(),
			mockFunc: func(store *MockStore, code uuid.UUID) {
				store.On("FindOne",
					mock.Anything,
					&entities.User{},
					map[string]interface{}{FieldCode: code},
				).Return(false, nil)
			},
			expectedUser: nil,
			expectError:  nil,
		},
		{
			name: "error finding user",
			code: uuid.New(),
			mockFunc: func(store *MockStore, code uuid.UUID) {
				store.On("FindOne",
					mock.Anything,
					&entities.User{},
					map[string]interface{}{FieldCode: code},
				).Return(false, errors.New("error finding user"))
			},
			expectedUser: nil,
			expectError:  errors.New("error finding user"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			store := &MockStore{}
			repo := NewUserRepository(store)

			if tc.name == "find one by code" {
				tc.expectedUser = &entities.User{
					Code: tc.code,
				}
			}

			tc.mockFunc(store, tc.code)
			user, err := repo.GetByCode(ctx, tc.code)

			if tc.expectError != nil {
				assert.Equal(t, tc.expectError, err)
			} else {
				if err != nil {
					t.Fatalf("Error getting user by code: %v", err)
				}
				assert.Equal(t, tc.expectedUser, user)
			}
		})
	}
}

func Test_UserRepository_Create(t *testing.T) {
	ctx := context.Background()

	testCase := []struct {
		name        string
		user        *entities.User
		mockFunc    func(*MockStore)
		expectError error
	}{
		{
			name: "create user",
			user: &entities.User{
				Email: "test@test.com",
			},
			mockFunc: func(store *MockStore) {
				store.On("Create", mock.Anything, mock.Anything).Return(nil)
			},
			expectError: nil,
		},
		{
			name: "error creating user",
			user: &entities.User{
				Email: "test@test.com",
			},
			mockFunc: func(store *MockStore) {
				store.On("Create", mock.Anything, mock.Anything).Return(errors.New("error creating user"))
			},
			expectError: errors.New("error creating user"),
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			store := &MockStore{}
			repo := NewUserRepository(store)

			tc.mockFunc(store)
			err := repo.Create(ctx, tc.user)

			if tc.expectError != nil {
				assert.Equal(t, tc.expectError, err)
			} else {
				if err != nil {
					t.Fatalf("Error creating user: %v", err)
				}
			}
		})
	}
}
