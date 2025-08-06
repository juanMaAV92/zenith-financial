package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type jwtConfig struct {
	SecretKey       string
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
	Issuer          string
	SigningMethod   jwt.SigningMethod
}

var JWTConfig *jwtConfig

func InitJWTConfig(secretKey, issuer string, accessTokenTTL, refreshTokenTTL time.Duration, signingMethod jwt.SigningMethod) {
	JWTConfig = &jwtConfig{
		SecretKey:       secretKey,
		AccessTokenTTL:  accessTokenTTL,
		RefreshTokenTTL: refreshTokenTTL,
		Issuer:          issuer,
		SigningMethod:   signingMethod,
	}
}

func generateToken(userCode uuid.UUID, ttl time.Duration, tokenType string) (string, error) {
	if JWTConfig.SigningMethod == nil {
		return "", errors.New("JWTConfig or SigningMethod not initialized")
	}
	now := time.Now()
	claims := jwt.MapClaims{
		"user_code": userCode.String(),
		"exp":       now.Add(ttl).Unix(),
		"iat":       now.Unix(),
		"iss":       JWTConfig.Issuer,
		"type":      "access",
	}
	token := jwt.NewWithClaims(JWTConfig.SigningMethod, claims)
	return token.SignedString([]byte(JWTConfig.SecretKey))
}

func GenerateAccessToken(userCode uuid.UUID) (string, error) {
	if JWTConfig == nil {
		return "", errors.New("JWTConfig not initialized")
	}
	return generateToken(userCode, JWTConfig.AccessTokenTTL, "access")
}

func GenerateRefreshToken(userCode uuid.UUID) (string, error) {
	if JWTConfig == nil {
		return "", errors.New("JWTConfig not initialized")
	}
	return generateToken(userCode, JWTConfig.RefreshTokenTTL, "refresh")
}

func validateToken(tokenString string) (*jwt.Token, error) {
	if JWTConfig == nil || JWTConfig.SigningMethod == nil {
		return nil, errors.New("JWTConfig or SigningMethod not initialized")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(JWTConfig.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token, nil
}

func ParseClaims(tokenString string) (jwt.MapClaims, error) {
	token, err := validateToken(tokenString)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}
	return claims, nil
}
