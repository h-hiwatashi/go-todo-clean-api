-- TODO

-- sample
-- create table if not exists todos (
--   todo_id integer unsigned auto_increment primary key,
--   title varchar(100) not null,
--   contents text not null,
--   username varchar(100) not null,
--   nice integer not null,
--   created_at datetime
-- );

-- create table if not exists comments (
--   comment_id integer unsigned auto_increment primary key,
--   todo_id integer unsigned not null,
--   message text not null,
--   created_at datetime,
--   foreign key (todo_id) references todos(todo_id)
-- );

-- insert into todos (title, contents, username, nice, created_at) values
-- ('firstPost', 'This is my first blog', 'user name', 2, now());
-- insert into todos (title, contents, username, nice) values
-- ('2nd', 'Second blog post', 'user name', 4);
-- insert into comments (todo_id, message, created_at) values
-- (1, '1st comment yeah', now());
-- insert into comments (todo_id, message) values
-- (1, 'welcome');