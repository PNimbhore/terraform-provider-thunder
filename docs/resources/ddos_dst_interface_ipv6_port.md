---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_dst_interface_ipv6_port Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_dst_interface_ipv6_port: Configure port for local interface IPs
  PLACEHOLDER
---

# thunder_ddos_dst_interface_ipv6_port (Resource)

`thunder_ddos_dst_interface_ipv6_port`: Configure port for local interface IPs

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_interface_ipv6_port" "thunder_ddos_dst_interface_ipv6_port" {
  addr     = "2001:db8:3333:4444:5555:6666:7777:8888"
  deny     = 1
  port_num = 75
  protocol = "tcp"
  user_tag = "4"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `addr` (String) Addr
- `port_num` (Number) Port Number
- `protocol` (String) 'tcp': tcp; 'udp': udp; 'http-probe': http port for interface health check;

### Optional

- `deny` (Number) Blacklist and Drop all incoming packets for protocol
- `glid` (String) Global limit ID
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


