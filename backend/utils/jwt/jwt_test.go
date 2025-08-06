package jwt

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const (
	secret     = "test_secret"
	issuer     = "test_issuer"
	accessTTL  = time.Minute * 15
	refreshTTL = time.Hour * 24
)

var signingMethod = jwt.SigningMethodHS256

func Test_GenerateAccessTokenAndRefreshToken(t *testing.T) {
	InitJWTConfig(secret, issuer, accessTTL, refreshTTL, signingMethod)
	userCode := uuid.New()

	tests := []struct {
		name        string
		generate    func(uuid.UUID) (string, error)
		ttl         time.Duration
		expectError bool
	}{
		{
			name:     "Generate valid access token",
			generate: GenerateAccessToken,
			ttl:      accessTTL,
		},
		{
			name:     "Generate valid refresh token",
			generate: GenerateRefreshToken,
			ttl:      refreshTTL,
		},
		{
			name: "JWTConfig not initialized",
			generate: func(u uuid.UUID) (string, error) {
				// Backup and clear config
				old := JWTConfig
				JWTConfig = nil
				defer func() { JWTConfig = old }()
				return GenerateAccessToken(u)
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenStr, err := tt.generate(userCode)
			if tt.expectError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.NotEmpty(t, tokenStr)

			// Validate token
			parsed, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})
			assert.NoError(t, err)
			assert.True(t, parsed.Valid)

			claims, ok := parsed.Claims.(jwt.MapClaims)
			assert.True(t, ok)
			assert.Equal(t, userCode.String(), claims["user_code"])
			assert.Equal(t, issuer, claims["iss"])
			assert.NotZero(t, claims["exp"])
			assert.NotZero(t, claims["iat"])
		})
	}
}

func Test_validateToken(t *testing.T) {
	InitJWTConfig(secret, issuer, accessTTL, refreshTTL, signingMethod)

	userCode := uuid.New()
	tokenStr, err := GenerateAccessToken(userCode)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenStr)

	JWTConfig = nil
	t.Run("JWTConfig not initialized", func(t *testing.T) {
		_, err := validateToken(tokenStr)
		assert.Error(t, err)
	})

	InitJWTConfig(secret, issuer, accessTTL, refreshTTL, signingMethod)

	t.Run("Invalid token", func(t *testing.T) {
		_, err := validateToken(tokenStr + "tampered")
		assert.Error(t, err)
	})

	t.Run("Valid token", func(t *testing.T) {
		token, err := validateToken(tokenStr)
		assert.NoError(t, err)
		assert.NotNil(t, token)
		assert.True(t, token.Valid)
	})
}

func Test_ParseClaims(t *testing.T) {
	InitJWTConfig(secret, issuer, accessTTL, refreshTTL, signingMethod)

	userCode := uuid.New()
	tokenStr, err := GenerateAccessToken(userCode)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenStr)

	claims, err := ParseClaims(tokenStr)
	assert.NoError(t, err)
	assert.NotNil(t, claims)
	assert.Equal(t, userCode.String(), claims["user_code"])
	assert.Equal(t, issuer, claims["iss"])
	assert.NotZero(t, claims["exp"])
	assert.NotZero(t, claims["iat"])
}
