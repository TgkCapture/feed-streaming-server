package utils

import (
	"fmt"
	"path/filepath"
	"time"
)

func GenerateUniqueFilename(originalFilename string) string {
	fileExt := filepath.Ext(originalFilename)
	
	timestamp := time.Now().UnixNano()
	
	filename := fmt.Sprintf("%s_%d%s", originalFilename, timestamp, fileExt)
	return filename
}
