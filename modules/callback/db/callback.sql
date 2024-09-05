CREATE TABLE `zero_callback_records` (
    `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键ID',
    `callback_func_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '回调方法名',
    `params` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '参数，格式 json',
    `retry_cnt` int unsigned NOT NULL DEFAULT '1' COMMENT '重试次数',
    `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态 -1 回调异常 0 待回调 1 已回调完成',
    `remark` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '回调异常信息',
    `expected_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '期望执行时间',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='回调记录';