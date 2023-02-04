package filestore

import (
	"encoding/json"
	"io"
	"os"

	"github.com/goccy/go-yaml"
)

func writeJSONToFile(file string, dataInterface interface{}) (err error) {
	jsonData, err := json.MarshalIndent(dataInterface, "", "  ")
	if err != nil {
		return
	}

	// create a file with a supplied name
	if jsonFile, err := os.Create(file); err != nil {
		return err
	} else if _, err = jsonFile.Write(jsonData); err != nil {
		return err
	}

	return nil
}

func readJSONFromFile(file string, dataInterface interface{}) (err error) {
	jsonFile, err := os.Open(file)
	// if we os.Open returns an error then handle it
	if err != nil {
		return err
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer func() {
		if err = jsonFile.Close(); err != nil {
			return
		}
	}()

	// read our opened xmlFile as a byte array.
	byteValue, _ := io.ReadAll(jsonFile)
	if err = json.Unmarshal(byteValue, dataInterface); err != nil {
		return err
	}

	// return the json byte value.
	return nil
}

func writeYamlToFile(file string, dataInterface interface{}) (err error) {
	yamlData, err := yaml.Marshal(&dataInterface)
	if err != nil {
		return
	}

	// create a file with a supplied name
	yamlFile, err := os.Create(file)
	if err != nil {
		return
	}

	if _, err = yamlFile.Write(yamlData); err != nil {
		return
	}

	return
}

func readYamlFromFile(file string, dataInterface interface{}) (err error) {
	yamlFile, err := os.Open(file)
	if err != nil {
		return
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer func() {
		if err := yamlFile.Close(); err != nil {
			return
		}
	}()

	byteValue, _ := io.ReadAll(yamlFile)
	if err = yaml.Unmarshal(byteValue, dataInterface); err != nil {
		return
	}

	return
}
