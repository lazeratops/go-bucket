package httptest

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAvailableReleases(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name             string
		bbStatusCode     int
		bbBody           string
		wantReleaseCount int
		wantErr          error
	}{
		{
			name:         "bad status code",
			bbStatusCode: http.StatusInternalServerError,
			wantErr:      ErrFailedAPICall,
		},
		{
			name:         "one available book",
			bbStatusCode: http.StatusOK,
			// This example body was retrieved from the ReleasePopulator API docs
			bbBody: `[
			  {
				"id": 2355,
				"bookName": "The Little Giraffe",
				"authorName": "G. Neckton",
				"isAvailable": false
			  },
			  {
				"id": 123,
				"bookName": "The Big Pelican",
				"authorName": "P. Birdster",
				"isAvailable": true
			  }
			]`,
			wantReleaseCount: 1,
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.bbStatusCode)
				_, err := w.Write([]byte(tc.bbBody))
				require.NoError(t, err)
			}))

			defer testServer.Close()

			rp := NewReleasePopulator()
			rp.apiURL = testServer.URL

			gotReleases, gotErr := rp.GetAvailableReleases()
			require.ErrorIs(t, gotErr, tc.wantErr)
			if gotErr == nil {
				require.Len(t, gotReleases, tc.wantReleaseCount)
			}
		})
	}
}
