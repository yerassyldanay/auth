// Code generated by sqlc. DO NOT EDIT.
// source: model_phone.sql

package database

import (
	"context"
	"database/sql"
)

const phoneConfirmById = `-- name: PhoneConfirmById :one
update phones set is_phone_confirmed = $1, phone_confirmed_at = $2
where id = $3 returning id, country_code, number, phone_confirmed_at, is_phone_confirmed
`

type PhoneConfirmByIdParams struct {
	IsPhoneConfirmed sql.NullBool `json:"is_phone_confirmed"`
	PhoneConfirmedAt sql.NullTime `json:"phone_confirmed_at"`
	ID               int64        `json:"id"`
}

func (q *Queries) PhoneConfirmById(ctx context.Context, arg PhoneConfirmByIdParams) (Phone, error) {
	row := q.db.QueryRowContext(ctx, phoneConfirmById, arg.IsPhoneConfirmed, arg.PhoneConfirmedAt, arg.ID)
	var i Phone
	err := row.Scan(
		&i.ID,
		&i.CountryCode,
		&i.Number,
		&i.PhoneConfirmedAt,
		&i.IsPhoneConfirmed,
	)
	return i, err
}

const phoneConfirmByNumberAndCountryCode = `-- name: PhoneConfirmByNumberAndCountryCode :one
update phones set is_phone_confirmed = $1, phone_confirmed_at = $2
where number = $3 and country_code = $4 returning id, country_code, number, phone_confirmed_at, is_phone_confirmed
`

type PhoneConfirmByNumberAndCountryCodeParams struct {
	IsPhoneConfirmed sql.NullBool `json:"is_phone_confirmed"`
	PhoneConfirmedAt sql.NullTime `json:"phone_confirmed_at"`
	Number           string       `json:"number"`
	CountryCode      string       `json:"country_code"`
}

func (q *Queries) PhoneConfirmByNumberAndCountryCode(ctx context.Context, arg PhoneConfirmByNumberAndCountryCodeParams) (Phone, error) {
	row := q.db.QueryRowContext(ctx, phoneConfirmByNumberAndCountryCode,
		arg.IsPhoneConfirmed,
		arg.PhoneConfirmedAt,
		arg.Number,
		arg.CountryCode,
	)
	var i Phone
	err := row.Scan(
		&i.ID,
		&i.CountryCode,
		&i.Number,
		&i.PhoneConfirmedAt,
		&i.IsPhoneConfirmed,
	)
	return i, err
}

const phoneCreateOne = `-- name: PhoneCreateOne :one
insert into phones (country_code, number, phone_confirmed_at, is_phone_confirmed)
values ($1, $2, $3, $4) returning id, country_code, number, phone_confirmed_at, is_phone_confirmed
`

type PhoneCreateOneParams struct {
	CountryCode      string       `json:"country_code"`
	Number           string       `json:"number"`
	PhoneConfirmedAt sql.NullTime `json:"phone_confirmed_at"`
	IsPhoneConfirmed sql.NullBool `json:"is_phone_confirmed"`
}

func (q *Queries) PhoneCreateOne(ctx context.Context, arg PhoneCreateOneParams) (Phone, error) {
	row := q.db.QueryRowContext(ctx, phoneCreateOne,
		arg.CountryCode,
		arg.Number,
		arg.PhoneConfirmedAt,
		arg.IsPhoneConfirmed,
	)
	var i Phone
	err := row.Scan(
		&i.ID,
		&i.CountryCode,
		&i.Number,
		&i.PhoneConfirmedAt,
		&i.IsPhoneConfirmed,
	)
	return i, err
}

const phoneDeleteById = `-- name: PhoneDeleteById :one
delete from phones where id = $1 returning id, country_code, number, phone_confirmed_at, is_phone_confirmed
`

func (q *Queries) PhoneDeleteById(ctx context.Context, id int64) (Phone, error) {
	row := q.db.QueryRowContext(ctx, phoneDeleteById, id)
	var i Phone
	err := row.Scan(
		&i.ID,
		&i.CountryCode,
		&i.Number,
		&i.PhoneConfirmedAt,
		&i.IsPhoneConfirmed,
	)
	return i, err
}

const phoneDeleteByNumberAndCountryCode = `-- name: PhoneDeleteByNumberAndCountryCode :one
delete from phones where country_code = $1 and number = $2 returning id, country_code, number, phone_confirmed_at, is_phone_confirmed
`

type PhoneDeleteByNumberAndCountryCodeParams struct {
	CountryCode string `json:"country_code"`
	Number      string `json:"number"`
}

func (q *Queries) PhoneDeleteByNumberAndCountryCode(ctx context.Context, arg PhoneDeleteByNumberAndCountryCodeParams) (Phone, error) {
	row := q.db.QueryRowContext(ctx, phoneDeleteByNumberAndCountryCode, arg.CountryCode, arg.Number)
	var i Phone
	err := row.Scan(
		&i.ID,
		&i.CountryCode,
		&i.Number,
		&i.PhoneConfirmedAt,
		&i.IsPhoneConfirmed,
	)
	return i, err
}

const phoneGetById = `-- name: PhoneGetById :one
select id, country_code, number, phone_confirmed_at, is_phone_confirmed from phones where id = $1 limit 1
`

func (q *Queries) PhoneGetById(ctx context.Context, id int64) (Phone, error) {
	row := q.db.QueryRowContext(ctx, phoneGetById, id)
	var i Phone
	err := row.Scan(
		&i.ID,
		&i.CountryCode,
		&i.Number,
		&i.PhoneConfirmedAt,
		&i.IsPhoneConfirmed,
	)
	return i, err
}

const phoneGetByNumberAndCountryCode = `-- name: PhoneGetByNumberAndCountryCode :one
select id, country_code, number, phone_confirmed_at, is_phone_confirmed from phones where country_code = $1 and number = $2 limit 1
`

type PhoneGetByNumberAndCountryCodeParams struct {
	CountryCode string `json:"country_code"`
	Number      string `json:"number"`
}

func (q *Queries) PhoneGetByNumberAndCountryCode(ctx context.Context, arg PhoneGetByNumberAndCountryCodeParams) (Phone, error) {
	row := q.db.QueryRowContext(ctx, phoneGetByNumberAndCountryCode, arg.CountryCode, arg.Number)
	var i Phone
	err := row.Scan(
		&i.ID,
		&i.CountryCode,
		&i.Number,
		&i.PhoneConfirmedAt,
		&i.IsPhoneConfirmed,
	)
	return i, err
}