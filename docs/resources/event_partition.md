---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_event_partition Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_event_partition: module partition
  PLACEHOLDER
---

# thunder_event_partition (Resource)

`thunder_event_partition`: module partition

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_event_partition" "thunder_event_partition" {
  email      = "on"
  logging    = "on"
  user_tag   = "87"
  vnp_events = "part-create"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `vnp_events` (String) 'part-create': Create new partition; 'part-del': Delete a partition;

### Optional

- `email` (String) 'on': enable this action; 'off': disable this action;
- `logging` (String) 'on': enable this action; 'off': disable this action;
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

