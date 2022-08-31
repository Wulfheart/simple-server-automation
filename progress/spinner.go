package progress

import (
	"bytes"
	"sync"
	"time"

	"github.com/gosuri/uiprogress/util/strutil"
)

// I have no idea why this has to be set multiple time
var charset = []string{"⣾", "⣾", "⣽", "⣽", "⣻", "⣻", "⢿", "⢿", "⡿", "⡿", "⣟", "⣟", "⣯", "⣯", "⣷", "⣷"}

// Spinner represents a progress bar
type Spinner struct {

	// TimeStated is time progress began
	TimeStarted time.Time

	// timeElased is the time elapsed for the progress
	timeElapsed time.Duration

	mtx          *sync.RWMutex
	currentIndex int
}

// DecoratorFunc is a function that can be prepended and appended to the progress bar
type DecoratorFunc func(b *Spinner) string

// NewBar returns a new progress bar
func NewBar(total int) *Spinner {
	return &Spinner{

		mtx: &sync.RWMutex{},
	}
}

// Bytes returns the byte presentation of the progress bar
func (b *Spinner) Bytes() []byte {

	// add fill and empty bits
	var buf bytes.Buffer
	buf.Write([]byte(b.GetCurrentCursor()))
	return buf.Bytes()
}

// String returns the string representation of the bar
func (b *Spinner) String() string {
	// b.IncrementCurrentIndex()
	return string(b.Bytes())
}

// TimeElapsed returns the time elapsed
func (b *Spinner) TimeElapsed() time.Duration {
	b.mtx.RLock()
	defer b.mtx.RUnlock()
	return b.timeElapsed
}

// TimeElapsedString returns the formatted string represenation of the time elapsed
func (b *Spinner) TimeElapsedString() string {
	return strutil.PrettyTime(b.TimeElapsed())
}

func (b *Spinner) GetCurrentCursor() string {
	result := charset[b.currentIndex]
	return result
}

func (b *Spinner) IncrementCurrentIndex() {
	if b.currentIndex < len(charset)-1 {
		b.currentIndex++
	} else {
		b.currentIndex = 0
	}
}

func (s *Spinner) Done() {

}
