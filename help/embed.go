package help

import (
	"embed"
	"path"
	"strings"
)

//go:embed *.txt
var content embed.FS

var files, _ = content.ReadDir(".")

var Help map[string]string = make(map[string]string)

func init() {
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		name := file.Name()
		name = strings.TrimSuffix(name, path.Ext(name))

		rawContents, _ := content.ReadFile(file.Name())
		Help[name] = string(rawContents)
	}
}
