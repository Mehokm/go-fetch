package fetch

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFileAdapter_CanReadJson(t *testing.T) {
	fa, err := NewFileAdapter("test_data/file.json")

	if err != nil {
		t.Error(err)
	}

	data := []Service{
		Service{
			Name: "svc1json",
			Host: "localhost",
			Port: "8080",
		},
		Service{
			Name: "svc2json",
			Host: "localhost",
			Port: "9000",
		},
	}

	assert.Equal(t, data, fa.services)
}

func TestFileAdapter_CanReadYaml(t *testing.T) {
	fa, err := NewFileAdapter("test_data/file.yaml")

	if err != nil {
		t.Error(err)
	}

	data := []Service{
		Service{
			Name: "svc1yaml",
			Host: "localhost",
			Port: "8080",
		},
		Service{
			Name: "svc2yaml",
			Host: "localhost",
			Port: "9000",
		},
	}

	assert.Equal(t, data, fa.services)
}

func TestFileAdapter_ReturnsErrorInvalidFormat(t *testing.T) {
	_, err := NewFileAdapter("test_data/file.txt")

	assert.EqualError(t, err, "fileadapter: cannot unmarshal file.  File must be type 'json' or 'yaml'")
}

func TestFileAdapter_ReturnsErrorServiceNotFound(t *testing.T) {
	fa, _ := NewFileAdapter("test_data/file.yaml")

	_, err := fa.GetService("not_there")

	assert.EqualError(t, err, "fileadapter: cannot find service by name 'not_there'")
}

func TestFileAdapter_ReturnsCorrectService(t *testing.T) {
	fa, _ := NewFileAdapter("test_data/file.yaml")

	svc, _ := fa.GetService("svc1yaml")

	data := Service{
		Name: "svc1yaml",
		Host: "localhost",
		Port: "8080",
	}

	assert.Equal(t, svc, data)
}

func TestFileAdapter_FileWatcher(t *testing.T) {
	// create tmp file
	filepath := fmt.Sprintf("test_data/tmp%v", time.Now().Unix())

	file, err := os.Create(filepath)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	}

	fa, _ := NewFileWatcherAdapter(filepath)

	// start fsnotify
	go fa.StartWatcher()
	defer fa.CloseWatcher()

	// test svc not there
	_, err = fa.GetService("svc1watch")

	assert.EqualError(t, err, "fileadapter: cannot find service by name 'svc1watch'")

	file.Write([]byte(
		`{
			"services": [{
				"name": "svc1watch",
				"host": "localhost",
				"port": "8181"
			}]
		}
	`))

	file.Sync()

	time.Sleep(time.Second * 1)

	data := Service{
		Name: "svc1watch",
		Host: "localhost",
		Port: "8181",
	}

	// test file adapter updated the serices
	svc, err := fa.GetService("svc1watch")

	assert.NoError(t, err)
	assert.Equal(t, data, svc)

	os.Remove(filepath)
}
