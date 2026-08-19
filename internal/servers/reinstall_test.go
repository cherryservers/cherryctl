package servers

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/cherryservers/cherryctl/internal/fakes"
	"github.com/cherryservers/cherrygo/v4"
)

func TestReinstall(t *testing.T) {
	tmpDir := t.TempDir()
	userdataPath := filepath.Join(tmpDir, "userdata")
	ipxePath := filepath.Join(tmpDir, "ipxe")
	err := os.WriteFile(userdataPath, []byte("test-userdata"), 0644)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = os.WriteFile(ipxePath, []byte("test-ipxe"), 0644)
	if err != nil {
		t.Fatal(err.Error())
	}

	cases := []struct {
		title       string
		args        []string
		wantReqBody *cherrygo.ReinstallServerFields
	}{
		{
			title: "only required args",
			args: []string{
				"1",
				"--hostname",
				"test-hostname",
				"--image",
				"test-image",
				"--password",
				"test-password",
			},
			wantReqBody: &cherrygo.ReinstallServerFields{
				Image:    "test-image",
				Hostname: "test-hostname",
				Password: "test-password",
				SSHKeys:  []string{},
			},
		},
		{
			title: "all args except ipxe",
			args: []string{
				"1",
				"--image",
				"test-image",
				"--hostname",
				"test-hostname",
				"--password",
				"test-password",
				"--ssh-keys",
				"1,2",
				"--userdata-file",
				userdataPath,
				"--os-partition-size",
				"1",
			},
			wantReqBody: &cherrygo.ReinstallServerFields{
				Image:           "test-image",
				Hostname:        "test-hostname",
				Password:        "test-password",
				SSHKeys:         []string{"1", "2"},
				UserData:        "dGVzdC11c2VyZGF0YQ==",
				OSPartitionSize: 1,
			},
		},
		{
			title: "ipxe no image",
			args: []string{
				"1",
				"--hostname",
				"test-hostname",
				"--password",
				"test-password",
				"--persist-ipxe",
				"--ipxe-file",
				ipxePath,
			},
			wantReqBody: &cherrygo.ReinstallServerFields{
				Image:       "custom_ipxe_install",
				Hostname:    "test-hostname",
				Password:    "test-password",
				IPXE:        "dGVzdC1pcHhl", // base64
				PersistIPXE: true,
				SSHKeys:     []string{},
			},
		},
		{
			title: "ipxe with image custom_ipxe_install",
			args: []string{
				"1",
				"--hostname",
				"test-hostname",
				"--password",
				"test-password",
				"--image",
				"custom_ipxe_install",
				"--persist-ipxe",
				"--ipxe-file",
				ipxePath,
			},
			wantReqBody: &cherrygo.ReinstallServerFields{
				Image:       "custom_ipxe_install",
				Hostname:    "test-hostname",
				Password:    "test-password",
				IPXE:        "dGVzdC1pcHhl", // base64
				PersistIPXE: true,
				SSHKeys:     []string{},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.title, func(t *testing.T) {
			var (
				fakeSvc fakes.ServerService
				buf     bytes.Buffer
			)
			fakeSvc.SetReinstall(reinstallOK)
			cmd := setupCobraReinstallCommand(setupCommand(&fakeSvc, nil), &buf, tc.args)

			err := cmd.Execute()
			if err != nil {
				t.Fatal(err.Error())
			}

			if len(fakeSvc.Calls) != 1 {
				t.Fatalf("want 1 api call, got %d", len(fakeSvc.Calls))
			}
			fakeSvc.Calls[0].AssertMethod(t, "Reinstall")
			fakeSvc.Calls[0].AssertParams(t, 1, tc.wantReqBody)

			const wantOut = "Server 1 reinstall has been started.\n"
			if got, _ := io.ReadAll(&buf); string(got) != wantOut {
				t.Fatalf("want output: %q, got: %q", wantOut, got)
			}
		})
	}
}

func TestReinstallWithErrorsExpected(t *testing.T) {
	tmpDir := t.TempDir()
	ipxePath := filepath.Join(tmpDir, "ipxe")
	err := os.WriteFile(ipxePath, []byte("test-ipxe"), 0644)
	if err != nil {
		t.Fatal(err.Error())
	}

	cases := []struct {
		title        string
		args         []string
		reinstallFn  fakes.ServerReinstallFunc
		wantMsg      *regexp.Regexp
		wantSvcCalls int
	}{
		{
			title: "no id",
			args: []string{
				"--hostname",
				"test-hostname",
				"--image",
				"test-image",
				"--password",
				"test-password",
			},
			reinstallFn:  reinstallOK,
			wantMsg:      regexp.MustCompile(`^accepts 1 arg\(s\), received 0$`),
			wantSvcCalls: 0,
		},
		{
			title: "non int id",
			args: []string{
				"a",
				"--hostname",
				"test-hostname",
				"--image",
				"test-image",
				"--password",
				"test-password",
			},
			reinstallFn:  reinstallOK,
			wantMsg:      regexp.MustCompile(`^invalid server id \"a\": .*$`),
			wantSvcCalls: 0,
		},
		{
			title:        "no hostname",
			args:         []string{"1", "--image", "test-image", "--password", "test-password"},
			reinstallFn:  reinstallOK,
			wantMsg:      regexp.MustCompile(`^required flag\(s\) \"hostname\" not set$`),
			wantSvcCalls: 0,
		},
		{
			title:        "no image",
			args:         []string{"1", "--hostname", "test-hostname", "--password", "test-password"},
			reinstallFn:  reinstallOK,
			wantMsg:      regexp.MustCompile(`^required flag\(s\) \"image\" not set$`),
			wantSvcCalls: 0,
		},
		{
			title:        "no password",
			args:         []string{"1", "--hostname", "test-hostname", "--image", "test-image"},
			reinstallFn:  reinstallOK,
			wantMsg:      regexp.MustCompile(`^required flag\(s\) \"password\" not set$`),
			wantSvcCalls: 0,
		},
		{
			title: "missing userdata file",
			args: []string{
				"1",
				"--hostname",
				"test-hostname",
				"--image",
				"test-image",
				"--password",
				"test-password",
				"--userdata-file",
				"no-file",
			},
			reinstallFn:  reinstallOK,
			wantMsg:      regexp.MustCompile(`^failed to read user-data file: .+$`),
			wantSvcCalls: 0,
		},
		{
			title: "missing ipxe file",
			args: []string{
				"1",
				"--hostname",
				"test-hostname",
				"--image",
				"test-image",
				"--password",
				"test-password",
				"--ipxe-file",
				"no-file",
			},
			reinstallFn:  reinstallOK,
			wantMsg:      regexp.MustCompile(`^failed to read ipxe file: .+$`),
			wantSvcCalls: 0,
		},
		{
			title: "persist-ipxe with no ipxe-file",
			args: []string{
				"1",
				"--hostname",
				"test-hostname",
				"--image",
				"test-image",
				"--password",
				"test-password",
				"--persist-ipxe",
			},
			reinstallFn:  reinstallOK,
			wantMsg:      regexp.MustCompile("^\"persist-ipxe\" requires \"ipxe-file\"$"),
			wantSvcCalls: 0,
		},
		{
			title: "incompatible image with ipxe",
			args: []string{
				"1",
				"--hostname",
				"test-hostname",
				"--image",
				"test-image",
				"--password",
				"test-password",
				"--ipxe-file",
				ipxePath,
			},
			reinstallFn:  reinstallOK,
			wantMsg:      regexp.MustCompile("image \"test-image\" is not compatible with iPXE: set image to \"custom_ipxe_install\" or leave it unset"),
			wantSvcCalls: 0,
		},
		{
			title: "api error",
			args: []string{
				"1",
				"--hostname",
				"test-hostname",
				"--image",
				"test-image",
				"--password",
				"test-password",
			},
			reinstallFn:  reinstallErr,
			wantMsg:      regexp.MustCompile(`^failed to reinstall server 1: test-error$`),
			wantSvcCalls: 1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.title, func(t *testing.T) {
			var (
				fakeSvc fakes.ServerService
				buf     bytes.Buffer
			)
			fakeSvc.SetReinstall(tc.reinstallFn)
			cmd := setupCobraReinstallCommand(setupCommand(&fakeSvc, nil), &buf, tc.args)

			err := cmd.Execute()
			if err == nil {
				t.Fatal("error shouldn't be nil")
			}
			if !tc.wantMsg.MatchString(err.Error()) {
				t.Fatalf("expected error msg that matches regex %q, got %q", tc.wantMsg, err.Error())
			}

			if len(fakeSvc.Calls) != tc.wantSvcCalls {
				t.Errorf("want %d api call, got %d", tc.wantSvcCalls, len(fakeSvc.Calls))
			}
		})
	}
}
