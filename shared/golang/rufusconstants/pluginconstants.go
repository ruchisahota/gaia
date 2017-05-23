package rufusconstants

import (
	"fmt"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/gaia/rufusmodels/current/golang"
)

// PluginInfo carries information about a plugin
type PluginInfo struct {
	Name        string
	Description string
	Version     float32
}

func (p *PluginInfo) String() string {
	return fmt.Sprintf("<Plugin name=%s version=%f description=%s>", p.Name, p.Version, p.Description)
}

// PluginInfoFunc is the method that describe the plugin
type PluginInfoFunc func() *PluginInfo

// PluginExecuteFunc is the executed method of the plugin
// It returns an OutputData and an error.
type PluginExecuteFunc func(elemental.Operation, rufusmodels.RemoteProcessorModeValue, elemental.Identifiable, []string) error
