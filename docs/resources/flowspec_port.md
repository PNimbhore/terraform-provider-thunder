---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_flowspec_port Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_flowspec_port: Configure Port for a Flowspec
  PLACEHOLDER
---

# thunder_flowspec_port (Resource)

`thunder_flowspec_port`: Configure Port for a Flowspec

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_flowspec_port" "thunder_flowspec_port" {

  name           = "test"
  port_attribute = "eq"
  port_num       = 50
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name
- `port_attribute` (String) 'eq': Match only packets on a given port; 'gt': Match only packets with a greater port number; 'lt': Match only packets with a lower port number; 'range': match only packets in the range of port numbers;
- `port_num` (Number) Specify the port number

### Optional

- `port_num_end` (Number) Specify the port number
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


