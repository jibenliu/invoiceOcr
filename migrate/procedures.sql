DELIMITER $$
DROP PROCEDURE IF EXISTS get_table_by_name $$
CREATE PROCEDURE get_table_by_name(IN table_name VARCHAR(255), OUT i INT(1))
BEGIN
    SET @select_sql =
            CONCAT('SELECT IF(EXISTS(SELECT * FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA="', table_name,
                   '" OR TABLE_NAME="', table_name, '"), 1, 0) INTO @is_exists FROM INFORMATION_SCHEMA.TABLES LIMIT 1');
    PREPARE stmt FROM @select_sql;
    EXECUTE stmt;
    DEALLOCATE PREPARE stmt;
    SET i = @is_exists;
END $$
DELIMITER ;

/**---------------------------------通用表起始---------------------------------**/
DELIMITER $$
DROP PROCEDURE IF EXISTS create_tenant_api_charge_log_table_by_week $$
CREATE PROCEDURE create_tenant_api_charge_log_table_by_week(IN week VARCHAR(255))
BEGIN
    DECLARE table_exists INT DEFAULT 0;
    CALL get_table_by_name(CONCAT('p_tenant_api_charge_log_', week), table_exists);
    IF !table_exists THEN
        SET @create_sql = CONCAT('CREATE TABLE p_tenant_api_charge_log_', week, '(
            `id`          int(11)     NOT NULL AUTO_INCREMENT COMMENT "主键",
            `tenant_id`   varchar(50) NOT NULL COMMENT "租户id",
            `tenant_code` varchar(50) NOT NULL COMMENT "租户代码",
            `api_code`    varchar(50) NOT NULL DEFAULT "" COMMENT "接口类型",
            `api_count`   int(1)      NOT NULL DEFAULT 0 COMMENT "耗用次数",
            `created_at`  int(11)     NOT NULL COMMENT "创建时间",
            PRIMARY KEY (`id`),
            INDEX `idx_tenant_id` (`tenant_id`) USING BTREE,
            INDEX `idx_tenant_code` (`tenant_code`) USING BTREE
            ) ENGINE = InnoDB
          DEFAULT CHARSET = utf8 COMMENT ="租户接口消耗记录表"');
        PREPARE stmt FROM @create_sql;
        EXECUTE stmt;
    END IF;
END $$

DELIMITER $$
DROP PROCEDURE IF EXISTS create_api_runtime_log_table_by_week $$
CREATE PROCEDURE create_api_runtime_log_table_by_week(IN week VARCHAR(255))
BEGIN
    DECLARE table_exists INT DEFAULT 0;
    CALL get_table_by_name(CONCAT('p_api_runtime_log_', week), table_exists);
    IF !table_exists THEN
        SET @create_sql = CONCAT('CREATE TABLE p_api_runtime_log_', week, '(
        `id`                     int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT "主键id",
        `tenant_code`            varchar(50)      NOT NULL DEFAULT "" COMMENT "租户代码",
        `tenant_name`            varchar(200)     NOT NULL DEFAULT "" COMMENT "租户名称",
        `log_id`                 varchar(50)      NOT NULL DEFAULT "" COMMENT "日志序列号",
        `channel`                varchar(50)      NOT NULL DEFAULT "" COMMENT "渠道",
        `tag`                    varchar(50)      NOT NULL DEFAULT "" COMMENT "标签",
        `request_params`         longtext COMMENT "请求参数",
        `encrypt_request_params` longtext COMMENT "加密后请求参数",
        `url`                    varchar(1000)    NOT NULL DEFAULT "" COMMENT "请求地址",
        `response_data`          longtext COMMENT "响应内容",
        `decrypt_data`           longtext COMMENT "解密后响应内容",
        `extra`                  longtext COMMENT "其他",
        `level`                  tinyint(1)       NOT NULL DEFAULT "0" COMMENT "日志级别",
        `log_time`               int(11)          NOT NULL DEFAULT "0" COMMENT "日志时间",
        `created_at`             int(11)          NOT NULL DEFAULT "0" COMMENT "创建时间",
        PRIMARY KEY (`id`),
        INDEX `idx_tenant_code` (`tenant_code`),
        INDEX `idx_log_id` (`log_id`),
        INDEX `idx_channel` (`channel`),
        INDEX `idx_tag` (`tag`),
        INDEX `idx_level` (`level`),
        INDEX `idx_log_time` (`log_time`)
    ) ENGINE = InnoDB
      DEFAULT CHARSET = utf8 COMMENT ="接口日志表";');
        PREPARE stmt FROM @create_sql;
        EXECUTE stmt;
    END IF;
END $$
/**---------------------------------通用表起始---------------------------------**/


/**---------------------------------查验表起始---------------------------------**/
DELIMITER $$
DROP PROCEDURE IF EXISTS create_image_ocr $$
CREATE PROCEDURE create_image_ocr(IN n INT(10))
BEGIN
    DECLARE i, table_exists INT DEFAULT 0;
    WHILE i < n
        DO
            CALL get_table_by_name(CONCAT('y_invoice_image_', i), table_exists);
            if !table_exists THEN
                SET @create_sql = CONCAT('CREATE TABLE y_invoice_image_', i, '(
					  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT "自增id",
					  `image_url` varchar(500) NOT NULL COMMENT "图片地址",
					  `ocr_status` tinyint(1) NOT NULL COMMENT "ocr状态码",
					  `tenant_code` varchar(50) NOT NULL COMMENT "创建租户代码",
					  `created_at` int(11) NOT NULL COMMENT "创建时间",
					  `updated_at` int(11) NOT NULL COMMENT "最后更新时间",
					  PRIMARY KEY (`id`),
					  INDEX `idx_image_url` (`image_url`(255)) USING BTREE,
					  INDEX `idx_tenant_code` (`tenant_code`) USING BTREE
					) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8');
                PREPARE stmt FROM @create_sql;
                EXECUTE stmt;
            END IF;
            SET i = i + 1;
        END WHILE;
END $$
DELIMITER ;

CALL create_image_ocr(32);
DROP PROCEDURE IF EXISTS create_image_ocr;

DELIMITER $$
DROP PROCEDURE IF EXISTS create_image_ocr_details $$
CREATE PROCEDURE create_image_ocr_details(IN n INT(10))
BEGIN
    DECLARE i, table_exists INT DEFAULT 0;
    WHILE i < n
        DO
            CALL get_table_by_name(CONCAT('y_invoice_image_', 'detail_', i), table_exists);
            if !table_exists THEN
                SET @create_sql = CONCAT('CREATE TABLE y_invoice_image_detail_', i, '(
					`id` int(11) NOT NULL AUTO_INCREMENT COMMENT "自增id",
					`image_id` int(11) NOT NULL COMMENT "图片id",
					`invoice_type` varchar(50) NULL COMMENT "发票类别",
					`ocr_content` varchar(6000) NULL COMMENT "OCR结构体",
					PRIMARY KEY (`id`),
					INDEX `idx_image_url`(`image_id`) USING BTREE
				)');
                PREPARE stmt FROM @create_sql;
                EXECUTE stmt;
            END IF;
            SET i = i + 1;
        END WHILE;
END $$
DELIMITER ;

CALL create_image_ocr_details(32);
DROP PROCEDURE IF EXISTS create_image_ocr_details;


DELIMITER $$
DROP PROCEDURE IF EXISTS create_invoice_table_by_month $$
CREATE PROCEDURE create_invoice_table_by_month(IN monthStr VARCHAR(255))
BEGIN
    DECLARE table_exists INT DEFAULT 0;
    CALL get_table_by_name(CONCAT('y_invoice_', monthStr), table_exists);
    IF !table_exists THEN
        SET @create_sql = CONCAT('CREATE TABLE y_invoice_', monthStr, '(
			`id` int(11) NOT NULL AUTO_INCREMENT COMMENT "主键",
			  `invoice_type` varchar(5) DEFAULT "" COMMENT "发票类型(01-增值税专用发票 03-机动车发票 04-增值税普通发票 10-电子发票 11-卷式发票 14-电子发票通行费)",
			  `tenant_code` varchar(50) DEFAULT "" COMMENT "识别租户代码",
			  `invoice_no` varchar(50) DEFAULT "" COMMENT "发票号码",
			  `invoice_code` varchar(50) DEFAULT "" COMMENT "发票代码",
			  `seller_taxno` varchar(50) DEFAULT "" COMMENT "销方税号",
			  `seller_name` varchar(150) DEFAULT "" COMMENT "销方名称",
			  `seller_address` varchar(350) DEFAULT "" COMMENT "销方地址",
			  `seller_tel` varchar(50) DEFAULT "" COMMENT "销方电话",
			  `seller_bank_name` varchar(150) DEFAULT "" COMMENT "销方银行名称",
			  `seller_bank_account` varchar(50) DEFAULT "" COMMENT "销方银行账号",
			  `buyer_taxno` varchar(50) DEFAULT "" COMMENT "购方税号",
			  `buyer_name` varchar(150) DEFAULT "" COMMENT "购方名称",
			  `buyer_address` varchar(350) DEFAULT "" COMMENT "购方地址",
			  `buyer_tel` varchar(50) DEFAULT "" COMMENT "购方电话",
			  `buyer_bank_name` varchar(150) DEFAULT "" COMMENT "购方银行名称",
			  `buyer_bank_account` varchar(50) DEFAULT "" COMMENT "购方银行账号",
			  `amount_with_tax` decimal(18,2) DEFAULT "0.00" COMMENT "含税金额",
			  `amount_without_tax` decimal(18,2) DEFAULT "0.00" COMMENT "不含税金额",
			  `capital_amount_with_tax` varchar(100) DEFAULT "0.00" COMMENT "含税金额大写",
			  `tax_amount` decimal(18,2) DEFAULT "0.00" COMMENT "税额",
			  `deductions` decimal(18,2) DEFAULT "0.00" COMMENT "可抵扣金额",
			  `remark` varchar(4000) DEFAULT "" COMMENT "备注",
			  `is_valid` tinyint(1) DEFAULT 0 COMMENT "作废标记",
			  `payee` varchar(20) DEFAULT "" COMMENT "收款人",
			  `checker` varchar(20) DEFAULT "" COMMENT "复核人",
			  `drawer` varchar(20) DEFAULT "" COMMENT "开票人",
			  `invoice_date` int(11) DEFAULT "0" COMMENT "开票日期",
			  `check_code` varchar(50) DEFAULT "" COMMENT "校验码",
			  `machine_code` varchar(50) DEFAULT "" COMMENT "机器码",
			  `verify_times` smallint(5) DEFAULT 1 COMMENT "发票查验次数",
			  `flag` int(1) DEFAULT "1" COMMENT "发票标识 1 公司，2个人",
			  `receiver` varchar(50) DEFAULT "" COMMENT "电票接收人",
			  `receiver_phone` varchar(15) DEFAULT "" COMMENT "电票接收人电话",
			  `receiver_email` varchar(50) DEFAULT "" COMMENT "电票接收人邮箱",
			  `created_at` int(11) DEFAULT 0 COMMENT "创建时间",
			  `updated_at` int(11) DEFAULT 0 COMMENT "最后更新时间",
			  PRIMARY KEY (`id`),
			  INDEX `idx_code_no_index` (`invoice_code`,`invoice_no`,`invoice_date`,`amount_with_tax`) USING BTREE
		) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;');
        PREPARE stmt FROM @create_sql;
        EXECUTE stmt;
    END IF;
    CALL get_table_by_name(CONCAT('y_invoice_detail_', monthStr), table_exists);
    IF !table_exists THEN
        SET @create_detail_sql = CONCAT('CREATE TABLE y_invoice_detail_', monthStr, '(
			`id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT "主键",
			`invoice_id` int(11) DEFAULT "0" COMMENT "发票id",
			`item_name` varchar(150) DEFAULT "" COMMENT "货物及服务名称（电子发票通行费：项目名称）",
			`specification` varchar(50) DEFAULT "" COMMENT "型号规格（电子发票通行费：车牌号）",
			`quantity_unit` varchar(20) DEFAULT "" COMMENT "单位",
			`quantity` decimal(18,8) DEFAULT "0.00000000" COMMENT "数量",
			`price_type` tinyint(4) DEFAULT "0" COMMENT "价格方式（0=>不含税价、1=>含税价）",
			`tax_rate` decimal(18,2) DEFAULT "0.00" COMMENT "税率",
			`price` decimal(18,8) DEFAULT "0.00000000" COMMENT "单价",
			`amount_with_tax` decimal(18,2) DEFAULT "0.00" COMMENT "含税金额",
			`amount_without_tax` decimal(18,2) DEFAULT "0.00" COMMENT "不含税金额",
			`tax_amount` decimal(18,2) DEFAULT "0.00" COMMENT "税额",
			`deductions` decimal(18,2) DEFAULT "0.00" COMMENT "可抵扣金额",
			`goods_version` varchar(20) DEFAULT "" COMMENT "编码版本号",
			`goods_tax_no` varchar(20) DEFAULT "" COMMENT "税收分类编码",
			`goods_tax_name` varchar(100) DEFAULT "" COMMENT "商品税务名称",
			`tax_pre_con` varchar(10) DEFAULT "" COMMENT "零税率标志 免征  不征税 零税率",
			PRIMARY KEY (`id`)
			) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;');
        PREPARE stmt FROM @create_detail_sql;
        EXECUTE stmt;
    END IF;
END $$
DELIMITER ;

/**---------------------------------查验表结束---------------------------------**/