package components

import "testing"

func TestJSComponent(t *testing.T) {
	expectedResult := "testComponent(variable,object.key,functionCall(),arrayValue[0],\"NULL\",)"
	jsComponentName := "testComponent"

	result := JSComponent(jsComponentName, "variable", "object.key", "functionCall()", "arrayValue[0]", "NULL")
	if result != expectedResult {
		t.Error("Result is different from expected")
	}

}
