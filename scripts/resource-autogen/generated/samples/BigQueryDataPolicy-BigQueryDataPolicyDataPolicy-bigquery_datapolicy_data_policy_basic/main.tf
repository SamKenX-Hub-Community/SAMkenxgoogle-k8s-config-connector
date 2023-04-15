/**
 * Copyright 2022 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

```hcl
resource "google_bigquery_datapolicy_data_policy" "data_policy" {
    location         = "us-central1"
    data_policy_id   = "data_policy"
    policy_tag       = google_data_catalog_policy_tag.policy_tag.name
    data_policy_type = "COLUMN_LEVEL_SECURITY_POLICY"
  }

  resource "google_data_catalog_policy_tag" "policy_tag" {
    taxonomy     = google_data_catalog_taxonomy.taxonomy.id
    display_name = "Low security"
    description  = "A policy tag normally associated with low security items"
  }
  
  resource "google_data_catalog_taxonomy" "taxonomy" {
    region                 = "us-central1"
    display_name           = "taxonomy"
    description            = "A collection of policy tags"
    activated_policy_types = ["FINE_GRAINED_ACCESS_CONTROL"]
  }
```
