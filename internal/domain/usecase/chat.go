//go:generate mockgen -source=chat.go -destination=chat_mock.go -package=usecase

// コメントを適宜つける

// contextを使って、タイムアウト（30秒ぐらい）とか、traceID（uuid.New()）、sessionID、time.Now()とか渡したい
// `ctx.WithValue(ctx, "traceID", "1234")` みたいな感じで
// `context.Background()` で初期化して、`ctx := context.WithValue(context.Background(), "traceID", "1234")` みたいな感じで
// ログ出力するときに、`slog.Error(ctx, "chat successful")` みたいな感じでcontextの中身を表示する。そうすることでログを追いやすくなる
// contextにドメイン知識を持たせることはない
// contextの中身を変更して新たにそのcontextを引き回したいとかじゃない限り、contextを返すことはない

// ログをきちんと書くようにする。Info, Error, Warn, Debug, Trace, Fatal, Panic とか
// loggerの設定もきちんとやる。環境変数を使って、development, production, staging とかでログの出力先を変えるとか

// 認証機能を実装する

// フロントエンドはRemixを使う

// クエリを実行するときにデータベースが変更されないようにする（そもそもこのアプリケーションに書き込み権限を与えない？dryrunして事前に検知する？）

package usecase

import (
	"github.com/rihib/querychat/internal/domain/entity"
)

type ChatUsecase struct {
	llm  LLM
	repo ChatRepository
}

type LLM interface {
	Ask(prompt entity.OptimizedPrompt) (*entity.LLMOutput, error)
}

type ChatRepository interface {
	Exec(output entity.LLMOutput) ([]map[string]interface{}, error)
}

func NewChatUsecase(llm LLM, repo ChatRepository) *ChatUsecase {
	return &ChatUsecase{llm: llm, repo: repo}
}

func (cu *ChatUsecase) Ask(prompt entity.OptimizedPrompt) (*entity.LLMOutput, error) {
	return cu.llm.Ask(prompt)
}

func (cu *ChatUsecase) Exec(output entity.LLMOutput) ([]map[string]interface{}, error) {
	return cu.repo.Exec(output)
}
