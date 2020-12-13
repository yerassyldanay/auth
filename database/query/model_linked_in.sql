-- name: LinkedInGetById :one
select * from linked_ins where id = $1 limit 1;

-- name: LinkedInCreateOne :one
insert into linked_ins (login, linked_in_confirmed_at)
values ($1, $2) returning *;

-- name: LinkedInDeleteById :one
delete from linked_ins where id = $1 returning *;

-- name: LinkedInDeleteByLogin :one
delete from linked_ins where login = $1 returning *;

-- name: LinkedInConfirmByLogin :one
update linked_ins set linked_in_confirmed_at = $1
where login = $2 returning *;

-- name: LinkedInConfirmById :one
update linked_ins set linked_in_confirmed_at = $1
where id = $2 returning *;

