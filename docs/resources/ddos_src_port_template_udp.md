---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_src_port_template_udp Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_src_port_template_udp: UDP template configuration
  PLACEHOLDER
---

# thunder_ddos_src_port_template_udp (Resource)

`thunder_ddos_src_port_template_udp`: UDP template configuration

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_src_port_template_udp" "thunder_ddos_src_port_template_udp" {
  drop_ntp_monlist = 1
  filter_list {
    udp_filter_seq       = 4
    udp_filter_regex     = "609"
    udp_filter_unmatched = 1
    udp_filter_action    = "blacklist-src"
    user_tag             = "64"
  }
  max_payload_size = 620
  min_payload_size = 82
  user_tag         = "114"
  name             = "testing"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) DDOS UDP Template Name

### Optional

- `drop_ntp_monlist` (Number) Drop NTP monlist request/response
- `filter_list` (Block List) (see [below for nested schema](#nestedblock--filter_list))
- `max_payload_size` (Number) Maximum UDP payload size for each single packet
- `min_payload_size` (Number) Minimum UDP payload size for each single packet
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--filter_list"></a>
### Nested Schema for `filter_list`

Required:

- `udp_filter_seq` (Number) Sequence number

Optional:

- `byte_offset_filter` (String) Filter Expression using Berkeley Packet Filter syntax
- `udp_filter_action` (String) 'blacklist-src': Also blacklist the source when action is taken; 'whitelist-src': Whitelist the source after filter passes, packets are dropped until then; 'count-only': Take no action and continue processing the next filter;
- `udp_filter_regex` (String) Regex Expression
- `udp_filter_unmatched` (Number) action taken when it does not match
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object


