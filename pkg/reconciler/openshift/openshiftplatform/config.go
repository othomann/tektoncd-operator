package openshiftplatform

import (
	k8sInstallerSet "github.com/tektoncd/operator/pkg/reconciler/kubernetes/tektoninstallerset"
	openshiftAddon "github.com/tektoncd/operator/pkg/reconciler/openshift/tektonaddon"
	"github.com/tektoncd/operator/pkg/reconciler/openshift/tektonchain"
	openshiftConfig "github.com/tektoncd/operator/pkg/reconciler/openshift/tektonconfig"
	"github.com/tektoncd/operator/pkg/reconciler/openshift/tektonhub"
	openshiftPipeline "github.com/tektoncd/operator/pkg/reconciler/openshift/tektonpipeline"
	openshiftTrigger "github.com/tektoncd/operator/pkg/reconciler/openshift/tektontrigger"
	"github.com/tektoncd/operator/pkg/reconciler/platform"
)

const (
	ControllerTektonAddon platform.ControllerName = "tektonaddon"
	PlatformNameOpenShift string                  = "openshift"
)

var (
	openshiftControllers = platform.ControllerMap{
		platform.ControllerTektonConfig:   openshiftConfig.NewController,
		platform.ControllerTektonPipeline: openshiftPipeline.NewController,
		platform.ControllerTektonTrigger:  openshiftTrigger.NewController,
		platform.ControllerTektonHub:      tektonhub.NewController,
		platform.ControllerTektonChain:    tektonchain.NewController,
		ControllerTektonAddon:             openshiftAddon.NewController,
		// there is no openshift specific extension for TektonInstallerSet Reconciler
		platform.ControllerTektonInstallerSet: k8sInstallerSet.NewController,
	}
)
