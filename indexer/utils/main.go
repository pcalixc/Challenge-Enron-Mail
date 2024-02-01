package utils

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"indexer/models"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	_ "net/http/pprof"
)

// ListFolderFiles devuelve una lista de archivos en una carpeta.
func ListFiles(path string) ([]string, error) {
	var folderList []string

	outputDirRead, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open folder %s: %v", path, err)
	}
	defer outputDirRead.Close()

	outputDirFiles, err := outputDirRead.Readdir(0)
	if err != nil {
		return nil, fmt.Errorf("could not read files from folder %s: %v", path, err)
	}

	for outputIndex := range outputDirFiles {
		outputFileHere := outputDirFiles[outputIndex]

		outputNameHere := outputFileHere.Name()
		folderList = append(folderList, outputNameHere)
	}

	return folderList, nil
}

// IsDirectory verifica si una ruta corresponde a un directorio.
func IsDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
	}
	if fileInfo.IsDir() {
		return true
	} else {
		return false
	}
}

// MapEmailHeaders estructura los datos del correo según los encabezados.
func MapEmailHeaders(key string, value string, emailStruct models.EnronMail) models.EnronMail {
	switch key {
	case "Message-ID":
		emailStruct.MessageID = value
	case "Date":
		emailStruct.Date = value
	case "From":
		emailStruct.From = value
	case "To":
		emailStruct.To = value
	case "Subject":
		emailStruct.Subject = value
	case "Mime-Version":
		emailStruct.MimeVersion = value
	case "Content-Type":
		emailStruct.ContentType = value
	case "Content-Transfer-Encoding":
		emailStruct.ContentTransferEncoding = value
	case "X-From":
		emailStruct.XFrom = value
	case "X-To":
		emailStruct.XTo = value
	case "X-cc":
		emailStruct.XCc = value
	case "X-bcc":
		emailStruct.XBcc = value
	case "X-Folder":
		emailStruct.XFolder = value
	case "X-Origin":
		emailStruct.XOrigin = value
	case "X-FileName":
		emailStruct.XFileName = value
	}

	return emailStruct
}

// ConvertEmailFileToJSON convierte un archivo de correo a formato JSON.
func ConvertEmailFileToJSON(filePath string) models.EnronMail {
	var bodyLines strings.Builder
	var emailStructure models.EnronMail
	var bodyStarted bool

	// We read the email file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// We read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// If the line is empty, it indicates the start of the email body
		if line == "" {
			bodyStarted = true
			continue
		}

		if bodyStarted {
			// Store body lines in a slice
			bodyLines.WriteString(line)
			bodyLines.WriteString("\n")
		} else {
			// Parse email headers
			parts := strings.SplitN(line, ": ", 2)
			if len(parts) == 2 {
				key := parts[0]
				value := parts[1]
				// Create an Email object
				emailStructure = MapEmailHeaders(key, value, emailStructure)
			}
		}
	}

	// Add the body to the email structure
	emailStructure.Content = bodyLines.String()

	return emailStructure
}

// SendDataToIndex envía datos al índice utilizando HTTP.
func SendDataToIndex(data []models.EnronMail) error {
	newData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	reader := bytes.NewReader(newData)

	req, err := http.NewRequest("POST", "http://localhost:5080/api/default/test3/_json", reader)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("root@example.com", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Body: " + string(body))

	return nil
}

func IndexEmailFolder(path string) error {
	id := 0
	batchSize := 100
	var dataBatch []models.EnronMail

	var indexEmail func(string) error
	indexEmail = func(path string) error {
		if IsDirectory(path) {
			folders, err := ListFiles(path)
			if err != nil {
				return fmt.Errorf("error while listing files in path %s: %v", path, err)
			}
			for _, folder := range folders {
				err := indexEmail(path + "/" + folder)
				if err != nil {
					return fmt.Errorf("error while indexing email in path %s: %v", path, err)
				}
			}
		} else {
			jsonEmail := ConvertEmailFileToJSON(path)
			dataBatch = append(dataBatch, jsonEmail)

			if len(dataBatch) == batchSize {
				err := SendDataToIndex(dataBatch)
				if err != nil {
					return fmt.Errorf("error while %s: %v", path, err)
				}
				dataBatch = nil // Limpiar el lote después de enviar
			}
			id++
		}
		return nil
	}

	// Llamar a la función interna para iniciar el proceso de indexación
	if err := indexEmail(path); err != nil {
		return err
	}

	// Verificar si hay algún correo en el lote pendiente de enviar
	if len(dataBatch) > 0 {
		err := SendDataToIndex(dataBatch)
		if err != nil {
			return err
		}
	}

	return nil
}
