package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func Loadenvir() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading your enironment you peasant")
	}
}
