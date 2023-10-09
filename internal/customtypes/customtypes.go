package customtypes

type Tconfig struct {
	API struct {
		MaxConcurrentRequests int `yaml:"max_concurrent_requests"`
		NumberOfRequests      int `yaml:"number_of_requests"`
		Endpoints             []struct {
			Endpoint string `yaml:"endpoint"`
			Method   string `yaml:"method"`
			JSON     string `yaml:"JSON"`
		} `yaml:"endpoints"`
	} `yaml:"API"`
}
