---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_health_monitor_method_smtp Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_health_monitor_method_smtp: SMTP type
  PLACEHOLDER
---

# thunder_health_monitor_method_smtp (Resource)

`thunder_health_monitor_method_smtp`: SMTP type

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_health_monitor_method_smtp" "thunder_health_monitor_method_smtp" {

  name          = "tf_test"
  mail_from     = "93"
  rcpt_to       = "116"
  smtp          = 1
  smtp_domain   = "35"
  smtp_port     = 251
  smtp_starttls = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name

### Optional

- `mail_from` (String) Specify SMTP Sender
- `rcpt_to` (String) Specify SMTP Receiver
- `smtp` (Number) SMTP type
- `smtp_domain` (String) Specify domain name of 'helo' command
- `smtp_port` (Number) Specify SMTP port, default is 25 (Port Number (default 25))
- `smtp_starttls` (Number) Check the STARTTLS support at helo response
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


