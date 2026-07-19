package prometheus

func GetStorage() (*StorageMetric, error) {

	result, err := Query("node_filesystem_avail_bytes")

	if err != nil {
		return nil, err
	}

	value, err := ParseMetricValue(result)
	if err != nil {
		return nil, err
	}

	return &StorageMetric{
		AvailableGB: value / 1024 / 1024 / 1024,
	}, nil
}
