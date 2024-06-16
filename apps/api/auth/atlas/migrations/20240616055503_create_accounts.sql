-- Create "accounts" table
CREATE TABLE `accounts` (`id` bigint NOT NULL AUTO_INCREMENT, `email` varchar(255) NOT NULL, `password` varchar(255) NOT NULL, `created_at` timestamp NOT NULL, `update_at` timestamp NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `email` (`email`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
