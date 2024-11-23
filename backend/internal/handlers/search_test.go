package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuzzySearch(t *testing.T) {
	// テストケース
	tests := []struct {
		name     string
		query    string
		corpus   []string
		expected []string
	}{
		{
			name:     "部分一致するデータがある場合",
			query:    "テス",
			corpus:   []string{"テストケース", "サンプル", "テストドライブ"},
			expected: []string{"テストケース", "テストドライブ"},
		},
		{
			name:     "一致するデータがない場合",
			query:    "テスト",
			corpus:   []string{"サンプル", "データ"},
			expected: []string{},
		},
		{
			name:     "空のクエリを指定した場合",
			query:    "",
			corpus:   []string{"サンプル", "データ"},
			expected: []string{"サンプル", "データ"},
		},
	}

	// テスト実行
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FuzzySearch(tt.query, tt.corpus)
			assert.ElementsMatch(t, tt.expected, result, "テスト失敗: %s", tt.name)
		})

	}
}
