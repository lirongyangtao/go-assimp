package post_processing

import (
	"context"
	"github.com/gorustyt/go-assimp/core"
	"github.com/gorustyt/go-assimp/driver/base/iassimp"
)

var (
	_ iassimp.PostProcessing = (*SortByPTypeProcess)(nil)
)

type SortByPTypeProcess struct {
}

func (s SortByPTypeProcess) IsActive(pFlags int) bool {
	return (iassimp.AiPostProcessSteps(pFlags) & iassimp.AiProcess_SortByPType) != 0
}

func (s SortByPTypeProcess) Execute(pScene *core.AiScene) {
	//TODO implement me
	panic("implement me")
}

func (s SortByPTypeProcess) SetupProperties(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}
