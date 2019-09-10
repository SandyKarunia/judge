package sandbox

import (
	"errors"
	"os"
)

var MoeErrorIsNotReady = errors.New("moe sandbox is not ready yet, please call Initialize() before use")
var MoeErrorIsGone = errors.New("moe sandbox is gone and should not be used anymore")

type moeSandbox struct {
	// if true, means the sandbox has been destroyed and should not be used anymore
	isDestroyed bool

	// if true, means the sandbox has been initialized and ready to use
	isInitialized bool

	// boxID is the box id of this sandbox
	boxID int
}

func (m *moeSandbox) Initialize() error {
	// it has been initialized before, skip this step
	if m.isInitialized {
		return nil
	}
	panic("implement me")
}

func (m *moeSandbox) CopyFile(source *os.File) error {
	if err := m.beforeUseCheck(); err != nil {
		return err
	}
	panic("implement me")
}

func (m *moeSandbox) GetFile(fileName string) (*os.File, error) {
	if err := m.beforeUseCheck(); err != nil {
		return nil, err
	}
	panic("implement me")
}

func (m *moeSandbox) Run(commands ...string) error {
	if err := m.beforeUseCheck(); err != nil {
		return err
	}
	panic("implement me")
}

func (m *moeSandbox) Destroy() {
	// it has been destroyed before, skip this step
	if m.isDestroyed {
		return
	}
	panic("implement me")
}

func (m *moeSandbox) beforeUseCheck() error {
	// sandbox has been destroyed
	if m.isDestroyed {
		return MoeErrorIsGone
	}
	// sandbox has not been initialized
	if !m.isInitialized {
		return MoeErrorIsNotReady
	}
	return nil
}