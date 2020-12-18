
create table roles (
                       id          bigserial               primary key,
                       name        varchar                 not null ,
                       description     varchar             default ''
);

create table permissions (
                             id          bigserial               primary key,
                             name        varchar                 not null ,
                             description     varchar             default ''
);

create table role_and_permissions (
                                      role_id     bigint not null ,
                                      permission_id       bigint not null
);

alter table role_and_permissions add foreign key (role_id) references roles (id);
alter table role_and_permissions add foreign key (permission_id) references permissions (id);

