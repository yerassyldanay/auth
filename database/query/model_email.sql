
-- name: EmailCreateOne :one
insert into emails (address) values ($1) returning *;

-- name: EmailCreateConfirmedOne :one
insert into emails (address, email_confirmed_at) values ($1, $2) returning *;

-- name: EmailConfirmById :one
update emails set email_confirmed_at = $1 where id = $2 returning *;

-- name: EmailConfirmByAddress :one
update emails set email_confirmed_at = $1 where address = $2 returning *;

-- name: EmailDeleteById :one
delete from emails where id = $1 returning *;

-- name: EmailDeleteByAddress :one
delete from emails where address = $1 returning *;

-- name: EmailGetOneById :one
select * from emails where id = $1 limit 1;

-- name: EmailGetOneByAddress :one
select * from emails where address = $1 limit 1;
