# Geth Breaker
A simple to use tool to break into your Ethereum JSON keystore wallet with a password list.


## Installation
Download the executable built for your operating system in the [GethBreaker Releases](https://github.com/hunterlong/gethbreaker/releases/latest) section.
Once you have the file downloaded, make it executable.

```
chmod +x gethbreaker-osx-x64
mv gethbreaker-osx-x64 /usr/local/bin/gethbreaker
gethbreaker
```

## Usage
```
gethbreaker wallet.json passwords.txt

...
Trying: 50RT8deU
Trying: d1TzXD5g
Trying: 6ODGo3il
Trying: 2cb2Co6C
Trying: 7TxGdUk4
Trying: password123

Found password!!!
File:     wallet.json
Address:  0x75DC4a5084260183c79223F63e1F18E811059bee
Password: password123
```