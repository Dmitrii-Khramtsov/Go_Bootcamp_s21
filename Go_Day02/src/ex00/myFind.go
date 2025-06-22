package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func parseFlags() (bool, bool, bool, string, string) {
	var (
		showFiles    bool
		showDirs     bool
		showSymlinks bool
		ext          string
	)

	flag.BoolVar(&showFiles, "f", false, "Показать файлы")
	flag.BoolVar(&showDirs, "d", false, "Показать директории")
	flag.BoolVar(&showSymlinks, "sl", false, "Показать символические ссылки")
	flag.StringVar(&ext, "ext", "", "Расширение файла для фильтрации (работает ТОЛЬКО при указании -f)")

	flag.Parse()

	if flag.NArg() != 1 {
		log.Fatalf("Неправильное количество аргументов. Использование: %s [опции] <путь>", os.Args[0])
	}
	path := flag.Arg(0)

	return showFiles, showDirs, showSymlinks, ext, path
}

func findEntries(showFiles, showDirs, showSymlinks bool, ext, path string) {
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			if os.IsPermission(err) {
				return nil // пропускаем ошибки доступа
			}
			return err
		}
		switch {
		case info.IsDir() && showDirs:
			fmt.Println(path)
		case !info.IsDir() && showFiles:
			if ext == "" || filepath.Ext(path) == "."+ext {
				fmt.Println(path)
			}
		case info.Mode()&os.ModeSymlink != 0 && showSymlinks:
			handleSymlink(path)
		}
		return nil
	})
}

func handleSymlink(path string) {
	link, err := filepath.EvalSymlinks(path)
	if err != nil {
		fmt.Printf("%s -> [broken]\n", path)
		return
	}
	fmt.Printf("%s -> %s\n", path, link)
}

func main() {
	showFiles, showDirs, showSymlinks, ext, path := parseFlags()
	findEntries(showFiles, showDirs, showSymlinks, ext, path)
}
