# メモ

## タスクリスト

### リファクタ

#### コメント

- [ ] コメントを適宜つける

#### context

- [ ] contextを使って、タイムアウト（30秒ぐらい）とか、traceID（uuid.New()）、sessionID、time.Now()とか渡したい
  - `ctx.WithValue(ctx, "traceID", "1234")` みたいな感じで
  - `context.Background()` で初期化して、`ctx := context.WithValue(context.Background(), "traceID", "1234")` みたいな感じで
  - contextにドメイン知識を持たせることはない
  - contextの中身を変更して新たにそのcontextを引き回したいとかじゃない限り、contextを返すことはない

#### ログ

- [ ] ログをきちんと書くようにする。Info, Error, Warn, Debug, Trace, Fatal, Panic とか
- [ ] loggerの設定もきちんとやる。環境変数を使って、development, production, staging とかでログの出力先を変えるとか
- [ ] ログ出力するときに、`slog.Error(ctx, "chat successful")` みたいな感じでcontextの中身を表示する。そうすることでログを追いやすくなる

### API

- [ ] REST API化する（自動生成せずに標準ライブラリのみで手書き）
- [ ] OpenAPI specを修正する

### 機能追加

- [ ] 認証機能を実装する
- [ ] Remixでフロントエンドを作る
- [ ] クエリを実行するときにデータベースが変更されないようにする（そもそもこのアプリケーションに書き込み権限を与えない？dryrunして事前に検知する？）
