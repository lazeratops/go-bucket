package mocking

import (
	"github.com/golang/mock/gomock"
	"github.com/lazeratops/go-bucket/external-api-testing/mocking/mocks"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"strings"
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
			mockCtrl := gomock.NewController(t)
			mockCommunicator := mocks.NewMockCommunicator(mockCtrl)
			body := io.NopCloser(strings.NewReader(tc.bbBody))
			mockCommunicator.EXPECT().GetNewReleases().Return(&http.Response{
				StatusCode: tc.bbStatusCode,
				Body:       body,
			}, nil).Times(1)

			// Since the test is in the same package we can access the unexported communicator field
			// But if you prefer putting the test in a separate package you could also do this
			// via an Option or passing the mock to the constructor func, for example.
			// (Although that would also expose that override functionality to other callers, which you
			// may or may not want to do).
			rp := NewReleasePopulator()
			rp.communicator = mockCommunicator
			gotReleases, gotErr := rp.GetAvailableReleases()
			require.ErrorIs(t, gotErr, tc.wantErr)
			if gotErr == nil {
				require.Len(t, gotReleases, tc.wantReleaseCount)
			}
		})
	}
}
