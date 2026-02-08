package handlers

import (
	"encoding/json"
	"net/http"

	"kasir-api/services"
)

type ReportHandler struct {
	service *services.TransactionService
}

func NewReportHandler(service *services.TransactionService) *ReportHandler {
	return &ReportHandler{service: service}
}

// HandleReportHariIni - GET /api/report/hari-ini
func (h *ReportHandler) HandleReportHariIni(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	summary, err := h.service.ReportHariIni()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summary)
}

// HandleReport - GET /api/report?start_date=YYYY-MM-DD&end_date=YYYY-MM-DD
func (h *ReportHandler) HandleReport(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	summary, err := h.service.ReportRange(startDate, endDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summary)
}
