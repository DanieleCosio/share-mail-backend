package components

func JSComponent(name string, arguments ...string) string {
	if len(arguments) == 0 {
		return name
	}

	jsString := name + "("
	for _, param := range arguments {
		if param == "NULL" {
			jsString = jsString + "\"" + param + "\","
			continue
		}

		jsString = jsString + param + ","
	}
	jsString = jsString + ")"

	return jsString

}
