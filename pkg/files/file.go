package files

import (
	"fmt"
	"log"
	"os"
	"time"
)

func RemoveAfter(fullPath string, duration time.Duration) {
	go func() {
		time.Sleep(duration)
		err := os.Remove(fullPath)
		if err != nil {
			log.Printf("Error removing file: %v", err)
		} else {
			log.Printf("File %s removed successfully", fullPath)
		}
	}()
}
func GenerateUniqueNameFile(ext string) string {
	currentDate := time.Now()
	path := fmt.Sprintf("%s/%s.%s", ext, currentDate.Format("20060102150405"), ext)
	return path
}
