---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_acos_cloud_integration Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_acos_cloud_integration: Setting for heterogeneous cloud environment integration
  PLACEHOLDER
---

# thunder_acos_cloud_integration (Resource)

`thunder_acos_cloud_integration`: Setting for heterogeneous cloud environment integration

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_acos_cloud_integration" "thunder_acos_cloud_integration" {
  dummy = 0
  ecosystem {
    dummy = 0
    consul {
      host_name = "125"
      port      = 62058
      action    = "disable"
    }
    oracle {
      host_name        = "18"
      port             = 20567
      compartment_id   = "73"
      tenancy_id       = "80"
      user_id          = "56"
      fingerprint      = "61"
      private_key_path = "96"
      action           = "disable"
    }
    k8s {
      action = "disable"
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `dummy` (Number) dummy to make intermediate obj to single
- `ecosystem` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ecosystem))
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--ecosystem"></a>
### Nested Schema for `ecosystem`

Optional:

- `consul` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ecosystem--consul))
- `dummy` (Number) dummy to make intermediate obj to single
- `k8s` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ecosystem--k8s))
- `oracle` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ecosystem--oracle))
- `uuid` (String) uuid of the object

<a id="nestedblock--ecosystem--consul"></a>
### Nested Schema for `ecosystem.consul`

Optional:

- `action` (String) 'enable': Enable Configuration; 'disable': Disable Configuration;
- `health_check_interval` (String) '5': 5 seconds; '10': 10 seconds; '15': 15 seconds; '20': 20 seconds;
- `host_name` (String) Configure the host name for bootstrap server(e.g www.a10networks.com)
- `ipv4_address` (String) Configure the bootstrap server's IPv4 address (the host IPv4 address)
- `ipv6_address` (String) Configure the bootstrap server's IPv6 address (the host IPv6 address)
- `port` (Number) Configure the http port to use (port 8500)
- `service_label` (Block List) (see [below for nested schema](#nestedblock--ecosystem--consul--service_label))
- `uuid` (String) uuid of the object

<a id="nestedblock--ecosystem--consul--service_label"></a>
### Nested Schema for `ecosystem.consul.service_label`

Optional:

- `service_label_name` (String) Name service group to be monitored



<a id="nestedblock--ecosystem--k8s"></a>
### Nested Schema for `ecosystem.k8s`

Optional:

- `action` (String) 'enable': Enable Configuration; 'disable': Disable Configuration;
- `cluster_config_file` (String) Enter cluster config file name
- `health_check_interval` (String) '5': 5 seconds; '10': 10 seconds; '15': 15 seconds; '20': 20 seconds;
- `service_label` (Block List) (see [below for nested schema](#nestedblock--ecosystem--k8s--service_label))
- `uuid` (String) uuid of the object

<a id="nestedblock--ecosystem--k8s--service_label"></a>
### Nested Schema for `ecosystem.k8s.service_label`

Optional:

- `service_label_name` (String) Name service group to be monitored



<a id="nestedblock--ecosystem--oracle"></a>
### Nested Schema for `ecosystem.oracle`

Optional:

- `action` (String) 'enable': Enable Configuration; 'disable': Disable Configuration;
- `compartment_id` (String) OCI compartment  id
- `fingerprint` (String)
- `health_check_interval` (String) '5': 5 seconds; '10': 10 seconds; '15': 15 seconds; '20': 20 seconds;
- `host_name` (String) Configure the host name for bootstrap server(e.g www.a10networks.com)
- `ipv4_address` (String) Configure the bootstrap server's IPv4 address (the host IPv4 address)
- `ipv6_address` (String) Configure the bootstrap server's IPv6 address (the host IPv6 address)
- `port` (Number) Configure the http port to use (port 8500)
- `private_key_path` (String)
- `service_label` (Block List) (see [below for nested schema](#nestedblock--ecosystem--oracle--service_label))
- `tenancy_id` (String) OCI tenancy  id
- `user_id` (String) OCI user id
- `uuid` (String) uuid of the object

<a id="nestedblock--ecosystem--oracle--service_label"></a>
### Nested Schema for `ecosystem.oracle.service_label`

Optional:

- `service_label_name` (String) Name service group to be monitored


