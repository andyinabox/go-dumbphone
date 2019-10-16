package main

import (
	_ "github.com/sakirsensoy/genv/dotenv/autoload"

	"fmt"

	"github.com/andyinabox/go-dumbphone/internal/config"

	"github.com/andyinabox/go-dumbphone/pkg/directions"
)

func main() {
	fmt.Println(config.DirectionsConfig.HomeAddress)

}
