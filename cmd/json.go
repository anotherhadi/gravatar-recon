package main

import (
	"encoding/json"
	"os"

	"github.com/charmbracelet/log"
)

func writeJson(outputPath string, data any) {
	if outputPath == "" {
		return
	}
	file, err := os.Create(outputPath)
	if err != nil {
		log.Error("Failed to create JSON file", "err", err)
		return
	}

	defer func() {
		_ = file.Close()
	}()

	as_json, _ := json.MarshalIndent(data, "", "\t")
	_, err = file.Write(as_json)
	if err != nil {
		log.Error("Failed to write to JSON file", "err", err)
		return
	}
	log.Info("JSON output written to file", "file", outputPath)
}
