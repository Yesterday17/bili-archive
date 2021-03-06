package utils

import (
	"time"
)

// Status:
// -1     Error
// 0      Not Started
// 1      Running
// 2      Finished
type Progress struct {
	Title    string       `json:"title"`
	Status   int          `json:"status"`
	Message  string       `json:"message"`
	Progress ProgressData `json:"data"`
	callback func(pg *Progress)
}

type ProgressData struct {
	Size     int64     `json:"size"`
	Progress int64     `json:"progress"`
	Increase int64     `json:"increase"`
	Time     time.Time `json:"time"`
}

func NewProgress(title string, size int64, callback func(pg *Progress)) *Progress {
	return &Progress{
		Title:  title,
		Status: 0,
		Progress: ProgressData{
			Size:     size,
			Progress: 0,
			Time:     time.Now(),
			Increase: 0,
		},
		callback: callback,
	}
}

func (pg *Progress) Add(len int64) {
	pg.Progress.Progress += len
}

func (pg *Progress) Write(p []byte) (n int, err error) {
	n = len(p)
	pg.Progress.Progress += int64(n)
	pg.Progress.Increase = int64(n)
	pg.Progress.Time = time.Now()

	if pg.Status != 2 && pg.callback != nil {
		pg.callback(pg)
	}
	return
}

func (pg *Progress) Start() {
	pg.Status = 1
	if pg.callback != nil {
		pg.callback(pg)
	}
}

func (pg *Progress) Finish() {
	pg.Status = 2
	if pg.callback != nil {
		pg.callback(pg)
	}
}

func (pg *Progress) Error(err error) {
	pg.Status = -1
	pg.Message = err.Error()
	if pg.callback != nil {
		pg.callback(pg)
	}
}
