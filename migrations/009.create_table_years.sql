CREATE TABLE if not exists Years(
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    name varchar(4),
    description varchar(40)
)