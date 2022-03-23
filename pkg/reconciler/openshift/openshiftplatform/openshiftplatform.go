package openshiftplatform

import (
	"github.com/tektoncd/operator/pkg/reconciler/platform"
	"knative.dev/pkg/injection"
	"log"
)

type OpenShiftPlt struct {
	platform.PlatformConfig
	supportedControllers platform.ControllerMap
	enabledControllers   platform.ControllerMap
}

func NewOpenShiftPlatform(pc platform.PlatformConfig) *OpenShiftPlt {
	plt := OpenShiftPlt{
		supportedControllers: openshiftControllers,
		enabledControllers:   platform.ControllerMap{}}

	err := platform.ValidateControllerNames(pc.ControllerNames, plt.supportedControllers)
	if err != nil {
		log.Fatalf("invalid input: %v", err)
	}
	plt.PlatformConfig = pc
	plt.PlatformConfig.Name = PlatformNameOpenShift
	return &plt
}

func (op *OpenShiftPlt) SupportedControllers() platform.ControllerMap {
	return op.supportedControllers
}

func (op *OpenShiftPlt) EnableControllers(ctrlrNames []platform.ControllerName) {
	for _, rc := range ctrlrNames {
		op.enabledControllers[rc] = op.supportedControllers[rc]
	}
}

func (op *OpenShiftPlt) ActiveControllers() []injection.ControllerConstructor {
	result := []injection.ControllerConstructor{}
	for _, ic := range op.enabledControllers {
		result = append(result, ic)
	}
	return result
}

func (op *OpenShiftPlt) PlatformParams() platform.PlatformConfig {
	return op.PlatformConfig
}
