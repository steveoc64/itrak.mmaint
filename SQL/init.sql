drop table if exists roles;
create table roles (
	id integer primary key,
	name text
);
insert into roles (id,name) values 
	(1,'Staff'),
	(2,'Site Manager'),
	(3,'Vendor'),
	(100,'Admin');
select * from roles;

drop table if exists users;
create table users (
	id serial primary key,
	username varchar(32),
	passwd varchar(32),
	role integer
);
delete from users;
insert into users (username,passwd,role) values 
	('worker1','abc123',1),
	('worker2','abc123',1),
	('worker3','abc123',1),
	('sitemgr1','abc123',2),
	('sitemgr2','abc123',2),
	('sitemgr3','abc123',2),
	('subbie1','abc123',3),
	('subbie2','abc123',3),
	('subbie3','abc123',3),
	('admin','abc123',100);
select * from users;

drop table if exists sites;
drop table if exists site;
create table site (
	id serial primary key,
	name varchar(32),
	address text,
	phone text,
	contactname text,
	lat numeric(12,10),
	lon numeric(12,10)
);
insert into site (name) values 
	('R&D Workshop'),
	('SBS Edinburgh'),
	('SBS Newcastle'),
	('SBS Minto'),
	('SBS Tomago'),
	('SBS Chinderah'),
	('SBS Victoria'),
	('Thermoloc'),
	('Fab Shop'),
	('USA Connecticut');
select * from site;


drop table if exists person;
create table person (
	id serial primary key,
	user_id integer,
	name text,
	email text,
	phone text,
	hourlyrate integer,
	comments text,
	alternate text,
	role integer,
	location integer,
	calendar integer
);
insert into person (name,email,phone,hourlyrate,comments,alternate,role) 
	select personname,email,phone,hourlyrate,comments,alternate,1 
	from fm_person;
select * from person;

drop table if exists equip_type;
create table equip_type (
	id serial primary key,
	name text,
	is_consumable boolean,
	is_asset boolean
);
insert into equip_type (name,is_consumable,is_asset) values 
	('Consumables',TRUE,FALSE),
	('Mech',FALSE,TRUE),
	('Spare Parts',FALSE,FALSE);
select * from equip_type;


drop table if exists equipment;
create table equipment (
	id serial primary key,
	name text,
	descr text,
	comments text,
	modelno text,
	serialno text,
	location integer,
	parent_id integer,
	category integer,
	calendar integer,
	vendor integer
);
insert into equipment (name,descr,comments,modelno,serialno)
	select name,descr,comments,modelno,serialno
	from fm_equipment;
select * from equipment;

drop table if exists vendor_rating;
create table vendor_rating (
	id serial primary key,
	name text
);
insert into vendor_rating (id,name) values 
	(1,'Good'),
	(2,'Average'),
	(3,'Poor');
select * from vendor_rating;

drop table if exists vendor;
create table vendor (
	id serial primary key,
	name text,
	descr text,
	comments text,
	account text,
	maincontact text,
	servicecontact text,
	partscontact text,
	othercontact text,
	rating int
);
insert into vendor (name,descr,comments,account,maincontact,servicecontact,partscontact,othercontact)
	select name,descr,comments,account,maincontact,servicecontact,partscontact,othercontact
	from fm_vendor;
select * from vendor;






);




