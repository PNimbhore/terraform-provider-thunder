---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_import_store Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_import_store: Create store name for remote url
  PLACEHOLDER
---

# thunder_import_store (Resource)

`thunder_import_store`: Create store name for remote url

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_import_store" "thunder_import_store" {
  create      = 1
  delete      = 0
  name        = "test"
  remote_file = "test"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `create` (Number) Create an import store profile
- `delete` (Number) Delete an import store profile
- `name` (String) profile name to store remote url
- `remote_file` (String)

### Read-Only

- `id` (String) The ID of this resource.


