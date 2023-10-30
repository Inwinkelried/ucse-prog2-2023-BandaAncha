package utils

import "time"

type FiltroEnvio struct {
	PatenteCamion string
	Estado        string
	UltimaParada  string
	FechaMenor    time.Time
	FechaMayor    time.Time
}
