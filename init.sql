CREATE TABLE if not exists Users (
    id    integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    name  varchar(30) NOT NULL
);