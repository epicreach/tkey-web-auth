# tkey-web-authenticator

A simple command-line tool to interface with a Tillitis TKey for secure authentication.

## How to run

Extract the `tar.gz` or `.zip` file and navigate to the folder.
Run the binary file using

```bash
./tkeyauth
```

or

```cmd
.\tkeyauth.exe
```

## Build Instructions

You can also manually build the binary file if you wish.

### Requirements

To build the project, you'll need:

- [Go 1.23+](https://golang.org/dl/)
- `make`
- [Gpg4win](https://gpg4win.org/download.html)

### Build Using `make`

Clone the repo and run the following commands in your terminal:

```bash
git clone https://github.com/epicreach/tkey-web-authenticator.git
cd tkey-web-authenticator
make
```

You should now have a folder called `dist`. In this folder you will find a binary file called `tkeyauth`
To run the binary file use:

```bash
cd dist
./tkeyauth
```

#### To remove the binary

```bash
make clean
```

### Build Without `make`

In the root folder:

```cmd
go build -o dist\tkeyauth .\cmd\main.go
cd dist
.\tkeyauth.exe
```
