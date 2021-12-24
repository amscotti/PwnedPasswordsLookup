package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"syscall"

	"golang.org/x/term"
)

type PasswordHash struct {
	prefix string
	suffix string
}

const APIEndpoint = "https://api.pwnedpasswords.com/range/"

func (h *PasswordHash) lookupPassword() (bool, error) {
	rsp, err := http.Get(APIEndpoint + h.prefix)
	if err != nil {
		return false, err
	}
	defer rsp.Body.Close()

	data, err := io.ReadAll(rsp.Body)
	if err != nil {
		return false, err
	}

	hashes := strings.Split(string(data), "\n")
	for _, hash := range hashes {
		suffix := strings.Split(hash, ":")[0]
		if suffix == strings.ToUpper(h.suffix) {
			return true, nil
		}
	}

	return false, nil
}

func getPassword() (PasswordHash, error) {
	fmt.Print("Password to lookup: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return PasswordHash{}, err
	}
	fmt.Println("")

	hasher := sha1.New()
	hasher.Write(bytePassword)
	sha := hex.EncodeToString(hasher.Sum(nil))

	return PasswordHash{prefix: sha[0:5], suffix: sha[5:]}, nil
}

func main() {
	hash, err := getPassword()
	if err != nil {
		log.Fatal(err)
	}

	found, err := hash.lookupPassword()
	if err != nil {
		log.Fatal(err)
	}

	if found {
		fmt.Println("Your password was found in the haveibeenpwned.com database.")
	} else {
		fmt.Println("Looks like your password wasn't found in the haveibeenpwned.com database!")
	}
}
