---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_file_bw_list_local Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_file_bw_list_local: black white list file information and management commands
  PLACEHOLDER
---

# thunder_file_bw_list_local (Resource)

`thunder_file_bw_list_local`: black white list file information and management commands

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_file_bw_list_local" "thunder_file_bw_list_local" {
  action      = "import"
  dst_file    = "test_123"
  file        = "test_1"
  file_handle = "/mnt/c/Users/PNimbhore/WorkSpace Terraform/sample-certificates/class-list-ac1-a10"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `action` (String) 'check': check; 'create': create; 'delete': delete; 'export': export; 'import': import; 'replace': replace;
- `dst_file` (String) destination file name for copy and rename action
- `file` (String) bw-list file name
- `file_handle` (String) full path of the uploaded file
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.


