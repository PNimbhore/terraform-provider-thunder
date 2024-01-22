

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemCpuDataCpuOper struct {
    
    Oper SystemCpuDataCpuOperOper `json:"oper"`

}
type DataSystemCpuDataCpuOper struct {
    DtSystemCpuDataCpuOper SystemCpuDataCpuOper `json:"data-cpu"`
}


type SystemCpuDataCpuOperOper struct {
    NumberOfCpu int `json:"number-of-cpu"`
    NumberOfDataCpu int `json:"number-of-data-cpu"`
    CpuUsage []SystemCpuDataCpuOperOperCpuUsage `json:"cpu-usage"`
}


type SystemCpuDataCpuOperOperCpuUsage struct {
    CpuId int `json:"cpu-id"`
    Sec1 int `json:"sec1"`
    Sec5 int `json:"sec5"`
    Sec10 int `json:"sec10"`
    Sec30 int `json:"sec30"`
    Sec60 int `json:"sec60"`
    DcpuStr string `json:"dcpu-str"`
}

func (p *SystemCpuDataCpuOper) GetId() string{
    return "1"
}

func (p *SystemCpuDataCpuOper) getPath() string{
    return "system-cpu/data-cpu/oper"
}

func (p *SystemCpuDataCpuOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemCpuDataCpuOper,error) {
logger.Println("SystemCpuDataCpuOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemCpuDataCpuOper
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
