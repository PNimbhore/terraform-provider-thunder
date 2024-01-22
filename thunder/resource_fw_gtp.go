package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwGtp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_gtp`: Configure GTP\n\n__PLACEHOLDER__",
		CreateContext: resourceFwGtpCreate,
		UpdateContext: resourceFwGtpUpdate,
		ReadContext:   resourceFwGtpRead,
		DeleteContext: resourceFwGtpDelete,

		Schema: map[string]*schema.Schema{
			"apn_log_periodicity": {
				Type: schema.TypeInt, Optional: true, Description: "Periodic Logging Frequency(In Minutes)",
			},
			"apn_prefix": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"apn_prefix_list": {
				Type: schema.TypeString, Optional: true, Description: "Class List (Class List Name)",
			},
			"echo_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 120, Description: "echo message timeout (minutes) (echo-timeout (default 120))",
			},
			"gtp_value": {
				Type: schema.TypeString, Optional: true, Description: "'enable': Enable GTP Inspection;",
			},
			"insertion_mode": {
				Type: schema.TypeString, Optional: true, Description: "'monitor': Enable inline view-only mode; 'skip-state-checks': Enable skip stateful checks mode;",
			},
			"ne_v4_log_periodicity": {
				Type: schema.TypeInt, Optional: true, Description: "Periodic Logging Frequency(In Minutes)",
			},
			"ne_v6_log_periodicity": {
				Type: schema.TypeInt, Optional: true, Description: "Periodic Logging Frequency(In Minutes)",
			},
			"network_element": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"network_element_list_v4": {
				Type: schema.TypeString, Optional: true, Description: "Class List (Class List Name)",
			},
			"network_element_list_v6": {
				Type: schema.TypeString, Optional: true, Description: "Class List (Class List Name)",
			},
			"path_mgmt_logging": {
				Type: schema.TypeString, Optional: true, Description: "'enable-log': Enable Log for Path Management;",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'out-of-session-memory': Out of Tunnel Memory; 'no-fwd-route': No Forward Route; 'no-rev-route': No Reverse Route; 'gtp-smp-created': GTP SMP Created; 'gtp-smp-marked-deleted': GTP SMP Marked Deleted; 'gtp-smp-deleted': GTP SMP Deleted; 'smp-creation-failed': GTP-U SMP Helper Session Creation Failed; 'gtp-smp-path-created': GTP SMP PATH Created; 'gtp-smp-path-freed': GTP SMP PATH MEM freed; 'gtp-smp-path-allocated': GTP SMP PATH MEM allocated; 'gtp-smp-path-creation-failed': GTP SMP PATH creation Failed; 'gtp-smp-path-check-failed': GTP SMP PATH check Failed; 'gtp-smp-check-failed': GTP SMP check Failed; 'gtp-smp-session-count-check-failed': GTP-U session count is not in range of 0-11 in GTP-C SMP; 'gtp-c-ref-count-smp-exceeded': GTP-C session count on C-smp exceeded 2; 'gtp-u-smp-in-rml-with-sess': GTP-U smp is marked RML with U-session; 'gtp-u-pkt-fwd-conn-create': GTP-U pkt fwded while creating conn with gtp toggling; 'gtp-c-pkt-fwd-conn-create': GTP-C pkt fwded while creating conn with gtp toggling; 'gtp-echo-pkt-fwd-conn-create': GTP-ECHO pkt fwded while creating conn with gtp toggling; 'gtp-tunnel-rate-limit-entry-create-success': GTP Tunnel Level Rate Limit Entry Create Success; 'gtp-tunnel-rate-limit-entry-create-failure': GTP Tunnel Level Rate Limit Entry Create Failure; 'gtp-tunnel-rate-limit-entry-deleted': GTP Tunnel Level Rate Limit Entry Deleted; 'gtp-rate-limit-smp-created': GTP Rate Limit SMP Created; 'gtp-rate-limit-smp-freed': GTP Rate Limit SMP Freed; 'gtp-rate-limit-smp-create-failure': GTP Rate Limit SMP Create Failure; 'gtp-rate-limit-t3-ctr-create-failure': GTP Rate Limit Dynamic Counters Create Failure; 'gtp-rate-limit-entry-create-failure': GTP Rate Limit Entry Create Failure; 'gtp-echo-conn-created': GTP Echo Request Conn Created; 'gtp-echo-conn-deleted': GTP Echo Request conn Deleted; 'gtp-node-restart-echo': GTP Node Restoration due to Recovery IE in Echo; 'gtp-c-echo-path-failure': GTP-C Path Failure due to Echo; 'drop-vld-gtp-echo-out-of-state-': GTP Echo Out of State Drop; 'drop-vld-gtp-echo-ie-len-exceed-msg-len': GTP Echo IE Length Exceeds Message Length; 'gtp-create-session-request-retransmit': GTP-C Retransmitted Create Session Request; 'gtp-add-bearer-request-retransmit': GTP-C Retransmitted Add Bearer Request; 'gtp-delete-session-request-retransmit': GTP-C Retransmitted Delete Session Request; 'gtp-handover-request-retransmit': GTP Handover Request Retransmit; 'gtp-del-bearer-request-retransmit': GTP-C Retransmitted Delete Bearer Request; 'gtp-add-bearer-response-retransmit': GTP-C Retransmitted Add Bearer Response; 'gtp-create-session-request-retx-drop': GTP-C Retransmitted Create Session Request dropped; 'gtp-u-out-of-state-drop': GTP-U Out of state Drop; 'gtp-c-handover-request-out-of-state-drop': GTP-C Handover Request Out of state Drop; 'gtp-v1-c-nsapi-not-found-in-delete-req': GTPv1-C NSAPI Not Found in GTP Request; 'gtp-v2-c-bearer-not-found-in-delete-req': GTPv2-C Bearer Not Found in GTP Request; 'gtp-v2-c-bearer-not-found-in-delete-resp': GTPv2-C Bearer Not Found in GTP Response; 'gtp-multiple-handover-request': GTP Multiple Handover Request; 'gtp-rr-message-drop': GTP Message Dropped in RR Mode; 'gtp-rr-echo-message-dcmsg': GTP Echo Message Sent to home CPU in RR Mode; 'gtp-rr-c-message-dcmsg': GTP-C Message Sent to home CPU in RR Mode; 'drop-gtp-frag-or-jumbo-pkt': GTP Fragmented or JUMBO packet Drop; 'response-with-reject-cause-forwarded': GTP-C Response with Reject Cause Forwarded; 'gtp-c-message-forwarded-without-conn': GTP-C Message Forwarded without Conn; 'gtp-v0-c-ver-not-supp': GTPv0-C Version not supported indication; 'gtp-v1-c-ver-not-supp': GTPv1-C Version not supported indication; 'gtp-v2-c-ver-not-supp': GTPv2-C Version not supported indication; 'gtp-v1-extn-hdt-notif': GTPV1 Supported Extension header notification; 'gtp-u-error-ind': GTP-U Error Indication; 'gtp-c-handover-in-progress-with-conn': GTP-C mesg matching conn with HO In Progress; 'gtp-ho-in-progress-handover-request': GTP-C ho mesg matching conn with HO In Progress; 'gtp-correct-conn-ho-in-progress-handover-request': GTP-C ho mesg matching correct conn(reuse teid) with HO In Progress; 'gtp-wrong-conn-ho-in-progress-handover-request': GTP-C ho mesg matching wrong conn(new teid) with HO In Progress; 'gtp-ho-in-progress-handover-response': GTP-C ho response matching a conn with HO In Progress; 'gtp-ho-in-progress-c-mesg': GTP-C other than ho mesg matching conn with HO In Progress; 'gtp-unset-ho-flag-reuse-teid': GTP-C SGW reuse teid with ho and unset ho flag; 'gtp-refresh-c-conn-reuse-teid': GTP-C SGW reuse teid with ho and refresh old conn; 'gtp-rematch-smp-matching-conn': GTP-C rematch smp with packet matching conn; 'gtp-wrong-conn-handover-request': GTP-C ho mesg matching wrong conn(new teid) with no HO flag; 'gtp-refresh-conn-set-ho-flag-latest': GTP-C SGW refresh old conn and set ho flag on latest smp; 'gtp-c-process-pkt-drop': GTP-C process pkt drop; 'gtp-c-fwd-pkt-drop': GTP-C fwd pkt drop; 'gtp-c-rev-pkt-drop': GTP-C rev pkt drop; 'gtp-c-fwd-v1-other': GTP-C fwd v1 other messages; 'gtp-c-fwd-v2-other': GTP-C fwd v2 other messages; 'gtp-c-rev-v1-other': GTP-C rev v1 other messages; 'gtp-c-rev-v2-other': GTP-C rev v2 other messages; 'gtp-c-going-thru-fw-lookup': GTP-C mesg going thru fw lookup can be resp or l5 mesg not matching smp; 'gtp-c-conn-create-pkt-drop': GTP-C conn creation drop; 'gtp-c-pkt-fwd-conn-create-no-fteid': GTP-C pkt fwded while creating conn when no FTEID; 'gtp-inter-pu-mstr-to-bld-dcmsg-fail': GTP inter-PU dcmsg failed from Master to Blade; 'gtp-inter-pu-mstr-to-bld-dcmsg-sent': GTP inter-PU Master to Blade dcmsg sent; 'gtp-inter-pu-mstr-to-bld-dcmsg-recv': GTP inter-PU dcmsg received on blade; 'gtp-inter-pu-mstr-to-bld-query-sent': GTP inter-PU query sent from Master to Blade; 'gtp-inter-pu-mstr-to-bld-query-recv': GTP inter-PU GTP-C query received on Blade; 'gtp-inter-pu-mstr-to-bld-query-resp-sent': GTP inter-PU GTP-C query response sent from Master to Blade; 'gtp-inter-pu-bld-to-mstr-dcmsg-fail': GTP inter-PU dcmsg failed from Blade to Master; 'gtp-inter-pu-bld-to-mstr-dcmsg-sent': GTP inter-PU Blade to Master dcmsg sent; 'gtp-inter-pu-bld-to-mstr-dcmsg-recv': GTP inter-PU dcmsg received on Master; 'gtp-inter-pu-bld-to-mstr-query-sent': GTP inter-PU query sent from Blade to Master; 'gtp-inter-pu-bld-to-mstr-query-recv': GTP inter-PU GTP-C query received on Master; 'gtp-inter-pu-bld-to-mstr-query-resp-sent': GTP inter-PU GTP-C query response sent from Blade to Master; 'gtp-mstr-to-bld-query-resp-fail': GTP inter-PU dcmsg of query response failed from Master to Blade; 'gtp-bld-to-mstr-query-resp-fail': GTP inter-PU dcmsg of query response failed from Blade to Master; 'gtp-c-smp-refer-stale-idx': GTP-C SMP referring stale C-conn idx; 'gtp-smp-dec-sess-count-check-failed': GTP-U session count is 0 in GTP-C SMP; 'gtp-c-freed-conn-check': GTP-C freed conn accessed; 'gtp-c-conn-not-in-rml-when-freed': GTP-C conn not in rml when tuple is freed; 'gtp-v0-c-uplink-ingress-packets': GTPv0-C Uplink Ingress Packets; 'gtp-v0-c-uplink-egress-packets': GTPv0-C Uplink Egress Packets; 'gtp-v0-c-downlink-ingress-packets': GTPv0-C Downlink Ingress Packets; 'gtp-v0-c-downlink-egress-packets': GTPv0-C Downlink Egress Packets; 'gtp-v0-c-uplink-ingress-bytes': GTPv0-C Uplink Ingress Bytes; 'gtp-v0-c-uplink-egress-bytes': GTPv0-C Uplink Egress Bytes; 'gtp-v0-c-downlink-ingress-bytes': GTPv0-C Downlink Ingress Bytes; 'gtp-v0-c-downlink-egress-bytes': GTPv0-C Downlink Egress Bytes; 'gtp-v1-c-uplink-ingress-packets': GTPv1-C Uplink Ingress Packets; 'gtp-v1-c-uplink-egress-packets': GTPv1-C Uplink Egress Packets;",
						},
						"counters2": {
							Type: schema.TypeString, Optional: true, Description: "'gtp-v1-c-downlink-ingress-packets': GTPv1-C Downlink Ingress Packets; 'gtp-v1-c-downlink-egress-packets': GTPv1-C Downlink Egress Packets; 'gtp-v1-c-uplink-ingress-bytes': GTPv1-C Uplink Ingress Bytes; 'gtp-v1-c-uplink-egress-bytes': GTPv1-C Uplink Egress Bytes; 'gtp-v1-c-downlink-ingress-bytes': GTPv1-C Downlink Ingress Bytes; 'gtp-v1-c-downlink-egress-bytes': GTPv1-C Downlink Egress Bytes; 'gtp-v2-c-uplink-ingress-packets': GTPv2-C Uplink Ingress Packets; 'gtp-v2-c-uplink-egress-packets': GTPv2-C Uplink Egress Packets; 'gtp-v2-c-downlink-ingress-packets': GTPv2-C Downlink Ingress Packets; 'gtp-v2-c-downlink-egress-packets': GTPv2-C Downlink Egress Packets; 'gtp-v2-c-uplink-ingress-bytes': GTPv2-C Uplink Ingress Bytes; 'gtp-v2-c-uplink-egress-bytes': GTPv2-C Uplink Egress Bytes; 'gtp-v2-c-downlink-ingress-bytes': GTPv2-C Downlink Ingress Bytes; 'gtp-v2-c-downlink-egress-bytes': GTPv2-C Downlink Egress Bytes; 'gtp-u-uplink-ingress-packets': GTP-U Uplink Ingress Packets; 'gtp-u-uplink-egress-packets': GTP-U Uplink Egress Packets; 'gtp-u-downlink-ingress-packets': GTP-U Downlink Ingress Packets; 'gtp-u-downlink-egress-packets': GTP-U Downlink Egress Packets; 'gtp-u-uplink-ingress-bytes': GTP-U Uplink Ingress Bytes; 'gtp-u-uplink-egress-bytes': GTP-U Uplink Egress Bytes; 'gtp-u-downlink-ingress-bytes': GTP-U Downlink Ingress Bytes; 'gtp-u-downlink-egress-bytes': GTP-U Downlink Egress Bytes; 'gtp-v0-c-create-synced': GTPv0-C Tunnel Create Synced; 'gtp-v1-c-create-synced': GTPv1-C Tunnel Create Synced; 'gtp-v2-c-create-synced': GTPv2-C Tunnel Create Synced; 'gtp-v0-c-delete-synced': GTPv0-C Tunnel Delete Synced; 'gtp-v1-c-delete-synced': GTPv1-C Tunnel Delete Synced; 'gtp-v2-c-delete-synced': GTPv2-C Tunnel Delete Synced; 'gtp-v0-c-create-sync-rx': GTPv0-C Tunnel Create Sync Received on Standby; 'gtp-v1-c-create-sync-rx': GTPv1-C Tunnel Create Sync Received on Standby; 'gtp-v2-c-create-sync-rx': GTPv2-C Tunnel Create Sync Received on Standby; 'gtp-v0-c-delete-sync-rx': GTPv0-C Tunnel Delete Sync Received on Standby; 'gtp-v1-c-delete-sync-rx': GTPv1-C Tunnel Delete Sync Received on Standby; 'gtp-v2-c-delete-sync-rx': GTPv2-C Tunnel Delete Sync Received on Standby; 'gtp-handover-synced': GTP Handover Synced; 'gtp-handover-sync-rx': GTP Handover Sync Received on Standby; 'gtp-smp-add-bearer-synced': GTP SMP Add Bearer Synced; 'gtp-smp-del-bearer-synced': GTP SMP Del Bearer Synced; 'gtp-smp-additional-bearer-synced': GTP SMP Additional Bearer Synced; 'gtp-smp-add-bearer-sync-rx': GTP SMP Add Bearer Sync Received on Standby; 'gtp-smp-del-bearer-sync-rx': GTP SMP Del Bearer Sync Received on Standby; 'gtp-smp-additional-bearer-sync-rx': GTP SMP Additional Bearer Sync Received on Standby; 'gtp-add-bearer-sync-not-rx-on-standby': GTP Add Bearer Sync Not Received on Standby; 'gtp-add-bearer-sync-with-periodic-update-on-standby': GTP Bearer Added on Standby with Periodic Sync; 'gtp-delete-bearer-sync-with-periodic-update-on-standby': GTP Bearer Deleted on Standy with Periodic Sync; 'gtp-v0-c-echo-create-synced': GTPv0-C Echo Create Synced; 'gtp-v1-c-echo-create-synced': GTPv1-C Echo Create Synced; 'gtp-v2-c-echo-create-synced': GTPv2-C Echo Create Synced; 'gtp-v0-c-echo-create-sync-rx': GTPv0-C-Echo Create Sync Received on Standby; 'gtp-v1-c-echo-create-sync-rx': GTPv1-C-Echo Create Sync Received on Standby; 'gtp-v2-c-echo-create-sync-rx': GTPv2-C-Echo Create Sync Received on Standby; 'gtp-v0-c-echo-del-synced': GTPv0-C Echo Delete Synced; 'gtp-v1-c-echo-del-synced': GTPv1-C Echo Delete Synced; 'gtp-v2-c-echo-del-synced': GTPv2-C Echo Delete Synced; 'gtp-v0-c-echo-del-sync-rx': GTPv0-C-Echo Delete Sync Received on Standby; 'gtp-v1-c-echo-del-sync-rx': GTPv1-C-Echo Delete Sync Received on Standby; 'gtp-v2-c-echo-del-sync-rx': GTPv2-C-Echo Delete Sync Received on Standby; 'drop-gtp-conn-creation-standby': GTP Conn creation on Standby Drop; 'gtp-u-synced-before-control': GTP-U Tunnel synced before corresponding GTP-C; 'gtp-c-l5-synced-before-l3': GTP-C L5 conn synced before corresponding L3 GTP-C conn; 'gtp-smp-path-del-synced': GTP SMP path delete Synced; 'gtp-smp-path-del-sync-rx': GTP SMP path delete Sync Received on Standby; 'gtp-not-enabled-on-standby': GTP Not Enabled on Standby; 'gtp-ip-version-v4-v6': GTP IP versions of V4&V6 in FTEID; 'drop-gtp-ip-version-mismatch-fteid': GTP IP version mismatch for req & response FTEIDs; 'drop-gtp-ip-version-mismatch-ho-fteid': GTP IP version mismatch in Handover SGW FTEID; 'gtp-u-message-length-mismatch': GTP-U Message Length Mismatch Across Layers; 'gtp-path-message-length-mismatch': GTP-Path Message Length Mismatch Across Layers; 'drop-gtp-missing-cond-ie-bearer-ctx': Missing conditional IE in bearer context Drop; 'drop-gtp-bearer-not-found-in-resp': GTP Bearer not found in response; 'gtp-stateless-forward': GTP Stateless Forward; 'gtp-l3-conn-deleted': GTP L3 conn deleted; 'gtp-l5-conn-created': GTP L5 conn created; 'gtp-monitor-forward': GTP messages forwarded via monitor mode; 'gtp-u_inner-ip-not-present': GTP-U inner IP not present; 'gtp-ext_hdr-incorrect-length': GTP Extension header incorrect length;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceFwGtpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwGtpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwGtp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwGtpRead(ctx, d, meta)
	}
	return diags
}

func resourceFwGtpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwGtpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwGtp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwGtpRead(ctx, d, meta)
	}
	return diags
}
func resourceFwGtpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwGtpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwGtp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwGtpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwGtpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwGtp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectFwGtpApnPrefix369(d []interface{}) edpt.FwGtpApnPrefix369 {

	var ret edpt.FwGtpApnPrefix369
	return ret
}

func getObjectFwGtpNetworkElement370(d []interface{}) edpt.FwGtpNetworkElement370 {

	var ret edpt.FwGtpNetworkElement370
	return ret
}

func getSliceFwGtpSamplingEnable(d []interface{}) []edpt.FwGtpSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.FwGtpSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwGtpSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		oi.Counters2 = in["counters2"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFwGtp(d *schema.ResourceData) edpt.FwGtp {
	var ret edpt.FwGtp
	ret.Inst.ApnLogPeriodicity = d.Get("apn_log_periodicity").(int)
	ret.Inst.ApnPrefix = getObjectFwGtpApnPrefix369(d.Get("apn_prefix").([]interface{}))
	ret.Inst.ApnPrefixList = d.Get("apn_prefix_list").(string)
	ret.Inst.EchoTimeout = d.Get("echo_timeout").(int)
	ret.Inst.GtpValue = d.Get("gtp_value").(string)
	ret.Inst.InsertionMode = d.Get("insertion_mode").(string)
	ret.Inst.NeV4LogPeriodicity = d.Get("ne_v4_log_periodicity").(int)
	ret.Inst.NeV6LogPeriodicity = d.Get("ne_v6_log_periodicity").(int)
	ret.Inst.NetworkElement = getObjectFwGtpNetworkElement370(d.Get("network_element").([]interface{}))
	ret.Inst.NetworkElementListV4 = d.Get("network_element_list_v4").(string)
	ret.Inst.NetworkElementListV6 = d.Get("network_element_list_v6").(string)
	ret.Inst.PathMgmtLogging = d.Get("path_mgmt_logging").(string)
	ret.Inst.SamplingEnable = getSliceFwGtpSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
