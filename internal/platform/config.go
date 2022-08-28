package platform

import "log"

type Scope string

const (
	ScopeLocal  Scope = "LOCAL"
	ScopeDocker Scope = "DOCKER"
)

func (s Scope) GetDBConnectionURL() string {
	switch s {
	case ScopeLocal:
		return "postgres://postgres:postgres@localhost:5435/postgres"
	case ScopeDocker:
		return "postgres://postgres:postgres@postgres-class-api:5432/postgres"
	}
	log.Fatal("can't get connection url by scope")
	return ""
}

func (s Scope) GetAccountServiceURL() string {
	switch s {
	case ScopeLocal:
		return "http://localhost:5187"
	case ScopeDocker:
		return "http://host.docker.internal:5187"
	}
	log.Fatal("can't get account service url by scope")
	return ""
}
