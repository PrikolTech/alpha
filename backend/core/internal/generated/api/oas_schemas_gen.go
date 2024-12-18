// Code generated by ogen, DO NOT EDIT.

package api

import (
	"time"
)

// Мета данные.
// Ref: #/components/schemas/Meta
type Meta struct {
	Page         int `json:"page"`
	TotalPages   int `json:"totalPages"`
	Per          int `json:"per"`
	TotalRecords int `json:"totalRecords"`
}

// GetPage returns the value of Page.
func (s *Meta) GetPage() int {
	return s.Page
}

// GetTotalPages returns the value of TotalPages.
func (s *Meta) GetTotalPages() int {
	return s.TotalPages
}

// GetPer returns the value of Per.
func (s *Meta) GetPer() int {
	return s.Per
}

// GetTotalRecords returns the value of TotalRecords.
func (s *Meta) GetTotalRecords() int {
	return s.TotalRecords
}

// SetPage sets the value of Page.
func (s *Meta) SetPage(val int) {
	s.Page = val
}

// SetTotalPages sets the value of TotalPages.
func (s *Meta) SetTotalPages(val int) {
	s.TotalPages = val
}

// SetPer sets the value of Per.
func (s *Meta) SetPer(val int) {
	s.Per = val
}

// SetTotalRecords sets the value of TotalRecords.
func (s *Meta) SetTotalRecords(val int) {
	s.TotalRecords = val
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

// NewOptNilString returns new OptNilString with value set to v.
func NewOptNilString(v string) OptNilString {
	return OptNilString{
		Value: v,
		Set:   true,
	}
}

// OptNilString is optional nullable string.
type OptNilString struct {
	Value string
	Set   bool
	Null  bool
}

// IsSet returns true if OptNilString was set.
func (o OptNilString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNilString) Reset() {
	var v string
	o.Value = v
	o.Set = false
	o.Null = false
}

// SetTo sets value to v.
func (o *OptNilString) SetTo(v string) {
	o.Set = true
	o.Null = false
	o.Value = v
}

// IsSet returns true if value is Null.
func (o OptNilString) IsNull() bool { return o.Null }

// SetNull sets value to null.
func (o *OptNilString) SetToNull() {
	o.Set = true
	o.Null = true
	var v string
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNilString) Get() (v string, ok bool) {
	if o.Null {
		return v, false
	}
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNilString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Пользователь.
// Ref: #/components/schemas/User
type User struct {
	Email      string       `json:"email"`
	FirstName  string       `json:"firstName"`
	MiddleName OptNilString `json:"middleName"`
	LastName   string       `json:"lastName"`
	CreatedAt  OptDateTime  `json:"createdAt"`
	UpdatedAt  OptDateTime  `json:"updatedAt"`
}

// GetEmail returns the value of Email.
func (s *User) GetEmail() string {
	return s.Email
}

// GetFirstName returns the value of FirstName.
func (s *User) GetFirstName() string {
	return s.FirstName
}

// GetMiddleName returns the value of MiddleName.
func (s *User) GetMiddleName() OptNilString {
	return s.MiddleName
}

// GetLastName returns the value of LastName.
func (s *User) GetLastName() string {
	return s.LastName
}

// GetCreatedAt returns the value of CreatedAt.
func (s *User) GetCreatedAt() OptDateTime {
	return s.CreatedAt
}

// GetUpdatedAt returns the value of UpdatedAt.
func (s *User) GetUpdatedAt() OptDateTime {
	return s.UpdatedAt
}

// SetEmail sets the value of Email.
func (s *User) SetEmail(val string) {
	s.Email = val
}

// SetFirstName sets the value of FirstName.
func (s *User) SetFirstName(val string) {
	s.FirstName = val
}

// SetMiddleName sets the value of MiddleName.
func (s *User) SetMiddleName(val OptNilString) {
	s.MiddleName = val
}

// SetLastName sets the value of LastName.
func (s *User) SetLastName(val string) {
	s.LastName = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *User) SetCreatedAt(val OptDateTime) {
	s.CreatedAt = val
}

// SetUpdatedAt sets the value of UpdatedAt.
func (s *User) SetUpdatedAt(val OptDateTime) {
	s.UpdatedAt = val
}

// UserCreateCreated is response for UserCreate operation.
type UserCreateCreated struct{}

func (*UserCreateCreated) userCreateRes() {}

// Ref: #/components/schemas/UserCreateDomainError
type UserCreateDomainError struct {
	Message string `json:"message"`
}

// GetMessage returns the value of Message.
func (s *UserCreateDomainError) GetMessage() string {
	return s.Message
}

// SetMessage sets the value of Message.
func (s *UserCreateDomainError) SetMessage(val string) {
	s.Message = val
}

func (*UserCreateDomainError) userCreateRes() {}

// Ref: #/components/schemas/UserCreateRequest
type UserCreateRequest struct {
	Email      string       `json:"email"`
	FirstName  string       `json:"firstName"`
	MiddleName OptNilString `json:"middleName"`
	LastName   string       `json:"lastName"`
}

// GetEmail returns the value of Email.
func (s *UserCreateRequest) GetEmail() string {
	return s.Email
}

// GetFirstName returns the value of FirstName.
func (s *UserCreateRequest) GetFirstName() string {
	return s.FirstName
}

// GetMiddleName returns the value of MiddleName.
func (s *UserCreateRequest) GetMiddleName() OptNilString {
	return s.MiddleName
}

// GetLastName returns the value of LastName.
func (s *UserCreateRequest) GetLastName() string {
	return s.LastName
}

// SetEmail sets the value of Email.
func (s *UserCreateRequest) SetEmail(val string) {
	s.Email = val
}

// SetFirstName sets the value of FirstName.
func (s *UserCreateRequest) SetFirstName(val string) {
	s.FirstName = val
}

// SetMiddleName sets the value of MiddleName.
func (s *UserCreateRequest) SetMiddleName(val OptNilString) {
	s.MiddleName = val
}

// SetLastName sets the value of LastName.
func (s *UserCreateRequest) SetLastName(val string) {
	s.LastName = val
}

// Ref: #/components/schemas/UserCreateValidationError
type UserCreateValidationError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

// GetField returns the value of Field.
func (s *UserCreateValidationError) GetField() string {
	return s.Field
}

// GetReason returns the value of Reason.
func (s *UserCreateValidationError) GetReason() string {
	return s.Reason
}

// SetField sets the value of Field.
func (s *UserCreateValidationError) SetField(val string) {
	s.Field = val
}

// SetReason sets the value of Reason.
func (s *UserCreateValidationError) SetReason(val string) {
	s.Reason = val
}

func (*UserCreateValidationError) userCreateRes() {}

// Список пользователей.
// Ref: #/components/schemas/UserGetAllResponse
type UserGetAllResponse struct {
	Data []User `json:"data"`
	Meta Meta   `json:"meta"`
}

// GetData returns the value of Data.
func (s *UserGetAllResponse) GetData() []User {
	return s.Data
}

// GetMeta returns the value of Meta.
func (s *UserGetAllResponse) GetMeta() Meta {
	return s.Meta
}

// SetData sets the value of Data.
func (s *UserGetAllResponse) SetData(val []User) {
	s.Data = val
}

// SetMeta sets the value of Meta.
func (s *UserGetAllResponse) SetMeta(val Meta) {
	s.Meta = val
}
