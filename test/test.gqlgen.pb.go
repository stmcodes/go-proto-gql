package test

import (
	context "context"
)

type ServiceResolvers struct{ Service ServiceServer }

func (s *ServiceResolvers) ServiceTestQ(ctx context.Context, in *Hello) (*OutHello, error) {
	return s.Service.TestQ(ctx, in)
}
func (s *ServiceResolvers) ServiceTestM(ctx context.Context, in *Hello) (*Hello, error) {
	return s.Service.TestM(ctx, in)
}

type QueryResolvers struct{ Service QueryServer }

func (s *QueryResolvers) QueryQuery1(ctx context.Context, in *Hello) (*Hello, error) {
	return s.Service.Query1(ctx, in)
}
func (s *QueryResolvers) QueryQuery2(ctx context.Context, in *Hello) (*Hello, error) {
	return s.Service.Query2(ctx, in)
}
func (s *QueryResolvers) QueryMutate1(ctx context.Context, in *Hello) (*Hello, error) {
	return s.Service.Mutate1(ctx, in)
}

type HelloInput = Hello
type OutHelloInput = OutHello
