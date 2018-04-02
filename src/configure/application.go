package configure

import (
	"github.com/fsnotify/fsnotify"
	log "github.com/kataras/golog"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func Application() {
	currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	viper.AddConfigPath(currentDir + "/config")
	viper.SetConfigType("json")
	viper.SetConfigName("settings") // name of config file (without extension)

	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		log.Fatal(err)
	}

	// look changes
	go viper.OnConfigChange(func(e fsnotify.Event) {
		log.Info("Config file changed: ", e.Name)
	})

	log.Debug("Config loaded: ")

	configKeys := viper.AllKeys()
	for _, key := range configKeys {
		log.Debug("[", key, "] = ", viper.Get(key))
	}
}
