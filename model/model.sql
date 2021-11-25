-- we don't know how to generate root <with-no-name> (class Root) :(
create table Project
(
    Id INTEGER
        constraint Project_pk
            primary key autoincrement,
    Name text not null,
    Description text
);

create table Category
(
    Id INTEGER
        constraint Category_pk
            primary key autoincrement,
    Name TEXT not null,
    Color TEXT not null,
    ProjectId INTEGER not null
        references Project
            on update cascade on delete cascade
);

create unique index Category_Color_uindex
    on Category (Color);

create unique index Category_Name_uindex
    on Category (Name);

create unique index Project_Name_uindex
    on Project (Name);

create table User
(
    Id INTEGER
        constraint table_name_pk
            primary key autoincrement,
    Name text not null,
    Password text not null,
    Role INTEGER default 0 not null
);

create table Payment
(
    Id INTEGER
        constraint Payment_pk
            primary key autoincrement,
    Name text not null,
    Amount INTEGER not null,
    Date text not null,
    ProjectId INTEGER not null,
    CategoryId INTEGER
        references Category
                  on update set null on delete set null,
    PayerId INTEGER
        references User
            on update cascade on delete cascade,
    PayeeId INTEGER
        references User
            on update cascade on delete cascade
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

