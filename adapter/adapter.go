package adapter

import (
	flats2 "flats-parser/repositories"
)

func AdsToFlats[T any](flatType int8, ads []T, getID func(T) int) []flats2.Flat {
	var flats []flats2.Flat

	for _, ad := range ads {
		flat := flats2.Flat{
			RemoteID: getID(ad),
			Type:     flatType,
		}
		flats = append(flats, flat)
	}

	return flats
}
