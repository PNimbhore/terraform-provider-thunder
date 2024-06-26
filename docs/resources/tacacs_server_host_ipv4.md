---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_tacacs_server_host_ipv4 Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_tacacs_server_host_ipv4: Specify the hostname of TACACS+ server
  PLACEHOLDER
---

# thunder_tacacs_server_host_ipv4 (Resource)

`thunder_tacacs_server_host_ipv4`: Specify the hostname of TACACS+ server

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_tacacs_server_host_ipv4" "thunder_tacacs_server_host_ipv4" {
  ipv4_addr = "10.10.10.10"
  secret {
    source_eth   = 2
    secret_value = "test"
    port_cfg {
      port           = 22
      timeout        = 3
      monitor        = 1
      username       = "pramod"
      password       = 1
      password_value = "pramod"
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `ipv4_addr` (String) IPV4 address of TACACS+ server

### Optional

- `secret` (Block List, Max: 1) (see [below for nested schema](#nestedblock--secret))
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--secret"></a>
### Nested Schema for `secret`

Optional:

- `port_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--secret--port_cfg))
- `secret_value` (String) The TACACS+ server's secret
- `source_eth` (Number) Ethernet interface (Port number)
- `source_ip` (String) IP address
- `source_lif` (Number) Logical interface (Lif interface number)
- `source_loopback` (Number) Loopback interface (Port number)
- `source_trunk` (Number) Trunk interface (Trunk interface number)
- `source_ve` (Number) Virtual ethernet interface (Virtual ethernet interface number)

<a id="nestedblock--secret--port_cfg"></a>
### Nested Schema for `secret.port_cfg`

Optional:

- `monitor` (Number) Specify monitor TACACS+ server
- `password` (Number) Specify the user password
- `password_value` (String) The user password
- `port` (Number) Specify the port number used by TACACS+ server.( default port is 49) (Port number (default 49))
- `timeout` (Number) Specify the maximum time allowed for setting up a connection with the TACACS+ server. (default timeout is 12 seconds) (Maximum time allowed for setting up a connection with the TACACS+ server, in seconds (default 12))
- `username` (String) Specify the username


