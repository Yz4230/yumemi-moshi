# 「株式会社ゆめみ」さんのサーバーサイド応募者向けの模試を解いてみました✏️

株式会社ゆめみ：
[ゆめみ | DX,SoE,UX/UI,アジャイル,DevOps,グロースハック](https://www.yumemi.co.jp)

問題文URL：
[【新卒・中途採用】サーバーサイドエンジニア応募者向けの模試 | ゆめみ](https://www.yumemi.co.jp/serverside_recruit)

## テスト方法
```bash
git clone https://github.com/Yz4230/yumemi-moshi.git
cd yumemi-moshi
go test -v ./...
```

## テスト結果
```bash
=== RUN   TestValidateHeader
--- PASS: TestValidateHeader (0.00s)
=== RUN   TestValidatePlayerID
--- PASS: TestValidatePlayerID (0.00s)
=== RUN   TestParseCSV
--- PASS: TestParseCSV (0.00s)
PASS
ok  	github.com/Yz4230/yumemi-moshi	0.075s
```
