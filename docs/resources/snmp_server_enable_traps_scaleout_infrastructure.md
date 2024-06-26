---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_snmp_server_enable_traps_scaleout_infrastructure Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_snmp_server_enable_traps_scaleout_infrastructure: Enable Intrastructure group traps
  PLACEHOLDER
---

# thunder_snmp_server_enable_traps_scaleout_infrastructure (Resource)

`thunder_snmp_server_enable_traps_scaleout_infrastructure`: Enable Intrastructure group traps

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_enable_traps_scaleout_infrastructure" "thunder_snmp_server_enable_traps_scaleout_infrastructure" {
  all = 0
  cluster {
    election                   = 0
    master_calling_re_election = 0
    node_status                = 0
  }
  master_node {
    traffic_map_distribution   = 0
    vserver_traffic_map_update = 0
  }
  service_node {
    local_device_disabled = 0
    service_master        = 0
    traffic_map_update    = 0
  }
  test_send_all_traps = 0
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `all` (Number) Enable all infra traps
- `cluster` (Block List, Max: 1) (see [below for nested schema](#nestedblock--cluster))
- `master_node` (Block List, Max: 1) (see [below for nested schema](#nestedblock--master_node))
- `service_node` (Block List, Max: 1) (see [below for nested schema](#nestedblock--service_node))
- `test_send_all_traps` (Number) Send all infra traps
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--cluster"></a>
### Nested Schema for `cluster`

Optional:

- `election` (Number) Enable election status trap
- `master_calling_re_election` (Number) Enable re-election trap
- `node_status` (Number) Enable active node status trap
- `uuid` (String) uuid of the object


<a id="nestedblock--master_node"></a>
### Nested Schema for `master_node`

Optional:

- `traffic_map_distribution` (Number) Enable Traffic-map distribution trap
- `uuid` (String) uuid of the object
- `vserver_traffic_map_update` (Number) Enable VServer Traffic-map trap


<a id="nestedblock--service_node"></a>
### Nested Schema for `service_node`

Optional:

- `local_device_disabled` (Number) Enable local device disabled trap
- `service_master` (Number) Enable service-master trap
- `traffic_map_update` (Number) Enable traffic map update trap
- `uuid` (String) uuid of the object


