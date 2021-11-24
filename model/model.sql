create table Project
(
    Id INTEGER
        constraint Project_pk
            primary key autoincrement,
    Name text not null,
    Description text
);

create unique index Project_Name_uindex
    on Project (Name);

create table Tag
(
    Id INTEGER
        constraint Tag_pk
            primary key autoincrement,
    Name text not null
);

create unique index Tag_Name_uindex
    on Tag (Name);

create table User
(
    Id INTEGER
        constraint table_name_pk
            primary key autoincrement,
    Name text not null,
    Password text not null,
    Role INTEGER default 0 not null
);

create table ProjectUser
(
    Id INTEGER
        constraint ProjectUser_pk
            primary key autoincrement,
    ProjectId INTEGER not null
        references Project
            on update cascade on delete cascade,
    UserId INTEGER not null
        references User
            on update cascade on delete cascade
);

create unique index table_name_username_uindex
    on User (Name);
