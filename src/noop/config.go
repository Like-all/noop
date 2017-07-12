package main

import "gopkg.in/ini.v1"

func loadConfig()(cfg *File)  {
  cfg,err := ini.Load("/etc/noop/noop.conf")
  if err != nil {
    log.Fatalf("Could not open config file: %s", err.Error)
  }
  return
}
