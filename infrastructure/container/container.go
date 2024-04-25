package container

import (
	"github.com/MorrisMorrison/retfig/api"
	"github.com/MorrisMorrison/retfig/persistence/database"
	"github.com/MorrisMorrison/retfig/persistence/repositories"
	"github.com/MorrisMorrison/retfig/services"
)

type RepositoryContainer struct {
	EventRepository        *repositories.EventRepository
	PresentRepository      *repositories.PresentRepository
	VoteRepository         *repositories.VoteRepository
	CommentRepository      *repositories.CommentRepository
	ParticipantRrepository *repositories.ParticipantRepository
	ClaimRepository        *repositories.ClaimRepository
}

type ServiceContainer struct {
	EventService       *services.EventService
	PresentService     *services.PresentService
	VoteService        *services.VoteService
	CommentService     *services.CommentService
	ParticipantService *services.ParticipantService
	ClaimService       *services.ClaimService
}

type APIContainer struct {
	EventAPI       *api.EventAPI
	PresentAPI     *api.PresentAPI
	VoteAPI        *api.VoteAPI
	CommentAPI     *api.CommentAPI
	ParticipantAPI *api.ParticipantAPI
	ClaimAPI       *api.ClaimAPI
}

func NewRepositoryContainer(dbConn *database.Connection) *RepositoryContainer {
	eventRepo := repositories.NewEventRepository(dbConn)
	presentRepo := repositories.NewPresentRepository(dbConn)
	voteRepo := repositories.NewVoteRepository(dbConn)
	commentRepo := repositories.NewCommentRepository(dbConn)
	participantRepo := repositories.NewParticipantRepository(dbConn)
	claimRepository := repositories.NewClaimRepository(dbConn)

	return &RepositoryContainer{
		EventRepository:        eventRepo,
		PresentRepository:      presentRepo,
		VoteRepository:         voteRepo,
		CommentRepository:      commentRepo,
		ParticipantRrepository: participantRepo,
		ClaimRepository:        claimRepository,
	}
}

func NewServiceContainer(repositoryContainer *RepositoryContainer) *ServiceContainer {
	voteService := services.NewVoteService(repositoryContainer.VoteRepository)
	commentService := services.NewCommentService(repositoryContainer.CommentRepository)
	claimService := services.NewClaimService(repositoryContainer.ClaimRepository)
	presentService := services.NewPresentService(repositoryContainer.PresentRepository, voteService, commentService, claimService)
	participantService := services.NewParticipantService(repositoryContainer.ParticipantRrepository)
	eventService := services.NewEventService(repositoryContainer.EventRepository, presentService, participantService)

	return &ServiceContainer{
		VoteService:        voteService,
		CommentService:     commentService,
		PresentService:     presentService,
		EventService:       eventService,
		ParticipantService: participantService,
		ClaimService:       claimService,
	}
}

func NewAPIContainer(services *ServiceContainer) *APIContainer {
	return &APIContainer{
		PresentAPI:     api.NewPresentAPI(services.PresentService),
		CommentAPI:     api.NewCommentAPI(services.CommentService, services.PresentService),
		EventAPI:       api.NewEventAPI(services.EventService),
		VoteAPI:        api.NewVoteAPI(services.VoteService, services.PresentService),
		ParticipantAPI: api.NewParticipantAPI(services.ParticipantService, services.EventService),
		ClaimAPI:       api.NewClaimAPI(services.ClaimService, services.PresentService),
	}
}
