package goodreads

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func cannedClient(testResponse string) *httptest.Server {
	return httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open(testResponse)
		if err != nil {
			log.Fatalf("unable to open test data file: %s", testResponse)
		}
		http.ServeContent(w, r, testResponse, time.Now(), f)
	}))
}

func TestSearch(t *testing.T) {
	server := cannedClient("test_data/search.xml")
	httpClient := server.Client()
	defer server.Close()

	client := NewClientWithHttpClientAndRootUrl("api-key", httpClient, server.URL)
	works, err := client.Search("Ender's Game")
	if err != nil {
		t.Error(err, "search must not return an error")
	}

	if len(works) != 20 {
		t.Errorf("invalid number of works in response: %d", len(works))
	}
}
