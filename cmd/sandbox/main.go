package main

import (
	v1 "github.com/fengdotdev/golibs-1driveclient/v1"
)

func main() {

	ModeConfig()
}

func ModeConfig() {
	// Configura el cliente de datos
	err := v1.WriteEmptyFileDataClient()
	if err != nil {
		panic(err)
	}
}

func ModeRun() {
	// Lee la configuración del cliente de datos
	err := v1.UpdateConfigFromFile()
	if err != nil {
		panic(err)
	}
	token := v1.Authorize()
	v1.ListRootFiles(token)
}
