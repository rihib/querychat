# メモ

## タスクリスト

### 改善

#### ログ

- [x] ログをきちんと書くようにする
- [x] loggerの設定
- [ ] ユニットテストの場合はmain関数のloggerの設定が適用できていないので修正する
- [ ] contextに経過時間も入れたい（リクエストを受け取ってからどのぐらいの処理時間が経過したか）
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
  - それぞれのLLMのAPIに合わせて入力サイズの確認を行う
- [ ] 値渡しと参照渡しの使い分けが混在しているので、統一する（規則を決める）
- [ ] Google Go Style GuideやEffective Goに沿った形に修正する
- [ ] READMEを整備する

#### ビルド・テスト

- [x] launch.jsonを整備して実行できるようにする
- [ ] ユニットテストのカバレッジを100%にする
- [ ] ５種類のテストケースを用意すると良い。サンプルテストケース、コーナーテストケース、最大テストケース、最小テストケース、ランダムテストケース（毎回ランダム）
- [ ] ローカルでdocker-composeを使ってテスト環境を立ち上げられるようにする
- [ ] プロダクションビルドをできるようにする
- [ ] mainを保護ブランチにし、PRを作成すると自動生成、リント、ビルド、テストを行うCIをGitHub Actionsで作る
- [ ] バージョンフォーマットはA.B.C
  - メジャーバージョン：プロトタイプの段階では0、本番リリースしたら１、以降は大きく変更を加えた時にのみインクリメントする
  - マイナーバージョン：新しい機能を追加するたびにバージョンを上げる
  - パッチバージョン：修正やリファクタリングを行った時にバージョンを上げる

### API

- [ ] grpc-gatewayを使ってREST APIを提供する（OpenAPIも自動生成する）
  - こうすることでデフォルトのWebクライアントを使うのではなく、マイクロサービスの１つとして組み込むと言ったこともできるようになる
- [ ] 入力サイズの確認
  - 例えば下記のようにする

    ```go
    func handler(w http.ResponseWriter, r *http.Request) {
      // 1MBの制限を設定
      r.Body = http.MaxBytesReader(w, r.Body, 1048576)
      // ここでリクエストボディを処理
    }
    ```

### 機能追加

- [ ] 認証機能を実装する
  - セッションIDはredisに保存する
- [ ] Remixでフロントエンドを作る
  - ユーザーが使うようと、管理者が使うように分ける。管理者はAPIキーの設定、スキーマファイルのアップロード、ユーザーの管理などを行う
- [ ] クエリを実行するときにデータベースが変更されないようにする（そもそもこのアプリケーションに書き込み権限を与えない？dryrunして事前に検知する？）
- [ ] MySQL, PostgreSQLにも対応する
- [ ] 表現力を増やす
  - 現状はx,y軸を持つデータのみを扱えるが、それ以外に円グラフとか表とかも扱えるようにしたい
    - 例えば「どのようなデータ構造になっているか」という質問を受けたときに現状はそれを表現できない
    - その他にも標準偏差とかそういうBIツールが扱うような複雑なデータ分析とかも表現できるようにしたい
    - データ分析を民主化するというテーマを掲げて、専門性を持ったデータ分析者がいなくても、誰でも本格的なデータ分析ができるようにしたい
- [ ] 利用できるLLMモデルを増やす
  - Goを使うのが難しければgRPCを使う
