package main

func main() {
	dbMain := DatabaseStore{}
	fileMain := FileStore{}
	RecordVictory(&dbMain, "John", 100)
	RecordVictory(&fileMain, "John", 100)
}
