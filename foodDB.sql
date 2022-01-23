create database food_delivery;
use food_delivery;

CREATE TABLE `foods` (
  `id` int AUTO_INCREMENT,
  `restaurantID` int,
  `name` varchar(100),
  `price` float,
  `rate` float,
  `categoryID` int,
  `description` varchar(255),
  `created` timestamp,
  `update` timestamp,
  PRIMARY KEY (`id`, `restaurantID`)
);

CREATE TABLE `restaurants` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(100),
  `addressID` int,
  `phone_number` int,
  `rate` float,
  `created` timestamp,
  `update` timestamp
);

CREATE TABLE `categories` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(100)
);

CREATE TABLE `users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `full_name` varchar(255),
  `email` varchar(255),
  `addressID` int,
  `phone_number` int,
  `created` timestamp,
  `update` timestamp,
  `role` tinyint
);

CREATE TABLE `orders` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `userID` int UNIQUE NOT NULL,
  `status` varchar(255),
  `created` timestamp,
  `update` timestamp,
  `foodID` int,
  `quantity` int DEFAULT 1
);

CREATE TABLE `restaurants_categories` (
  `restaurantID` int,
  `categoryID` int,
  `created` timestamp,
  `update` timestamp
);

CREATE TABLE `addresses` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `full_name` varchar(255),
  `phone_number` int,
  `state` varchar(80),
  `city` varchar(50),
  `street` varchar(100),
  `address` varchar(100)
);

CREATE TABLE `favorites_restaurants` (
  `userID` int,
  `restaurantID` int
);

CREATE TABLE `favorites_foods` (
  `userID` int,
  `foodID` int
);

CREATE TABLE `review_restaurants` (
  `userID` int,
  `restaurantID` int,
  `content` varchar(255),
  `created` timestamp,
  `update` timestamp
);

CREATE TABLE `review_foods` (
  `userID` int,
  `foodID` int,
  `content` varchar(255),
  `created` timestamp,
  `update` timestamp
);

 ALTER TABLE `restaurants` ADD FOREIGN KEY (`id`) REFERENCES `foods` (`restaurantID`); 

ALTER TABLE `orders` ADD FOREIGN KEY (`id`) REFERENCES `order_foods` (`orderID`);

ALTER TABLE `foods` ADD FOREIGN KEY (`id`) REFERENCES `order_foods` (`foodID`);

ALTER TABLE `users` ADD FOREIGN KEY (`id`) REFERENCES `orders` (`userID`);

ALTER TABLE `foods` ADD FOREIGN KEY (`categoryID`) REFERENCES `categories` (`id`);

ALTER TABLE `restaurants` ADD FOREIGN KEY (`id`) REFERENCES `restaurants_categories` (`restaurantID`);

ALTER TABLE `categories` ADD FOREIGN KEY (`id`) REFERENCES `restaurants_categories` (`categoryID`);

ALTER TABLE `addresses` ADD FOREIGN KEY (`id`) REFERENCES `restaurants` (`addressID`);

ALTER TABLE `addresses` ADD FOREIGN KEY (`id`) REFERENCES `users` (`addressID`);

ALTER TABLE `users` ADD FOREIGN KEY (`id`) REFERENCES `favorites_restaurants` (`userID`);

ALTER TABLE `users` ADD FOREIGN KEY (`id`) REFERENCES `favorites_foods` (`userID`);

ALTER TABLE `restaurants` ADD FOREIGN KEY (`id`) REFERENCES `favorites_restaurants` (`restaurantID`);

ALTER TABLE `foods` ADD FOREIGN KEY (`id`) REFERENCES `favorites_foods` (`foodID`);

ALTER TABLE `users` ADD FOREIGN KEY (`id`) REFERENCES `review_foods` (`userID`);

ALTER TABLE `users` ADD FOREIGN KEY (`id`) REFERENCES `review_restaurants` (`userID`);

ALTER TABLE `restaurants` ADD FOREIGN KEY (`id`) REFERENCES `review_restaurants` (`restaurantID`);

ALTER TABLE `foods` ADD FOREIGN KEY (`id`) REFERENCES `review_foods` (`foodID`);
