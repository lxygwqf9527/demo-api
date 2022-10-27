CREATE TABLE `resource` (
  `id` char(64) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL COMMENT '资源的实例ID',
  `vendor` tinyint(1) NOT NULL,
  `region` varchar(64) CHARACTER SET latin1  COLLATE latin1_swedish_ci NOT NULL,
  `create_at` bigint(13) NOT NULL,
  `expire_at` bigint(13) DEFAULT NULL,
  `type` varchar(120) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `status` varchar(255) CHARACTER SET latin1  COLLATE latin1_swedish_ci NOT NULL,
  `update_at` bigint(13) DEFAULT NULL,
  `sync_at` bigint(13) DEFAULT NULL,
  `account` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `public_ip` varchar(64) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `private_ip` varchar(64) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  PRIMARY KEY (`id`),
  KEY `name` (`name`) USING BTREE,
  KEY `status` (`status`) USING HASH,
  KEY `private_ip` (`public_ip`) USING BTREE,
  KEY `public_ip` (`public_ip`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;



CREATE TABLE `host` (
  `resource_id` varchar(64) NOT NULL,
  `cpu` tinyint NOT NULL,
  `memory` int NOT NULL,
  `gpu_amount` tinyint DEFAULT NULL,
  `gpu_spec` varchar(255) DEFAULT NULL,
  `os_type` varchar(255) DEFAULT NULL,
  `os_name` varchar(255) DEFAULT NULL,
  `serial_number` varchar(120) DEFAULT NULL,
  PRIMARY KEY (`resource_id`) 
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;