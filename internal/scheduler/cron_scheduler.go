package scheduler

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

type Scheduler interface {
	ScheduleTask(cronExpression string, task func()) (taskID int, err error)
	StartScheduler()
	StopScheduler()
	RemoveTask(taskID int) error
}

type cronScheduler struct {
	cron       *cron.Cron
	tasks      map[int]cron.EntryID
	nextTaskID int
}

func NewCronScheduler() Scheduler {
	return &cronScheduler{
		cron:  cron.New(),
		tasks: make(map[int]cron.EntryID),
	}
}

func (cs *cronScheduler) ScheduleTask(cronExpression string, task func()) (taskID int, err error) {
	id, err := cs.cron.AddFunc(cronExpression, task)
	if err != nil {
		return 0, err
	}

	cs.tasks[cs.nextTaskID] = id
	cs.nextTaskID++

	return cs.nextTaskID - 1, nil
}

func (cs *cronScheduler) StartScheduler() {
	cs.cron.Start()
}

func (cs *cronScheduler) StopScheduler() {
	cs.cron.Stop()
}

func (cs *cronScheduler) RemoveTask(taskID int) error {
	if entryID, exists := cs.tasks[taskID]; exists {
		cs.cron.Remove(entryID)
		delete(cs.tasks, taskID)
		return nil
	}

	return fmt.Errorf("task with ID %d not found", taskID)
}
