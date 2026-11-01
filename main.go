package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fileFlag := flag.String("f", "", "Путь к файлу для дампа")
	noAscii := flag.Bool("a", false, "Отключить вывод ASCII-колонки")

	flag.Parse()

	if *fileFlag != "" {
		ProcessFile(*fileFlag, *noAscii)
		return
	}

	remainingFlags := flag.Args()

	if len(remainingFlags) > 0 {
		data, err := DecodeHex(remainingFlags)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		dump := HexDump(data, *noAscii, 0)

		fmt.Println(dump)
		return
	}

	fmt.Println("Ошибка: не указаны данные для дампа")
	fmt.Fprintln(os.Stderr, "Использование:\n  hexdump -f <путь_к_файлу>  - для работы с файлом\n  hexdump <hex_строка>        - для работы со строкой")
	os.Exit(1)
}

func DecodeHex(args []string) ([]byte, error) {
	raw := strings.Join(args, " ")
	words := strings.Fields(raw)
	cleared := strings.Join(words, "")

	if len(cleared) == 0 {
		return nil, fmt.Errorf("error: No hex string provided.")
	}

	data, err := hex.DecodeString(cleared)

	if err != nil {
		return nil, fmt.Errorf("error decoding hex string: %w", err)
	}

	return data, nil
}

func ProcessFile(path string, noAscii bool) {
	file, err := os.Open(path)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "Ошибка: файл '%s' не найден", path)
		} else {
			fmt.Fprintln(os.Stderr, "Ошибка открытия файла: ", err)
		}
		os.Exit(1)
	}

	defer file.Close()

	buffer := make([]byte, 4096)
	globalOffset := 0

	for {
		n, err := file.Read(buffer)
		if n > 0 {
			chunk := buffer[:n]
			dump := HexDump(chunk, noAscii, globalOffset)

			fmt.Print(dump)

			globalOffset += n
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка чтения файла:", err)
			os.Exit(1)
		}
	}
}

func HexDump(b []byte, noAscii bool, baseOffset int) string {
	var sb strings.Builder
	for i := 0; i < len(b); i += 16 {
		end := i + 16
		if end > len(b) {
			end = len(b)
		}
		line := b[i:end]
		fmt.Fprintf(&sb, "%08x: ", baseOffset+i)
		for j := 0; j < 16; j++ {
			if j < len(line) {
				fmt.Fprintf(&sb, "%02x ", line[j])
			} else {
				fmt.Fprintf(&sb, "   ")
			}
		}
		if !noAscii {
			fmt.Fprintf(&sb, "| ")
			for _, c := range line {
				if c >= 32 && c <= 126 {
					sb.WriteByte(c)
				} else {
					sb.WriteByte('.')
				}
			}
			sb.WriteByte('|')
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
