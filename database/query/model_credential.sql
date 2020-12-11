
-- name: CredentialCreateOne :one
insert into credentials (username, password, created_at) values ($1, $2, $3) returning *;

-- name: CredentialGetByUsername :one
select * from credentials where username = $1 limit 1;

-- name: CredentialGetById :one
select * from credentials where id = $1 limit 1;

-- name: CredentialResetPasswordByUsername :one
update credentials set password = $1 where username = $2 returning *;

-- name: CredentialResetPasswordById :one
update credentials set password = $1 where id = $2 returning *;

