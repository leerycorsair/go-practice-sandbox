create table if not exists orders (
    order_uid varchar(255),
    track_number varchar(255),
    entry varchar(50),
    locale varchar(50),
    internal_signature varchar(255),
    customer_id varchar(255),
    delivery_service varchar(255),
    shardkey varchar(10),
    sm_id integer,
    date_created timestamp,
    oof_shard varchar(10),
    primary key (order_uid)
);

create table if not exists payments (
    order_uid varchar(255),
    transaction varchar(255),
    request_id varchar(255),
    currency varchar(10),
    provider varchar(255),
    amount integer,
    payment_dt integer,
    bank varchar(255),
    delivery_cost integer,
    goods_total integer,
    custom_fee integer,
    primary key (order_uid),
    foreign key (order_uid) references orders(order_uid) on delete cascade
);

create table if not exists items (
    item_id serial,
    order_uid varchar(255),
    chrt_id integer,
    track_number varchar(255),
    price integer,
    rid varchar(255),
    name varchar(255),
    sale integer,
    size varchar(25),
    total_price integer,
    nm_id integer,
    brand varchar(255),
    status integer,
    primary key (item_id),
    foreign key (order_uid) references orders(order_uid) on delete cascade
);

create table if not exists deliveries (
    order_uid varchar(255),
    name varchar(255),
    phone varchar(20),
    zip varchar(255),
    city varchar(255),
    address varchar(255),
    region varchar(255),
    email varchar(255),
    primary key (order_uid),
    foreign key (order_uid) references orders(order_uid) on delete cascade
);