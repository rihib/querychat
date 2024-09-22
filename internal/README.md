# メモ

## タスクリスト

### 改善

#### ログ

- [x] ログをきちんと書くようにする
- [x] loggerの設定
- [ ] ユニットテストの場合はmain関数のloggerの設定が適用できていないので修正する
- [ ] contextの中身をログと一緒に表示するようにする
  - [Go1.21 log/slogパッケージ超入門](https://zenn.dev/88888888_kota/articles/7e97ff874083cf)
  - [slog を触る(Group, Context)](https://zenn.dev/kyoshigai/articles/bc90cc776dea2c)
- [ ] ログを見やすくする
  - [Creating a pretty console logger using Go's slog package](https://dusted.codes/creating-a-pretty-console-logger-using-gos-slog-package)
  - [Logging in Go with Slog: The Ultimate Guide](https://betterstack.com/community/guides/logging/logging-in-go/)
  - [Pretty handler for structured Logging with slog](https://github.com/go-slog-handler/slog-handler)
  - [Go公式の構造化ロガー（予定）のslogの出力を見やすくしてみる](https://zenn.dev/mizutani/articles/golang-clog-handler)

#### リファクタ

- [ ] 入力サイズの確認
- [ ] 値渡しと参照渡しの使い分けが混在しているので、統一する（規則を決める）
- [ ] Google Go Style GuideやEffective Goに沿った形に修正する

#### テスト

- [ ] ユニットテストのカバレッジを100%にする
- [ ] ５種類のテストケースを用意すると良い。サンプルテストケース、コーナーテストケース、最大テストケース、最小テストケース、ランダムテストケース（毎回ランダム）

### API

- [ ] OpenAPI specを修正する
- [ ] REST API化する（[oapi-codegen](https://github.com/oapi-codegen/oapi-codegen)）

### 機能追加

- [ ] 認証機能を実装する
  - セッションIDはredisに保存する
- [ ] Remixでフロントエンドを作る
  - ユーザーが使うようと、管理者が使うように分ける。管理者はAPIキーの設定、スキーマファイルのアップロード、ユーザーの管理などを行う
- [ ] GitHub ActionsでCI/CDパイプラインを作る
- [ ] クエリを実行するときにデータベースが変更されないようにする（そもそもこのアプリケーションに書き込み権限を与えない？dryrunして事前に検知する？）
- [ ] MySQL, PostgreSQLにも対応する（ローカルでdocker-composeを使ってテスト環境を立ち上げられるようにする）
- [ ] 表現力を増やす
  - 現状はx,y軸を持つデータのみを扱えるが、それ以外に円グラフとか表とかも扱えるようにしたい
    - 例えば「どのようなデータ構造になっているか」という質問を受けたときに現状はそれを表現できない
    - その他にも標準偏差とかそういうBIツールが扱うような複雑なデータ分析とかも表現できるようにしたい
    - データ分析を民主化するというテーマを掲げて、専門性を持ったデータ分析者がいなくても、誰でも本格的なデータ分析ができるようにしたい
