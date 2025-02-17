package main

import (
	"github.com/devstream-io/devstream/internal/pkg/plugin/harbordocker"
	"github.com/devstream-io/devstream/pkg/util/log"
)

// NAME is the name of this DevStream plugin.
const NAME = "harbor-docker"

// Plugin is the type used by DevStream core. It's a string.
type Plugin string

// Create implements the create of harbor-docker.
func (p Plugin) Create(options map[string]interface{}) (map[string]interface{}, error) {
	return harbordocker.Create(options)
}

// Update implements the update of harbor-docker.
func (p Plugin) Update(options map[string]interface{}) (map[string]interface{}, error) {
	return harbordocker.Update(options)
}

// Delete implements the delete of harbor-docker.
func (p Plugin) Delete(options map[string]interface{}) (bool, error) {
	return harbordocker.Delete(options)
}

// Read implements the read of harbor-docker.
func (p Plugin) Read(options map[string]interface{}) (map[string]interface{}, error) {
	return harbordocker.Read(options)
}

// DevStreamPlugin is the exported variable used by the DevStream core.
var DevStreamPlugin Plugin

func main() {
	log.Infof("%T: %s is a plugin for DevStream. Use it with DevStream.\n", DevStreamPlugin, NAME)
}
