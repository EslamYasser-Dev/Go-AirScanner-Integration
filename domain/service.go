package domain

type ScannerService interface {
	ListScanners() ([]Scanner, error)
	GetJobs(scannerIP string) ([]ScanJob, error)
}
