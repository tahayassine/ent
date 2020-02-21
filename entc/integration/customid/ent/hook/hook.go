// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/entc/integration/customid/ent"
)

type BlobFunc func(context.Context, *ent.BlobMutation) (ent.Value, error)

func (f BlobFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.BlobMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BlobMutation", m)
	}
	return f(ctx, mv)
}

type CarFunc func(context.Context, *ent.CarMutation) (ent.Value, error)

func (f CarFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CarMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CarMutation", m)
	}
	return f(ctx, mv)
}

type GroupFunc func(context.Context, *ent.GroupMutation) (ent.Value, error)

func (f GroupFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GroupMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GroupMutation", m)
	}
	return f(ctx, mv)
}

type PetFunc func(context.Context, *ent.PetMutation) (ent.Value, error)

func (f PetFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PetMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PetMutation", m)
	}
	return f(ctx, mv)
}

type UserFunc func(context.Context, *ent.UserMutation) (ent.Value, error)

func (f UserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.UserMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserMutation", m)
	}
	return f(ctx, mv)
}

func On(hk ent.Hook, ops ...string) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			for i := range ops {
				if ops[i] == m.Op() {
					return hk(next).Mutate(ctx, m)
				}
			}
			return next.Mutate(ctx, m)
		})
	}
}
