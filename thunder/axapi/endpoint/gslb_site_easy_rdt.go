

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbSiteEasyRdt struct {
	Inst struct {

    AgingTime int `json:"aging-time" dval:"10"`
    BindGeoloc int `json:"bind-geoloc"`
    IgnoreCount int `json:"ignore-count" dval:"5"`
    Ipv6Mask int `json:"ipv6-mask" dval:"128"`
    Limit int `json:"limit" dval:"16383"`
    Mask string `json:"mask" dval:"/32"`
    Overlap int `json:"overlap"`
    RangeFactor int `json:"range-factor" dval:"25"`
    SmoothFactor int `json:"smooth-factor" dval:"10"`
    Uuid string `json:"uuid"`
    SiteName string 

	} `json:"easy-rdt"`
}

func (p *GslbSiteEasyRdt) GetId() string{
    return "1"
}

func (p *GslbSiteEasyRdt) getPath() string{
    return "gslb/site/" +p.Inst.SiteName + "/easy-rdt"
}

func (p *GslbSiteEasyRdt) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbSiteEasyRdt::Post")
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

func (p *GslbSiteEasyRdt) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbSiteEasyRdt::Get")
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
func (p *GslbSiteEasyRdt) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbSiteEasyRdt::Put")
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

func (p *GslbSiteEasyRdt) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbSiteEasyRdt::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
