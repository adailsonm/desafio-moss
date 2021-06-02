create table order
(
    id serial not null
        constraint pizza_pk
            primary key,
    number_order varchar(100) not null,
    client_name varchar(100) not null,
    address varchar(150) not null,
    estimated_time_of_arrival timestamp not null,
    last_update timestamp default now() not null,
    status int not null,
    pizzas json not null
);