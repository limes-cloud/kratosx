package web

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/limes-cloud/kratosx/config"
)

func TestApplyInjectionsSupportsScriptAndHTML(t *testing.T) {
	injections, err := prepareInjections([]*config.WebInjection{
		{
			Marker: "<!--__WEB_CONFIG__-->",
			Type:   "script",
			Value: map[string]any{
				"apiBaseUrl": "https://gw.example.com",
				"retry":      2,
			},
		},
		{
			Marker: "<!--HTML_CONFIG-->",
			Type:   "html",
			Value:  `<meta name="runtime" content="enabled">`,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	content := []byte("<head>\n<!--__WEB_CONFIG__-->\n<!--HTML_CONFIG-->\n</head>")
	got, found := applyInjections(content, injections)
	body := string(got)

	if !found["<!--__WEB_CONFIG__-->"] || !found["<!--HTML_CONFIG-->"] {
		t.Fatalf("markers were not found: %#v", found)
	}
	if !strings.Contains(body, `window.__WEB_CONFIG__={"apiBaseUrl":"https://gw.example.com","retry":2};`) {
		t.Fatalf("script injection mismatch: %s", body)
	}
	if !strings.Contains(body, `<meta name="runtime" content="enabled">`) {
		t.Fatalf("html injection mismatch: %s", body)
	}
	if strings.Count(body, "<!--__WEB_CONFIG__-->") != 1 || strings.Count(body, "<!--HTML_CONFIG-->") != 1 {
		t.Fatalf("markers should be retained exactly once: %s", body)
	}
}

func TestPrepareInjectionsEscapesScriptPayload(t *testing.T) {
	injections, err := prepareInjections([]*config.WebInjection{{
		Marker: "<!--CONFIG-->",
		Type:   "script",
		Value:  "</script><script>alert(1)</script>",
	}})
	if err != nil {
		t.Fatal(err)
	}
	got, _ := applyInjections([]byte("<!--CONFIG-->"), injections)
	if strings.Contains(string(got), "</script><script>alert(1)") {
		t.Fatalf("script payload was not escaped: %s", got)
	}
}

func TestPrepareInjectionsRejectsInvalidConfig(t *testing.T) {
	tests := []struct {
		name string
		item *config.WebInjection
	}{
		{name: "missing marker", item: &config.WebInjection{Type: "script", Value: map[string]any{}}},
		{name: "missing type", item: &config.WebInjection{Marker: "<!--CONFIG-->"}},
		{name: "marker without identifier", item: &config.WebInjection{Marker: "<!--APP CONFIG-->", Type: "script", Value: map[string]any{}}},
		{name: "invalid name", item: &config.WebInjection{Marker: "<!--window.alert(1)-->", Type: "script", Value: map[string]any{}}},
		{name: "invalid html value", item: &config.WebInjection{Marker: "<!--CONFIG-->", Type: "html", Value: map[string]any{}}},
		{name: "unsupported type", item: &config.WebInjection{Marker: "<!--CONFIG-->", Type: "style"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := prepareInjections([]*config.WebInjection{tt.item}); err == nil {
				t.Fatal("prepareInjections() accepted invalid config")
			}
		})
	}
}

func TestPreloadInjectsOnlyHTML(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "index.html"), []byte("<!--CONFIG-->"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "app.css"), []byte("<!--CONFIG-->"), 0o644); err != nil {
		t.Fatal(err)
	}

	injections, err := prepareInjections([]*config.WebInjection{{
		Marker: "<!--CONFIG-->",
		Type:   "html",
		Value:  "<meta charset=\"utf-8\">",
	}})
	if err != nil {
		t.Fatal(err)
	}

	c := newCache()
	found := c.preload(dir, defaultPreloadExts, injections)
	if !found["<!--CONFIG-->"] {
		t.Fatal("marker was not found during preload")
	}

	html, _ := c.Get(filepath.Join(dir, "index.html"))
	if !strings.Contains(string(html), `<meta charset="utf-8">`) {
		t.Fatalf("html file was not injected: %s", html)
	}
	css, _ := c.Get(filepath.Join(dir, "app.css"))
	if string(css) != "<!--CONFIG-->" {
		t.Fatalf("non-html file was modified: %s", css)
	}
}
