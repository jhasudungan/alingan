-- DDL for Alingan
-- Run this to create schema

create schema core;

create table core.owner (
	owner_id varchar(200) unique not null primary key,
	owner_name varchar (200) not null,
	owner_type varchar (20) not null,
	owner_email varchar (200) unique not null,
	password varchar (200) not null,
	is_active boolean not null,
	created_date timestamp not null,
	last_modified timestamp not null
);

create table core.store (
	store_id varchar(200) unique not null primary key,
	owner_id varchar(200) not null,
	store_name varchar(200) not null,
	store_address varchar(200) not null,
	is_active boolean not null,
	created_date timestamp not null,
	last_modified timestamp not null
);

create table core.agent (
	agent_id varchar(200) unique not null primary key,
	store_id varchar(200) not null,
	agent_name varchar(200) not null,
	agent_email varchar(200) not null unique,
	agent_password varchar(200) not null,
	is_active boolean not null,
	created_date timestamp not null,
	last_modified timestamp not null
);

create table core.product (
	product_id varchar(200) unique not null primary key,
	owner_id varchar(200) not null,
	product_name varchar(200) not null,
	product_measurement_unit varchar(200) not null,
	product_price numeric not null,
	is_active boolean not null,
	created_date timestamp not null,
	last_modified timestamp not null
);

create table core.transaction (
	transaction_id varchar(200) unique not null primary key,
	transaction_date timestamp not null,
	agent_id varchar(200) not null,
	store_id varchar(200) not null,
	transaction_total numeric not null,
	created_date timestamp not null,
	last_modified timestamp not null
);

create table core.transaction_item (
	transaction_item_id varchar(200) unique not null primary key,
	product_id varchar(200) not null,
	transaction_id varchar(200) not null,
	used_price numeric not null,
	buy_quantity numeric not null,
	created_date timestamp not null,
	last_modified timestamp not null
);

create table core.product_image (
	product_image_id varchar(200) unique not null primary key,
	product_id varchar(200) not null,
	location_path varchar(200) not null,
	created_date timestamp not null,
	last_modified timestamp not null
);

create table core.transaction_receipt (
	transaction_receipt_id varchar(200) unique not null primary key,
	transaction_id varchar(200) not null,
	location_path varchar(200) not null,
	created_date timestamp not null,
	last_modified timestamp not null
);
	

































