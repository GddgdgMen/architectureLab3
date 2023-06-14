package lang

import (
	"image"
	"reflect"
	"strings"
	"testing"

	"github.com/GddgdgMen/architectureLab3/painter"
)

func TestParser_Parse(t *testing.T) {
	tests := []struct {
		name    string
		command string
		want    []painter.Operation
		wantErr bool
	}{
		{
			name:    "white",
			command: "white",
			want:    []painter.Operation{painter.OperationFunc(painter.WhiteFill)},
			wantErr: false,
		},
		{
			name:    "green",
			command: "green",
			want:    []painter.Operation{painter.OperationFunc(painter.GreenFill)},
			wantErr: false,
		},
		{
			name:    "bgrect",
			command: "bgrect 0 0 100 100",
			want: []painter.Operation{painter.OperationFunc(painter.ResetScreen),
				&painter.BgRectangle{Rect: image.Rect(0, 0, 100, 100)}},
			wantErr: false,
		},
		{
			name:    "figure",
			command: "figure 200 200",
			want: []painter.Operation{painter.OperationFunc(painter.ResetScreen),
				&painter.Figure{X: 200, Y: 200}},
			wantErr: false,
		},
		{
			name:    "move",
			command: "move 100 100",
			want: []painter.Operation{painter.OperationFunc(painter.ResetScreen),
				&painter.Move{X: 100, Y: 100}},
			wantErr: false,
		},
		{
			name:    "reset",
			command: "reset",
			want:    []painter.Operation{painter.OperationFunc(painter.ResetScreen)},
			wantErr: false,
		},
		{
			name:    "update",
			command: "update",
			want: []painter.Operation{painter.OperationFunc(painter.ResetScreen),
				painter.UpdateOp},
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
			p := &Parser{}
			got, err := p.Parse(strings.NewReader(tt.command))
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(got) != len(tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
				return
			}

			for i := 0; i < len(got); i++ {
				if reflect.TypeOf(got[i]) != reflect.TypeOf(tt.want[i]) {
					t.Errorf("Parse() got = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}
