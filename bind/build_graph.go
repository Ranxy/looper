package bind

import (
	"bytes"
	"container/list"
	"fmt"
	"strings"

	"github.com/Ranxy/looper/symbol"
	"github.com/Ranxy/looper/syntax"
)

type ControlFlowGraph struct {
	Start    *BasicBlock
	End      *BasicBlock
	Blocks   *DataMaps[BasicBlock]
	Branches *DataMaps[BasicBlockBranch]
}

func CreateControlFlow(body *BoundBlockStatements) *ControlFlowGraph {
	blocks := BuildBasicBlock(body)
	dms := DataMapsFromSlice[BasicBlock](blocks)
	graphBuild := NewGraphBuild()
	return graphBuild.Build(dms)
}

func ALlPathReturn(body *BoundBlockStatements) bool {
	graph := CreateControlFlow(body)
	for iter := graph.End.Incoming.Iter(); iter.HasNext(); {
		branch := iter.Next()
		if len(branch.From.Statements) == 0 {
			return false
		}
		lastStatemtn := branch.From.Statements[len(branch.From.Statements)-1]
		if lastStatemtn.Kind() != BoundNodeKindReturnStatement {
			return false
		}
	}

	return true
}

type BasicBlock struct {
	IsStart    bool
	IsEnd      bool
	Statements []Boundstatement
	Incoming   *DataMaps[BasicBlockBranch]
	Outgoing   *DataMaps[BasicBlockBranch]
}

func NewBasicBlock(start bool, end bool) *BasicBlock {
	return &BasicBlock{
		IsStart:    start,
		IsEnd:      end,
		Statements: []Boundstatement{},
		Incoming:   NewDataMaps[BasicBlockBranch](),
		Outgoing:   NewDataMaps[BasicBlockBranch](),
	}
}

func (b *BasicBlock) String() string {
	if b.IsStart {
		return "<start>"
	}
	if b.IsEnd {
		return "<end>"
	}
	buf := bytes.NewBuffer([]byte{})
	for _, statement := range b.Statements {
		WriteTo(buf, statement)
	}

	return buf.String()
}

type DataMaps[T any] struct {
	m map[*T]*list.Element
	*list.List
}

func NewDataMaps[T any]() *DataMaps[T] {
	return &DataMaps[T]{
		m:    map[*T]*list.Element{},
		List: list.New(),
	}
}

func DataMapsFromSlice[T any](slices []*T) *DataMaps[T] {
	dm := &DataMaps[T]{
		m:    map[*T]*list.Element{},
		List: list.New(),
	}

	for _, s := range slices {
		dm.Add(s)
	}
	return dm
}

func (b *DataMaps[T]) Add(block *T) {
	e := b.List.PushBack(block)
	b.m[block] = e
}

func (b *DataMaps[T]) Remove(block *T) {
	e := b.m[block]
	delete(b.m, block)
	b.List.Remove(e)
}

func (b *DataMaps[T]) Len() int {
	return len(b.m)
}

func (b *DataMaps[T]) First() *T {
	e := b.List.Front()
	if e == nil {
		return nil
	}
	value := e.Value.(*T)
	return value
}

func (b *DataMaps[T]) Next(block *T) *T {
	e := b.m[block]
	if e == nil {
		return nil
	}

	next := e.Next()
	if next == nil {
		return nil
	}

	return next.Value.(*T)
}

func (b *DataMaps[T]) Iter() *DataMapsIter[T] {

	return &DataMapsIter[T]{
		D: b,
		e: b.List.Front(),
	}
}

type DataMapsIter[T any] struct {
	D *DataMaps[T]
	e *list.Element
}

func (d *DataMapsIter[T]) Next() *T {
	if d.e == nil {
		return nil
	}
	res := d.e.Value.(*T)
	d.e = d.e.Next()
	return res
}
func (d *DataMapsIter[T]) HasNext() bool {
	return d.e != nil
}

type BasicBlockBranch struct {
	From      *BasicBlock
	To        *BasicBlock
	Condition BoundExpression
}

func (b *BasicBlockBranch) String() string {
	if b.Condition == nil {
		return ""
	} else {
		buf := bytes.NewBuffer([]byte{})
		WriteTo(buf, b.Condition)
		return buf.String()
	}
}

func BuildBasicBlock(block *BoundBlockStatements) []*BasicBlock {
	statements := []Boundstatement{}
	blocks := []*BasicBlock{}

	endBlock := func() {
		if len(statements) > 0 {
			block := NewBasicBlock(false, false)

			block.Statements = append(block.Statements, statements...)
			blocks = append(blocks, block)
			statements = statements[:0]
		}
	}

	startBlock := func() {
		endBlock()
	}

	for _, statement := range block.Statement {
		switch statement.Kind() {
		case BoundNodeKindLabelStatement:
			startBlock()
			statements = append(statements, statement)
		case BoundNodeKindGotoStatement, BoundNodeKindConditionalGotoStatement, BoundNodeKindReturnStatement:
			statements = append(statements, statement)
			startBlock()

		case BoundNodeKindVariableDeclaration, BoundNodeKindExpressionStatement:
			statements = append(statements, statement)
		default:
			panic("Unexceted statement " + statement.Kind().String())
		}
	}

	endBlock()
	return blocks
}

type GraphBuilder struct {
	blockFromStatement map[Boundstatement]*BasicBlock
	blockFromLabel     map[*BoundLabel]*BasicBlock
	branch             *DataMaps[BasicBlockBranch]
	start              *BasicBlock
	end                *BasicBlock
}

func NewGraphBuild() *GraphBuilder {
	return &GraphBuilder{
		blockFromStatement: map[Boundstatement]*BasicBlock{},
		blockFromLabel:     map[*BoundLabel]*BasicBlock{},
		branch:             NewDataMaps[BasicBlockBranch](),
		start:              NewBasicBlock(true, false),
		end:                NewBasicBlock(false, true),
	}
}

func (g *GraphBuilder) Connect(from, to *BasicBlock, condition BoundExpression) {
	if ble, ok := condition.(*BoundLiteralExpression); ok {
		value, ok := ble.Value.(bool)
		if ok && value {
			condition = nil
		} else {
			return
		}
	}

	branch := &BasicBlockBranch{
		From:      from,
		To:        to,
		Condition: condition,
	}
	from.Outgoing.Add(branch)
	to.Incoming.Add(branch)
	g.branch.Add(branch)
}

func (g *GraphBuilder) RemoveBlock(blocks *DataMaps[BasicBlock], block *BasicBlock) {

	for branch := range block.Incoming.m {
		branch.From.Outgoing.Remove(branch)
		g.branch.Remove(branch)
	}
	for branch := range block.Outgoing.m {
		branch.To.Incoming.Remove(branch)
		g.branch.Remove(branch)
	}

	blocks.Remove(block)
}

func (g *GraphBuilder) Negate(condition BoundExpression) BoundExpression {
	if literal, ok := condition.(*BoundLiteralExpression); ok {
		value := literal.Value.(bool)
		return NewBoundLiteralExpression(!value)
	}

	op := BindBoundUnaryOperator(syntax.SyntaxKindBangToken, symbol.TypeBool)

	return NewBoundUnaryExpression(op, condition)
}

func (g *GraphBuilder) Build(blocks *DataMaps[BasicBlock]) *ControlFlowGraph {
	if blocks.Len() == 0 {
		g.Connect(g.start, g.end, nil)
	} else {
		g.Connect(g.start, blocks.First(), nil)
	}

	iter := blocks.Iter()
	for {
		block := iter.Next()
		if block == nil {
			break
		}

		for _, statement := range block.Statements {
			g.blockFromStatement[statement] = block
			if labelStatement, ok := statement.(*LabelStatement); ok {
				g.blockFromLabel[labelStatement.Label] = block
			}
		}

	}
	iter = blocks.Iter()
	for iter.HasNext() {
		current := iter.Next()
		next := func() *BasicBlock {
			n := blocks.Next(current)
			if n == nil {
				return g.end
			} else {
				return n
			}
		}()

		for idx, statement := range current.Statements {
			isLastStatementInBlock := idx == len(current.Statements)-1
			switch statement.Kind() {
			case BoundNodeKindGotoStatement:
				gs := statement.(*GotoStatement)
				toBLock := g.blockFromLabel[gs.Label]
				g.Connect(current, toBLock, nil)
			case BoundNodeKindConditionalGotoStatement:
				cgs := statement.(*ConditionalGotoStatement)
				thenBlock := g.blockFromLabel[cgs.Label]
				elseBlock := next
				negatedCondition := g.Negate(cgs.Condition)
				var thenCondition BoundExpression
				var elseCondition BoundExpression
				if cgs.JumpIfFalse {
					thenCondition = negatedCondition
					elseCondition = cgs.Condition
				} else {
					thenCondition = cgs.Condition
					elseCondition = negatedCondition
				}

				g.Connect(current, thenBlock, thenCondition)
				g.Connect(current, elseBlock, elseCondition)
			case BoundNodeKindReturnStatement:
				g.Connect(current, g.end, nil)
			case BoundNodeKindVariableDeclaration:
				fallthrough
			case BoundNodeKindLabelStatement:
				fallthrough
			case BoundNodeKindExpressionStatement:
				if isLastStatementInBlock {
					g.Connect(current, next, nil)
				}

			default:
				panic(fmt.Sprintf("Unexcepted statement %s", statement.Kind()))
			}
		}

	}

SCAN_AGAIN:
	iter = blocks.Iter()
	for iter.HasNext() {
		first := iter.Next()
		if first.Incoming.Len() == 0 {
			g.RemoveBlock(blocks, first)
			goto SCAN_AGAIN
		}
	}

	blocks.PushFront(g.start)
	blocks.PushBack(g.end)

	return &ControlFlowGraph{
		Start:    g.start,
		End:      g.end,
		Blocks:   blocks,
		Branches: g.branch,
	}

}

func (g *ControlFlowGraph) WriteTo(w *strings.Builder) {
	quoto := func(text string) string {
		return "\"" + strings.ReplaceAll(text, "\"", "\\\"") + "\""
	}
	w.WriteString("digraph G {")
	blockIds := map[*BasicBlock]string{}

	idfunc := func() func() string {
		i := 0
		return func() string {
			id := fmt.Sprintf("N%d", i)
			i += 1
			return id
		}
	}()

	for iter := g.Blocks.Iter(); iter.HasNext(); {
		block := iter.Next()
		blockIds[block] = idfunc()
	}

	for iter := g.Blocks.Iter(); iter.HasNext(); {
		block := iter.Next()
		id := blockIds[block]
		label := quoto(strings.ReplaceAll(block.String(), "\n", "\\l"))
		w.WriteString(fmt.Sprintf("    %s [label = %s shape=box]", id, label))
		w.WriteByte('\n')
	}

	for iter := g.Branches.Iter(); iter.HasNext(); {
		branch := iter.Next()
		fromID := blockIds[branch.From]
		toID := blockIds[branch.To]
		label := quoto(branch.String())
		w.WriteString(fmt.Sprintf("    %s -> %s [label = %s]", fromID, toID, label))
		w.WriteByte('\n')
	}

	w.WriteString("}")

}
