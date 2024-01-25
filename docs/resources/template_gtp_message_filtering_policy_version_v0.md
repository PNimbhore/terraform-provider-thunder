---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_template_gtp_message_filtering_policy_version_v0 Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_template_gtp_message_filtering_policy_version_v0: Configure Message Filtering Policy for GTPv0 Control Messages,
  PLACEHOLDER
---

# thunder_template_gtp_message_filtering_policy_version_v0 (Resource)

`thunder_template_gtp_message_filtering_policy_version_v0`: Configure Message Filtering Policy for GTPv0 Control Messages,

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_template_gtp_message_filtering_policy_version_v0" "thunder_template_gtp_message_filtering_policy_version_v0" {

  name                  = "test"
  create_aa_pdp         = "enable"
  create_pdp            = "enable"
  delete_aa_pdp         = "enable"
  delete_pdp            = "enable"
  enable_disable_action = "enable"
  gtp_pdu               = "enable"
  pdu_notification      = "enable"
  reserved_messages     = "disable"
  update_pdp            = "enable"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `enable_disable_action` (String) 'enable': Enable Message Filtering on version; 'disable': Disable Message Filtering on version;
- `name` (String) Name

### Optional

- `create_aa_pdp` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `create_pdp` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `delete_aa_pdp` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `delete_pdp` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `gtp_pdu` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `message_type` (Number) Specify the Message Type
- `pdu_notification` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `reserved_messages` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `update_pdp` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

