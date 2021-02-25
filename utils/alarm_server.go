package utils

//n9e告警拒收数据
import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	ETC_SETTLEMENT_ORIGINAL_MSG_REFUSE_WARN       = "etcsettlement.originalmessage.refuse.warn"       //拒收告警
	ETC_SETTLEMENT_CLEAR_MSG_CLEARERROR_WARN      = "etcsettlement.clearmessage.clearerror.warn"      //清分处理失败告警
	ETC_SETTLEMENT_CLEAR_MSG_CLEARCLEARERROR_WARN = "etcsettlement.clearmessage.clearcheckerror.warn" //清分核对失败告警
)

type N9EPostData struct {
	Metric    string `json:"metric"`    //指标
	Endpoint  string `json:"endpoint"`  //端点
	Tags      string `json:"tags"`      //标签
	Value     int64  `json:"value"`     //值
	Timestamp int64  `json:"timestamp"` //时间戳
	Step      int64  `json:"step"`      //步
}

type N9EPostDataPost struct {
	DataArray []N9EPostData `json:"data"`
}

func (n9e *N9EPostDataPost) PostData(metric string, value int64, tags map[string]string) {

	sg_deviceid, getDeviderr := GetDeviceID()
	if getDeviderr != nil {
		log.Error(getDeviderr)
	}
	log.Println("sg_deviceid", sg_deviceid) //188EFFA448ED8751  \08AAA8D5EC80A037
	var (
		jsonData   string
		itemValue  N9EPostData
		tagsString string
	)

	if len(sg_deviceid) != 16 {
		return
	}
	for k, v := range tags {
		k = strings.TrimSpace(k)
		k = strings.TrimSpace(k)
		if tagsString != "" {
			tagsString += ","
		}
		tagsString += fmt.Sprintf("%v=%v", k, v)
	}
	itemValue.Timestamp = time.Now().Unix()
	itemValue.Step = 20
	itemValue.Tags = tagsString
	itemValue.Metric = metric
	itemValue.Value = value
	itemValue.Endpoint = sg_deviceid

	var array N9EPostDataPost
	array.DataArray = append(array.DataArray, itemValue)
	jsonData = GetJSONArray(&array)
	if jsonData == "" {
		return
	}
	//http://etcgateway.jstxb.com:5810/api/transfer/push
	var res, err = HttpPostJson("http://etcgateway.jstxb.com:5810/api/transfer/push", jsonData, time.Second)
	if err != nil {
		log.Error(err)
		return
	}
	log.Println(string(res))
	return
}

func GetJSONArray(n9e *N9EPostDataPost) (findString string) {
	var err error
	var bsData, _ = json.Marshal(n9e.DataArray)
	var exp *regexp.Regexp
	exp, err = regexp.Compile("\\[{1}.*\\]{1}")
	if err != nil {
		return
	}
	var myfindString = exp.FindStringSubmatch(string(bsData))
	if len(myfindString) == 0 {
		return
	}

	if len(myfindString) > 0 {
		findString = myfindString[0]
	}
	return
}

func HttpPostJson(url string, jsonParam string, timeout time.Duration) (res []byte, err error) {
	reqest, err := http.NewRequest("POST", url, bytes.NewReader([]byte(jsonParam)))
	if err != nil {
		return
	}
	reqest.Header.Set("Content-Type", "application/json; encoding=utf-8")
	c := http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: timeout,
			}).DialContext,
		},
	}
	resp, err := c.Do(reqest)
	if err != nil {
		return
	}
	var body []byte
	if resp.StatusCode == 200 {
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}
		defer func() {
			_ = resp.Body.Close()
		}()

	} else {
		err = errors.New(fmt.Sprintf("StatusCode=%v error", resp.StatusCode))
	}
	res = body
	return
}

//告警的使用
func UsePostData(jsbh, jsts int) {
	var n9e N9EPostDataPost
	tags := make(map[string]string, 0)

	tags["拒收原始消息包号"] = strconv.Itoa(jsbh)
	tags["拒收条数"] = strconv.Itoa(jsts)

	tzsj := time.Now().Format("2006-01-02T15:04:05")
	tags["通知时间"] = tzsj

	n9e.PostData(ETC_SETTLEMENT_ORIGINAL_MSG_REFUSE_WARN, 1, tags)
}

//清分核对失败告警
func ClearErrorUsePostData(qfbh int, qfmbr string) {
	var n9e N9EPostDataPost
	tags := make(map[string]string, 0)
	tags["清分消息包号"] = strconv.Itoa(qfbh)
	tags["清分目标日"] = qfmbr
	tzsj := time.Now().Format("2006-01-02T15:04:05")
	tags["通知时间"] = tzsj

	n9e.PostData(ETC_SETTLEMENT_CLEAR_MSG_CLEARCLEARERROR_WARN, 1, tags)
}
