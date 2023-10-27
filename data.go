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
}

type BudgetItem struct {
    ID string
    Name string
    Amount float64
    Remaining float64
}

func init() {
    income := []BudgetItem{
        {
            ID: "BI1",
            Name: "Paycheck 1",
            Amount: 10000,
            Remaining: 0,
        },
        {
            ID: "BI2",
            Name: "Paycheck 2",
            Amount: 10000,
            Remaining: 0,
        },
    }

    housing := []BudgetItem{
        {
            ID: "BI1",
            Name: "Mortgage",
            Amount: 3000,
            Remaining: 3000,
        },
    }

    categories := []Category{
        {
            ID: "C1",
            Name: "Income",
            Items: income,
        },
        {
            ID: "C2",
            Name: "Housing",
            Items: housing,
        },
    }

    data = Budget{
        ID: "B1",
        Name: "October",
        Categories: categories,
    }
}
