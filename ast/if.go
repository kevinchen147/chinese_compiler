package ast

import (
	"io"
)

type If struct {
	nodeBase      //行列信息
	Test     Node //条件
	Body     Node //执行块
	Alt      Node //后续如else等
}

func NewIf(pos PositionHolder, test Node, body Node, alt Node) *If {
	n := &If{Test: test, Body: body, Alt: alt}
	n.init(pos)
	return n
}

func (n *If) dump(w io.Writer, nest int) {
	header(n, w, nest, true)
	tag("条件", w, nest+1)
	dumpNode(n.Test, w, nest+2)
	tag("执行块", w, nest+1)
	dumpNode(n.Body, w, nest+2)
	tag("后续", w, nest+1)
	dumpNode(n.Alt, w, nest+2)
}
