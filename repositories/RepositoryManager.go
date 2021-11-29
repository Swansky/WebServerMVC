package repositories

import (
	"awesomeProject1/bdd"
	"awesomeProject1/config"
)

type RepositoryManager struct {
	userRepository  *UserRepository
	databaseManager *bdd.DatabaseManager
}

func NewRepositoryManager() *RepositoryManager {
	repositoryManager := new(RepositoryManager)
	repositoryManager.init()
	return repositoryManager
}

func (receiver *RepositoryManager) init() {
	settings := config.GetSettings()
	databaseConfig := settings.Database
	var credentials = bdd.NewDatabaseCredentials(databaseConfig.Host, databaseConfig.Username,
		databaseConfig.Password, databaseConfig.DatabaseName, databaseConfig.Port)

	receiver.databaseManager = bdd.NewDatabaseManager(credentials)
	receiver.userRepository = NewUserRepository(receiver.databaseManager)
}
