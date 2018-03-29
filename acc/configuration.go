package acc

type Configuration struct {
}

func (c *Configuration) GetString(key string) string {
	return ""
}

func (c *Configuration) GetInt(key string) int {
	return 0
}

func (c *Configuration) GetBoolean(key string) bool {
	return false
}
