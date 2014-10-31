// Package uid generates an URL safe string.
//
//  id, err := Uid(10)
//  if err != nil {
//    panic(err)
//  }
//  fmt.Println(id)
//  // 9BZ1sApAX4
package uid

import (
	"crypto/rand"
	"strings"
)

const UIDCHARS = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

// Uid generates an URL safe string of given length.
func Gen(length int) (string, error) {
	b := make([]byte, length)
	res := make([]string, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	for i, val := range b {
		res[i] = string(UIDCHARS[int(val)%len(UIDCHARS)])
	}
	return strings.Join(res, ""), nil
}