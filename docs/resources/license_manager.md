---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_license_manager Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_license_manager: Configure license manager
  PLACEHOLDER
---

# thunder_license_manager (Resource)

`thunder_license_manager`: Configure license manager

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_license_manager" "thunder_license_manager" {
  bandwidth_base         = 23443
  bandwidth_unrestricted = 1
  connect {
    connect = 0
  }
  host_list {
    host_ipv4 = "10.0.1.2"
    host_ipv6 = "7b5c:2272:a2a7:94c7:4004:8cbc:666f:06f1"
    port      = 443
  }
  instance_name = "1.4.7.6"
  interval      = 1
  overage {
    days    = 277
    hours   = 3
    minutes = 28
    seconds = 32
    gb      = 49686
    mb      = 219
    kb      = 646
    bytes   = 1019
  }
  reminder_list {
    reminder_value = 5885
  }
  sn = "22"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `bandwidth_base` (Number) Configure feature bandwidth base (Mb)
- `bandwidth_unrestricted` (Number) Set the bandwidth to maximum
- `connect` (Block List, Max: 1) (see [below for nested schema](#nestedblock--connect))
- `host_list` (Block List) (see [below for nested schema](#nestedblock--host_list))
- `instance_name` (String) Configure instance name [format: (string).(string).(string).(string)]
- `interval` (Number) Configure interval profile (1 monthly, 2 daily, 3 hourly)
- `ng_waf_module` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ng_waf_module))
- `overage` (Block List, Max: 1) (see [below for nested schema](#nestedblock--overage))
- `reminder_list` (Block List) (see [below for nested schema](#nestedblock--reminder_list))
- `sn` (String) serial number
- `use_mgmt_port` (Number) Use management port to connect license server
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--connect"></a>
### Nested Schema for `connect`

Optional:

- `connect` (Number) Connect to license manager to activate
- `uuid` (String) uuid of the object


<a id="nestedblock--host_list"></a>
### Nested Schema for `host_list`

Required:

- `host_ipv4` (String) license server ip address (length:1-31)
- `host_ipv6` (String) Configure license manager server ipv6-address

Optional:

- `port` (Number) Configure the license manager port, default is 443
- `uuid` (String) uuid of the object


<a id="nestedblock--ng_waf_module"></a>
### Nested Schema for `ng_waf_module`

Optional:

- `access_key_id` (String) access-key
- `secret_access_key` (String)


<a id="nestedblock--overage"></a>
### Nested Schema for `overage`

Optional:

- `bytes` (Number) Number of bytes
- `days` (Number) Number of days
- `gb` (Number) Number of GB
- `hours` (Number) Number of hours
- `kb` (Number) Number of KB
- `mb` (Number) Number of MB
- `minutes` (Number) Number of minutes
- `seconds` (Number) Number of seconds
- `uuid` (String) uuid of the object


<a id="nestedblock--reminder_list"></a>
### Nested Schema for `reminder_list`

Required:

- `reminder_value` (Number) Configure reminder for grace time (Hour)

Optional:

- `uuid` (String) uuid of the object


