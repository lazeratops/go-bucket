package httptest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	defaultAPIURL = "https://bookbeta.com/api/v1"
)

func releasesURL(apiURL string) string {
	return fmt.Sprintf("%s/releases", apiURL)
}

// ErrFailedAPICall is returned when we get an error or bad response from the ReleasePopulator API
var ErrFailedAPICall = errors.New("bad response from ReleasePopulator API")

// Release represents a book release, as per ReleasePopulator API
type Release struct {
	ID          int64  `json:"id"`
	BookName    string `json:"bookName"`
	AuthorName  string `json:"authorName"`
	IsAvailable bool   `json:"isAvailable"`
}

// ReleasePopulator recommends new book releases to users.
type ReleasePopulator struct {
	apiURL string
}

// NewReleasePopulator returns an instance of ReleasePopulator.
func NewReleasePopulator() *ReleasePopulator {
	return &ReleasePopulator{apiURL: defaultAPIURL}
}

// GetAvailableReleases returns a slice of releases that are marked as available.
func (rp *ReleasePopulator) GetAvailableReleases() ([]Release, error) {
	res, err := http.Get(releasesURL(rp.apiURL))
	if err != nil {
		return nil, fmt.Errorf("failed to get releases: %v: %w", err, ErrFailedAPICall)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d - %s: %w", res.StatusCode, res.Status, ErrFailedAPICall)
	}
	var releases []Release
	if err := json.NewDecoder(res.Body).Decode(&releases); err != nil {
		return nil, fmt.Errorf("failed to decode body into Release slice: %w", err)
	}
	var availableReleases []Release
	for _, r := range releases {
		if r.IsAvailable {
			availableReleases = append(availableReleases, r)
		}
	}
	return availableReleases, nil
}
