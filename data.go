package main

var data []Category

type Category struct {
    ID string
    Name string
}

func init() {
    data = []Category{
        {
            ID: "1",
            Name: "Income",
        },
        {
            ID: "2",
            Name: "Housing",
        },
    }
}
