package iniutil

import (
	"fmt"
	"log"
	"strings"

	"gopkg.in/ini.v1"
)

var (
	cfg *ini.File
)

func init() {
	var err error
	cfg, err = ini.Load("config")
	if err != nil {
		log.Print("Config file not found")
	}
}

func LoadConfig() {
	var err error
	cfg, err = ini.Load("config")
	if err != nil {
		log.Print("Config file not found")
	}
}

func CreateEmptyIni() *ini.File {
	return ini.Empty()
}

func FindCategories(supplier string) []string {

	var secs []string
	for _, section := range cfg.Sections() {
		if strings.Contains(strings.ToLower(supplier), strings.ToLower(section.Name())) {
			if section.HasKey("kategooria") {
				key, err := section.GetKey("kategooria")
				if err != nil {
					log.Fatal(err)
				}
				key.Value()
				secs = append(secs, strings.Split(key.Value(), ",")...)
			}
		}
	}
	return secs
}

func SupplierInConfig(supplier string) bool {
	supplier = strings.ToLower(supplier)

	for _, section := range cfg.Sections() {
		name := strings.ToLower(section.Name())

		// skip DEFAULT section
		if name == "default" {
			continue
		}

		if strings.Contains(supplier, name) {
			return true
		}
	}
	return false
}

func AddConfigEntry(secName string, category string) {
	fmt.Println("ADDING CONFIG ENTRY", secName, category)
	section, err := cfg.NewSection(secName)
	if err != nil {
		panic(err)
	}
	section.NewKey("kategooria", category)

	cfg.SaveTo("config")
}
