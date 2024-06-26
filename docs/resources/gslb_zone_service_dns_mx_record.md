---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_gslb_zone_service_dns_mx_record Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_gslb_zone_service_dns_mx_record: Specify DNS MX Record
  PLACEHOLDER
---

# thunder_gslb_zone_service_dns_mx_record (Resource)

`thunder_gslb_zone_service_dns_mx_record`: Specify DNS MX Record

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_zone_service_dns_mx_record" "thunder_gslb_zone_service_dns_mx_record" {

  name         = "a11"
  service_name = "s25"
  service_port = 33159
  mx_name      = "2"
  priority     = 21510
  sampling_enable {
    counters1 = "all"
  }
  ttl = 0
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `mx_name` (String) Specify Domain Name
- `name` (String) Name
- `service_name` (String) ServiceName
- `service_port` (String) ServicePort

### Optional

- `priority` (Number) Specify Priority
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--sampling_enable))
- `ttl` (Number) Specify TTL
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--sampling_enable"></a>
### Nested Schema for `sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'hits': Number of times the record has been used;


