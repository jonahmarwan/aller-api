package multi

import (
	"bufio"
	"os"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func ReadMeta(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	win16le := unicode.UTF16(unicode.LittleEndian, unicode.ExpectBOM)
	decoder := win16le.NewDecoder()

	utf8reader := transform.NewReader(file, decoder)

	isBeginning := true
	result := []string{}
	scanner := bufio.NewScanner(utf8reader)
	for scanner.Scan() {
		line := scanner.Bytes()
		var numchar int = 0

		//FIX SYNTAX
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
