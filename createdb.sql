PRAGMA foreign_keys = ON;

create table Budget (
    budgetID string primary key,
    name string
);

create table Category (
    categoryID string primary key,
    budgetID string,
    name string,

    foreign key(budgetID) references Budget(budgetID)
);

create table BudgetItem (
    itemID string primary key,
    categoryID string,
    name string,
    amount float,
    remaining float,

    foreign key(categoryID) references Category(categoryID)
);

