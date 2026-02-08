package services

import (
	"errors"
	"fmt"
	"kasir-api/models"
	"kasir-api/repositories"
	"time"
)

type TransactionService struct {
	repo *repositories.TransactionRepository
}

func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Checkout(items []models.CheckoutItem, useLock bool) (*models.Transaction, error) {
	return s.repo.CreateTransaction(items)
}

// ReportHariIni returns sales summary for today (server local time).
func (s *TransactionService) ReportHariIni() (*models.SalesSummaryResponse, error) {
	now := time.Now()
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	end := start.AddDate(0, 0, 1)
	return s.repo.GetSalesSummary(start, end)
}

// ReportRange returns sales summary for the date range [startDate, endDate] (inclusive).
// Dates must be in YYYY-MM-DD format; both are required and startDate must be <= endDate.
func (s *TransactionService) ReportRange(startDate, endDate string) (*models.SalesSummaryResponse, error) {
	if startDate == "" || endDate == "" {
		return nil, errors.New("start_date and end_date are required")
	}
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start_date: %w", err)
	}
	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, fmt.Errorf("invalid end_date: %w", err)
	}
	if start.After(end) {
		return nil, errors.New("start_date must be before or equal to end_date")
	}
	start = time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location())
	end = time.Date(end.Year(), end.Month(), end.Day(), 0, 0, 0, 0, end.Location()).AddDate(0, 0, 1)
	return s.repo.GetSalesSummary(start, end)
}