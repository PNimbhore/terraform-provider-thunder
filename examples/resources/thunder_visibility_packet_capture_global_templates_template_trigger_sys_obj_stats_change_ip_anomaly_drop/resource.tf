provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_ip_anomaly_drop" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_ip_anomaly_drop" {

  name = "test"
  trigger_stats_rate {
    threshold_exceeded_by = 5
    duration              = 60
    land                  = 1
    emp_frg               = 1
    emp_mic_frg           = 1
    opt                   = 1
    frg                   = 1
    bad_ip_hdrlen         = 1
    bad_ip_flg            = 1
    bad_ip_ttl            = 1
    no_ip_payload         = 1
    over_ip_payload       = 1
    bad_ip_payload_len    = 1
    bad_ip_frg_offset     = 1
    csum                  = 1
    pod                   = 1
    bad_tcp_urg_offset    = 1
    tcp_sht_hdr           = 1
    tcp_bad_iplen         = 1
    tcp_null_frg          = 1
    tcp_null_scan         = 1
    tcp_syn_fin           = 1
    tcp_xmas              = 1
    tcp_xmas_scan         = 1
    tcp_syn_frg           = 1
    tcp_frg_hdr           = 1
    tcp_bad_csum          = 1
    udp_srt_hdr           = 1
    udp_bad_len           = 1
    udp_kerb_frg          = 1
    udp_port_lb           = 1
    udp_bad_csum          = 1
    runt_ip_hdr           = 1
    runt_tcp_udp_hdr      = 1
    ipip_tnl_msmtch       = 1
    tcp_opt_err           = 1
    ipip_tnl_err          = 1
    vxlan_err             = 1
    nvgre_err             = 1
    gre_pptp_err          = 1
    ipv6_eh_hbh           = 1
    ipv6_eh_dest          = 1
    ipv6_eh_routing       = 1
    ipv6_eh_frag          = 1
    ipv6_eh_ah            = 1
    ipv6_eh_esp           = 1
    ipv6_eh_mobility      = 1
    ipv6_eh_none          = 1
    ipv6_eh_other         = 1
    ipv6_eh_malformed     = 1
  }
}