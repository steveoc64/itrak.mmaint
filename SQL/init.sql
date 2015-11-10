drop table if exists users;
create table users (
	username varchar(32),
	passwd varchar(32),
	role varchar(16)
);
delete from users;
insert into users (username,passwd,role) values ('worker1','abc123','Worker');
insert into users (username,passwd,role) values ('worker2','abc123','Worker');
insert into users (username,passwd,role) values ('worker3','abc123','Worker');
insert into users (username,passwd,role) values ('sitemgr1','abc123','SiteMgr');
insert into users (username,passwd,role) values ('sitemgr2','abc123','SiteMgr');
insert into users (username,passwd,role) values ('sitemgr3','abc123','SiteMgr');
insert into users (username,passwd,role) values ('subbie1','abc123','Vendor');
insert into users (username,passwd,role) values ('subbie2','abc123','Vendor');
insert into users (username,passwd,role) values ('subbie3','abc123','Vendor');
insert into users (username,passwd,role) values ('admin','abc123','Admin');

drop table if exists sites;
create table sites (
	sitename varchar(32)
	address text
	phone text
	contactname text
	lat numeric(12,10)
	lon numeric(12,10)
);
