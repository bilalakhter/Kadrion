package toolinfo

import (
	"fmt"
)

func toolInfo() {
	fmt.Println("\n Kadrion - A Continuous Testing CLI Tool")
	fmt.Println("\n Usage:")
	fmt.Println("\n    kadrion apply tconfig.yaml")
	fmt.Println("\n Available scopes to use are:")
	fmt.Println("\n    api     Performance testing of an API Endpoint")
	fmt.Println("\n    cluster     Validate Kube deployment")
	fmt.Println("\n Documentation for configuration can be found at\n\n https://www.kadriontestops.tech/docs")
	fmt.Println()
	fmt.Println(" Additional Commands --help and --version")
	fmt.Println()
}
