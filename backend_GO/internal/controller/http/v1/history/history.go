package history

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
	mdw "github.com/kurochkinivan/commit_history/internal/controller/http/v1/middleware"
)

type HistoryUseCase interface {
	SaveReport(ctx context.Context, file []byte) (string, error)
	GetReports(ctx context.Context) ([]string, error)
	GetReportByName(ctx context.Context, filename string) ([]byte, error)
}

type Controller struct {
	historyUseCase HistoryUseCase
}

func New(historyUseCase HistoryUseCase) *Controller {
	return &Controller{
		historyUseCase: historyUseCase,
	}
}

func (c *Controller) Register(r *httprouter.Router) {
	r.GET("/api/v1/history-service/history", mdw.LoggingMiddleware(c.getAllReports))
	r.GET("/api/v1/history-service/history/:filename", mdw.LoggingMiddleware(c.getReportByName))
	r.POST("/api/v1/history-service/history", mdw.LoggingMiddleware(c.saveReport))
}

type saveReportResponse struct {
	Url string `json:"report_url"`
}

func (c *Controller) saveReport(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	reportData, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	url, err := c.historyUseCase.SaveReport(r.Context(), reportData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(saveReportResponse{
		Url: url,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}

type getAllReportsResponse struct {
	Urls []string `json:"urls"`
}

func (c *Controller) getAllReports(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	urls, err := c.historyUseCase.GetReports(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(getAllReportsResponse{
		Urls: urls,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}

func (c *Controller) getReportByName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	filename := ps.ByName("filename")
	if filename == "" {
		http.Error(w, "filename is empty", http.StatusBadRequest)
		return
	}

	reportData, err := c.historyUseCase.GetReportByName(r.Context(), filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(reportData)
}
