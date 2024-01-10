// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package vertexai_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccVertexAIFeatureOnlineStoreFeatureview_vertexAiFeatureonlinestoreFeatureviewExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckVertexAIFeatureOnlineStoreFeatureviewDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccVertexAIFeatureOnlineStoreFeatureview_vertexAiFeatureonlinestoreFeatureviewExample(context),
			},
			{
				ResourceName:            "google_vertex_ai_feature_online_store_featureview.featureview",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "feature_online_store", "region", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccVertexAIFeatureOnlineStoreFeatureview_vertexAiFeatureonlinestoreFeatureviewExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_vertex_ai_feature_online_store" "featureonlinestore" {
  name = "tf_test_example_feature_view%{random_suffix}"
  labels = {
    foo = "bar"
  }
  region = "us-central1"
  bigtable {
    auto_scaling {
      min_node_count         = 1
      max_node_count         = 2
      cpu_utilization_target = 80
    }
  }
}

resource "google_bigquery_dataset" "tf-test-dataset" {

  dataset_id    = "tf_test_example_feature_view%{random_suffix}"
  friendly_name = "test"
  description   = "This is a test description"
  location      = "US"
}

resource "google_bigquery_table" "tf-test-table" {
  deletion_protection = false

  dataset_id = google_bigquery_dataset.tf-test-dataset.dataset_id
  table_id   = "tf_test_example_feature_view%{random_suffix}"
  schema     = <<EOF
  [
  {
    "name": "entity_id",
    "mode": "NULLABLE",
    "type": "STRING",
    "description": "Test default entity_id"
  },
    {
    "name": "test_entity_column",
    "mode": "NULLABLE",
    "type": "STRING",
    "description": "test secondary entity column"
  },
  {
    "name": "feature_timestamp",
    "mode": "NULLABLE",
    "type": "TIMESTAMP",
    "description": "Default timestamp value"
  }
]
EOF
}

resource "google_vertex_ai_feature_online_store_featureview" "featureview" {
  name                 = "tf_test_example_feature_view%{random_suffix}"
  region               = "us-central1"
  feature_online_store = google_vertex_ai_feature_online_store.featureonlinestore.name
  sync_config {
    cron = "0 0 * * *"
  }
  big_query_source {
    uri               = "bq://${google_bigquery_table.tf-test-table.project}.${google_bigquery_table.tf-test-table.dataset_id}.${google_bigquery_table.tf-test-table.table_id}"
    entity_id_columns = ["test_entity_column"]

  }
}

data "google_project" "project" {
  provider = google
}
`, context)
}

func testAccCheckVertexAIFeatureOnlineStoreFeatureviewDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_vertex_ai_feature_online_store_featureview" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{VertexAIBasePath}}projects/{{project}}/locations/{{region}}/featureOnlineStores/{{feature_online_store}}/featureViews/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("VertexAIFeatureOnlineStoreFeatureview still exists at %s", url)
			}
		}

		return nil
	}
}