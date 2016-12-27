package fetch

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"

	"gopkg.in/yaml.v2"
)

var fileAdapter FileAdapter
var mutex sync.Mutex

type FileAdapter struct {
	filepath string
	services []Service
	watcher  *fsnotify.Watcher
}

func NewFileAdapter(filepath string) (FileAdapter, error) {
	err := loadData(filepath)

	return fileAdapter, err
}

func NewFileWatcherAdapter(filepath string) (FileAdapter, error) {
	err := loadData(filepath)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fileAdapter, err
	}

	err = watcher.Add(filepath)
	if err != nil {
		return fileAdapter, err
	}

	fileAdapter.watcher = watcher

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

func (fa FileAdapter) StartWatcher() {
	if fa.watcher != nil {
		for {
			select {
			case event := <-fa.watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					mutex.Lock()
					loadData(fa.filepath)
					mutex.Unlock()
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

func loadData(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	var svcs []Service

	err = unmarshal(data, &svcs)

	fileAdapter = FileAdapter{filepath, svcs, nil}

	return err
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
