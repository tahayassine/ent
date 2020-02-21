// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/examples/edgeindex/ent/city"
	"github.com/facebookincubator/ent/examples/edgeindex/ent/predicate"
	"github.com/facebookincubator/ent/schema/field"
)

// CityDelete is the builder for deleting a City entity.
type CityDelete struct {
	config
	hooks      []Hook
	mutation   *CityMutation
	predicates []predicate.City
}

// Where adds a new predicate to the delete builder.
func (cd *CityDelete) Where(ps ...predicate.City) *CityDelete {
	cd.predicates = append(cd.predicates, ps...)
	return cd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cd *CityDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cd.hooks) == 0 {
		affected, err = cd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cd.mutation = mutation
			affected, err = cd.sqlExec(ctx)
			return affected, err
		})
		for _, hook := range cd.hooks {
			mut = hook(mut)
		}
		if _, err := mut.Mutate(ctx, cd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cd *CityDelete) ExecX(ctx context.Context) int {
	n, err := cd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cd *CityDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: city.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: city.FieldID,
			},
		},
	}
	if ps := cd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, cd.driver, _spec)
}

// CityDeleteOne is the builder for deleting a single City entity.
type CityDeleteOne struct {
	cd *CityDelete
}

// Exec executes the deletion query.
func (cdo *CityDeleteOne) Exec(ctx context.Context) error {
	n, err := cdo.cd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{city.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cdo *CityDeleteOne) ExecX(ctx context.Context) {
	cdo.cd.ExecX(ctx)
}
