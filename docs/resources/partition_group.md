---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_partition_group Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_partition_group: Modify a partition group
  PLACEHOLDER
---

# thunder_partition_group (Resource)

`thunder_partition_group`: Modify a partition group

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_partition_group" "thunder_partition_group" {

  member_list {
    member = "test"
  }
  partition_group_name = "test"
  user_tag             = "test_part_group"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `partition_group_name` (String) Partition Group Name

### Optional

- `member_list` (Block List) (see [below for nested schema](#nestedblock--member_list))
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--member_list"></a>
### Nested Schema for `member_list`

Optional:

- `member` (String) Partition Name


