---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_acos_events_message_selector_rule Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_acos_events_message_selector_rule: Configure rules to select messages for which logging is enabled/blocked
  PLACEHOLDER
---

# thunder_acos_events_message_selector_rule (Resource)

`thunder_acos_events_message_selector_rule`: Configure rules to select messages for which logging is enabled/blocked

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_acos_events_message_selector_rule" "thunder_acos_events_message_selector_rule" {

  name          = "test"
  action        = "send"
  index         = 26
  severity_oper = "equal-and-higher"
  severity_val  = "emergency"
  user_tag      = "116"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `index` (Number) Specify rule index - rules are applied in numeric order
- `name` (String) Name

### Optional

- `action` (String) 'send': log messages selected by this rule will be sent; 'drop': log messages selected by this rule will be dropped;
- `message_id` (String) Select a specific message by message-id and optionally severity
- `message_id_scope` (String) 'all': Log messages at this level and all sub-trees; 'node-only': Log messages at this node only; 'children-only': Log messages at all sub-trees only; 'log-field-only': Log message for this Log Field only;
- `severity_oper` (String) 'equal-and-higher': emergency is highest, debugging lowest; 'equal': single severity;
- `severity_val` (String) 'emergency': System unusable log messages (Most Important); 'alert': Action must be taken immediately; 'critical': Critical conditions; 'error': Error conditions; 'warning': Warning conditions; 'notification': Normal but significant conditions; 'information': Informational messages; 'debugging': Debug level messages (Least Important);
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


