

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsInc struct {
	Inst struct {

    JwtAuthorizeFailure int `json:"jwt-authorize-failure"`
    JwtMissingClaim int `json:"jwt-missing-claim"`
    JwtMissingToken int `json:"jwt-missing-token"`
    JwtOtherError int `json:"jwt-other-error"`
    JwtSignatureFailure int `json:"jwt-signature-failure"`
    JwtTokenExpired int `json:"jwt-token-expired"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"trigger-stats-inc"`
}

func (p *VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsInc) GetId() string{
    return "1"
}

func (p *VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsInc) getPath() string{
    return "visibility/packet-capture/object-templates/aam-jwt-authorization-tmpl/" +p.Inst.Name + "/trigger-stats-inc"
}

func (p *VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsInc) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsInc::Post")
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

func (p *VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsInc) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsInc::Get")
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
func (p *VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsInc) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsInc::Put")
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

func (p *VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsInc) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VisibilityPacketCaptureObjectTemplatesAamJwtAuthorizationTmplTriggerStatsInc::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
