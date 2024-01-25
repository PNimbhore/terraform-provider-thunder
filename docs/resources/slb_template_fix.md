---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_template_fix Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_slb_template_fix: FIX template
  PLACEHOLDER
---

# thunder_slb_template_fix (Resource)

`thunder_slb_template_fix`: FIX template

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_fix" "test_thunder_slb_template_fix" {
  name             = "test_fix"
  logging          = "init"
  insert_client_ip = 1
  tag_switching {
    switching_type = "sender-comp-id"
    equals         = "test"
    service_group  = "sg1"
  }
  user_tag = "test"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) FIX Template Name

### Optional

- `insert_client_ip` (Number) Insert client ip to tag 11447
- `logging` (String) 'init': init only log; 'term': termination only log; 'both': both initial and termination log;
- `tag_switching` (Block List) (see [below for nested schema](#nestedblock--tag_switching))
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--tag_switching"></a>
### Nested Schema for `tag_switching`

Optional:

- `equals` (String) Equals (Tag String)
- `service_group` (String) Create a Service Group comprising Servers (Service Group Name)
- `switching_type` (String) 'sender-comp-id': Select service group based on SenderCompID; 'target-comp-id': Select service group based on TargetCompID;

