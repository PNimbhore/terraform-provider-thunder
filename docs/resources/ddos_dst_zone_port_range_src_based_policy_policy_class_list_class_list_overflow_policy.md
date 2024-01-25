---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_dst_zone_port_range_src_based_policy_policy_class_list_class_list_overflow_policy Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_dst_zone_port_range_src_based_policy_policy_class_list_class_list_overflow_policy: Configure dynamic entry count overflow policy for class-list
  PLACEHOLDER
---

# thunder_ddos_dst_zone_port_range_src_based_policy_policy_class_list_class_list_overflow_policy (Resource)

`thunder_ddos_dst_zone_port_range_src_based_policy_policy_class_list_class_list_overflow_policy`: Configure dynamic entry count overflow policy for class-list

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_port_range_src_based_policy_policy_class_list_class_list_overflow_policy" "thunder_ddos_dst_zone_port_range_src_based_policy_policy_class_list_class_list_overflow_policy" {
  zone_name             = "test"
  port_range_start      = 20
  port_range_end        = 22
  protocol              = "tcp"
  src_based_policy_name = "test"
  class_list_name       = "test"
  dummy_name            = "configuration"
  glid                  = "3"
  action                = "deny"
  log_enable            = 1
  log_periodic          = 1
  user_tag              = "test"

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `class_list_name` (String) ClassListName
- `dummy_name` (String) 'configuration': Configure overflow policy for class-list;
- `port_range_end` (String) PortRangeEnd
- `port_range_start` (String) PortRangeStart
- `protocol` (String) Protocol
- `src_based_policy_name` (String) SrcBasedPolicyName
- `zone_name` (String) ZoneName

### Optional

- `action` (String) 'bypass': Always permit for the Source to bypass all feature & limit checks; 'deny': Blacklist incoming packets for service;
- `glid` (String) Global limit ID
- `log_enable` (Number) Enable logging
- `log_periodic` (Number) Enable log periodic
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object
- `zone_template` (Block List, Max: 1) (see [below for nested schema](#nestedblock--zone_template))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--zone_template"></a>
### Nested Schema for `zone_template`

Optional:

- `dns` (String) DDOS dns template
- `encap` (String) DDOS encap template (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)
- `http` (String) DDOS http template
- `logging` (String) DDOS logging template
- `quic` (String) DDOS quic template
- `sip` (String) DDOS sip template
- `ssl_l4` (String) DDOS ssl-l4 template
- `tcp` (String) DDOS tcp template
- `udp` (String) DDOS udp template

