package wg

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

const (
	PearTemplate = "\n[Peer]\nPublicKey = %sAllowedIPs = 10.0.0.%d/32\n"
)

func AppendPear(pubKey string, id int, pathToConf string) error {
	f, err := os.OpenFile(pathToConf, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.WriteString(fmt.Sprintf(PearTemplate, pubKey, id)); err != nil {
		return err
	}

	return nil
}

func RemovePear(pubKey string, id int, pathToConf string, pathToImages string) error {
	if err := os.Remove(fmt.Sprintf(filepath.Join(pathToImages, "%d_wg.png"), id)); err != nil {
		return err
	}

	input, err := os.ReadFile(pathToConf)
	if err != nil {
		return err
	}

	output := bytes.Replace(input, []byte(fmt.Sprintf(PearTemplate, pubKey, id)), []byte(""), -1)

	err = os.WriteFile(pathToConf, []byte(output), 0777)
	if err != nil {
		return err
	}

	return nil
}
