package repo

import (
	"database/sql"
	"fmt"
	"lifedash/model"
)

type ExpenseRepo struct {
	db *sql.DB
}

func NewExpenseRepo(db *sql.DB) *ExpenseRepo {
	return &ExpenseRepo{db}
}

func (er *ExpenseRepo) GetExpenseDistribution(userId int) ([]model.Expense, error) {
	query := "SELECT date, category, amount FROM expense WHERE user_id = ?"
	expenses := []model.Expense{}
	results, err := er.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	for results.Next() {
		e := model.Expense{}
		err = results.Scan(&e.Date, &e.Category, &e.Amount)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, e)
	}
	for _, e := range expenses {
		fmt.Println(e)
	}
	return expenses, nil
}
