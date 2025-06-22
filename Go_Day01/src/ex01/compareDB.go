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

// Функция для парсинга флагов командной строки для compareDB
func parseCompareFlags() (string, string, error) {
	oldFlag := flag.String("old", "", "path to the original database file")
	newFlag := flag.String("new", "", "path to the new database file")

	flag.Parse()

	if *oldFlag == "" || *newFlag == "" {
		return "", "", fmt.Errorf("both --old and --new flags are required")
	}

	return *oldFlag, *newFlag, nil
}

// Функция для создания карты тортов
func createCakeMap(cakes []Cake) map[string]Cake {
	cakeMap := make(map[string]Cake)
	for _, cake := range cakes {
		cakeMap[cake.Name] = cake
	}
	return cakeMap
}

// Функция для сравнения тортов
func compareCakes(oldCakesMap, newCakesMap map[string]Cake) {
	for name, oldCake := range oldCakesMap {
		newCake, exists := newCakesMap[name]
		if !exists {
			fmt.Printf("REMOVED cake \"%s\"\n", name)
			continue
		}

		if oldCake.Time != newCake.Time {
			fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", name, newCake.Time, oldCake.Time)
		}

		compareIngredients(oldCake, newCake, name)
	}

	for name := range newCakesMap {
		if _, exists := oldCakesMap[name]; !exists {
			fmt.Printf("ADDED cake \"%s\"\n", name)
		}
	}
}

// Функция для создания карты ингредиентов
func createIngredientMap(ingredients []Ingredient) map[string]Ingredient {
	ingredientMap := make(map[string]Ingredient)
	for _, ingredient := range ingredients {
		ingredientMap[ingredient.Name] = ingredient
	}
	return ingredientMap
}

// Функция для сравнения ингредиентов
func compareIngredients(oldCake, newCake Cake, cakeName string) {
	oldIngredientsMap := createIngredientMap(oldCake.Ingredients)
	newIngredientsMap := createIngredientMap(newCake.Ingredients)

	for ingredientName, oldIngredient := range oldIngredientsMap {
		newIngredient, exists := newIngredientsMap[ingredientName]
		if !exists {
			fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n", ingredientName, cakeName)
			continue
		}

		if oldIngredient.Count != newIngredient.Count {
			fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", ingredientName, cakeName, newIngredient.Count, oldIngredient.Count)
		}

		compareUnits(oldIngredient, newIngredient, ingredientName, cakeName)
	}

	for ingredientName := range newIngredientsMap {
		if _, exists := oldIngredientsMap[ingredientName]; !exists {
			fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n", ingredientName, cakeName)
		}
	}
}

// Функция для сравнения единиц измерения
func compareUnits(oldIngredient, newIngredient Ingredient, ingredientName, cakeName string) {
	if oldIngredient.Unit != newIngredient.Unit {
		if oldIngredient.Unit == "" {
			fmt.Printf("ADDED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n", newIngredient.Unit, ingredientName, cakeName)
		} else if newIngredient.Unit == "" {
			fmt.Printf("REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n", oldIngredient.Unit, ingredientName, cakeName)
		} else {
			fmt.Printf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", ingredientName, cakeName, newIngredient.Unit, oldIngredient.Unit)
		}
	}
}

// Основная функция для сравнения баз данных
func compareDatabases(oldData, newData Data) {
	// Создаем карты тортов для старой и новой базы данных
	oldCakesMap := createCakeMap(oldData.Cakes)
	newCakesMap := createCakeMap(newData.Cakes)

	// Сравниваем торты
	compareCakes(oldCakesMap, newCakesMap)
}

func main() {
	// Парсим флаги командной строки
	oldFilePath, newFilePath, err := parseCompareFlags()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Получаем ридеры для старой и новой базы данных
	oldReader, err := getReader(oldFilePath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	newReader, err := getReader(newFilePath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Читаем данные из файлов
	oldData, err := oldReader.Read(oldFilePath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	newData, err := newReader.Read(newFilePath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Сравниваем базы данных
	compareDatabases(oldData, newData)
}
