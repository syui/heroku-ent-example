// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/bits"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = bits.LeadingZeros64
	_ = big.Rat{}
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = attribute.KeyValue{}
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
	_ = codes.Unset
)

type CreateTodoReq struct {
	Title string  "json:\"title\""
	Done  OptBool "json:\"done\""
}

type CreateUsersReq struct {
	User      string      "json:\"user\""
	First     OptInt      "json:\"first\""
	Start     OptBool     "json:\"start\""
	Draw      OptInt      "json:\"draw\""
	CreatedAt OptDateTime "json:\"created_at\""
	UpdatedAt OptDateTime "json:\"updated_at\""
}

// DeleteTodoNoContent is response for DeleteTodo operation.
type DeleteTodoNoContent struct{}

func (*DeleteTodoNoContent) deleteTodoRes() {}

// DeleteUsersNoContent is response for DeleteUsers operation.
type DeleteUsersNoContent struct{}

func (*DeleteUsersNoContent) deleteUsersRes() {}

// DrawDoneNoContent is response for DrawDone operation.
type DrawDoneNoContent struct{}

// DrawStartNoContent is response for DrawStart operation.
type DrawStartNoContent struct{}

type ListTodoOKApplicationJSON []TodoList

func (ListTodoOKApplicationJSON) listTodoRes() {}

type ListUsersOKApplicationJSON []UsersList

func (ListUsersOKApplicationJSON) listUsersRes() {}

// MarkDoneNoContent is response for MarkDone operation.
type MarkDoneNoContent struct{}

// NewOptBool returns new OptBool with value set to v.
func NewOptBool(v bool) OptBool {
	return OptBool{
		Value: v,
		Set:   true,
	}
}

// OptBool is optional bool.
type OptBool struct {
	Value bool
	Set   bool
}

// IsSet returns true if OptBool was set.
func (o OptBool) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptBool) Reset() {
	var v bool
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptBool) SetTo(v bool) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptBool) Get() (v bool, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptBool) Or(d bool) bool {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptDateTime returns new OptDateTime with value set to v.
func NewOptDateTime(v time.Time) OptDateTime {
	return OptDateTime{
		Value: v,
		Set:   true,
	}
}

// OptDateTime is optional time.Time.
type OptDateTime struct {
	Value time.Time
	Set   bool
}

// IsSet returns true if OptDateTime was set.
func (o OptDateTime) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptDateTime) Reset() {
	var v time.Time
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptDateTime) SetTo(v time.Time) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptDateTime) Get() (v time.Time, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptDateTime) Or(d time.Time) time.Time {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

type R400 struct {
	Code   int    "json:\"code\""
	Status string "json:\"status\""
	Errors jx.Raw "json:\"errors\""
}

func (*R400) createTodoRes()  {}
func (*R400) createUsersRes() {}
func (*R400) deleteTodoRes()  {}
func (*R400) deleteUsersRes() {}
func (*R400) listTodoRes()    {}
func (*R400) listUsersRes()   {}
func (*R400) readTodoRes()    {}
func (*R400) readUsersRes()   {}
func (*R400) updateTodoRes()  {}
func (*R400) updateUsersRes() {}

type R404 struct {
	Code   int    "json:\"code\""
	Status string "json:\"status\""
	Errors jx.Raw "json:\"errors\""
}

func (*R404) deleteTodoRes()  {}
func (*R404) deleteUsersRes() {}
func (*R404) listTodoRes()    {}
func (*R404) listUsersRes()   {}
func (*R404) readTodoRes()    {}
func (*R404) readUsersRes()   {}
func (*R404) updateTodoRes()  {}
func (*R404) updateUsersRes() {}

type R409 struct {
	Code   int    "json:\"code\""
	Status string "json:\"status\""
	Errors jx.Raw "json:\"errors\""
}

func (*R409) createTodoRes()  {}
func (*R409) createUsersRes() {}
func (*R409) deleteTodoRes()  {}
func (*R409) deleteUsersRes() {}
func (*R409) listTodoRes()    {}
func (*R409) listUsersRes()   {}
func (*R409) readTodoRes()    {}
func (*R409) readUsersRes()   {}
func (*R409) updateTodoRes()  {}
func (*R409) updateUsersRes() {}

type R500 struct {
	Code   int    "json:\"code\""
	Status string "json:\"status\""
	Errors jx.Raw "json:\"errors\""
}

func (*R500) createTodoRes()  {}
func (*R500) createUsersRes() {}
func (*R500) deleteTodoRes()  {}
func (*R500) deleteUsersRes() {}
func (*R500) listTodoRes()    {}
func (*R500) listUsersRes()   {}
func (*R500) readTodoRes()    {}
func (*R500) readUsersRes()   {}
func (*R500) updateTodoRes()  {}
func (*R500) updateUsersRes() {}

// Ref: #/components/schemas/TodoCreate
type TodoCreate struct {
	ID    int     "json:\"id\""
	Title string  "json:\"title\""
	Done  OptBool "json:\"done\""
}

func (*TodoCreate) createTodoRes() {}

// Ref: #/components/schemas/TodoList
type TodoList struct {
	ID    int     "json:\"id\""
	Title string  "json:\"title\""
	Done  OptBool "json:\"done\""
}

// Ref: #/components/schemas/TodoRead
type TodoRead struct {
	ID    int     "json:\"id\""
	Title string  "json:\"title\""
	Done  OptBool "json:\"done\""
}

func (*TodoRead) readTodoRes() {}

// Ref: #/components/schemas/TodoUpdate
type TodoUpdate struct {
	ID    int     "json:\"id\""
	Title string  "json:\"title\""
	Done  OptBool "json:\"done\""
}

func (*TodoUpdate) updateTodoRes() {}

type UpdateTodoReq struct {
	Title OptString "json:\"title\""
	Done  OptBool   "json:\"done\""
}

type UpdateUsersReq struct {
	Start     OptBool     "json:\"start\""
	Draw      OptInt      "json:\"draw\""
	CreatedAt OptDateTime "json:\"created_at\""
	UpdatedAt OptDateTime "json:\"updated_at\""
}

// Ref: #/components/schemas/UsersCreate
type UsersCreate struct {
	ID        int         "json:\"id\""
	User      string      "json:\"user\""
	First     OptInt      "json:\"first\""
	Start     OptBool     "json:\"start\""
	Draw      OptInt      "json:\"draw\""
	CreatedAt OptDateTime "json:\"created_at\""
	UpdatedAt OptDateTime "json:\"updated_at\""
}

func (*UsersCreate) createUsersRes() {}

// Ref: #/components/schemas/UsersList
type UsersList struct {
	ID        int         "json:\"id\""
	User      string      "json:\"user\""
	First     OptInt      "json:\"first\""
	Start     OptBool     "json:\"start\""
	Draw      OptInt      "json:\"draw\""
	CreatedAt OptDateTime "json:\"created_at\""
	UpdatedAt OptDateTime "json:\"updated_at\""
}

// Ref: #/components/schemas/UsersRead
type UsersRead struct {
	ID        int         "json:\"id\""
	User      string      "json:\"user\""
	First     OptInt      "json:\"first\""
	Start     OptBool     "json:\"start\""
	Draw      OptInt      "json:\"draw\""
	CreatedAt OptDateTime "json:\"created_at\""
	UpdatedAt OptDateTime "json:\"updated_at\""
}

func (*UsersRead) readUsersRes() {}

// Ref: #/components/schemas/UsersUpdate
type UsersUpdate struct {
	ID        int         "json:\"id\""
	User      string      "json:\"user\""
	First     OptInt      "json:\"first\""
	Start     OptBool     "json:\"start\""
	Draw      OptInt      "json:\"draw\""
	CreatedAt OptDateTime "json:\"created_at\""
	UpdatedAt OptDateTime "json:\"updated_at\""
}

func (*UsersUpdate) updateUsersRes() {}
