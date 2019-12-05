CREATE TABLE `employees` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `city` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4;

INSERT INTO `employees` (`id`, `name`, `city`)
VALUES
	(1, 'Yoga Hanggara', 'Yogyakarta'),
	(2, 'Aldy Ginanjar', 'Bandung');
