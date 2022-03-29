package platform

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	ErrSharedMainNameEmpty = fmt.Errorf("sharedMainName cannot be empty string")
	ErrControllerNamesNil  = fmt.Errorf("ControllerNames slice should be non-nil")
)

func NewConfigFromFlags() PlatformConfig {
	config, err := newConfig(flagsConfigReader)
	if err != nil {
		log.Fatalf("unable to read platform from flags: %v", err)
	}
	return config
}

func NewConfigFromEnv() PlatformConfig {
	config, err := newConfig(envConfigReader)
	if err != nil {
		log.Fatalf("unable to read platform from env: %v", err)
	}
	return config
}

func envConfigReader(pc *PlatformConfig) error {
	ctrlArgs := os.Getenv(EnvControllerNames)
	c := os.Getenv(EnvControllerSubsetName)
	pc.SharedMainName = c
	pc.ControllerNames = strSliceToReconcilerNamesSlice(ctrlArgs)
	return nil
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
	if err := validateConfig(&config); err != nil {
		return PlatformConfig{}, err
	}
	return config, nil
}

func validateConfig(pc *PlatformConfig) error {
	violations := []string{}

	if len(pc.SharedMainName) == 0 {
		violations = append(violations, ErrSharedMainNameEmpty.Error())
	}
	// TODO: set a maximum length for pc.SharedMainName

	if pc.ControllerNames == nil {
		violations = append(violations, ErrControllerNamesNil.Error())
	}
	if len(violations) == 0 {
		return nil
	}
	return fmt.Errorf(strings.Join(violations, ","))
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
