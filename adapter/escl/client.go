package escl

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/EslamYasser-Dev/Go-AirScanner-Integration/domain"
	"github.com/EslamYasser-Dev/Go-AirScanner-Integration/port"
)

type EsclClient struct{}

func NewEsclClient() port.ScannerJobPort {
	return &EsclClient{}
}

func (e *EsclClient) FetchJobs(scanner domain.Scanner) ([]domain.ScanJob, error) {
	url := fmt.Sprintf("http://%s/eSCL/ScanJobs", scanner.IP)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("bad response: %s", resp.Status)
	}

	raw, _ := io.ReadAll(resp.Body)
	if !strings.Contains(string(raw), "JobId") {
		return []domain.ScanJob{}, nil
	}

	// Sample parser (not full XML)
	var jobs []domain.ScanJob
	lines := strings.Split(string(raw), "\n")
	for _, line := range lines {
		if strings.Contains(line, "<scan:JobId>") {
			id := strings.TrimPrefix(line, "<scan:JobId>")
			id = strings.TrimSuffix(id, "</scan:JobId>")
			jobs = append(jobs, domain.ScanJob{ID: id})
		}
	}
	return jobs, nil
}
