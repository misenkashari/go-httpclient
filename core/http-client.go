package core

func (c *HttpClient) Get(path string, config ...AdditionalConfig) (Response, error) {
	var conf *AdditionalConfig
	if len(config) > 0 {
		conf = &config[0]
	}
	return c.sendAsyncRequest("GET", path, nil, conf)
}
