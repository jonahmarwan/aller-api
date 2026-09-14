package multi

import (
	"bufio"
	"os"
	"strings"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func ReadMeta(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	win16le := unicode.UTF16(unicode.LittleEndian, unicode.ExpectBOM)
	decoder := win16le.NewDecoder()

	utf8reader := transform.NewReader(file, decoder)

	isBeginning := true
	resultBuff := []string{}
	scanner := bufio.NewScanner(utf8reader)
	for scanner.Scan() {
		line := scanner.Bytes()
		if isBeginning {
			if !(line[0] == 0x23) {
				isBeginning = false
				continue
			}
			resultBuff = append(resultBuff, string(line[1:]))
		}
	}
	result := strings.Join(resultBuff, "\n")
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return result, nil
}
