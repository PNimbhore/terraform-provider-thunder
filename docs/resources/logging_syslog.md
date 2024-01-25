---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_logging_syslog Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_logging_syslog: Set logging level which sent to syslog host
  PLACEHOLDER
---

# thunder_logging_syslog (Resource)

`thunder_logging_syslog`: Set logging level which sent to syslog host

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_logging_syslog" "thunder_logging_syslog" {
  syslog_levelname = "alert"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `syslog_levelname` (String) 'disable': Do not send log to syslog; 'emergency': System unusable log messages      (severity=0); 'alert': Action must be taken immediately  (severity=1); 'critical': Critical conditions               (severity=2); 'error': Error conditions                  (severity=3); 'warning': Warning conditions                (severity=4); 'notification': Normal but significant conditions (severity=5); 'information': Informational messages            (severity=6); 'debugging': Debug level messages              (severity=7);
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

