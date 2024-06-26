---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_automatic_update_config Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_automatic_update_config: Configure software update schedule
  PLACEHOLDER
---

# thunder_automatic_update_config (Resource)

`thunder_automatic_update_config`: Configure software update schedule

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_automatic_update_config" "thunder_automatic_update_config" {
  debug              = 1
  disable_ssl_verify = 1
  feature_name       = "ca-bundle"
  schedule           = 1
  weekly             = 1
  week_day           = "Wednesday"
  week_time          = "11:11"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `feature_name` (String) 'app-fw': Application Firewall Configuration; 'ca-bundle': CA Certificate Bundle; 'a10-threat-intel': A10 Threat intel class list; 'central-cert-pin-list': Central updated cert pinning list;

### Optional

- `daily` (Number) Every day
- `day_time` (String) Time of day to update (hh:mm) in 24 hour local time
- `debug` (Number) Enable libcurl debug option
- `disable_ssl_verify` (Number) Disable peer server certificate verification
- `schedule` (Number)
- `uuid` (String) uuid of the object
- `week_day` (String) 'Monday': Monday; 'Tuesday': Tuesday; 'Wednesday': Wednesday; 'Thursday': Thursday; 'Friday': Friday; 'Saturday': Saturday; 'Sunday': Sunday;
- `week_time` (String) Time of day to update (hh:mm) in 24 hour local time
- `weekly` (Number) Every week

### Read-Only

- `id` (String) The ID of this resource.


