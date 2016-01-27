package main

import (
	"testing"
)

type Case []string

var SimpleCases = []Case{
	Case{"SET ex 10", ""},
	Case{"GET ex", "10"},
	Case{"UNSET ex", ""},
	Case{"GET ex", "NULL"},
}

var TransactionCases_1 = []Case{
	Case{"BEGIN", ""},
	Case{"SET a 10", ""},
	Case{"GET a", "10"},
	Case{"BEGIN", ""},
	Case{"SET a 20", ""},
	Case{"GET a", "20"},
	Case{"ROLLBACK", ""},
	Case{"GET a", "10"},
	Case{"ROLLBACK", ""},
	Case{"GET a", "NULL"},
}

var TransactionCases_2 = []Case{
	Case{"BEGIN", ""},
	Case{"SET a 30", ""},
	Case{"BEGIN", ""},
	Case{"SET a 40", ""},
	Case{"COMMIT", ""},
	Case{"GET a", "40"},
	Case{"ROLLBACK", "NO TRANSACTION"},
}

var TransactionCases_3 = []Case{
	Case{"SET a 50", ""},
	Case{"BEGIN", ""},
	Case{"GET a", "50"},
	Case{"SET a 60", ""},
	Case{"BEGIN", ""},
	Case{"UNSET a", ""},
	Case{"GET a", "NULL"},
	Case{"ROLLBACK", ""},
	Case{"GET a", "60"},
	Case{"COMMIT", ""},
	Case{"GET a", "60"},
}

var TransactionCases_4 = []Case{
	Case{"SET a 10", ""},
	Case{"BEGIN", ""},
	Case{"NUMEQUALTO 10", "1"},
	Case{"BEGIN", ""},
	Case{"UNSET a", ""},
	Case{"NUMEQUALTO 10", "0"},
	Case{"ROLLBACK", ""},
	Case{"NUMEQUALTO 10", "1"},
	Case{"COMMIT", ""},
}

func TestSimpleCommands(t *testing.T) {
	DB := NewDB()
	for _, cmd := range SimpleCases {
		in, out := cmd[0], cmd[1]
		result := Router(DB, in)
		if result != out {
			t.Fail()
		}
	}
}

func TestTransaction_1(t *testing.T) {
	DB := NewDB()
	for _, cmd := range TransactionCases_1 {
		in, out := cmd[0], cmd[1]
		result := Router(DB, in)
		if result != out {
			t.Fail()
		}
	}
}

func TestTransaction_2(t *testing.T) {
	DB := NewDB()
	for _, cmd := range TransactionCases_2 {
		in, out := cmd[0], cmd[1]
		result := Router(DB, in)
		if result != out {
			t.Fail()
		}
	}
}

func TestTransaction_3(t *testing.T) {
	DB := NewDB()
	for _, cmd := range TransactionCases_3 {
		in, out := cmd[0], cmd[1]
		result := Router(DB, in)
		if result != out {
			t.Fail()
		}
	}
}

func TestTransaction_4(t *testing.T) {
	DB := NewDB()
	for _, cmd := range TransactionCases_4 {
		in, out := cmd[0], cmd[1]
		result := Router(DB, in)
		if result != out {
			t.Fail()
		}
	}
}
