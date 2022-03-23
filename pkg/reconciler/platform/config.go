package platform

import (
	"flag"
	"log"
	"strings"
)

func NewConfigFromFlags() PlatformConfig {
	config, err := newConfig(flagsConfigReader)
	if err != nil {
		log.Fatalf("unable to read platform from flags: %v", err)
	}
	return config
}

func flagsConfigReader(pc *PlatformConfig) error {
	ctrlArgs := flag.String(
		"controllers",
		"",
		"comma separated list of names of controllers to be enabled (\"\" enables all controllers)",
	)
	c := flag.String("controller-subset-name",
		"tekton-operator",
		"name of the sharedMain process used in leader election")
	flag.Parse()
	pc.SharedMainName = *c
	pc.ControllerNames = strSliceToReconcilerNamesSlice(*ctrlArgs)
	return nil
}

func newConfig(inFn configReader) (PlatformConfig, error) {
	config := PlatformConfig{}
	err := inFn(&config)
	if err != nil {
		return PlatformConfig{}, err
	}
	return config, nil
}

func strSliceToReconcilerNamesSlice(s string) []ControllerName {
	result := []ControllerName{}
	if len(s) == 0 {
		return result
	}
	for _, val := range strings.Split(s, ",") {
		result = append(result, ControllerName(val))
	}
	return result
}
