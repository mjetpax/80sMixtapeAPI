package models

import (
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

// Cassette struct for storing songs in cassette format.
type Cassette struct {
	Type  string       `json:"tape_type"`
	SideA CassetteSide `json:"side_a"`
	SideB CassetteSide `json:"side_b"`
}

// CassetteSide is one side of a cassette.
type CassetteSide struct {
	Songs         []Song `json:"songs"`
	DurationLabel string `json:"play_time"`
	Duration      int    `json:"-"`
}

// AddSong adds a song to a cassette side.
func (cs *CassetteSide) AddSong(song Song) {
	cs.Songs = append(cs.Songs, song)
	cs.Duration += song.Duration
	cs.SetDurationLabel()
}

// SetDurationLabel convert duration to duration label string, ex. 125 == "2:05"
func (cs *CassetteSide) SetDurationLabel() {
	min := cs.Duration / 60
	minutes := "0"
	if min >= 1 {
		minutes = strconv.Itoa(min)
	}

	seconds := strconv.Itoa(cs.Duration % 60)
	if len(seconds) < 2 {
		seconds = "0" + seconds
	}

	cs.DurationLabel = minutes + ":" + seconds
}

// FetchCassette query database random songs and create cassette.
// Args cassetteTypes C120, C90, C60
func FetchCassette(db *sqlx.DB, cassetteType string) *Cassette {

	sideLength, typeLabel := GetCassetteMeta(cassetteType)

	// Average song length in the 80's was about 4 minutes.
	// Add some extra as margin, need enough songs to fill both sides of the tape. ((sideLen*2)/4)+10
	songCount := sideLength/2 + 10

	songs := []Song{}

	query := `SELECT id, title, artist, year,
	video, duration_label, duration
	FROM songs
	ORDER BY random()
	Limit $1`

	db.Select(&songs, query, songCount)

	// create cassette
	cassette := Cassette{Type: typeLabel}

	for i := 0; i < len(songs); i++ {

		// Fill Side A, duration is in seconds
		if cassette.SideA.Duration+songs[i].Duration < sideLength*60 {
			cassette.SideA.AddSong(songs[i])
			continue
		}

		// Fill Side B
		if cassette.SideB.Duration+songs[i].Duration < sideLength*60 {
			cassette.SideB.AddSong(songs[i])
			continue
		}

		break
	}

	return &cassette
}

// GetCassetteMeta takes cassette type C120, C90, C60.
// Returns tape single side duration in minutes (int) and type label (string).
func GetCassetteMeta(cassetteType string) (int, string) {
	cassetteType = strings.ToUpper(cassetteType)
	cassetteType = strings.Replace(cassetteType, "/", "", 1)

	switch cassetteType {
	case "C120":
		return 60, cassetteType

	case "C90":
		return 45, cassetteType

	default:
		// default is 60 minute tape
		return 30, "C60"
	}
}
