create table IF NOT EXISTS test (
    id serial not null,
    "name" varchar(255) NOT NULL,
    age int,
    PRIMARY KEY (id)
);