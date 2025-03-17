create table if not exists todos (
todo_id integer unsigned auto_increment primary key,
title varchar(100) not null,
contents text not null,
user_name varchar(100) not null,
nice_num integer not null,
created_at datetime
);
create table if not exists comments (
comment_id integer unsigned auto_increment primary key,
todo_id integer unsigned not null,
message text not null,
created_at datetime,
foreign key (todo_id) references todos(todo_id)
);
insert into todos (title, contents, user_name, nice_num, created_at) values('first todo','This is my first todo','user name', 2, now());
insert into todos (title, contents, user_name, nice_num) values('2nd','Second todo post','user name', 4);
insert into comments (todo_id, message, created_at) values
(1,'1st comment yeah', now());
insert into comments (todo_id, message) values (1,'welcome');