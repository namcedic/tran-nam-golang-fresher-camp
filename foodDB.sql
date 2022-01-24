CREATE DATABASE food_delivery;
USE food_delivery;



CREATE TABLE `foods` (
  `id` int AUTO_INCREMENT,
  `restaurantID` int,
  `name` varchar(255),
  `price` int,
  `rated` float,
  `categoryID` int,
  `description` varchar(255),
  `create` timestamp,
  `update` timestamp,
  PRIMARY KEY (`id`, `restaurantID`)
);

CREATE TABLE `restaurants` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `addressID` int,
  `phone_number` int,
  `rated` float,
  `create` timestamp,
  `update` timestamp
);

CREATE TABLE `categories` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255)
);

CREATE TABLE `users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `full_name` varchar(255),
  `email` varchar(255),
  `addressID` int,
  `phone_number` int,
  `create` timestamp,
  `update` timestamp
);

CREATE TABLE `orders` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `userID` int UNIQUE NOT NULL,
  `status` varchar(255),
  `create` timestamp,
  `update` timestamp
);

CREATE TABLE `order_foods` (
  `orderID` int,
  `foodID` int,
  `quantity` int DEFAULT 1
);

CREATE TABLE `restaurants_categories` (
  `restaurantID` int,
  `categoryID` int,
  `create` timestamp,
  `update` timestamp
);

CREATE TABLE `addresses` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `full_name` varchar(255),
  `phone_number` int,
  `state` varchar(255),
  `city` varchar(255),
  `street` varchar(255),
  `adress` varchar(255)
);

CREATE TABLE `favorites_restaurants` (
  `userID` int,
  `restaurantID` int
);

CREATE TABLE `favorites_foods` (
  `userID` int,
  `foodID` int
);

CREATE TABLE `reviews_restaurants` (
  `id` int,
  `userID` int,
  `restaurantID` int,
  `content` varchar(255)
);

CREATE TABLE `reviews_foods` (
  `userID` int,
  `foodID` int,
  `content` varchar(255)
);

ALTER TABLE `foods` ADD FOREIGN KEY (`restaurantID`) REFERENCES `restaurants` (`id`);

ALTER TABLE `restaurants` ADD FOREIGN KEY (`addressID`) REFERENCES `addresses` (`id`);

ALTER TABLE `favorites_restaurants` ADD FOREIGN KEY (`userID`) REFERENCES `users` (`id`);

ALTER TABLE `favorites_restaurants` ADD FOREIGN KEY (`restaurantID`) REFERENCES `restaurants` (`id`);

ALTER TABLE `order_foods` ADD FOREIGN KEY (`foodID`) REFERENCES `foods` (`id`);

ALTER TABLE `orders` ADD FOREIGN KEY (`userID`) REFERENCES `users` (`id`);

ALTER TABLE `users` ADD FOREIGN KEY (`addressID`) REFERENCES `addresses` (`id`);

ALTER TABLE `restaurants_categories` ADD FOREIGN KEY (`categoryID`) REFERENCES `categories` (`id`);

ALTER TABLE `restaurants_categories` ADD FOREIGN KEY (`restaurantID`) REFERENCES `restaurants` (`id`);

ALTER TABLE `order_foods` ADD FOREIGN KEY (`orderID`) REFERENCES `orders` (`id`);

ALTER TABLE `foods` ADD FOREIGN KEY (`categoryID`) REFERENCES `categories` (`id`);

ALTER TABLE `favorites_foods` ADD FOREIGN KEY (`userID`) REFERENCES `users` (`id`);

ALTER TABLE `favorites_foods` ADD FOREIGN KEY (`foodID`) REFERENCES `foods` (`id`);

ALTER TABLE `reviews_restaurants` ADD FOREIGN KEY (`userID`) REFERENCES `users` (`id`);

ALTER TABLE `reviews_foods` ADD FOREIGN KEY (`userID`) REFERENCES `users` (`id`);

ALTER TABLE `reviews_foods` ADD FOREIGN KEY (`foodID`) REFERENCES `foods` (`id`);

ALTER TABLE `reviews_restaurants` ADD FOREIGN KEY (`restaurantID`) REFERENCES `restaurants` (`id`);
