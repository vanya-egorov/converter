// main.go

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"project_folder/internal"
)

func main() {
	// Использование флага для указания имени файла
	filename := flag.String("file", "input.json", "Имя входного файла JSON")
	flag.Parse()

	// Вызов экспортируемой функции для конвертации
	yamlData, err := internal.ConvertJsonToYaml(*filename)
	if err != nil {
		fmt.Println("Ошибка конвертации:", err)
		return
	}

	// Запись полученных байтов в файл output.yaml
	err = ioutil.WriteFile("output.yaml", yamlData, 0644)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		return
	}

	fmt.Println("Конвертация выполнена успешно! Результат записан в output.yaml")
}
