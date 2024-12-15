package main

import (
	"fmt"

	"github.com/mohits-git/go-aggregator/internal/config"
)

func main() {
  cfg, err := config.Read()
  if err != nil {
    fmt.Println("Error reading config file\n", err)
    return
  }

  fmt.Println("OLDER CONFIG: ", cfg)

  cfg.SetUser("mohit")

  newcfg, err := config.Read()
  if err != nil {
    fmt.Println("Error reading config file\n", err)
    return
  }
  fmt.Println("NEW CONFIG: ", newcfg)
}
