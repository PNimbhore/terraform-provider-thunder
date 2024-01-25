---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_fw_server Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_fw_server: Firewall logging Server
  PLACEHOLDER
---

# thunder_fw_server (Resource)

`thunder_fw_server`: Firewall logging Server

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_fw_server" "thunder_fw_server" {
  name         = "test"
  fqdn_name    = "test.com"
  action       = "enable"
  health_check = "ping"
  user_tag     = "testing"
  sampling_enable {
    counters1 = "all"
  }
  port_list {
    port_number  = 30
    protocol     = "tcp"
    action       = "enable"
    health_check = "ping"
    user_tag     = "test_user"
    sampling_enable {
      counters1 = "all"
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Server Name

### Optional

- `action` (String) 'enable': Enable this Real Server; 'disable': Disable this Real Server;
- `fqdn_name` (String) Server hostname
- `health_check` (String) Health Check Monitor (Health monitor name)
- `health_check_disable` (Number) Disable configured health check configuration
- `host` (String) IP Address
- `port_list` (Block List) (see [below for nested schema](#nestedblock--port_list))
- `resolve_as` (String) 'resolve-to-ipv4': Use A Query only to resolve FQDN; 'resolve-to-ipv6': Use AAAA Query only to resolve FQDN; 'resolve-to-ipv4-and-ipv6': Use A as well as AAAA Query to resolve FQDN;
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--sampling_enable))
- `server_ipv6_addr` (String) IPV6 address
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--port_list"></a>
### Nested Schema for `port_list`

Required:

- `port_number` (Number) Port Number
- `protocol` (String) 'tcp': TCP Port; 'udp': UDP Port;

Optional:

- `action` (String) 'enable': enable; 'disable': disable;
- `health_check` (String) Health Check (Monitor Name)
- `health_check_disable` (Number) Disable health check
- `packet_capture_template` (String) Name of the packet capture template to be bind with this object
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--port_list--sampling_enable))
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

<a id="nestedblock--port_list--sampling_enable"></a>
### Nested Schema for `port_list.sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'curr_conn': Current connections; 'curr_req': Current requests; 'total_req': Total requests; 'total_req_succ': Total request success; 'total_fwd_bytes': Forward bytes; 'total_fwd_pkts': Forward packets; 'total_rev_bytes': Reverse bytes; 'total_rev_pkts': Reverse packets; 'total_conn': Total connections; 'last_total_conn': Last total connections; 'peak_conn': Peak connections; 'es_resp_200': Response status 200; 'es_resp_300': Response status 300; 'es_resp_400': Response status 400; 'es_resp_500': Response status 500; 'es_resp_other': Response status other; 'es_req_count': Total proxy request; 'es_resp_count': Total proxy Response; 'es_resp_invalid_http': Total non-http response; 'total_rev_pkts_inspected': Total reverse packets inspected; 'total_rev_pkts_inspected_good_status_code': Total reverse packets with good status code inspected; 'response_time': Response time; 'fastest_rsp_time': Fastest response time; 'slowest_rsp_time': Slowest response time;



<a id="nestedblock--sampling_enable"></a>
### Nested Schema for `sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'curr-conn': Current connections; 'total-conn': Total connections; 'fwd-pkt': Forward packets; 'rev-pkt': Reverse Packets; 'peak-conn': Peak connections;

