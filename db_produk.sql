-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 21 Apr 2023 pada 06.43
-- Versi server: 10.4.24-MariaDB
-- Versi PHP: 8.1.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_produk`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `like_review`
--

CREATE TABLE `like_review` (
  `ID_REVIEW` int(11) NOT NULL,
  `ID_MEMBER` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `like_review`
--

INSERT INTO `like_review` (`ID_REVIEW`, `ID_MEMBER`) VALUES
(1, 1);

-- --------------------------------------------------------

--
-- Struktur dari tabel `member`
--

CREATE TABLE `member` (
  `ID_MEMBER` int(11) NOT NULL,
  `USERNAME` varchar(225) NOT NULL,
  `GENDER` varchar(225) NOT NULL,
  `SKINTYPE` varchar(225) NOT NULL,
  `SKINCOLOR` varchar(225) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `member`
--

INSERT INTO `member` (`ID_MEMBER`, `USERNAME`, `GENDER`, `SKINTYPE`, `SKINCOLOR`) VALUES
(1, 'Jamaludin21', 'Men', 'Exotic', 'Brown'),
(4, 'Lutfi28', 'Men', 'Exotic', 'Black');

-- --------------------------------------------------------

--
-- Struktur dari tabel `product`
--

CREATE TABLE `product` (
  `ID_PRODUCT` int(11) NOT NULL,
  `NAME_PRODUCT` varchar(225) NOT NULL,
  `PRICE` varchar(225) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `product`
--

INSERT INTO `product` (`ID_PRODUCT`, `NAME_PRODUCT`, `PRICE`) VALUES
(1, 'Nike Predator 2023', '5000000'),
(2, 'Jersey Aero Sport 2020', '2000000');

-- --------------------------------------------------------

--
-- Struktur dari tabel `review_product`
--

CREATE TABLE `review_product` (
  `ID_REVIEW` int(11) NOT NULL,
  `ID_PRODUCT` int(11) NOT NULL,
  `ID_MEMBER` int(11) NOT NULL,
  `DESC_REVIEW` varchar(225) NOT NULL,
  `JUMLAH_LIKE_REVIEW` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `review_product`
--

INSERT INTO `review_product` (`ID_REVIEW`, `ID_PRODUCT`, `ID_MEMBER`, `DESC_REVIEW`, `JUMLAH_LIKE_REVIEW`) VALUES
(1, 1, 1, 'Sepatu nya bagus banget dan awet, rekomendasi untuk pemakaian jangka panjang', 0),
(2, 2, 4, 'Baju nya adem banget', 0);

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `like_review`
--
ALTER TABLE `like_review`
  ADD KEY `ID_REVIEW` (`ID_REVIEW`),
  ADD KEY `LIKE_MEMBER` (`ID_MEMBER`);

--
-- Indeks untuk tabel `member`
--
ALTER TABLE `member`
  ADD PRIMARY KEY (`ID_MEMBER`);

--
-- Indeks untuk tabel `product`
--
ALTER TABLE `product`
  ADD PRIMARY KEY (`ID_PRODUCT`);

--
-- Indeks untuk tabel `review_product`
--
ALTER TABLE `review_product`
  ADD PRIMARY KEY (`ID_REVIEW`),
  ADD KEY `ID_PRODUCT` (`ID_PRODUCT`),
  ADD KEY `ID_MEMBER` (`ID_MEMBER`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `member`
--
ALTER TABLE `member`
  MODIFY `ID_MEMBER` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT untuk tabel `product`
--
ALTER TABLE `product`
  MODIFY `ID_PRODUCT` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT untuk tabel `review_product`
--
ALTER TABLE `review_product`
  MODIFY `ID_REVIEW` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `like_review`
--
ALTER TABLE `like_review`
  ADD CONSTRAINT `ID_REVIEW` FOREIGN KEY (`ID_REVIEW`) REFERENCES `review_product` (`ID_REVIEW`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `LIKE_MEMBER` FOREIGN KEY (`ID_MEMBER`) REFERENCES `review_product` (`ID_MEMBER`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Ketidakleluasaan untuk tabel `review_product`
--
ALTER TABLE `review_product`
  ADD CONSTRAINT `ID_MEMBER` FOREIGN KEY (`ID_MEMBER`) REFERENCES `member` (`ID_MEMBER`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `ID_PRODUCT` FOREIGN KEY (`ID_PRODUCT`) REFERENCES `product` (`ID_PRODUCT`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
