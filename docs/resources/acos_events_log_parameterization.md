---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_acos_events_log_parameterization Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_acos_events_log_parameterization: Harmony controller log parameterization configuration
  PLACEHOLDER
---

# thunder_acos_events_log_parameterization (Resource)

`thunder_acos_events_log_parameterization`: Harmony controller log parameterization configuration

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_acos_events_log_parameterization" "thunder_acos_events_log_parameterization" {
  log_rate = 10
  message_selector {
    rule_list {
      index         = 88
      action        = "drop"
      severity_oper = "equal"
      severity_val  = "alert"
      user_tag      = "9"
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `log_rate` (Number) Max number of parameterized logs sent per second
- `message_selector` (Block List, Max: 1) (see [below for nested schema](#nestedblock--message_selector))
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--message_selector"></a>
### Nested Schema for `message_selector`

Optional:

- `rule_list` (Block List) (see [below for nested schema](#nestedblock--message_selector--rule_list))
- `uuid` (String) uuid of the object

<a id="nestedblock--message_selector--rule_list"></a>
### Nested Schema for `message_selector.rule_list`

Required:

- `index` (Number) Specify rule index - rules are applied in numeric order

Optional:

- `action` (String) 'send': log messages selected by this rule will be sent (Default); 'drop': log messages selected by this rule will be dropped;
- `message_id` (String) Select a specific message by message-id and optionally severity
- `message_id_scope` (String) 'all': Log messages at this level and all sub-trees; 'node-only': Log messages at this node only; 'children-only': Log messages at all sub-trees only; 'log-field-only': Log message for this Log Field only;
- `severity_oper` (String) 'equal-and-higher': emergency is highest, debugging lowest; 'equal': single severity;
- `severity_val` (String) 'emergency': System unusable log messages (Most Important); 'alert': Action must be taken immediately; 'critical': Critical conditions; 'error': Error conditions; 'warning': Warning conditions; 'notification': Normal but significant conditions; 'information': Informational messages; 'debugging': Debug level messages (Least Important);
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

