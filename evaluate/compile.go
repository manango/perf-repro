package evaluate

import (
	"context"

	"eval/evaluate/simpleboolean"
)

type DSLRuleEvaluator func() interface{}

func Compile(ctx context.Context, expr string) (DSLRuleEvaluator, error) {
	parseCtx, err := simpleboolean.ParseExpression(expr)
	if err != nil {
		return nil, err
	}

	return func() interface{} {
		v := &simpleboolean.BaseSimpleBooleanVisitor{}
		res := parseCtx.Accept(v)
		return res
	}, nil
}
