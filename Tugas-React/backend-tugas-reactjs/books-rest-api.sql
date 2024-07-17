/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 80030 (8.0.30)
 Source Host           : localhost:3306
 Source Schema         : books-rest-api

 Target Server Type    : MySQL
 Target Server Version : 80030 (8.0.30)
 File Encoding         : 65001

 Date: 17/07/2024 00:16:35
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for age_rating_categories
-- ----------------------------
DROP TABLE IF EXISTS `age_rating_categories`;
CREATE TABLE `age_rating_categories`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `description` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_users_age_rating_categories`(`user_id` ASC) USING BTREE,
  CONSTRAINT `fk_users_age_rating_categories` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of age_rating_categories
-- ----------------------------
INSERT INTO `age_rating_categories` VALUES (1, 'PG', 'Parental Guidance Suggested', '2024-07-15 23:45:17.000', '2024-07-15 23:45:17.000', 1);
INSERT INTO `age_rating_categories` VALUES (2, 'R', 'Restricted', '2024-07-15 23:45:17.000', '2024-07-15 23:45:17.000', 2);

-- ----------------------------
-- Table structure for books
-- ----------------------------
DROP TABLE IF EXISTS `books`;
CREATE TABLE `books`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `description` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `image_url` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `release_year` bigint NULL DEFAULT NULL,
  `price` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `total_page` bigint NULL DEFAULT NULL,
  `thickness` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of books
-- ----------------------------
INSERT INTO `books` VALUES (1, 'Sample Book', 'This is a sample book.', 'http://example.com/image.png', 1990, '9.99', 150, '', '2024-07-16 22:54:25.906', '2024-07-16 22:54:25.906');
INSERT INTO `books` VALUES (2, 'Sample Book', 'This is a sample book.', 'http://example.com/image.png', 1990, '9.99', 150, '', '2024-07-17 00:01:31.529', '2024-07-17 00:01:31.529');
INSERT INTO `books` VALUES (3, 'Sample Book', 'This is a sample book.', 'http://example.com/image.png', 1990, '9.99', 150, '', '2024-07-17 00:01:39.348', '2024-07-17 00:01:39.348');
INSERT INTO `books` VALUES (4, 'New Book Title', 'A detailed description of the new book.', 'http://example.com/image.jpg', 2021, '19.99', 350, '', '2024-07-17 00:08:50.134', '2024-07-17 00:08:50.134');
INSERT INTO `books` VALUES (5, 'Sample Book', 'This is a sample book.', 'http://example.com/image.png', 1990, '9.99', 150, '', '2024-07-17 00:10:22.754', '2024-07-17 00:10:22.754');

-- ----------------------------
-- Table structure for footballers
-- ----------------------------
DROP TABLE IF EXISTS `footballers`;
CREATE TABLE `footballers`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `nationality` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `age` bigint NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of footballers
-- ----------------------------
INSERT INTO `footballers` VALUES (1, 'adam', 'indo', 21, '2024-07-16 22:42:31.000', '2024-07-16 22:42:36.000');
INSERT INTO `footballers` VALUES (2, 'rizky', 'malay', 20, '2024-07-16 22:42:38.000', '2024-07-16 22:42:44.000');
INSERT INTO `footballers` VALUES (3, 'faras', 'zimabgwqwe', 12, '2024-07-16 22:44:58.420', '2024-07-16 22:44:58.420');
INSERT INTO `footballers` VALUES (4, 'bian', 'indo', 10, '2024-07-16 22:47:40.263', '2024-07-16 22:47:40.263');
INSERT INTO `footballers` VALUES (5, '', '', 0, '2024-07-16 23:29:05.410', '2024-07-16 23:29:05.410');
INSERT INTO `footballers` VALUES (6, 'faras', 'zimabgwqwe', 12, '2024-07-16 23:36:32.413', '2024-07-16 23:36:32.413');
INSERT INTO `footballers` VALUES (7, 'John Doe', 'American', 30, '2024-07-17 00:00:23.481', '2024-07-17 00:00:23.481');
INSERT INTO `footballers` VALUES (8, 'John Doe', 'American', 30, '2024-07-17 00:00:49.073', '2024-07-17 00:00:49.073');
INSERT INTO `footballers` VALUES (9, 'John Doe', 'American', 30, '2024-07-17 00:03:02.937', '2024-07-17 00:03:02.937');

-- ----------------------------
-- Table structure for movies
-- ----------------------------
DROP TABLE IF EXISTS `movies`;
CREATE TABLE `movies`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `year` bigint NULL DEFAULT NULL,
  `age_rating_category_id` bigint UNSIGNED NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_age_rating_categories_movies`(`age_rating_category_id` ASC) USING BTREE,
  CONSTRAINT `fk_age_rating_categories_movies` FOREIGN KEY (`age_rating_category_id`) REFERENCES `age_rating_categories` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of movies
-- ----------------------------
INSERT INTO `movies` VALUES (1, 'The Great Adventure', 2021, 1, '2024-07-15 23:45:17.000', '2024-07-15 23:45:17.000');
INSERT INTO `movies` VALUES (2, 'Night Walk', 2022, 2, '2024-07-15 23:45:17.000', '2024-07-15 23:45:17.000');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `email` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_users_username`(`username` ASC) USING BTREE,
  UNIQUE INDEX `uni_users_email`(`email` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, 'adam', 'admin@gmail.com', '1234qwer', '2024-07-15 23:45:17.000', '2024-07-15 23:45:17.000');
INSERT INTO `users` VALUES (2, 'jane_smith', 'jane.smith@example.com', '1234qwer', '2024-07-15 23:45:17.000', '2024-07-15 23:45:17.000');
INSERT INTO `users` VALUES (4, 'adam2', 'adam@gmail.com', '$2a$10$8EdVii1o7.3Li68QhAjJ3OQxEfXAl9Nxx/y67YNngw2VdSiJcU8gG', '2024-07-16 00:14:13.463', '2024-07-16 00:14:13.463');

SET FOREIGN_KEY_CHECKS = 1;
