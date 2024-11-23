package spreadsheet

import (
	"context"
	"errors"
	"log"

	"bsearch/backend/config"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// SpreadsheetService は Google スプレッドシートからデータを取得するサービス
type SpreadsheetService struct {
	sheetsService *sheets.Service
}

// NewSpreadsheetService は SpreadsheetService を初期化する
func NewSpreadsheetService() (*SpreadsheetService, error) {
	// config パッケージから認証ファイルのパスを取得
	credentialsPath := config.GetCredentialsPath()

	// Google Sheets API クライアントを作成
	ctx := context.Background()
	sheetsService, err := sheets.NewService(ctx, option.WithCredentialsFile(credentialsPath))
	if err != nil {
		return nil, errors.New("failed to create Google Sheets service: " + err.Error())
	}

	return &SpreadsheetService{sheetsService: sheetsService}, nil
}

// GetSpreadsheetData はスプレッドシートから指定範囲のデータを取得する
func (s *SpreadsheetService) GetSpreadsheetData(spreadsheetID, readRange string) ([][]interface{}, error) {
	resp, err := s.sheetsService.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		return nil, errors.New("failed to retrieve data from spreadsheet: " + err.Error())
	}

	if resp.Values == nil || len(resp.Values) == 0 {
		log.Println("No data found in the spreadsheet")
		return [][]interface{}{}, nil
	}

	return resp.Values, nil
}
