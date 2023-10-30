package main

var data Budget

type Budget struct {
    ID string
    Name string
    Categories []Category
}

type Category struct {
    ID string
    Name string
    Items []BudgetItem
    BudgetID string
}

type BudgetItem struct {
    ID string
    Name string
    Amount float64
    Remaining float64
    BudgetID string
}

func init() {
    income := []BudgetItem{
        {
            ID: "BI1",
            Name: "Paycheck 1",
            Amount: 10000,
            Remaining: 0,
            BudgetID: "B1",
        },
        {
            ID: "BI2",
            Name: "Paycheck 2",
            Amount: 10000,
            Remaining: 0,
            BudgetID: "B1",
        },
    }

    housing := []BudgetItem{
        {
            ID: "BI1",
            Name: "Mortgage",
            Amount: 3000,
            Remaining: 3000,
            BudgetID: "B1",
        },
    }

    categories := []Category{
        {
            ID: "C1",
            Name: "Income",
            Items: income,
            BudgetID: "B1",
        },
        {
            ID: "C2",
            Name: "Housing",
            Items: housing,
            BudgetID: "B1",
        },
    }

    data = Budget{
        ID: "B1",
        Name: "October",
        Categories: categories,
    }
}
