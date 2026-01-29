package lazyload

import (
	"fmt"
	"sync"
	"time"
)

// Config represents a heavy object (e.g., Database connection, Big cache)
type Config struct {
	APIKey   string
	BaseURL  string
	LoadedAt time.Time
}

// Global state
type ConfigManager struct {
	instance *Config
	once     sync.Once
}



// loadHeavyConfig simulates a slow I/O operation (e.g., reading a 5GB file)
// DO NOT MODIFY THIS FUNCTION
func loadHeavyConfig() *Config {
	fmt.Println("⚠️  Heavy Load Started... (This should only print ONCE)")
	time.Sleep(5000 * time.Millisecond) // Simulate delay

	return &Config{
		APIKey:   "SECRET_KEY_123",
		BaseURL:  "https://api.agoda.com",
		LoadedAt: time.Now(),
	}
}

// GetConfig returns the singleton instance of Config.
// TODO: Implement this function using sync.Once to ensure thread safety.
func (cm *ConfigManager) GetConfig() *Config {
	cm.once.Do(func() {
		cm.instance = loadHeavyConfig()
		fmt.Println("⚠️  Done Loading")
	})
	
	return cm.instance
}

func Run() {
	var wg sync.WaitGroup
	configManager := ConfigManager{}
	
	// We simulate 10 concurrent requests to the config loader
	// If your logic is correct, "Heavy Load Started" prints ONCE.
	// If incorrect, it might print multiple times.
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			// Call the lazy loader
			cfg := configManager.GetConfig()
			
			// Verify we got the object
			fmt.Printf("Worker %d got config loaded at: %s\n", id, cfg.LoadedAt.Format("15:04:05.000"))
		}(i)
	}
	
	wg.Wait()
	fmt.Println("✅ All workers finished.")
}