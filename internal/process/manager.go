package process

import (
	"fmt"

	"github.com/devbookhq/devbookd/internal/smap"
	"go.uber.org/zap"
)

type Manager struct {
	procs *smap.Map[ID, Process]
}

func NewManager() *Manager {
	return &Manager{
		procs: smap.New[ID, Process](),
	}
}

func (m *Manager) Remove(id ID) {
	proc, ok := m.procs.Get(id)

	if !ok {
		return
	}

	proc.Kill()
	m.procs.Remove(id)
}

func (m *Manager) Get(id ID) (*Process, bool) {
	return m.procs.Get(id)
}

func (m *Manager) Add(id ID, cmd string, envVars *map[string]string, rootdir string, logger *zap.SugaredLogger) (*Process, error) {
	proc, err := New(id, cmd, envVars, rootdir, logger)
	if err != nil {
		return nil, fmt.Errorf("error starting new process with id '%s': %+v", id, err)
	}

	m.procs.Insert(proc.ID, proc)
	return proc, nil
}
