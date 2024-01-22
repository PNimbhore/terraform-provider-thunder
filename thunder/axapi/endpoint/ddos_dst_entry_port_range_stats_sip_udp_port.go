

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntryPortRangeStats8 struct {
	Inst struct {

    PortRangeEnd int `json:"port-range-end"`
    PortRangeStart int `json:"port-range-start"`
    Protocol string `json:"protocol"`
    Stats DdosDstEntryPortRangeStats8Stats `json:"stats"`
    DstEntryName string 

	} `json:"stats"`
}


type DdosDstEntryPortRangeStats8Stats struct {
    SipUdpPort DdosDstEntryPortRangeStats8StatsSipUdpPort `json:"sip-udp-port"`
}


type DdosDstEntryPortRangeStats8StatsSipUdpPort struct {
    Request_method_ack int `json:"request_method_ack"`
    Request_method_bye int `json:"request_method_bye"`
    Request_method_cancel int `json:"request_method_cancel"`
    Request_method_invite int `json:"request_method_invite"`
    Request_method_info int `json:"request_method_info"`
    Request_method_message int `json:"request_method_message"`
    Request_method_notify int `json:"request_method_notify"`
    Request_method_options int `json:"request_method_options"`
    Request_method_prack int `json:"request_method_prack"`
    Request_method_publish int `json:"request_method_publish"`
    Request_method_register int `json:"request_method_register"`
    Request_method_refer int `json:"request_method_refer"`
    Request_method_subscribe int `json:"request_method_subscribe"`
    Request_method_update int `json:"request_method_update"`
    Request_method_unknown int `json:"request_method_unknown"`
    Request_unknown_version int `json:"request_unknown_version"`
    Keep_alive_msg int `json:"keep_alive_msg"`
    Rate1_limit_exceed int `json:"rate1_limit_exceed"`
    Rate2_limit_exceed int `json:"rate2_limit_exceed"`
    Rate3_limit_exceed int `json:"rate3_limit_exceed"`
    Src_rate1_limit_exceed int `json:"src_rate1_limit_exceed"`
    Src_rate2_limit_exceed int `json:"src_rate2_limit_exceed"`
    Src_rate3_limit_exceed int `json:"src_rate3_limit_exceed"`
    Response_1xx int `json:"response_1xx"`
    Response_2xx int `json:"response_2xx"`
    Response_3xx int `json:"response_3xx"`
    Response_4xx int `json:"response_4xx"`
    Response_5xx int `json:"response_5xx"`
    Response_6xx int `json:"response_6xx"`
    Response_unknown int `json:"response_unknown"`
    Response_unknown_version int `json:"response_unknown_version"`
    Read_start_line_error int `json:"read_start_line_error"`
    Invalid_start_line_error int `json:"invalid_start_line_error"`
    Parse_start_line_error int `json:"parse_start_line_error"`
    Line_too_long int `json:"line_too_long"`
    Line_mem_allocated int `json:"line_mem_allocated"`
    Line_mem_freed int `json:"line_mem_freed"`
    Max_uri_len_exceed int `json:"max_uri_len_exceed"`
    Too_many_header int `json:"too_many_header"`
    Invalid_header int `json:"invalid_header"`
    Header_name_too_long int `json:"header_name_too_long"`
    Parse_header_fail_error int `json:"parse_header_fail_error"`
    Max_header_value_len_exceed int `json:"max_header_value_len_exceed"`
    Max_call_id_len_exceed int `json:"max_call_id_len_exceed"`
    Header_filter_match int `json:"header_filter_match"`
    Header_filter_not_match int `json:"header_filter_not_match"`
    Header_filter_none_match int `json:"header_filter_none_match"`
    Header_filter_action_drop int `json:"header_filter_action_drop"`
    Header_filter_action_blacklist int `json:"header_filter_action_blacklist"`
    Header_filter_action_whitelist int `json:"header_filter_action_whitelist"`
    Header_filter_action_default_pass int `json:"header_filter_action_default_pass"`
    Header_filter_filter1_match int `json:"header_filter_filter1_match"`
    Header_filter_filter2_match int `json:"header_filter_filter2_match"`
    Header_filter_filter3_match int `json:"header_filter_filter3_match"`
    Header_filter_filter4_match int `json:"header_filter_filter4_match"`
    Header_filter_filter5_match int `json:"header_filter_filter5_match"`
    Max_sdp_len_exceed int `json:"max_sdp_len_exceed"`
    Body_too_big int `json:"body_too_big"`
    Get_content_fail_error int `json:"get_content_fail_error"`
    Concatenate_msg int `json:"concatenate_msg"`
    Mem_alloc_fail_error int `json:"mem_alloc_fail_error"`
    Malform_request int `json:"malform_request"`
    Udp_auth int `json:"udp_auth"`
    Udp_auth_fail int `json:"udp_auth_fail"`
    Port_rcvd int `json:"port_rcvd"`
    Port_drop int `json:"port_drop"`
    Port_pkt_sent int `json:"port_pkt_sent"`
    Port_pkt_rate_exceed int `json:"port_pkt_rate_exceed"`
    Port_kbit_rate_exceed int `json:"port_kbit_rate_exceed"`
    Port_conn_rate_exceed int `json:"port_conn_rate_exceed"`
    Port_conn_limm_exceed int `json:"port_conn_limm_exceed"`
    Port_bytes int `json:"port_bytes"`
    Outbound_port_bytes int `json:"outbound_port_bytes"`
    Outbound_port_rcvd int `json:"outbound_port_rcvd"`
    Outbound_port_pkt_sent int `json:"outbound_port_pkt_sent"`
    Port_bytes_sent int `json:"port_bytes_sent"`
    Port_bytes_drop int `json:"port_bytes_drop"`
    Port_src_bl int `json:"port_src_bl"`
    Filter_auth_fail int `json:"filter_auth_fail"`
    Spoof_detect_fail int `json:"spoof_detect_fail"`
    Sess_create int `json:"sess_create"`
    Filter_action_blacklist int `json:"filter_action_blacklist"`
    Filter_action_drop int `json:"filter_action_drop"`
    Filter_action_default_pass int `json:"filter_action_default_pass"`
    Filter_action_whitelist int `json:"filter_action_whitelist"`
    Exceed_drop_prate_src int `json:"exceed_drop_prate_src"`
    Exceed_drop_crate_src int `json:"exceed_drop_crate_src"`
    Exceed_drop_climit_src int `json:"exceed_drop_climit_src"`
    Exceed_drop_brate_src int `json:"exceed_drop_brate_src"`
    Outbound_port_bytes_sent int `json:"outbound_port_bytes_sent"`
    Outbound_port_drop int `json:"outbound_port_drop"`
    Outbound_port_bytes_drop int `json:"outbound_port_bytes_drop"`
    Udp_auth_retry_fail int `json:"udp_auth_retry_fail"`
    Exceed_drop_brate_src_pkt int `json:"exceed_drop_brate_src_pkt"`
    Port_kbit_rate_exceed_pkt int `json:"port_kbit_rate_exceed_pkt"`
    Udp_retry_init int `json:"udp_retry_init"`
    Conn_prate_excd int `json:"conn_prate_excd"`
    Ntp_monlist_req int `json:"ntp_monlist_req"`
    Ntp_monlist_resp int `json:"ntp_monlist_resp"`
    Wellknown_sport_drop int `json:"wellknown_sport_drop"`
    Payload_too_small int `json:"payload_too_small"`
    Payload_too_big int `json:"payload_too_big"`
    Udp_auth_retry_gap_drop int `json:"udp_auth_retry_gap_drop"`
    Src_udp_auth int `json:"src_udp_auth"`
    Udp_auth_pass int `json:"udp_auth_pass"`
    Udp_auth_drop int `json:"udp_auth_drop"`
    Bl int `json:"bl"`
    Src_drop int `json:"src_drop"`
    Frag_rcvd int `json:"frag_rcvd"`
    Frag_drop int `json:"frag_drop"`
    Sess_create_inbound int `json:"sess_create_inbound"`
    Sess_create_outbound int `json:"sess_create_outbound"`
    Sess_aged int `json:"sess_aged"`
    Udp_retry_pass int `json:"udp_retry_pass"`
    Udp_retry_gap_drop int `json:"udp_retry_gap_drop"`
    Filter1_match int `json:"filter1_match"`
    Filter2_match int `json:"filter2_match"`
    Filter3_match int `json:"filter3_match"`
    Filter4_match int `json:"filter4_match"`
    Filter5_match int `json:"filter5_match"`
    Filter_none_match int `json:"filter_none_match"`
    Src_payload_too_small int `json:"src_payload_too_small"`
    Src_payload_too_big int `json:"src_payload_too_big"`
    Src_ntp_monlist_req int `json:"src_ntp_monlist_req"`
    Src_ntp_monlist_resp int `json:"src_ntp_monlist_resp"`
    Src_well_known_port int `json:"src_well_known_port"`
    Src_conn_pkt_rate_excd int `json:"src_conn_pkt_rate_excd"`
    Src_udp_retry_init int `json:"src_udp_retry_init"`
    Src_filter_action_blacklist int `json:"src_filter_action_blacklist"`
    Src_filter_action_drop int `json:"src_filter_action_drop"`
    Src_filter_action_default_pass int `json:"src_filter_action_default_pass"`
    Src_filter_action_whitelist int `json:"src_filter_action_whitelist"`
    Src_frag_drop int `json:"src_frag_drop"`
    Frag_timeout int `json:"frag_timeout"`
    Src_udp_auth_fail int `json:"src_udp_auth_fail"`
    Policy_drop int `json:"policy_drop"`
    Policy_violation int `json:"policy_violation"`
    Src_udp_auth_drop int `json:"src_udp_auth_drop"`
    Src_udp_retry_gap_drop int `json:"src_udp_retry_gap_drop"`
    Src_filter1_match int `json:"src_filter1_match"`
    Src_filter2_match int `json:"src_filter2_match"`
    Src_filter3_match int `json:"src_filter3_match"`
    Src_filter4_match int `json:"src_filter4_match"`
    Src_filter5_match int `json:"src_filter5_match"`
    Src_filter_none_match int `json:"src_filter_none_match"`
    Src_filter_total_not_match int `json:"src_filter_total_not_match"`
    Src_filter_auth_fail int `json:"src_filter_auth_fail"`
    Filter_total_not_match int `json:"filter_total_not_match"`
    Sflow_internal_samples_packed int `json:"sflow_internal_samples_packed"`
    Sflow_external_samples_packed int `json:"sflow_external_samples_packed"`
    Sflow_internal_packets_sent int `json:"sflow_internal_packets_sent"`
    Sflow_external_packets_sent int `json:"sflow_external_packets_sent"`
    Exceed_action_tunnel int `json:"exceed_action_tunnel"`
    Src_udp_auth_timeout int `json:"src_udp_auth_timeout"`
    Src_udp_retry_pass int `json:"src_udp_retry_pass"`
    Pattern_recognition_proceeded int `json:"pattern_recognition_proceeded"`
    Pattern_not_found int `json:"pattern_not_found"`
    Pattern_recognition_generic_error int `json:"pattern_recognition_generic_error"`
    Pattern_filter1_match int `json:"pattern_filter1_match"`
    Pattern_filter2_match int `json:"pattern_filter2_match"`
    Pattern_filter3_match int `json:"pattern_filter3_match"`
    Pattern_filter4_match int `json:"pattern_filter4_match"`
    Pattern_filter5_match int `json:"pattern_filter5_match"`
    Pattern_filter_drop int `json:"pattern_filter_drop"`
    Pattern_recognition_sampling_started int `json:"pattern_recognition_sampling_started"`
    Pattern_recognition_pattern_changed int `json:"pattern_recognition_pattern_changed"`
    Dst_hw_drop int `json:"dst_hw_drop"`
    Token_authentication_mismatched int `json:"token_authentication_mismatched"`
    Token_authentication_invalid int `json:"token_authentication_invalid"`
    Token_authentication_curr_salt_matched int `json:"token_authentication_curr_salt_matched"`
    Token_authentication_prev_salt_matched int `json:"token_authentication_prev_salt_matched"`
    Token_authentication_session_created int `json:"token_authentication_session_created"`
    Token_authentication_session_created_fail int `json:"token_authentication_session_created_fail"`
    Snat_fail int `json:"snat_fail"`
    Exceed_action_drop int `json:"exceed_action_drop"`
}

func (p *DdosDstEntryPortRangeStats8) GetId() string{
    return "1"
}

func (p *DdosDstEntryPortRangeStats8) getPath() string{
    return "ddos/dst/entry/" +p.Inst.DstEntryName + "/port-range/" +strconv.Itoa(p.Inst.PortRangeStart)+"+" +strconv.Itoa(p.Inst.PortRangeEnd)+"+"+p.Inst.Protocol+"/stats?sip-udp-port=true"
}

func (p *DdosDstEntryPortRangeStats8) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryPortRangeStats8::Post")
    headers := axapi.GenRequestHeader(authToken)
        payloadBytes, err := axapi.SerializeToJson(p)
        if err != nil {
            logger.Println("Failed to serialize struct as json", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
    return err
}

func (p *DdosDstEntryPortRangeStats8) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryPortRangeStats8::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return err
}
func (p *DdosDstEntryPortRangeStats8) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryPortRangeStats8::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *DdosDstEntryPortRangeStats8) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryPortRangeStats8::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
