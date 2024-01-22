

package endpoint
import (
        "io/ioutil"
        "os"
        "net/http"
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FileDnssecDsOper struct {
	Inst struct {

    Oper FileDnssecDsOperOper `json:"oper"`
    FileContent []byte `json:"-"`
    File string `json:"file"`
    FileHandle string `json:"file-handle"`

	} `json:"oper"`
}
type DeleteFileDnssecDsOper struct {
        Inst struct {
            FileName string `json:"filename"`
        } `json:"oper"`
}


type FileDnssecDsOperOper struct {
    FileList []FileDnssecDsOperOperFileList `json:"file-list"`
}


type FileDnssecDsOperOperFileList struct {
    File string `json:"file"`
    Parent string `json:"parent"`
}

func (p *FileDnssecDsOper) GetId() string{
    return "1"
}

func (p *FileDnssecDsOper) getPath() string{
    return "file/dnssec-ds/oper"
}

func (p *FileDnssecDsOper) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileDnssecDsOper::Post")
    headers := axapi.GenRequestHeader(authToken)
        f, error := os.Open(p.Inst.FileHandle)
        if error != nil {
            logger.Println("Failed to open a file: ", error)
            return error
        }
        data, error := ioutil.ReadAll(f)
        defer f.Close()
        if error != nil {
            logger.Println("Failed to read file: ", error)
            return error
        }
        s := &FileDnssecDsOper{}
            s.Inst.Oper = p.Inst.Oper
            s.Inst.FileContent = data
            s.Inst.File = p.Inst.File
            s.Inst.FileHandle =  p.Inst.File
        _, err := axapi.NormalizeMultipartObject(http.MethodPost, p.getPath(), s.Inst.File, s.Inst.FileContent, s, headers, host, logger)
    return err
}

func (p *FileDnssecDsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileDnssecDsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, "file/oper/oper", "", nil, headers, logger)
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
func (p *FileDnssecDsOper) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileDnssecDsOper::Put")
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

func (p *FileDnssecDsOper) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileDnssecDsOper::Delete")
    headers := axapi.GenRequestHeader(authToken)
        s := &DeleteFileDnssecDsOper{}
        s.Inst.FileName = p.Inst.File
        payloadBytes, err := json.Marshal(s)
        logger.Println(s)
        if err != nil {
            logger.Println("json.Marshal() failed with error", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, "delete/oper", payloadBytes, headers, logger)
    return err
}
