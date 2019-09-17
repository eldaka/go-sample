package xfers

import (
	"log"
	"os"
	"time"

	gcfg "gopkg.in/gcfg.v1"
)

// ServerConfig represent server config
type ServerConfig struct {
	Name        string
	Environment string
	HTTPTimeout time.Duration
}

// APIConfig represent Xfers API config
type APIConfig struct {
	BaseAddress     string
	GetAccountInfo  string
	GetTransferInfo string
}

// Config represent bot module config root
type Config struct {
	Server ServerConfig
	API    APIConfig
}

var (
	// Main represent MainConfig
	Main Config
)

//NewMainConfig initialize config
func NewMainConfig() {
	dirList := []string{"/etc/xfers", "files/etc/xfers"}

	for _, dir := range dirList {
		environ := os.Getenv("TKPENV")
		if environ == "" {
			environ = "development"
		}

		fname := dir + "/app." + environ + ".ini"
		err := gcfg.ReadFileInto(&Main, fname)
		if err == nil {
			log.Println("read config from ", fname)
			return
		}
	}

	log.Fatal("failed to read config")
}
