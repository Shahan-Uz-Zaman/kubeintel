package prometheus

func GetNetwork() (*NetworkMetric, error) {

	// Receive bytes/sec
	rxResult, err := Query(`sum(rate(node_network_receive_bytes_total[5m]))`)
	if err != nil {
		return nil, err
	}

	// Transmit bytes/sec
	txResult, err := Query(`sum(rate(node_network_transmit_bytes_total[5m]))`)
	if err != nil {
		return nil, err
	}

	rx, err := ParseMetricValue(rxResult)
	if err != nil {
		return nil, err
	}

	tx, err := ParseMetricValue(txResult)
	if err != nil {
		return nil, err
	}

	return &NetworkMetric{
		// Convert Bytes/sec → MB/sec
		Receive:  rx / 1024 / 1024,
		Transmit: tx / 1024 / 1024,
	}, nil
}
