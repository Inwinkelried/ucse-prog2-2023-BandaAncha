package model

type Parada struct {
	Ciudad       string  `bson:"ciudad"`
	KmRecorridos float64 `bson:"kmRecorridos"`
}
