package core

import (
	"github.com/gorustyt/go-assimp/common"
	"github.com/gorustyt/go-assimp/common/pb_msg"
	"math"
)

type AiCamera struct {
	/** The name of the camera.
	 *
	 *  There must be a node in the scenegraph with the same name.
	 *  This node specifies the position of the camera in the scene
	 *  hierarchy and can be animated.
	 */
	Name string

	/** Position of the camera relative to the coordinate space
	 *  defined by the corresponding node.
	 *
	 *  The default value is 0|0|0.
	 */
	Position *common.AiVector3D

	/** 'Up' - vector of the camera coordinate system relative to
	 *  the coordinate space defined by the corresponding node.
	 *
	 *  The 'right' vector of the camera coordinate system is
	 *  the cross product of  the up and lookAt vectors.
	 *  The default value is 0|1|0. The vector
	 *  may be normalized, but it needn't.
	 */
	Up *common.AiVector3D

	/** 'LookAt' - vector of the camera coordinate system relative to
	 *  the coordinate space defined by the corresponding node.
	 *
	 *  This is the viewing direction of the user.
	 *  The default value is 0|0|1. The vector
	 *  may be normalized, but it needn't.
	 */
	LookAt *common.AiVector3D

	/** Horizontal field of view angle, in radians.
	 *
	 *  The field of view angle is the angle between the center
	 *  line of the screen and the left or right border.
	 *  The default value is 1/4PI.
	 */
	HorizontalFOV float32

	/** Distance of the near clipping plane from the camera.
	 *
	 * The value may not be 0.f (for arithmetic reasons to prevent
	 * a division through zero). The default value is 0.1f.
	 */
	ClipPlaneNear float32

	/** Distance of the far clipping plane from the camera.
	 *
	 * The far clipping plane must, of course, be further away than the
	 * near clipping plane. The default value is 1000.f. The ratio
	 * between the near and the far plane should not be too
	 * large (between 1000-10000 should be ok) to avoid floating-point
	 * inaccuracies which could lead to z-fighting.
	 */
	ClipPlaneFar float32

	/** Screen aspect ratio.
	 *
	 * This is the ration between the width and the height of the
	 * screen. Typical values are 4/3, 1/2 or 1/1. This value is
	 * 0 if the aspect ratio is not defined in the source file.
	 * 0 is also the default value.
	 */
	Aspect float32

	/** Half horizontal orthographic width, in scene units.
	 *
	 *  The orthographic width specifies the half width of the
	 *  orthographic view box. If non-zero the camera is
	 *  orthographic and the mAspect should define to the
	 *  ratio between the orthographic width and height
	 *  and mHorizontalFOV should be set to 0.
	 *  The default value is 0 (not orthographic).
	 */
	OrthographicWidth float32
}

func (ai *AiCamera) FromPbMsg(p *pb_msg.AiCamera) *AiCamera {
	if p == nil {
		return nil
	}
	ai.Name = p.Name
	ai.Position = (&common.AiVector3D{}).FromPbMsg(p.Position)
	ai.Up = (&common.AiVector3D{}).FromPbMsg(p.Up)
	ai.LookAt = (&common.AiVector3D{}).FromPbMsg(p.LookAt)
	ai.HorizontalFOV = p.HorizontalFOV
	ai.ClipPlaneNear = p.ClipPlaneNear
	ai.ClipPlaneFar = p.ClipPlaneFar
	ai.Aspect = p.Aspect
	ai.OrthographicWidth = p.OrthographicWidth
	return ai
}

func (ai *AiCamera) Clone() *AiCamera {
	if ai == nil {
		return nil
	}
	r := &AiCamera{}
	r.Name = ai.Name
	r.Position = ai.Position.Clone()
	r.Up = ai.Up.Clone()
	r.LookAt = ai.LookAt.Clone()
	r.HorizontalFOV = ai.HorizontalFOV
	r.ClipPlaneNear = ai.ClipPlaneNear
	r.ClipPlaneFar = ai.ClipPlaneFar
	r.Aspect = ai.Aspect
	r.OrthographicWidth = ai.OrthographicWidth
	return r
}

func (ai *AiCamera) ToPbMsg() *pb_msg.AiCamera {
	if ai == nil {
		return nil
	}
	r := &pb_msg.AiCamera{}
	r.Name = ai.Name
	r.Position = ai.Position.ToPbMsg()
	r.Up = ai.Up.ToPbMsg()
	r.LookAt = ai.LookAt.ToPbMsg()
	r.HorizontalFOV = ai.HorizontalFOV
	r.ClipPlaneNear = ai.ClipPlaneNear
	r.ClipPlaneFar = ai.ClipPlaneFar
	r.Aspect = ai.Aspect
	r.OrthographicWidth = ai.OrthographicWidth
	return r
}

func NewAiCamera() *AiCamera {
	return &AiCamera{
		Up:                common.NewAiVector3D3(0., 1., 0.),
		LookAt:            common.NewAiVector3D3(0., 0., 1.),
		HorizontalFOV:     0.25 * math.Pi,
		ClipPlaneNear:     1.0,
		ClipPlaneFar:      1000.,
		Aspect:            0.,
		OrthographicWidth: 0.,
	}
}
