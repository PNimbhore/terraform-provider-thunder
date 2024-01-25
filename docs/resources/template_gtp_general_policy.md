---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_template_gtp_general_policy Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_template_gtp_general_policy: Configure GTP General Policy
  PLACEHOLDER
---

# thunder_template_gtp_general_policy (Resource)

`thunder_template_gtp_general_policy`: Configure GTP General Policy

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_template_gtp_general_policy" "thunder_template_gtp_general_policy" {
  name                   = "test"
  handover_timeout       = 2
  max_mesg_length_action = "drop"
  maximum_message_length = 1500
  tunnel_timeout         = 1440
  user_tag               = "100"
  v0_action              = "drop"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Specify name of the GTP General Policy

### Optional

- `handover_timeout` (Number) Tunnel Inactivity Timeout during Handover in minutes (default: 2 mins)
- `max_mesg_length_action` (String) 'monitor': Forward failed packet; 'drop': drop packet failing check(Default);
- `maximum_message_length` (Number) Maximum message length for a GTP message in bytes
- `tunnel_timeout` (Number) Tunnel Inactivity Timeout in minutes (default: 1440 minutes)
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object
- `v0_action` (String) 'permit': Permit GTP-C version 0; 'drop': Drop GTP-C version 0(Default);

### Read-Only

- `id` (String) The ID of this resource.

