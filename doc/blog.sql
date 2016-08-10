--需要切换到blog数据库,并且使用blog用户登录

--创建blog模式
drop schema if exists blog;
create schema blog authorization blog;
comment on schema blog is 'blog模式';

--设置查询路径
set search_path to '$user',public;

--创建账号表
drop table if exists blog.account;
create table blog.account (
    id serial primary key not null,
    email varchar(100) unique not null,
    name varchar(10) unique not null,
    password char(32) not null,
    salt varchar(32) not null,
    create_time timestamp with time zone not null default now(),
    status integer not null default 1,
    reason varchar(1000) not null default ''
);
--表和字段备注
comment on table blog.account is '账号表';
comment on column blog.account.id is '用户id';
comment on column blog.account.email is '用户邮箱';
comment on column blog.account.name is '用户昵称';
comment on column blog.account.password is '用户密码';
comment on column blog.account.salt is '干扰码';
comment on column blog.account.create_time is '创建时间';
comment on column blog.account.status is '状态码:1-正常,2-禁止登录';
comment on column blog.account.reason is '处于当前状态的原因';

--创建分类表
drop table if exists blog.category;
create table blog.category (
    id serial primary key not null,
    name varchar(100),
    create_time timestamp with time zone not null default now(),
    status integer not null default 1
);
--表和字段备注
comment on table blog.category is '分类表';
comment on column blog.category.id is '分类id';
comment on column blog.category.name is '分类名称';
comment on column blog.category.create_time is '创建时间';
comment on column blog.category.status is '状态码:1-正常,2-隐藏,3-删除';

--创建文章表
drop table if exists blog.article;
create table blog.article (
    id serial primary key not null,
    category integer not null,
    title varchar(1000) not null,
    content text not null,
    top integer not null default 0,
    create_time timestamp with time zone not null default now(),
    update_time timestamp with time zone not null default now(),
    status integer not null default 1
);
--表和字段备注
comment on table blog.article is '文章表';
comment on column blog.article.id is '文章id';
comment on column blog.article.title is '文章标题';
comment on column blog.article.content is '文章内容';
comment on column blog.article.top is '置顶数';
comment on column blog.article.create_time is '创建时间';
comment on column blog.article.update_time is '更新时间';
comment on column blog.article.status is '状态码:1-正常,2-隐藏,3-删除';

--创建回复表
drop table if exists blog.reply;
create table blog.reply (
    id serial primary key not null,
    article integer not null,
    account integer not null,
    reply integer not null,
    content text not null,
    create_time timestamp with time zone not null default now(),
    status integer not null  default 1
);
--表和字段备注
comment on table blog.reply is '回复表';
comment on column blog.reply.id is '回复id';
comment on column blog.reply.article is '文章id';
comment on column blog.reply.account is '用户id';
comment on column blog.reply.reply is '回复的某个回复id';
comment on column blog.reply.content is '回复内容';
comment on column blog.reply.create_time is '回复时间';
comment on column blog.reply.status is '状态码:1-正常,2-已隐藏,3-已屏蔽,4-已删除';

