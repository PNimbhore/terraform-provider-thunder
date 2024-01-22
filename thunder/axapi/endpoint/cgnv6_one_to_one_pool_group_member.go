

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6OneToOnePoolGroupMember struct {
	Inst struct {

    PoolName string `json:"pool-name"`
    Uuid string `json:"uuid"`
    PoolGroupName string 

	} `json:"member"`
}

func (p *Cgnv6OneToOnePoolGroupMember) GetId() string{
    return url.QueryEscape(p.Inst.PoolName)
}

func (p *Cgnv6OneToOnePoolGroupMember) getPath() string{
    return "cgnv6/one-to-one/pool-group/" +p.Inst.PoolGroupName + "/member"
}

func (p *Cgnv6OneToOnePoolGroupMember) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6OneToOnePoolGroupMember::Post")
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

func (p *Cgnv6OneToOnePoolGroupMember) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6OneToOnePoolGroupMember::Get")
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
func (p *Cgnv6OneToOnePoolGroupMember) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6OneToOnePoolGroupMember::Put")
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

func (p *Cgnv6OneToOnePoolGroupMember) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6OneToOnePoolGroupMember::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
