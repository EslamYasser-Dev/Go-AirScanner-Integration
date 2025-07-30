package mdns

import (
	"context"
	"time"

	"github.com/EslamYasser-Dev/Go-AirScanner-Integration/domain"
	"github.com/EslamYasser-Dev/Go-AirScanner-Integration/port"
	"github.com/grandcat/zeroconf"
)

type MDNSScannerDiscovery struct{}

func NewMDNSScannerDiscovery() port.ScannerDiscoveryPort {
	return &MDNSScannerDiscovery{}
}

func (m *MDNSScannerDiscovery) DiscoverScanners() ([]domain.Scanner, error) {
	var scanners []domain.Scanner

	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		return nil, err
	}

	entries := make(chan *zeroconf.ServiceEntry)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func(results <-chan *zeroconf.ServiceEntry) {
		for entry := range results {
			for _, ip := range entry.AddrIPv4 {
				scanners = append(scanners, domain.Scanner{
					Name: entry.Instance,
					IP:   ip.String(),
				})
			}
		}
	}(entries)

	err = resolver.Browse(ctx, "_uscan._tcp", "local.", entries)
	if err != nil {
		return nil, err
	}
	<-ctx.Done()
	return scanners, nil
}
