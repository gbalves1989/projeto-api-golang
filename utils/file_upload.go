package utils

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/gbalves1989/projeto-api-golang/constants"
)

func SaveUploadedFile(file *multipart.FileHeader, destinationDir string) string {
	src, err := file.Open()
	if err != nil {
		Log(ERROR, "Erro ao abrir arquivo enviado: %v", err)
		return ""
	}
	defer src.Close()

	err = os.MkdirAll(destinationDir, os.ModePerm)
	if err != nil {
		Log(ERROR, "Erro ao criar diretório de upload '%s': %v", destinationDir, err)
		return ""
	}

	uniqueFileName := GenerateUniqueFileName(file.Filename)
	filePath := filepath.Join(destinationDir, uniqueFileName)

	out, err := os.Create(filePath)
	if err != nil {
		Log(ERROR, "Erro ao criar arquivo de destino '%s': %v", filePath, err)
		return ""
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		Log(ERROR, "Erro ao copiar arquivo para '%s': %v", filePath, err)
		return ""
	}

	relativeURL := strings.Replace(filePath, destinationDir, "/uploads", 1)
	return relativeURL
}

func ValidateImageFile(file *multipart.FileHeader) bool {
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/gif":  true,
	}
	if _, ok := allowedTypes[file.Header.Get("Content-Type")]; !ok {
		return false
	}

	if file.Size > constants.MAX_UPLOAD_SIZE {
		return false
	}

	return true
}
