create table gas_prices (
  id serial primary key,
  timestamp timestamp not null,
  gas_price bigint not null
);

---- create above / drop below ----

drop table if exists gas_prices cascade;