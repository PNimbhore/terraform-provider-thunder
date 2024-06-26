---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_file_license_local Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_file_license_local: license file information and management commands
  PLACEHOLDER
---

# thunder_file_license_local (Resource)

`thunder_file_license_local`: license file information and management commands

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_file_license_local" "thunder_file_license_local" {
  action      = "import"
  dst_file    = "test"
  file        = "test"
  file_handle = "/mnt/c/Users/PNimbhore/WorkSpace Terraform/sample-certificates/glm.txt"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `action` (String) 'create': create; 'import': import; 'export': export; 'copy': copy; 'rename': rename; 'check': check; 'replace': replace; 'delete': delete;
- `device` (Number) Device (Device ID)
- `dst_file` (String) destination file name for copy and rename action
- `file` (String) license local file name
- `file_handle` (String) full path of the uploaded file
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


