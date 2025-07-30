package main

import (
	"github.com/EslamYasser-Dev/Go-AirScanner-Integration/adapter/escl"
	"github.com/EslamYasser-Dev/Go-AirScanner-Integration/adapter/mdns"
	"github.com/EslamYasser-Dev/Go-AirScanner-Integration/application"
)

func main() {
	mdns := mdns.NewMDNSScannerDiscovery()
	escl := escl.NewEsclClient()

	useCase := application.NewDiscoverScannersUseCase(mdns, escl)
	useCase.Execute()
}
