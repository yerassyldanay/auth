-- name: PhoneCreateOne :one
insert into phones (country_code, number, phone_confirmed_at, is_phone_confirmed)
values ($1, $2, $3, $4) returning *;

-- name: PhoneDeleteById :one
delete from phones where id = $1 returning *;

-- name: PhoneDeleteByNumberAndCountryCode :one
delete from phones where country_code = $1 and number = $2 returning *;

-- name: PhoneConfirmById :one
update phones set is_phone_confirmed = $1, phone_confirmed_at = $2
where id = $3 returning *;

-- name: PhoneConfirmByNumberAndCountryCode :one
update phones set is_phone_confirmed = $1, phone_confirmed_at = $2
where number = $3 and country_code = $4 returning *;

-- name: PhoneGetById :one
select * from phones where id = $1 limit 1;

-- name: PhoneGetByNumberAndCountryCode :one
select * from phones where country_code = $1 and number = $2 limit 1 ;
