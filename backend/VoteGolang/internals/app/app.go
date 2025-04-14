package app

import (
	"VoteGolang/internals/app/conf"
	"VoteGolang/internals/app/connect"
	"VoteGolang/internals/data"
	"VoteGolang/internals/deliveries"
	"VoteGolang/internals/handlers"
	"VoteGolang/internals/repositories"
	auth2 "VoteGolang/internals/services/auth"
	"VoteGolang/internals/usecases"
	"VoteGolang/pkg/domain"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type App struct {
	Config *conf.Config
	DB     *gorm.DB
}

func NewApp() (*App, *usecases.AuthUseCase, domain.TokenManager, error) {
	config := conf.LoadConfig()
	db, err := connect.ConnectDB(config)

	if err != nil {
		return nil, nil, nil, err
	}
	//
	//err = MigrateAllTables(db)
	//if err != nil {
	//	return nil, nil, nil, err
	//}
	app := &App{
		Config: config,
		DB:     db,
	}

	userRepo := repositories.NewUserRepository(db)
	tokenManager := domain.NewJwtToken(config.JWTSecret)
	authUseCase := usecases.NewAuthUseCase(userRepo, tokenManager)

	return app, authUseCase, tokenManager, nil
}

func (a *App) Run(authUseCase *usecases.AuthUseCase, tokenManager domain.TokenManager) {
	log.Println("Starting server on port 8080...")

	mux := http.NewServeMux()
	//auth
	authHandler := auth2.NewAuthHandler(authUseCase, tokenManager)
	deliveries.LoginRegisterRoutes(mux, authHandler, tokenManager)

	_, ok := tokenManager.(domain.TokenManager)
	if !ok {
		log.Fatal("Token manager is not of type *auth.JwtTokenManager")
	}

	// Candidate
	candidateHandler := handlers.NewCandidateHandler(
		usecases.NewCandidateUseCase(
			repositories.NewCandidateRepository(a.DB),
			repositories.NewVoteRepository(a.DB),
		),
		tokenManager.(*domain.JwtToken),
	)
	deliveries.RegisterCandidateRoutes(mux, candidateHandler, tokenManager)

	// General News
	generalNewsHandler := handlers.NewGeneralNewsHandler(
		usecases.NewGeneralNewsUseCase(
			repositories.NewGeneralNewsRepository(a.DB),
		),
		tokenManager.(*domain.JwtToken),
	)
	deliveries.RegisterGeneralNewsRoutes(mux, generalNewsHandler, tokenManager)

	//Petitions
	petitionsHandler := handlers.NewPetitionHandler(
		usecases.NewPetitionUseCase(
			repositories.NewPetitionRepository(a.DB),
			repositories.NewPetitionVoteRepository(a.DB),
		),
		tokenManager.(*domain.JwtToken),
	)
	deliveries.RegisterPetitionRoutes(mux, petitionsHandler, tokenManager)

	//start
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func MigrateAllTables(db *gorm.DB) error {
	err := db.AutoMigrate(
		&data.User{},
		&data.Candidate{},
		&data.GeneralNews{},
		&data.Petition{},
		&data.PetitionVote{},
		&data.Vote{},
	)
	if err != nil {
		return err
	}
	log.Println("Database tables migrated successfully")
	return nil
}
