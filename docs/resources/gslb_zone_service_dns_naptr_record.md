---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_gslb_zone_service_dns_naptr_record Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_gslb_zone_service_dns_naptr_record: Specify DNS NAPTR Record
  PLACEHOLDER
---

# thunder_gslb_zone_service_dns_naptr_record (Resource)

`thunder_gslb_zone_service_dns_naptr_record`: Specify DNS NAPTR Record

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_zone_service_dns_naptr_record" "thunder_gslb_zone_service_dns_naptr_record" {

  name         = "a11"
  service_name = "s25"
  service_port = 33159
  flag         = "1"
  naptr_target = "29"
  order        = 48506
  preference   = 55779
  regexp       = 0
  sampling_enable {
    counters1 = "all"
  }
  service_proto = "40"
  ttl           = 0
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `flag` (String) Specify the flag (e.g., a, s). Default is empty flag
- `name` (String) Name
- `naptr_target` (String) Specify the replacement or regular expression
- `service_name` (String) ServiceName
- `service_port` (String) ServicePort
- `service_proto` (String) Specify Service and Protocol

### Optional

- `order` (Number) Specify Order
- `preference` (Number) Specify Preference
- `regexp` (Number) Return the regular expression
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--sampling_enable))
- `ttl` (Number) Specify TTL
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--sampling_enable"></a>
### Nested Schema for `sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'naptr-hits': Number of times the NAPTR has been used;


