---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cgnv6_logging_nat_quota_exceeded Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cgnv6_logging_nat_quota_exceeded: Change logging level for NAT Quota Exceeded
  PLACEHOLDER
---

# thunder_cgnv6_logging_nat_quota_exceeded (Resource)

`thunder_cgnv6_logging_nat_quota_exceeded`: Change logging level for NAT Quota Exceeded

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_logging_nat_quota_exceeded" "thunder_cgnv6_logging_nat_quota_exceeded" {
  level = "critical"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `level` (String) 'warning': Log level Warning (Default); 'critical': Log level Critical; 'notice': Log level Notice;
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


