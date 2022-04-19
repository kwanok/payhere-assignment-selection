-- auto-generated definition
create table pays
(
    id         char(36)                             not null
        primary key,
    user_id    char(36)                             not null,
    price      int                                  not null,
    memo       text                                 null,
    removed    tinyint(1) default 0                 not null,
    created_at datetime   default CURRENT_TIMESTAMP null,
    updated_at datetime   default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    removed_at datetime                             null,
    constraint pays_users_id_fk
        foreign key (user_id) references users (id)
);

