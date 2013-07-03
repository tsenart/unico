package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Visit struct {
	UserId    int
	Timestamp uint64
}

func ParseVisit(b []byte) (Visit, error) {
	var visit Visit
	var err error

	tokens := strings.Split(string(b), " ")
	if len(tokens) != 9 {
		return visit, fmt.Errorf("Couldn't parse log visit %s", b)
	}
	urlParts := strings.Split(tokens[5], "/")
	if len(urlParts) == 0 {
		return visit, fmt.Errorf("Couldn't parse user id. %s", tokens[5])
	}
	if _, err = fmt.Sscanf(urlParts[1], "%d", &(visit.UserId)); err != nil {
		return visit, fmt.Errorf("Couldn't parse user id. %s", tokens[5])
	}
	visit.Timestamp, err = strconv.ParseUint(tokens[3][1:len(tokens[3])-1], 10, 64)
	if err != nil {
		return visit, err
	}

	return visit, nil
}
