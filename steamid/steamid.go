package steamid

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io"
	"strconv"
)

const InvalidSteamID = 76561197960265729

func ValidateString(steamID string) (uint64, error) {
	sid, err := strconv.Atoi(steamID)
	if err != nil {
		return 0, fmt.Errorf("invalid steamID: %s", steamID)
	}

	if err = Validate(uint64(sid)); err != nil {
		return 0, err
	}

	return uint64(sid), nil
}

func Validate(steamID uint64) error {
	if steamID < 0x110000100000000 || steamID > 0x01100001FFFFFFFF {
		return fmt.Errorf("steamID %d is out of valid range", steamID)
	}

	return nil
}

func FromTicket(ticket string) (uint64, error) {
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(ticket)))
	_, err := base64.StdEncoding.Decode(dst, []byte(ticket))
	if err != nil {
		return 0, fmt.Errorf("cannot decode steam ticket: %w", err)
	}

	var steamID uint64
	r := bytes.NewReader(dst)
	if _, err = r.Seek(12, io.SeekStart); err != nil {
		return 0, fmt.Errorf("cannot seek steam ticket: %w", err)
	}

	err = binary.Read(r, binary.LittleEndian, &steamID)
	if err != nil {
		return 0, fmt.Errorf("cannot read steamID from ticket: %w", err)
	}

	return steamID, nil
}
