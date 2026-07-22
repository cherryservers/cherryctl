package servers

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cherryservers/cherryctl/internal/fakes"
	"github.com/cherryservers/cherryctl/internal/outputs"
	"github.com/cherryservers/cherrygo/v4"
)

type fakeDeps struct {
	svc  *fakes.ServerService
	out  *fakes.Outputer
	opts *cherrygo.GetOptions
}

func (fd fakeDeps) GetOpts() *cherrygo.GetOptions {
	return fd.opts
}

func (fd fakeDeps) Client() cherrygo.ServersService {
	return fd.svc
}

func (fd fakeDeps) Outputer() outputs.Outputer {
	return fd.out
}

func newTrue() *bool {
	b := true
	return &b
}

func TestCreate(t *testing.T) {
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
		wantReqBody *cherrygo.CreateServer
	}{
		{
			title: "only required args",
			args:  []string{"--project-id", "1", "--region", "test-region", "--plan", "test-plan"},
			wantReqBody: &cherrygo.CreateServer{
				ProjectID:     1,
				Region:        "test-region",
				Plan:          "test-plan",
				SSHKeys:       []string{},
				IPAddresses:   []string{},
				Tags:          &map[string]string{},
				ConfigureIPv6: new(bool),
			},
		},
		{
			title: "all args",
			args: []string{
				"--project-id",
				"1",
				"--region",
				"test-region",
				"--plan",
				"test-plan",
				"--hostname",
				"test-hostname",
				"--image",
				"test-image",
				"--ssh-keys",
				"1,2",
				"--ip-addresses",
				"1,2",
				"--os-partition-size",
				"1",
				"--tags",
				"env=test",
				"--spot-instance",
				"--storage-id",
				"1",
				"--cycle",
				"test-cycle",
				"--discount",
				"test-discount",
				"--enable-ipv6",
				"--ipxe-file",
				ipxePath,
				"--userdata-file",
				userdataPath,
			},
			wantReqBody: &cherrygo.CreateServer{
				ProjectID:       1,
				Region:          "test-region",
				Plan:            "test-plan",
				Hostname:        "test-hostname",
				Image:           "test-image",
				SSHKeys:         []string{"1", "2"},
				IPAddresses:     []string{"1", "2"},
				OSPartitionSize: 1,
				Tags:            &map[string]string{"env": "test"},
				SpotInstance:    true,
				StorageID:       1,
				Cycle:           "test-cycle",
				DiscountCode:    "test-discount",
				ConfigureIPv6:   newTrue(),
				IPXE:            "dGVzdC1pcHhl", // base64
				UserData:        "dGVzdC11c2VyZGF0YQ==",
			},
		},
		{
			title: "shorthands",
			args:  []string{"-p", "1", "--region", "test-region", "--plan", "test-plan"},
			wantReqBody: &cherrygo.CreateServer{
				ProjectID:     1,
				Region:        "test-region",
				Plan:          "test-plan",
				SSHKeys:       []string{},
				IPAddresses:   []string{},
				Tags:          &map[string]string{},
				ConfigureIPv6: new(bool),
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.title, func(t *testing.T) {
			var (
				fakeSvc fakes.ServerService
				fakeOut fakes.Outputer
			)
			dep := fakeDeps{svc: &fakeSvc, out: &fakeOut}
			c := Command{Deps: dep}

			cmd := c.Create()
			cmd.SetArgs(tc.args)
			cmd.SilenceUsage = true

			err := cmd.Execute()
			if err != nil {
				t.Fatal(err.Error())
			}

			if len(fakeSvc.Calls) != 1 {
				t.Fatalf("want 1 api call, got %d", len(fakeSvc.Calls))
			}
			fakeSvc.Calls[0].AssertMethod(t, "Create")
			fakeSvc.Calls[0].AssertParams(t, tc.wantReqBody)

			if len(fakeOut.Calls) != 1 {
				t.Fatalf("want 1 output call, got %d", len(fakeOut.Calls))
			}
			wantTh := []string{"ID", "Name", "Hostname", "Image", "State", "Region"}
			wantTd := [][]string{{"1", "", "", "", "", ""}}
			fakeOut.Calls[0].Assert(t, cherrygo.Server{ID: 1}, wantTh, wantTd)

		})
	}
}

func TestCreateWithErrorsExpected(t *testing.T) {
	cases := []struct {
		title string
		args  []string
	}{
		{
			title: "no project",
			args:  []string{"--region", "test-region", "--plan", "test-plan"},
		},
		{
			title: "no plan",
			args:  []string{"--project-id", "1", "--region", "test-region"},
		},
		{
			title: "no region",
			args:  []string{"--project-id", "1", "--plan", "test-plan"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.title, func(t *testing.T) {
			c := Command{}
			cmd := c.Create()
			cmd.SetArgs(tc.args)
			cmd.SilenceUsage = true

			err := cmd.Execute()
			if err == nil {
				t.Fatal("error expected")
			}
		})
	}
}
