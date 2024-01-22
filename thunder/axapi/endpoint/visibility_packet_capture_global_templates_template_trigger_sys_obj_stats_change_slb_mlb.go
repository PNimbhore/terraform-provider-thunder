

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMlb struct {
	Inst struct {

    TriggerStatsInc VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMlbTriggerStatsInc2051 `json:"trigger-stats-inc"`
    TriggerStatsRate VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMlbTriggerStatsRate2052 `json:"trigger-stats-rate"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"slb-mlb"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMlbTriggerStatsInc2051 struct {
    Mlb_dcmsg_error int `json:"mlb_dcmsg_error"`
    Uuid string `json:"uuid"`
}


type VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMlbTriggerStatsRate2052 struct {
    ThresholdExceededBy int `json:"threshold-exceeded-by" dval:"5"`
    Duration int `json:"duration" dval:"60"`
    Mlb_dcmsg_error int `json:"mlb_dcmsg_error"`
    Uuid string `json:"uuid"`
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMlb) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMlb) getPath() string{
    return "visibility/packet-capture/global-templates/template/" +p.Inst.Name + "/trigger-sys-obj-stats-change/slb-mlb"
}

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMlb) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMlb::Post")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMlb) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMlb::Get")
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
func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMlb) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMlb::Put")
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

func (p *VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMlb) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureGlobalTemplatesTemplateTriggerSysObjStatsChangeSlbMlb::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
