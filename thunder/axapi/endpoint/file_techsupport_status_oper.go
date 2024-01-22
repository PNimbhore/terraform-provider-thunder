

package endpoint
import (
        "io/ioutil"
        "os"
        "net/http"
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FileTechsupportStatusOper struct {
	Inst struct {

    Oper FileTechsupportStatusOperOper `json:"oper"`
    FileContent []byte `json:"-"`
    File string `json:"file"`
    FileHandle string `json:"file-handle"`

	} `json:"oper"`
}
type DeleteFileTechsupportStatusOper struct {
        Inst struct {
            FileName string `json:"filename"`
        } `json:"oper"`
}


type FileTechsupportStatusOperOper struct {
    Message string `json:"message"`
}

func (p *FileTechsupportStatusOper) GetId() string{
    return "1"
}

func (p *FileTechsupportStatusOper) getPath() string{
    return "file/techsupport/status/oper"
}

func (p *FileTechsupportStatusOper) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileTechsupportStatusOper::Post")
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
        s := &FileTechsupportStatusOper{}
            s.Inst.Oper = p.Inst.Oper
            s.Inst.FileContent = data
            s.Inst.File = p.Inst.File
            s.Inst.FileHandle =  p.Inst.File
        _, err := axapi.NormalizeMultipartObject(http.MethodPost, p.getPath(), s.Inst.File, s.Inst.FileContent, s, headers, host, logger)
    return err
}

func (p *FileTechsupportStatusOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileTechsupportStatusOper::Get")
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
func (p *FileTechsupportStatusOper) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileTechsupportStatusOper::Put")
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

func (p *FileTechsupportStatusOper) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileTechsupportStatusOper::Delete")
    headers := axapi.GenRequestHeader(authToken)
        s := &DeleteFileTechsupportStatusOper{}
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
