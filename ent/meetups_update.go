// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/confus1on/meetmeup/ent/meetups"
	"github.com/confus1on/meetmeup/ent/predicate"
	"github.com/confus1on/meetmeup/ent/users"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// MeetupsUpdate is the builder for updating Meetups entities.
type MeetupsUpdate struct {
	config
	name        *string
	description *string
	usersID     *int
	addusersID  *int
	user        map[int]struct{}
	removedUser map[int]struct{}
	predicates  []predicate.Meetups
}

// Where adds a new predicate for the builder.
func (mu *MeetupsUpdate) Where(ps ...predicate.Meetups) *MeetupsUpdate {
	mu.predicates = append(mu.predicates, ps...)
	return mu
}

// SetName sets the name field.
func (mu *MeetupsUpdate) SetName(s string) *MeetupsUpdate {
	mu.name = &s
	return mu
}

// SetDescription sets the description field.
func (mu *MeetupsUpdate) SetDescription(s string) *MeetupsUpdate {
	mu.description = &s
	return mu
}

// SetUsersID sets the usersID field.
func (mu *MeetupsUpdate) SetUsersID(i int) *MeetupsUpdate {
	mu.usersID = &i
	mu.addusersID = nil
	return mu
}

// AddUsersID adds i to usersID.
func (mu *MeetupsUpdate) AddUsersID(i int) *MeetupsUpdate {
	if mu.addusersID == nil {
		mu.addusersID = &i
	} else {
		*mu.addusersID += i
	}
	return mu
}

// AddUserIDs adds the user edge to Users by ids.
func (mu *MeetupsUpdate) AddUserIDs(ids ...int) *MeetupsUpdate {
	if mu.user == nil {
		mu.user = make(map[int]struct{})
	}
	for i := range ids {
		mu.user[ids[i]] = struct{}{}
	}
	return mu
}

// AddUser adds the user edges to Users.
func (mu *MeetupsUpdate) AddUser(u ...*Users) *MeetupsUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return mu.AddUserIDs(ids...)
}

// RemoveUserIDs removes the user edge to Users by ids.
func (mu *MeetupsUpdate) RemoveUserIDs(ids ...int) *MeetupsUpdate {
	if mu.removedUser == nil {
		mu.removedUser = make(map[int]struct{})
	}
	for i := range ids {
		mu.removedUser[ids[i]] = struct{}{}
	}
	return mu
}

// RemoveUser removes user edges to Users.
func (mu *MeetupsUpdate) RemoveUser(u ...*Users) *MeetupsUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return mu.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (mu *MeetupsUpdate) Save(ctx context.Context) (int, error) {
	return mu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MeetupsUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MeetupsUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MeetupsUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *MeetupsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   meetups.Table,
			Columns: meetups.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: meetups.FieldID,
			},
		},
	}
	if ps := mu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := mu.name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: meetups.FieldName,
		})
	}
	if value := mu.description; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: meetups.FieldDescription,
		})
	}
	if value := mu.usersID; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: meetups.FieldUsersID,
		})
	}
	if value := mu.addusersID; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: meetups.FieldUsersID,
		})
	}
	if nodes := mu.removedUser; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   meetups.UserTable,
			Columns: []string{meetups.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: users.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.user; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   meetups.UserTable,
			Columns: []string{meetups.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: users.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// MeetupsUpdateOne is the builder for updating a single Meetups entity.
type MeetupsUpdateOne struct {
	config
	id          int
	name        *string
	description *string
	usersID     *int
	addusersID  *int
	user        map[int]struct{}
	removedUser map[int]struct{}
}

// SetName sets the name field.
func (muo *MeetupsUpdateOne) SetName(s string) *MeetupsUpdateOne {
	muo.name = &s
	return muo
}

// SetDescription sets the description field.
func (muo *MeetupsUpdateOne) SetDescription(s string) *MeetupsUpdateOne {
	muo.description = &s
	return muo
}

// SetUsersID sets the usersID field.
func (muo *MeetupsUpdateOne) SetUsersID(i int) *MeetupsUpdateOne {
	muo.usersID = &i
	muo.addusersID = nil
	return muo
}

// AddUsersID adds i to usersID.
func (muo *MeetupsUpdateOne) AddUsersID(i int) *MeetupsUpdateOne {
	if muo.addusersID == nil {
		muo.addusersID = &i
	} else {
		*muo.addusersID += i
	}
	return muo
}

// AddUserIDs adds the user edge to Users by ids.
func (muo *MeetupsUpdateOne) AddUserIDs(ids ...int) *MeetupsUpdateOne {
	if muo.user == nil {
		muo.user = make(map[int]struct{})
	}
	for i := range ids {
		muo.user[ids[i]] = struct{}{}
	}
	return muo
}

// AddUser adds the user edges to Users.
func (muo *MeetupsUpdateOne) AddUser(u ...*Users) *MeetupsUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return muo.AddUserIDs(ids...)
}

// RemoveUserIDs removes the user edge to Users by ids.
func (muo *MeetupsUpdateOne) RemoveUserIDs(ids ...int) *MeetupsUpdateOne {
	if muo.removedUser == nil {
		muo.removedUser = make(map[int]struct{})
	}
	for i := range ids {
		muo.removedUser[ids[i]] = struct{}{}
	}
	return muo
}

// RemoveUser removes user edges to Users.
func (muo *MeetupsUpdateOne) RemoveUser(u ...*Users) *MeetupsUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return muo.RemoveUserIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (muo *MeetupsUpdateOne) Save(ctx context.Context) (*Meetups, error) {
	return muo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MeetupsUpdateOne) SaveX(ctx context.Context) *Meetups {
	m, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return m
}

// Exec executes the query on the entity.
func (muo *MeetupsUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MeetupsUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *MeetupsUpdateOne) sqlSave(ctx context.Context) (m *Meetups, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   meetups.Table,
			Columns: meetups.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  muo.id,
				Type:   field.TypeInt,
				Column: meetups.FieldID,
			},
		},
	}
	if value := muo.name; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: meetups.FieldName,
		})
	}
	if value := muo.description; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: meetups.FieldDescription,
		})
	}
	if value := muo.usersID; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: meetups.FieldUsersID,
		})
	}
	if value := muo.addusersID; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: meetups.FieldUsersID,
		})
	}
	if nodes := muo.removedUser; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   meetups.UserTable,
			Columns: []string{meetups.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: users.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.user; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   meetups.UserTable,
			Columns: []string{meetups.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: users.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	m = &Meetups{config: muo.config}
	_spec.Assign = m.assignValues
	_spec.ScanValues = m.scanValues()
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return m, nil
}