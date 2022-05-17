// Code generated from SimpleBoolean.g4 by ANTLR 4.10.1. DO NOT EDIT.

package simpleboolean // SimpleBoolean
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by SimpleBooleanParser.
type SimpleBooleanVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SimpleBooleanParser#parse.
	VisitParse(ctx *ParseContext) interface{}

	// Visit a parse tree produced by SimpleBooleanParser#binaryExpression.
	VisitBinaryExpression(ctx *BinaryExpressionContext) interface{}

	// Visit a parse tree produced by SimpleBooleanParser#decimalExpression.
	VisitDecimalExpression(ctx *DecimalExpressionContext) interface{}

	// Visit a parse tree produced by SimpleBooleanParser#boolExpression.
	VisitBoolExpression(ctx *BoolExpressionContext) interface{}

	// Visit a parse tree produced by SimpleBooleanParser#conditionalExpression.
	VisitConditionalExpression(ctx *ConditionalExpressionContext) interface{}

	// Visit a parse tree produced by SimpleBooleanParser#calcExpression.
	VisitCalcExpression(ctx *CalcExpressionContext) interface{}

	// Visit a parse tree produced by SimpleBooleanParser#sqrtExpression.
	VisitSqrtExpression(ctx *SqrtExpressionContext) interface{}

	// Visit a parse tree produced by SimpleBooleanParser#notExpression.
	VisitNotExpression(ctx *NotExpressionContext) interface{}

	// Visit a parse tree produced by SimpleBooleanParser#parenExpression.
	VisitParenExpression(ctx *ParenExpressionContext) interface{}

	// Visit a parse tree produced by SimpleBooleanParser#deprecatedConditionalExpression.
	VisitDeprecatedConditionalExpression(ctx *DeprecatedConditionalExpressionContext) interface{}

	// Visit a parse tree produced by SimpleBooleanParser#comparatorExpression.
	VisitComparatorExpression(ctx *ComparatorExpressionContext) interface{}

	// Visit a parse tree produced by SimpleBooleanParser#comparator.
	VisitComparator(ctx *ComparatorContext) interface{}

	// Visit a parse tree produced by SimpleBooleanParser#binary.
	VisitBinary(ctx *BinaryContext) interface{}

	// Visit a parse tree produced by SimpleBooleanParser#boolean.
	VisitBoolean(ctx *BooleanContext) interface{}
}
