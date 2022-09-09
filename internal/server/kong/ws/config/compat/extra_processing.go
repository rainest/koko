package compat

import (
	"fmt"
	"strings"

	"github.com/kong/koko/internal/json"
	"github.com/kong/koko/internal/resource"
	"github.com/kong/koko/internal/server/kong/ws/config"
	"github.com/kong/koko/internal/versioning"
	"github.com/samber/lo"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"go.uber.org/zap"
)

var versionOlderThan300 = versioning.MustNewRange("< 3.0.0")

var (
	statsdDefaultIdentifiers = map[string]string{
		"consumer_identifier":  "custom_id",
		"service_identifier":   "service_name_or_host",
		"workspace_identifier": "workspace_id",
	}

	defaultMetrics = map[string][]string{
		"kong_latency":                    {"service_identifier"},
		"latency":                         {"service_identifier"},
		"request_count":                   {"service_identifier"},
		"request_per_user":                {"consumer_identifier", "service_identifier"},
		"request_size":                    {"service_identifier"},
		"response_size":                   {"service_identifier"},
		"shdict_usage":                    {"service_identifier"},
		"status_count":                    {"service_identifier"},
		"status_count_per_user":           {"consumer_identifier", "service_identifier"},
		"status_count_per_user_per_route": {"consumer_identifier", "service_identifier"},
		"status_count_per_workspace":      {"workspace_identifier"},
		"unique_users":                    {"consumer_identifier", "service_identifier"},
		"upstream_latency":                {"service_identifier"},
	}
)

const (
	awsLambdaExclusiveFieldChangeID = "P121"
)

func init() {
	err := config.ChangeRegistry.Register(config.Change{
		Metadata: config.ChangeMetadata{
			ID:       awsLambdaExclusiveFieldChangeID,
			Severity: config.ChangeSeverityError,
			Description: "For the 'aws-lambda' plugin, " +
				"'config.aws_region' and 'config.host' fields are set. " +
				"These fields were mutually exclusive for Kong gateway " +
				"versions < 3.0. The plugin configuration has been changed " +
				"to remove the 'config.host' field.",
			Resolution:       standardUpgradeMessage("3.0"),
			DocumentationURL: "",
		},
		SemverRange: versionsPre300,
		// none since the logic is hard-coded instead
		Update: config.ConfigTableUpdates{},
	})
	if err != nil {
		panic(err)
	}
}

// correctAWSLambdaMutuallyExclusiveFields handles 'aws_region' and 'host' fields, which were
// mutually exclusive until Kong version 2.8 but both are accepted in 3.x. If both are set
// with DPs < 3.x, the 'host' field will be dropped in order to prevent a failure in the DP.
func correctAWSLambdaMutuallyExclusiveFields(
	payload string,
	dataPlaneVersionStr string,
	tracker *config.ChangeTracker,
	logger *zap.Logger,
) string {
	pluginName := "aws-lambda"
	processedPayload := payload
	results := gjson.Get(processedPayload, "config_table.plugins.#(name=aws-lambda)#")
	indexUpdate := 0
	for _, res := range results.Array() {
		updatedRaw := res.Raw
		awsRegionResult := res.Get("config.aws_region")
		hostResult := res.Get("config.host")
		if awsRegionResult.Exists() && hostResult.Exists() {
			var (
				err      error
				pluginID = res.Get("id").String()
			)
			err = tracker.TrackForResource(awsLambdaExclusiveFieldChangeID,
				config.ResourceInfo{
					Type: string(resource.TypePlugin),
					ID:   pluginID,
				})
			if err != nil {
				logger.Error("failed to track version compatibility"+
					" change",
					zap.String("change-id", awsLambdaExclusiveFieldChangeID),
					zap.String("resource-type", "plugin"))
			}
			if updatedRaw, err = sjson.Delete(updatedRaw, "config.host"); err != nil {
				logger.With(zap.String("plugin", pluginName)).
					With(zap.String("field", "host")).
					With(zap.String("data-plane", dataPlaneVersionStr)).
					With(zap.Error(err)).
					Error("plugin configuration field was not removed from configuration")
			} else {
				logger.With(zap.String("plugin", pluginName)).
					With(zap.String("field", "host")).
					With(zap.String("data-plane", dataPlaneVersionStr)).
					Warn("removing plugin configuration field which is incompatible with data plane")
			}
		}

		// Update the processed payload
		resIndex := res.Index - indexUpdate
		updatedPayload := processedPayload[:resIndex] + updatedRaw +
			processedPayload[resIndex+len(res.Raw):]
		indexUpdate += len(processedPayload) - len(updatedPayload)
		processedPayload = updatedPayload
	}
	return processedPayload
}

func correctHTTPLogHeadersField(payload string) (string, error) {
	for i, plugin := range gjson.Get(payload, "config_table.plugins").Array() {
		if plugin.Get("name").Str != "http-log" {
			continue
		}

		for _, headers := range plugin.Get("config.headers").Array() {
			// When `headers.Type` is not a JSON object, no headers have been set, so we'll re-write it to be null.
			// This matches the behavior of the gateway when the headers field exists, but no value is defined.
			var newHeadersIface interface{}

			// Handle transforming the headers on each plugin from an object consisting
			// of a single string (`{"header-1": "value-1"}`), to an object consisting
			// of a single string value in an array (`{"header-1": ["value-1"]}`).
			if headers.Type == gjson.JSON {
				newHeaders := make(map[string][]string, len(headers.Indexes))
				newHeadersIface = newHeaders
				for key, values := range headers.Map() {
					// In <=2.8, while it is possible to set a header with an empty array of values,
					// the data plane won't send the header to the HTTP log endpoint with no value to
					// the defined HTTP endpoint. To match this behavior, we'll remove the header.
					if len(values.Str) == 0 {
						continue
					}

					newHeaders[key] = []string{values.Str}
				}
			}

			// Replace the headers for the http-log plugin.
			var err error
			if payload, err = sjson.SetOptions(
				payload,
				fmt.Sprintf("config_table.plugins.%d.config.headers", i),
				newHeadersIface,
				&sjson.Options{Optimistic: true},
			); err != nil {
				return "", err
			}
		}
	}

	return payload, nil
}

// correctStatsdIdentifiers addresses the needed statsd 3.0 schema changes for older DPs.
//
// The changes done in 3.0 remove some 'hard-coded' default values for metric identifiers and
// introduces specific schema fields to define those defaults. This function ensures the new
// schema works for older DPs by setting the missing default values for the metric identifiers.
func correctStatsdIdentifiers(
	payload string,
	dataPlaneVersionStr string,
	_ *config.ChangeTracker,
	logger *zap.Logger,
) string {
	log := logger.With(zap.String("plugin", "statsd")).
		With(zap.String("data-plane", dataPlaneVersionStr))

	processedPayload := payload
	indexUpdate := 0
	for _, res := range gjson.Get(processedPayload, "config_table.plugins.#(name=statsd)#").Array() {
		var (
			err         error
			metricsJSON []interface{}
			metrics     []string
			updatedRaw  = res.Raw
		)
		for _, metric := range gjson.Get(updatedRaw, "config.metrics").Array() {
			metricRaw := metric.Raw
			name := metric.Get("name").String()
			// Only evaluate default metrics.
			identifiers, ok := defaultMetrics[name]
			if !ok {
				metrics = append(metrics, metricRaw)
				continue
			}
			for key, defaultValue := range statsdDefaultIdentifiers {
				identifier := metric.Get(key)
				if lo.Contains(identifiers, key) && (!identifier.Exists() || identifier.String() == "") {
					log := log.With(zap.String("metric", name)).
						With(zap.String("field", key)).
						With(zap.String("condition", "missing value")).
						With(zap.String("new-value", defaultValue))
					if metricRaw, err = sjson.Set(metricRaw, key, defaultValue); err != nil {
						log.With(zap.Error(err)).
							Error("plugin configuration field was not updated in configuration")
					} else {
						log.Warn("updating plugin configuration field which is incompatible with data plane")
					}
				}
			}
			metrics = append(metrics, metricRaw)
		}
		metricsBytes := []byte(fmt.Sprintf("[%v]", strings.Join(metrics, ",")))
		if err = json.Unmarshal(metricsBytes, &metricsJSON); err != nil {
			log.With(zap.Error(err)).Error("plugin configuration field was not updated in configuration")
		}
		if updatedRaw, err = sjson.Set(updatedRaw, "config.metrics", metricsJSON); err != nil {
			log.With(zap.Any("new-value", metricsJSON)).
				With(zap.Error(err)).
				Error("plugin configuration field was not updated in configuration")
		}

		// Update the processed payload.
		resIndex := res.Index - indexUpdate
		updatedPayload := processedPayload[:resIndex] + updatedRaw +
			processedPayload[resIndex+len(res.Raw):]
		indexUpdate += len(processedPayload) - len(updatedPayload)
		processedPayload = updatedPayload
	}

	return processedPayload
}

func VersionCompatibilityExtraProcessing(payload string, dataPlaneVersion versioning.Version,
	tracker *config.ChangeTracker, logger *zap.Logger,
) (string, error) {
	dataPlaneVersionStr := dataPlaneVersion.String()
	processedPayload := payload

	if versionOlderThan300(dataPlaneVersion) {
		// 'aws_region' and 'host' are mutually exclusive for DP < 3.x
		processedPayload = correctAWSLambdaMutuallyExclusiveFields(
			processedPayload, dataPlaneVersionStr, tracker, logger)

		// The `headers` field on the `http-log` plugin changed from an array of strings to just a single string
		// for DP's >= 3.0. As such, we need to transform the headers back to a single string within an array.
		var err error
		if processedPayload, err = correctHTTPLogHeadersField(processedPayload); err != nil {
			return "", err
		}

		// Correct default metrics identifier for statsd in 2.x.x.x.
		processedPayload = correctStatsdIdentifiers(processedPayload, dataPlaneVersionStr, tracker, logger)
	}

	return processedPayload, nil
}
