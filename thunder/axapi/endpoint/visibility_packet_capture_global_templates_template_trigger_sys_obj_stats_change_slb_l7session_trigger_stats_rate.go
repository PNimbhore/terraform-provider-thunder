

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL7sessionTriggerStatsRate struct {
	Inst struct {

    Conn_not_exist int `json:"conn_not_exist"`
    Data_cb_failed int `json:"data_cb_failed"`
    Duration int `json:"duration" dval:"60"`
    Err_cb_failed int `json:"err_cb_failed"`
    Err_event int `json:"err_event"`
    Hps_fwdreq_fail int `json:"hps_fwdreq_fail"`
    Server_conn_failed int `json:"server_conn_failed"`
    Server_select_fail int `json:"server_select_fail"`
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Uuid string `json:"uuid"`
    Wbuf_cb_failed int `json:"wbuf_cb_failed"`
    Name string 

	} `json:"trigger-stats-rate"`
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL7sessionTriggerStatsRate) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL7sessionTriggerStatsRate) getPath() string{
    return "visibility/packet-capture/global-templates/template/" +p.Inst.Name + "/trigger-sys-obj-stats-change/slb-l7session/trigger-stats-rate"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL7sessionTriggerStatsRate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL7sessionTriggerStatsRate::Post")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL7sessionTriggerStatsRate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL7sessionTriggerStatsRate::Get")
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
func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL7sessionTriggerStatsRate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL7sessionTriggerStatsRate::Put")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL7sessionTriggerStatsRate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbL7sessionTriggerStatsRate::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
