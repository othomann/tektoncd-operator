package platform

import "knative.dev/pkg/injection"

type configReader func(config *PlatformConfig) error

type PlatformConfig struct {
	Name            string
	ControllerNames []ControllerName
	SharedMainName  string
}

type PlatformNameKey struct{}

type ControllerName string

type ControllerMap map[ControllerName]injection.ControllerConstructor

func (rm ControllerMap) Keys() []ControllerName {
	result := []ControllerName{}
	for k := range rm {
		result = append(result, k)
	}
	return result
}

type Platform interface {
	PlatformParams() PlatformConfig
	SupportedControllers() ControllerMap
	EnableControllers([]ControllerName)
	ActiveControllers() []injection.ControllerConstructor
}
