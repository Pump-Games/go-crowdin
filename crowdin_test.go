package crowdin

import (
	"net/http"
	"testing"
)

func TestNew_setToken(t *testing.T) {
	g := New("token", "project-name")
	if g.config.token != "token" {
		t.Errorf("Expected %v, got %v", "abc", g.config.token)
	}
	if g.config.project != "project-name" {
		t.Errorf("Expected %v, got %v", "project-name", g.config.project)
	}
}

func TestNew_setAPIBaseURL(t *testing.T) {
	g := New("token", "project-name")
	if g.config.apiBaseURL != apiBaseURL {
		t.Errorf("Expected %v, got %v", apiBaseURL, g.config.apiBaseURL)
	}
}

func TestNew_setStreamBaseURL(t *testing.T) {
	g := New("token", "project-name")
	if g.config.apiAccountBaseURL != apiAccountBaseURL {
		t.Errorf("Expected %v, got %v", apiAccountBaseURL, g.config.apiAccountBaseURL)
	}
}

func TestGitter_SetClient(t *testing.T) {
	setup()
	defer teardown()

	c := &http.Client{}
	crowdin.SetClient(c)

	if crowdin.config.client != c {
		t.Logf("Expected %v, got %v", c, crowdin.config.client)
	}
}
