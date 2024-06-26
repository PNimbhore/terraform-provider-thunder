---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_src_port_template_tcp Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_src_port_template_tcp: TCP template Configuration
  PLACEHOLDER
---

# thunder_ddos_src_port_template_tcp (Resource)

`thunder_ddos_src_port_template_tcp`: TCP template Configuration

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_src_port_template_tcp" "thunder_ddos_src_port_template_tcp" {

  filter_list {
    tcp_filter_seq       = 1
    tcp_filter_regex     = "1049"
    tcp_filter_unmatched = 1
    tcp_filter_action    = "blacklist-src"
    user_tag             = "89"
  }
  user_tag = "24"
  name     = "testing"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)

### Optional

- `filter_list` (Block List) (see [below for nested schema](#nestedblock--filter_list))
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--filter_list"></a>
### Nested Schema for `filter_list`

Required:

- `tcp_filter_seq` (Number) Sequence number

Optional:

- `byte_offset_filter` (String) Filter Expression using Berkeley Packet Filter syntax
- `tcp_filter_action` (String) 'blacklist-src': Also blacklist the source when action is taken; 'whitelist-src': Whitelist the source after filter passes, packets are dropped until then; 'count-only': Take no action and continue processing the next filter;
- `tcp_filter_regex` (String) Regex Expression
- `tcp_filter_unmatched` (Number) action taken when it does not match
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object


