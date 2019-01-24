package engine

import (
	"github.com/pkg/errors"
	"gonum.org/v1/gonum/graph"
	"gorgonia.org/gorgonia/node"
	"gorgonia.org/gorgonia/ops"
)

// Code generated by genapi, which is a API generation tool for Gorgonia. DO NOT EDIT.

// Abs performs a pointwise abs.
func Abs(a *Node) (*Node, error) { return unaryOpNode(newElemUnaryOp(absOpType, a), a) }

// AbsOp ...
func NewAbsOperation() Operation {
	return func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		it := getOrderedChildren(g, n)
		if it.Len() != 1 {
			return nil, errors.New("Unexpected number of children")
		}
		children := make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		return newElemUnaryOp(absOpType, children[0]), nil
	}
}

// Sign performs a pointwise sign.
func Sign(a *Node) (*Node, error) { return unaryOpNode(newElemUnaryOp(signOpType, a), a) }

// SignOp ...
func NewSignOperation() Operation {
	return func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		it := getOrderedChildren(g, n)
		if it.Len() != 1 {
			return nil, errors.New("Unexpected number of children")
		}
		children := make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		return newElemUnaryOp(signOpType, children[0]), nil
	}
}

// Ceil performs a pointwise ceil.
func Ceil(a *Node) (*Node, error) { return unaryOpNode(newElemUnaryOp(ceilOpType, a), a) }

// CeilOp ...
func NewCeilOperation() Operation {
	return func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		it := getOrderedChildren(g, n)
		if it.Len() != 1 {
			return nil, errors.New("Unexpected number of children")
		}
		children := make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		return newElemUnaryOp(ceilOpType, children[0]), nil
	}
}

// Floor performs a pointwise floor.
func Floor(a *Node) (*Node, error) { return unaryOpNode(newElemUnaryOp(floorOpType, a), a) }

// FloorOp ...
func NewFloorOperation() Operation {
	return func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		it := getOrderedChildren(g, n)
		if it.Len() != 1 {
			return nil, errors.New("Unexpected number of children")
		}
		children := make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		return newElemUnaryOp(floorOpType, children[0]), nil
	}
}

// Sin performs a pointwise sin.
func Sin(a *Node) (*Node, error) { return unaryOpNode(newElemUnaryOp(sinOpType, a), a) }

// SinOp ...
func NewSinOperation() Operation {
	return func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		it := getOrderedChildren(g, n)
		if it.Len() != 1 {
			return nil, errors.New("Unexpected number of children")
		}
		children := make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		return newElemUnaryOp(sinOpType, children[0]), nil
	}
}

// Cos performs a pointwise cos.
func Cos(a *Node) (*Node, error) { return unaryOpNode(newElemUnaryOp(cosOpType, a), a) }

// CosOp ...
func NewCosOperation() Operation {
	return func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		it := getOrderedChildren(g, n)
		if it.Len() != 1 {
			return nil, errors.New("Unexpected number of children")
		}
		children := make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		return newElemUnaryOp(cosOpType, children[0]), nil
	}
}

// Exp performs a pointwise exp.
func Exp(a *Node) (*Node, error) { return unaryOpNode(newElemUnaryOp(expOpType, a), a) }

// ExpOp ...
func NewExpOperation() Operation {
	return func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		it := getOrderedChildren(g, n)
		if it.Len() != 1 {
			return nil, errors.New("Unexpected number of children")
		}
		children := make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		return newElemUnaryOp(expOpType, children[0]), nil
	}
}

// Log performs a pointwise log.
func Log(a *Node) (*Node, error) { return unaryOpNode(newElemUnaryOp(lnOpType, a), a) }

// LogOp ...
func NewLogOperation() Operation {
	return func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		it := getOrderedChildren(g, n)
		if it.Len() != 1 {
			return nil, errors.New("Unexpected number of children")
		}
		children := make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		return newElemUnaryOp(lnOpType, children[0]), nil
	}
}

// Log2 performs a pointwise log2.
func Log2(a *Node) (*Node, error) { return unaryOpNode(newElemUnaryOp(log2OpType, a), a) }

// Log2Op ...
func NewLog2Operation() Operation {
	return func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		it := getOrderedChildren(g, n)
		if it.Len() != 1 {
			return nil, errors.New("Unexpected number of children")
		}
		children := make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		return newElemUnaryOp(log2OpType, children[0]), nil
	}
}

// Neg performs a pointwise neg.
func Neg(a *Node) (*Node, error) { return unaryOpNode(newElemUnaryOp(negOpType, a), a) }

// NegOp ...
func NewNegOperation() Operation {
	return func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		it := getOrderedChildren(g, n)
		if it.Len() != 1 {
			return nil, errors.New("Unexpected number of children")
		}
		children := make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		return newElemUnaryOp(negOpType, children[0]), nil
	}
}

// Square performs a pointwise square.
func Square(a *Node) (*Node, error) { return unaryOpNode(newElemUnaryOp(squareOpType, a), a) }

// SquareOp ...
func NewSquareOperation() Operation {
	return func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		it := getOrderedChildren(g, n)
		if it.Len() != 1 {
			return nil, errors.New("Unexpected number of children")
		}
		children := make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		return newElemUnaryOp(squareOpType, children[0]), nil
	}
}

// Sqrt performs a pointwise sqrt.
func Sqrt(a *Node) (*Node, error) { return unaryOpNode(newElemUnaryOp(sqrtOpType, a), a) }

// SqrtOp ...
func NewSqrtOperation() Operation {
	return func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		it := getOrderedChildren(g, n)
		if it.Len() != 1 {
			return nil, errors.New("Unexpected number of children")
		}
		children := make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		return newElemUnaryOp(sqrtOpType, children[0]), nil
	}
}

// Inverse performs a pointwise inverse.
func Inverse(a *Node) (*Node, error) { return unaryOpNode(newElemUnaryOp(inverseOpType, a), a) }

// InverseOp ...
func NewInverseOperation() Operation {
	return func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		it := getOrderedChildren(g, n)
		if it.Len() != 1 {
			return nil, errors.New("Unexpected number of children")
		}
		children := make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		return newElemUnaryOp(inverseOpType, children[0]), nil
	}
}

// InverseSqrt performs a pointwise inversesqrt.
func InverseSqrt(a *Node) (*Node, error) { return unaryOpNode(newElemUnaryOp(inverseSqrtOpType, a), a) }

// InverseSqrtOp ...
func NewInverseSqrtOperation() Operation {
	return func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		it := getOrderedChildren(g, n)
		if it.Len() != 1 {
			return nil, errors.New("Unexpected number of children")
		}
		children := make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		return newElemUnaryOp(inverseSqrtOpType, children[0]), nil
	}
}

// Cube performs a pointwise cube.
func Cube(a *Node) (*Node, error) { return unaryOpNode(newElemUnaryOp(cubeOpType, a), a) }

// CubeOp ...
func NewCubeOperation() Operation {
	return func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		it := getOrderedChildren(g, n)
		if it.Len() != 1 {
			return nil, errors.New("Unexpected number of children")
		}
		children := make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		return newElemUnaryOp(cubeOpType, children[0]), nil
	}
}

// Tanh performs a pointwise tanh.
func Tanh(a *Node) (*Node, error) { return unaryOpNode(newElemUnaryOp(tanhOpType, a), a) }

// TanhOp ...
func NewTanhOperation() Operation {
	return func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		it := getOrderedChildren(g, n)
		if it.Len() != 1 {
			return nil, errors.New("Unexpected number of children")
		}
		children := make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		return newElemUnaryOp(tanhOpType, children[0]), nil
	}
}

// Sigmoid performs a pointwise sigmoid.
func Sigmoid(a *Node) (*Node, error) { return unaryOpNode(newElemUnaryOp(sigmoidOpType, a), a) }

// SigmoidOp ...
func NewSigmoidOperation() Operation {
	return func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		it := getOrderedChildren(g, n)
		if it.Len() != 1 {
			return nil, errors.New("Unexpected number of children")
		}
		children := make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		return newElemUnaryOp(sigmoidOpType, children[0]), nil
	}
}

// Log1p performs a pointwise log1p.
func Log1p(a *Node) (*Node, error) { return unaryOpNode(newElemUnaryOp(log1pOpType, a), a) }

// Log1pOp ...
func NewLog1pOperation() Operation {
	return func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		it := getOrderedChildren(g, n)
		if it.Len() != 1 {
			return nil, errors.New("Unexpected number of children")
		}
		children := make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		return newElemUnaryOp(log1pOpType, children[0]), nil
	}
}

// Expm1 performs a pointwise expm1.
func Expm1(a *Node) (*Node, error) { return unaryOpNode(newElemUnaryOp(expm1OpType, a), a) }

// Expm1Op ...
func NewExpm1Operation() Operation {
	return func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		it := getOrderedChildren(g, n)
		if it.Len() != 1 {
			return nil, errors.New("Unexpected number of children")
		}
		children := make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		return newElemUnaryOp(expm1OpType, children[0]), nil
	}
}

// Softplus performs a pointwise softplus.
func Softplus(a *Node) (*Node, error) { return unaryOpNode(newElemUnaryOp(softplusOpType, a), a) }

// SoftplusOp ...
func NewSoftplusOperation() Operation {
	return func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		it := getOrderedChildren(g, n)
		if it.Len() != 1 {
			return nil, errors.New("Unexpected number of children")
		}
		children := make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		return newElemUnaryOp(softplusOpType, children[0]), nil
	}
}

// Add perfors a pointwise add operation.
func Add(a, b *Node) (*Node, error) { return binOpNode(newElemBinOp(addOpType, a, b), a, b) }

// START_ADDOP OMIT

// AddOp is broadcastable...
func NewAddOperation(leftAxes, rightAxes []byte) Operation {
	return func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		it := getOrderedChildren(g, n)
		if it.Len() != 2 {
			return nil, errors.New("AddOperation: Unexpected number of children")
		}
		children := make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		x := children[0]
		y := children[1]

		if leftAxes != nil || rightAxes != nil {
			builder, ok := g.(graph.DirectedWeightedBuilder)
			if !ok {
				return nil, errors.Errorf("Broadcast needs to modify the graph but is not a DirectedWeightedBuilder")
			}
			_, ok = g.(graph.EdgeRemover)
			if !ok {
				return nil, errors.Errorf("Broadcast needs to modify the graph but is not an EdgeRemover")
			}

			pattern := newBroadcastPattern(leftAxes, rightAxes)
			broadcastOn := pattern.on()
			switch {
			case len(broadcastOn[0]) != 0:
				// Remove the link from n to x
				g.(graph.EdgeRemover).RemoveEdge(n.ID(), x.ID())
				broadcastedX := builder.NewNode().(*Node)
				broadcastedX.name = n.(*Node).name + "_broadcastedX"
				builder.AddNode(broadcastedX)
				// Link it to the input tensor
				builder.SetWeightedEdge(builder.NewWeightedEdge(n, broadcastedX, 0.0))
				builder.SetWeightedEdge(builder.NewWeightedEdge(broadcastedX, x, 0.0))
				builder.SetWeightedEdge(builder.NewWeightedEdge(broadcastedX, y, 1.0))

				bcastOp := newBroadcastOperation(second, broadcastOn[0])
				err := g.(*ExprGraph).ApplyOp(bcastOp, broadcastedX)
				if err != nil {
					return nil, err
				}
				//x = broadcastedX
			case len(broadcastOn[1]) != 0:
				// Remove the link from n to x
				g.(graph.EdgeRemover).RemoveEdge(n.ID(), y.ID())
				broadcastedY := builder.NewNode().(*Node)
				broadcastedY.name = n.(*Node).name + "_broadcastedY"
				builder.AddNode(broadcastedY)
				// Link it to the input tensor
				builder.SetWeightedEdge(builder.NewWeightedEdge(n, broadcastedY, 0.0))
				builder.SetWeightedEdge(builder.NewWeightedEdge(broadcastedY, x, 0.0))
				builder.SetWeightedEdge(builder.NewWeightedEdge(broadcastedY, y, 1.0))

				bcastOp := newBroadcastOperation(first, broadcastOn[1])
				err := g.(*ExprGraph).ApplyOp(bcastOp, broadcastedY)
				if err != nil {
					return nil, err
				}
				//y = broadcastedY
			}
		}
		it = getOrderedChildren(g, n)
		if it.Len() != 2 {
			return nil, errors.New("AddOperation: Unexpected number of children")
		}
		children = make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		return newElemBinOp(addOpType, children[0], children[1]), nil
	}
}

// END_ADDOP OMIT

// Sub perfors a pointwise sub operation.
func Sub(a, b *Node) (*Node, error) { return binOpNode(newElemBinOp(subOpType, a, b), a, b) }

// SubOp ...
func NewSubOperation() Operation {
	return func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		it := getOrderedChildren(g, n)
		if it.Len() != 2 {
			return nil, errors.New("Unexpected number of children")
		}
		children := make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		return newElemBinOp(subOpType, children[0], children[1]), nil
	}
}

// HadamardProd perfors a pointwise hadamardprod operation.
func HadamardProd(a, b *Node) (*Node, error) { return binOpNode(newElemBinOp(mulOpType, a, b), a, b) }

// HadamardProdOp ...
func NewHadamardProdOperation() Operation {
	return func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		it := getOrderedChildren(g, n)
		if it.Len() != 2 {
			return nil, errors.New("Unexpected number of children")
		}
		children := make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		return newElemBinOp(mulOpType, children[0], children[1]), nil
	}
}

// HadamardDiv perfors a pointwise hadamarddiv operation.
func HadamardDiv(a, b *Node) (*Node, error) { return binOpNode(newElemBinOp(divOpType, a, b), a, b) }

// HadamardDivOp ...
func NewHadamardDivOperation() Operation {
	return func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		it := getOrderedChildren(g, n)
		if it.Len() != 2 {
			return nil, errors.New("Unexpected number of children")
		}
		children := make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		return newElemBinOp(divOpType, children[0], children[1]), nil
	}
}

// Pow perfors a pointwise pow operation.
func Pow(a, b *Node) (*Node, error) { return binOpNode(newElemBinOp(powOpType, a, b), a, b) }

// PowOp ...
func NewPowOperation() Operation {
	return func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		it := getOrderedChildren(g, n)
		if it.Len() != 2 {
			return nil, errors.New("Unexpected number of children")
		}
		children := make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		return newElemBinOp(powOpType, children[0], children[1]), nil
	}
}

// Lt perfors a pointwise lt operation.
// retSame indicates if the data type of the return value should be the same as the input data type. It defaults to Bool otherwise.
func Lt(a, b *Node, retSame bool) (*Node, error) {
	op := newElemBinOp(ltOpType, a, b)
	op.retSame = retSame
	return binOpNode(op, a, b)
}

// LtOp ...
func NewLtOperation(retSame bool) Operation {
	return func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		it := getOrderedChildren(g, n)
		if it.Len() != 2 {
			return nil, errors.New("Unexpected number of children")
		}
		children := make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		op := newElemBinOp(ltOpType, children[0], children[1])
		op.retSame = retSame
		return op, nil
	}
}

// Gt perfors a pointwise gt operation.
// retSame indicates if the data type of the return value should be the same as the input data type. It defaults to Bool otherwise.
func Gt(a, b *Node, retSame bool) (*Node, error) {
	op := newElemBinOp(gtOpType, a, b)
	op.retSame = retSame
	return binOpNode(op, a, b)
}

// GtOp ...
func NewGtOperation(retSame bool) Operation {
	return func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		it := getOrderedChildren(g, n)
		if it.Len() != 2 {
			return nil, errors.New("Unexpected number of children")
		}
		children := make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		op := newElemBinOp(gtOpType, children[0], children[1])
		op.retSame = retSame
		return op, nil
	}
}

// Lte perfors a pointwise lte operation.
// retSame indicates if the data type of the return value should be the same as the input data type. It defaults to Bool otherwise.
func Lte(a, b *Node, retSame bool) (*Node, error) {
	op := newElemBinOp(lteOpType, a, b)
	op.retSame = retSame
	return binOpNode(op, a, b)
}

// LteOp ...
func NewLteOperation(retSame bool) Operation {
	return func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		it := getOrderedChildren(g, n)
		if it.Len() != 2 {
			return nil, errors.New("Unexpected number of children")
		}
		children := make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		op := newElemBinOp(lteOpType, children[0], children[1])
		op.retSame = retSame
		return op, nil
	}
}

// Gte perfors a pointwise gte operation.
// retSame indicates if the data type of the return value should be the same as the input data type. It defaults to Bool otherwise.
func Gte(a, b *Node, retSame bool) (*Node, error) {
	op := newElemBinOp(gteOpType, a, b)
	op.retSame = retSame
	return binOpNode(op, a, b)
}

// GteOp ...
func NewGteOperation(retSame bool) Operation {
	return func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		it := getOrderedChildren(g, n)
		if it.Len() != 2 {
			return nil, errors.New("Unexpected number of children")
		}
		children := make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		op := newElemBinOp(gteOpType, children[0], children[1])
		op.retSame = retSame
		return op, nil
	}
}

// Eq perfors a pointwise eq operation.
// retSame indicates if the data type of the return value should be the same as the input data type. It defaults to Bool otherwise.
func Eq(a, b *Node, retSame bool) (*Node, error) {
	op := newElemBinOp(eqOpType, a, b)
	op.retSame = retSame
	return binOpNode(op, a, b)
}

// EqOp ...
func NewEqOperation(retSame bool) Operation {
	return func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		it := getOrderedChildren(g, n)
		if it.Len() != 2 {
			return nil, errors.New("Unexpected number of children")
		}
		children := make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		op := newElemBinOp(eqOpType, children[0], children[1])
		op.retSame = retSame
		return op, nil
	}
}

// Ne perfors a pointwise ne operation.
// retSame indicates if the data type of the return value should be the same as the input data type. It defaults to Bool otherwise.
func Ne(a, b *Node, retSame bool) (*Node, error) {
	op := newElemBinOp(neOpType, a, b)
	op.retSame = retSame
	return binOpNode(op, a, b)
}

// NeOp ...
func NewNeOperation(retSame bool) Operation {
	return func(g graph.WeightedDirected, n node.Node) (ops.Op, error) {
		it := getOrderedChildren(g, n)
		if it.Len() != 2 {
			return nil, errors.New("Unexpected number of children")
		}
		children := make([]*Node, it.Len())
		for i := 0; it.Next(); i++ {
			children[i] = it.Node().(*Node)
		}
		op := newElemBinOp(neOpType, children[0], children[1])
		op.retSame = retSame
		return op, nil
	}
}
