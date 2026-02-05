package main

import "fmt"

type ScoreSaver interface {
	SaveScore(username string, score int) error
}

type DatabaseStore struct{}

type FileStore struct{}

func (d *DatabaseStore) SaveScore(username string, score int) error {
	mainLine := fmt.Sprintf("Connecting to SQL... Saving %d for %s.", score, username)
	fmt.Println(mainLine)
	return nil
}

func (f *FileStore) SaveScore(username string, score int) error {
	mainLine := fmt.Sprintf("Writing to highscores.txt: %d - %s.", score, username)
	fmt.Println(mainLine)
	return nil
}

func RecordVictory(s ScoreSaver, name string, points int) {
	if points > 100 {
		points += 10
	}
	(s).SaveScore(name, points)
}
