package config

import "testing"

func TestResolvePlaceholdersFromEnv(t *testing.T) {
	t.Setenv("DB_USER", "kratosx")
	t.Setenv("DB_PASSWORD", "secret")

	data := map[string]any{
		"database": map[string]any{
			"connect": map[string]any{
				"username": "${DB_USER}",
				"password": "${DB_PASSWORD}",
				"host":     "${DB_HOST:127.0.0.1}",
				"dsn":      "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST:127.0.0.1}/app",
			},
		},
	}

	if err := resolvePlaceholders(data); err != nil {
		t.Fatal(err)
	}

	connect := data["database"].(map[string]any)["connect"].(map[string]any)
	if connect["username"] != "kratosx" {
		t.Fatalf("username = %v", connect["username"])
	}
	if connect["password"] != "secret" {
		t.Fatalf("password = %v", connect["password"])
	}
	if connect["host"] != "127.0.0.1" {
		t.Fatalf("host = %v", connect["host"])
	}
	if connect["dsn"] != "postgres://kratosx:secret@127.0.0.1/app" {
		t.Fatalf("dsn = %v", connect["dsn"])
	}
}

func TestResolvePlaceholdersKeepsConfigReference(t *testing.T) {
	t.Setenv("APP_NAME", "from-env")

	data := map[string]any{
		"APP_NAME": "from-config",
		"app": map[string]any{
			"name": "${APP_NAME}",
			"id":   "${APP_NAME}-api",
		},
		"tags": []any{"$APP_NAME", "${MISSING:default}"},
	}

	if err := resolvePlaceholders(data); err != nil {
		t.Fatal(err)
	}

	app := data["app"].(map[string]any)
	if app["name"] != "from-config" {
		t.Fatalf("name = %v", app["name"])
	}
	if app["id"] != "from-config-api" {
		t.Fatalf("id = %v", app["id"])
	}

	tags := data["tags"].([]any)
	if tags[0] != "from-config" {
		t.Fatalf("tags[0] = %v", tags[0])
	}
	if tags[1] != "default" {
		t.Fatalf("tags[1] = %v", tags[1])
	}
}
