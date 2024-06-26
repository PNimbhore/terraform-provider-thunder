---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_snmp_server_group Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_snmp_server_group: Define a User Security Model group
  PLACEHOLDER
---

# thunder_snmp_server_group (Resource)

`thunder_snmp_server_group`: Define a User Security Model group

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_group" "thunder_snmp_server_group" {

  groupname = "test-group"
  read      = "test-view"
  v3        = "noauth"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `groupname` (String) Name of the group

### Optional

- `read` (String) specify a read view for the group (read view name)
- `uuid` (String) uuid of the object
- `v3` (String) 'auth': group using the authNoPriv Security Level; 'noauth': group using the noAuthNoPriv Security Level; 'priv': group using SNMPv3 authPriv security level;

### Read-Only

- `id` (String) The ID of this resource.


