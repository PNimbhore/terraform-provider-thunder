---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cgnv6_dns64_virtualserver_port Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cgnv6_dns64_virtualserver_port: Virtual Port
  PLACEHOLDER
---

# thunder_cgnv6_dns64_virtualserver_port (Resource)

`thunder_cgnv6_dns64_virtualserver_port`: Virtual Port

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_dns64_virtualserver_port" "thunder_cgnv6_dns64_virtualserver_port" {

  name        = "test"
  port_number = 22866
  protocol    = "dns-udp"
  sampling_enable {
    counters1 = "all"
  }
  user_tag = "15"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name
- `port_number` (Number) Port
- `protocol` (String) 'dns-udp': DNS service over UDP;

### Optional

- `acl_id_list` (Block List) (see [below for nested schema](#nestedblock--acl_id_list))
- `acl_name_list` (Block List) (see [below for nested schema](#nestedblock--acl_name_list))
- `action` (String) 'enable': Enable; 'disable': Disable;
- `auto` (Number) Configure auto NAT for the vport
- `packet_capture_template` (String) Name of the packet capture template to be bind with this object
- `pool` (String) Specify NAT pool or pool group
- `precedence` (Number) Set auto NAT pool as higher precedence for source NAT
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--sampling_enable))
- `service_group` (String) Bind a Service Group to this Virtual Server (Service Group Name)
- `template_dns` (String) DNS template (DNS template name)
- `template_policy` (String) Policy Template (Policy template name)
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--acl_id_list"></a>
### Nested Schema for `acl_id_list`

Optional:

- `acl_id` (Number) ACL id VPORT
- `acl_id_seq_num` (Number) Specify ACL precedence (sequence-number)
- `acl_id_src_nat_pool` (String) Policy based Source NAT (NAT Pool or Pool Group)


<a id="nestedblock--acl_name_list"></a>
### Nested Schema for `acl_name_list`

Optional:

- `acl_name` (String) Apply an access list name (Named Access List)
- `acl_name_seq_num` (Number) Specify ACL precedence (sequence-number)
- `acl_name_src_nat_pool` (String) Policy based Source NAT (NAT Pool or Pool Group)


<a id="nestedblock--sampling_enable"></a>
### Nested Schema for `sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'curr_conn': Current connection; 'total_l4_conn': Total L4 connections; 'total_l7_conn': Total L7 connections; 'toatal_tcp_conn': Total TCP connections; 'total_conn': Total connections; 'total_fwd_bytes': Total forward bytes; 'total_fwd_pkts': Total forward packets; 'total_rev_bytes': Total reverse bytes; 'total_rev_pkts': Total reverse packets; 'total_dns_pkts': Total DNS packets; 'total_mf_dns_pkts': Total MF DNS packets; 'es_total_failure_actions': Total failure actions; 'compression_bytes_before': Data into compression engine; 'compression_bytes_after': Data out of compression engine; 'compression_hit': Number of requests compressed; 'compression_miss': Number of requests NOT compressed; 'compression_miss_no_client': Compression miss no client; 'compression_miss_template_exclusion': Compression miss template exclusion; 'curr_req': Current requests; 'total_req': Total requests; 'total_req_succ': Total successful requests; 'peak_conn': Peak connections; 'curr_conn_rate': Current connection rate; 'last_rsp_time': Last response time; 'fastest_rsp_time': Fastest response time; 'slowest_rsp_time': Slowest response time;


