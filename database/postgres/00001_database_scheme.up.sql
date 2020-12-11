create table users (
    id              bigserial    primary key ,
    first_name                   varchar           not null,
    last_name                   varchar           not null,
    avatar_uri           varchar,
    credential_id         bigserial,
    email_id        bigserial,
    phone_id        bigserial,
    linked_in_id     bigserial,
    created_at         timestamptz             default 'now()',
    is_confirmed       boolean     default true
);

create table credentials (
    id          bigserial           primary key,
    username    varchar             unique not null,
    password    varchar             not null,
    created_at    timestamptz             default 'now()'
);

create table emails (
    id          bigserial           primary key ,
    address     varchar             unique ,
    email_confirmed_at    timestamptz             default 'now()',
    is_email_confirmed    boolean             default false
);

create table phones (
    id              bigserial           primary key ,
    country_code           varchar             not null,
    number          varchar             not null,
    phone_confirmed_at    timestamptz             default 'now()',
    is_phone_confirmed    boolean             default false
);

create table country_codes (
    country_code        varchar         primary key ,
    country             varchar         not null ,
    continent           varchar
);

create table linked_ins (
    id          bigserial           primary key ,
    login       varchar             unique,
    linked_in_confirmed_at    timestamptz             default 'now()',
    is_linked_in_confirmed    boolean             default false
);

alter table users add foreign key (credential_id) references credentials (id);
alter table users add foreign key (email_id) references emails (id);
alter table users add foreign key (phone_id) references phones (id);
alter table users add foreign key (linked_in_id) references linked_ins (id);

alter table phones add foreign key (country_code) references country_codes (country_code);
