create table if not exists gas_prices (
  id serial primary key,
  timestamp timestamp not null,
  gas_price integer not null
);

drop table if exists gas_prices cascade;