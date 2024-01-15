package meta

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/alexliesenfeld/health"
	log "github.com/sirupsen/logrus"
)

type JsonResultWriter struct {
}

func NewHealthChecker(healthCheckConfig HealthCheckConfig) health.Checker {
	var (
		cacheDuration   = time.Millisecond * time.Duration(healthCheckConfig.CacheDuration)
		refreshInterval = time.Millisecond * time.Duration(healthCheckConfig.RefreshInterval)
		initialDelay    = time.Millisecond * time.Duration(healthCheckConfig.InitialDelay)
	)

	return health.NewChecker(
		health.WithCacheDuration(cacheDuration),
		health.WithPeriodicCheck(refreshInterval, initialDelay, fiberHealthCheck()),
		health.WithStatusListener(func(ctx context.Context, state health.CheckerState) {
			log.StandardLogger().Infof("Overall system health status has changed to %v.", strings.ToUpper(string(state.Status)))
		}),
	)
}

func fiberHealthCheck() health.Check {
	return health.Check{
		Name: "gin",
		Check: func(ctx context.Context) error {
			return nil
		},
		StatusListener: func(ctx context.Context, name string, state health.CheckState) {
			log.StandardLogger().Infof("%v status has changed to %v.", name, strings.ToUpper(string(state.Status)))
		},
	}
}

func (rw *JsonResultWriter) Write(result *health.CheckerResult, statusCode int, w http.ResponseWriter, r *http.Request) error {
	checkerResult := checkerResult(result)
	jsonResp, err := json.Marshal(checkerResult)
	if err != nil {
		return fmt.Errorf("cannot marshal response: %w", err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	_, err = w.Write(jsonResp)
	return err
}

type CheckerResult struct {
	Status  health.AvailabilityStatus `json:"status"`
	Details *map[string]CheckResult   `json:"details,omitempty"`
}

type CheckResult struct {
	Status    health.AvailabilityStatus `json:"status"`
	Timestamp time.Time                 `json:"timestamp,omitempty"`
	Error     *string                   `json:"error,omitempty"`
}

func checkerResult(result *health.CheckerResult) *CheckerResult {
	return &CheckerResult{
		Status:  result.Status,
		Details: checkResult(result.Details),
	}
}

func checkResult(details *map[string]health.CheckResult) *map[string]CheckResult {
	results := make(map[string]CheckResult)
	for name, result := range *details {
		checkResult := CheckResult{
			Status:    result.Status,
			Timestamp: *result.Timestamp,
			Error:     result.Error,
		}
		results[name] = checkResult
	}
	return &results
}
