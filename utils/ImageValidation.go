package utils

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func IsValidImage(file *gin.Context, formKey string) bool {
    uploadedFile, err := file.FormFile(formKey)
    if err != nil {
        return false
    }

    extension := filepath.Ext(uploadedFile.Filename)
    return extension == ".png" || extension == ".jpeg" || extension == ".jpg"
}