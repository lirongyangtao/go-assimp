package post_processing

import (
	"context"
	"github.com/gorustyt/go-assimp/core"
	"github.com/gorustyt/go-assimp/driver/base/iassimp"
)

var (
	_ iassimp.PostProcessing = (*SplitByBoneCountProcess)(nil)
)

type SplitByBoneCountProcess struct {
}

func (s SplitByBoneCountProcess) IsActive(pFlags int) bool {
	return (iassimp.AiPostProcessSteps(pFlags) & iassimp.AiProcess_SplitByBoneCount) != 0
}

func (s SplitByBoneCountProcess) Execute(pScene *core.AiScene) {
	//TODO implement me
	panic("implement me")
}

func (s SplitByBoneCountProcess) SetupProperties(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}
