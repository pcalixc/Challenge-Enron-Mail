package utilsv2

import (
	"fmt"
	"indexer/models"
	"os"
	"path/filepath"
	"runtime"
)

type Result struct {
	filePath string
	email    models.EnronMail
	err      error
}

func ProcessEmailDirectory(pathName string) error {
	workerPool := NewWorkerPool(runtime.NumCPU(), 1000)
	fmt.Println(runtime.NumCPU())
	workerPool.Start()

	err := readFilesInDir(pathName, workerPool)
	if err != nil {
		// Manejar el error al leer los archivos
		fmt.Printf("Error al leer archivos en el directorio: %v\n", err)
	}

	workerPool.Close()
	return nil
}

// readFilesInDir lee los nombres de archivo en el directorio y sus subdirectorios y los envía al pool de trabajadores para su procesamiento.
func readFilesInDir(dir string, workerPool *WorkerPool) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		filePath := filepath.Join(dir, file.Name())
		if file.IsDir() {
			// Si es un directorio, llamar recursivamente a la función para leer los archivos en el subdirectorio
			if err := readFilesInDir(filePath, workerPool); err != nil {
				return err
			}
		} else {
			// Si es un archivo, enviarlo al pool de trabajadores para su procesamiento
			workerPool.SubmitFile(filePath)
		}
	}
	return nil
}
