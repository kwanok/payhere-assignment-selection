-- auto-generated definition
create table users
(
    id         char(36)
        primary key,
    name       varchar(64)                        null,
    email      varchar(256)                       not null,
    password   varchar(256)                       not null,
    created_at datetime default CURRENT_TIMESTAMP not null,
    updated_at datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP
);

