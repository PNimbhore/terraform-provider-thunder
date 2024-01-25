---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_template_client_ssh Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_slb_template_client_ssh: Client Side SSH Template
  PLACEHOLDER
---

# thunder_slb_template_client_ssh (Resource)

`thunder_slb_template_client_ssh`: Client Side SSH Template

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_client_ssh" "Slb_Template_Client_Ssh_Test" {
  user_tag              = "ClientSSH"
  forward_proxy_enable  = 1
  passphrase            = "clientkey"
  forward_proxy_hostkey = "SLBCERT"
  name                  = "CLientSSH"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Client SSH Template Name

### Optional

- `forward_proxy_enable` (Number) Enable SSH forward proxy
- `forward_proxy_hostkey` (String) Specify private-key (Key Name)
- `passphrase` (String) Password Phrase
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

