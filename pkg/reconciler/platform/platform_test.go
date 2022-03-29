package platform

import "testing"

func TestValidateControllerNames(t *testing.T) {
	tests := []struct {
		description    string
		ctrlNames      []ControllerName
		supportedCtrls ControllerMap
		expectedError  error
	}{
		{
			description: "return nil when all give controller names are supported (all names are present in the map)",
			ctrlNames:   []ControllerName{ControllerName("ctrl1"), ControllerName("ctrl1")},
			supportedCtrls: ControllerMap{
				ControllerName("ctrl1"): nil,
				ControllerName("ctrl2"): nil,
			},
			expectedError: nil,
		},
		{
			description: "return error some of the ControllerName are not supported",
			ctrlNames:   []ControllerName{ControllerName("ctrl1"), ControllerName("ctrlx")},
			supportedCtrls: ControllerMap{
				ControllerName("ctrl1"): nil,
				ControllerName("ctrl2"): nil,
			},
			expectedError: errorMsg("ctrlx", []ControllerName{ControllerName("ctrl1"), ControllerName("ctrl2")}),
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			err := ValidateControllerNames(test.ctrlNames, test.supportedCtrls)
			if test.expectedError == nil {
				assertNoError(t, err)
				return
			} else {

				assertError(t, err, test.expectedError)
			}

		})
	}
}
