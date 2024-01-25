---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cloud_services_cloud_provider_aws_log Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cloud_services_cloud_provider_aws_log: AWS log configuration
  PLACEHOLDER
---

# thunder_cloud_services_cloud_provider_aws_log (Resource)

`thunder_cloud_services_cloud_provider_aws_log`: AWS log configuration

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cloud_services_cloud_provider_aws_log" "test" {
  action         = "enable"
  log_group_name = "testing"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `action` (String) 'enable': Enable AWS CloudWatch; 'disable': Disable AWS CloudWatch (default);
- `active_partitions` (String) Specifies the thunder active partition name separated by a comma for multiple values
- `log_group_name` (String) Specifies the log group name under which all logs are sent to AWS CloudWatch
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

