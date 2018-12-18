package helpers

import (
	"strconv"
	"strings"
)

type LoginRequest struct {
	UserName          string
	PassMD5           string

	OsuBuild          string
	UTCOffset         int
	DisplayLocation   bool
	SecurityHash      string
	BlockNonFriendsDM bool
}


func ParseLoginRequest(body []byte) *LoginRequest {
	data := string(body)
	retVal := LoginRequest{}

	lines := strings.Split(data, "\n")
	if len(lines) < 3 {
		return nil
	}

	retVal.UserName = lines[0]
	retVal.PassMD5  = lines[1]

	clientData := lines[2]
	lines = strings.Split(clientData, "|")
	if len(lines) < 5 {
		return nil
	}
	var err error = nil

	retVal.OsuBuild = lines[0]
	retVal.UTCOffset, err = strconv.Atoi(lines[1])
	if err != nil {
		return nil
	}

	retVal.DisplayLocation, err = strconv.ParseBool(lines[2])
	if err != nil {
		return nil
	}

	retVal.SecurityHash = lines[3]
	retVal.BlockNonFriendsDM, err = strconv.ParseBool(lines[4])
	if err != nil {
		return nil
	}

	return &retVal
}
