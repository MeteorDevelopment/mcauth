package codes

import (
	"errors"
	"go.minekube.com/gate/pkg/util/uuid"
	"math/rand"
	"sync"
	"time"
)

type codeData struct {
	time time.Time
	id   uuid.UUID
}

var codes []codeData
var mutex sync.Mutex

var ErrInvalidCode = errors.New("invalid code")

func init() {
	codes = make([]codeData, 1000000)
}

func New(id uuid.UUID) int {
	mutex.Lock()
	defer mutex.Unlock()

	start := rand.Intn(len(codes))
	i := start

	for {
		data := codes[i]

		if data.time.Add(time.Minute * 5).Before(time.Now()) {
			codes[i] = codeData{}
			data = codeData{}
		}

		if data.time.IsZero() {
			codes[i] = codeData{
				time: time.Now(),
				id:   id,
			}

			return i
		}

		i++
		if i >= len(codes) {
			i = 0
		}

		if i == start {
			break
		}
	}

	return -1
}

func Retrieve(code int) (uuid.UUID, error) {
	mutex.Lock()
	defer mutex.Unlock()

	data := codes[code]

	if data.time.Add(time.Minute * 5).Before(time.Now()) {
		return uuid.UUID{}, ErrInvalidCode
	}

	codes[code] = codeData{}

	return data.id, nil
}
