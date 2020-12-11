alter table users add column role_id bigserial
    references roles (id);