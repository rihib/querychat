package app_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rihib/querychat/internal/app"
	"github.com/rihib/querychat/internal/domain/entity"
	"github.com/rihib/querychat/internal/domain/usecase"
	"github.com/stretchr/testify/assert"
)

func TestChat(t *testing.T) {
	type want struct {
		vd  *entity.VisualizableData
		err error
	}
	cases := []struct {
		name   string
		prompt string
		want   want
		setup  func() (*entity.LLMOutput, []map[string]interface{}, error)
	}{
		{
			name:   "success",
			prompt: "What are the total purchases per user?",
			want: want{
				err: nil,
			},
			setup: func() (*entity.LLMOutput, []map[string]interface{}, error) {
				query := "SELECT user_name, SUM(amount) AS total_amount FROM purchases GROUP BY user_id, user_name"
				chart := `{"type": "bar", "x": "UserName", "y": "TotalAmount"}`
				output, err := entity.NewLLMOutput(query, chart)
				if err != nil {
					return nil, nil, err
				}
				datas := []map[string]interface{}{
					{"UserName": "Alice", "TotalAmount": 100},
					{"UserName": "Bob", "TotalAmount": 200},
					{"UserName": "Charlie", "TotalAmount": 300},
				}
				return output, datas, nil
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()

			// Arrange
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockLLM := usecase.NewMockLLM(ctrl)
			mockRepo := usecase.NewMockChatRepository(ctrl)

			if tt.want.err == nil {
				output, datas, err := tt.setup()
				if err != nil {
					t.Fatalf("failed to setup: %v", err)
				}
				wantVD, err := entity.NewVisualizableData(datas, *output)
				if err != nil {
					t.Fatalf("failed to create visualizable data: %v", err)
				}
				tt.want.vd = wantVD
				mockLLM.EXPECT().Ask(gomock.Any()).Return(output, nil)
				mockRepo.EXPECT().Exec(*output).Return(datas, nil)
			}

			qcConfig, err := entity.NewQueryChatConfig(tt.prompt, "system", "user", "db", "schema")
			if err != nil {
				t.Fatalf("failed to create query chat config: %v", err)
			}

			// Act
			vd, err := app.Chat(*qcConfig, mockLLM, mockRepo)

			// Assert
			assert.Equal(t, tt.want.vd, vd)
			assert.Equal(t, tt.want.err, err)
		})
	}
}
