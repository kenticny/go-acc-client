package acc

type Configuration struct {
	cfg map[string]interface{}
}

// GetString returns a string type config by key
func (c *Configuration) GetString(key string) string {
	if _, ok := c.cfg[key]; !ok {
		return ""
	}
	if _, ok := c.cfg[key].(string); !ok {
		return ""
	}
	return c.cfg[key].(string)
}

// GetNumber returns a float64 type config by key
func (c *Configuration) GetNumber(key string) float64 {
	if _, ok := c.cfg[key]; !ok {
		return 0
	}
	if _, ok := c.cfg[key].(float64); !ok {
		return 0
	}
	return c.cfg[key].(float64)
}

// GetBoolean returns a boolean type config by key
func (c *Configuration) GetBoolean(key string) bool {
	if _, ok := c.cfg[key]; !ok {
		return false
	}
	if _, ok := c.cfg[key].(bool); !ok {
		return false
	}
	return c.cfg[key].(bool)
}
