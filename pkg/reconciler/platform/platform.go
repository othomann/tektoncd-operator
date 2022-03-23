package platform

import (
	"context"
	"fmt"
	"knative.dev/pkg/injection/sharedmain"
	"log"
	"strings"
)

func EnableControllers(p Platform) {
	pParams := p.PlatformParams()
	if len(pParams.ControllerNames) == 0 {
		enableAllSupportedControllers(p)
		return
	}
	p.EnableControllers(pParams.ControllerNames)
}

func enableAllSupportedControllers(p Platform) {
	rNames := p.SupportedControllers()
	p.EnableControllers(rNames.Keys())
}

func ValidateControllerNames(cNames []ControllerName, supported ControllerMap) error {
	invalidNames := strings.Builder{}
	for _, cName := range cNames {
		if _, ok := supported[cName]; !ok {
			invalidNames.WriteString(string(cName))
			invalidNames.WriteString(",")
		}
	}
	if len(invalidNames.String()) != 0 {
		return fmt.Errorf("un-identified controller names: %s, supported names: %v", invalidNames.String(), supported.Keys())
	}
	return nil
}

func ContextWithPlatformName(pName string) context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, PlatformNameKey{}, pName)
	return ctx
}

func StartMain(p Platform) {
	pParams := p.PlatformParams()
	//ctx := ContextWithPlatformName(pParams.Name)
	log.Printf("sharedMainName: %v\n", pParams.SharedMainName)
	log.Printf("asdfasdfcontrollers: %v\n", pParams.ControllerNames)
	sharedmain.Main(pParams.SharedMainName,
		p.ActiveControllers()...,
	)
}
