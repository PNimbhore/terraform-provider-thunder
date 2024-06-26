---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_gslb_zone_service_dns_cname_record_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_gslb_zone_service_dns_cname_record_stats: Statistics for the object dns-cname-record
  PLACEHOLDER
---

# thunder_gslb_zone_service_dns_cname_record_stats (Data Source)

`thunder_gslb_zone_service_dns_cname_record_stats`: Statistics for the object dns-cname-record

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_zone_service_dns_cname_record_stats" "thunder_gslb_zone_service_dns_cname_record_stats" {

  name         = "a11"
  service_name = "s25"
  service_port = 33159
  alias_name   = "3"
}
output "get_gslb_zone_service_dns_cname_record_stats" {
  value = ["${data.thunder_gslb_zone_service_dns_cname_record_stats.thunder_gslb_zone_service_dns_cname_record_stats}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `alias_name` (String) Specify the alias name
- `name` (String) Name
- `service_name` (String) ServiceName
- `service_port` (String) ServicePort

### Optional

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--stats))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--stats"></a>
### Nested Schema for `stats`

Optional:

- `cname_hits` (Number) Number of times the CNAME has been used


