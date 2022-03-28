package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/nicolerobin/tiny_sqlite/constant"
	"github.com/nicolerobin/tiny_sqlite/entity"
)

func promt() {
	fmt.Print("db > ")
}

func metaCommand(input string) error {
	if input[:5] == ".exit" {
		fmt.Println("Bye!")
		os.Exit(0)
	} else {
		return errors.New("unknown meta command")
	}

	return nil
}

func prepareStatement(input string) (entity.Statement, error) {
	stmt := entity.Statement{}
	if len(input) < 6 {
		return stmt, errors.New("unknown statement")
	}

	var err error
	switch input[:6] {
	case "insert":
		stmt.Type = constant.STATEMENT_INSERT
		fmt.Sscanf(input, "insert %d %s %s\n", &stmt.Row.Id, &stmt.Row.Name, &stmt.Row.Email)
	case "select":
		stmt.Type = constant.STATEMENT_SELECT
	default:
		err = errors.New("unknown statement")
	}
	return stmt, err
}

func execStatement(stmt *entity.Statement) error {
	switch stmt.Type {
	case constant.STATEMENT_INSERT:
		fmt.Println("this is a insert statement")
	case constant.STATEMENT_SELECT:
		fmt.Println("this is a select statement")
	}

	return nil
}

func main() {
	for {
		promt()
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("reader.ReadString() failed, err:%s\n", err)
			continue
		}
		input = input[:len(input)-1]
		if len(input) == 0 {
			continue
		}

		if input[:1] == "." {
			err := metaCommand(input)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}

		stmt, err := prepareStatement(input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("stmt:%+v\n", stmt)
	}
}
