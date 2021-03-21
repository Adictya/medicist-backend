// Code generated by sqlc. DO NOT EDIT.
// source: user.sql

package db

import (
	"context"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO users (
	full_name,mobile,type
	) VALUES (
	$1, $2, $3
)
RETURNING id, full_name, created_at, mobile, type
`

type CreateAccountParams struct {
	FullName string `json:"full_name"`
	Mobile   int64  `json:"mobile"`
	Type     int32  `json:"type"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createAccount, arg.FullName, arg.Mobile, arg.Type)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.CreatedAt,
		&i.Mobile,
		&i.Type,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteAccount, id)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT id, full_name, created_at, mobile, type FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getAccount, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.CreatedAt,
		&i.Mobile,
		&i.Type,
	)
	return i, err
}

const listAccounts = `-- name: ListAccounts :many
SELECT id, full_name, created_at, mobile, type FROM users
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListAccountsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAccounts(ctx context.Context, arg ListAccountsParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listAccounts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FullName,
			&i.CreatedAt,
			&i.Mobile,
			&i.Type,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAccount = `-- name: UpdateAccount :one
UPDATE users
SET type = $2
WHERE id = $1
RETURNING id, full_name, created_at, mobile, type
`

type UpdateAccountParams struct {
	ID   int32 `json:"id"`
	Type int32 `json:"type"`
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateAccount, arg.ID, arg.Type)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.CreatedAt,
		&i.Mobile,
		&i.Type,
	)
	return i, err
}
