package plugin

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/nomad-autoscaler/plugins"
	"github.com/hashicorp/nomad-autoscaler/plugins/apm"
	"github.com/hashicorp/nomad-autoscaler/plugins/base"
	"github.com/hashicorp/nomad-autoscaler/sdk"
)

const (
	// pluginName is the name of the plugin
	pluginName = "files-apm"
)

var (
	PluginID = plugins.PluginID{
		Name:       pluginName,
		PluginType: sdk.PluginTypeAPM,
	}

	PluginConfig = &plugins.InternalPluginConfig{
		Factory: func(l hclog.Logger) interface{} { return NewFilesAPMPlugin(l) },
	}

	pluginInfo = &base.PluginInfo{
		Name:       pluginName,
		PluginType: sdk.PluginTypeAPM,
	}
)

type APMPlugin struct {
	logger hclog.Logger
}

func NewFilesAPMPlugin(log hclog.Logger) apm.APM {
	return &APMPlugin{
		logger: log,
	}
}

func (a *APMPlugin) PluginInfo() (*base.PluginInfo, error) {
	return pluginInfo, nil
}

func (a *APMPlugin) SetConfig(config map[string]string) error {
	return nil
}

func (a *APMPlugin) Query(query string, r sdk.TimeRange) (sdk.TimestampedMetrics, error) {
	content, err := ioutil.ReadFile(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query: %v", err)
	}
	num, err := strconv.ParseFloat(string(content), 10)
	if err != nil {
		return nil, fmt.Errorf("failed to query: %v", err)
	}

	var result sdk.TimestampedMetrics

	// Generate one value per second.
	repeat := int(r.To.Sub(r.From).Seconds())

	for i := 1; i <= repeat; i++ {
		ts := r.From.Add(time.Duration(i) * time.Second).UTC()
		result = append(result, sdk.TimestampedMetric{Value: num, Timestamp: ts})
	}
	return result, nil
}

func (a *APMPlugin) QueryMultiple(query string, r sdk.TimeRange) ([]sdk.TimestampedMetrics, error) {
	m, err := a.Query(query, r)
	if err != nil {
		return nil, err
	}
	return []sdk.TimestampedMetrics{m}, nil
}
