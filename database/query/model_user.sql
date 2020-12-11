
-- name: UserCreateOne :one
insert into users (first_name, last_name, credential_id, email_id, phone_id, linked_in_id, is_confirmed, role_id, created_at)
values ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning *;

-- name: UserSetAvatarUriById :one
update users set avatar_uri = $1 where id = $2 returning *;

-- name: UserUpdateConfirmedById :one
update users set is_confirmed = $1 where id = $2 returning *;

-- name: UserGetById :one
select * from users where id = $1 limit 1;

-- name: UserDeleteById :one
delete from users where id = $1 returning *;

-- name: UserGetPreloadedById :one
select u.*, e.*, p.id, li.* from users u join credentials c on c.id = u.credential_id
join emails e on e.id = u.email_id
join phones p on p.id = u.phone_id
join linked_ins li on li.id = u.linked_in_id
where u.id = $1;

