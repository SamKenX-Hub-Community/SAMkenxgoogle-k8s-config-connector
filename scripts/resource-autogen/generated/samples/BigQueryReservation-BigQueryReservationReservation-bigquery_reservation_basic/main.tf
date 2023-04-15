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
resource "google_bigquery_reservation" "reservation" {
	name           = "my-reservation"
	location       = "us-west2"
	// Set to 0 for testing purposes
	// In reality this would be larger than zero
	slot_capacity     = 0
	edition = "STANDARD"
	ignore_idle_slots = true
	concurrency       = 0
	autoscale {
   	  max_slots = 100
    }
}
```
