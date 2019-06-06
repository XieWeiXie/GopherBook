CREATE TABLE `wechat_persons` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `nick_name` varchar(10) DEFAULT NULL,
  `account_string` varchar(15) DEFAULT NULL,
  `account_qr` varchar(255) DEFAULT NULL,
  `gender` int(11) DEFAULT NULL,
  `location` varchar(255) DEFAULT NULL,
  `signal_person` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_wechat_persons_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `wechat_address` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `person_id` int(10) unsigned DEFAULT NULL,
  `name` varchar(10) DEFAULT NULL,
  `phone` varchar(11) DEFAULT NULL,
  `location_info` varchar(64) NOT NULL,
  `detail` varchar(128) NOT NULL,
  `code` varchar(6) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_wechat_address_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `wechat_receipt` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `person_id` int(10) unsigned DEFAULT NULL,
  `type` int(11) DEFAULT NULL,
  `name` varchar(32) DEFAULT NULL,
  `tax_number` varchar(32) DEFAULT NULL,
  `company_address` varchar(64) DEFAULT NULL,
  `phone` varchar(11) DEFAULT NULL,
  `bank` varchar(32) DEFAULT NULL,
  `bank_count` varchar(18) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_wechat_receipt_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



