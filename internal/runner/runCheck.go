package runner

import (
	"fmt"
	"os"
)

func CheckPath(path string) error {
	// Проверка существования пути
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("путь %s не существует", path)
		}
		return fmt.Errorf("ошибка при проверке пути: %v", err)
	}

	// Проверка, является ли путь директорией
	fileInfo, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("ошибка при получении информации о файле: %v", err)
	}

	if !fileInfo.IsDir() {
		return fmt.Errorf("путь %s не является директорией", path)
	}

	return nil
}
