package port

import "github.com/EslamYasser-Dev/Go-AirScanner-Integration/domain"

type ScannerDiscoveryPort interface {
	DiscoverScanners() ([]domain.Scanner, error)
}

type ScannerJobPort interface {
	FetchJobs(scanner domain.Scanner) ([]domain.ScanJob, error)
}
