-- LiteFlow 数据库初始化脚本 (UUID版本)
-- 使用方法: mysql -u root -p < init_database_uuid.sql

-- 创建数据库（如果不存在）
CREATE DATABASE IF NOT EXISTS `liteflow` 
DEFAULT CHARACTER SET utf8mb4 
COLLATE utf8mb4_unicode_ci;

-- 使用数据库
USE `liteflow`;

-- 创建链路配置表
DROP TABLE IF EXISTS `liteflow_chain`;
CREATE TABLE `liteflow_chain` (
  `id` varchar(36) NOT NULL COMMENT '主键ID(UUID)',
  `chain_id` varchar(64) NOT NULL COMMENT '链路ID',
  `chain_name` varchar(128) NOT NULL COMMENT '链路名称',
  `chain_desc` text COMMENT '链路描述',
  `el_data` text NOT NULL COMMENT 'EL表达式数据',
  `enable` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否启用 1:启用 0:禁用',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_chain_id` (`chain_id`),
  KEY `idx_chain_name` (`chain_name`),
  KEY `idx_enable` (`enable`),
  KEY `idx_create_time` (`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='LiteFlow链路配置表';

-- 创建脚本配置表
DROP TABLE IF EXISTS `liteflow_script`;
CREATE TABLE `liteflow_script` (
  `id` varchar(36) NOT NULL COMMENT '主键ID(UUID)',
  `script_id` varchar(64) NOT NULL COMMENT '脚本ID',
  `script_name` varchar(128) NOT NULL COMMENT '脚本名称',
  `script_type` varchar(32) NOT NULL COMMENT '脚本类型(groovy,js,python等)',
  `script_data` longtext NOT NULL COMMENT '脚本内容',
  `script_desc` text COMMENT '脚本描述',
  `enable` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否启用 1:启用 0:禁用',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_script_id` (`script_id`),
  KEY `idx_script_name` (`script_name`),
  KEY `idx_script_type` (`script_type`),
  KEY `idx_enable` (`enable`),
  KEY `idx_create_time` (`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='LiteFlow脚本配置表';

-- 创建节点配置表
DROP TABLE IF EXISTS `liteflow_node`;
CREATE TABLE `liteflow_node` (
  `id` varchar(36) NOT NULL COMMENT '主键ID(UUID)',
  `node_id` varchar(64) NOT NULL COMMENT '节点ID',
  `node_name` varchar(128) NOT NULL COMMENT '节点名称',
  `node_type` varchar(32) NOT NULL DEFAULT 'common' COMMENT '节点类型(common,switch,for,while等)',
  `class_name` varchar(255) COMMENT '节点实现类名',
  `script_id` varchar(64) COMMENT '关联脚本ID',
  `node_desc` text COMMENT '节点描述',
  `enable` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否启用 1:启用 0:禁用',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_node_id` (`node_id`),
  KEY `idx_node_name` (`node_name`),
  KEY `idx_node_type` (`node_type`),
  KEY `idx_script_id` (`script_id`),
  KEY `idx_enable` (`enable`),
  KEY `idx_create_time` (`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='LiteFlow节点配置表';

-- 创建流程执行日志表
DROP TABLE IF EXISTS `liteflow_log`;
CREATE TABLE `liteflow_log` (
  `id` varchar(36) NOT NULL COMMENT '主键ID(UUID)',
  `request_id` varchar(64) NOT NULL COMMENT '请求ID',
  `chain_id` varchar(64) NOT NULL COMMENT '链路ID',
  `chain_name` varchar(128) COMMENT '链路名称',
  `node_id` varchar(64) COMMENT '节点ID',
  `node_name` varchar(128) COMMENT '节点名称',
  `execute_status` varchar(32) NOT NULL COMMENT '执行状态(SUCCESS,FAILED,RUNNING)',
  `execute_time` bigint(20) COMMENT '执行耗时(毫秒)',
  `error_msg` text COMMENT '错误信息',
  `input_data` longtext COMMENT '输入数据',
  `output_data` longtext COMMENT '输出数据',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_request_id` (`request_id`),
  KEY `idx_chain_id` (`chain_id`),
  KEY `idx_node_id` (`node_id`),
  KEY `idx_execute_status` (`execute_status`),
  KEY `idx_create_time` (`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='LiteFlow执行日志表';

-- 插入示例数据（使用UUID函数生成ID）
INSERT INTO `liteflow_chain` (`id`, `chain_id`, `chain_name`, `chain_desc`, `el_data`) VALUES
(UUID(), 'test_chain', '测试链路', '这是一个测试用的简单链路', 'THEN(a, b, c)'),
(UUID(), 'complex_chain', '复杂业务链路', '包含条件判断和循环的复杂链路', 'THEN(a, IF(x, THEN(b, c), d), e)');

INSERT INTO `liteflow_node` (`id`, `node_id`, `node_name`, `node_type`, `node_desc`) VALUES
(UUID(), 'a', '节点A', 'common', '普通业务节点A'),
(UUID(), 'b', '节点B', 'common', '普通业务节点B'),
(UUID(), 'c', '节点C', 'common', '普通业务节点C'),
(UUID(), 'd', '节点D', 'common', '普通业务节点D'),
(UUID(), 'e', '节点E', 'common', '普通业务节点E'),
(UUID(), 'x', '条件节点X', 'switch', '条件判断节点X');

INSERT INTO `liteflow_script` (`id`, `script_id`, `script_name`, `script_type`, `script_data`, `script_desc`) VALUES
(UUID(), 'demo_script', '演示脚本', 'groovy', 'def result = "Hello from Groovy Script"\nreturn result', '这是一个演示用的Groovy脚本');

-- 显示创建结果
SELECT '数据库初始化完成！' AS message;
SELECT COUNT(*) AS chain_count FROM liteflow_chain;
SELECT COUNT(*) AS node_count FROM liteflow_node;
SELECT COUNT(*) AS script_count FROM liteflow_script;