package fetch

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"

	"gopkg.in/yaml.v2"
)

var fileAdapter FileAdapter

type FileAdapter struct {
	filepath string
	services []Service
	watcher  *fsnotify.Watcher
}

func NewFileAdapter(filepath string) (FileAdapter, error) {
	file, err := os.Open(filepath)

	if err != nil {
		return fileAdapter, err
	}

	data, err := ioutil.ReadAll(file)

	if err != nil {
		return fileAdapter, err
	}

	var svcs []Service

	err = unmarshal(data, &svcs)

	fileAdapter = FileAdapter{filepath, svcs, nil}

	return fileAdapter, err
}

func NewFileWatcherAdapter(filepath string) (FileAdapter, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fileAdapter, err
	}

	err = watcher.Add(filepath)
	if err != nil {
		return fileAdapter, err
	}

	file, err := os.Open(filepath)

	if err != nil {
		return fileAdapter, err
	}

	data, err := ioutil.ReadAll(file)

	if err != nil {
		return fileAdapter, err
	}

	var svcs []Service

	err = unmarshal(data, &svcs)

	fileAdapter = FileAdapter{filepath, svcs, watcher}

	return fileAdapter, err
}

func (fa FileAdapter) GetService(s string) (Service, error) {
	for _, v := range fileAdapter.services {
		if strings.EqualFold(s, v.Name) {
			return v, nil
		}
	}

	return Service{}, fmt.Errorf("fileadapter: cannot find service by name '%s'", s)
}

func (fa FileAdapter) Watch() {
	if fa.watcher != nil {
		for {
			select {
			case event := <-fa.watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err := <-fa.watcher.Errors:
				if err != nil {
					log.Println("error:", err)
				}
			}
		}
	}
}

func (fa FileAdapter) CloseWatcher() {
	fa.watcher.Close()
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
