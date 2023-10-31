insert into Budget(budgetID, name) values("B1", "October");

insert into Category(categoryID, budgetID, name) values ("C1", "B1", "Income");
insert into Category(categoryID, budgetID, name) values ("C2", "B1", "Housing");

insert into BudgetItem(itemID, categoryID, name, amount, remaining) values ("BI1", "C1", "Paycheck 1", 10000, 0);
insert into BudgetItem(itemID, categoryID, name, amount, remaining) values ("BI2", "C1", "Paycheck 2", 10000, 0);
insert into BudgetItem(itemID, categoryID, name, amount, remaining) values ("BI3", "C2", "Mortgage", 3000, 3000);
