package main

import (
	"archive/tar"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
)

// Получение информации о файле
func getFileInfo(logPath string) (os.FileInfo, error) {
	filePath, err := os.Stat(logPath)
	if err != nil {
		return nil, err
	}
	return filePath, nil
}

// Получение времени модификации лог-файла
func getModificationTime(fileInfo os.FileInfo) int64 {
	return fileInfo.ModTime().Unix()
}

func getFileName(logPath string) string {
	return filepath.Base(logPath)
}

// Создание имени архива
func createArchiveName(logPath string, mtime int64) string {
	baseName := getFileName(logPath)
	return fmt.Sprintf("%s_%d.tar.gz", baseName, mtime)
}

// / Проверка существования директории и её создание, если она не существует
func ensureDirectoryExists(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, 0755)
	}
	return nil
}

// Создание архивного файла и писателей
func createWriters(archivePath string) (*os.File, *gzip.Writer, *tar.Writer, error) {
	archiveFile, err := os.Create(archivePath)
	if err != nil {
		return nil, nil, nil, err
	}

	gzipWriter := gzip.NewWriter(archiveFile)
	tarWriter := tar.NewWriter(gzipWriter)

	return archiveFile, gzipWriter, tarWriter, err
}

// Закрытие всех писателей и файлов
func closeWriters(archiveFile *os.File, gzipWriter *gzip.Writer, tarWriter *tar.Writer) {
	tarWriter.Close()
	gzipWriter.Close()
	archiveFile.Close()
}

// Проверка существования лог-файла и директории архива
func checkAndPrepare(logPath, archiveDir string) (string, error) {
	fileInfo, err := getFileInfo(logPath)
	if err != nil {
		return "", fmt.Errorf("log file %s does not exist: %v", logPath, err)
	}

	if err := ensureDirectoryExists(archiveDir); err != nil {
		return "", fmt.Errorf("error creating archive directory %s: %v", archiveDir, err)
	}

	mtime := getModificationTime(fileInfo)
	archiveName := createArchiveName(logPath, mtime)
	archivePath := filepath.Join(archiveDir, archiveName)

	return archivePath, nil
}

// Добавление лог-файла в tar архив
func addLogFileToTar(logPath string, tarWriter *tar.Writer, fileInfo os.FileInfo) error {
	logFile, err := os.Open(logPath)
	if err != nil {
		return err
	}
	defer logFile.Close()

	baseName := getFileName(logPath)
	header := &tar.Header{
		Name: baseName,
		Mode: 0600,
		Size: fileInfo.Size(),
	}
	if err := tarWriter.WriteHeader(header); err != nil {
		return err
	}

	if _, err := io.Copy(tarWriter, logFile); err != nil {
		return err
	}

	return nil
}

// Основная функция для архивирования лог-файла
func archiveLog(logPath, archiveDir string, wg *sync.WaitGroup) {
	defer wg.Done()

	archivePath, err := checkAndPrepare(logPath, archiveDir)
	if err != nil {
		log.Printf("Error preparing archive for %s: %v", logPath, err)
		return
	}

	archiveFile, gzipWriter, tarWriter, err := createWriters(archivePath)
	if err != nil {
		log.Printf("Error creating archive file for %s: %v", logPath, err)
		return
	}
	defer closeWriters(archiveFile, gzipWriter, tarWriter)

	fileInfo, err := getFileInfo(logPath)
	if err != nil {
		log.Printf("Error getting file info for %s: %v", logPath, err)
		return
	}

	if err := addLogFileToTar(logPath, tarWriter, fileInfo); err != nil {
		log.Printf("Error adding log file to tar archive for %s: %v", logPath, err)
		return
	}

	log.Printf("Archived %s to %s", logPath, archivePath)
}

func main() {
	archiveDir := flag.String("a", ".", "Archive directory")
	flag.Parse()

	logPaths := flag.Args()
	if len(logPaths) == 0 {
		log.Fatal("No log files specified")
	}

	var wg sync.WaitGroup
	for _, logPath := range logPaths {
		wg.Add(1)
		go archiveLog(logPath, *archiveDir, &wg)
	}
	wg.Wait()
}
