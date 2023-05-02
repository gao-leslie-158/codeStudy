
-- 用户表创建
create table user(
                     id bigint not null auto_increment,
                     user_id bigint not null,
                     username varchar(64) collate utf8mb4_general_ci not null ,
                     password varchar(64) collate utf8mb4_general_ci not null ,
                     email varchar(64) collate utf8mb4_general_ci,
                     gendar tinyint not null default 0,
                     create_time timestamp null default current_timestamp,
                     update_time timestamp null default current_timestamp on update current_timestamp,
                     primary key(id),
                     unique key idx_username (username) using btree ,
                     unique key idx_user_id (user_id) using btree
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 创建社区表
drop table if exists community;
create table community(
                          id bigint not null auto_increment,
                          community_id bigint unsigned not null,
                          community_name varchar(128) collate utf8mb4_general_ci not null,
                          introduction varchar(256) collate utf8mb4_general_ci not null ,
                          create_time timestamp not null default current_timestamp,
                          update_time timestamp not null default current_timestamp on update current_timestamp,
                          primary key(id),
                          unique key idx_community_id (community_id) ,
                          unique key idx_community_name (community_name)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 向community写入数据

insert into community values ('1','1','GO','Golang','2022-11-01 00:10:20','2022-12-12 12:45:56');
insert into community values ('2','2','leetcode','动态规划','2022-10-01 00:10:20','2022-10-12 12:45:56');
insert into community values ('3','3','game','王者荣耀','2021-11-01 00:10:20','2021-12-12 12:45:56');
insert into community values ('4','4','eat','穿越火线','2022-01-01 00:10:20','2022-01-12 12:45:56');
insert into community values ('5','5','life','生活日常','2021-12-01 00:10:20','2021-12-12 12:45:56');


-- 创建 post 帖子信息表
drop table if exists post;
create table post(
    id bigint not null auto_increment,
    post_id bigint not null comment '帖子id',
    title varchar(128) collate utf8mb4_general_ci not null comment '标题',
    content varchar(8192) collate utf8mb4_general_ci not null comment '帖子内容',
    author_id bigint not null comment '坐着的用户id',
    community_id bigint not null comment '所属社区',
    status tinyint not null default '1' comment '帖子状态',
    create_time timestamp null default current_timestamp,
    update_time timestamp null default current_timestamp on update current_timestamp,
    primary key(id),
    unique key idx_post_id (post_id),
    key idx_author_id (author_id),
    key idx_community_id (community_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;