package logic

import (
    "fmt"
    // "os"
    // "github.com/joho/godotenv"
    // _ "github.com/go-sql-driver/mysql"
)

func CalcDiff(list1, list2 []string) []string {
    s := make(map[string]struct{}, len(list1))
    for _, data := range list2 {
        s[data] = struct{}{}
    }
    r := make([]string, 0, len(list2))
    for _, data := range list1 {
        if _, ok := s[data]; ok {
            continue
        }
        r = append(r, data)
    }
	fmt.Println(r)
    return r
}