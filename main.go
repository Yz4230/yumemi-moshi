package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type playLog struct {
	Ts       time.Time
	PlayerID string
	Score    int
}

func validateHeader(header string) bool {
	headerElements := strings.Split(header, ",")
	if len(headerElements) != 3 {
		return false
	}
	if headerElements[0] != "create_timestamp" ||
		headerElements[1] != "player_id" ||
		headerElements[2] != "score" {
		return false
	}
	return true
}

func validatePlayerID(playerID string) bool {
	playerIDRegExp := regexp.MustCompile("^[a-zA-Z0-9]+$")
	return playerIDRegExp.MatchString(playerID)
}

func parseRow(row string) (*playLog, error) {
	rowElements := strings.Split(row, ",")
	if len(rowElements) != 3 {
		return nil, errors.New("invalid number of row elements")
	}
	time, err := time.Parse("2006/01/02 15:04", rowElements[0])
	if err != nil {
		return nil, err
	}
	playerID := rowElements[1]
	if !validatePlayerID(playerID) {
		return nil, errors.New("invalid player id: " + rowElements[1])
	}
	score, err := strconv.Atoi(rowElements[2])
	if err != nil {
		return nil, err
	}
	if score < 1 {
		return nil, errors.New("invalid score: " + rowElements[2])
	}
	return &playLog{
		Ts:       time,
		PlayerID: playerID,
		Score:    score,
	}, nil
}

func parseCSV(reader io.Reader) ([]*playLog, error) {
	var logs []*playLog
	bufReader := bufio.NewReader(reader)
	header, _, err := bufReader.ReadLine()
	if err != nil {
		return nil, err
	} else if !validateHeader(string(header)) {
		return nil, fmt.Errorf("invalid header: %s", header)
	}
	lineCount := 2 // header has already been read
	for {
		line, _, err := bufReader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		log, err := parseRow(string(line))
		if err != nil {
			return nil, fmt.Errorf("invalid row at line %d: %s", lineCount, err.Error())
		}
		logs = append(logs, log)
		lineCount++
	}
	return logs, nil
}

func groupLogsByPlayerID(logs []*playLog) map[string][]*playLog {
	result := make(map[string][]*playLog)
	for _, log := range logs {
		if logs, ok := result[log.PlayerID]; ok {
			result[log.PlayerID] = append(logs, log)
		} else {
			result[log.PlayerID] = []*playLog{log}
		}
	}
	return result
}

func sumScore(logs []*playLog) int {
	scoreSum := 0
	for _, log := range logs {
		scoreSum += log.Score
	}
	return scoreSum
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: <executable> [filename]")
		return
	}
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()
	logs, err := parseCSV(f)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	playerIDLogs := groupLogsByPlayerID(logs)

	type scoreMean struct {
		PlayerID string
		Mean     int
	}
	scoreMeans := make([]*scoreMean, 0)
	for playerID, logs := range playerIDLogs {
		scoreSum := sumScore(logs)
		mean := float64(scoreSum) / float64(len(logs))
		scoreMeans = append(
			scoreMeans,
			&scoreMean{
				PlayerID: playerID,
				Mean:     int(math.Round(mean)),
			},
		)
	}

	sort.Slice(scoreMeans, func(i, j int) bool {
		return scoreMeans[i].Mean > scoreMeans[j].Mean
	})

	rankCount := 1
	for i := 0; i < len(scoreMeans); i++ {
		fmt.Printf("%d,%s,%d\n", rankCount, scoreMeans[i].PlayerID, scoreMeans[i].Mean)
		if i+1 < len(scoreMeans) && scoreMeans[i].Mean != scoreMeans[i+1].Mean {
			rankCount = i + 2
		}
	}
}
