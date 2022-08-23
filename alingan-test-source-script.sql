-- script for repository and query test
-- run before repository test

-- Create The Owner
insert into core."owner"
(owner_id, owner_name, owner_type, owner_email, "password", is_active, created_date, last_modified)
values('owner-001', 'Toko Prima', 'organization', 'tokoprima@gmail.co.id', '$2a$14$SgHKgXnfxVOA.fHMN3yOSORkJq7J5XcboPm9mBikPKXzaKPNL2.o.', true, now(), now());

-- Create Store for "owner-001"
insert into core.store
(store_id, owner_id, store_name, store_address, is_active, created_date, last_modified)
values('str-001', 'owner-001', 'Toko Prima 1 - Salemba', 'Jakarta, Indonesia', true, now(), now());

insert into core.store
(store_id, owner_id, store_name, store_address, is_active, created_date, last_modified)
values('str-002', 'owner-001', 'Toko Prima 2 - Kramat Raya', 'Jakarta, Indonesia', true, now(), now());

-- Create Product for "owner-001"
insert into core.product
(product_id, owner_id, product_name, product_measurement_unit, product_price, is_active, created_date, last_modified)
values('prd-001', 'owner-001', 'Indomie Goreng', 'bungkus', 2500, true, now(), now());

insert into core.product
(product_id, owner_id, product_name, product_measurement_unit, product_price, is_active, created_date, last_modified)
values('prd-002', 'owner-001', 'Telur Ayam', 'bungkus', 2500, true, now(), now());

-- Create Agent for "owner-001" store "str-001"
insert into core.agent
(agent_id, store_id, agent_name, agent_email, agent_password, is_active, created_date, last_modified)
values('agent-001', 'str-001', 'Budi', 'budi@gmial.com', 'budi123', true, now(), now());

-- Create Transaction
insert into core."transaction"
(transaction_id, transaction_date, agent_id, store_id, transaction_total, created_date, last_modified)
values('trx-001', now(), 'agent-001', 'str-001', 5000, now(), now());

insert into core."transaction"
(transaction_id, transaction_date, agent_id, store_id, transaction_total, created_date, last_modified)
values('trx-002', now(), 'agent-001', 'str-001', 10000, now(), now());

-- Create Transaction item for "trx-001"
insert into core.transaction_item
(transaction_item_id, product_id, transaction_id, used_price, buy_quantity, created_date, last_modified)
values('trx-item-001', 'prd-001', 'trx-001', 2500, 2, now(), now());

insert into core.transaction_item
(transaction_item_id, product_id, transaction_id, used_price, buy_quantity, created_date, last_modified)
values('trx-item-002', 'prd-002', 'trx-002', 2500, 4, now(), now());

-- Product Image for prd-001 and prd-002
INSERT INTO core.product_image (product_image_id, product_id, location_path, created_date, last_modified) 
VALUES('prd-img-001', 'prd-001', 'https://via.placeholder.com/300', now(), now());

INSERT INTO core.product_image (product_image_id, product_id, location_path, created_date, last_modified) 
VALUES('prd-img-002', 'prd-002', 'https://via.placeholder.com/300', now(), now());