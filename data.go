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
}

func init() {
    categories := []Category{
        {
            ID: "1",
            Name: "Income",
        },
        {
            ID: "2",
            Name: "Housing",
        },
    }
    data = Budget{
        ID: "1",
        Name: "October",
        Categories: categories,
    }
}
