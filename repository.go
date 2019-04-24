package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type RegisterdUrl struct {
	path string
	url  string
}

type MappingValues map[string]string

func Init(filePath string) *MappingValues {
	mappingValues := make(MappingValues)

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Print(err)
		return &mappingValues
	}

	if err := yaml.Unmarshal(data, &mappingValues); err != nil {
		log.Fatalf("error: %v", err)
	}

	return &mappingValues
}

func (m *MappingValues) Save(filePath string) {
	if m == nil || len(*m) == 0 {
		return
	}

	data, err := yaml.Marshal(m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if err := ioutil.WriteFile(filePath, data, 0644); err != nil {
		log.Fatal(err)
	}
}

func (m *MappingValues) Add(r *RegisterdUrl) {
	(*m)[r.path] = r.url
}

func (m *MappingValues) CleanUp(path string) {
	if _, ok := (*m)[path]; ok {
		delete(*m, path)
	}
}
