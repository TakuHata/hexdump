package main

import (
	"strings"
	"testing"
)

func TestHexDump(t *testing.T) {
	// Передаем сырые байты (длина 14 байт — поместится в одну строку)
	input := []byte("Hello, World!\n")

	t.Run("With ASCII", func(t *testing.T) {
		result := HexDump(input, false, 0)

		if !strings.Contains(result, "00000000:") {
			t.Errorf("Ожидалось смещение 00000000")
		}
		// Проверяем, что буквы "H", "e", "l", "l", "o" перевелись в правильный хекс
		if !strings.Contains(result, "48 65 6c 6c 6f") {
			t.Errorf("Ожидались hex-значения для 'Hello'")
		}
		// Проверяем ASCII колонку с учетом вашего пробела в начале
		if !strings.Contains(result, "| Hello, World!.|") {
			t.Errorf("Ожидалась правильная ASCII колонка для строки")
		}
	})

	t.Run("No ASCII", func(t *testing.T) {
		result := HexDump(input, true, 0)
		if strings.Contains(result, "|") {
			t.Errorf("При noAscii=true символов '|' быть не должно")
		}
	})
}

func BenchmarkHexDump(b *testing.B) {
	data := make([]byte, 1024)
	for i := range data {
		data[i] = byte(i % 256)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = HexDump(data, false, 0)
	}
}
