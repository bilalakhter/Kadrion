package customtypes

type Tconfig struct {
	API []struct {
		Endpoint string                 `yaml:"endpoint"`
		Method   string                 `yaml:"method"`
		JSON     map[string]interface{} `yaml:"JSON"`
	} `yaml:"API"`
}
