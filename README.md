# golang_get_rss_feeds
任意のサイトのRSSフィードを取得・LINE通知するアプリです。

## 使用技術
___
* 言語  
 GO  
* 外部ライブラリ  
 gofeed   
 gopkg
* エディター  
Visual Studio Code  
※任意のものでOKです。

## 動かすための準備
___
* GOのインストール  
※MacのHomebrewでのインストール手順です。  
brew install go
* gofeedのインストール  
go get github.com/mmcdole/gofeed
* gopkgのインストール  
go get gopkg.in/ini.v
* VSCodeにプロジェクトをインポート  
※この際の注意点として、プロジェクトは$GOPATH配下に配置する必要がある。  
※例えば、Githubからクローンしてきたプロジェクトであれば、  
$GOPATH/src/github.com/{gitのusername}/プロジェクト名  
main.goを開きF5キーで実行 
 * LINE Developer登録し、ボット作成  
 [詳しくはこちら](https://developers.line.biz/ja/docs/messaging-api/building-bot/)

