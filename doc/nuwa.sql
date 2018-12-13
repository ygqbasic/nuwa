-- MySQL dump 10.13  Distrib 5.6.40, for Win64 (x86_64)
--
-- Host: localhost    Database: nuwa
-- ------------------------------------------------------
-- Server version	5.6.40

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `sys_backend_conf`
--

DROP TABLE IF EXISTS `sys_backend_conf`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_backend_conf` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '序号',
  `appname` varchar(50) NOT NULL,
  `appversion` varchar(50) NOT NULL,
  `deploy` varchar(50) NOT NULL,
  `configtext` varchar(5000) DEFAULT NULL,
  `tag` int(11) DEFAULT '0' COMMENT '状态',
  `createuser` varchar(20) DEFAULT NULL COMMENT '创建人',
  `createdate` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `changeuser` varchar(20) DEFAULT NULL COMMENT '修改人',
  `changedate` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='后台APP配置';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_backend_conf`
--

LOCK TABLES `sys_backend_conf` WRITE;
/*!40000 ALTER TABLE `sys_backend_conf` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_backend_conf` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_backend_user`
--

DROP TABLE IF EXISTS `sys_backend_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_backend_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `real_name` varchar(255) NOT NULL DEFAULT '',
  `user_name` varchar(255) NOT NULL DEFAULT '',
  `user_pwd` varchar(255) NOT NULL DEFAULT '',
  `is_super` tinyint(1) NOT NULL DEFAULT '0',
  `status` int(11) NOT NULL DEFAULT '0',
  `mobile` varchar(16) NOT NULL DEFAULT '',
  `email` varchar(256) NOT NULL DEFAULT '',
  `avatar` varchar(256) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='后台用户管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_backend_user`
--

LOCK TABLES `sys_backend_user` WRITE;
/*!40000 ALTER TABLE `sys_backend_user` DISABLE KEYS */;
INSERT INTO `sys_backend_user` VALUES (1,'超级管理员','admin','21232f297a57a5a743894a0e4a801fc3',1,1,'13812345678','ygqbasic@gmail.com','/static/upload/a290X290.jpg'),
(5,'guest','guest','e10adc3949ba59abbe56e057f20f883e',0,1,'13812345678','13812345678@qq.com','/static/upload/psb.jpg');
/*!40000 ALTER TABLE `sys_backend_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_backend_user_rms_roles`
--

DROP TABLE IF EXISTS `sys_backend_user_rms_roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_backend_user_rms_roles` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `rms_backend_user_id` int(11) NOT NULL,
  `rms_role_id` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_backend_user_rms_roles`
--

LOCK TABLES `sys_backend_user_rms_roles` WRITE;
/*!40000 ALTER TABLE `sys_backend_user_rms_roles` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_backend_user_rms_roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_logintrace`
--

DROP TABLE IF EXISTS `sys_logintrace`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_logintrace` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '序号',
  `user` varchar(30) DEFAULT NULL COMMENT '用户名',
  `remoteAddr` varchar(50) DEFAULT NULL COMMENT 'IP地址',
  `loginTime` datetime DEFAULT NULL COMMENT '登录时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_logintrace`
--

LOCK TABLES `sys_logintrace` WRITE;
/*!40000 ALTER TABLE `sys_logintrace` DISABLE KEYS */;
INSERT INTO `sys_logintrace` VALUES (1,'admin','localhost','2018-12-10 16:01:30'),
(2,'admin','localhost','2018-12-10 16:01:30'),
(3,'admin','localhost','2018-12-10 16:01:31'),
(4,'admin','localhost','2018-12-10 16:01:32'),
(5,'admin','localhost','2018-12-10 16:08:21'),
(6,'admin','localhost','2018-12-10 16:12:03'),
(7,'admin','localhost','2018-12-10 16:59:27'),
(8,'admin','localhost','2018-12-11 15:44:25'),
(9,'admin','localhost','2018-12-11 16:21:26');
/*!40000 ALTER TABLE `sys_logintrace` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_resource`
--

DROP TABLE IF EXISTS `sys_resource`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_resource` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `rtype` int(11) NOT NULL DEFAULT '0',
  `name` varchar(64) NOT NULL DEFAULT '',
  `parent_id` int(11) DEFAULT NULL,
  `seq` int(11) NOT NULL DEFAULT '0',
  `icon` varchar(32) NOT NULL DEFAULT '',
  `url_for` varchar(256) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=88 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='资源管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_resource`
--

LOCK TABLES `sys_resource` WRITE;
/*!40000 ALTER TABLE `sys_resource` DISABLE KEYS */;
INSERT INTO `sys_resource` VALUES (7,1,'权限管理',8,100,'fa fa-balance-scale',''),
(8,0,'系统菜单',NULL,900,'',''),
(9,1,'资源管理',7,100,'fa fa-television','ResourceController.Index'),
(12,1,'角色管理',7,200,'fa fa-users','RoleController.Index'),
(13,1,'用户管理',7,300,'fa fa-user','BackendUserController.Index'),
(14,1,'系统管理',8,120,'fa fa-gears',''),
(21,0,'业务菜单',NULL,200,'',''),
(23,1,'日志管理',14,109,'fa fa-building-o','LoginTraceController.Index'),
(25,2,'编辑',9,100,'fa fa-pencil','ResourceController.Edit'),
(26,2,'编辑',13,100,'fa fa-pencil','BackendUserController.Edit'),
(27,2,'删除',9,100,'fa fa-trash','ResourceController.Delete'),
(29,2,'删除',13,100,'fa fa-trash','BackendUserController.Delete'),
(30,2,'编辑',12,100,'fa fa-pencil','RoleController.Edit'),
(31,2,'删除',12,100,'fa fa-trash','RoleController.Delete'),
(32,2,'分配资源',12,100,'fa fa-th','RoleController.Allocate'),
(35,1,'仪表1',75,1,'fa fa-wpforms','HomeController.Index'),
(41,1,'项目管理',21,210,'fa fa-credit-card',''),
(44,1,'项目管理',41,211,'fa fa-inbox',''),
(45,1,'应用管理',41,212,'fa fa-feed',''),
(48,2,'编辑',44,100,'fa fa-pencil','EquipmentVendorController.Edit'),
(49,2,'删除',44,101,'fa fa-trash','EquipmentVendorController.Delete'),
(50,2,'编辑',45,100,'fa fa-pencil','EquipmentGatewayController.Edit'),
(51,2,'删除',45,100,'fa fa-trash','EquipmentGatewayController.Delete'),
(66,1,'后台APP配置',14,108,'fa fa-th','BackendConfController.Index'),
(67,2,'编辑',66,100,'fa fa-pencil','BackendConfController.Edit'),
(68,2,'删除',66,100,'fa fa-trash','BackendConfController.Delete'),
(69,1,'数据管理',21,300,'fa fa-database',''),
(71,1,'部署日志查询',69,100,'fa fa-suitcase','TotalActivePowerMinuteController.Index'),
(72,1,'系统工具',8,100,'fa fa-magnet',''),
(73,1,'图标信息',72,100,'fa fa-gg','IconsController.Index'),
(74,1,'Websocket测试',72,100,'fa fa-skyatlas','WebsocketWidgetController.Index'),
(75,1,'仪表盘',NULL,100,'menu-icon fa fa-tachometer',''),
(76,1,'仪表2',75,2,'fa fa-tv','HomeController.Index2'),
(77,1,'Kubernetes',21,100,'fa fa-cloud',''),
(78,1,'群集管理',77,100,'fa fa-cubes',''),
(79,2,'编辑',78,100,'fa fa-pencil',''),
(80,2,'删除',78,101,'fa fa-trash','EquipmentCustomerController.Delete'),
(81,1,'主机信息',77,101,'fa fa-desktop',''),
(82,2,'编辑',81,100,'fa fa-pencil','EquipmentRoomController.Edit'),
(83,2,'删除',81,101,'fa fa-trash','EquipmentRoomController.Delete'),
(84,1,'系统参数管理',14,100,'fa fa-list-alt','SystemValController.Index'),
(85,2,'编辑',84,100,'fa fa-pencil','SystemValController.Edit'),
(86,2,'删除',84,101,'fa fa-trash','SystemValController.Delete'),
(87,1,'应用日志查询',69,100,'fa fa-building','TotalDtuRowsController.Index');
/*!40000 ALTER TABLE `sys_resource` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_role`
--

DROP TABLE IF EXISTS `sys_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `seq` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='角色管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role`
--

LOCK TABLES `sys_role` WRITE;
/*!40000 ALTER TABLE `sys_role` DISABLE KEYS */;
INSERT INTO `sys_role` VALUES (22,'超级管理员',1),
(24,'角色管理员',10),(25,'客户管理',5),(26,'运维人员',4);
/*!40000 ALTER TABLE `sys_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_role_backenduser_rel`
--

DROP TABLE IF EXISTS `sys_role_backenduser_rel`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_role_backenduser_rel` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL,
  `backend_user_id` int(11) NOT NULL,
  `created` datetime NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=80 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role_backenduser_rel`
--

LOCK TABLES `sys_role_backenduser_rel` WRITE;
/*!40000 ALTER TABLE `sys_role_backenduser_rel` DISABLE KEYS */;
INSERT INTO `sys_role_backenduser_rel` VALUES (69,22,1,'2018-02-23 17:06:50'),
(72,25,5,'2018-03-15 17:57:17'),
(74,25,6,'2018-04-04 13:44:18'),
(76,25,7,'2018-04-13 11:44:43'),
(77,25,8,'2018-05-11 11:22:03'),
(78,26,4,'2018-05-27 17:49:19'),
(79,25,4,'2018-05-27 17:49:19');
/*!40000 ALTER TABLE `sys_role_backenduser_rel` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_role_resource_rel`
--

DROP TABLE IF EXISTS `sys_role_resource_rel`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_role_resource_rel` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL,
  `resource_id` int(11) NOT NULL,
  `created` datetime NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1340 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role_resource_rel`
--

LOCK TABLES `sys_role_resource_rel` WRITE;
/*!40000 ALTER TABLE `sys_role_resource_rel` DISABLE KEYS */;
INSERT INTO `sys_role_resource_rel` VALUES (448,24,8,'2017-12-19 06:40:16'),
(449,24,14,'2017-12-19 06:40:16'),
(450,24,23,'2017-12-19 06:40:16'),
(1199,25,75,'2018-05-27 17:45:12'),
(1200,25,35,'2018-05-27 17:45:12'),
(1201,25,76,'2018-05-27 17:45:12'),
(1202,25,21,'2018-05-27 17:45:12'),
(1203,25,77,'2018-05-27 17:45:12'),
(1204,25,78,'2018-05-27 17:45:12'),
(1205,25,81,'2018-05-27 17:45:12'),
(1206,25,41,'2018-05-27 17:45:12'),
(1207,25,44,'2018-05-27 17:45:12'),(1208,25,45,'2018-05-27 17:45:12'),
(1217,25,69,'2018-05-27 17:45:12'),(1219,25,71,'2018-05-27 17:45:12'),
(1245,26,75,'2018-06-22 17:44:31'),(1246,26,35,'2018-06-22 17:44:31'),
(1247,26,76,'2018-06-22 17:44:31'),(1248,26,21,'2018-06-22 17:44:31'),
(1249,26,77,'2018-06-22 17:44:31'),(1250,26,78,'2018-06-22 17:44:31'),
(1251,26,79,'2018-06-22 17:44:31'),(1252,26,81,'2018-06-22 17:44:31'),
(1253,26,82,'2018-06-22 17:44:31'),(1254,26,41,'2018-06-22 17:44:31'),
(1255,26,44,'2018-06-22 17:44:31'),(1256,26,45,'2018-06-22 17:44:31'),
(1267,26,69,'2018-06-22 17:44:31'),(1269,26,71,'2018-06-22 17:44:31'),
(1270,26,87,'2018-06-22 17:44:31'),(1271,22,75,'2018-06-22 17:44:59'),
(1272,22,35,'2018-06-22 17:44:59'),(1273,22,76,'2018-06-22 17:44:59'),
(1274,22,21,'2018-06-22 17:44:59'),(1275,22,77,'2018-06-22 17:44:59'),
(1276,22,78,'2018-06-22 17:44:59'),(1277,22,79,'2018-06-22 17:44:59'),
(1278,22,80,'2018-06-22 17:44:59'),(1279,22,81,'2018-06-22 17:44:59'),
(1280,22,82,'2018-06-22 17:44:59'),(1281,22,83,'2018-06-22 17:44:59'),
(1282,22,41,'2018-06-22 17:44:59'),(1283,22,44,'2018-06-22 17:44:59'),
(1284,22,48,'2018-06-22 17:44:59'),(1285,22,49,'2018-06-22 17:44:59'),
(1286,22,45,'2018-06-22 17:44:59'),(1287,22,50,'2018-06-22 17:44:59'),
(1288,22,51,'2018-06-22 17:44:59'),(1313,22,69,'2018-06-22 17:44:59'),
(1315,22,71,'2018-06-22 17:44:59'),(1316,22,87,'2018-06-22 17:44:59'),
(1317,22,8,'2018-06-22 17:44:59'),(1318,22,7,'2018-06-22 17:44:59'),
(1319,22,9,'2018-06-22 17:44:59'),(1320,22,25,'2018-06-22 17:44:59'),
(1321,22,27,'2018-06-22 17:44:59'),(1322,22,12,'2018-06-22 17:44:59'),
(1323,22,30,'2018-06-22 17:44:59'),(1324,22,31,'2018-06-22 17:44:59'),
(1325,22,32,'2018-06-22 17:44:59'),(1326,22,13,'2018-06-22 17:44:59'),
(1327,22,26,'2018-06-22 17:44:59'),(1328,22,29,'2018-06-22 17:44:59'),
(1329,22,72,'2018-06-22 17:44:59'),(1330,22,73,'2018-06-22 17:44:59'),
(1331,22,74,'2018-06-22 17:44:59'),(1332,22,14,'2018-06-22 17:44:59'),
(1333,22,84,'2018-06-22 17:44:59'),(1334,22,85,'2018-06-22 17:44:59'),
(1335,22,86,'2018-06-22 17:44:59'),(1336,22,66,'2018-06-22 17:44:59'),
(1337,22,67,'2018-06-22 17:44:59'),(1338,22,68,'2018-06-22 17:44:59'),
(1339,22,23,'2018-06-22 17:44:59');
/*!40000 ALTER TABLE `sys_role_resource_rel` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_val`
--

DROP TABLE IF EXISTS `sys_val`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sys_val` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '序号',
  `code` varchar(20) DEFAULT NULL COMMENT '代码',
  `desc` varchar(60) DEFAULT NULL COMMENT '描述',
  `value` varchar(20) DEFAULT NULL COMMENT '变量',
  `uplimit` varchar(20) DEFAULT NULL COMMENT '上限',
  `step` varchar(10) DEFAULT NULL COMMENT '步长',
  `tag` int(11) DEFAULT '0' COMMENT '状态:0.正常1.删除',
  `createuser` varchar(20) DEFAULT NULL COMMENT '创建人',
  `createdate` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `changeuser` varchar(20) DEFAULT NULL COMMENT '修改人',
  `changedate` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_val`
--

LOCK TABLES `sys_val` WRITE;
/*!40000 ALTER TABLE `sys_val` DISABLE KEYS */;
INSERT INTO `sys_val` VALUES (1,'customerno','客户编号','1014','9999','1',0,'root','2018-05-24 17:20:35','超级管理员','2018-05-27 14:00:24'),
(2,'dtuno','DTU编号','900000000042','999999999999','1',1,'root','2018-05-24 17:22:08','超级管理员','2018-05-27 15:33:08'),
(3,'gatewayno','通讯方式编号','1002','9999','1',0,'root','2018-05-24 17:25:53','root','2018-05-24 17:25:53'),
(4,'metertypeno','电表类型编号','1009','9999','1',0,'root','2018-05-24 17:26:41','root','2018-05-24 17:26:41'),
(5,'vendorno','设备供应商编号','1005','9999','1',0,'root','2018-05-24 17:28:00','root','2018-05-24 17:28:00');
/*!40000 ALTER TABLE `sys_val` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Temporary table structure for view `v_sys_backend_conf`
--

DROP TABLE IF EXISTS `v_sys_backend_conf`;
/*!50001 DROP VIEW IF EXISTS `v_sys_backend_conf`*/;
SET @saved_cs_client     = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE VIEW `v_sys_backend_conf` AS SELECT 
 1 AS `id`,
 1 AS `appname`,
 1 AS `appversion`,
 1 AS `deploy`,
 1 AS `configtext`*/;
SET character_set_client = @saved_cs_client;

--
-- Final view structure for view `v_sys_backend_conf`
--

/*!50001 DROP VIEW IF EXISTS `v_sys_backend_conf`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_general_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`%` SQL SECURITY DEFINER */
/*!50001 VIEW `v_sys_backend_conf` AS select `sys_backend_conf`.`id` AS `id`,`sys_backend_conf`.`appname` AS `appname`,`sys_backend_conf`.`appversion` AS `appversion`,`sys_backend_conf`.`deploy` AS `deploy`,`sys_backend_conf`.`configtext` AS `configtext` from `sys_backend_conf` where (`sys_backend_conf`.`tag` = 0) */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-12-11 17:39:21
