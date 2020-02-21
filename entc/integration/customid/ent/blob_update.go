// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/customid/ent/blob"
	"github.com/facebookincubator/ent/entc/integration/customid/ent/predicate"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
)

// BlobUpdate is the builder for updating Blob entities.
type BlobUpdate struct {
	config
	hooks      []Hook
	mutation   *BlobMutation
	predicates []predicate.Blob
}

// Where adds a new predicate for the builder.
func (bu *BlobUpdate) Where(ps ...predicate.Blob) *BlobUpdate {
	bu.predicates = append(bu.predicates, ps...)
	return bu
}

// SetUUID sets the uuid field.
func (bu *BlobUpdate) SetUUID(u uuid.UUID) *BlobUpdate {
	bu.mutation.SetUUID(u)
	return bu
}

// SetParentID sets the parent edge to Blob by id.
func (bu *BlobUpdate) SetParentID(id uuid.UUID) *BlobUpdate {
	bu.mutation.SetParentID(id)
	return bu
}

// SetNillableParentID sets the parent edge to Blob by id if the given value is not nil.
func (bu *BlobUpdate) SetNillableParentID(id *uuid.UUID) *BlobUpdate {
	if id != nil {
		bu = bu.SetParentID(*id)
	}
	return bu
}

// SetParent sets the parent edge to Blob.
func (bu *BlobUpdate) SetParent(b *Blob) *BlobUpdate {
	return bu.SetParentID(b.ID)
}

// AddLinkIDs adds the links edge to Blob by ids.
func (bu *BlobUpdate) AddLinkIDs(ids ...uuid.UUID) *BlobUpdate {
	bu.mutation.AddLinkIDs(ids...)
	return bu
}

// AddLinks adds the links edges to Blob.
func (bu *BlobUpdate) AddLinks(b ...*Blob) *BlobUpdate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bu.AddLinkIDs(ids...)
}

// ClearParent clears the parent edge to Blob.
func (bu *BlobUpdate) ClearParent() *BlobUpdate {
	bu.mutation.ClearParent()
	return bu
}

// RemoveLinkIDs removes the links edge to Blob by ids.
func (bu *BlobUpdate) RemoveLinkIDs(ids ...uuid.UUID) *BlobUpdate {
	bu.mutation.RemoveLinkIDs(ids...)
	return bu
}

// RemoveLinks removes links edges to Blob.
func (bu *BlobUpdate) RemoveLinks(b ...*Blob) *BlobUpdate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bu.RemoveLinkIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (bu *BlobUpdate) Save(ctx context.Context) (int, error) {
	if len(bu.mutation.ParentIDs()) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"parent\"")
	}
	var (
		err      error
		affected int
	)
	if len(bu.hooks) == 0 {
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BlobMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bu.mutation = mutation
			affected, err = bu.sqlSave(ctx)
			return affected, err
		})
		for _, hook := range bu.hooks {
			mut = hook(mut)
		}
		if _, err := mut.Mutate(ctx, bu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BlobUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BlobUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BlobUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BlobUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   blob.Table,
			Columns: blob.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: blob.FieldID,
			},
		},
	}
	if ps := bu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: blob.FieldUUID,
		})
	}
	if bu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   blob.ParentTable,
			Columns: []string{blob.ParentColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: blob.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   blob.ParentTable,
			Columns: []string{blob.ParentColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: blob.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := bu.mutation.RemovedLinksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   blob.LinksTable,
			Columns: blob.LinksPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: blob.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.LinksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   blob.LinksTable,
			Columns: blob.LinksPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: blob.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blob.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// BlobUpdateOne is the builder for updating a single Blob entity.
type BlobUpdateOne struct {
	config
	hooks    []Hook
	mutation *BlobMutation
}

// SetUUID sets the uuid field.
func (buo *BlobUpdateOne) SetUUID(u uuid.UUID) *BlobUpdateOne {
	buo.mutation.SetUUID(u)
	return buo
}

// SetParentID sets the parent edge to Blob by id.
func (buo *BlobUpdateOne) SetParentID(id uuid.UUID) *BlobUpdateOne {
	buo.mutation.SetParentID(id)
	return buo
}

// SetNillableParentID sets the parent edge to Blob by id if the given value is not nil.
func (buo *BlobUpdateOne) SetNillableParentID(id *uuid.UUID) *BlobUpdateOne {
	if id != nil {
		buo = buo.SetParentID(*id)
	}
	return buo
}

// SetParent sets the parent edge to Blob.
func (buo *BlobUpdateOne) SetParent(b *Blob) *BlobUpdateOne {
	return buo.SetParentID(b.ID)
}

// AddLinkIDs adds the links edge to Blob by ids.
func (buo *BlobUpdateOne) AddLinkIDs(ids ...uuid.UUID) *BlobUpdateOne {
	buo.mutation.AddLinkIDs(ids...)
	return buo
}

// AddLinks adds the links edges to Blob.
func (buo *BlobUpdateOne) AddLinks(b ...*Blob) *BlobUpdateOne {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return buo.AddLinkIDs(ids...)
}

// ClearParent clears the parent edge to Blob.
func (buo *BlobUpdateOne) ClearParent() *BlobUpdateOne {
	buo.mutation.ClearParent()
	return buo
}

// RemoveLinkIDs removes the links edge to Blob by ids.
func (buo *BlobUpdateOne) RemoveLinkIDs(ids ...uuid.UUID) *BlobUpdateOne {
	buo.mutation.RemoveLinkIDs(ids...)
	return buo
}

// RemoveLinks removes links edges to Blob.
func (buo *BlobUpdateOne) RemoveLinks(b ...*Blob) *BlobUpdateOne {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return buo.RemoveLinkIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (buo *BlobUpdateOne) Save(ctx context.Context) (*Blob, error) {
	if len(buo.mutation.ParentIDs()) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"parent\"")
	}
	var (
		err  error
		node *Blob
	)
	if len(buo.hooks) == 0 {
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BlobMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			buo.mutation = mutation
			node, err = buo.sqlSave(ctx)
			return node, err
		})
		for _, hook := range buo.hooks {
			mut = hook(mut)
		}
		if _, err := mut.Mutate(ctx, buo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BlobUpdateOne) SaveX(ctx context.Context) *Blob {
	b, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return b
}

// Exec executes the query on the entity.
func (buo *BlobUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BlobUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BlobUpdateOne) sqlSave(ctx context.Context) (b *Blob, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   blob.Table,
			Columns: blob.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: blob.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, fmt.Errorf("missing Blob.ID for update")
	}
	_spec.Node.ID.Value = id
	if value, ok := buo.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: blob.FieldUUID,
		})
	}
	if buo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   blob.ParentTable,
			Columns: []string{blob.ParentColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: blob.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   blob.ParentTable,
			Columns: []string{blob.ParentColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: blob.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := buo.mutation.RemovedLinksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   blob.LinksTable,
			Columns: blob.LinksPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: blob.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.LinksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   blob.LinksTable,
			Columns: blob.LinksPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: blob.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	b = &Blob{config: buo.config}
	_spec.Assign = b.assignValues
	_spec.ScanValues = b.scanValues()
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blob.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return b, nil
}
