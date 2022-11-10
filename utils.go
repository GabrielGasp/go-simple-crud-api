package main

import (
	"encoding/json"
	"os"
)

func FindAirplaneById(airplanes []airplane, id string) (int, airplane, bool) {
	for i, a := range airplanes {
		if a.ID == id {
			return i, a, true
		}
	}

	return 0, airplane{}, false
}

func WriteAirplanesToFile(filename string, data []airplane) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, jsonData, 0644)
}

func ReadAirplanesFromFile(filename string) ([]airplane, error) {
	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return []airplane{}, nil
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var airplanes []airplane

	err = json.Unmarshal(data, &airplanes)
	if err != nil {
		return nil, err
	}

	return airplanes, nil
}
