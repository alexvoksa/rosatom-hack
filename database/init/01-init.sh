#!/usr/bin/env bash
set -e
psql -v ON_ERROR_STOP=1 -U "$PGUSER" -d "$PGDATABASE" <<-EOSQL
  CREATE USER $DB_USER WITH LOGIN PASSWORD '$DB_PASS';
  CREATE DATABASE $DB_NAME;
  GRANT ALL PRIVILEGES ON DATABASE $DB_NAME TO $DB_USER;
  \connect $DB_NAME $DB_USER
  BEGIN;

  CREATE TABLE IF NOT EXISTS classifications (
    id int unique not null GENERATED BY DEFAULT AS IDENTITY (START WITH 1 INCREMENT BY 1) PRIMARY KEY,
    okpo bigint not null,
    oktmo_code bigint not null,
    oktmo_name varchar not null,
  CONSTRAINT classifications_uc UNIQUE (okpo, oktmo_code, oktmo_name)
  );

  CREATE TABLE IF NOT EXISTS suppliers (
    ogrn bigint unique not null PRIMARY KEY,
    name varchar unique not null,
    short_name varchar not null,
    email varchar,
    phone varchar,
    address varchar,
    inn bigint not null,
    kpp bigint not null,
    registered_at TIMESTAMP,
    reg_num varchar not null,
    classification int not null,
    description varchar not null,
    reputation int not null default 0,
    sold_amount float not null default 0.0,
    successful_tenders int not null default 0,
    unsuccessful_tenders int not null default 0,
    is_innovate boolean not null default false,
    CONSTRAINT suppliers_classification_fk
      FOREIGN KEY (classification)
        REFERENCES classifications(id)
          ON DELETE SET NULL
          ON UPDATE CASCADE,
    CONSTRAINT suppliers_uc UNIQUE (name, inn, kpp)
  );

  CREATE TABLE IF NOT EXISTS customers (
    id int unique not null GENERATED BY DEFAULT AS IDENTITY (START WITH 1 INCREMENT BY 1) PRIMARY KEY,
    reg_num varchar not null,
    name varchar unique not null,
    email varchar,
    phone varchar,
    address varchar,
    inn bigint not null,
    kpp bigint not null,
    registered_at TIMESTAMP,
    classification int not null,
    description varchar not null,
    reputation int not null default 0,
    ordered_amount float not null default 0.0,
    successful_tenders int not null default 0,
    unsuccessful_tenders int not null default 0,
    CONSTRAINT customers_classification_fk
      FOREIGN KEY (classification)
        REFERENCES classifications(id)
          ON DELETE SET NULL
          ON UPDATE CASCADE,
    CONSTRAINT customers_uc UNIQUE (name, inn, kpp)
  );

  CREATE TABLE IF NOT EXISTS goods (
    id int unique not null GENERATED BY DEFAULT AS IDENTITY (START WITH 1 INCREMENT BY 1) PRIMARY KEY,
    name varchar,
    code varchar,
    taxes float not null default 0.0,
    price float not null default 0.0,
    currency varchar (3) not null default 'RUB',
    description varchar,
    CONSTRAINT goods_uc UNIQUE (name, code)
  );

  CREATE TABLE IF NOT EXISTS org_codes (
    id int unique not null GENERATED BY DEFAULT AS IDENTITY (START WITH 1 INCREMENT BY 1) PRIMARY KEY,
    name varchar,
    code varchar,
    CONSTRAINT org_codes_uc UNIQUE (name, code)
  );

  CREATE TABLE IF NOT EXISTS units (
    id int unique not null GENERATED BY DEFAULT AS IDENTITY (START WITH 1 INCREMENT BY 1) PRIMARY KEY,
    name varchar,
    code varchar,
    national_name varchar,
    international_name varchar,
    national_code varchar,
    international_code varchar,
    CONSTRAINT units_uc UNIQUE (name, code, national_name, international_name, national_code, international_code)
  );

  CREATE TABLE IF NOT EXISTS tenders (
    id varchar unique not null PRIMARY KEY,
    reg_num varchar not null,
    name varchar not null,
    tender_goods varchar not null,
    resolution int not null,
    winner bigint,
    price float not null default 0.0,
    name varchar (3) not null default 'RUB',
    customer int not null,
    published_at TIMESTAMP not null,
    url varchar not null,
    CONSTRAINT winning_supplier_fk
     FOREIGN KEY(winner)
       REFERENCES suppliers(ogrn)
         ON DELETE SET NULL
         ON UPDATE CASCADE,
    CONSTRAINT customer_fk
     FOREIGN KEY(customer)
       REFERENCES customers(id)
         ON DELETE SET NULL
         ON UPDATE CASCADE
  );
  CREATE TABLE IF NOT EXISTS tender_suppliers (
    id int unique not null GENERATED BY DEFAULT AS IDENTITY (START WITH 1 INCREMENT BY 1) PRIMARY KEY,
    tender_id varchar,
    supplier_id bigint,

    CONSTRAINT tender_tender_fk
     FOREIGN KEY(tender_id)
       REFERENCES tenders(id)
         ON DELETE SET NULL
         ON UPDATE CASCADE,

    CONSTRAINT tender_supplier_fk
     FOREIGN KEY(supplier_id)
       REFERENCES suppliers(ogrn)
         ON DELETE SET NULL
         ON UPDATE CASCADE
  );

  CREATE TABLE IF NOT EXISTS tender_goods (
    id int unique not null GENERATED BY DEFAULT AS IDENTITY (START WITH 1 INCREMENT BY 1) PRIMARY KEY,
    tender_id varchar not null,
    goods_id int not null,
    CONSTRAINT tender_id_fk
      FOREIGN KEY(tender_id)
        REFERENCES tenders(id)
          ON DELETE SET NULL
          ON UPDATE CASCADE,
    CONSTRAINT goods_id_fk
      FOREIGN KEY(goods_id)
        REFERENCES goods(id)
          ON DELETE SET NULL
          ON UPDATE CASCADE
  );

  CREATE TABLE IF NOT EXISTS supplier_goods (
    id int unique not null GENERATED BY DEFAULT AS IDENTITY (START WITH 1 INCREMENT BY 1) PRIMARY KEY,
    supplier_id bigint not null,
    goods_id int not null,
    goods_amount float not null,
    CONSTRAINT goods_id_supplier_goods_fk
      FOREIGN KEY(goods_id)
        REFERENCES goods(id)
          ON DELETE SET NULL
          ON UPDATE CASCADE,
    CONSTRAINT supplier_id_supplier_goods_fk
      FOREIGN KEY(supplier_id)
        REFERENCES suppliers(ogrn)
          ON DELETE SET NULL
          ON UPDATE CASCADE
  );

  CREATE TABLE IF NOT EXISTS customer_goods (
    id int unique not null GENERATED BY DEFAULT AS IDENTITY (START WITH 1 INCREMENT BY 1) PRIMARY KEY,
    customer_id int not null,
    goods_id int not null,
    goods_amount float not null,
    CONSTRAINT goods_id_customer_goods_fk
      FOREIGN KEY(goods_id)
        REFERENCES goods(id)
          ON DELETE SET NULL
          ON UPDATE CASCADE,
    CONSTRAINT customer_id_customer_goods_fk
      FOREIGN KEY(customer_id)
        REFERENCES customers(id)
          ON DELETE SET NULL
          ON UPDATE CASCADE
  );

  INSERT INTO goods (code, name)
  VALUES
    $(cat ../constant_files/okpd2.txt)
  ON CONFLICT DO NOTHING
  ;
  INSERT INTO org_codes (code, name)
  VALUES
    $(cat ../constant_files/okved2.txt)
  ON CONFLICT DO NOTHING
  ;
  INSERT INTO units (code, name, national_name, international_name, national_code, international_code)
  VALUES
    $(cat ../constant_files/okei2.txt)
  ON CONFLICT DO NOTHING
  ;
  COMMIT;
EOSQL