package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/mmcdole/gofeed"
	"gopkg.in/ini.v1"
)

type MessageBody struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type Message struct {
	To       string        `json:"to"`
	Messages []MessageBody `json:"messages"`
}

func main() {
	// プロパティファイルの読み込み
	prop, _ := ini.Load("./resources/props.ini")

	// RSS取得
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(prop.Section("rss").Key("url").String())
	items := feed.Items

	// 取得したRSSをメッセージ送信用のテキストに設定
	message := Message{}
	message.To = prop.Section("bot").Key("targetuser").String()
	message.Messages = append(message.Messages, MessageBody{"text", "おはようございます。今日の新着情報です。\n\n"})
	rssCount, _ := prop.Section("rss").Key("notice_count").Int()
	for index, item := range items {
		message.Messages[0].Text += item.Title + "\n" + item.Link + "\n\n"
		if index == rssCount {
			// 指定件数読み込んだらループ終了
			break
		}
	}

	// HTTPリクエスト生成
	reqJSON, _ := json.Marshal(&message)
	reqBodyStr := bytes.NewBuffer(reqJSON)
	req, err := http.NewRequest("POST", prop.Section("bot").Key("url").String(), reqBodyStr)
	req.Header.Set("Authorization", "Bearer "+prop.Section("bot").Key("token").String())
	req.Header.Set("Content-Type", "application/json")

	// HTTPリクエスト送信（正常に送信されればLINE通知される）
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
