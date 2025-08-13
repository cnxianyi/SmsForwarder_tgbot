package sms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var BaseUrl string

type BaseReq struct {
	Data      Data   `json:"data"`
	Timestamp int64  `json:"timestamp"`
	Sign      string `json:"sign"`
}

type Data struct {
	SimSlot     int    `json:"sim_slot"`
	PhoneNumber string `json:"phone_numbers"`
	MsgContent  string `json:"msg_content"`
	Type        int    `json:"type"`
	PageNum     int    `json:"page_num"`
	PageSize    int    `json:"page_size"`
	Keyword     string `json:"keyword"`
	Phone       string `json:"phone_number"`
}

type Response struct {
	Timestamp int             `json:"timestamp"`
	Code      int             `json:"code"`
	Msg       string          `json:"msg"`
	Data      json.RawMessage `json:"data"` // 使用 json.RawMessage
}

func Post(url string, data Data) (string, error) {

	t, s := GetSign()
	j := BaseReq{
		Data:      data,
		Timestamp: t,
		Sign:      s,
	}

	jsonData, err := json.Marshal(j)
	if err != nil {
		return "", err
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", BaseUrl+url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("创建请求失败: %w", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("发送请求失败: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应失败: %w", err)
	}

	// 检查 HTTP 状态码
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("服务器返回非正常状态码: %d, 响应体: %s", resp.StatusCode, string(body))
	}

	var res Response
	err = json.Unmarshal(body, &res)
	if err != nil {
		return "", fmt.Errorf("解析响应体失败: %w", err)
	}

	// 新增代码：使用 json.MarshalIndent 格式化 data 字段
	formattedData, err := json.MarshalIndent(res.Data, "", "    ")
	if err != nil {
		return "", fmt.Errorf("格式化 data 字段失败: %w", err)
	}

	message := fmt.Sprintf("```json\n%s\n```", string(formattedData))

	// 返回格式化后的 JSON 字符串
	return message, nil
}
