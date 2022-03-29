package platform

// Reconcilers common to all platforms
const (
	ControllerTektonConfig       ControllerName = "tektonconfig"
	ControllerTektonPipeline     ControllerName = "tektonpipeline"
	ControllerTektonTrigger      ControllerName = "tektontrigger"
	ControllerTektonInstallerSet ControllerName = "tektoninstallerset"
	ControllerTektonHub          ControllerName = "tektonhub"
	ControllerTektonChain        ControllerName = "tektonchain"
	EnvControllerNames           string         = "CONTROLLERS_NAMES"
	EnvControllerSubsetName      string         = "CONTROLLER_SUBSET_NAME"
)
