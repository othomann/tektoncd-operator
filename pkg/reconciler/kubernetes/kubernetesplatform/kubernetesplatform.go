package kubernetesplatform

import (
	"github.com/tektoncd/operator/pkg/reconciler/platform"
	"knative.dev/pkg/injection"
	"log"
)

type KubernetesPlt struct {
	platform.PlatformConfig
	supportedControllers platform.ControllerMap
	enabledControllers   platform.ControllerMap
}

func NewKubernetesPlatform(pc platform.PlatformConfig) *KubernetesPlt {
	plt := KubernetesPlt{
		supportedControllers: kubernetesControllers,
		enabledControllers:   platform.ControllerMap{},
	}
	err := platform.ValidateControllerNames(pc.ControllerNames, plt.supportedControllers)
	if err != nil {
		log.Fatalf("invalid input: %v", err)
	}
	plt.PlatformConfig = pc
	plt.PlatformConfig.Name = PlatformNameKubernetes
	return &plt
}

func (kp *KubernetesPlt) SupportedControllers() platform.ControllerMap {
	return kp.supportedControllers
}

func (kp *KubernetesPlt) EnableControllers(rcs []platform.ControllerName) {
	for _, rc := range rcs {
		kp.enabledControllers[rc] = kp.supportedControllers[rc]
	}
}

func (kp *KubernetesPlt) ActiveControllers() []injection.ControllerConstructor {
	result := []injection.ControllerConstructor{}
	for _, ic := range kp.enabledControllers {
		result = append(result, ic)
	}
	return result
}

func (kp *KubernetesPlt) PlatformParams() platform.PlatformConfig {
	return kp.PlatformConfig
}
