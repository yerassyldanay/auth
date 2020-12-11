
-- name: EmailCreateOne :one
insert into emails (address) values ($1) returning *;

-- name: EmailCreateConfirmedOne :one
insert into emails (address, email_confirmed_at, is_email_confirmed) values ($1, $2, $3) returning *;

-- name: EmailConfirmById :one
update emails set is_email_confirmed = true, email_confirmed_at = now() where id = $1 returning *;

-- name: EmailConfirmByAddress :one
update emails set is_email_confirmed = true, email_confirmed_at = now() where address = $1 returning *;

-- name: EmailDeleteById :one
delete from emails where id = $1 returning *;

-- name: EmailDeleteByAddress :one
delete from emails where id = $1 returning *;

