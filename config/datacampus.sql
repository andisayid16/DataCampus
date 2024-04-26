-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Apr 26, 2024 at 04:41 AM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `datacampus`
--

-- --------------------------------------------------------

--
-- Table structure for table `dosen`
--

CREATE TABLE `dosen` (
  `id` int(11) NOT NULL,
  `nama_lengkap` varchar(100) NOT NULL,
  `nip` varchar(16) NOT NULL,
  `jenis_kelamin` tinyint(1) NOT NULL,
  `tempat_lahir` varchar(100) NOT NULL,
  `tanggal_lahir` date NOT NULL,
  `alamat` text NOT NULL,
  `no_hp` varchar(18) NOT NULL,
  `fakultas` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `dosen`
--

INSERT INTO `dosen` (`id`, `nama_lengkap`, `nip`, `jenis_kelamin`, `tempat_lahir`, `tanggal_lahir`, `alamat`, `no_hp`, `fakultas`) VALUES
(4, 'John Doe', '1234567890', 1, 'Jakarta', '2000-01-01', 'Jl. Merdeka No. 1', '081234567890', 'Teknik Informatika'),
(5, 'Jane Doe', '0987654321', 2, 'Bandung', '1995-05-05', 'Jl. Braga No. 2', '081234567890', 'Ekonomi dan Bisnis'),
(6, 'Alice Doe', '2345678901', 2, 'Surabaya', '1998-08-08', 'Jl. Sudirman No. 3', '081234567890', 'Ilmu Komputer'),
(7, 'Bob Doe', '3456789012', 1, 'Medan', '1990-11-11', 'Jl. Thamrin No. 4', '081234567890', 'Teknik Elektro'),
(8, 'Charlie Doe', '4567890123', 1, 'Yogyakarta', '1993-02-02', 'Jl. Malioboro No. 5', '081234567890', 'Teknik Mesin'),
(9, 'David Doe', '5678901234', 1, 'Semarang', '1996-05-05', 'Jl. Pemuda No. 6', '081234567890', 'Teknik Sipil'),
(10, 'Eve Doe', '6789012345', 2, 'Palembang', '1999-08-08', 'Jl. Sudirman No. 7', '081234567890', 'Teknik Kimia'),
(11, 'Frank Doe', '7890123456', 1, 'Makassar', '1992-11-11', 'Jl. Ahmad Yani No. 8', '081234567890', 'Teknik Industri'),
(12, 'Grace Doe', '8901234567', 2, 'Denpasar', '1995-02-02', 'Jl. Diponegoro No. 9', '081234567890', 'Teknik Pertanian'),
(13, 'Heidi Doe', '9012345678', 1, 'Pontianak', '1998-05-05', 'Jl. Gajah Mada No.10', '081234567890', 'Teknik Lingkungan'),
(14, 'Ivan Doe', '0123456789', 1, 'Manado', '1991-11-11', 'Jl. Sam Ratulangi No. 11', '081234567890', 'Teknik Perkapalan'),
(15, 'Judy Doe', '1234567890', 1, 'Bali', '1996-05-05', 'Jl. Monas No. 12', '081234567890', 'Teknik Geologi'),
(16, 'Kevin Doe', '2345678901', 1, 'Lombok', '1993-02-02', 'Jl. Sumbawa No. 13', '081234567890', 'Teknik Informatika'),
(17, 'Lily Doe', '3456789012', 2, 'Batam', '1990-11-11', 'Jl. Barelang No. 14', '081234567890', 'Teknik Elektro'),
(18, 'Mike Doe', '4567890123', 1, 'Pekanbaru', '1997-05-05', 'Jl. Riau No. 15', '081234567890', 'Teknik Mesin'),
(19, 'Nancy Doe', '5678901234', 2, 'Padang', '1994-02-02', 'Jl. Sumatera No. 16', '081234567890', 'Teknik Sipil'),
(20, 'Oliver Doe', '6789012345', 2, 'Banda Aceh', '1991-11-11', 'Jl. Aceh No.17', '081234567890', 'Teknik Kimia'),
(21, 'Penny Doe', '7890123456', 2, 'Jayapura', '1998-05-05', 'Jl. Papua No. 18', '081234567890', 'Teknik Industri'),
(22, 'Quinn Doe', '8901234567', 2, 'Yogyakarta', '1995-02-02', 'Jl. Malioboro No. 19', '081234567890', 'Teknik Pertanian'),
(23, 'Roy Doe', '9012345678', 1, 'Bali', '1992-11-11', 'Jl. Kuta No. 20', '081234567890', 'Teknik Lingkungan');

-- --------------------------------------------------------

--
-- Table structure for table `fakultas`
--

CREATE TABLE `fakultas` (
  `id` int(11) NOT NULL,
  `data_fakultas` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `fakultas`
--

INSERT INTO `fakultas` (`id`, `data_fakultas`) VALUES
(3, 'TEKNIK INFORMATIKA'),
(4, 'Farmasi'),
(5, 'Farmasi'),
(6, 'Kesenian'),
(7, 'Judo'),
(8, 'Kesenian');

-- --------------------------------------------------------

--
-- Table structure for table `mahasiswa`
--

CREATE TABLE `mahasiswa` (
  `id` int(11) NOT NULL,
  `nama_lengkap` varchar(50) NOT NULL,
  `nim` varchar(20) NOT NULL,
  `jenis_kelamin` tinyint(1) NOT NULL,
  `tempat_lahir` varchar(100) NOT NULL,
  `tanggal_lahir` date NOT NULL,
  `alamat` text NOT NULL,
  `no_hp` varchar(18) NOT NULL,
  `nama_fakultas` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `mahasiswa`
--

INSERT INTO `mahasiswa` (`id`, `nama_lengkap`, `nim`, `jenis_kelamin`, `tempat_lahir`, `tanggal_lahir`, `alamat`, `no_hp`, `nama_fakultas`) VALUES
(7, 'riris pratiwi', '121341212', 2, 'Balikpapan', '1997-02-02', 'JL. JENDRAL A. YANI RT. 20 NO. 8 KEL. KARANG REJO BALIKPAPAN', '081346320751', 'Farmasi'),
(8, 'riris pratiwi', '121341212', 2, 'Balikpapan', '1997-02-02', 'JL. JENDRAL A. YANI RT. 20 NO. 8 KEL. KARANG REJO BALIKPAPAN', '081346320751', 'Farmasi');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(10) UNSIGNED NOT NULL,
  `nama_lengkap` varchar(50) NOT NULL,
  `email` varchar(50) NOT NULL,
  `username` varchar(30) NOT NULL,
  `password` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `nama_lengkap`, `email`, `username`, `password`) VALUES
(5, 'suyadi', 'suyadi@mail.com', 'suyadi', '$2a$10$FWTqTxWUL7qkVlopHhdHyu6KiVMrIsj.cOXd3Oe83jyCXzn0OOzi6');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `dosen`
--
ALTER TABLE `dosen`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `fakultas`
--
ALTER TABLE `fakultas`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `mahasiswa`
--
ALTER TABLE `mahasiswa`
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
-- AUTO_INCREMENT for table `dosen`
--
ALTER TABLE `dosen`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=25;

--
-- AUTO_INCREMENT for table `fakultas`
--
ALTER TABLE `fakultas`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `mahasiswa`
--
ALTER TABLE `mahasiswa`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
