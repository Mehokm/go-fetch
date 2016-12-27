package fetch

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type FileAdapter struct {
	filepath string
	services []Service
}

func NewFileAdapter(filepath string) (FileAdapter, error) {
	file, err := os.Open(filepath)

	if err != nil {
		return FileAdapter{}, err
	}

	data, err := ioutil.ReadAll(file)

	if err != nil {
		return FileAdapter{}, err
	}

	var svcs []Service

	err = unmarshal(data, &svcs)

	return FileAdapter{filepath, svcs}, err
}

func (fa FileAdapter) GetService(s string) (Service, error) {
	for _, v := range fa.services {
		if strings.EqualFold(s, v.Name) {
			return v, nil
		}
	}

	return Service{}, fmt.Errorf("fileadapter: cannot find service by name '%s'", s)
}

func unmarshal(data []byte, s *[]Service) error {
	ss := struct {
		Services *[]Service
	}{s}

	if json.Unmarshal(data, &ss) != nil {
		if yaml.Unmarshal(data, &ss) != nil {
			return errors.New("fileadapter: cannot unmarshal file.  File must be type 'json' or 'yaml'")
		}
	}

	return nil
}
