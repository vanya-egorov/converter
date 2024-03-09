// internal/converter.go

package internal

import (
	"encoding/json"
	"io/ioutil"
	"project_folder/models"

	"gopkg.in/yaml.v2"
)

func ConvertJsonToYaml(filename string) ([]byte, error) {
	// Чтение данных из файла JSON
	jsonData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Анмаршалинг данных в структуру Employee
	var employee models.Employee
	err = json.Unmarshal(jsonData, &employee)
	if err != nil {
		return nil, err
	}

	// Проверка валидности полей
	err = employee.Validate()
	if err != nil {
		return nil, err
	}

	// Маршалинг Employee в формат YAML
	yamlData, err := yaml.Marshal(&employee)
	if err != nil {
		return nil, err
	}

	return yamlData, nil
}
