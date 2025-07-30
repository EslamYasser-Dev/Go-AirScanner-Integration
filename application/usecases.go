package application

import (
	"fmt"

	"github.com/EslamYasser-Dev/Go-AirScanner-Integration/port"
)

type DiscoverScannersUseCase struct {
	Discovery port.ScannerDiscoveryPort
	Jobs      port.ScannerJobPort
}

func NewDiscoverScannersUseCase(d port.ScannerDiscoveryPort, j port.ScannerJobPort) *DiscoverScannersUseCase {
	return &DiscoverScannersUseCase{
		Discovery: d,
		Jobs:      j,
	}
}

func (uc *DiscoverScannersUseCase) Execute() {
	scanners, err := uc.Discovery.DiscoverScanners()
	if err != nil {
		fmt.Println("❌ Error discovering scanners:", err)
		return
	}

	for _, s := range scanners {
		fmt.Printf("🖨️ %s (%s)\n", s.Name, s.IP)
		jobs, err := uc.Jobs.FetchJobs(s)
		if err != nil {
			fmt.Println("⚠️ Failed to get jobs:", err)
			continue
		}
		for _, job := range jobs {
			fmt.Printf("   📄 Job ID: %s\n", job.ID)
		}
	}
}
