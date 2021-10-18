SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for phone_account
-- ----------------------------
DROP TABLE IF EXISTS `mini_program_account`;
CREATE TABLE `mini_program_account` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_identity` char(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户标识',
  `open_id` char(28) COLLATE utf8mb4_general_ci NOT NULL COMMENT 'open_id',
  `nick_name` varchar(128) COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户昵称',
  `avatar_url` varchar(1024) COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户头像',
  `gender` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '用户性别：0-未知，1-男性，2-女性，默认-0',
  `country` varchar(128) COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户所在国家',
  `province` varchar(128) COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户所在省份',
  `city` varchar(128) COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户所在城市',
  `language` varchar(28) COLLATE utf8mb4_general_ci NOT NULL COMMENT '所用语言：en-英语，zh_CN-简体中文，zh_TW-繁体中文',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uidx_user_identity` (`user_identity`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='微信小程序登录表';

SET FOREIGN_KEY_CHECKS = 1;