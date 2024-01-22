

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosProtectionMultiPuZoneDistribution struct {
	Inst struct {

    CpuThresholdPerEntry int `json:"cpu-threshold-per-entry" dval:"60"`
    CpuThresholdPerPu int `json:"cpu-threshold-per-pu" dval:"80"`
    DistributionMethod string `json:"distribution-method" dval:"traffic-rate"`
    RateKbitThreshold int `json:"rate-kbit-threshold" dval:"150000000"`
    RatePktThreshold int `json:"rate-pkt-threshold" dval:"55000000"`
    Uuid string `json:"uuid"`

	} `json:"multi-pu-zone-distribution"`
}

func (p *DdosProtectionMultiPuZoneDistribution) GetId() string{
    return "1"
}

func (p *DdosProtectionMultiPuZoneDistribution) getPath() string{
    return "ddos/protection/multi-pu-zone-distribution"
}

func (p *DdosProtectionMultiPuZoneDistribution) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosProtectionMultiPuZoneDistribution::Post")
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

func (p *DdosProtectionMultiPuZoneDistribution) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosProtectionMultiPuZoneDistribution::Get")
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
func (p *DdosProtectionMultiPuZoneDistribution) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosProtectionMultiPuZoneDistribution::Put")
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

func (p *DdosProtectionMultiPuZoneDistribution) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosProtectionMultiPuZoneDistribution::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
