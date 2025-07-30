package escl

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/EslamYasser-Dev/Go-AirScanner-Integration/domain"
)

func (e *EsclClient) HasADF(scanner domain.Scanner) (bool, error) {
	url := fmt.Sprintf("http://%s/eSCL/ScannerCapabilities", scanner.IP)
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	return strings.Contains(string(body), "<pwg:Source>Feeder</pwg:Source>"), nil
}
