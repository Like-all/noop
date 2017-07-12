package main

import "gopkg.in/pelletier/go-toml.v1"
import "io/ioutil"
import "log"

const configFilePath string = "/etc/noop/noop.conf"

type Noop struct {
  ConfVendorDirectory     string
  ConfAvailableDirectory  string
  ConfEnabledDirectory    string
  EntryPoint              string
  Listen                  string
}

type Config struct {
  Noop Noop
}

func loadConfig()(c *Config, err error) {
  var fileData []byte
  if fileData, err = ioutil.ReadFile(configFilePath); err != nil {
      log.Fatalf("Error reading config file: %s\n", err.Error())
  }
  c = new(Config)
  err = toml.Unmarshal(fileData, &c)
  return
}
