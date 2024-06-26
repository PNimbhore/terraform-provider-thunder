---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_locale Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_locale: Set locale for the CLI startup
  PLACEHOLDER
---

# thunder_locale (Resource)

`thunder_locale`: Set locale for the CLI startup

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_locale" "thunder_locale" {
  value = "zh_CN.UTF-8"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `test` (Block List, Max: 1) (see [below for nested schema](#nestedblock--test))
- `uuid` (String) uuid of the object
- `value` (String) 'en_US.UTF-8': English locale for the USA, encoding with UTF-8 (default); 'zh_CN.UTF-8': Chinese locale for PRC, encoding with UTF-8; 'zh_CN.GB18030': Chinese locale for PRC, encoding with GB18030; 'zh_CN.GBK': Chinese locale for PRC, encoding with GBK; 'zh_CN.GB2312': Chinese locale for PRC, encoding with GB2312; 'zh_TW.UTF-8': Chinese locale for Taiwan, encoding with UTF-8; 'zh_TW.BIG5': Chinese locale for Taiwan, encoding with BIG5; 'zh_TW.EUCTW': Chinese locale for Taiwan, encoding with EUC-TW; 'ja_JP.UTF-8': Japanese locale for Japan, encoding with UTF-8; 'ja_JP.EUC-JP': Japanese locale for Japan, encoding with EUC-JP;

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--test"></a>
### Nested Schema for `test`

Optional:

- `locale` (String) 'zh_CN': Chinese locale for PRC; 'zh_TW': Chinese locale for Taiwan; 'ja_JP': Japanese locale for Japan;


