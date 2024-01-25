---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cgnv6_nat_pool Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cgnv6_nat_pool: Configure CGNv6 NAT pool
  PLACEHOLDER
---

# thunder_cgnv6_nat_pool (Resource)

`thunder_cgnv6_nat_pool`: Configure CGNv6 NAT pool

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_nat_pool" "thunder_cgnv6_nat_pool" {
  end_address                            = "10.10.10.20"
  netmask                                = "/24"
  max_users_per_ip                       = 5
  partition                              = "test"
  per_batch_port_usage_warning_threshold = 34
  pool_name                              = "testPool"
  port_batch_v2_size                     = "64"
  shared                                 = 1
  simultaneous_batch_allocation          = 1
  start_address                          = "10.10.10.10"
  tcp_time_wait_interval                 = 5
  vrid                                   = 3
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `pool_name` (String) Specify pool name

### Optional

- `all` (Number) Share with all partitions
- `end_address` (String) Configure end IP address of NAT pool
- `exclude_ip` (Block List) (see [below for nested schema](#nestedblock--exclude_ip))
- `group` (String) Share with a partition group (Partition Group Name)
- `max_users_per_ip` (Number) Number of users that can be assigned to a NAT IP
- `netmask` (String) Configure mask for pool
- `partition` (String) Share with a single partition (Partition Name)
- `per_batch_port_usage_warning_threshold` (Number) Configure warning log threshold for per batch port usage (default: disabled) (Number of ports)
- `port_batch_v2_size` (String) '64': Allocate 64 ports at a time; '128': Allocate 128 ports at a time; '256': Allocate 256 ports at a time; '512': Allocate 512 ports at a time; '1024': Allocate 1024 ports at a time; '2048': Allocate 2048 ports at a time; '4096': Allocate 4096 ports at a time;
- `shared` (Number) Share this pool with other partitions (default: not shared)
- `simultaneous_batch_allocation` (Number) Allocate same TCP and UDP batches at once
- `start_address` (String) Configure start IP address of NAT pool
- `tcp_time_wait_interval` (Number) Minutes before TCP NAT ports can be reused
- `usable_nat_ports` (Number) Configure usable NAT ports
- `usable_nat_ports_end` (Number) End Port of Usable NAT Ports
- `usable_nat_ports_start` (Number) Start Port of Usable NAT Ports (needs to be even)
- `uuid` (String) uuid of the object
- `vrid` (Number) Configure VRRP-A vrid (Specify ha VRRP-A vrid)

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--exclude_ip"></a>
### Nested Schema for `exclude_ip`

Optional:

- `exclude_ip_end` (String) Address range end
- `exclude_ip_start` (String) Single IP address or IP address range start

