package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// Общий тип данных для возврата
type Cake struct {
	Name        string       `json:"name" xml:"name"`
	Time        string       `json:"time" xml:"stovetime"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>item"`
}

type Ingredient struct {
	Name  string `json:"ingredient_name" xml:"itemname"`
	Count string `json:"ingredient_count" xml:"itemcount"`
	Unit  string `json:"ingredient_unit,omitempty" xml:"itemunit,omitempty"`
}

type Data struct {
	Cakes []Cake `json:"cake" xml:"cake"`
}

// Интерфейс DBReader для чтения данных из файлов
type DBReader interface {
	Read(filePath string) (Data, error)
}

// Реализация DBReader для JSON файлов
type JSONReader struct{}

func (j JSONReader) Read(filePath string) (Data, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return Data{}, fmt.Errorf("cannot read file: %w", err)
	}

	var jsonData Data
	err = json.Unmarshal(fileContent, &jsonData)
	if err != nil {
		return Data{}, fmt.Errorf("cannot unmarshal JSON: %w", err)
	}

	return jsonData, nil
}

// Реализация DBReader для XML файлов
type XMLReader struct{}

func (x XMLReader) Read(filePath string) (Data, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return Data{}, fmt.Errorf("cannot read file: %w", err)
	}

	var xmlData Data
	err = xml.Unmarshal(fileContent, &xmlData)
	if err != nil {
		return Data{}, fmt.Errorf("cannot unmarshal XML: %w", err)
	}

	return xmlData, nil
}

// Функция для выбора подходящего ридера на основе расширения файла
func getReader(filePath string) (DBReader, error) {
	switch getFileFormat(filePath) {
	case "json":
		return JSONReader{}, nil
	case "xml":
		return XMLReader{}, nil
	default:
		return nil, fmt.Errorf("unsupported file format")
	}
}

// Функция для определения формата файла (json или xml)
func getFileFormat(filePath string) string {
	if strings.HasSuffix(filePath, ".json") {
		return "json"
	} else if strings.HasSuffix(filePath, ".xml") {
		return "xml"
	}
	return ""
}

// Функция для маршалинга данных в JSON или XML
func marshalData(data Data, fileFormat string) ([]byte, error) {
	if fileFormat == "json" {
		return json.MarshalIndent(data, "", "    ")
	} else if fileFormat == "xml" {
		return xml.MarshalIndent(data, "", "    ")
	}
	return nil, fmt.Errorf("unsupported file format")
}

// Функция для парсинга флагов командной строки
func parseFlags() (string, error) {
	// Определяем флаг -f для указания пути к файлу
	fileFlag := flag.String("f", "", "path to JSON or XML file")

	// Парсим флаги
	flag.Parse()

	// Проверяем, был ли указан флаг -f
	if *fileFlag == "" {
		return "", fmt.Errorf("the -f flag is required")
	}

	return *fileFlag, nil
}

func main() {
	// Парсим флаги командной строки и получаем путь к файлу
	filePath, err := parseFlags()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Получаем подходящий ридер на основе расширения файла
	reader, err := getReader(filePath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Читаем данные из файла
	data, err := reader.Read(filePath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Определяем формат файла
	fileFormat := getFileFormat(filePath)

	// Преобразуем данные в JSON или XML и выводим
	res, err := marshalData(data, fileFormat)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if fileFormat == "json" {
		fmt.Printf("JSON data:\n%s\n", res)
	} else if fileFormat == "xml" {
		fmt.Printf("XML data:\n%s\n", res)
	}
}
