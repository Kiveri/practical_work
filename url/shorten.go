package main

import (
	"encoding/base64"
	"fmt"
	"hash/crc32"
)

// ShortenURL сокращает длинный URL с использованием CRC32
func ShortenURL(longURL string) string {
	// Вычисляем контрольную сумму CRC32
	crc32Hash := crc32.ChecksumIEEE([]byte(longURL))

	// Преобразуем контрольную сумму в строку (base64 или hex)
	// Здесь используется base64 для более компактного представления
	hashString := base64.URLEncoding.EncodeToString([]byte(fmt.Sprintf("%08x", crc32Hash)))

	// Укоротим строку до первых 6-8 символов (опционально)
	shortHash := hashString[:10]

	return shortHash
}

func main() {
	longURL := "https://example.com/some/very/long/url/with/parameters?query=123"
	shortURL := ShortenURL(longURL)

	fmt.Printf("Original URL: %s\n", longURL)
	fmt.Printf("Shortened URL: %s\n", shortURL)
}
