-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        5.7.26 - MySQL Community Server (GPL)
-- 服务器操作系统:                      Win64
-- HeidiSQL 版本:                  12.5.0.6677
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- 导出 personalbill 的数据库结构
CREATE DATABASE IF NOT EXISTS `personalbill` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;
USE `personalbill`;

-- 导出  表 personalbill.bill 结构
CREATE TABLE IF NOT EXISTS `bill` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `u_id` int(11) DEFAULT NULL COMMENT '用户id',
  `kind` tinyint(1) DEFAULT NULL COMMENT '1:收入,2:支出',
  `status` tinyint(1) DEFAULT '1',
  `money` float(10,2) DEFAULT NULL COMMENT '金额，单位：元',
  `category_id` int(11) DEFAULT NULL COMMENT '类别id',
  `date` date DEFAULT NULL,
  `note` varchar(255) DEFAULT NULL COMMENT '备注',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- 正在导出表  personalbill.bill 的数据：1 rows
/*!40000 ALTER TABLE `bill` DISABLE KEYS */;
INSERT INTO `bill` (`id`, `u_id`, `kind`, `status`, `money`, `category_id`, `date`, `note`, `created_at`, `updated_at`) VALUES
	(1, 1, 1, 1, 10.20, 1, '2023-10-16', NULL, '2023-10-16 22:36:32', '2023-10-16 22:36:35');
/*!40000 ALTER TABLE `bill` ENABLE KEYS */;

-- 导出  表 personalbill.category 结构
CREATE TABLE IF NOT EXISTS `category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(20) DEFAULT NULL,
  `pid` int(11) DEFAULT NULL,
  `kind` tinyint(1) NOT NULL COMMENT '1:收入,2:支出',
  `icon` varchar(255) DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- 正在导出表  personalbill.category 的数据：1 rows
/*!40000 ALTER TABLE `category` DISABLE KEYS */;
INSERT INTO `category` (`id`, `title`, `pid`, `kind`, `icon`, `status`, `created_at`, `updated_at`) VALUES
	(1, '餐饮', 0, 1, NULL, 1, '2023-10-16 22:28:54', '2023-10-16 22:28:56');
/*!40000 ALTER TABLE `category` ENABLE KEYS */;

-- 导出  表 personalbill.user 结构
CREATE TABLE IF NOT EXISTS `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `phone_num` varchar(11) NOT NULL,
  `password` varchar(255) DEFAULT NULL,
  `nickname` varchar(32) DEFAULT NULL,
  `gender` int(1) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `birthday` date DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`,`phone_num`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- 正在导出表  personalbill.user 的数据：1 rows
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` (`id`, `phone_num`, `password`, `nickname`, `gender`, `avatar`, `birthday`, `created_at`, `updated_at`) VALUES
	(1, '18883362533', '$2a$10$Bou3PxCLW1XflWrDHqfplONpTvyEHxDghuamCrDLo741Zz7EaWAGu', '哈哈哈', 1, '111', '2023-10-16', '2023-10-16 21:19:40', '2023-10-16 21:19:43');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
