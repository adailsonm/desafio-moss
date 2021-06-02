create table pizza
(
    id          serial           not null
        constraint pizza_pk_id
            primary key,
    name        varchar(150)     not null,
    price       double precision not null,
    ingredients json
);

alter table pizza
    owner to postgres;