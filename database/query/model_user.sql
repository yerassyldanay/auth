
-- name: UserCreateOne :one
insert into users (name, password, description, email_id, phone_id, linked_in_id, role_id, created_at)
values ($1, $2, $3, $4, $5, $6, $7, $8) returning *;

-- name: UserSetAvatarUriById :one
update users set avatar_uri = $1 where id = $2 returning *;

-- name: UserGetById :one
select * from users where id = $1 limit 1;

-- name: UserDeleteById :one
delete from users where id = $1 returning *;

-- name: UserGetPreloadedById :one
select u.*, e.*, p.id, li.* from users u
join emails e on e.id = u.email_id
join phones p on p.id = u.phone_id
join linked_ins li on li.id = u.linked_in_id
where u.id = $1;

-- name: UserGetOneByEmailId :one
select u.* from users u join emails e on e.id = u.email_id
where e.id = $1 limit 1;

-- name: UserGetOneByEmailAddress :one
select u.* from users u join emails e on e.id = u.email_id
where e.address = $1 limit 1;

