

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type DdosZoneTemplateIcmpV6Type struct {
	Inst struct {

    DstCodeOtherRate int `json:"dst-code-other-rate"`
    DstCodeOtherRateAction string `json:"dst-code-other-rate-action"`
    DstCodeOtherRateActionListName string `json:"dst-code-other-rate-action-list-name"`
    IcmpTypeAction string `json:"icmp-type-action"`
    IcmpTypeActionListName string `json:"icmp-type-action-list-name"`
    SrcCodeOtherRate int `json:"src-code-other-rate"`
    SrcCodeOtherRateAction string `json:"src-code-other-rate-action"`
    SrcCodeOtherRateActionListName string `json:"src-code-other-rate-action-list-name"`
    TypeNumber int `json:"type-number"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    V6DstCodeCfg []DdosZoneTemplateIcmpV6TypeV6DstCodeCfg `json:"v6-dst-code-cfg"`
    V6DstRateCfg DdosZoneTemplateIcmpV6TypeV6DstRateCfg `json:"v6-dst-rate-cfg"`
    V6SrcCodeCfg []DdosZoneTemplateIcmpV6TypeV6SrcCodeCfg `json:"v6-src-code-cfg"`
    V6SrcRateCfg DdosZoneTemplateIcmpV6TypeV6SrcRateCfg `json:"v6-src-rate-cfg"`
    IcmpTmplName string 

	} `json:"type"`
}


type DdosZoneTemplateIcmpV6TypeV6DstCodeCfg struct {
    DstCodeNumber int `json:"dst-code-number"`
    DstCodeRate int `json:"dst-code-rate"`
    DstCodeRateActionListName string `json:"dst-code-rate-action-list-name"`
    DstCodeRateAction string `json:"dst-code-rate-action"`
}


type DdosZoneTemplateIcmpV6TypeV6DstRateCfg struct {
    DstTypeRate int `json:"dst-type-rate"`
    DstTypeRateActionListName string `json:"dst-type-rate-action-list-name"`
    DstTypeRateAction string `json:"dst-type-rate-action"`
}


type DdosZoneTemplateIcmpV6TypeV6SrcCodeCfg struct {
    SrcCodeNumber int `json:"src-code-number"`
    SrcCodeRate int `json:"src-code-rate"`
    SrcCodeRateActionListName string `json:"src-code-rate-action-list-name"`
    SrcCodeRateAction string `json:"src-code-rate-action"`
}


type DdosZoneTemplateIcmpV6TypeV6SrcRateCfg struct {
    SrcTypeRate int `json:"src-type-rate"`
    SrcTypeRateActionListName string `json:"src-type-rate-action-list-name"`
    SrcTypeRateAction string `json:"src-type-rate-action"`
}

func (p *DdosZoneTemplateIcmpV6Type) GetId() string{
    return strconv.Itoa(p.Inst.TypeNumber)
}

func (p *DdosZoneTemplateIcmpV6Type) getPath() string{
    return "ddos/zone-template/icmp-v6/" +p.Inst.IcmpTmplName + "/type"
}

func (p *DdosZoneTemplateIcmpV6Type) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateIcmpV6Type::Post")
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

func (p *DdosZoneTemplateIcmpV6Type) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateIcmpV6Type::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
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
func (p *DdosZoneTemplateIcmpV6Type) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateIcmpV6Type::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
    return err
}

func (p *DdosZoneTemplateIcmpV6Type) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateIcmpV6Type::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
