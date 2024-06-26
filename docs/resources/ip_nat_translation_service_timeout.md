---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ip_nat_translation_service_timeout Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ip_nat_translation_service_timeout: Specify any custom service timeout
  PLACEHOLDER
---

# thunder_ip_nat_translation_service_timeout (Resource)

`thunder_ip_nat_translation_service_timeout`: Specify any custom service timeout

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_nat_translation_service_timeout" "thunder_ip_nat_translation_service_timeout" {
  port         = 37642
  service_type = "tcp"
  timeout_type = "age"
  timeout_val  = 13303
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `port` (Number) Port Number
- `service_type` (String) 'tcp': TCP Protocol; 'udp': UDP Protocol;

### Optional

- `timeout_type` (String) 'age': Expiration time; 'fast': Use Fast aging;
- `timeout_val` (Number) Timeout in seconds (Interval of 60 seconds)
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


