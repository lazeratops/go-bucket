package mocking

import "net/http"

//go:generate mockgen -package mocks -destination=./mocks/mock_communicator.go -source=communicator.go

// Communicator makes requests to the ReleasePopulator API
type Communicator interface {
	GetNewReleases() (*http.Response, error)
}

type bbCommunicator struct{}

// GetNewReleases retrieves all new releases from the ReleasePopulator API
func (c *bbCommunicator) GetNewReleases() (*http.Response, error) {
	return http.Get(releasesURL)
}
