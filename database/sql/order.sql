create table "order"
(
    id serial not null,
    number_order varchar(100),
    client_name varchar not null,
    address text not null,
    pizzas json not null,
    estimated_time_of_arrival timestamp not null,
    last_update date default now() not null,
    status int default 1 not null
);

