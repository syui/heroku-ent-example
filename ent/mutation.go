// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"t/ent/predicate"
	"t/ent/todo"
	"t/ent/users"
	"time"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeTodo  = "Todo"
	TypeUsers = "Users"
)

// TodoMutation represents an operation that mutates the Todo nodes in the graph.
type TodoMutation struct {
	config
	op            Op
	typ           string
	id            *int
	title         *string
	_done         *bool
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Todo, error)
	predicates    []predicate.Todo
}

var _ ent.Mutation = (*TodoMutation)(nil)

// todoOption allows management of the mutation configuration using functional options.
type todoOption func(*TodoMutation)

// newTodoMutation creates new mutation for the Todo entity.
func newTodoMutation(c config, op Op, opts ...todoOption) *TodoMutation {
	m := &TodoMutation{
		config:        c,
		op:            op,
		typ:           TypeTodo,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTodoID sets the ID field of the mutation.
func withTodoID(id int) todoOption {
	return func(m *TodoMutation) {
		var (
			err   error
			once  sync.Once
			value *Todo
		)
		m.oldValue = func(ctx context.Context) (*Todo, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Todo.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTodo sets the old Todo of the mutation.
func withTodo(node *Todo) todoOption {
	return func(m *TodoMutation) {
		m.oldValue = func(context.Context) (*Todo, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TodoMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TodoMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *TodoMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *TodoMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Todo.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetTitle sets the "title" field.
func (m *TodoMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the value of the "title" field in the mutation.
func (m *TodoMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// OldTitle returns the old "title" field's value of the Todo entity.
// If the Todo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TodoMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTitle is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitle: %w", err)
	}
	return oldValue.Title, nil
}

// ResetTitle resets all changes to the "title" field.
func (m *TodoMutation) ResetTitle() {
	m.title = nil
}

// SetDone sets the "done" field.
func (m *TodoMutation) SetDone(b bool) {
	m._done = &b
}

// Done returns the value of the "done" field in the mutation.
func (m *TodoMutation) Done() (r bool, exists bool) {
	v := m._done
	if v == nil {
		return
	}
	return *v, true
}

// OldDone returns the old "done" field's value of the Todo entity.
// If the Todo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TodoMutation) OldDone(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDone is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDone requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDone: %w", err)
	}
	return oldValue.Done, nil
}

// ClearDone clears the value of the "done" field.
func (m *TodoMutation) ClearDone() {
	m._done = nil
	m.clearedFields[todo.FieldDone] = struct{}{}
}

// DoneCleared returns if the "done" field was cleared in this mutation.
func (m *TodoMutation) DoneCleared() bool {
	_, ok := m.clearedFields[todo.FieldDone]
	return ok
}

// ResetDone resets all changes to the "done" field.
func (m *TodoMutation) ResetDone() {
	m._done = nil
	delete(m.clearedFields, todo.FieldDone)
}

// Where appends a list predicates to the TodoMutation builder.
func (m *TodoMutation) Where(ps ...predicate.Todo) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *TodoMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Todo).
func (m *TodoMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TodoMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.title != nil {
		fields = append(fields, todo.FieldTitle)
	}
	if m._done != nil {
		fields = append(fields, todo.FieldDone)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TodoMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case todo.FieldTitle:
		return m.Title()
	case todo.FieldDone:
		return m.Done()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TodoMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case todo.FieldTitle:
		return m.OldTitle(ctx)
	case todo.FieldDone:
		return m.OldDone(ctx)
	}
	return nil, fmt.Errorf("unknown Todo field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TodoMutation) SetField(name string, value ent.Value) error {
	switch name {
	case todo.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case todo.FieldDone:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDone(v)
		return nil
	}
	return fmt.Errorf("unknown Todo field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TodoMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TodoMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TodoMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Todo numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TodoMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(todo.FieldDone) {
		fields = append(fields, todo.FieldDone)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TodoMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TodoMutation) ClearField(name string) error {
	switch name {
	case todo.FieldDone:
		m.ClearDone()
		return nil
	}
	return fmt.Errorf("unknown Todo nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TodoMutation) ResetField(name string) error {
	switch name {
	case todo.FieldTitle:
		m.ResetTitle()
		return nil
	case todo.FieldDone:
		m.ResetDone()
		return nil
	}
	return fmt.Errorf("unknown Todo field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TodoMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TodoMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TodoMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TodoMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TodoMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TodoMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TodoMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Todo unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TodoMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Todo edge %s", name)
}

// UsersMutation represents an operation that mutates the Users nodes in the graph.
type UsersMutation struct {
	config
	op            Op
	typ           string
	id            *int
	user          *string
	first         *int
	addfirst      *int
	draw          *int
	adddraw       *int
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Users, error)
	predicates    []predicate.Users
}

var _ ent.Mutation = (*UsersMutation)(nil)

// usersOption allows management of the mutation configuration using functional options.
type usersOption func(*UsersMutation)

// newUsersMutation creates new mutation for the Users entity.
func newUsersMutation(c config, op Op, opts ...usersOption) *UsersMutation {
	m := &UsersMutation{
		config:        c,
		op:            op,
		typ:           TypeUsers,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUsersID sets the ID field of the mutation.
func withUsersID(id int) usersOption {
	return func(m *UsersMutation) {
		var (
			err   error
			once  sync.Once
			value *Users
		)
		m.oldValue = func(ctx context.Context) (*Users, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Users.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUsers sets the old Users of the mutation.
func withUsers(node *Users) usersOption {
	return func(m *UsersMutation) {
		m.oldValue = func(context.Context) (*Users, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UsersMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UsersMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UsersMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UsersMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Users.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetUser sets the "user" field.
func (m *UsersMutation) SetUser(s string) {
	m.user = &s
}

// User returns the value of the "user" field in the mutation.
func (m *UsersMutation) User() (r string, exists bool) {
	v := m.user
	if v == nil {
		return
	}
	return *v, true
}

// OldUser returns the old "user" field's value of the Users entity.
// If the Users object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UsersMutation) OldUser(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUser is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUser requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUser: %w", err)
	}
	return oldValue.User, nil
}

// ResetUser resets all changes to the "user" field.
func (m *UsersMutation) ResetUser() {
	m.user = nil
}

// SetFirst sets the "first" field.
func (m *UsersMutation) SetFirst(i int) {
	m.first = &i
	m.addfirst = nil
}

// First returns the value of the "first" field in the mutation.
func (m *UsersMutation) First() (r int, exists bool) {
	v := m.first
	if v == nil {
		return
	}
	return *v, true
}

// OldFirst returns the old "first" field's value of the Users entity.
// If the Users object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UsersMutation) OldFirst(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFirst is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFirst requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFirst: %w", err)
	}
	return oldValue.First, nil
}

// AddFirst adds i to the "first" field.
func (m *UsersMutation) AddFirst(i int) {
	if m.addfirst != nil {
		*m.addfirst += i
	} else {
		m.addfirst = &i
	}
}

// AddedFirst returns the value that was added to the "first" field in this mutation.
func (m *UsersMutation) AddedFirst() (r int, exists bool) {
	v := m.addfirst
	if v == nil {
		return
	}
	return *v, true
}

// ClearFirst clears the value of the "first" field.
func (m *UsersMutation) ClearFirst() {
	m.first = nil
	m.addfirst = nil
	m.clearedFields[users.FieldFirst] = struct{}{}
}

// FirstCleared returns if the "first" field was cleared in this mutation.
func (m *UsersMutation) FirstCleared() bool {
	_, ok := m.clearedFields[users.FieldFirst]
	return ok
}

// ResetFirst resets all changes to the "first" field.
func (m *UsersMutation) ResetFirst() {
	m.first = nil
	m.addfirst = nil
	delete(m.clearedFields, users.FieldFirst)
}

// SetDraw sets the "draw" field.
func (m *UsersMutation) SetDraw(i int) {
	m.draw = &i
	m.adddraw = nil
}

// Draw returns the value of the "draw" field in the mutation.
func (m *UsersMutation) Draw() (r int, exists bool) {
	v := m.draw
	if v == nil {
		return
	}
	return *v, true
}

// OldDraw returns the old "draw" field's value of the Users entity.
// If the Users object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UsersMutation) OldDraw(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDraw is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDraw requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDraw: %w", err)
	}
	return oldValue.Draw, nil
}

// AddDraw adds i to the "draw" field.
func (m *UsersMutation) AddDraw(i int) {
	if m.adddraw != nil {
		*m.adddraw += i
	} else {
		m.adddraw = &i
	}
}

// AddedDraw returns the value that was added to the "draw" field in this mutation.
func (m *UsersMutation) AddedDraw() (r int, exists bool) {
	v := m.adddraw
	if v == nil {
		return
	}
	return *v, true
}

// ClearDraw clears the value of the "draw" field.
func (m *UsersMutation) ClearDraw() {
	m.draw = nil
	m.adddraw = nil
	m.clearedFields[users.FieldDraw] = struct{}{}
}

// DrawCleared returns if the "draw" field was cleared in this mutation.
func (m *UsersMutation) DrawCleared() bool {
	_, ok := m.clearedFields[users.FieldDraw]
	return ok
}

// ResetDraw resets all changes to the "draw" field.
func (m *UsersMutation) ResetDraw() {
	m.draw = nil
	m.adddraw = nil
	delete(m.clearedFields, users.FieldDraw)
}

// SetCreatedAt sets the "created_at" field.
func (m *UsersMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *UsersMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Users entity.
// If the Users object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UsersMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ClearCreatedAt clears the value of the "created_at" field.
func (m *UsersMutation) ClearCreatedAt() {
	m.created_at = nil
	m.clearedFields[users.FieldCreatedAt] = struct{}{}
}

// CreatedAtCleared returns if the "created_at" field was cleared in this mutation.
func (m *UsersMutation) CreatedAtCleared() bool {
	_, ok := m.clearedFields[users.FieldCreatedAt]
	return ok
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *UsersMutation) ResetCreatedAt() {
	m.created_at = nil
	delete(m.clearedFields, users.FieldCreatedAt)
}

// SetUpdatedAt sets the "updated_at" field.
func (m *UsersMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *UsersMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Users entity.
// If the Users object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UsersMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (m *UsersMutation) ClearUpdatedAt() {
	m.updated_at = nil
	m.clearedFields[users.FieldUpdatedAt] = struct{}{}
}

// UpdatedAtCleared returns if the "updated_at" field was cleared in this mutation.
func (m *UsersMutation) UpdatedAtCleared() bool {
	_, ok := m.clearedFields[users.FieldUpdatedAt]
	return ok
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *UsersMutation) ResetUpdatedAt() {
	m.updated_at = nil
	delete(m.clearedFields, users.FieldUpdatedAt)
}

// Where appends a list predicates to the UsersMutation builder.
func (m *UsersMutation) Where(ps ...predicate.Users) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *UsersMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Users).
func (m *UsersMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UsersMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.user != nil {
		fields = append(fields, users.FieldUser)
	}
	if m.first != nil {
		fields = append(fields, users.FieldFirst)
	}
	if m.draw != nil {
		fields = append(fields, users.FieldDraw)
	}
	if m.created_at != nil {
		fields = append(fields, users.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, users.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UsersMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case users.FieldUser:
		return m.User()
	case users.FieldFirst:
		return m.First()
	case users.FieldDraw:
		return m.Draw()
	case users.FieldCreatedAt:
		return m.CreatedAt()
	case users.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UsersMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case users.FieldUser:
		return m.OldUser(ctx)
	case users.FieldFirst:
		return m.OldFirst(ctx)
	case users.FieldDraw:
		return m.OldDraw(ctx)
	case users.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case users.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Users field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UsersMutation) SetField(name string, value ent.Value) error {
	switch name {
	case users.FieldUser:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUser(v)
		return nil
	case users.FieldFirst:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFirst(v)
		return nil
	case users.FieldDraw:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDraw(v)
		return nil
	case users.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case users.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Users field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UsersMutation) AddedFields() []string {
	var fields []string
	if m.addfirst != nil {
		fields = append(fields, users.FieldFirst)
	}
	if m.adddraw != nil {
		fields = append(fields, users.FieldDraw)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UsersMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case users.FieldFirst:
		return m.AddedFirst()
	case users.FieldDraw:
		return m.AddedDraw()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UsersMutation) AddField(name string, value ent.Value) error {
	switch name {
	case users.FieldFirst:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddFirst(v)
		return nil
	case users.FieldDraw:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddDraw(v)
		return nil
	}
	return fmt.Errorf("unknown Users numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UsersMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(users.FieldFirst) {
		fields = append(fields, users.FieldFirst)
	}
	if m.FieldCleared(users.FieldDraw) {
		fields = append(fields, users.FieldDraw)
	}
	if m.FieldCleared(users.FieldCreatedAt) {
		fields = append(fields, users.FieldCreatedAt)
	}
	if m.FieldCleared(users.FieldUpdatedAt) {
		fields = append(fields, users.FieldUpdatedAt)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UsersMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UsersMutation) ClearField(name string) error {
	switch name {
	case users.FieldFirst:
		m.ClearFirst()
		return nil
	case users.FieldDraw:
		m.ClearDraw()
		return nil
	case users.FieldCreatedAt:
		m.ClearCreatedAt()
		return nil
	case users.FieldUpdatedAt:
		m.ClearUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Users nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UsersMutation) ResetField(name string) error {
	switch name {
	case users.FieldUser:
		m.ResetUser()
		return nil
	case users.FieldFirst:
		m.ResetFirst()
		return nil
	case users.FieldDraw:
		m.ResetDraw()
		return nil
	case users.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case users.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Users field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UsersMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UsersMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UsersMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UsersMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UsersMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UsersMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UsersMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Users unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UsersMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Users edge %s", name)
}
