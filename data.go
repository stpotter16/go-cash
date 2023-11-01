package main

type Budget struct {
	ID         string
	Name       string
	Categories []Category
}

type Category struct {
	ID       string
	Name     string
	Items    []BudgetItem
	BudgetID string
}

type BudgetItem struct {
	ID         string
	Name       string
	Amount     float64
	Remaining  float64
	CategoryID string
}

func dbLoadBudget(budgetID string) Budget {
	budgetRow := db.QueryRow(`select Budget.name from Budget where Budget.budgetId = ?`, budgetID)

	var budget Budget
	budget.ID = budgetID
	err := budgetRow.Scan(&budget.Name)
	if err != nil {
		// TODO - something better here?
		return Budget{}
	}

	categoryRows, err := db.Query(`select Category.categoryId, Category.name from Category where Category.budgetId = ?`, budgetID)
	if err != nil {
		// TODO - something better here?
		return Budget{}
	}
	for categoryRows.Next() {
		var category Category
		category.BudgetID = budgetID
		err = categoryRows.Scan(&category.ID, &category.Name)
		if err != nil {
			// TODO - something better here?
			return Budget{}
		}

		budgetItemRows, err := db.Query(`select BudgetItem.itemID, BudgetItem.name, BudgetItem.amount, BudgetItem.remaining from BudgetItem where BudgetItem.categoryId = ?`, category.ID)
		if err != nil {
			// TODO - something better here?
			return Budget{}
		}

		for budgetItemRows.Next() {
			var budgetItem BudgetItem
			budgetItem.CategoryID = category.ID
			err = budgetItemRows.Scan(&budgetItem.ID, &budgetItem.Name, &budgetItem.Amount, &budgetItem.Remaining)
			if err != nil {
				// TODO - something better here?
				return Budget{}
			}
			category.Items = append(category.Items, budgetItem)
		}
		budget.Categories = append(budget.Categories, category)
	}

	return budget
}
