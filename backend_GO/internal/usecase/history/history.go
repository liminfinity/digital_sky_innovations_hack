package history

import (
	"context"
	"fmt"

	"github.com/kurochkinivan/commit_history/internal/entity"
)

type ReportSaver interface {
	Save(ctx context.Context, file []byte) (string, error)
	GetAll(ctx context.Context) ([]entity.HistoryContent, error)
	GetByName(ctx context.Context, filename string) ([]byte, error)
}

type History struct {
	reportSaver ReportSaver
}

func New(reportSaver ReportSaver) *History {
	return &History{
		reportSaver: reportSaver,
	}
}

func (h *History) SaveReport(ctx context.Context, file []byte) (string, error) {
	const op = "history.SaveReport"

	url, err := h.reportSaver.Save(ctx, file)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return url, nil
}

func (h *History) GetReports(ctx context.Context) ([]string, error) {
	const op = "history.GetReports"

	contents, err := h.reportSaver.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	urls := make([]string, 0, len(contents))
	for _, content := range contents {
		urls = append(urls, content.Name)
	}

	return urls, nil
}

func (h *History) GetReportByName(ctx context.Context, filename string) ([]byte, error) {
	const op = "history.GetReportByName"

	contents, err := h.reportSaver.GetByName(ctx, filename)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return contents, nil
}
