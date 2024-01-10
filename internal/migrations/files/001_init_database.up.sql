create table if not exists users
(
    id             uuid                     not null
        constraint users_pk primary key,
    web3Address    text                     not null,
    first_name     text                              default null,
    last_name      text                              default null,
    username       varchar(255)                      default null,
    deactivated_at timestamp with time zone          default null,
    created_at     timestamp with time zone not null default now(),
    updated_at     timestamp with time zone not null default now()
);

create unique index web3Address_uix on users (web3Address);
create unique index username_uix on users (username);