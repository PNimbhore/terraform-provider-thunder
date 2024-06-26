---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_template_tcp_filter Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_template_tcp_filter: TCP Filter Configuration
  PLACEHOLDER
---

# thunder_ddos_template_tcp_filter (Resource)

`thunder_ddos_template_tcp_filter`: TCP Filter Configuration

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_template_tcp_filter" "thunder_ddos_template_tcp_filter" {
  name                 = "test"
  tcp_filter_action    = "blacklist-src"
  tcp_filter_regex     = "278"
  tcp_filter_seq       = 2
  tcp_filter_unmatched = 1
  user_tag             = "49"

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name
- `tcp_filter_seq` (Number) Sequence number

### Optional

- `byte_offset_filter` (String) Filter Expression using Berkeley Packet Filter syntax
- `tcp_filter_action` (String) 'blacklist-src': Also blacklist the source when action is taken; 'whitelist-src': Whitelist the source after filter passes, packets are dropped until then; 'count-only': Take no action and continue processing the next filter;
- `tcp_filter_regex` (String) Regex Expression
- `tcp_filter_unmatched` (Number) action taken when it does not match
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


