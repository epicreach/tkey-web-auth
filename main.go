package main

import (
	"crypto"
	"crypto/ed25519"
	"crypto/sha512"
	"fmt"
	"log"
	"os"

	"github.com/tillitis/tkeyclient"
	"golang.org/x/crypto/ssh"
)

const progname = "tkey-device-signer"

var le = log.New(os.Stderr, "", 0)

func main() {
	devPath := "/dev/ttyACM0"

	// tk := tkeyclient.New()
	// err := tk.Connect(devPath)
	// if err != nil {
	// 	panic("Failed to connect to TKEY")
	// }
	// fmt.Printf("Connected to TKEY\n")

	// nameVer, err := tk.GetNameVersion()
	// if err != nil {
	// 	panic("Failed to get Name and Version")
	// }
	// fmt.Printf("Firmware name: %s\n", nameVer.Name0)
	// fmt.Printf("Name1: %s\n", nameVer.Name1)
	// fmt.Printf("Version: %d\n", nameVer.Version)

	exit := func(code int) {
		os.Exit(code)
	}

	signer := NewSigner(devPath, tkeyclient.SerialSpeed, false, "", "", exit)
	if !signer.connect() {
		le.Printf("Connect failed")
		return
	}

	defer signer.disconnect()

	fmt.Printf("Connected to TKEY\n")

	pub, err := signer.tkSigner.GetPubkey()

	if err != nil {
		fmt.Println("Error getting Public Key")
		return
	}

	sshPub, _ := ssh.NewPublicKey(ed25519.PublicKey(pub))

	fmt.Printf("Public key is: \n%s\n", ssh.MarshalAuthorizedKey(sshPub))

	testHash := hash("test")

	sig, err := signer.Sign(nil, testHash[:], crypto.Hash(0))
	if err != nil {
		le.Printf("Sign failed: %s\n", err)
		return
	}
	fmt.Printf("Signature: %x\n", string(sig))

	if ed25519.Verify(pub, testHash[:], sig) {
		fmt.Println("Signature is valid!")
		return
	}
	fmt.Println("Signature is wrong :(")
}

func hash(strToHash string) [64]byte {
	return sha512.Sum512([]byte(strToHash))
}
