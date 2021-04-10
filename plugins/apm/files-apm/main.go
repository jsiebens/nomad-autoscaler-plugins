package main

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/nomad-autoscaler/plugins"
	"github.com/jsiebens/nomad-autoscaler-plugins/plugins/apm/files-apm/plugin"
)

func main() {
	plugins.Serve(factory)
}

// factory returns a new instance of the Files APM plugin.
func factory(log hclog.Logger) interface{} {
	return plugin.NewFilesAPMPlugin(log)
}
