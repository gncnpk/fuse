package util

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"log/slog"
	"net/url"
	"os"
	"strings"
)

var exePath = "mgsvtpp.exe"
var backupPath = "mgsvtpp.exe.vendor"
var exeUrlOffset = int64(35746160)
var oldURL = "https://mgstpp-game.konamionline.com/tppstm/gate"

func SplitByteString(input []byte, width int) [][]byte {
	var result [][]byte

	for i := 0; i < len(input); i += width {
		end := i + width
		if end > len(input) {
			end = len(input)
		}

		chunk := input[i:end]

		result = append(result, chunk)
	}

	return result
}

func Md5sum(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hash := md5.New()
	if _, err = io.Copy(hash, file); err != nil {
		return nil, err
	}

	hashInBytes := hash.Sum(nil)
	return hashInBytes, nil
}

func Backup() error {
	slog.Info("creating exe backup")

	_, err := os.Stat(backupPath)
	if err == nil {
		slog.Info("exe backup already exists")
		return nil
	}

	_, err = os.Stat(exePath)
	if err != nil {
		return fmt.Errorf("stat: %w", err)
	}

	sum, err := Md5sum(exePath)
	if err != nil {
		return fmt.Errorf("md5 checksum: %w", err)
	}

	// 7cc5f282b068f741adda2bb1076fb721
	// version 1.0.15.3
	target := []byte{0x7c, 0xc5, 0xf2, 0x82, 0xb0, 0x68, 0xf7, 0x41, 0xad, 0xda, 0x2b, 0xb1, 0x07, 0x6f, 0xb7, 0x21}
	if bytes.Compare(sum, target) != 0 {
		return fmt.Errorf("exe checksum mismatch, got %s, want %s", fmt.Sprintf("%x", sum), fmt.Sprintf("%x", target))
	}

	source, err := os.Open(exePath)
	if err != nil {
		return fmt.Errorf("source: %w", err)
	}
	defer source.Close()

	backupFile, err := os.Create(backupPath)
	if err != nil {
		return fmt.Errorf("backup file: %w", err)
	}
	defer backupFile.Close()

	if _, err = backupFile.ReadFrom(source); err != nil {
		return fmt.Errorf("copy data: %w", err)
	}

	return nil
}

func ValidateURL(clientURL string) (string, error) {
	if !strings.HasPrefix(clientURL, "http") {
		return "", fmt.Errorf("missing url protocol, expected http://%s or https://%s, got ", clientURL, clientURL)
	}

	clientURL = strings.TrimSpace(clientURL)
	clientURL = strings.TrimSuffix(clientURL, "/")

	if !strings.HasSuffix(clientURL, "/tppstm/gate") {
		clientURL = clientURL + "/tppstm/gate"
	}

	u, err := url.Parse(clientURL)
	if err != nil {
		return "", fmt.Errorf("invalid url: %w", err)
	}
	clientURL = u.String()

	return clientURL, nil
}

func Patch(gateURL string) error {
	var err error

	if gateURL, err = ValidateURL(gateURL); err != nil {
		return err
	}

	file, err := os.OpenFile(exePath, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("open: %w", err)
	}
	defer file.Close()

	_, err = file.Seek(exeUrlOffset, io.SeekStart)
	if err != nil {
		return fmt.Errorf("seek: %w", err)
	}

	padLen := len(oldURL) - len(gateURL)
	pad := make([]byte, padLen)
	_, err = file.Write([]byte(gateURL))
	if err != nil {
		return fmt.Errorf("write url patch: %w", err)
	}

	_, err = file.Write(pad)
	if err != nil {
		return fmt.Errorf("write url patch padding: %w", err)
	}

	return nil
}
