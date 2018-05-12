
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table Apis (
  id    serial PRIMARY KEY,
  name  varchar(120),
  path  varchar(120),
  method  varchar(10),
  access_path  varchar(120)
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE Apis;
