-- name: PhoneCreateOne :one
insert into phones (country_code, number, phone_confirmed_at)
values ($1, $2, $3) returning *;

-- name: PhoneDeleteById :one
delete from phones where id = $1 returning *;

-- name: PhoneDeleteByCountryCodeAndNumber :one
delete from phones where number = $1 and country_code = $2 returning *;

-- name: PhoneConfirmById :one
update phones set phone_confirmed_at = $1 where id = $2 returning *;

-- name: PhoneConfirmByNumberAndCountryCode :one
update phones set phone_confirmed_at = $1
where number = $2 and country_code = $3 returning *;

-- name: PhoneGetById :one
select * from phones where id = $1 limit 1;

-- name: PhoneGetByNumberAndCountryCode :one
select * from phones where country_code = $1 and number = $2 limit 1 ;
