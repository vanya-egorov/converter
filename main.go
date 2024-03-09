// main.go

package main

import (
	"converter/internal"
	"converter/ioutil"
	"flag"
	"fmt"
)

func main() {
	filename := flag.String("file", "input.json", "Имя входного файла JSON")
	flag.Parse()

	yamlData, err := internal.ConvertJsonToYaml(*filename)
	if err != nil {
		fmt.Println("Ошибка конвертации:", err)
		return
	}

	err = ioutil.WriteFile("output.yaml", yamlData, 0644)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		return
	}

	fmt.Println("Конвертация выполнена успешно! Результат записан в output.yaml")
}
