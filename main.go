package main

import (
	"fmt"
	"github-trending-scraper/internal/config"
	"github-trending-scraper/internal/lib"
	"github-trending-scraper/internal/scheduler"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	// 创建输出目录
	err := os.MkdirAll("output", 0755)
	if err != nil && !os.IsExist(err) {
		log.Fatal("Failed to create output directory:", err)
	}
	beginTime := time.Now()

	wg := new(sync.WaitGroup)
	wg.Add(3)
	go func() {
		defer wg.Done()
		cfg := config.LoadConfig(lib.Daily)
		scheduler.Schedule(cfg)
	}()
	go func() {
		defer wg.Done()
		cfg := config.LoadConfig(lib.Weekly)
		scheduler.Schedule(cfg)
	}()
	go func() {
		defer wg.Done()
		cfg := config.LoadConfig(lib.Monthly)
		scheduler.Schedule(cfg)
	}()
	wg.Wait()
	fmt.Println("All tasks completed, since: ", time.Since(beginTime))
}
