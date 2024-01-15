package meta_test

import (
	"context"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/chayut-eam/test-golang-hugeman/handler/meta"

	"github.com/alexliesenfeld/health"
	"github.com/stretchr/testify/assert"
)

const st = `# will be replaced by gitlab ci
name=Membership Point
version=5.9.0.0
built_timestamp=2022-09-28 00:00:00
commit=test commit
commit_timestamp=2022-09-28 00:00:00`

func TestHealth(t *testing.T) {
	context.Background()
	model := meta.HealthCheckConfig{
		CacheDuration:   5000,
		RefreshInterval: 5000,
		InitialDelay:    1000,
	}
	t.Run("test Success", func(t *testing.T) {
		meta.NewHealthChecker(model)
	})

}

func TestHandlerInfo(t *testing.T) {

	model := meta.HealthCheckConfig{
		CacheDuration:   5000,
		RefreshInterval: 5000,
		InitialDelay:    1000,
	}
	handlers := meta.NewMetaHandlerImpl(st, model)
	t.Run("test Success", func(t *testing.T) {
		handlers.Info()
	})
}

func TestHandlerHealth(t *testing.T) {
	model := meta.HealthCheckConfig{
		CacheDuration:   5000,
		RefreshInterval: 5000,
		InitialDelay:    1000,
	}
	handlers := meta.NewMetaHandlerImpl(st, model)
	t.Run("test Success", func(t *testing.T) {
		handlers.Health()
	})
}

func TestInfo(t *testing.T) {

	// success
	want := meta.InfoResponse{
		Name:            "Membership Point",
		Version:         "5.9.0.0",
		BuiltTimestamp:  "2022-09-28 00:00:00",
		Commit:          "test commit",
		CommitTimestamp: "2022-09-28 00:00:00",
	}
	st := `# will be replaced by gitlab ci
	name=Membership Point
	version=5.9.0.0
	built_timestamp=2022-09-28 00:00:00
	commit=test commit
	commit_timestamp=2022-09-28 00:00:00`

	result := meta.ParseInfo(st)

	assert.Equalf(t, want, result, "")

	st = `# will be replaced by gitlab ci
	name
	version=5.9.0.0
	built_timestamp=2022-09-28 00:00:00
	commit=test commit
	commit_timestamp=2022-09-28 00:00:00`

	meta.ParseInfo(st)

}

func TestWrite(t *testing.T) {
	handler := meta.JsonResultWriter{}

	detail := &map[string]health.CheckResult{}
	now := time.Now()
	stringErr := "error"
	(*detail)["test"] = health.CheckResult{
		Status:    health.StatusUp,
		Timestamp: &now,
		Error:     &stringErr,
	}

	result := &health.CheckerResult{
		Status:  health.StatusUp,
		Details: detail,
	}
	req := httptest.NewRequest("GET", "/", nil)

	rr := httptest.NewRecorder()

	handler.Write(result, 200, rr, req)
}
