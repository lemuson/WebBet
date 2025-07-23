package backend

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, `{"error": "Не удалось получить файл"}`, http.StatusBadRequest)
		return
	}
	defer file.Close()

	uploadDir := "main/frontend/uploads"
	os.MkdirAll(uploadDir, os.ModePerm)

	filename := header.Filename
	filepath := filepath.Join(uploadDir, filename)
	outFile, err := os.Create(filepath)
	if err != nil {
		http.Error(w, `{"error": "Ошибка при сохранении файла"}`, http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		http.Error(w, `{"error": "Ошибка при копировании файла"}`, http.StatusInternalServerError)
		return
	}

	relativePath := "/uploads/" + filename

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"path": "%s"}`, relativePath)))
}
