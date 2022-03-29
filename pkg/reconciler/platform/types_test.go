package platform

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestKeys(t *testing.T) {
	tests := []struct {
		description  string
		cMap         ControllerMap
		expectedKeys []ControllerName
	}{
		{
			description: "returns map keys when map is non-empty",
			cMap: ControllerMap{
				ControllerName("key1"): nil,
				ControllerName("key2"): nil,
			},
			expectedKeys: []ControllerName{ControllerName("key1"), ControllerName("key2")},
		},
		{
			description:  "returns empty slice map is empty",
			cMap:         ControllerMap{},
			expectedKeys: []ControllerName{},
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got := test.cMap.Keys()
			if diff := cmp.Diff(got, test.expectedKeys); diff != "" {
				t.Errorf("expected: %v, got: %v", test.expectedKeys, got)
			}
		})
	}
}
