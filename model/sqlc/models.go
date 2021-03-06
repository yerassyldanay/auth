// Code generated by sqlc. DO NOT EDIT.

package database

import (
	"database/sql"
)

type CountryCode struct {
	CountryCode string         `json:"country_code"`
	Country     string         `json:"country"`
	IsoCode     sql.NullString `json:"iso_code"`
	Continent   sql.NullString `json:"continent"`
}

type Email struct {
	ID               int64        `json:"id"`
	Address          string       `json:"address"`
	EmailConfirmedAt sql.NullTime `json:"email_confirmed_at"`
}

type LinkedIn struct {
	ID                  int64          `json:"id"`
	Login               sql.NullString `json:"login"`
	LinkedInConfirmedAt sql.NullTime   `json:"linked_in_confirmed_at"`
}

type Permission struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
}

type Phone struct {
	ID               int64        `json:"id"`
	CountryCode      string       `json:"country_code"`
	Number           string       `json:"number"`
	PhoneConfirmedAt sql.NullTime `json:"phone_confirmed_at"`
}

type Role struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
}

type RoleAndPermission struct {
	RoleID       int64 `json:"role_id"`
	PermissionID int64 `json:"permission_id"`
}

type User struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Password    sql.NullString `json:"password"`
	AvatarUri   sql.NullString `json:"avatar_uri"`
	EmailID     sql.NullInt64  `json:"email_id"`
	PhoneID     sql.NullInt64  `json:"phone_id"`
	LinkedInID  sql.NullInt64  `json:"linked_in_id"`
	CreatedAt   sql.NullTime   `json:"created_at"`
	RoleID      sql.NullInt64  `json:"role_id"`
}
