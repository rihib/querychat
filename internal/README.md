# メモ

## タスクリスト

### リファクタ

#### その他

- [ ] よりわかりやすいコードにする
- [ ] コメントを適宜つける
- [ ] ファイルパスを相対パスにする
- [ ] 値渡しと参照渡しの使い分けが混在しているので、統一する（規則を決める）
- [ ] VDには実行したクエリも入れて返したい
- [ ] 入力サイズの確認
- [ ] InfoでLLMへの入力応答を出す

#### context

- [ ] contextを使って、タイムアウト（30秒ぐらい）とか、traceID（uuid.New()）、sessionID、time.Now()とか渡したい
  - `ctx.WithValue(ctx, "traceID", "1234")` みたいな感じで
  - `context.Background()` で初期化して、`ctx := context.WithValue(context.Background(), "traceID", "1234")` みたいな感じで
  - contextにドメイン知識を持たせることはない
  - contextの中身を変更して新たにそのcontextを引き回したいとかじゃない限り、contextを返すことはない

#### ログ

- [ ] ログをきちんと書くようにする。Info, Error, Warn, Debug, Trace, Fatal, Panic とか
- [ ] loggerの設定もきちんとやる。環境変数を使って、development, production, staging とかでログの出力先を変えるとか
- [ ] ログを見やすくする
  - [Creating a pretty console logger using Go's slog package](https://dusted.codes/creating-a-pretty-console-logger-using-gos-slog-package)
  - [Logging in Go with Slog: The Ultimate Guide](https://betterstack.com/community/guides/logging/logging-in-go/)
  - [Pretty handler for structured Logging with slog](https://github.com/go-slog-handler/slog-handler)
  - [Go公式の構造化ロガー（予定）のslogの出力を見やすくしてみる](https://zenn.dev/mizutani/articles/golang-clog-handler)
- [ ] ログ出力するときに、`slog.Error(ctx, "chat successful")` みたいな感じでcontextの中身を表示する。そうすることでログを追いやすくなる

### API

- [ ] REST API化する（自動生成せずに標準ライブラリのみで手書き）
- [ ] OpenAPI specを修正する

### 機能追加

- [ ] 認証機能を実装する
- [ ] Remixでフロントエンドを作る
- [ ] GitHub ActionsでCI/CDパイプラインを作る
- [ ] クエリを実行するときにデータベースが変更されないようにする（そもそもこのアプリケーションに書き込み権限を与えない？dryrunして事前に検知する？）
- [ ] MySQL, PostgreSQLにも対応する
