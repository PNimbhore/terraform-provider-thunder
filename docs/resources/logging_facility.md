---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_logging_facility Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_logging_facility: Facility parameter for syslog messages
  PLACEHOLDER
---

# thunder_logging_facility (Resource)

`thunder_logging_facility`: Facility parameter for syslog messages

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_logging_facility" "thunder_logging_facility" {
  facilityname = "local1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `facilityname` (String) 'local0': Local use; 'local1': Local use; 'local2': Local use; 'local3': Local use; 'local4': Local use; 'local5': Local use; 'local6': Local use; 'local7': Local use;  (Facility parameter for syslog messages)
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


