package output

type OutputConfig struct {
	// Addresses of stations, by callsign
	addresses map[int]string

	// Messages, by number
	messages map[int]string
}
