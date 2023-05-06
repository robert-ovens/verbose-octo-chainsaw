package main_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/cucumber/godog"
	"github.com/nsf/jsondiff"
	compute "github.com/robert-ovens/verbose-octo-chainsaw/api"
	"github.com/robert-ovens/verbose-octo-chainsaw/backend"
	"github.com/stretchr/testify/assert"
)

func (api *apiFeature) theFollowingInstances(arg1 *godog.Table) error {
	for i := 0; i < len(arg1.Rows); i++ {
		api.backend.Instances = append(api.backend.Instances, backend.Instance{
			Id:       arg1.Rows[i].Cells[0].Value,
			Image:    arg1.Rows[i].Cells[1].Value,
			SwapSize: arg1.Rows[i].Cells[2].Value,
			Type:     arg1.Rows[i].Cells[3].Value,
			Label:    arg1.Rows[i].Cells[4].Value,
		})
	}
	return nil
}

func (api *apiFeature) theListOfInstancesIsRequested() error {
	req, err := http.NewRequest("GET", "/", strings.NewReader(api.request))
	req.Header.Add("Accept", api.acceptHeader)

	if err != nil {
		return err
	}

	api.controller.List(api.resp, req)
	return nil
}

func (api *apiFeature) theResponseWillBe(arg1 *godog.DocString) error {
	resp, err := io.ReadAll(api.resp.Body)
	if err != nil {
		return err
	}

	opts := jsondiff.DefaultConsoleOptions()
	opts.PrintTypes = false
	comp, ff := jsondiff.Compare(resp, []byte(arg1.Content), &opts)
	if comp != 0 {
		return fmt.Errorf(ff)
	}
	return nil
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {

}

func InitializeScenario(ctx *godog.ScenarioContext) {
	api := &apiFeature{}

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {

		api.resetResponse(sc)
		return ctx, nil
	})
	ctx.Step(`^the following instances$`, api.theFollowingInstances)
	ctx.Step(`^the list of instances is requested$`, api.theListOfInstancesIsRequested)
	ctx.Step(`^the response will be$`, api.theResponseWillBe)

}
func TestFeatures(t *testing.T) {

	f, err := os.Create("./report.json")
	check(err)

	suite := godog.TestSuite{
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options: &godog.Options{
			Format:   "cucumber",
			Output:   f,
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
	f.Close()
}

func (a *apiFeature) resetResponse(interface{}) {
	a.resp = httptest.NewRecorder()
	a.backend = &MockBackend{}

	router := compute.NewDefaultApiController(compute.NewDefaultApiService(a.backend))
	a.controller = router.(*compute.DefaultApiController)
}

type apiFeature struct {
	resp         *httptest.ResponseRecorder
	request      string
	acceptHeader string
	controller   *compute.DefaultApiController
	backend      *MockBackend
}

func assertExpectedAndActual(a expectedAndActualAssertion, expected, actual interface{}, msgAndArgs ...interface{}) error {
	var t asserter
	a(&t, expected, actual, msgAndArgs...)
	return t.err
}
func (a *asserter) Errorf(format string, args ...interface{}) {
	a.err = fmt.Errorf(format, args...)
}

type asserter struct {
	err error
}
type expectedAndActualAssertion func(t assert.TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type MockBackend struct {
	Instances []backend.Instance
}

func NewMockBackend() MockBackend {
	return MockBackend{}
}

func (m MockBackend) GetInstances() ([]backend.Instance, error) {
	return m.Instances, nil
}
