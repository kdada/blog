--PostgreSQL
--本文件中的指令只需要执行一次即可

--创建用户
create user blog with password 'blog';

--创建表空间,位置修改为正确的位置
--create tablespace blog owner blog location 'C:\Program Files\PostgreSQL\data\blog';
create tablespace blog owner blog location '/var/lib/pgsql/blog';

--创建数据库
create database blog owner blog tablespace blog encoding utf8;
comment on database blog is 'blog数据库';