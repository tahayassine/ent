// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/examples/start/ent/car"
	"github.com/facebookincubator/ent/examples/start/ent/group"
	"github.com/facebookincubator/ent/examples/start/ent/user"
)

const (
	// Operation types.
	OpCreate    = "Create"
	OpDelete    = "Delete"
	OpDeleteOne = "DeleteOne"
	OpUpdate    = "Update"
	OpUpdateOne = "UpdateOne"

	// Node types.
	TypeCar   = "Car"
	TypeGroup = "Group"
	TypeUser  = "User"
)

// CarMutation represents an operation that mutate the Cars
// nodes in the graph.
type CarMutation struct {
	config
	op, typ       string
	id            *int
	model         *string
	registered_at *time.Time
	clearedFields map[string]bool
	owner         map[int]struct{}
	clearedowner  bool
}

var _ ent.Mutation = (*CarMutation)(nil)

// newCarMutation creates new mutation for $n.Name.
func newCarMutation(c config, op string) *CarMutation {
	return &CarMutation{
		op:            op,
		typ:           TypeCar,
		clearedFields: make(map[string]bool),
	}
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *CarMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetModel sets the model field.
func (m *CarMutation) SetModel(s string) {
	m.model = &s
}

// Model returns the model value in the mutation.
func (m *CarMutation) Model() (r string, exists bool) {
	v := m.model
	if v == nil {
		return
	}
	return *v, true
}

// ResetModel reset all changes of the model field.
func (m *CarMutation) ResetModel() {
	m.model = nil
}

// SetRegisteredAt sets the registered_at field.
func (m *CarMutation) SetRegisteredAt(t time.Time) {
	m.registered_at = &t
}

// RegisteredAt returns the registered_at value in the mutation.
func (m *CarMutation) RegisteredAt() (r time.Time, exists bool) {
	v := m.registered_at
	if v == nil {
		return
	}
	return *v, true
}

// ResetRegisteredAt reset all changes of the registered_at field.
func (m *CarMutation) ResetRegisteredAt() {
	m.registered_at = nil
}

// SetOwnerID sets the owner edge to User by id.
func (m *CarMutation) SetOwnerID(id int) {
	if m.owner == nil {
		m.owner = make(map[int]struct{})
	}
	m.owner[id] = struct{}{}
}

// ClearOwner clears the owner edge to User.
func (m *CarMutation) ClearOwner() {
	m.clearedowner = true
}

// OwnerCleared returns if the edge owner was cleared.
func (m *CarMutation) OwnerCleared() bool {
	return m.clearedowner
}

// OwnerIDs returns the owner ids in the mutation.
func (m *CarMutation) OwnerIDs() (ids []int) {
	for id := range m.owner {
		ids = append(ids, id)
	}
	return
}

// ResetOwner reset all changes of the owner edge.
func (m *CarMutation) ResetOwner() {
	m.owner = nil
	m.clearedowner = false
}

// Op returns the operation name.
func (m *CarMutation) Op() string {
	return m.op
}

// Type returns the node type of this mutation (Car).
func (m *CarMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *CarMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.model != nil {
		fields = append(fields, car.FieldModel)
	}
	if m.registered_at != nil {
		fields = append(fields, car.FieldRegisteredAt)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *CarMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case car.FieldModel:
		return m.Model()
	case car.FieldRegisteredAt:
		return m.RegisteredAt()
	}
	return nil, false
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *CarMutation) SetField(name string, value ent.Value) error {
	switch name {
	case car.FieldModel:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetModel(v)
		return nil
	case car.FieldRegisteredAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRegisteredAt(v)
		return nil
	}
	return fmt.Errorf("unknown Car field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *CarMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *CarMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *CarMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Car numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *CarMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *CarMutation) FieldCleared(name string) bool {
	return m.clearedFields[name]
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *CarMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Car nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *CarMutation) ResetField(name string) error {
	switch name {
	case car.FieldModel:
		m.ResetModel()
		return nil
	case car.FieldRegisteredAt:
		m.ResetRegisteredAt()
		return nil
	}
	return fmt.Errorf("unknown Car field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *CarMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.owner != nil {
		edges = append(edges, car.EdgeOwner)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *CarMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case car.EdgeOwner:
		ids := make([]int, 0, len(m.owner))
		for id := range m.owner {
			ids = append(ids, id)
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *CarMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *CarMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *CarMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedowner {
		edges = append(edges, car.EdgeOwner)
	}
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *CarMutation) EdgeCleared(name string) bool {
	switch name {
	case car.EdgeOwner:
		return m.clearedowner
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *CarMutation) ClearEdge(name string) error {
	switch name {
	case car.EdgeOwner:
		m.ClearOwner()
		return nil
	}
	return fmt.Errorf("unknown Car unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *CarMutation) ResetEdge(name string) error {
	switch name {
	case car.EdgeOwner:
		m.ResetOwner()
		return nil
	}
	return fmt.Errorf("unknown Car edge %s", name)
}

// GroupMutation represents an operation that mutate the Groups
// nodes in the graph.
type GroupMutation struct {
	config
	op, typ       string
	id            *int
	name          *string
	clearedFields map[string]bool
	users         map[int]struct{}
	removedusers  map[int]struct{}
}

var _ ent.Mutation = (*GroupMutation)(nil)

// newGroupMutation creates new mutation for $n.Name.
func newGroupMutation(c config, op string) *GroupMutation {
	return &GroupMutation{
		op:            op,
		typ:           TypeGroup,
		clearedFields: make(map[string]bool),
	}
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *GroupMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the name field.
func (m *GroupMutation) SetName(s string) {
	m.name = &s
}

// Name returns the name value in the mutation.
func (m *GroupMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// ResetName reset all changes of the name field.
func (m *GroupMutation) ResetName() {
	m.name = nil
}

// AddUserIDs adds the users edge to User by ids.
func (m *GroupMutation) AddUserIDs(ids ...int) {
	if m.users == nil {
		m.users = make(map[int]struct{})
	}
	for i := range ids {
		m.users[ids[i]] = struct{}{}
	}
}

// RemoveUserIDs removes the users edge to User by ids.
func (m *GroupMutation) RemoveUserIDs(ids ...int) {
	if m.removedusers == nil {
		m.removedusers = make(map[int]struct{})
	}
	for i := range ids {
		m.removedusers[ids[i]] = struct{}{}
	}
}

// RemovedUsers returns the removed ids of users.
func (m *GroupMutation) RemovedUsersIDs() (ids []int) {
	for id := range m.removedusers {
		ids = append(ids, id)
	}
	return
}

// UsersIDs returns the users ids in the mutation.
func (m *GroupMutation) UsersIDs() (ids []int) {
	for id := range m.users {
		ids = append(ids, id)
	}
	return
}

// ResetUsers reset all changes of the users edge.
func (m *GroupMutation) ResetUsers() {
	m.users = nil
	m.removedusers = nil
}

// Op returns the operation name.
func (m *GroupMutation) Op() string {
	return m.op
}

// Type returns the node type of this mutation (Group).
func (m *GroupMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *GroupMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, group.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *GroupMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case group.FieldName:
		return m.Name()
	}
	return nil, false
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *GroupMutation) SetField(name string, value ent.Value) error {
	switch name {
	case group.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown Group field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *GroupMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *GroupMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *GroupMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Group numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *GroupMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *GroupMutation) FieldCleared(name string) bool {
	return m.clearedFields[name]
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *GroupMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Group nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *GroupMutation) ResetField(name string) error {
	switch name {
	case group.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown Group field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *GroupMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.users != nil {
		edges = append(edges, group.EdgeUsers)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *GroupMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case group.EdgeUsers:
		ids := make([]int, 0, len(m.users))
		for id := range m.users {
			ids = append(ids, id)
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *GroupMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedusers != nil {
		edges = append(edges, group.EdgeUsers)
	}
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *GroupMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case group.EdgeUsers:
		ids := make([]int, 0, len(m.removedusers))
		for id := range m.removedusers {
			ids = append(ids, id)
		}
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *GroupMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *GroupMutation) EdgeCleared(name string) bool {
	switch name {
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *GroupMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Group unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *GroupMutation) ResetEdge(name string) error {
	switch name {
	case group.EdgeUsers:
		m.ResetUsers()
		return nil
	}
	return fmt.Errorf("unknown Group edge %s", name)
}

// UserMutation represents an operation that mutate the Users
// nodes in the graph.
type UserMutation struct {
	config
	op, typ       string
	id            *int
	age           *int
	addage        *int
	name          *string
	clearedFields map[string]bool
	cars          map[int]struct{}
	removedcars   map[int]struct{}
	groups        map[int]struct{}
	removedgroups map[int]struct{}
}

var _ ent.Mutation = (*UserMutation)(nil)

// newUserMutation creates new mutation for $n.Name.
func newUserMutation(c config, op string) *UserMutation {
	return &UserMutation{
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]bool),
	}
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetAge sets the age field.
func (m *UserMutation) SetAge(i int) {
	m.age = &i
	m.addage = nil
}

// Age returns the age value in the mutation.
func (m *UserMutation) Age() (r int, exists bool) {
	v := m.age
	if v == nil {
		return
	}
	return *v, true
}

// AddAge adds i to age.
func (m *UserMutation) AddAge(i int) {
	if m.addage != nil {
		*m.addage += i
	} else {
		m.addage = &i
	}
}

// AddedAge returns the value that was added to the age field in this mutation.
func (m *UserMutation) AddedAge() (r int, exists bool) {
	v := m.addage
	if v == nil {
		return
	}
	return *v, true
}

// ResetAge reset all changes of the age field.
func (m *UserMutation) ResetAge() {
	m.age = nil
	m.addage = nil
}

// SetName sets the name field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the name value in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// ResetName reset all changes of the name field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// AddCarIDs adds the cars edge to Car by ids.
func (m *UserMutation) AddCarIDs(ids ...int) {
	if m.cars == nil {
		m.cars = make(map[int]struct{})
	}
	for i := range ids {
		m.cars[ids[i]] = struct{}{}
	}
}

// RemoveCarIDs removes the cars edge to Car by ids.
func (m *UserMutation) RemoveCarIDs(ids ...int) {
	if m.removedcars == nil {
		m.removedcars = make(map[int]struct{})
	}
	for i := range ids {
		m.removedcars[ids[i]] = struct{}{}
	}
}

// RemovedCars returns the removed ids of cars.
func (m *UserMutation) RemovedCarsIDs() (ids []int) {
	for id := range m.removedcars {
		ids = append(ids, id)
	}
	return
}

// CarsIDs returns the cars ids in the mutation.
func (m *UserMutation) CarsIDs() (ids []int) {
	for id := range m.cars {
		ids = append(ids, id)
	}
	return
}

// ResetCars reset all changes of the cars edge.
func (m *UserMutation) ResetCars() {
	m.cars = nil
	m.removedcars = nil
}

// AddGroupIDs adds the groups edge to Group by ids.
func (m *UserMutation) AddGroupIDs(ids ...int) {
	if m.groups == nil {
		m.groups = make(map[int]struct{})
	}
	for i := range ids {
		m.groups[ids[i]] = struct{}{}
	}
}

// RemoveGroupIDs removes the groups edge to Group by ids.
func (m *UserMutation) RemoveGroupIDs(ids ...int) {
	if m.removedgroups == nil {
		m.removedgroups = make(map[int]struct{})
	}
	for i := range ids {
		m.removedgroups[ids[i]] = struct{}{}
	}
}

// RemovedGroups returns the removed ids of groups.
func (m *UserMutation) RemovedGroupsIDs() (ids []int) {
	for id := range m.removedgroups {
		ids = append(ids, id)
	}
	return
}

// GroupsIDs returns the groups ids in the mutation.
func (m *UserMutation) GroupsIDs() (ids []int) {
	for id := range m.groups {
		ids = append(ids, id)
	}
	return
}

// ResetGroups reset all changes of the groups edge.
func (m *UserMutation) ResetGroups() {
	m.groups = nil
	m.removedgroups = nil
}

// Op returns the operation name.
func (m *UserMutation) Op() string {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.age != nil {
		fields = append(fields, user.FieldAge)
	}
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldAge:
		return m.Age()
	case user.FieldName:
		return m.Name()
	}
	return nil, false
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAge(v)
		return nil
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *UserMutation) AddedFields() []string {
	var fields []string
	if m.addage != nil {
		fields = append(fields, user.FieldAge)
	}
	return fields
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case user.FieldAge:
		return m.AddedAge()
	}
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	case user.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAge(v)
		return nil
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	return m.clearedFields[name]
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldAge:
		m.ResetAge()
		return nil
	case user.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.cars != nil {
		edges = append(edges, user.EdgeCars)
	}
	if m.groups != nil {
		edges = append(edges, user.EdgeGroups)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeCars:
		ids := make([]int, 0, len(m.cars))
		for id := range m.cars {
			ids = append(ids, id)
		}
	case user.EdgeGroups:
		ids := make([]int, 0, len(m.groups))
		for id := range m.groups {
			ids = append(ids, id)
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	if m.removedcars != nil {
		edges = append(edges, user.EdgeCars)
	}
	if m.removedgroups != nil {
		edges = append(edges, user.EdgeGroups)
	}
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeCars:
		ids := make([]int, 0, len(m.removedcars))
		for id := range m.removedcars {
			ids = append(ids, id)
		}
	case user.EdgeGroups:
		ids := make([]int, 0, len(m.removedgroups))
		for id := range m.removedgroups {
			ids = append(ids, id)
		}
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeCars:
		m.ResetCars()
		return nil
	case user.EdgeGroups:
		m.ResetGroups()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
