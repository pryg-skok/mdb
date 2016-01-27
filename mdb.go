package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type state struct {
	storage map[string]string
	count   map[string]int
	prev    *state
}

type mDB struct {
	states []*state
	top    *state
}

var Transaction bool

func NewState(prev *state) *state {
	return &state{
		make(map[string]string),
		make(map[string]int),
		prev,
	}
}

func NewDB() *mDB {
	states := make([]*state, 0)
	top := NewState(nil)
	states = append(states, top)
	return &mDB{states, top}
}

func (db mDB) Get(key string) string {
	if val, ok := db.top.Get(key); ok {
		return val
	}
	return "NULL"
}

func (db *mDB) Set(key string, val string) {
	db.top.Set(key, val)
}

func (db *mDB) Unset(key string) {
	db.top.Unset(key)
}

func (db mDB) NumEqualTo(val string) int {
	return db.top.NumEqualTo(val)
}

func (bl state) Get(key string) (string, bool) {
	if value, ok := bl.storage[key]; ok {
		return value, ok
	} else if bl.prev != nil {
		return bl.prev.Get(key)
	}
	return "", false
}

func (bl *state) Set(key string, val string) {
	bl.Unset(key)
	bl.storage[key] = val
	bl.count[val] += 1
}

func (bl *state) Unset(key string) {
	if curval, ok := bl.Get(key); ok {
		bl.count[curval] = bl.NumEqualTo(curval) - 1
	}
	if Transaction {
		bl.storage[key] = "NULL"
	} else {
		delete(bl.storage, key)
	}
}

func (bl state) NumEqualTo(val string) int {
	if cnt, ok := bl.count[val]; ok {
		return cnt
	} else if bl.prev != nil {
		return bl.prev.NumEqualTo(val)
	}
	return 0
}

func (db *mDB) BeginTransaction() {
	Transaction = true
	top := NewState(db.top)
	db.top = top
	db.states = append(db.states, top)
}

func (db *mDB) Rollback() string {
	if !Transaction {
		return "NO TRANSACTION"
	}
	state := db.states[len(db.states)-2]
	db.states = db.states[:len(db.states)-1]
	db.top = state
	return ""
}

func (db *mDB) Commit() string {
	if !Transaction {
		return "NO TRANSACTION"
	}
	newtop := NewState(nil)
	tounset := make([]string, 4)
	for i := len(db.states) - 1; i >= 0; i-- {
		bl := db.states[i]
		for key, val := range bl.storage {
			k, ok := newtop.Get(key)
			if !ok {
				newtop.Set(key, val)
			} else if val == "NULL" {
				tounset = append(tounset, k)
			} else { // key already in new map so continue
				continue
			}
		}
	}
	Transaction = false
	// cleanup Unset keys
	for _, key := range tounset {
		newtop.Unset(key)
	}
	return ""
}

func Router(DB *mDB, cmdArgs string) string {
	argList := strings.Fields(cmdArgs)
	switch argList[0] {
	case "END":
		break
	case "SET":
		DB.Set(argList[1], argList[2])
	case "GET":
		return DB.Get(argList[1])
	case "UNSET":
		DB.Unset(argList[1])
	case "NUMEQUALTO":
		return strconv.Itoa(DB.NumEqualTo(argList[1]))
	case "BEGIN":
		DB.BeginTransaction()
	case "ROLLBACK":
		return DB.Rollback()
	case "COMMIT":
		return DB.Commit()
	}
	return ""
}

func main() {
	DB := NewDB()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Router(DB, scanner.Text())
	}
}
