package tomat

import (
	"context"
	"gopkg.in/guregu/null.v4"
	"time"
)

// Timer timer's model
type Timer struct {
	ID           string
	Duration     time.Duration
	StartAt      time.Time
	EndAt        time.Time
	StopAt       null.Time
	LastPauseAt  null.Time
	LastResumeAt null.Time
	FinishAt     null.Time
	Completed    bool
}

// TimerHandler handle timer actions
type TimerHandler interface {
	Start(ctx context.Context, startAt time.Time, dur time.Duration) (*Timer, error)
	Pause(ctx context.Context, timer *Timer, pauseAt time.Time) error
	Resume(ctx context.Context, timer *Timer, resumeAt time.Time) error
	Stop(ctx context.Context, timer *Timer, stopAt time.Time) error
	Finish(ctx context.Context, timer *Timer, finishAt time.Time) error
}
