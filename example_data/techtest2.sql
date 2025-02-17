-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Feb 17, 2025 at 01:05 AM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `techtest2`
--

-- --------------------------------------------------------

--
-- Table structure for table `transaction`
--

CREATE TABLE `transaction` (
  `IDTransaction` varchar(30) NOT NULL,
  `IDAlias` varchar(30) DEFAULT NULL,
  `IDUser` varchar(30) DEFAULT NULL,
  `IDReference` varchar(30) DEFAULT NULL,
  `Type` varchar(50) DEFAULT NULL,
  `TransactionType` varchar(50) DEFAULT NULL,
  `Amount` int(20) DEFAULT NULL,
  `Remarks` text DEFAULT NULL,
  `Status` varchar(20) DEFAULT NULL,
  `BalanceStart` int(20) DEFAULT NULL,
  `BalanceEnd` int(20) DEFAULT NULL,
  `CreateDate` datetime DEFAULT NULL,
  `UpdateDate` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `transaction`
--

INSERT INTO `transaction` (`IDTransaction`, `IDAlias`, `IDUser`, `IDReference`, `Type`, `TransactionType`, `Amount`, `Remarks`, `Status`, `BalanceStart`, `BalanceEnd`, `CreateDate`, `UpdateDate`) VALUES
('1a297be1-bc67-4e9d-95ee-be51e0', 'transfer_id', '', '', 'Topup', 'KREDIT', 50000, '', 'Success', 0, 50000, '2025-02-16 23:33:53', '2025-02-16 23:33:53'),
('4ca29bc5-285f-4663-a480-9fe0a9', 'transfer_id', '', '', 'Topup', 'KREDIT', 50000, '', 'Pending', 0, 50000, '2025-02-16 23:35:32', '2025-02-16 23:35:32'),
('57881c7a-81d7-4958-93ec-7b94ab', 'payment_id', '', '', 'Payment', 'DEBIT', 10000, 'Bayar Listrik', 'Pending', 50000, 40000, '2025-02-16 23:37:05', '2025-02-16 23:37:05'),
('c06fd743-361e-4751-9324-959e1e', 'transfer_id', '', '', 'Topup', 'KREDIT', 50000, '', 'Pending', 0, 50000, '2025-02-16 23:37:05', '2025-02-16 23:37:05'),
('ec2efaf8-56fb-4067-99c2-c04584', 'payment_id', '', '', 'Payment', 'DEBIT', 10000, 'Bayar Listrik', 'Pending', 50000, 40000, '2025-02-16 23:33:53', '2025-02-16 23:33:53'),
('ecaab02e-df27-46c9-9a5e-caa160', 'payment_id', '', '', 'Payment', 'DEBIT', 10000, 'Bayar Listrik', 'Pending', 50000, 40000, '2025-02-16 23:35:32', '2025-02-16 23:35:32');

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `IDUser` varchar(30) NOT NULL,
  `FirstName` varchar(50) DEFAULT NULL,
  `LastName` varchar(50) DEFAULT NULL,
  `Phone` varchar(20) DEFAULT NULL,
  `Address` text DEFAULT NULL,
  `Pin` text DEFAULT NULL,
  `CreateDate` datetime DEFAULT NULL,
  `ModifyDate` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`IDUser`, `FirstName`, `LastName`, `Phone`, `Address`, `Pin`, `CreateDate`, `ModifyDate`) VALUES
('236b4108-bf64-4956-b737-93ad66', 'Bramantio', 'Galih', '081998998155', 'Jakarta', '9267420324f5cbf1c231096eff7f75ee', '2025-02-16 23:33:47', '2025-02-16 23:33:47'),
('d55c1af0-4559-48ee-b16c-d23793', 'Norman', 'Osbourne', '081998998155', 'Jakarta', '9267420324f5cbf1c231096eff7f75ee', '2025-02-16 23:33:47', '2025-02-16 23:33:47');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `transaction`
--
ALTER TABLE `transaction`
  ADD PRIMARY KEY (`IDTransaction`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`IDUser`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
