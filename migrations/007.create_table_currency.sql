CREATE TABLE if not exists Currency(
    id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    name varchar (10) unique,
    description varchar(100)
)