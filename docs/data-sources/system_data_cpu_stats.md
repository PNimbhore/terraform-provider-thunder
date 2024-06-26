---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_system_data_cpu_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_system_data_cpu_stats: Statistics for the object data-cpu
  PLACEHOLDER
---

# thunder_system_data_cpu_stats (Data Source)

`thunder_system_data_cpu_stats`: Statistics for the object data-cpu

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_system_data_cpu_stats" "thunder_system_data_cpu_stats" {
}
output "get_system_data_cpu_stats" {
  value = ["${data.thunder_system_data_cpu_stats.thunder_system_data_cpu_stats}"]
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

- `cpu_1` (Number) Data CPU-1
- `cpu_10` (Number) Data CPU-10
- `cpu_100` (Number) Data CPU-100
- `cpu_11` (Number) Data CPU-11
- `cpu_12` (Number) Data CPU-12
- `cpu_13` (Number) Data CPU-13
- `cpu_14` (Number) Data CPU-14
- `cpu_15` (Number) Data CPU-15
- `cpu_16` (Number) Data CPU-16
- `cpu_17` (Number) Data CPU-17
- `cpu_18` (Number) Data CPU-18
- `cpu_19` (Number) Data CPU-19
- `cpu_2` (Number) Data CPU-2
- `cpu_20` (Number) Data CPU-20
- `cpu_21` (Number) Data CPU-21
- `cpu_22` (Number) Data CPU-22
- `cpu_23` (Number) Data CPU-23
- `cpu_24` (Number) Data CPU-24
- `cpu_25` (Number) Data CPU-25
- `cpu_26` (Number) Data CPU-26
- `cpu_27` (Number) Data CPU-27
- `cpu_28` (Number) Data CPU-28
- `cpu_29` (Number) Data CPU-29
- `cpu_3` (Number) Data CPU-3
- `cpu_30` (Number) Data CPU-30
- `cpu_31` (Number) Data CPU-31
- `cpu_32` (Number) Data CPU-32
- `cpu_33` (Number) Data CPU-33
- `cpu_34` (Number) Data CPU-34
- `cpu_35` (Number) Data CPU-35
- `cpu_36` (Number) Data CPU-36
- `cpu_37` (Number) Data CPU-37
- `cpu_38` (Number) Data CPU-38
- `cpu_39` (Number) Data CPU-39
- `cpu_4` (Number) Data CPU-4
- `cpu_40` (Number) Data CPU-40
- `cpu_41` (Number) Data CPU-41
- `cpu_42` (Number) Data CPU-42
- `cpu_43` (Number) Data CPU-43
- `cpu_44` (Number) Data CPU-44
- `cpu_45` (Number) Data CPU-45
- `cpu_46` (Number) Data CPU-46
- `cpu_47` (Number) Data CPU-47
- `cpu_48` (Number) Data CPU-48
- `cpu_49` (Number) Data CPU-49
- `cpu_5` (Number) Data CPU-5
- `cpu_50` (Number) Data CPU-50
- `cpu_51` (Number) Data CPU-51
- `cpu_52` (Number) Data CPU-52
- `cpu_53` (Number) Data CPU-53
- `cpu_54` (Number) Data CPU-54
- `cpu_55` (Number) Data CPU-55
- `cpu_56` (Number) Data CPU-56
- `cpu_57` (Number) Data CPU-57
- `cpu_58` (Number) Data CPU-58
- `cpu_59` (Number) Data CPU-59
- `cpu_6` (Number) Data CPU-6
- `cpu_60` (Number) Data CPU-60
- `cpu_61` (Number) Data CPU-61
- `cpu_62` (Number) Data CPU-62
- `cpu_63` (Number) Data CPU-63
- `cpu_64` (Number) Data CPU-64
- `cpu_65` (Number) Data CPU-65
- `cpu_66` (Number) Data CPU-66
- `cpu_67` (Number) Data CPU-67
- `cpu_68` (Number) Data CPU-68
- `cpu_69` (Number) Data CPU-69
- `cpu_7` (Number) Data CPU-7
- `cpu_70` (Number) Data CPU-70
- `cpu_71` (Number) Data CPU-71
- `cpu_72` (Number) Data CPU-72
- `cpu_73` (Number) Data CPU-73
- `cpu_74` (Number) Data CPU-74
- `cpu_75` (Number) Data CPU-75
- `cpu_76` (Number) Data CPU-76
- `cpu_77` (Number) Data CPU-77
- `cpu_78` (Number) Data CPU-78
- `cpu_79` (Number) Data CPU-79
- `cpu_8` (Number) Data CPU-8
- `cpu_80` (Number) Data CPU-80
- `cpu_81` (Number) Data CPU-81
- `cpu_82` (Number) Data CPU-82
- `cpu_83` (Number) Data CPU-83
- `cpu_84` (Number) Data CPU-84
- `cpu_85` (Number) Data CPU-85
- `cpu_86` (Number) Data CPU-86
- `cpu_87` (Number) Data CPU-87
- `cpu_88` (Number) Data CPU-88
- `cpu_89` (Number) Data CPU-89
- `cpu_9` (Number) Data CPU-9
- `cpu_90` (Number) Data CPU-90
- `cpu_91` (Number) Data CPU-91
- `cpu_92` (Number) Data CPU-92
- `cpu_93` (Number) Data CPU-93
- `cpu_94` (Number) Data CPU-94
- `cpu_95` (Number) Data CPU-95
- `cpu_96` (Number) Data CPU-96
- `cpu_97` (Number) Data CPU-97
- `cpu_98` (Number) Data CPU-98
- `cpu_99` (Number) Data CPU-99
- `data_cpu_number` (Number) Number of data cpus


