---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_threat_intel_webroot_database_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_threat_intel_webroot_database_oper: Operational Status for the object webroot-database
  PLACEHOLDER
---

# thunder_threat_intel_webroot_database_oper (Data Source)

`thunder_threat_intel_webroot_database_oper`: Operational Status for the object webroot-database

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_threat_intel_webroot_database_oper" "thunder_threat_intel_webroot_database_oper" {
}
output "get_threat_intel_webroot_database_oper" {
  value = ["${data.thunder_threat_intel_webroot_database_oper.thunder_threat_intel_webroot_database_oper}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--oper))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--oper"></a>
### Nested Schema for `oper`

Optional:

- `botnets` (Number)
- `connection_status` (String)
- `dos_attacks` (Number)
- `failure_reason` (String)
- `last_successful_connection` (String)
- `last_update_time` (String)
- `mobile_threats` (Number)
- `name` (String)
- `next_update_time` (String)
- `phishing` (Number)
- `proxy` (Number)
- `reputation` (Number)
- `scanners` (Number)
- `size` (String)
- `spam_sources` (Number)
- `status` (String)
- `tor_proxy` (Number)
- `total_entries` (Number)
- `version` (Number)
- `web_attacks` (Number)
- `windows_exploits` (Number)


