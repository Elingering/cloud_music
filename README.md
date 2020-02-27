# cloud_music
网易云音乐爬虫

项目有三个分支master，concurrentEngine，distributed分别对应单机版，并发版，分布式版。
项目从https://music.163.com//discover/artist 开始爬取，依次获取歌手分类列表，歌手列表，歌曲列表，歌曲热评。
爬虫使用广度优先算法，代码中限制了每个列表的广度。

## master
支持爬取内容打印到终端或保存到~/data目录
## concurrent
爬取内容使用es存储

`
cd cloud_music
go run main.go
`
## distributed
爬取内容使用es存储
开启saveer服务：

`
cd cloud_music/distributed/persist/server
go run itemsaver.go -port=1234
`

开启worker服务：

`
cd cloud_music/distributed/worker/server
go run worker.go -port=9000
开启多个worker
cd cloud_music/distributed/worker/server
go run worker.go -port=9001
`

启动爬虫：

`
cd cloud_music
go run main.go -itemsaver_host=":1234" -worker_hosts=":9000,:9001"
`
