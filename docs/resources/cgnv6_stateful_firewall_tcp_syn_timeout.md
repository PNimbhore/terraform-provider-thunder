---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cgnv6_stateful_firewall_tcp_syn_timeout Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cgnv6_stateful_firewall_tcp_syn_timeout: Configure TCP SYNtimeout
  PLACEHOLDER
---

# thunder_cgnv6_stateful_firewall_tcp_syn_timeout (Resource)

`thunder_cgnv6_stateful_firewall_tcp_syn_timeout`: Configure TCP SYNtimeout

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_stateful_firewall_tcp_syn_timeout" "thunder_cgnv6_stateful_firewall_tcp_syn_timeout" {

  syn_timeout_val = 5

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `syn_timeout_val` (Number) Set Seconds session can remain in half-open state before being deleted (default: 4 seconds)
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


