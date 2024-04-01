package main

const (
	walkAcceleration = 0.5
	walkSpeed        = 24
)

type WalkState struct {
	player *Player
	input  *Joypad
}

func NewWalkState(player *Player, input *Joypad) *WalkState {
	return &WalkState{
		player: player,
		input:  input,
	}
}

func (ws *WalkState) Update() {
	ws.updateTargetVelocity()
	ws.updateVelocity()
	ws.applyVelocity()
	ws.checkBoundaries()
	ws.handleJump()
	ws.handlePivot()
}

func (ws *WalkState) updateTargetVelocity() {
	switch {
	case ws.input.HoldDown[Left]:
		ws.player.TargetVelocityX = -walkSpeed
	case ws.input.HoldDown[Right]:
		ws.player.TargetVelocityX = walkSpeed
	default:
		ws.player.TargetVelocityX = 0
	}
}

func (ws *WalkState) updateVelocity() {
	if ws.input.HoldDown[B] {
		ws.player.setState(ws.player.running)
		ws.player.MotionState = Walk
		return
	}

	switch {
	case ws.player.VelocityX < ws.player.TargetVelocityX:
		ws.player.VelocityX += walkAcceleration
	case ws.player.VelocityX > ws.player.TargetVelocityX:
		ws.player.VelocityX -= walkAcceleration
	default:
		if ws.player.VelocityX == 0 {
			ws.player.setState(ws.player.idle)
			ws.player.MotionState = Idle
		}
	}
}

func (ws *WalkState) applyVelocity() {
	ws.player.PositionX += ws.player.VelocityX
	ws.player.Sprite.X = int(SubpixelsToPx(ws.player.PositionX))
}

func (ws *WalkState) checkBoundaries() {
	switch {
	case ws.player.Sprite.X > RightBound:
		ws.player.Sprite.X = RightBound
		ws.player.PositionX = PxToSubpixels(px(RightBound))
		ws.player.setState(ws.player.idle)
		ws.player.MotionState = Idle
		ws.player.VelocityX = 0
	case ws.player.Sprite.X < LeftBound:
		ws.player.Sprite.X = LeftBound
		ws.player.PositionX = PxToSubpixels(px(LeftBound))
		ws.player.setState(ws.player.idle)
		ws.player.MotionState = Idle
		ws.player.VelocityX = 0
	}
}

func (ws *WalkState) handleJump() {
	if ws.input.JustPressed[A] {
		ws.player.setState(ws.player.jumping)
		ws.player.MotionState = Airborne
	}
}

func (ws *WalkState) handlePivot() {
	if (ws.input.HoldDown[Left] || ws.input.HoldDown[Right] || ws.input.JustPressed[Left] || ws.input.JustPressed[Right]) &&
		((ws.player.TargetVelocityX > 0 && ws.player.VelocityX < 0) || (ws.player.TargetVelocityX < 0 && ws.player.VelocityX > 0)) {
		ws.player.MotionState = Pivot
		ws.player.setState(ws.player.pivoting)
	}
}
