create table pizza
(
    id serial not null
        constraint pizza_pk
            primary key,
    name varchar(150) not null,
    ingredients json not null,
    price float8 not null
);