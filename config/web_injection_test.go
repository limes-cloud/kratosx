package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/go-kratos/kratos/v2/config/file"
)

func TestScanWebInjection(t *testing.T) {
	dir := t.TempDir()
	content := `
server:
  http:
    webServerDir: ./web
    injection:
      - marker: "<!--__APP_CONFIG__-->"
        type: script
        value:
          apiBaseUrl: https://gw.example.com
          retry: 2
      - marker: "<!--APP_META-->"
        type: html
        value: '<meta name="app" content="psou">'
`
	if err := os.WriteFile(filepath.Join(dir, "app.yaml"), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	cfg := New(file.NewSource(dir))
	if err := cfg.Load(); err != nil {
		t.Fatal(err)
	}

	httpCfg := cfg.App().Server.Http
	if httpCfg.WebServerDir != "./web" {
		t.Fatalf("webServerDir = %q", httpCfg.WebServerDir)
	}
	if len(httpCfg.Injection) != 2 {
		t.Fatalf("injection count = %d", len(httpCfg.Injection))
	}
	if httpCfg.Injection[0].Marker != "<!--__APP_CONFIG__-->" || httpCfg.Injection[0].Type != "script" {
		t.Fatalf("script injection = %#v", httpCfg.Injection[0])
	}
	value, ok := httpCfg.Injection[0].Value.(map[string]any)
	if !ok {
		t.Fatalf("script value type = %T", httpCfg.Injection[0].Value)
	}
	if value["apiBaseUrl"] != "https://gw.example.com" {
		t.Fatalf("apiBaseUrl = %v", value["apiBaseUrl"])
	}
	if httpCfg.Injection[1].Value != `<meta name="app" content="psou">` {
		t.Fatalf("html value = %#v", httpCfg.Injection[1].Value)
	}
}
