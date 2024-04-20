package context

import (
	"github.com/MorrisMorrison/retfig/infrastructure/container"
	"github.com/MorrisMorrison/retfig/persistence/database"
)

type ApplicationContext struct {
	APIs         *container.APIContainer
	Services     *container.ServiceContainer
	Repositories *container.RepositoryContainer
	DbConn       *database.Connection
}

func NewApplicationContext() *ApplicationContext {
	dbConn := database.NewConnection()
	repositories := container.NewRepositoryContainer(dbConn)
	services := container.NewServiceContainer(repositories)
	apis := container.NewAPIContainer(services)

	return &ApplicationContext{
		APIs:         apis,
		Services:     services,
		Repositories: repositories,
		DbConn:       dbConn,
	}
}
