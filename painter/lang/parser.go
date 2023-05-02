package lang

import (
	"bufio"
	"fmt"
	"image"
	"io"
	"strconv"
	"strings"

	"github.com/GddgdgMen/architectureLab3/painter"
)

// Parser уміє прочитати дані з вхідного io.Reader та повернути список операцій представлені вхідним скриптом.
type Parser struct {
	lastBgColor painter.Operation
	lastBgRect  *painter.BgRectangle
	figures     []*painter.Figure
	moveOps     []painter.Operation
	updateOp    painter.Operation
}

func (p *Parser) initialize() {
	if p.lastBgColor == nil {
		p.lastBgColor = painter.OperationFunc(painter.ResetScreen)
	}
	if p.updateOp != nil {
		p.updateOp = nil
	}
}

// Parse reads and parses input from the provided io.Reader and returns the corresponding list of painter.Operation.
func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	p.initialize()
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() { // loop through the input stream using the scanner
		commandLine := scanner.Text()

		err := p.parse(commandLine) // parse the command line into an operation
		if err != nil {
			return nil, err
		}
	}
	return p.finalResult(), nil
}

func (p *Parser) finalResult() []painter.Operation {
	var res []painter.Operation
	if p.lastBgColor != nil {
		res = append(res, p.lastBgColor)
	}
	if p.lastBgRect != nil {
		res = append(res, p.lastBgRect)
	}
	if len(p.moveOps) != 0 {
		res = append(res, p.moveOps...)
	}
	p.moveOps = nil
	if len(p.figures) != 0 {
		println(len(p.figures))
		for _, figure := range p.figures {
			res = append(res, figure)
		}
	}
	if p.updateOp != nil {
		res = append(res, p.updateOp)
	}
	return res
}

func (p *Parser) resetState() {
	p.lastBgColor = nil
	p.lastBgRect = nil
	p.figures = nil
	p.moveOps = nil
	p.updateOp = nil
}

func (p *Parser) parse(cmdl string) error {
	args := strings.Split(cmdl, " ")
	instruction := args[0]

	var iArgs []int
	for _, arg := range args[1:] {
		fArg, err := strconv.ParseFloat(arg, 64)
		if err == nil {
			iArgs = append(iArgs, int(fArg*800))
		}
	}

	switch instruction {
	case "white":
		p.lastBgColor = painter.OperationFunc(painter.WhiteFill)
	case "green":
		p.lastBgColor = painter.OperationFunc(painter.GreenFill)
	case "bgrect":
		err := validate(args, 4)
		if err != nil {
			return err
		}
		p.lastBgRect = &painter.BgRectangle{Rect: image.Rect(iArgs[0], iArgs[1], iArgs[2], iArgs[3])}
	case "figure":
		err := validate(args, 2)
		if err != nil {
			return err
		}
		figure := painter.Figure{X: iArgs[0], Y: iArgs[1]}
		p.figures = append(p.figures, &figure)
	case "move":
		err := validate(args, 2)
		if err != nil {
			return err
		}
		moveOp := painter.Move{X: iArgs[0], Y: iArgs[1], Figures: p.figures}
		p.moveOps = append(p.moveOps, &moveOp)
	case "reset":
		p.resetState()
		p.lastBgColor = painter.OperationFunc(painter.ResetScreen)
	case "update":
		p.updateOp = painter.UpdateOp
	default:
		return fmt.Errorf("could not parse command %v", cmdl)
	}
	return nil
}

func validate(args []string, expected int) error {
	if len(args) != expected+1 {
		return fmt.Errorf("wrong number of arguments for '%s', expected %v", args[0], expected)
	}
	var command = args[0]
	for _, arg := range args[1:] {
		_, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			return fmt.Errorf("invalid argument for '%s': '%s' is not a number", command, arg)
		}
	}
	return nil
}
