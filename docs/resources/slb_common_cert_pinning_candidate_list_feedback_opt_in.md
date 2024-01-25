---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_common_cert_pinning_candidate_list_feedback_opt_in Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_slb_common_cert_pinning_candidate_list_feedback_opt_in: Opt-in to improve A10’s certificate-pinning sites list
  PLACEHOLDER
---

# thunder_slb_common_cert_pinning_candidate_list_feedback_opt_in (Resource)

`thunder_slb_common_cert_pinning_candidate_list_feedback_opt_in`: Opt-in to improve A10’s certificate-pinning sites list

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_common_cert_pinning_candidate_list_feedback_opt_in" "thunder_slb_common_cert_pinning_candidate_list_feedback_opt_in" {
  daily    = 1
  enable   = 1
  schedule = 1
  day_time = "12:45"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `daily` (Number) Every day
- `day_time` (String) Time of day to update (hh:mm) in 24 hour local time
- `enable` (Number) Enable the feedback function
- `schedule` (Number) schedule the uploading time, default is daily 00:00
- `use_mgmt_port` (Number) Use management port to connect
- `uuid` (String) uuid of the object
- `week_day` (String) 'Monday': Monday; 'Tuesday': Tuesday; 'Wednesday': Wednesday; 'Thursday': Thursday; 'Friday': Friday; 'Saturday': Saturday; 'Sunday': Sunday;
- `week_time` (String) Time of day to update (hh:mm) in 24 hour local time
- `weekly` (Number) Every week

### Read-Only

- `id` (String) The ID of this resource.

