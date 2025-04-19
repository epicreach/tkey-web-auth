# tkey-web-authenticator

A simple command-line tool to interface with a Tillitis TKey for secure authentication.

---

## Getting Started

### Option 1: Use a Pre-Built Binary

1. Visit the [Releases Page](https://github.com/epicreach/tkey-web-authenticator/releases).
2. Download the appropriate `tar.gz` or `.zip` file for your operating system.
3. Extract the downloaded file and navigate to the extracted folder.
4. Run the binary:

   **Linux/macOS:**

   ```bash
   ./tkeyauth
   ```

   **Windows:**

   ```cmd
   .\tkeyauth.exe
   ```

---

### Option 2: Build the Binary Yourself

If you prefer to build the binary from source, follow the instructions below.

## Build Instructions

### Prerequisites

Ensure you have the following installed:

- [Go 1.23+](https://golang.org/dl/)
- [Gpg4win](https://gpg4win.org/download.html) (required for Windows users)
- `make` (optional)

---

### Build Using `make`

1. Clone the repository:

   ```bash
   git clone https://github.com/epicreach/tkey-web-authenticator.git
   cd tkey-web-authenticator
   ```

2. Build the project:

   ```bash
   make
   ```

3. The binary will be created in the root directory:

   ```bash
   ./tkeyauth
   ```

4. To clean up build artifacts:
   ```bash
   make clean
   ```

---

### Build Without `make`

1. Clone the repository:

   ```bash
   git clone https://github.com/epicreach/tkey-web-authenticator.git
   cd tkey-web-authenticator
   ```

2. Build the binary manually:

   ```bash
   go build -o tkeyauth ./cmd/main.go
   ```

3. Run the binary:

   **Linux/macOS:**

   ```bash
   ./tkeyauth
   ```

   **Windows:**

   ```cmd
   .\tkeyauth.exe
   ```

---

## License

This project is licensed under the BSD 2-Clause License. See the [LICENSE](LICENSE) file for details.

---
