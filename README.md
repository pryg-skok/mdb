# Simple in-memory database

The main idea behind the implementation of such key-value database with transactions is to
represent it as slice of small databases (states), which store only diff from previous one.
This reveals a simple way of implementing transactions, because we can just drop the "top"
state from slice and set the whole database to previous condition - perform rollbacks.

Because we store only diffs, each time when command (such GET or NUMEQUALTO) is requested,
we should calculate the global database's state. The main idea here is to store a pointer of
previous state within new state. Using this pointer we can recursively go through states and
thereby find requested item.

# System requirements

Nothing special: just Golang (say v. >= 1.0)
