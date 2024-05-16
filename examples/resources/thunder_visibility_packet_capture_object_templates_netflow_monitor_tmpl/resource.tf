provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_netflow_monitor_tmpl" "thunder_visibility_packet_capture_object_templates_netflow_monitor_tmpl" {

  name           = "test"
  capture_config = "test"
  trigger_stats_rate {
    threshold_exceeded_by                   = 5
    duration                                = 60
    nat44_records_sent_failure              = 1
    nat64_records_sent_failure              = 1
    dslite_records_sent_failure             = 1
    session_event_nat44_records_sent_failur = 1
    session_event_nat64_records_sent_failur = 1
    session_event_dslite_records_sent_failu = 1
    session_event_fw4_records_sent_failure  = 1
    session_event_fw6_records_sent_failure  = 1
    port_mapping_nat44_records_sent_failure = 1
    port_mapping_nat64_records_sent_failure = 1
    port_mapping_dslite_records_sent_failur = 1
    netflow_v5_records_sent_failure         = 1
    netflow_v5_ext_records_sent_failure     = 1
    port_batching_nat44_records_sent_failur = 1
    port_batching_nat64_records_sent_failur = 1
    port_batching_dslite_records_sent_failu = 1
    port_batching_v2_nat44_records_sent_fai = 1
    port_batching_v2_nat64_records_sent_fai = 1
    port_batching_v2_dslite_records_sent_fa = 1
    custom_session_event_nat44_creation_rec = 1
    custom_session_event_nat64_creation_rec = 1
    custom_session_event_dslite_creation_re = 1
    custom_session_event_nat44_deletion_rec = 1
    custom_session_event_nat64_deletion_rec = 1
    custom_session_event_dslite_deletion_re = 1
    custom_session_event_fw4_creation_recor = 1
    custom_session_event_fw6_creation_recor = 1
    custom_session_event_fw4_deletion_recor = 1
    custom_session_event_fw6_deletion_recor = 1
    custom_deny_reset_event_fw4_records_sen = 1
    custom_deny_reset_event_fw6_records_sen = 1
    custom_port_mapping_nat44_creation_reco = 1
    custom_port_mapping_nat64_creation_reco = 1
    custom_port_mapping_dslite_creation_rec = 1
    custom_port_mapping_nat44_deletion_reco = 1
    custom_port_mapping_nat64_deletion_reco = 1
    custom_port_mapping_dslite_deletion_rec = 1
    custom_port_batching_nat44_creation_rec = 1
    custom_port_batching_nat64_creation_rec = 1
    custom_port_batching_dslite_creation_re = 1
    custom_port_batching_nat44_deletion_rec = 1
    custom_port_batching_nat64_deletion_rec = 1
    custom_port_batching_dslite_deletion_re = 1
    custom_port_batching_v2_nat44_creation_ = 1
    custom_port_batching_v2_nat64_creation_ = 1
    custom_port_batching_v2_dslite_creation = 1
    custom_port_batching_v2_nat44_deletion_ = 1
    custom_port_batching_v2_nat64_deletion_ = 1
    custom_port_batching_v2_dslite_deletion = 1
    custom_gtp_c_tunnel_event_records_sent_ = 1
    custom_gtp_u_tunnel_event_records_sent_ = 1
    custom_gtp_deny_event_records_sent_fail = 1
    custom_gtp_info_event_records_sent_fail = 1
    custom_fw_iddos_entry_created_records_s = 1
    custom_fw_iddos_entry_deleted_records_s = 1
    custom_fw_sesn_limit_exceeded_records_s = 1
    custom_nat_iddos_l3_entry_created_recor = 1
    custom_nat_iddos_l3_entry_deleted_recor = 1
    custom_nat_iddos_l4_entry_created_recor = 1
    custom_nat_iddos_l4_entry_deleted_recor = 1
  }
  trigger_stats_severity {
    error          = 1
    error_alert    = 1
    error_warning  = 1
    error_critical = 1
    drop           = 1
    drop_alert     = 1
    drop_warning   = 1
    drop_critical  = 1
  }
  user_tag = "2"
}