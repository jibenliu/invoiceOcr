CREATE TABLE `p_tenant`
(
    `id`            int(11)       NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `tenant_code`   varchar(100)  NOT NULL COMMENT '租户代码',
    `tenant_name`   varchar(200)  NOT NULL COMMENT '租户名称',
    `contact_name`  varchar(50)   NOT NULL DEFAULT '' COMMENT '联系人',
    `contact_email` varchar(100)  NOT NULL DEFAULT '' COMMENT '联系人邮箱',
    `contact_phone` varchar(20)   NOT NULL DEFAULT '' COMMENT '联系人手机',
    `status`        tinyint(1)    NOT NULL DEFAULT 0 COMMENT '状态码',
    `app_id`        varchar(100)  NOT NULL DEFAULT '' COMMENT 'app_id',
    `public_key`    varchar(1000) NOT NULL DEFAULT '' COMMENT 'public_key',
    `secret_key`    varchar(1000) NOT NULL DEFAULT '' COMMENT 'secret_key',
    `created_at`    int(11)       NOT NULL DEFAULT 0 COMMENT '创建时间',
    `updated_at`    int(11)       NOT NULL DEFAULT 0 COMMENT '最后更新时间',
    PRIMARY KEY (`id`, `tenant_code`, `tenant_name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='租户表';


CREATE TABLE `p_tenant_status`
(
    `id`                    int(11)       NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `tenant_id`             varchar(100)  NOT NULL COMMENT '租户代码',
    `tenant_code`           varchar(100)  NOT NULL COMMENT '租户代码',
    `api_code`              varchar(50)   NOT NULL COMMENT '接口代码',
    `api_charge_status`     tinyint(1)    NOT NULL DEFAULT 0 COMMENT '接口状态',
    `api_charge_count`      int(10)       NOT NULL DEFAULT 0 COMMENT '充值次数',
    `api_charge_left_count` int(10)       NOT NULL DEFAULT 0 COMMENT '接口剩余次数',
    `app_id`                varchar(100)  NOT NULL DEFAULT '' COMMENT 'app_id',
    `public_key`            varchar(1000) NOT NULL DEFAULT '' COMMENT 'public_key',
    `secret_key`            varchar(1000) NOT NULL DEFAULT '' COMMENT 'secret_key',
    `created_at`            int(11)       NOT NULL DEFAULT 0 COMMENT '创建时间',
    `updated_at`            int(11)       NOT NULL DEFAULT 0 COMMENT '最后更新时间',
    PRIMARY KEY (`id`),
    INDEX `idx_tenant_id` (`tenant_id`) USING BTREE,
    INDEX `idx_tenant_code` (`tenant_code`) USING BTREE,
    INDEX `idx_api_code` (`api_code`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='租户表';