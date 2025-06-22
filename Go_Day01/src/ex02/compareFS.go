package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Структура для хранения ссылок
type Links struct {
	links []string
}

// Функция для открытия файла
func openFile(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// Функция для чтения файла и заполнения map
func readFileIntoMap(file *os.File) (map[string]struct{}, error) {
	fileMap := make(map[string]struct{})
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileMap[scanner.Text()] = struct{}{}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("cannot scanning a file: %w", err)
	}

	return fileMap, nil
}

// Функция для парсинга флагов командной строки для compareFS
func parseCompareFlags() (string, string, error) {
	oldFlag := flag.String("old", "", "path to the original database file")
	newFlag := flag.String("new", "", "path to the new database file")

	flag.Parse()

	if *oldFlag == "" || *newFlag == "" {
		return "", "", fmt.Errorf("both --old and --new flags are required")
	}

	return *oldFlag, *newFlag, nil
}

// Функция для проверки расширения файла
func checkFileExtension(filename string) error {
	ext := filepath.Ext(filename)
	if ext != ".txt" {
		return fmt.Errorf("invalid file extension: %s, expected .txt", ext)
	}
	return nil
}

// Функция для проверки существования файла
func checkFileExists(filename string) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %s", filename)
	}
	return nil
}

func main() {
	// Парсим флаги командной строки
	oldFilePath, newFilePath, err := parseCompareFlags()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Проверяем расширение файлов
	if err := checkFileExtension(oldFilePath); err != nil {
		log.Fatalf("error: %v", err)
	}
	if err := checkFileExtension(newFilePath); err != nil {
		log.Fatalf("error: %v", err)
	}

	// Проверяем существование файлов
	if err := checkFileExists(oldFilePath); err != nil {
		log.Fatalf("error: %v", err)
	}
	if err := checkFileExists(newFilePath); err != nil {
		log.Fatalf("error: %v", err)
	}

	// Открываем старый файл
	oldFile, err := openFile(oldFilePath)
	if err != nil {
		log.Fatalf("error opening old file: %v", err)
	}
	defer oldFile.Close()

	// Читаем старый файл и заполняем map
	oldFileMap, err := readFileIntoMap(oldFile)
	if err != nil {
		log.Fatalf("error reading old file: %v", err)
	}

	// Открываем новый файл
	newFile, err := openFile(newFilePath)
	if err != nil {
		log.Fatalf("error opening new file: %v", err)
	}
	defer newFile.Close()

	// Читаем новый файл и сравниваем его содержимое с содержимым старого файла
	scanner := bufio.NewScanner(newFile)
	for scanner.Scan() {
		line := scanner.Text()
		if _, exists := oldFileMap[line]; exists {
			delete(oldFileMap, line) // Удаляем элемент из map, если он существует
		} else {
			fmt.Printf("ADDED %s\n", line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error scanning new file: %v", err)
	}

	// Выводим удаленные элементы
	for line := range oldFileMap {
		fmt.Printf("REMOVED %s\n", line)
	}
}
