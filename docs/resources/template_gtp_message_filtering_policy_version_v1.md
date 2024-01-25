---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_template_gtp_message_filtering_policy_version_v1 Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_template_gtp_message_filtering_policy_version_v1: Configure Message Filtering Policy for GTPv1 Control Messages,
  PLACEHOLDER
---

# thunder_template_gtp_message_filtering_policy_version_v1 (Resource)

`thunder_template_gtp_message_filtering_policy_version_v1`: Configure Message Filtering Policy for GTPv1 Control Messages,

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_template_gtp_message_filtering_policy_version_v1" "thunder_template_gtp_message_filtering_policy_version_v1" {

  name                  = "test"
  create_mbms           = "disable"
  create_pdp            = "enable"
  delete_mbms           = "disable"
  delete_pdp            = "enable"
  enable_disable_action = "enable"
  gtp_pdu               = "enable"
  initiate_pdp          = "enable"
  mbms_deregistration   = "disable"
  mbms_notification     = "disable"
  mbms_registration     = "disable"
  mbms_session          = "disable"
  ms_info_change        = "enable"
  pdu_notification      = "enable"
  reserved_messages     = "disable"
  update_mbms           = "disable"
  update_pdp            = "enable"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `enable_disable_action` (String) 'enable': Enable Message Filtering on version; 'disable': Disable Message Filtering on version;
- `name` (String) Name

### Optional

- `create_mbms` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `create_pdp` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `delete_mbms` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `delete_pdp` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `gtp_pdu` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `initiate_pdp` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `mbms_deregistration` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `mbms_notification` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `mbms_registration` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `mbms_session` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `message_type` (Number) Specify the Message Type
- `ms_info_change` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `pdu_notification` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `reserved_messages` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `update_mbms` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `update_pdp` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

