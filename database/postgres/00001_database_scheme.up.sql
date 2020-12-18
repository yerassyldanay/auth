create table users (
    id              bigserial    primary key ,
    name                   varchar           not null,
    description             varchar            default '',
    password        varchar             default '',
    avatar_uri           varchar,
    email_id        bigint default 0,
    phone_id        bigint default 0,
    linked_in_id     bigint default 0,
    created_at         timestamptz             default 'now()'
);

create table emails (
    id          bigserial           primary key ,
    address     varchar             unique not null,
    email_confirmed_at      timestamptz             default 'now()'
);

create table phones (
    id              bigserial           primary key ,
    country_code           varchar      not null,
    number          varchar             not null,
    phone_confirmed_at    timestamptz             default 'now()',
    unique (country_code, number)
);

create table country_codes (
    country_code        varchar         primary key ,
    country             varchar         not null ,
    iso_code            varchar         default '',
    continent           varchar         default ''
);

create table linked_ins (
    id          bigserial           primary key ,
    login       varchar             unique,
    linked_in_confirmed_at    timestamptz             default 'now()'
);

alter table users add foreign key (email_id) references emails (id);
alter table users add foreign key (phone_id) references phones (id);
alter table users add foreign key (linked_in_id) references linked_ins (id);

alter table phones add foreign key (country_code) references country_codes (country_code);
