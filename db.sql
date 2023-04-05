-- MySQL dump 10.13  Distrib 8.0.32, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: dbsearch
-- ------------------------------------------------------
-- Server version	8.0.32

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `mydb`
--

DROP TABLE IF EXISTS `mydb`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `mydb` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `password` varchar(255) NOT NULL,
  `email` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `t1_password_username` (`username`,`password`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mydb`
--

LOCK TABLES `mydb` WRITE;
/*!40000 ALTER TABLE `mydb` DISABLE KEYS */;
INSERT INTO `mydb` VALUES (1,'校长','123456','1346466@qq.com'),(2,'校董','123457','4565564@qq.com'),(3,'p7wmHucW','8sJ9PzTcT2wo','user1@example.com'),(4,'JnCJr1q3','Wk9XByKbhzGZ','user2@example.com'),(5,'u8ukYKFS','zR7DjK94XyXf','user3@example.com'),(6,'3q6drrZP','e6hKkPQkUbJG','user4@example.com'),(7,'aKbJfKK6','M39thYFNGvgj','user5@example.com'),(8,'X9JDGnTw','jKvrZvjDhZjk','user6@example.com'),(9,'5WzFJCYt','2RVx5rLZQ2DE','user7@example.com'),(10,'wNhHP97y','pShDvKyx9aRf','user8@example.com'),(11,'Yd3CqjwK','m8rk7NMpGmEd','user9@example.com'),(12,'fThZJtPv','n7wHmbGhBrM5','user10@example.com');
/*!40000 ALTER TABLE `mydb` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-04-05 11:02:08
