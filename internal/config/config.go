package config

type Config struct {
	Frequency  string // "daily", "weekly", "monthly"
	OutputPath string
}

func LoadConfig(frequency string) Config {
	// 暂时使用硬编码配置
	return Config{
		Frequency:  frequency,
		OutputPath: "content/post",
	}
}
