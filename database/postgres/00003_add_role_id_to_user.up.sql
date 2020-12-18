alter table users add column role_id bigint
    references roles (id);