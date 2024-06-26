---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_gslb_site_slb_dev_vip_server_vip_server_name_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_gslb_site_slb_dev_vip_server_vip_server_name_stats: Statistics for the object vip-server-name
  PLACEHOLDER
---

# thunder_gslb_site_slb_dev_vip_server_vip_server_name_stats (Data Source)

`thunder_gslb_site_slb_dev_vip_server_vip_server_name_stats`: Statistics for the object vip-server-name

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_site_slb_dev_vip_server_vip_server_name_stats" "thunder_gslb_site_slb_dev_vip_server_vip_server_name_stats" {

  site_name   = "3"
  device_name = "25"
  vip_name    = "vs2"
}
output "get_gslb_site_slb_dev_vip_server_vip_server_name_stats" {
  value = ["${data.thunder_gslb_site_slb_dev_vip_server_vip_server_name_stats.thunder_gslb_site_slb_dev_vip_server_vip_server_name_stats}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `device_name` (String) DeviceName
- `site_name` (String) SiteName
- `vip_name` (String) Specify a VIP name for the SLB device

### Optional

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--stats))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--stats"></a>
### Nested Schema for `stats`

Optional:

- `dev_vip_hits` (Number) Number of times the service-ip was selected
- `dev_vip_recent` (Number) Recent hits


