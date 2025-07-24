package service

import "lifedash/repo"

type ExpenseService struct {
	repo *repo.ExpenseRepo
}

func NewExpenseService(repo *repo.ExpenseRepo) *ExpenseService {
	return &ExpenseService{repo}
}
