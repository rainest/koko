//go:build integration

package e2e

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/blang/semver/v4"
	"github.com/gavv/httpexpect/v2"
	"github.com/google/uuid"
	kongClient "github.com/kong/go-kong/kong"
	v1 "github.com/kong/koko/internal/gen/grpc/kong/admin/model/v1"
	"github.com/kong/koko/internal/json"
	"github.com/kong/koko/internal/private/test/run"
	"github.com/kong/koko/internal/test/util"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type update struct {
	field string
	value interface{}
}

type vcPlugins struct {
	name              string
	id                string
	config            string
	versionRange      string
	fieldUpdateChecks map[string][]update
	expectedConfig    string
}

func TestVersionCompatibility_Plugins(t *testing.T) {
	cleanup := run.Koko(t, run.KokoOpts{})
	defer cleanup()

	client := httpexpect.New(t, "http://localhost:3000")
	clusterClient := httpexpect.New(t, "http://localhost:4000")

	res := clusterClient.POST("/v1/clusters").WithBytes(cluster).Expect()
	res.Status(201)
	require.Equal(t, clusterID,
		res.JSON().Object().Value("id").String().Raw())
	require.Equal(t, "expired",
		res.JSON().Object().Value("state").String().Raw())

	conf := getKongConf()
	dpCleanup := run.KongDP(conf)
	defer dpCleanup()
	util.WaitForKong(t)
	util.WaitForKongAdminAPI(t)
	kongClient.RunWhenEnterprise(t, ">=2.3.0", kongClient.RequiredFeatures{})

	tests := []vcPlugins{
		// DP < 3.0
		//   - remove 'consumer_identifier_default', 'service_identifier_default',
		//     'workspace_identifier_default'
		{
			name: "statsd-advanced",
			id:   uuid.NewString(),
		},
	}
	for _, test := range tests {
		var config structpb.Struct
		if len(test.config) > 0 {
			require.Nil(t, json.ProtoJSONUnmarshal([]byte(test.config), &config))
		}

		plugin := &v1.Plugin{
			Id:        test.id,
			Name:      test.name,
			Config:    &config,
			Enabled:   wrapperspb.Bool(true),
			Protocols: []string{"http", "https"},
		}
		pluginBytes, err := json.ProtoJSONMarshal(plugin)
		require.Nil(t, err)
		res = client.POST("/v1/plugins").WithBytes(pluginBytes).WithQuery("cluster.id", clusterID).Expect()
		res.Status(201)
	}

	util.WaitFunc(t, func() error {
		err := ensurePlugins(tests)
		t.Log("plugin validation failed", err)
		return err
	})
}

func TestVersionCompatibilityTransformations_After_2_8_1(t *testing.T) {
	cleanup := run.Koko(t, run.KokoOpts{})
	defer cleanup()

	client := httpexpect.New(t, "http://localhost:3000")
	clusterClient := httpexpect.New(t, "http://localhost:4000")

	cluster := expiredCluster()
	res := clusterClient.POST("/v1/clusters").WithBytes(cluster).Expect()
	res.Status(http.StatusCreated)
	require.Equal(t, clusterID,
		res.JSON().Object().Value("id").String().Raw())
	require.Equal(t, "expired",
		res.JSON().Object().Value("state").String().Raw())

	tests := []vcPlugins{
		{
			name: "kafka-log",
			config: `{
				"authentication": {
					"mechanism": "SCRAM-SHA-512",
					"password": "test"
				},
				"topic": "koko",
				"cluster_name": "koko-cluster"
			}`,
			expectedConfig: `{
				"authentication": {
				  "mechanism": "SCRAM-SHA-512",
				  "password": "test",
				  "strategy": null,
				  "tokenauth": null,
				  "user": null
				},
				"bootstrap_servers": null,
				"cluster_name": "koko-cluster",
				"keepalive": 60000,
				"keepalive_enabled": false,
				"producer_async": true,
				"producer_async_buffering_limits_messages_in_memory": 50000,
				"producer_async_flush_timeout": 1000,
				"producer_request_acks": 1,
				"producer_request_limits_bytes_per_request": 1048576,
				"producer_request_limits_messages_per_request": 200,
				"producer_request_retries_backoff_timeout": 100,
				"producer_request_retries_max_attempts": 10,
				"producer_request_timeout": 2000,
				"security": {
				  "certificate_id": null,
				  "ssl": null
				},
				"timeout": 10000,
				"topic": "koko"
			  }`,
		},
		{
			name: "kafka-upstream",
			config: `{
				"authentication": {
					"mechanism": "SCRAM-SHA-512",
					"password": "test"
				},
				"topic": "koko",
				"cluster_name": "koko-cluster"
			}`,
			expectedConfig: `{
				"authentication": {
				  "mechanism": "SCRAM-SHA-512",
				  "password": "test",
				  "strategy": null,
				  "tokenauth": null,
				  "user": null
				},
				"bootstrap_servers": null,
				"cluster_name": "koko-cluster",
				"forward_body": true,
				"forward_headers": false,
				"forward_method": false,
				"forward_uri": false,
				"keepalive": 60000,
				"keepalive_enabled": false,
				"producer_async": true,
				"producer_async_buffering_limits_messages_in_memory": 50000,
				"producer_async_flush_timeout": 1000,
				"producer_request_acks": 1,
				"producer_request_limits_bytes_per_request": 1048576,
				"producer_request_limits_messages_per_request": 200,
				"producer_request_retries_backoff_timeout": 100,
				"producer_request_retries_max_attempts": 10,
				"producer_request_timeout": 2000,
				"security": {
				  "certificate_id": null,
				  "ssl": null
				},
				"timeout": 10000,
				"topic": "koko"
			  }`,
		},
		{
			name: "mtls-auth",
			config: `{
				"ca_certificates": [
					"abcdef12-3456-7890-abcd-ef1234567890"
				],
				"http_proxy_host": "test.com",
				"http_proxy_port": 80
			}`,
			expectedConfig: `{
				"anonymous": null,
				"authenticated_group_by": "CN",
				"ca_certificates": [
				  "abcdef12-3456-7890-abcd-ef1234567890"
				],
				"cache_ttl": 60,
				"cert_cache_ttl": 60000,
				"consumer_by": [
				  "username",
				  "custom_id"
				],
				"http_proxy_host": "test.com",
				"http_proxy_port": 80,
				"http_timeout": 30000,
				"https_proxy_host": null,
				"https_proxy_port": null,
				"revocation_check_mode": "IGNORE_CA_ERROR",
				"skip_consumer_lookup": false
			  }`,
		},
	}

	expectedConfig := &v1.TestingConfig{
		Plugins: make([]*v1.Plugin, 0, len(tests)),
	}

	for _, test := range tests {
		var config structpb.Struct
		require.NoError(t, json.ProtoJSONUnmarshal([]byte(test.config), &config))

		plugin := &v1.Plugin{
			Id:        uuid.NewString(),
			Name:      test.name,
			Config:    &config,
			Enabled:   wrapperspb.Bool(true),
			Protocols: []string{"http", "https"},
		}
		pluginBytes, err := json.ProtoJSONMarshal(plugin)
		require.NoError(t, err)
		res = client.POST("/v1/plugins").WithBytes(pluginBytes).WithQuery("cluster.id", clusterID).Expect()
		res.Status(http.StatusCreated)

		var expected structpb.Struct
		require.NoError(t, json.ProtoJSONUnmarshal([]byte(test.expectedConfig), &expected))
		expectedConfig.Plugins = append(expectedConfig.Plugins, &v1.Plugin{
			Id:        plugin.Id,
			Name:      plugin.Name,
			Config:    &expected,
			Enabled:   plugin.Enabled,
			Protocols: plugin.Protocols,
		})
	}

	dpCleanup := run.KongDP(getKongConf())
	defer dpCleanup()

	util.WaitForKong(t)
	util.WaitForKongAdminAPI(t)
	kongClient.RunWhenEnterprise(t, ">=2.8.1", kongClient.RequiredFeatures{})

	util.WaitFunc(t, func() error {
		err := util.EnsureConfig(expectedConfig)
		t.Log("plugin validation failed", err)
		return err
	})
}

// Ensure that extra-processing logic correctly injects metric identifier default values
// for DP < 3.0, which doesn't support setting default identifiers in the schema.
func TestVersionCompatibilityTransformations_StatsdAdvancedMetricsDefaults(t *testing.T) {
	cleanup := run.Koko(t, run.KokoOpts{})
	defer cleanup()

	client := httpexpect.New(t, "http://localhost:3000")
	clusterClient := httpexpect.New(t, "http://localhost:4000")

	cluster := expiredCluster()
	res := clusterClient.POST("/v1/clusters").WithBytes(cluster).Expect()
	res.Status(http.StatusCreated)
	require.Equal(t, clusterID,
		res.JSON().Object().Value("id").String().Raw())
	require.Equal(t, "expired",
		res.JSON().Object().Value("state").String().Raw())

	tests := []vcPlugins{
		// these 3 test metrics ensure that:
		//
		// 1. a non-default metric with no identifiers doesn't get filled with identifiers
		// 2. a default metric with identifiers gets filled with defaults
		// 3. a default metric with non-default identifiers doesn't get changed
		{
			config: `{
				"metrics": [
				  {
					"name": "cache_datastore_misses_total",
					"sample_rate": 1,
					"stat_type": "counter"
				  },
				  {
					"name": "response_size",
					"stat_type": "timer"
				  },
				  {
					"name": "request_count",
					"sample_rate": 1,
					"service_identifier": "service_name",
					"stat_type": "counter"
				  }
				]
			}`,
			expectedConfig: `{
				"allow_status_codes": null,
				"host": "localhost",
				"hostname_in_prefix": false,
				"metrics": [
					{
						"consumer_identifier": null,
						"name": "cache_datastore_misses_total",
						"sample_rate": 1,
						"service_identifier": null,
						"stat_type": "counter",
						"workspace_identifier": null
					},
					{
						"workspace_identifier": null,
						"name": "response_size",
						"stat_type": "timer",
						"sample_rate": null,
						"consumer_identifier": null,
						"service_identifier": "service_name_or_host"
					},
					{
						"service_identifier": "service_name",
						"stat_type": "counter",
						"sample_rate": 1,
						"consumer_identifier": null,
						"workspace_identifier": null,
						"name": "request_count"
					}
				],
				"port": 8125,
				"prefix": "kong",
				"udp_packet_size": 0,
				"use_tcp": false
			}`,
		},
	}
	expectedConfig := &v1.TestingConfig{
		Plugins: make([]*v1.Plugin, 0, len(tests)),
	}

	for _, test := range tests {
		var config structpb.Struct
		require.NoError(t, json.ProtoJSONUnmarshal([]byte(test.config), &config))

		plugin := &v1.Plugin{
			Id:        uuid.NewString(),
			Name:      "statsd-advanced",
			Config:    &config,
			Enabled:   wrapperspb.Bool(true),
			Protocols: []string{"http", "https"},
		}
		pluginBytes, err := json.ProtoJSONMarshal(plugin)
		require.NoError(t, err)
		res = client.POST("/v1/plugins").WithBytes(pluginBytes).WithQuery("cluster.id", clusterID).Expect()
		res.Status(http.StatusCreated)

		var expected structpb.Struct
		require.NoError(t, json.ProtoJSONUnmarshal([]byte(test.expectedConfig), &expected))
		expectedConfig.Plugins = append(expectedConfig.Plugins, &v1.Plugin{
			Id:        plugin.Id,
			Name:      plugin.Name,
			Config:    &expected,
			Enabled:   plugin.Enabled,
			Protocols: plugin.Protocols,
		})
	}

	dpCleanup := run.KongDP(getKongConf())
	defer dpCleanup()

	util.WaitForKong(t)
	util.WaitForKongAdminAPI(t)
	kongClient.RunWhenEnterprise(t, "<3.0.0", kongClient.RequiredFeatures{})

	util.WaitFunc(t, func() error {
		err := util.EnsureConfig(expectedConfig)
		t.Log("plugin validation failed", err)
		return err
	})
}

// Ensure that a plugin configured with exactly the same pre-3.0 default configuration,
// doesn't get changed via extra-processing logic.
func TestVersionCompatibilityTransformations_StatsdAdvancedDefaultConfig(t *testing.T) {
	cleanup := run.Koko(t, run.KokoOpts{})
	defer cleanup()

	client := httpexpect.New(t, "http://localhost:3000")
	clusterClient := httpexpect.New(t, "http://localhost:4000")

	cluster := expiredCluster()
	res := clusterClient.POST("/v1/clusters").WithBytes(cluster).Expect()
	res.Status(http.StatusCreated)
	require.Equal(t, clusterID,
		res.JSON().Object().Value("id").String().Raw())
	require.Equal(t, "expired",
		res.JSON().Object().Value("state").String().Raw())

	tests := []vcPlugins{
		{
			name:   "statsd-advanced",
			config: `{}`,
			expectedConfig: `{
				"allow_status_codes": null,
				"host": "localhost",
				"hostname_in_prefix": false,
				"metrics": [
					{
					  "workspace_identifier": null,
					  "name": "request_count",
					  "stat_type": "counter",
					  "sample_rate": 1,
					  "service_identifier": "service_name_or_host",
					  "consumer_identifier": null
					},
					{
					  "workspace_identifier": null,
					  "name": "latency",
					  "stat_type": "timer",
					  "sample_rate": null,
					  "service_identifier": "service_name_or_host",
					  "consumer_identifier": null
					},
					{
					  "workspace_identifier": null,
					  "name": "request_size",
					  "stat_type": "timer",
					  "sample_rate": null,
					  "service_identifier": "service_name_or_host",
					  "consumer_identifier": null
					},
					{
					  "workspace_identifier": null,
					  "name": "status_count",
					  "stat_type": "counter",
					  "sample_rate": 1,
					  "service_identifier": "service_name_or_host",
					  "consumer_identifier": null
					},
					{
					  "workspace_identifier": null,
					  "name": "response_size",
					  "stat_type": "timer",
					  "sample_rate": null,
					  "service_identifier": "service_name_or_host",
					  "consumer_identifier": null
					},
					{
					  "workspace_identifier": null,
					  "name": "unique_users",
					  "stat_type": "set",
					  "sample_rate": null,
					  "service_identifier": "service_name_or_host",
					  "consumer_identifier": "custom_id"
					},
					{
					  "workspace_identifier": null,
					  "name": "request_per_user",
					  "stat_type": "counter",
					  "sample_rate": 1,
					  "service_identifier": "service_name_or_host",
					  "consumer_identifier": "custom_id"
					},
					{
					  "workspace_identifier": null,
					  "name": "upstream_latency",
					  "stat_type": "timer",
					  "sample_rate": null,
					  "service_identifier": "service_name_or_host",
					  "consumer_identifier": null
					},
					{
					  "workspace_identifier": null,
					  "name": "kong_latency",
					  "stat_type": "timer",
					  "sample_rate": null,
					  "service_identifier": "service_name_or_host",
					  "consumer_identifier": null
					},
					{
					  "workspace_identifier": null,
					  "name": "status_count_per_user",
					  "stat_type": "counter",
					  "sample_rate": 1,
					  "service_identifier": "service_name_or_host",
					  "consumer_identifier": "custom_id"
					},
					{
					  "workspace_identifier": "workspace_id",
					  "name": "status_count_per_workspace",
					  "stat_type": "counter",
					  "sample_rate": 1,
					  "service_identifier": null,
					  "consumer_identifier": null
					},
					{
					  "workspace_identifier": null,
					  "name": "status_count_per_user_per_route",
					  "stat_type": "counter",
					  "sample_rate": 1,
					  "service_identifier": "service_name_or_host",
					  "consumer_identifier": "custom_id"
					},
					{
					  "workspace_identifier": null,
					  "name": "shdict_usage",
					  "stat_type": "gauge",
					  "sample_rate": 1,
					  "service_identifier": "service_name_or_host",
					  "consumer_identifier": null
					},
					{
					  "workspace_identifier": null,
					  "name": "cache_datastore_hits_total",
					  "stat_type": "counter",
					  "sample_rate": 1,
					  "service_identifier": null,
					  "consumer_identifier": null
					},
					{
					  "workspace_identifier": null,
					  "name": "cache_datastore_misses_total",
					  "stat_type": "counter",
					  "sample_rate": 1,
					  "service_identifier": null,
					  "consumer_identifier": null
					}
				],
				"port": 8125,
				"prefix": "kong",
				"udp_packet_size": 0,
				"use_tcp": false
			}`,
		},
	}
	expectedConfig := &v1.TestingConfig{
		Plugins: make([]*v1.Plugin, 0, len(tests)),
	}

	for _, test := range tests {
		var config structpb.Struct
		require.NoError(t, json.ProtoJSONUnmarshal([]byte(test.config), &config))

		plugin := &v1.Plugin{
			Id:        uuid.NewString(),
			Name:      "statsd-advanced",
			Config:    &config,
			Enabled:   wrapperspb.Bool(true),
			Protocols: []string{"http", "https"},
		}
		pluginBytes, err := json.ProtoJSONMarshal(plugin)
		require.NoError(t, err)
		res = client.POST("/v1/plugins").WithBytes(pluginBytes).WithQuery("cluster.id", clusterID).Expect()
		res.Status(http.StatusCreated)

		var expected structpb.Struct
		require.NoError(t, json.ProtoJSONUnmarshal([]byte(test.expectedConfig), &expected))
		expectedConfig.Plugins = append(expectedConfig.Plugins, &v1.Plugin{
			Id:        plugin.Id,
			Name:      plugin.Name,
			Config:    &expected,
			Enabled:   plugin.Enabled,
			Protocols: plugin.Protocols,
		})
	}

	dpCleanup := run.KongDP(getKongConf())
	defer dpCleanup()

	util.WaitForKong(t)
	util.WaitForKongAdminAPI(t)
	kongClient.RunWhenEnterprise(t, "<3.0.0", kongClient.RequiredFeatures{})

	util.WaitFunc(t, func() error {
		err := util.EnsureConfig(expectedConfig)
		t.Log("plugin validation failed", err)
		return err
	})
}

func ensurePlugins(plugins []vcPlugins) error {
	kongAdmin, err := kongClient.NewClient(util.BasedKongAdminAPIAddr, nil)
	if err != nil {
		return fmt.Errorf("create go client for kong: %v", err)
	}
	ctx := context.Background()
	info, err := kongAdmin.Root(ctx)
	if err != nil {
		return fmt.Errorf("fetching Kong Gateway info: %v", err)
	}
	dataPlaneVersion, err := kongClient.ParseSemanticVersion(kongClient.VersionFromInfo(info))
	if err != nil {
		return fmt.Errorf("parsing Kong Gateway version: %v", err)
	}
	dataPlanePlugins, err := kongAdmin.Plugins.ListAll(ctx)
	if err != nil {
		return fmt.Errorf("fetching plugins: %v", err)
	}

	// Remove plugins that may not be expected due to data plane version
	var expectedPlugins []vcPlugins
	for _, plugin := range plugins {
		addPlugin := true
		if len(plugin.versionRange) > 0 {
			version := semver.MustParseRange(plugin.versionRange)
			if !version(dataPlaneVersion) {
				addPlugin = false
			}
		}
		if addPlugin {
			expectedPlugins = append(expectedPlugins, plugin)
		}
	}

	// Because configurations may vary validation occurs via the name and ID for
	// removal items and a special test for updates will be performed which
	// verifies update occurred properly based on versions
	if len(expectedPlugins) != len(dataPlanePlugins) {
		return fmt.Errorf("plugins configured count does not match [%d != %d]", len(expectedPlugins), len(dataPlanePlugins))
	}
	var failedPlugins []string
	var missingPlugins []string
	for _, plugin := range expectedPlugins {
		found := false
		for _, dataPlanePlugin := range dataPlanePlugins {
			if plugin.name == *dataPlanePlugin.Name && plugin.id == *dataPlanePlugin.ID {
				// Ensure field updates occurred and validate
				if len(plugin.fieldUpdateChecks) > 0 {
					config, err := json.ProtoJSONMarshal(dataPlanePlugin.Config)
					if err != nil {
						return fmt.Errorf("marshal %s plugin config: %v", plugin.name, err)
					}
					configStr := string(config)

					for version, updates := range plugin.fieldUpdateChecks {
						version := semver.MustParseRange(version)
						if version(dataPlaneVersion) {
							for _, update := range updates {
								res := gjson.Get(configStr, update.field)
								if !res.Exists() || res.Value() != update.value {
									failedPlugins = append(failedPlugins, plugin.name)
									break
								}
							}
						}
					}
				}

				found = true
				break
			}
		}
		if !found {
			missingPlugins = append(missingPlugins, plugin.name)
		}
	}

	if len(missingPlugins) > 0 {
		return fmt.Errorf("failed to discover plugins %s", strings.Join(missingPlugins, ","))
	}
	if len(failedPlugins) > 0 {
		return fmt.Errorf("failed to validate plugin updates %s", strings.Join(failedPlugins, ","))
	}

	return nil
}
