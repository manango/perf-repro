// Code generated from SimpleBoolean.g4 by ANTLR 4.10.1. DO NOT EDIT.

package simpleboolean // SimpleBoolean
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseSimpleBooleanVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSimpleBooleanVisitor) VisitParse(ctx *ParseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleBooleanVisitor) VisitBinaryExpression(ctx *BinaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleBooleanVisitor) VisitDecimalExpression(ctx *DecimalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleBooleanVisitor) VisitBoolExpression(ctx *BoolExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleBooleanVisitor) VisitConditionalExpression(ctx *ConditionalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleBooleanVisitor) VisitCalcExpression(ctx *CalcExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleBooleanVisitor) VisitSqrtExpression(ctx *SqrtExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleBooleanVisitor) VisitNotExpression(ctx *NotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleBooleanVisitor) VisitParenExpression(ctx *ParenExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleBooleanVisitor) VisitDeprecatedConditionalExpression(ctx *DeprecatedConditionalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleBooleanVisitor) VisitComparatorExpression(ctx *ComparatorExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleBooleanVisitor) VisitComparator(ctx *ComparatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleBooleanVisitor) VisitBinary(ctx *BinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleBooleanVisitor) VisitBoolean(ctx *BooleanContext) interface{} {
	return v.VisitChildren(ctx)
}
