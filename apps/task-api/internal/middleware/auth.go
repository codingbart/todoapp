package middleware

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/MicahParks/keyfunc/v3"
	db "github.com/codingbart/todoapp/task-api/internal/db/postgresql"
	"github.com/codingbart/todoapp/task-api/internal/response"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type contextKey string

const UserIDKey contextKey = "userID"

type Auth interface {
	Protect(next http.Handler) http.Handler
}

type AuthMiddleware struct {
	jwks    keyfunc.Keyfunc
	queries *db.Queries
}

type userClaims struct {
	sub   string
	name  string
	email string
}

func NewAuthMiddleware(jwksURL string, queries *db.Queries) (*AuthMiddleware, error) {
	jwks, err := keyfunc.NewDefault([]string{jwksURL})
	if err != nil {
		return nil, err
	}

	return &AuthMiddleware{jwks: jwks, queries: queries}, nil
}

func (m *AuthMiddleware) Protect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr, ok := extractAuthToken(r)
		if !ok {
			response.Error(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		claims, ok := m.parseTokenClaims(tokenStr)
		if !ok {
			response.Error(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		uc := extractClaims(claims)

		user, err := m.syncUser(r.Context(), uc)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "Failed to sync user")
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, user.ID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserID(r *http.Request) (uuid.UUID, bool) {
	id, ok := r.Context().Value(UserIDKey).(uuid.UUID)
	return id, ok
}

func (m *AuthMiddleware) syncUser(ctx context.Context, uc userClaims) (db.User, error) {
	user, err := m.queries.FindUserByKeycloakId(ctx, uc.sub)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return db.User{}, err
	}

	if errors.Is(err, sql.ErrNoRows) {
		user, err = m.queries.SaveUser(ctx, db.SaveUserParams{
			KeycloakID: uc.sub,
			Name:       uc.name,
			Email:      uc.email,
		})
		if err != nil {
			return db.User{}, err
		}
	}

	return user, nil
}

func extractClaims(claims jwt.MapClaims) userClaims {
	sub, _ := claims["sub"].(string)
	name, _ := claims["name"].(string)
	email, _ := claims["email"].(string)
	return userClaims{sub: sub, name: name, email: email}
}

func extractAuthToken(r *http.Request) (string, bool) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return "", false
	}
	return strings.TrimPrefix(authHeader, "Bearer "), true
}

func (m *AuthMiddleware) parseTokenClaims(tokenStr string) (jwt.MapClaims, bool) {
	token, err := jwt.Parse(tokenStr, m.jwks.Keyfunc)
	if err != nil || !token.Valid {
		return nil, false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	return claims, ok
}
