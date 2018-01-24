-- MySQL dump 10.13  Distrib 5.7.20, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: mydb
-- ------------------------------------------------------
-- Server version	5.5.5-10.1.21-MariaDB

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
-- Table structure for table `bug`
--

DROP TABLE IF EXISTS `bug`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `bug` (
  `idBug` int(11) NOT NULL AUTO_INCREMENT,
  `bugName` varchar(90) NOT NULL,
  `bugDescription` varchar(150) NOT NULL,
  `category` varchar(45) NOT NULL,
  `solutionDescription` varchar(150) NOT NULL,
  `idDev` int(11) NOT NULL,
  `idTest` int(11) NOT NULL,
  `idProject` int(11) NOT NULL,
  `foundDate` date NOT NULL,
  `updateDate` date NOT NULL,
  PRIMARY KEY (`idBug`,`idDev`,`idTest`,`idProject`),
  KEY `fk_Bug_User1_idx` (`idTest`),
  KEY `fk_Bug_Project1_idx` (`idProject`),
  CONSTRAINT `fk_Bug_Project1` FOREIGN KEY (`idProject`) REFERENCES `project` (`idProject`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_Bug_User1` FOREIGN KEY (`idTest`) REFERENCES `user` (`idUser`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bug`
--

LOCK TABLES `bug` WRITE;
/*!40000 ALTER TABLE `bug` DISABLE KEYS */;
INSERT INTO `bug` VALUES (15,'adf','d','Dont know tech','abc',36,29,27,'0000-00-00','2017-12-21'),(19,'1','12','Dont know tech','12',34,29,26,'2017-12-01','2017-12-05'),(21,'2','223','Carelessness','',40,29,26,'2017-12-01','2017-12-07'),(22,'231','231','Logic','dsffa',36,29,28,'2017-12-07','2017-12-13'),(23,'123','123','Dont know tech','',40,29,30,'2017-12-07','0000-00-00'),(24,'123','kkj','Carelessness','',37,29,29,'2017-12-08','2017-12-11'),(25,'gjgk','.jkljkl','Dont know tech','',37,29,30,'2017-12-08','0000-00-00'),(26,'test ','ashe debug','Logic','good',36,29,26,'2017-12-08','2017-12-08'),(27,'d','d','Carelessness','',0,29,30,'2017-12-08','0000-00-00'),(28,'123','123','Dont understand spec','dafasf',36,29,28,'2017-12-11','2017-12-13'),(29,'ahaha','ahehe','Carelessness','',0,29,29,'2017-12-11','0000-00-00'),(30,'abc','cba','Logic','',0,29,29,'2017-12-11','0000-00-00'),(31,'mnb','bnm','Dont know tech','',0,29,28,'2017-12-11','0000-00-00'),(32,'Nham idproject voi idbug','khong delete duoc','Logic','sua lai kiem tra lai user truoc khi cho vao',36,29,26,'2017-12-13','2017-12-13');
/*!40000 ALTER TABLE `bug` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `company`
--

DROP TABLE IF EXISTS `company`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `company` (
  `idCompany` int(11) NOT NULL AUTO_INCREMENT,
  `companyDomain` varchar(45) NOT NULL,
  `status` int(11) NOT NULL,
  PRIMARY KEY (`idCompany`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `company`
--

LOCK TABLES `company` WRITE;
/*!40000 ALTER TABLE `company` DISABLE KEYS */;
INSERT INTO `company` VALUES (26,'123456',1);
/*!40000 ALTER TABLE `company` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `goose_db_version`
--

DROP TABLE IF EXISTS `goose_db_version`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `goose_db_version` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `version_id` bigint(20) NOT NULL,
  `is_applied` tinyint(1) NOT NULL,
  `tstamp` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goose_db_version`
--

LOCK TABLES `goose_db_version` WRITE;
/*!40000 ALTER TABLE `goose_db_version` DISABLE KEYS */;
INSERT INTO `goose_db_version` VALUES (1,0,1,'2017-12-13 08:09:47'),(2,1,1,'2017-12-13 08:09:47');
/*!40000 ALTER TABLE `goose_db_version` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `master`
--

DROP TABLE IF EXISTS `master`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `master` (
  `idmaster` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(45) NOT NULL,
  `password` varchar(45) NOT NULL,
  PRIMARY KEY (`idmaster`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `master`
--

LOCK TABLES `master` WRITE;
/*!40000 ALTER TABLE `master` DISABLE KEYS */;
INSERT INTO `master` VALUES (1,'master','master');
/*!40000 ALTER TABLE `master` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `position`
--

DROP TABLE IF EXISTS `position`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `position` (
  `idPosition` int(11) NOT NULL,
  `positionName` varchar(45) NOT NULL,
  PRIMARY KEY (`idPosition`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `position`
--

LOCK TABLES `position` WRITE;
/*!40000 ALTER TABLE `position` DISABLE KEYS */;
INSERT INTO `position` VALUES (0,'Admin'),(1,'Developer'),(2,'Tester');
/*!40000 ALTER TABLE `position` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `project`
--

DROP TABLE IF EXISTS `project`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `project` (
  `idProject` int(11) NOT NULL AUTO_INCREMENT,
  `projectName` varchar(100) NOT NULL,
  `projectDescription` varchar(150) NOT NULL,
  `beginDate` date NOT NULL,
  `finishDate` date NOT NULL,
  PRIMARY KEY (`idProject`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `project`
--

LOCK TABLES `project` WRITE;
/*!40000 ALTER TABLE `project` DISABLE KEYS */;
INSERT INTO `project` VALUES (26,'GoBug','Quan li bug va project','2017-01-05','2017-04-30'),(27,'BugManage','Nhu tren','2017-10-01','2017-11-30'),(28,'abc','abc','2017-12-07','2017-12-07'),(29,'123','123','2017-12-08','2017-12-08'),(30,'test','test','2017-12-03','2017-12-03');
/*!40000 ALTER TABLE `project` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `idUser` int(11) NOT NULL AUTO_INCREMENT,
  `userName` varchar(45) NOT NULL,
  `email` varchar(45) NOT NULL,
  `password` varchar(45) NOT NULL,
  `idCompany` int(11) NOT NULL,
  `idPosition` int(11) NOT NULL,
  `status` int(11) NOT NULL,
  PRIMARY KEY (`idUser`,`idCompany`,`idPosition`),
  KEY `fk_User_CongTy_idx` (`idCompany`),
  KEY `fk_User_vaitro1_idx` (`idPosition`),
  CONSTRAINT `fk_User_CongTy` FOREIGN KEY (`idCompany`) REFERENCES `company` (`idCompany`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_User_vaitro1` FOREIGN KEY (`idPosition`) REFERENCES `position` (`idPosition`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (25,'yasua','daxua997@gmail.com','c4ca4238a0b923820dcc509a6f75849b',26,0,1),(29,'1234','findbear997@gmail.com','c4ca4238a0b923820dcc509a6f75849b',26,2,1),(34,'sydra','ha@gmail.com','c4ca4238a0b923820dcc509a6f75849b',26,1,1),(35,'garen','ha2@gmail.com','c4ca4238a0b923820dcc509a6f75849b',26,2,1),(36,'ashe','ha3@gmail.com','c4ca4238a0b923820dcc509a6f75849b',26,1,1),(37,'morgana','ha4@gmail.com','c4ca4238a0b923820dcc509a6f75849b',26,1,1),(40,'jinx','ha5@gmail.com','c4ca4238a0b923820dcc509a6f75849b',26,1,1);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_project`
--

DROP TABLE IF EXISTS `user_project`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_project` (
  `idUser` int(11) NOT NULL,
  `idProject` int(11) NOT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`,`idUser`,`idProject`),
  KEY `fk_User_has_Project_Project1_idx` (`idProject`),
  KEY `fk_User_has_Project_User1_idx` (`idUser`),
  CONSTRAINT `fk_User_has_Project_Project1` FOREIGN KEY (`idProject`) REFERENCES `project` (`idProject`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_User_has_Project_User1` FOREIGN KEY (`idUser`) REFERENCES `user` (`idUser`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=59 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_project`
--

LOCK TABLES `user_project` WRITE;
/*!40000 ALTER TABLE `user_project` DISABLE KEYS */;
INSERT INTO `user_project` VALUES (25,26,21),(29,26,23),(25,27,24),(25,28,32),(25,29,33),(25,30,34),(29,28,41),(34,26,42),(34,27,43),(35,27,44),(36,27,45),(36,28,47),(29,29,48),(37,29,50),(37,28,51),(29,30,52),(34,30,53),(36,30,54),(36,26,55),(37,27,56),(34,28,57),(35,28,58);
/*!40000 ALTER TABLE `user_project` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-12-23 22:34:45
