package main

import (
	"encoding/json"
	"io/ioutil"
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
		return nil
	}

	return ioutil.WriteFile(filename, jsonData, 0644)
}

func ReadAirplanesFromFile(filename string) ([]airplane, error) {
	data, err := ioutil.ReadFile(filename)
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
