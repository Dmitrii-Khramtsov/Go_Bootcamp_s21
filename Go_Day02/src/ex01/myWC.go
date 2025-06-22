package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
)

func parseFlags() (bool, bool, bool, []string) {
	var (
		countLines bool
		countChars bool
		countWords bool
	)
	flag.BoolVar(&countLines, "l", false, "подсчет строк")
	flag.BoolVar(&countChars, "m", false, "подсчет символов")
	flag.BoolVar(&countWords, "w", false, "подсчет слов")

	flag.Parse()

	files := flag.Args()
	if len(files) == 0 {
		log.Fatalf("Не указаны файлы для обработки. Использование: %s [опции] <файлы>", os.Args[0])
	}

	// Проверка, что установлен только один флаг
	if countLines && countChars || countLines && countWords || countChars && countWords {
		log.Fatalf("Нельзя установить более одного флага. Использование: %s [опции] <файлы>", os.Args[0])
	}

	return countLines, countChars, countWords, files
}

// FileProcessor интерфейс для обработки файлов
type FileProcessor interface {
	Process(scanner *bufio.Scanner) (int, error)
}

func handleCount(file string) (*bufio.Scanner, *os.File, error) {
	data, err := os.Open(file)
	if err != nil {
		return nil, nil, err
	}

	scanner := bufio.NewScanner(data)
	return scanner, data, nil
}

type LineCounter struct{}

func (lc LineCounter) Process(scanner *bufio.Scanner) (int, error) {
	lines := 0
	for scanner.Scan() {
		lines++
	}
	return lines, scanner.Err()
}

type CharCounter struct{}

func (cc CharCounter) Process(scanner *bufio.Scanner) (int, error) {
	scanner.Split(bufio.ScanRunes)
	chars := 0
	for scanner.Scan() {
		chars++
	}
	return chars, scanner.Err()
}

type WordCounter struct{}

func (wc WordCounter) Process(scanner *bufio.Scanner) (int, error) {
	scanner.Split(bufio.ScanWords)
	words := 0
	for scanner.Scan() {
		words++
	}
	return words, scanner.Err()
}

// processFile обрабатывает файл в зависимости от режима подсчета.
func processFile(processor FileProcessor, file string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	scanner, data, err := handleCount(file)
	if err != nil {
		results <- fmt.Sprintf("Ошибка при открытии файла %s: %v", file, err)
		return
	}
	defer data.Close()
	count, err := processor.Process(scanner)
	if err != nil {
		results <- fmt.Sprintf("Ошибка при обработке файла %s: %v", file, err)
	} else {
		results <- fmt.Sprintf("%d\t%s", count, file)
	}
}

// getProcessor возвращает соответствующий процессор на основе установленных флагов
func getProcessor(countLine, countChar, countWord bool) FileProcessor {
	switch {
	case countLine:
		return LineCounter{}
	case countChar:
		return CharCounter{}
	case countWord:
		return WordCounter{}
	default:
		log.Fatalf("Не указан режим подсчета. Использование: %s [опции] <файлы>", os.Args[0])
		return nil // Этот код никогда не будет достигнут, но компилятор требует возвращения значения
	}
}

// processFiles обрабатывает файлы с использованием указанного процессора
func processFiles(processor FileProcessor, files []string) {
	var wg sync.WaitGroup
	results := make(chan string, len(files))

	for _, file := range files {
		wg.Add(1)
		go processFile(processor, file, &wg, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}
}

func main() {
	countLine, countChar, countWord, files := parseFlags()
	processor := getProcessor(countLine, countChar, countWord)
	processFiles(processor, files)
}
