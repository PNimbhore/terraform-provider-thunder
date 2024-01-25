---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_template_gtp_message_filtering_policy_version_v2 Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_template_gtp_message_filtering_policy_version_v2: Configure Message Filtering Policy for GTPv2 Control Messages
  PLACEHOLDER
---

# thunder_template_gtp_message_filtering_policy_version_v2 (Resource)

`thunder_template_gtp_message_filtering_policy_version_v2`: Configure Message Filtering Policy for GTPv2 Control Messages

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_template_gtp_message_filtering_policy_version_v2" "thunder_template_gtp_message_filtering_policy_version_v2" {

  name                  = "test"
  bearer_resource       = "enable"
  change_notification   = "enable"
  create_bearer         = "enable"
  create_session        = "enable"
  delete_bearer         = "enable"
  delete_command        = "enable"
  delete_pdn            = "enable"
  delete_session        = "enable"
  enable_disable_action = "enable"
  modify_bearer         = "enable"
  modify_command        = "enable"
  pgw_downlink_trigger  = "disable"
  remote_ue_report      = "enable"
  reserved_messages     = "disable"
  resume                = "enable"
  suspend               = "enable"
  trace_session         = "disable"
  update_bearer         = "enable"
  update_pdn            = "enable"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `enable_disable_action` (String) 'enable': Enable Message Filtering on version; 'disable': Disable Message Filtering on version;
- `name` (String) Name

### Optional

- `bearer_resource` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `change_notification` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `create_bearer` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `create_session` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `delete_bearer` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `delete_command` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `delete_pdn` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `delete_session` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `message_type` (Number) Specify the Message Type
- `modify_bearer` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `modify_command` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `pgw_downlink_trigger` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `remote_ue_report` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `reserved_messages` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `resume` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `suspend` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `trace_session` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `update_bearer` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `update_pdn` (String) 'enable': Enable the Message Type; 'disable': Disable the Message Type;
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

