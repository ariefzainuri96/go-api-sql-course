package store

import (
	"context"
	"database/sql"

	"github.com/ariefzainuri96/go-api-blogging/cmd/api/middleware"
	request "github.com/ariefzainuri96/go-api-blogging/cmd/api/request"
	"github.com/ariefzainuri96/go-api-blogging/cmd/api/response"
)

type AuthStore struct {
	db *sql.DB
}

func (store *AuthStore) Login(ctx context.Context, body request.LoginRequest) (response.LoginResponse, error) {
	query := `SELECT * FROM users_login WHERE email = $1 AND password = $2`

	row := store.db.QueryRowContext(ctx, query, body.Email, body.Password)

	var login response.LoginResponse

	token, err := middleware.GenerateToken(body.Email)

	if err != nil {
		return login, err
	}

	err = row.Scan(&login.ID, &login.Email, &login.CreatedAt)

	login.Token = token

	if err != nil {
		return login, err
	}

	return login, nil
}
