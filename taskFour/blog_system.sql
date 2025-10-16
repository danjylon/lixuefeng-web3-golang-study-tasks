/*M!999999\- enable the sandbox mode */ 
-- MariaDB dump 10.19-11.7.2-MariaDB, for Win64 (AMD64)
--
-- Host: localhost    Database: blog_system
-- ------------------------------------------------------
-- Server version	10.6.7-MariaDB-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*M!100616 SET @OLD_NOTE_VERBOSITY=@@NOTE_VERBOSITY, NOTE_VERBOSITY=0 */;

--
-- Table structure for table `comments`
--

DROP TABLE IF EXISTS `comments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `comments` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `content` varchar(255) NOT NULL,
  `post_id` bigint(20) unsigned DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_comments_deleted_at` (`deleted_at`),
  KEY `fk_users_comments` (`user_id`),
  KEY `fk_posts_comments` (`post_id`),
  CONSTRAINT `fk_posts_comments` FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`),
  CONSTRAINT `fk_users_comments` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comments`
--

LOCK TABLES `comments` WRITE;
/*!40000 ALTER TABLE `comments` DISABLE KEYS */;
INSERT INTO `comments` VALUES
(1,'2025-10-16 16:06:59.391','2025-10-16 16:06:59.391',NULL,'感谢楼主分享',3,1),
(2,'2025-10-16 16:15:06.751','2025-10-16 16:15:06.751',NULL,'写的不错',3,1),
(3,'2025-10-16 16:15:15.120','2025-10-16 16:15:15.120',NULL,'对我帮助很大',3,1);
/*!40000 ALTER TABLE `comments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `posts`
--

DROP TABLE IF EXISTS `posts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `posts` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` varchar(50) NOT NULL,
  `content` text NOT NULL,
  `comment_status` tinyint(1) DEFAULT 0,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_posts_deleted_at` (`deleted_at`),
  KEY `fk_users_posts` (`user_id`),
  CONSTRAINT `fk_users_posts` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `posts`
--

LOCK TABLES `posts` WRITE;
/*!40000 ALTER TABLE `posts` DISABLE KEYS */;
INSERT INTO `posts` VALUES
(1,'2025-10-16 14:15:51.736','2025-10-16 15:41:02.354','2025-10-16 15:51:38.126','模型定义2','假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。\n  - 要求 ：\n  - 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。\n  - 编写Go代码，使用Gorm创建这些模型对应的数据库表',0,1),
(2,'2025-10-16 14:19:16.869','2025-10-16 14:19:16.869',NULL,'关联查询','基于上述博客系统的模型定义。\n  - 要求 ：\n  - 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。\n  - 编写Go代码，使用Gorm查询评论数量最多的文章信息。',0,1),
(3,'2025-10-16 14:20:03.385','2025-10-16 16:06:59.392',NULL,'钩子函数','继续使用博客系统的模型。\n  - 要求 ：\n  - 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。\n  - 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 \"无评论\"。',1,2),
(4,'2025-10-16 14:20:37.362','2025-10-16 14:20:37.362',NULL,'使用SQL扩展库进行查询','假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。\n要求 ：\n编写Go代码，使用Sqlx查询 employees 表中所有部门为 \"技术部\" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。\n编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。',0,2),
(5,'2025-10-16 14:21:00.784','2025-10-16 14:21:00.784',NULL,'实现类型安全映射','假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。\n要求 ：\n定义一个 Book 结构体，包含与 books 表对应的字段。\n编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。',0,2),
(6,'2025-10-16 14:21:56.101','2025-10-16 14:21:56.101',NULL,'删除有序数组中的重复项','删除有序数组中的重复项给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。\n可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。',0,3),
(7,'2025-10-16 14:22:16.945','2025-10-16 14:22:16.945',NULL,'最长公共前缀','考察：字符串处理、循环嵌套\n题目：查找字符串数组中的最长公共前缀',0,3);
/*!40000 ALTER TABLE `posts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `username` varchar(30) NOT NULL,
  `pwd` varchar(100) NOT NULL,
  `phone` varchar(20) NOT NULL,
  `email` varchar(30) DEFAULT NULL,
  `post_num` bigint(20) DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_users_username` (`username`),
  UNIQUE KEY `uni_users_phone` (`phone`),
  UNIQUE KEY `uni_users_email` (`email`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES
(1,'2025-10-16 09:20:34.371','2025-10-16 15:51:38.127',NULL,'tom','$2a$10$WFfuiZGiAd1V7u89GHVbjOO30g2kIR6NTzBRJfwbqfnj/LDIu0GbK','+8613111111111',NULL,1),
(2,'2025-10-16 09:21:20.068','2025-10-16 14:21:00.784',NULL,'jerry','$2a$10$WXNALb.7D1Xw2dpl0l/N.eFWznM0DsvfNsJ62XXu/AH9Z5P7Qi.PG','+8613222222222','jerry@golang.com',3),
(3,'2025-10-16 09:56:47.489','2025-10-16 14:22:16.946',NULL,'jack','$2a$10$UgyDbEvaJwSpxrgkM2vMvuTiKSpqd0Tgc1bJ/JTodOeYTou3GlCH.','+8613333333333',NULL,2);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'blog_system'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*M!100616 SET NOTE_VERBOSITY=@OLD_NOTE_VERBOSITY */;

-- Dump completed on 2025-10-16 16:34:15
