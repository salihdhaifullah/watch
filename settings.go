package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Settings struct {
	Ignore             []string `json:"ignore"`
	IgnoreRegex        []string `json:"ignore_regex"`
	Delay              int      `json:"delay"`
	BackgroundCommands []string `json:"background_commands"`
	SetupCommands      []string `json:"setup_commands"`
	ParallelCommands   []string `json:"parallel_commands"`
}

func LoadSetting() (*Settings, error) {
	jsonSettings, err := os.ReadFile("./watch.config.json")
	if err != nil {
		fmt.Println("Error reading watch.config.json:", err)
		return nil, err
	}

	var settings Settings
	err = json.Unmarshal(jsonSettings, &settings)
	if err != nil {
		fmt.Println("Error parsing watch.config.json:", err)
		return nil, err
	}

	return &settings, nil
}
