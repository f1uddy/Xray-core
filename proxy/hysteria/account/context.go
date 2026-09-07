package account

import "context"

type validatorKey struct{}

func ContextWithValidator(ctx context.Context, v *Validator) context.Context {
	return context.WithValue(ctx, validatorKey{}, v)
}

func ValidatorFromContext(ctx context.Context) *Validator {
	v, _ := ctx.Value(validatorKey{}).(*Validator)
	return v
}
