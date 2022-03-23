package kubernetesplatform

import (
	k8sChain "github.com/tektoncd/operator/pkg/reconciler/kubernetes/tektonchain"
	k8sConfig "github.com/tektoncd/operator/pkg/reconciler/kubernetes/tektonconfig"
	k8sDashboard "github.com/tektoncd/operator/pkg/reconciler/kubernetes/tektondashboard"
	k8sHub "github.com/tektoncd/operator/pkg/reconciler/kubernetes/tektonhub"
	k8sInstallerSet "github.com/tektoncd/operator/pkg/reconciler/kubernetes/tektoninstallerset"
	k8sPipeline "github.com/tektoncd/operator/pkg/reconciler/kubernetes/tektonpipeline"
	k8sResult "github.com/tektoncd/operator/pkg/reconciler/kubernetes/tektonresult"
	k8sTrigger "github.com/tektoncd/operator/pkg/reconciler/kubernetes/tektontrigger"
	"github.com/tektoncd/operator/pkg/reconciler/platform"
)

const (
	ControllerTektonDashboard platform.ControllerName = "tektondashboard"
	ControllerTektonResults   platform.ControllerName = "tektonresults"
	PlatformNameKubernetes    string                  = "kubernetes"
)

var (
	kubernetesControllers = platform.ControllerMap{
		platform.ControllerTektonConfig:       k8sConfig.NewController,
		platform.ControllerTektonPipeline:     k8sPipeline.NewController,
		platform.ControllerTektonTrigger:      k8sTrigger.NewController,
		platform.ControllerTektonHub:          k8sHub.NewController,
		platform.ControllerTektonChain:        k8sChain.NewController,
		platform.ControllerTektonInstallerSet: k8sInstallerSet.NewController,
		ControllerTektonDashboard:             k8sDashboard.NewController,
		ControllerTektonResults:               k8sResult.NewController,
	}
)
