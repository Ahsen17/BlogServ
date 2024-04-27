SET NAMES utf8mb4;

DROP DATABASE IF EXISTS blogserv;
CREATE DATABASE blogserv;
USE blogserv;

CREATE TABLE `account` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `username` VARCHAR(16) NOT NULL COMMENT '用户账户',
    `password` VARCHAR(32) NOT NULL COMMENT '账户密码',
    `status` tinyint(9) NOT NULL DEFAULT 1 COMMENT '账户状态',
    `create_at` BIGINT NOT NULL DEFAULT 0 COMMENT '创建时间',
    `create_by` VARCHAR(16) NOT NULL COMMENT '创建账户人',
    `update_at` BIGINT NOT NULL DEFAULT 0 COMMENT '更新时间',
    `update_by` VARCHAR(16) NOT NULL COMMENT '修改账户人',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`username`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT '账户表';

INSERT INTO `account`(id, username, password, status, create_at, create_by, update_at, update_by) values (1, 'admin', 'admin', 1, unix_timestamp(now()), 'system', unix_timestamp(now()), 'system');


