package servers

import (
	"context"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"errors"

	"github.com/cherryservers/cherryctl/internal/fakes"
	"github.com/cherryservers/cherryctl/internal/outputs"
	"github.com/cherryservers/cherrygo/v4"
	"github.com/spf13/cobra"
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

func setupCommand(t *testing.T, svc *fakes.ServerService, out *fakes.Outputer, args []string) *cobra.Command {
	t.Helper()

	dep := fakeDeps{svc: svc, out: out}
	c := Command{Deps: dep}

	cmd := c.Create()
	cmd.SetArgs(args)
	cmd.SilenceUsage = true
	return cmd
}

func createOK(_ context.Context, _ *cherrygo.CreateServer) (cherrygo.Server, *cherrygo.Response, error) {
	return cherrygo.Server{ID: 1}, nil, nil
}

func createErr(_ context.Context, _ *cherrygo.CreateServer) (cherrygo.Server, *cherrygo.Response, error) {
	return cherrygo.Server{}, nil, errors.New("test-error")
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
		{
			title: "multiple tags",
			args: []string{
				"--project-id",
				"1",
				"--region",
				"test-region",
				"--plan",
				"test-plan",
				"--tags",
				"first=foo,second=bar",
			},
			wantReqBody: &cherrygo.CreateServer{
				ProjectID:     1,
				Region:        "test-region",
				Plan:          "test-plan",
				SSHKeys:       []string{},
				IPAddresses:   []string{},
				Tags:          &map[string]string{"first": "foo", "second": "bar"},
				ConfigureIPv6: new(bool),
			},
		},
		{
			title: "whitespace tags",
			args: []string{
				"--project-id",
				"1",
				"--region",
				"test-region",
				"--plan",
				"test-plan",
				"--tags",
				"  first  =  foo ,  second=  bar ",
			},
			wantReqBody: &cherrygo.CreateServer{
				ProjectID:     1,
				Region:        "test-region",
				Plan:          "test-plan",
				SSHKeys:       []string{},
				IPAddresses:   []string{},
				Tags:          &map[string]string{"first": "foo", "second": "bar"},
				ConfigureIPv6: new(bool),
			},
		},
		{
			title: "key only tag",
			args: []string{
				"--project-id",
				"1",
				"--region",
				"test-region",
				"--plan",
				"test-plan",
				"--tags",
				"foo",
			},
			wantReqBody: &cherrygo.CreateServer{
				ProjectID:     1,
				Region:        "test-region",
				Plan:          "test-plan",
				SSHKeys:       []string{},
				IPAddresses:   []string{},
				Tags:          &map[string]string{"foo": ""},
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
			fakeSvc.SetCreate(createOK)
			cmd := setupCommand(t, &fakeSvc, &fakeOut, tc.args)

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
		title           string
		args            []string
		createFn        fakes.ServerCreationFunc
		outputErr       error
		wantMsg         *regexp.Regexp
		wantSvcCalls    int
		wantOutputCalls int
	}{
		{
			title:           "no project",
			args:            []string{"--region", "test-region", "--plan", "test-plan"},
			createFn:        createOK,
			wantMsg:         regexp.MustCompile(`^required flag\(s\) \"project-id\" not set$`),
			wantSvcCalls:    0,
			wantOutputCalls: 0,
		},
		{
			title:           "no plan",
			args:            []string{"--project-id", "1", "--region", "test-region"},
			createFn:        createOK,
			wantMsg:         regexp.MustCompile(`^required flag\(s\) \"plan\" not set$`),
			wantSvcCalls:    0,
			wantOutputCalls: 0,
		},
		{
			title:           "no region",
			args:            []string{"--project-id", "1", "--plan", "test-plan"},
			createFn:        createOK,
			wantMsg:         regexp.MustCompile(`^required flag\(s\) \"region\" not set$`),
			wantSvcCalls:    0,
			wantOutputCalls: 0,
		},
		{
			title: "missing userdata file",
			args: []string{
				"--project-id",
				"1",
				"--region",
				"test-region",
				"--plan",
				"test-plan",
				"--userdata-file",
				"no-file",
			},
			createFn:        createOK,
			wantMsg:         regexp.MustCompile(`^failed to read user-data file: .+$`),
			wantSvcCalls:    0,
			wantOutputCalls: 0,
		},
		{
			title: "missing ipxe file",
			args: []string{
				"--project-id",
				"1",
				"--region",
				"test-region",
				"--plan",
				"test-plan",
				"--ipxe-file",
				"no-file",
			},
			createFn:        createOK,
			wantMsg:         regexp.MustCompile(`^failed to read ipxe file: .+$`),
			wantSvcCalls:    0,
			wantOutputCalls: 0,
		},
		{
			title:           "api error",
			args:            []string{"--project-id", "1", "--region", "test-region", "--plan", "test-plan"},
			createFn:        createErr,
			wantMsg:         regexp.MustCompile(`^Could not provision a server: test-error$`),
			wantSvcCalls:    1,
			wantOutputCalls: 0,
		},
		{
			title:           "output error",
			args:            []string{"--project-id", "1", "--region", "test-region", "--plan", "test-plan"},
			createFn:        createOK,
			outputErr:       errors.New("test-error"),
			wantMsg:         regexp.MustCompile(`^test-error$`),
			wantSvcCalls:    1,
			wantOutputCalls: 1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.title, func(t *testing.T) {
			var (
				fakeSvc fakes.ServerService
				fakeOut fakes.Outputer
			)
			fakeSvc.SetCreate(tc.createFn)
			fakeOut.Err = tc.outputErr
			cmd := setupCommand(t, &fakeSvc, &fakeOut, tc.args)

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

			if len(fakeOut.Calls) != tc.wantOutputCalls {
				t.Fatalf("want %d output call, got %d", tc.wantOutputCalls, len(fakeOut.Calls))
			}
		})
	}
}
