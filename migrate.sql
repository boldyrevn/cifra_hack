create table if not exists public."user"
(
    id              serial
        primary key,
    full_name       varchar(64),
    role            varchar(64) default 'employee'::character varying not null,
    balance         integer     default 0                             not null
        constraint positive_balance
            check (balance >= 0),
    email           varchar(64)
        constraint lol
            unique,
    hashed_password varchar(64)                                       not null
);

alter table public."user"
    owner to postgres;

create table if not exists public.zone
(
    id            integer default nextval('table_name_id_seq'::regclass) not null
        constraint table_name_pk
            primary key,
    title         varchar(64)                                            not null
        constraint zones_pk
            unique,
    current_count integer default 0                                      not null,
    max_count     integer default 0                                      not null
);

alter table public.zone
    owner to postgres;

create table if not exists public.user_zone
(
    user_id integer not null,
    zone_id integer not null,
    constraint user_zone_pk
        primary key (user_id, zone_id)
);

alter table public.user_zone
    owner to postgres;

create table if not exists public.user_stat
(
    user_id      integer           not null
        constraint user_stat_pk
            primary key
        constraint user_stat_user_id_fk
            references public."user",
    coffee_cups  integer default 0 not null,
    today_hours  integer default 0 not null,
    company_days integer default 1 not null
);

alter table public.user_stat
    owner to postgres;

create function public.book_zone(zid integer, uid integer) returns void
    language plpgsql
as
$$
BEGIN
    IF (SELECT current_count < max_count FROM zone WHERE zone.id = zid)
    THEN INSERT INTO user_zone VALUES (uid, zid);
    UPDATE zone SET current_count = current_count + 1 WHERE id = zid;
    ELSE RAISE EXCEPTION 'zone is full';
    END IF;
end;
$$;

alter function public.book_zone(integer, integer) owner to postgres;