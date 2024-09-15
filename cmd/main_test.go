package main

import (
	"testing"

	"github.com/rihib/querychat/internal/app"
	"github.com/stretchr/testify/assert"
)

func TestChat(t *testing.T) {
	tests := []struct {
		name    string
		prompt  string
		wantErr bool
	}{
		{
			name:    "Valid prompt",
			prompt:  "What are the monthly sales for 2013?",
			wantErr: false,
		},
		{
			name:    "Invalid prompt",
			prompt:  "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vd, err := app.Chat(tt.prompt)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, vd)
			}
		})
	}
}