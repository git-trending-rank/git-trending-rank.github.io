package main

type Config struct {
	Frequency  string // "daily", "weekly", "monthly"
	OutputPath string
}

func LoadConfig() Config {
	// 暂时使用硬编码配置
	return Config{
		Frequency:  "daily",
		OutputPath: "output",
	}
}
