/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 100432 (10.4.32-MariaDB)
 Source Host           : localhost:3306
 Source Schema         : db_university

 Target Server Type    : MySQL
 Target Server Version : 100432 (10.4.32-MariaDB)
 File Encoding         : 65001

 Date: 29/06/2024 03:56:54
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for mahasiswa
-- ----------------------------
DROP TABLE IF EXISTS `mahasiswa`;
CREATE TABLE `mahasiswa`  (
  `id` int NOT NULL,
  `nama` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of mahasiswa
-- ----------------------------
INSERT INTO `mahasiswa` VALUES (1, 'Budi', '2024-06-29 02:58:40', '2024-06-29 02:58:40');
INSERT INTO `mahasiswa` VALUES (2, 'Siti', '2024-06-29 02:58:40', '2024-06-29 02:58:40');
INSERT INTO `mahasiswa` VALUES (3, 'Agus', '2024-06-29 02:58:40', '2024-06-29 02:58:40');
INSERT INTO `mahasiswa` VALUES (4, 'Rina', '2024-06-29 02:58:40', '2024-06-29 02:58:40');

-- ----------------------------
-- Table structure for mata_kuliah
-- ----------------------------
DROP TABLE IF EXISTS `mata_kuliah`;
CREATE TABLE `mata_kuliah`  (
  `id` int NOT NULL,
  `nama` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of mata_kuliah
-- ----------------------------
INSERT INTO `mata_kuliah` VALUES (1, 'Matematika Lanjut', '2024-06-29 02:58:40', '2024-06-29 02:58:40');
INSERT INTO `mata_kuliah` VALUES (2, 'Fisika Dasar', '2024-06-29 02:58:40', '2024-06-29 02:58:40');
INSERT INTO `mata_kuliah` VALUES (3, 'Kimia Organik', '2024-06-29 02:58:40', '2024-06-29 02:58:40');
INSERT INTO `mata_kuliah` VALUES (4, 'Ilmu Komputer', '2024-06-29 02:58:40', '2024-06-29 02:58:40');

-- ----------------------------
-- Table structure for nilai
-- ----------------------------
DROP TABLE IF EXISTS `nilai`;
CREATE TABLE `nilai`  (
  `id` int NOT NULL,
  `indeks` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `skor` int NULL DEFAULT NULL,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `mata_kuliah_id` int NULL DEFAULT NULL,
  `mahasiswa_id` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `mata_kuliah_id`(`mata_kuliah_id` ASC) USING BTREE,
  INDEX `mahasiswa_id`(`mahasiswa_id` ASC) USING BTREE,
  CONSTRAINT `nilai_ibfk_1` FOREIGN KEY (`mata_kuliah_id`) REFERENCES `mata_kuliah` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `nilai_ibfk_2` FOREIGN KEY (`mahasiswa_id`) REFERENCES `mahasiswa` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of nilai
-- ----------------------------
INSERT INTO `nilai` VALUES (1, 'A', 85, '2024-06-29 02:58:40', '2024-06-29 02:58:40', 1, 1);
INSERT INTO `nilai` VALUES (2, 'B', 75, '2024-06-29 02:58:40', '2024-06-29 02:58:40', 2, 1);
INSERT INTO `nilai` VALUES (3, 'C', 65, '2024-06-29 02:58:40', '2024-06-29 02:58:40', 1, 2);
INSERT INTO `nilai` VALUES (4, 'D', 55, '2024-06-29 02:58:40', '2024-06-29 02:58:40', 3, 3);
INSERT INTO `nilai` VALUES (5, 'E', 45, '2024-06-29 02:58:40', '2024-06-29 02:58:40', 4, 4);
INSERT INTO `nilai` VALUES (6, 'A', 90, '2024-06-29 02:58:40', '2024-06-29 02:58:40', 4, 2);

SET FOREIGN_KEY_CHECKS = 1;
