INSERT INTO customers (reg_num, name, email, phone, address, inn, kpp, registered_at, description,reputation, ordered_amount, successful_tenders, unsuccessful_tenders)
values (?,?,?,?,?,?,?,?,?,?,?,?,?) ON CONFLICT (reg_num)
DO UPDATE SET
successful_tenders = customers.successful_tenders + ?,
unsuccessful_tenders = customers.unsuccessful_tenders + ?;

INSERT INTO suppliers (ogrn, name, short_name, email,phone, address, inn, kpp, registered_at, reg_num, classification,
description, reputation, sold_amount, successful_tenders, unsuccessful_tenders, is_innovate) 
values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON CONFLICT (ogrn)
DO UPDATE SET
successfull_tenders = suppliers.successfull_tenders + ?,
unsuccessfull_tenders = suppliers.unsuccessfull_tenders + ?;

INSERT INTO goods (name, code, taxes, price,currency, description)
values (?,?,?,?,?,?) ON CONFLICT DO NOTHING;

INSERT INTO org_codes (name, code) values (?,?) ON CONFLICT DO NOTHING;

INSERT INTO units (name, code, national_name, international_name, national_code, international_code)
values (?,?,?,?,?,?) ON CONFLICT DO NOTHING;

INSERT INTO tenders (id, reg_num, name, tender_goods, resolution, winner, price, currency, published_at, url)
values (?,?,?,?,?,?,?,?,?,?) ON CONFLICT DO NOTHING;

INSERT INTO tender_suppliers (tender_id, supplier_id) values (?,?) ON CONFLICT DO NOTHING;

INSERT INTO tender_goods  (tender_id, goods_id) values (?,?) ON CONFLICT DO NOTHING;

INSERT INTO supplier_goods  (supplier_id , goods_id, goods_amount) values (?,?,?) ON CONFLICT DO NOTHING;

INSERT INTO customer_goods (customer_id, goods_id, goods_amount) values (?,?,?) ON CONFLICT DO NOTHING;

