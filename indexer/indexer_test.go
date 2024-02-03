package main

import (
	"indexer/models"
	"indexer/utils"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestConvertEmailFileToJSON(t *testing.T) {
// 	// Crear un archivo de prueba con contenido de correo electrónico
// 	testEmailFile := "test_email.txt"
// 	defer func() {
// 		_ = os.Remove(testEmailFile)
// 	}()

// 	emailContent := `Message-ID: <13234904.1075861308849.JavaMail.evans@thyme>\n
// 	Date: Tue, 20 Nov 2001 12:07:00 -0800 (PST)
// 	From: alan.aronowitz@enron.com
// 	To: matthias.lee@enron.com
// 	Subject: Re: counterparty request to migrate trades
// 	Cc: elizabeth.sager@enron.com
// 	Mime-Version: 1.0
// 	Content-Type: text/plain; charset=us-ascii
// 	Content-Transfer-Encoding: 7bit
// 	Bcc: elizabeth.sager@enron.com
// 	X-From: Aronowitz, Alan </O=ENRON/OU=NA/CN=RECIPIENTS/CN=AARONOW>
// 	X-To: Lee, Matthias </O=ENRON/OU=NA/CN=RECIPIENTS/CN=EU/cn=Recipients/cn=mlee3>
// 	X-cc: Sager, Elizabeth </O=ENRON/OU=NA/CN=RECIPIENTS/CN=Esager>
// 	X-bcc:
// 	X-Folder: \ESAGER (Non-Privileged)\Sager, Elizabeth\Deleted Items
// 	X-Origin: Sager-E
// 	X-FileName: ESAGER (Non-Privileged).pst

// 	I have no objections, but I am copying Elizabeth Sager on this note to see if she has any issues. She is handling the Credit legal issues.
// 	 Alan Aronowitz`

// 	// Escribir contenido del correo electrónico en el archivo de prueba
// 	err := os.WriteFile(testEmailFile, []byte(emailContent), 0644)
// 	assert.NoError(t, err)

// 	// Llamar a la función bajo prueba
// 	result := utils.ConvertEmailFileToJSON(testEmailFile)

// 	// Verificar el resultado esperado
// 	expectedResult := models.EnronMail{
// 		MessageID:               "<13234904.1075861308849.JavaMail.evans@thyme>",
// 		Date:                    "Tue, 20 Nov 2001 12:07:00 -0800 (PST)",
// 		From:                    "alan.aronowitz@enron.com",
// 		To:                      "matthias.lee@enron.com",
// 		Subject:                 "Re: counterparty request to migrate trades",
// 		MimeVersion:             "1.0",
// 		ContentType:             "text/plain; charset=us-ascii",
// 		ContentTransferEncoding: "7bit",
// 		Content:                 "I have no objections, but I am copying Elizabeth Sager on this note to see if she has any issues. She is handling the Credit legal issues.\nAlan Aronowitz",
// 	}

// 	//assert.Equal(t, expectedResult.MessageID, result.MessageID)
// 	assert.Equal(t, expectedResult.Date, result.Date)
// }

func TestConvertEmailFileToJSONEmptyFile(t *testing.T) {
	// Crear un archivo de prueba vacío
	emptyFile := "empty_file.txt"
	_ = os.WriteFile(emptyFile, []byte(""), 0644)

	defer func() {
		_ = os.Remove(emptyFile)
	}()

	// Llamar a la función bajo prueba con el archivo vacío
	result := utils.ConvertEmailFileToJSON(emptyFile)

	// Verificar que el resultado tenga campos vacíos
	expectedResult := models.EnronMail{}
	assert.Equal(t, expectedResult, result)
}

func TestListFiles(t *testing.T) {
	// Create a new directory and some files
	testDir := "test_dir"
	testFiles := []string{"file1.txt", "file2.txt", "file3.txt"}

	err := os.Mkdir(testDir, 0755)
	assert.NoError(t, err)
	defer os.RemoveAll(testDir)

	for _, file := range testFiles {
		filePath := testDir + "/" + file
		err := os.WriteFile(filePath, []byte("content"), 0644)
		assert.NoError(t, err)
		defer os.Remove(filePath)
	}

	result, err := utils.ListFiles(testDir)

	// Verify
	assert.NoError(t, err)
	assert.ElementsMatch(t, testFiles, result)

}

func TestListFilesNonexistentDirectory(t *testing.T) {
	result, err := utils.ListFiles("nonexistent_directory")

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestMapEmailHeaders(t *testing.T) {
	// Crear una estructura EnronMail de prueba
	emailStruct := models.EnronMail{}

	// Llamar a la función bajo prueba para asignar valores a diferentes campos
	emailStruct = utils.MapEmailHeaders("Subject", "Test Subject", emailStruct)
	emailStruct = utils.MapEmailHeaders("From", "sender@example.com", emailStruct)
	emailStruct = utils.MapEmailHeaders("Date", "2022-01-01", emailStruct)

	// Verificar que los campos se asignen correctamente
	assert.Equal(t, "Test Subject", emailStruct.Subject)
	assert.Equal(t, "sender@example.com", emailStruct.From)
	assert.Equal(t, "2022-01-01", emailStruct.Date)

	// Verificar que otros campos no hayan sido modificados
	assert.Empty(t, emailStruct.MessageID)
	assert.Empty(t, emailStruct.To)
	// Añadir más verificaciones según sea necesario
}
