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
}

func init() {
    income := []BudgetItem{
        {
            ID: "1",
            Name: "Paycheck 1",
            Amount: 10000,
        },
        {
            ID: "2",
            Name: "Paycheck 2",
            Amount: 10000,
        },
    }

    housing := []BudgetItem{
        {
            ID: "1",
            Name: "Mortgage",
            Amount: 3000,
        },
    }

    categories := []Category{
        {
            ID: "1",
            Name: "Income",
            Items: income,
        },
        {
            ID: "2",
            Name: "Housing",
            Items: housing,
        },
    }

    data = Budget{
        ID: "1",
        Name: "October",
        Categories: categories,
    }
}
