# Pwned Passwords Lookup

A simple command line tool to look up if a password is on the [haveibeenpwned.com](https://haveibeenpwned.com/) password list. haveibeenpwned.com has hundreds of millions of real world passwords previously exposed in data breaches.

This using the API from `https://api.pwnedpasswords.com/`. Using the range endpoint to look up parts of a hash of the password, so the password's clear text is never sent to the server. The web version also uses the same API on the [https://haveibeenpwned.com/Passwords](https://haveibeenpwned.com/Passwords) page. Beside just using the API, this tool will ensure not to echo out the password in clear text as you are typing.

This is a port of a Nim version that I made sometime ago, [pwned_passwords_lookup](https://github.com/amscotti/pwned_passwords_lookup)

![PwnedPasswordsLookup Screenshot](https://github.com/amscotti/PwnedPasswordsLookup/blob/main/PwnedPasswordsLookup_screenshot.png?raw=true)

## To Build
* Run `go build`

## To Run
```bash
$ ./PwnedPasswordsLookup
Password to lookup:
Your password was found in the haveibeenpwned.com database.
```