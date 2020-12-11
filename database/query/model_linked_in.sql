-- name: LinkedInGetById :one
select * from linked_ins where id = $1 limit 1;

-- name: LinkedInCreateOne :one
insert into linked_ins (login, linked_in_confirmed_at, is_linked_in_confirmed)
values ($1, $2, $3) returning *;

-- name: LinkedInDeleteById :one
delete from linked_ins where id = $1 returning *;

-- name: LinkedInDeleteByLogin :one
delete from linked_ins where login = $1 returning *;

-- name: LinkedInConfirmByLogin :one
update linked_ins set is_linked_in_confirmed = $1, linked_in_confirmed_at = $2
where login = $3 returning *;

-- name: LinkedInConfirmById :one
update linked_ins set is_linked_in_confirmed = $1, linked_in_confirmed_at = $2
where id = $3 returning *;

