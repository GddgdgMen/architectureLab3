package tests

import (
	"github.com/GddgdgMen/architectureLab3/painter"
	"golang.org/x/exp/shiny/screen"
	"image"
	"image/color"
	"image/draw"
	"testing"
	"time"
)

func TestLoop_Post(t *testing.T) {
	var (
		l  painter.Loop
		mr mockReceiver
	)

	l = painter.Loop{
		Receiver: &mr,
	}

	l.Start(mockScreen{})

	l.Post(painter.OperationFunc(painter.GreenFill))
	l.Post(painter.OperationFunc(painter.WhiteFill))
	l.Post(painter.UpdateOp)

	if mr.lastTexture != nil {
		t.Fatal("Receiver got the texture too early")
	}

	time.Sleep(1 * time.Second)

	lt, ok := mr.lastTexture.(*mockTexture)

	if !ok {
		t.Fatal("Receiver still has not texture")
	}

	if lt.FillCounter != 2 {
		t.Error("Unexpected number of Fill calls:", lt.FillCounter)
	}
}

type mockReceiver struct {
	lastTexture screen.Texture
}

func (m *mockReceiver) Update(t screen.Texture) {
	m.lastTexture = t
}

type mockScreen struct {
}

func (m mockScreen) NewBuffer(size image.Point) (screen.Buffer, error) {
	panic("implement me")
}

func (m mockScreen) NewTexture(size image.Point) (screen.Texture, error) {
	return &mockTexture{size: size}, nil
}

func (m mockScreen) NewWindow(opts *screen.NewWindowOptions) (screen.Window, error) {
	panic("implement me")
}

type mockTexture struct {
	size        image.Point
	FillCounter int
}

func (m *mockTexture) Release() {}

func (m *mockTexture) Size() image.Point {
	return m.size
}

func (m *mockTexture) Bounds() image.Rectangle {
	return image.Rectangle{Max: m.size}
}

func (m *mockTexture) Upload(dp image.Point, src screen.Buffer, sr image.Rectangle) {
	panic("implement me")
}

func (m *mockTexture) Fill(dr image.Rectangle, src color.Color, op draw.Op) {
	m.FillCounter++
}
