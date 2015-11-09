drop table if exists users;
create table users (
	username varchar(32),
	passwd varchar(32)
);
delete from users;
insert into users (username,passwd) values ('worker1','abc123');
insert into users (username,passwd) values ('worker2','abc123');
insert into users (username,passwd) values ('worker3','abc123');
insert into users (username,passwd) values ('sitemgr1','abc123');
insert into users (username,passwd) values ('sitemgr2','abc123');
insert into users (username,passwd) values ('sitemgr3','abc123');
insert into users (username,passwd) values ('subbie1','abc123');
insert into users (username,passwd) values ('subbie2','abc123');
insert into users (username,passwd) values ('subbie3','abc123');
insert into users (username,passwd) values ('admin','abc123');

