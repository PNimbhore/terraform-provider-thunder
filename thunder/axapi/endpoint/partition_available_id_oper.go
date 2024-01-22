

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type PartitionAvailableIdOper struct {
    
    Oper PartitionAvailableIdOperOper `json:"oper"`

}
type DataPartitionAvailableIdOper struct {
    DtPartitionAvailableIdOper PartitionAvailableIdOper `json:"partition-available-id"`
}


type PartitionAvailableIdOperOper struct {
    RangeList []PartitionAvailableIdOperOperRangeList `json:"range-list"`
}


type PartitionAvailableIdOperOperRangeList struct {
    Start string `json:"start"`
    End string `json:"end"`
}

func (p *PartitionAvailableIdOper) GetId() string{
    return "1"
}

func (p *PartitionAvailableIdOper) getPath() string{
    return "partition-available-id/oper"
}

func (p *PartitionAvailableIdOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataPartitionAvailableIdOper,error) {
logger.Println("PartitionAvailableIdOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataPartitionAvailableIdOper
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
