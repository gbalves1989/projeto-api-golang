package utils

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

func GenerateUniqueFileName(originalFileName string) string {
	ext := filepath.Ext(originalFileName)
	fileNameWithoutExt := strings.TrimSuffix(originalFileName, ext)
	uniqueID := uuid.New().String()
	timestamp := time.Now().Format("20060102150405")
	safeFileName := sanitizeFileName(fileNameWithoutExt)

	return fmt.Sprintf("%s_%s_%s%s", safeFileName, timestamp, uniqueID, ext)
}

func sanitizeFileName(fileName string) string {
	s := strings.ReplaceAll(fileName, " ", "_")
	s = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_' || r == '-' {
			return r
		}

		return -1
	}, s)

	return s
}
