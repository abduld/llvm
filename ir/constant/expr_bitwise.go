// generated by gen.go using 'go generate'; DO NOT EDIT.

// === [ Bitwise expressions ] =================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#bitwise-binary-operations

package constant

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
)

// --- [ shl ] -----------------------------------------------------------------

// ExprShl represents a shift left expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#shl-instruction
type ExprShl struct {
	// Operands.
	x, y Constant
}

// NewShl returns a new shl expression based on the given operands.
func NewShl(x, y Constant) *ExprShl {
	return &ExprShl{x: x, y: y}
}

// Type returns the type of the constant expression.
func (expr *ExprShl) Type() types.Type {
	return expr.x.Type()
}

// Ident returns the string representation of the constant expression.
func (expr *ExprShl) Ident() string {
	x, y := expr.X(), expr.Y()
	return fmt.Sprintf("shl (%s %s, %s %s)",
		x.Type(),
		x.Ident(),
		y.Type(),
		y.Ident())
}

// Immutable ensures that only constants can be assigned to the Constant
// interface.
func (*ExprShl) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprShl) Simplify() Constant {
	panic("not yet implemented")
}

// X returns the x operand of the shl expression.
func (expr *ExprShl) X() Constant {
	return expr.x
}

// Y returns the y operand of the shl expression.
func (expr *ExprShl) Y() Constant {
	return expr.y
}

// --- [ lshr ] ----------------------------------------------------------------

// ExprLShr represents a logical shift right expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#lshr-instruction
type ExprLShr struct {
	// Operands.
	x, y Constant
}

// NewLShr returns a new lshr expression based on the given operands.
func NewLShr(x, y Constant) *ExprLShr {
	return &ExprLShr{x: x, y: y}
}

// Type returns the type of the constant expression.
func (expr *ExprLShr) Type() types.Type {
	return expr.x.Type()
}

// Ident returns the string representation of the constant expression.
func (expr *ExprLShr) Ident() string {
	x, y := expr.X(), expr.Y()
	return fmt.Sprintf("lshr (%s %s, %s %s)",
		x.Type(),
		x.Ident(),
		y.Type(),
		y.Ident())
}

// Immutable ensures that only constants can be assigned to the Constant
// interface.
func (*ExprLShr) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprLShr) Simplify() Constant {
	panic("not yet implemented")
}

// X returns the x operand of the lshr expression.
func (expr *ExprLShr) X() Constant {
	return expr.x
}

// Y returns the y operand of the lshr expression.
func (expr *ExprLShr) Y() Constant {
	return expr.y
}

// --- [ ashr ] ----------------------------------------------------------------

// ExprAShr represents an arithmetic shift right expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#ashr-instruction
type ExprAShr struct {
	// Operands.
	x, y Constant
}

// NewAShr returns a new ashr expression based on the given operands.
func NewAShr(x, y Constant) *ExprAShr {
	return &ExprAShr{x: x, y: y}
}

// Type returns the type of the constant expression.
func (expr *ExprAShr) Type() types.Type {
	return expr.x.Type()
}

// Ident returns the string representation of the constant expression.
func (expr *ExprAShr) Ident() string {
	x, y := expr.X(), expr.Y()
	return fmt.Sprintf("ashr (%s %s, %s %s)",
		x.Type(),
		x.Ident(),
		y.Type(),
		y.Ident())
}

// Immutable ensures that only constants can be assigned to the Constant
// interface.
func (*ExprAShr) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprAShr) Simplify() Constant {
	panic("not yet implemented")
}

// X returns the x operand of the ashr expression.
func (expr *ExprAShr) X() Constant {
	return expr.x
}

// Y returns the y operand of the ashr expression.
func (expr *ExprAShr) Y() Constant {
	return expr.y
}

// --- [ and ] -----------------------------------------------------------------

// ExprAnd represents an AND expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#and-instruction
type ExprAnd struct {
	// Operands.
	x, y Constant
}

// NewAnd returns a new and expression based on the given operands.
func NewAnd(x, y Constant) *ExprAnd {
	return &ExprAnd{x: x, y: y}
}

// Type returns the type of the constant expression.
func (expr *ExprAnd) Type() types.Type {
	return expr.x.Type()
}

// Ident returns the string representation of the constant expression.
func (expr *ExprAnd) Ident() string {
	x, y := expr.X(), expr.Y()
	return fmt.Sprintf("and (%s %s, %s %s)",
		x.Type(),
		x.Ident(),
		y.Type(),
		y.Ident())
}

// Immutable ensures that only constants can be assigned to the Constant
// interface.
func (*ExprAnd) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprAnd) Simplify() Constant {
	panic("not yet implemented")
}

// X returns the x operand of the and expression.
func (expr *ExprAnd) X() Constant {
	return expr.x
}

// Y returns the y operand of the and expression.
func (expr *ExprAnd) Y() Constant {
	return expr.y
}

// --- [ or ] ------------------------------------------------------------------

// ExprOr represents an OR expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#or-instruction
type ExprOr struct {
	// Operands.
	x, y Constant
}

// NewOr returns a new or expression based on the given operands.
func NewOr(x, y Constant) *ExprOr {
	return &ExprOr{x: x, y: y}
}

// Type returns the type of the constant expression.
func (expr *ExprOr) Type() types.Type {
	return expr.x.Type()
}

// Ident returns the string representation of the constant expression.
func (expr *ExprOr) Ident() string {
	x, y := expr.X(), expr.Y()
	return fmt.Sprintf("or (%s %s, %s %s)",
		x.Type(),
		x.Ident(),
		y.Type(),
		y.Ident())
}

// Immutable ensures that only constants can be assigned to the Constant
// interface.
func (*ExprOr) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprOr) Simplify() Constant {
	panic("not yet implemented")
}

// X returns the x operand of the or expression.
func (expr *ExprOr) X() Constant {
	return expr.x
}

// Y returns the y operand of the or expression.
func (expr *ExprOr) Y() Constant {
	return expr.y
}

// --- [ xor ] -----------------------------------------------------------------

// ExprXor represents an exclusive-OR expression.
//
// References:
//    http://llvm.org/docs/LangRef.html#xor-instruction
type ExprXor struct {
	// Operands.
	x, y Constant
}

// NewXor returns a new xor expression based on the given operands.
func NewXor(x, y Constant) *ExprXor {
	return &ExprXor{x: x, y: y}
}

// Type returns the type of the constant expression.
func (expr *ExprXor) Type() types.Type {
	return expr.x.Type()
}

// Ident returns the string representation of the constant expression.
func (expr *ExprXor) Ident() string {
	x, y := expr.X(), expr.Y()
	return fmt.Sprintf("xor (%s %s, %s %s)",
		x.Type(),
		x.Ident(),
		y.Type(),
		y.Ident())
}

// Immutable ensures that only constants can be assigned to the Constant
// interface.
func (*ExprXor) Immutable() {}

// Simplify returns a simplified version of the constant expression.
func (expr *ExprXor) Simplify() Constant {
	panic("not yet implemented")
}

// X returns the x operand of the xor expression.
func (expr *ExprXor) X() Constant {
	return expr.x
}

// Y returns the y operand of the xor expression.
func (expr *ExprXor) Y() Constant {
	return expr.y
}
