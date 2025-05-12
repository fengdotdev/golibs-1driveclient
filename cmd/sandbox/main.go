package main

import (
	v1 "github.com/fengdotdev/golibs-1driveclient/v1"
)

func main() {

	token := v1.Authorize()
	v1.ListRootFiles(token)

}
