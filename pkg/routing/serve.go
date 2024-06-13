package routing

import (
	"fmt"
	"go_first_project/pkg/config"
	"log"
)

func Serve() {
	r := GetRouter()

	config := config.GET()

	err := r.Run(fmt.Sprintf("%s:%v", config.Server.Host, config.Server.Port))

	if err != nil {
		log.Fatal("Error in routing")
		return
	}
}
