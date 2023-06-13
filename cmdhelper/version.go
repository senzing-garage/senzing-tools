package cmdhelper

import "fmt"

func Version(version string, iteration string) string {
	result := ""
	if iteration == "0" {
		result = version
	} else {
		result = fmt.Sprintf("%s-%s", version, iteration)
	}
	return result
}
