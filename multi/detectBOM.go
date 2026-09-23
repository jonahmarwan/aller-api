package multi

import (
	"bytes"
)

func DetectBOM(fileHead []byte) string {
	boms := map[string][]byte{
		"UTF-8":    {0xEF, 0xBB, 0xBF},
		"UTF-16LE": {0xFF, 0xFE},
		"UTF-16BE": {0xFE, 0xFF},
		"UTF-32LE": {0xFF, 0xFE, 0x00, 0x00},
		"UTF-32BE": {0x00, 0x00, 0xFE, 0xFF},
	}

	for name, bom := range boms {
		if bytes.HasPrefix(fileHead, bom) {
			return name
		}
	}
	return "UTF-8"
}
