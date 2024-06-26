---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_pattern_recognition_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_pattern_recognition_stats: Statistics for the object pattern-recognition
  PLACEHOLDER
---

# thunder_ddos_pattern_recognition_stats (Data Source)

`thunder_ddos_pattern_recognition_stats`: Statistics for the object pattern-recognition

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_pattern_recognition_stats" "thunder_ddos_pattern_recognition_stats" {
}
output "get_ddos_pattern_recognition_stats" {
  value = ["${data.thunder_ddos_pattern_recognition_stats.thunder_ddos_pattern_recognition_stats}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--stats))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--stats"></a>
### Nested Schema for `stats`

Optional:

- `filter_drop` (Number) Extracted Filter Drop
- `filter_match` (Number) Extracted Filter Match
- `generic_error` (Number) Pattern Recognition: Exceptions
- `not_found` (Number) Pattern Recognition: Pattern Not Found
- `proceeded` (Number) Pattern Recognition: Engine Started


