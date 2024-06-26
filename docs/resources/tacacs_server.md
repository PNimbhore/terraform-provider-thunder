---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_tacacs_server Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_tacacs_server: Configure TACACS+ servers
  PLACEHOLDER
---

# thunder_tacacs_server (Resource)

`thunder_tacacs_server`: Configure TACACS+ servers

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_tacacs_server" "thunder_tacacs_server" {
  monitor  = 1
  interval = 111
  host {
    ipv4_list {
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

    ipv6_list {
      ipv6_addr = "2001:db8:3333:4444:5555:6666:7777:8888"
      secret {
        source_eth   = 2
        secret_value = "test"
        port_cfg {
          port           = 22
          timeout        = 11
          monitor        = 1
          username       = "suraj123"
          password       = 1
          password_value = "suraj"
        }
      }
    }

    tacacs_hostname_list {
      hostname = "test"
      secret {
        source_eth   = 2
        secret_value = "test"
        port_cfg {
          port           = 22
          timeout        = 2
          monitor        = 1
          username       = "suraj"
          password       = 1
          password_value = "suraj"
        }
      }
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `host` (Block List, Max: 1) (see [below for nested schema](#nestedblock--host))
- `interval` (Number) The moniter interval in seconds (default 60)
- `monitor` (Number) Configure TACACS+ servers
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--host"></a>
### Nested Schema for `host`

Optional:

- `ipv4_list` (Block List) (see [below for nested schema](#nestedblock--host--ipv4_list))
- `ipv6_list` (Block List) (see [below for nested schema](#nestedblock--host--ipv6_list))
- `tacacs_hostname_list` (Block List) (see [below for nested schema](#nestedblock--host--tacacs_hostname_list))

<a id="nestedblock--host--ipv4_list"></a>
### Nested Schema for `host.ipv4_list`

Required:

- `ipv4_addr` (String) IPV4 address of TACACS+ server

Optional:

- `secret` (Block List, Max: 1) (see [below for nested schema](#nestedblock--host--ipv4_list--secret))
- `uuid` (String) uuid of the object

<a id="nestedblock--host--ipv4_list--secret"></a>
### Nested Schema for `host.ipv4_list.secret`

Optional:

- `port_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--host--ipv4_list--secret--port_cfg))
- `secret_value` (String) The TACACS+ server's secret
- `source_eth` (Number) Ethernet interface (Port number)
- `source_ip` (String) IP address
- `source_lif` (Number) Logical interface (Lif interface number)
- `source_loopback` (Number) Loopback interface (Port number)
- `source_trunk` (Number) Trunk interface (Trunk interface number)
- `source_ve` (Number) Virtual ethernet interface (Virtual ethernet interface number)

<a id="nestedblock--host--ipv4_list--secret--port_cfg"></a>
### Nested Schema for `host.ipv4_list.secret.port_cfg`

Optional:

- `monitor` (Number) Specify monitor TACACS+ server
- `password` (Number) Specify the user password
- `password_value` (String) The user password
- `port` (Number) Specify the port number used by TACACS+ server.( default port is 49) (Port number (default 49))
- `timeout` (Number) Specify the maximum time allowed for setting up a connection with the TACACS+ server. (default timeout is 12 seconds) (Maximum time allowed for setting up a connection with the TACACS+ server, in seconds (default 12))
- `username` (String) Specify the username




<a id="nestedblock--host--ipv6_list"></a>
### Nested Schema for `host.ipv6_list`

Required:

- `ipv6_addr` (String) IPV6 address of TACACS+ server

Optional:

- `secret` (Block List, Max: 1) (see [below for nested schema](#nestedblock--host--ipv6_list--secret))
- `uuid` (String) uuid of the object

<a id="nestedblock--host--ipv6_list--secret"></a>
### Nested Schema for `host.ipv6_list.secret`

Optional:

- `port_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--host--ipv6_list--secret--port_cfg))
- `secret_value` (String) The TACACS+ server's secret
- `source_eth` (Number) Ethernet interface (Port number)
- `source_ipv6` (String) IPV6 address
- `source_lif` (Number) Logical interface (Lif interface number)
- `source_loopback` (Number) Loopback interface (Port number)
- `source_trunk` (Number) Trunk interface (Trunk interface number)
- `source_ve` (Number) Virtual ethernet interface (Virtual ethernet interface number)

<a id="nestedblock--host--ipv6_list--secret--port_cfg"></a>
### Nested Schema for `host.ipv6_list.secret.port_cfg`

Optional:

- `monitor` (Number) Specify monitor TACACS+ server
- `password` (Number) Specify the user password
- `password_value` (String) The user password
- `port` (Number) Specify the port number used by TACACS+ server.( default port is 49) (Port number (default 49))
- `timeout` (Number) Specify the maximum time allowed for setting up a connection with the TACACS+ server. (default timeout is 12 seconds) (Maximum time allowed for setting up a connection with the TACACS+ server, in seconds (default 12))
- `username` (String) Specify the username




<a id="nestedblock--host--tacacs_hostname_list"></a>
### Nested Schema for `host.tacacs_hostname_list`

Required:

- `hostname` (String) Hostname of TACACS+ server

Optional:

- `secret` (Block List, Max: 1) (see [below for nested schema](#nestedblock--host--tacacs_hostname_list--secret))
- `uuid` (String) uuid of the object

<a id="nestedblock--host--tacacs_hostname_list--secret"></a>
### Nested Schema for `host.tacacs_hostname_list.secret`

Optional:

- `port_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--host--tacacs_hostname_list--secret--port_cfg))
- `secret_value` (String) The TACACS+ server's secret
- `source_eth` (Number) Ethernet interface (Port number)
- `source_ip` (String) IP address
- `source_ipv6` (String) IPV6 address
- `source_lif` (Number) Logical interface (Lif interface number)
- `source_loopback` (Number) Loopback interface (Port number)
- `source_trunk` (Number) Trunk interface (Trunk interface number)
- `source_ve` (Number) Virtual ethernet interface (Virtual ethernet interface number)

<a id="nestedblock--host--tacacs_hostname_list--secret--port_cfg"></a>
### Nested Schema for `host.tacacs_hostname_list.secret.port_cfg`

Optional:

- `monitor` (Number) Specify monitor TACACS+ server
- `password` (Number) Specify the user password
- `password_value` (String) The user password
- `port` (Number) Specify the port number used by TACACS+ server.( default port is 49) (Port number (default 49))
- `timeout` (Number) Specify the maximum time allowed for setting up a connection with the TACACS+ server. (default timeout is 12 seconds) (Maximum time allowed for setting up a connection with the TACACS+ server, in seconds (default 12))
- `username` (String) Specify the username


