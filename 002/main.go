package main

type Book struct {
	ID        int
	UserID    int
	Title     string
	AuthorIDs []int
}

type User struct {
	ID    int
	Name  int
	Email string
}

type Author struct {
	ID          int
	Name        string
	Description string
}

type Order struct {
}
