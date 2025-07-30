package service

import (
	"lifedash/model"
	"lifedash/repo"
)

type ExpenseService struct {
	repo *repo.ExpenseRepo
}

func NewExpenseService(repo *repo.ExpenseRepo) *ExpenseService {
	return &ExpenseService{repo}
}

func (es *ExpenseService) GetAggregatedExpensesForUser(userId int) ([]model.AggregatedExpense, error) {
	expenseDistribution, err := es.repo.GetExpenseDistribution(userId)
	if err != nil {
		return nil, err
	}
	return expenseDistribution, nil
}
