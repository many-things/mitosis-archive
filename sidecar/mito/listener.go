package mito

import (
	"github.com/many-things/mitosis/sidecar/utils"
)

type EventListenMgr interface {
	AddJob(job Job)
	Run() error
}

func (m eventMgr) AddJob(job Job) {
	m.jobs = append(m.jobs, job)
}

func (m eventMgr) Run() error {
	utils.ForEach(m.jobs, func(j Job) {
		m.errGroup.Go(func() error { return j(m.eventCtx) })
	})

	if err := m.errGroup.Wait(); err != nil {
		return err
	}

	return nil
}
