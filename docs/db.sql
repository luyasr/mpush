CREATE TABLE `user` (
                        `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键',
                        `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户名',
                        `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '密码',
                        `nickname` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '昵称',
                        `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '邮箱',
                        `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '电话',
                        `status` tinyint NOT NULL COMMENT '状态',
                        `role` tinyint NOT NULL COMMENT '角色',
                        `created_at` bigint NOT NULL COMMENT '创建时间',
                        `updated_at` bigint NOT NULL COMMENT '更新时间',
                        `deleted_at` bigint unsigned DEFAULT NULL COMMENT '删除时间',
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `idx_username_deleted_at` (`username`,`deleted_at`),
                        UNIQUE KEY `idx_nickname_deleted_at` (`nickname`,`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `token` (
                         `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键',
                         `user_id` bigint NOT NULL COMMENT '用户id',
                         `access_token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '登录token',
                         `refresh_token` varchar(255) NOT NULL COMMENT '刷新token',
                         `access_token_expired_at` bigint NOT NULL COMMENT '到期时间戳',
                         `refresh_token_expired_at` bigint NOT NULL COMMENT '刷新时间戳',
                         `created_at` bigint NOT NULL COMMENT '创建时间',
                         `updated_at` bigint NOT NULL COMMENT '更新时间',
                         `deleted_at` bigint DEFAULT NULL COMMENT '删除时间',
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `idx_token_deleted_at` (`access_token`,`deleted_at`) USING BTREE COMMENT '联合索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;