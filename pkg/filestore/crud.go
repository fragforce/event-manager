package filestore

func CreateConfig(path string, data interface{}) (err error) {
	switch path[4:] {
	case "json":
		err = writeJSONToFile(path, data)
	case ".yml", "yaml":
		err = writeYamlToFile(path, data)
	default:
		err = writeYamlToFile(path, data)
	}

	return
}

func ReadConfig(path string, data interface{}) (err error) {
	switch path[4:] {
	case "json":
		err = readJSONFromFile(path, data)
	case ".yml", "yaml":
		err = readYamlFromFile(path, data)
	default:
		err = readYamlFromFile(path, data)
	}

	return
}

func UpdateConfig(path string, data interface{}) (err error) {
	switch path[4:] {
	case "json":
		err = writeJSONToFile(path, data)
	case ".yml", "yaml":
		err = writeYamlToFile(path, data)
	default:
		err = writeYamlToFile(path, data)
	}

	return
}

func DeleteFile(path string) (err error) {
	return
}
