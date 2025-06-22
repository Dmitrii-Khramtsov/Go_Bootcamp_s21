package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Использование: ./myXargs <команда> [аргументы]")
	}

	// получаем команду и аргументы из аргументов командной строки
	command := os.Args[1]
	args := os.Args[2:]

	// читаем строки из стандартного ввода
	inputLines, err := readStdin()
	if err != nil {
		log.Fatalf("Ошибка при чтении стандартного ввода: %v", err)
	}

	// добавляем строки из стандартного ввода к аргументам команды
	args = append(args, inputLines...)

	// создаем и выполняем команду
	if err := executeCommand(command, args); err != nil {
		log.Fatalf("Ошибка при выполнении команды: %v", err)
	}
}

// читает строки из стандартного ввода и возвращает их в виде слайса строк
func readStdin() ([]string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var inputLines []string
	for scanner.Scan() {
		inputLines = append(inputLines, scanner.Text())
	}
	return inputLines, scanner.Err()
}

// создает и выполняет команду с переданными аргументами
func executeCommand(command string, args []string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
