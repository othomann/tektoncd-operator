package platform

import (
	"context"
	"fmt"
	"knative.dev/pkg/injection/sharedmain"
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
		return errorMsg(invalidNames.String(), supported.Keys())
	}
	return nil
}

func errorMsg(invalidNames string, validNames []ControllerName) error {
	return fmt.Errorf("un-identified controller names: %s supported names: %v", invalidNames, validNames)
}

func ContextWithPlatformName(pName string) context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, PlatformNameKey{}, pName)
	return ctx
}

func StartMain(p Platform) {
	pParams := p.PlatformParams()
	//ctx := ContextWithPlatformName(pParams.Name)
	sharedmain.Main(pParams.SharedMainName,
		p.ActiveControllers()...,
	)
}
