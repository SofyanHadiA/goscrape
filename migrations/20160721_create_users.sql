CREATE TABLE `users` (
 `id` int(11) NOT NULL,
 `role_id` int(11) NOT NULL,
 `name` varchar(100) DEFAULT NULL,
 `address` varchar(100) DEFAULT NULL,
 `password` varchar(100) DEFAULT NULL,
 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1