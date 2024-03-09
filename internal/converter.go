package internal

import (
	"converter/internal/models"
	"converter/ioutil"
	"encoding/json"

	"gopkg.in/yaml.v2"
)

func ConvertJsonToYaml(filename string) ([]byte, error) {
	jsonData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var employee models.Employee
	err = json.Unmarshal(jsonData, &employee)
	if err != nil {
		return nil, err
	}

	err = employee.Validate()
	if err != nil {
		return nil, err
	}

	yamlData, err := yaml.Marshal(&employee)
	if err != nil {
		return nil, err
	}

	return yamlData, nil
}
