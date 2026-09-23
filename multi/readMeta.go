package multi

import (
	"bufio"
	"io"
	"os"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/encoding/unicode/utf32"
	"golang.org/x/text/transform"
)

func ReadMeta(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	head := make([]byte, 4)
	_, err = file.Read(head)
	if err != nil && err != io.EOF {
		return nil, err
	}
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}
	encoding := DetectBOM(head)
	var scanner *bufio.Scanner
	switch encoding {
	case "UTF-16LE":
		win16le := unicode.UTF16(unicode.LittleEndian, unicode.ExpectBOM)
		decoder := win16le.NewDecoder()

		utf8reader := transform.NewReader(file, decoder)
		scanner = bufio.NewScanner(utf8reader)
	case "UTF-16BE":
		win16be := unicode.UTF16(unicode.BigEndian, unicode.ExpectBOM)
		decoder := win16be.NewDecoder()

		utf8reader := transform.NewReader(file, decoder)
		scanner = bufio.NewScanner(utf8reader)
	case "UTF-32LE":
		win32le := utf32.UTF32(utf32.LittleEndian, utf32.ExpectBOM)
		decoder := win32le.NewDecoder()

		utf8reader := transform.NewReader(file, decoder)
		scanner = bufio.NewScanner(utf8reader)
	case "UTF-32BE":
		win32be := utf32.UTF32(utf32.BigEndian, utf32.ExpectBOM)
		decoder := win32be.NewDecoder()

		utf8reader := transform.NewReader(file, decoder)
		scanner = bufio.NewScanner(utf8reader)
	case "UTF-8":
		scanner = bufio.NewScanner(file)
	default:
		panic("Unsupported encoding: " + encoding)
	}

	isBeginning := true
	result := []string{}
	for scanner.Scan() {
		line := scanner.Bytes()
		var numchar int = 0

		if isBeginning {
			if len(line) > 0 {
				switch line[0] {
				case 0x23:
					numchar = 1

				case 0x2F:
					numchar = 2

				default:
					isBeginning = false
					continue
				}
				if line[numchar] == 0x20 {
					numchar += 1
				}
			} else {
				continue
			}
			result = append(result, string(line[numchar:]))
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
