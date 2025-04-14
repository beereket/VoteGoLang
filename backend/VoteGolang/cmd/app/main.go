package main

import (
	"VoteGolang/internals/app"
	"log"
)

func main() {
	appInstance, authUseCase, tokenManager, err := app.NewApp()
	if err != nil {
		log.Fatalf("Error initializing app: %v", err)
	}

	appInstance.Run(authUseCase, tokenManager)
}
