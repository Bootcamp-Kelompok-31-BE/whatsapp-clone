package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/config"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/shared/models"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/shared/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type userIDKey string
type userEmailKey string

const (
	UserIDKey    userIDKey    = "user_id"
	UserEmailKey userEmailKey = "user_email"
)

type JWTClaims struct {
	UserID    uint   `json:"user_id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	IssuedAt  int64  `json:"iat"`
	ExpiresAt int64  `json:"exp"`
	jwt.RegisteredClaims
}

func GenerateToken(config *config.Config, userID uint, email, username string) (string, error) {
	now := time.Now()
	expiryTime := now.Add(time.Duration(config.JWT.ExpiredAt) * time.Second)

	claims := &JWTClaims{
		UserID:    userID,
		Email:     email,
		Username:  username,
		IssuedAt:  now.Unix(),
		ExpiresAt: expiryTime.Unix(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiryTime),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Subject:   strconv.Itoa(int(userID)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWT.SecretKey))
}

func ValidateToken(config *config.Config, tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(config.JWT.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("token expired")
	}

	return claims, nil
}

func JWTMiddleware(config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := util.LoggerFromContext(c.Request.Context())

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			logger.Warn("missing authorization header")
			c.JSON(http.StatusUnauthorized, models.NewErrResponse("authorization header required"))
			c.Abort()
			return
		}

		tokenParts := strings.SplitN(authHeader, " ", 2)
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			logger.Warn("invalid authorization header format")
			c.JSON(http.StatusUnauthorized, models.NewErrResponse("invalid authorization header format"))
			c.Abort()
			return
		}

		tokenString := tokenParts[1]
		if tokenString == "" {
			logger.Warn("missing token in authorization header")
			c.JSON(http.StatusUnauthorized, models.NewErrResponse("token required"))
			c.Abort()
			return
		}

		claims, err := ValidateToken(config, tokenString)
		if err != nil {
			logger.Warn("token validation failed", zap.Error(err))
			c.JSON(http.StatusUnauthorized, models.NewErrResponse("invalid or expired token"))
			c.Abort()
			return
		}

		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, UserIDKey, claims.UserID)
		ctx = context.WithValue(ctx, UserEmailKey, claims.Email)

		c.Request = c.Request.WithContext(ctx)

		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Set("username", claims.Username)
		c.Set("jwt_claims", claims)

		logger.Debug("user authenticated successfully",
			zap.Uint("user_id", claims.UserID),
			zap.String("email", claims.Email),
		)

		c.Next()
	}
}

func GetUserIDFromContext(ctx context.Context) (uint, bool) {
	userID, ok := ctx.Value(UserIDKey).(uint)
	return userID, ok
}

func GetUserEmailFromContext(ctx context.Context) (string, bool) {
	email, ok := ctx.Value(UserEmailKey).(string)
	return email, ok
}
