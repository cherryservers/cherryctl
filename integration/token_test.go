package integration

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/cherryservers/cherryctl/cmd"
)

func discardStdout(t *testing.T) {
	t.Helper()

	nullOut, err := os.OpenFile(os.DevNull, os.O_WRONLY, os.ModeDevice)
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	old := os.Stdout
	os.Stdout = nullOut

	t.Cleanup(func() {
		os.Stdout = old
		if err := nullOut.Close(); err != nil {
			t.Errorf("failed to close dev null file: %s", err.Error())
		}
	})
}

func tempFileWithContent(t *testing.T, name, content string) *os.File {
	t.Helper()

	dir := t.TempDir()

	temp, err := os.Create(filepath.Join(dir, name))
	if err != nil {
		t.Fatalf("failed to create temp file: %s", err.Error())
	}
	t.Cleanup(func() {
		if err = temp.Close(); err != nil {
			t.Errorf("failed to close temp file: %s", err.Error())
		}
		if err = os.Remove(temp.Name()); err != nil {
			t.Errorf("failed to remove temp file: %s", err.Error())
		}
	})
	if _, err = fmt.Fprint(temp, content); err != nil {
		t.Fatalf("failed to write to temp file: %s", err.Error())
	}

	return temp
}

func TestTokenConfigHierarchy(t *testing.T) {
	const (
		tokenVar = "CHERRY_AUTH_TOKEN"
		apiKeyVar = "CHERRY_API_KEY"
	)

	// If no config directory is set, cherryctl will try to fallback
	// to some default directory, potentially interfering with test results.
	defaultConfig := tempFileWithContent(t, "default.yaml", "")
	t.Setenv("CHERRY_CONFIG", filepath.Dir(defaultConfig.Name()))

	t.Run("set via env var", func(t *testing.T) {
		testTokenConfigHierarchy(t, func(_ *cmd.Cli) {
			t.Setenv(tokenVar, "abc")
		})
	})

	t.Run("set via token flag", func(t *testing.T) {
		testTokenConfigHierarchy(t, func(cli *cmd.Cli) {
			cli.MainCmd.PersistentFlags().Set("token", "abc")
		})
	})

	t.Run("set via auth-token flag", func(t *testing.T) {
		testTokenConfigHierarchy(t, func(cli *cmd.Cli) {
			cli.MainCmd.PersistentFlags().Set("auth-token", "abc")
		})
	})

	t.Run("auth-token flag beats token flag", func(t *testing.T) {
		testTokenConfigHierarchy(t, func(cli *cmd.Cli) {
			cli.MainCmd.PersistentFlags().Set("token", "bad")
			cli.MainCmd.PersistentFlags().Set("auth-token", "abc")
		})
	})

	t.Run("set via default config context file", func(t *testing.T) {
		cfg := tempFileWithContent(t, "default.yaml", "token: abc\n")

		testTokenConfigHierarchy(t, func(cli *cmd.Cli) {
			cli.MainCmd.PersistentFlags().Set("config", filepath.Dir(cfg.Name()))
		})
	})

	t.Run("set via custom config context file", func(t *testing.T) {
		cfg := tempFileWithContent(t, "custom.yaml", "token: abc\n")

		testTokenConfigHierarchy(t, func(cli *cmd.Cli) {
			cli.MainCmd.PersistentFlags().Set("config", filepath.Dir(cfg.Name()))
			cli.MainCmd.PersistentFlags().Set("context", "custom")
		})
	})

	t.Run("env var beats token flag", func(t *testing.T) {
		testTokenConfigHierarchy(t, func(cli *cmd.Cli) {
			t.Setenv(tokenVar, "abc")
			cli.MainCmd.PersistentFlags().Set("token", "bad")
		})
	})

	// This happens because CHERRY_AUTH_TOKEN and auth-token
	// configure the same viper value, but viper prioritizes CLI
	// flags over env variables.
	t.Run("auth-token flag beats env var", func(t *testing.T) {
		testTokenConfigHierarchy(t, func(cli *cmd.Cli) {
			cli.MainCmd.PersistentFlags().Set("auth-token", "abc")
			t.Setenv(tokenVar, "bad")
		})
	})

	t.Run("token flag beats config file", func(t *testing.T) {
		cfg := tempFileWithContent(t, "default.yaml", "token: bad\n")

		testTokenConfigHierarchy(t, func(cli *cmd.Cli) {
			cli.MainCmd.PersistentFlags().Set("config", filepath.Dir(cfg.Name()))
			cli.MainCmd.PersistentFlags().Set("token", "abc")
		})
	})
	t.Run("api-key flag beats other flags", func(t *testing.T) {
		testTokenConfigHierarchy(t, func(cli *cmd.Cli) {
			cli.MainCmd.PersistentFlags().Set("api-key", "abc")
			cli.MainCmd.PersistentFlags().Set("token", "bad")
			cli.MainCmd.PersistentFlags().Set("auth-token", "bad")
		})
	})

	t.Run("api-key flag beats env var", func(t *testing.T) {
		testTokenConfigHierarchy(t, func(cli *cmd.Cli) {
			t.Setenv(apiKeyVar, "bad")
			cli.MainCmd.PersistentFlags().Set("api-key", "abc")
		})
	})

	t.Run("api-key env var beats auth-token env var", func(t *testing.T) {
		testTokenConfigHierarchy(t, func(cli *cmd.Cli) {
			t.Setenv(tokenVar, "bad")
			t.Setenv(apiKeyVar, "abc")
		})
	})

	t.Run("api-key env var beats config file", func(t *testing.T) {
		cfg := tempFileWithContent(t, "default.yaml", "token: bad\n")

		testTokenConfigHierarchy(t, func(cli *cmd.Cli) {
			t.Setenv(apiKeyVar, "abc")
			cli.MainCmd.PersistentFlags().Set("config", filepath.Dir(cfg.Name()))
		})
	})

	t.Run("api-key config file beats other in-file configs", func(t *testing.T) {
		cfg := tempFileWithContent(t, "default.yaml", "token: bad\nauth-token: bad\napi-key: abc\n")

		testTokenConfigHierarchy(t, func(cli *cmd.Cli) {
			cli.MainCmd.PersistentFlags().Set("config", filepath.Dir(cfg.Name()))
		})
	})
}

func testTokenConfigHierarchy(t *testing.T, preExecute func(cli *cmd.Cli)) {
	cli := cmd.NewCli()
	discardStdout(t)

	mux := http.NewServeMux()
	svc := httptest.NewServer(mux)
	defer svc.Close()

	respBodyFile, err := os.Open(filepath.Join("testdata", "get_usr.json"))
	if err != nil {
		t.Fatalf("%s", err.Error())
	}
	defer respBodyFile.Close()

	mux.HandleFunc("GET /v1/user", func(w http.ResponseWriter, r *http.Request) {
		got := r.Header.Get("Authorization")
		if got != "Bearer abc" {
			t.Errorf("want token %q, got: %q", "abc", got)
		}
		_, err := io.Copy(w, respBodyFile)
		if err != nil {
			t.Errorf("failed to write api response: %s", err)
		}
	})

	cli.MainCmd.SetArgs([]string{"user", "get"})
	cli.MainCmd.PersistentFlags().Set("api-url", svc.URL)

	preExecute(cli)

	if err := cli.MainCmd.Execute(); err != nil {
		t.Fatalf("failed to execute command: %s", err.Error())
	}
}
