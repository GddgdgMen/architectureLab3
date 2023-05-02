package tests

import (
	"github.com/GddgdgMen/architectureLab3/painter"
	"github.com/GddgdgMen/architectureLab3/painter/lang"
	"image"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	type args struct {
		in io.Reader
	}
	tests := []struct {
		name    string
		command string
		want    painter.Operation
		wantErr bool
	}{
		{
			name:    "white",
			command: "white",
			want:    painter.OperationFunc(painter.WhiteFill),
			wantErr: false,
		},
		{
			name:    "green",
			command: "green",
			want:    painter.OperationFunc(painter.GreenFill),
			wantErr: false,
		},
		{
			name:    "bgrect",
			command: "bgrect 0 0 100 100",
			want:    &painter.BgRectangle{image.Rect(0, 0, 100, 100)},
			wantErr: false,
		},
		{
			name:    "figure",
			command: "figure 200 200",
			want:    &painter.Figure{X: 200, Y: 200},
			wantErr: false,
		},
		{
			name:    "move",
			command: "move 100 100",
			want:    &painter.Move{X: 100, Y: 100},
			wantErr: false,
		},
		{
			name:    "reset",
			command: "reset",
			want:    painter.OperationFunc(painter.ResetScreen),
			wantErr: false,
		},
		{
			name:    "update",
			command: "update",
			want:    painter.UpdateOp,
			wantErr: false,
		},
		{
			name:    "Wrong command",
			command: "non-existent command",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &lang.Parser{}
			got, err := p.Parse(strings.NewReader(tt.command))
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
