---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_snmp_server_host_ipv4_host Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_snmp_server_host_ipv4_host: Specify IPV4 hosts to receive SNMP traps
  PLACEHOLDER
---

# thunder_snmp_server_host_ipv4_host (Resource)

`thunder_snmp_server_host_ipv4_host`: Specify IPV4 hosts to receive SNMP traps

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_host_ipv4_host" "thunder_snmp_server_host_ipv4_host" {
  ipv4_addr   = "10.10.10.10"
  udp_port    = 162
  v1_v2c_comm = "20"
  version     = "v1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `ipv4_addr` (String) IPV4 address of SNMP trap host
- `version` (String) 'v1': Use SNMPv1; 'v2c': Use SNMPv2c; 'v3': User SNMPv3;

### Optional

- `udp_port` (Number) The trap host's UDP port number(default: 162)
- `user` (String) SNMPv3 user to send traps (User Name)
- `uuid` (String) uuid of the object
- `v1_v2c_comm` (String) SNMPv1/v2c community string

### Read-Only

- `id` (String) The ID of this resource.


