---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_fw_clear_session_filter Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_fw_clear_session_filter: Enable clear L4 session filter for fw (Default: Disable)
  PLACEHOLDER
---

# thunder_fw_clear_session_filter (Resource)

`thunder_fw_clear_session_filter`: Enable clear L4 session filter for fw (Default: Disable)

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_fw_clear_session_filter" "test_thunder_fw_clear_session_filter" {
  status = "enable"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `status` (String) 'disable': Disable clear L4 session filter for fw (Default: disabled); 'enable': Enable clear L4 session filter for fw;
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


