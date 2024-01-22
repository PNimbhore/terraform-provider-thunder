

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6SctpRateLimitDestination struct {
	Inst struct {

    Ip string `json:"ip"`
    RateLimit int `json:"rate-limit"`
    Uuid string `json:"uuid"`

	} `json:"destination"`
}

func (p *Cgnv6SctpRateLimitDestination) GetId() string{
    return p.Inst.Ip
}

func (p *Cgnv6SctpRateLimitDestination) getPath() string{
    return "cgnv6/sctp/rate-limit/destination"
}

func (p *Cgnv6SctpRateLimitDestination) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6SctpRateLimitDestination::Post")
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

func (p *Cgnv6SctpRateLimitDestination) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6SctpRateLimitDestination::Get")
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
func (p *Cgnv6SctpRateLimitDestination) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6SctpRateLimitDestination::Put")
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

func (p *Cgnv6SctpRateLimitDestination) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6SctpRateLimitDestination::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
