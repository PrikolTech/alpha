// Code generated by ogen, DO NOT EDIT.

package api

import (
	"time"

	"github.com/google/uuid"
)

// Доменная ошибка.
// Ref: #/components/schemas/DomainError
type DomainError struct {
	Message string `json:"message"`
}

// GetMessage returns the value of Message.
func (s *DomainError) GetMessage() string {
	return s.Message
}

// SetMessage sets the value of Message.
func (s *DomainError) SetMessage(val string) {
	s.Message = val
}

func (*DomainError) userCreateRes()  {}
func (*DomainError) userGetByIdRes() {}

// Мета данные.
// Ref: #/components/schemas/Meta
type Meta struct {
	// Номер страницы.
	Page int `json:"page"`
	// Общее количество страниц.
	TotalPages int `json:"totalPages"`
	// Количество записей на странице.
	PerPage OptInt `json:"perPage"`
	// Общее количество записей.
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

// GetPerPage returns the value of PerPage.
func (s *Meta) GetPerPage() OptInt {
	return s.PerPage
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

// SetPerPage sets the value of PerPage.
func (s *Meta) SetPerPage(val OptInt) {
	s.PerPage = val
}

// SetTotalRecords sets the value of TotalRecords.
func (s *Meta) SetTotalRecords(val int) {
	s.TotalRecords = val
}

// NewNilString returns new NilString with value set to v.
func NewNilString(v string) NilString {
	return NilString{
		Value: v,
	}
}

// NilString is nullable string.
type NilString struct {
	Value string
	Null  bool
}

// SetTo sets value to v.
func (o *NilString) SetTo(v string) {
	o.Null = false
	o.Value = v
}

// IsSet returns true if value is Null.
func (o NilString) IsNull() bool { return o.Null }

// SetNull sets value to null.
func (o *NilString) SetToNull() {
	o.Null = true
	var v string
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o NilString) Get() (v string, ok bool) {
	if o.Null {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o NilString) Or(d string) string {
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

// NewOptMeta returns new OptMeta with value set to v.
func NewOptMeta(v Meta) OptMeta {
	return OptMeta{
		Value: v,
		Set:   true,
	}
}

// OptMeta is optional Meta.
type OptMeta struct {
	Value Meta
	Set   bool
}

// IsSet returns true if OptMeta was set.
func (o OptMeta) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptMeta) Reset() {
	var v Meta
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptMeta) SetTo(v Meta) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptMeta) Get() (v Meta, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptMeta) Or(d Meta) Meta {
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

// Проект.
// Ref: #/components/schemas/Project
type Project struct {
	Name        string      `json:"name"`
	Description NilString   `json:"description"`
	Code        string      `json:"code"`
	CreatedAt   OptDateTime `json:"createdAt"`
	UpdatedAt   OptDateTime `json:"updatedAt"`
}

// GetName returns the value of Name.
func (s *Project) GetName() string {
	return s.Name
}

// GetDescription returns the value of Description.
func (s *Project) GetDescription() NilString {
	return s.Description
}

// GetCode returns the value of Code.
func (s *Project) GetCode() string {
	return s.Code
}

// GetCreatedAt returns the value of CreatedAt.
func (s *Project) GetCreatedAt() OptDateTime {
	return s.CreatedAt
}

// GetUpdatedAt returns the value of UpdatedAt.
func (s *Project) GetUpdatedAt() OptDateTime {
	return s.UpdatedAt
}

// SetName sets the value of Name.
func (s *Project) SetName(val string) {
	s.Name = val
}

// SetDescription sets the value of Description.
func (s *Project) SetDescription(val NilString) {
	s.Description = val
}

// SetCode sets the value of Code.
func (s *Project) SetCode(val string) {
	s.Code = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *Project) SetCreatedAt(val OptDateTime) {
	s.CreatedAt = val
}

// SetUpdatedAt sets the value of UpdatedAt.
func (s *Project) SetUpdatedAt(val OptDateTime) {
	s.UpdatedAt = val
}

// ProjectCreateCreated is response for ProjectCreate operation.
type ProjectCreateCreated struct{}

// Запрос на создание проекта.
// Ref: #/components/schemas/ProjectCreateRequest
type ProjectCreateRequest struct {
	Name        string    `json:"name"`
	Description NilString `json:"description"`
	Code        string    `json:"code"`
}

// GetName returns the value of Name.
func (s *ProjectCreateRequest) GetName() string {
	return s.Name
}

// GetDescription returns the value of Description.
func (s *ProjectCreateRequest) GetDescription() NilString {
	return s.Description
}

// GetCode returns the value of Code.
func (s *ProjectCreateRequest) GetCode() string {
	return s.Code
}

// SetName sets the value of Name.
func (s *ProjectCreateRequest) SetName(val string) {
	s.Name = val
}

// SetDescription sets the value of Description.
func (s *ProjectCreateRequest) SetDescription(val NilString) {
	s.Description = val
}

// SetCode sets the value of Code.
func (s *ProjectCreateRequest) SetCode(val string) {
	s.Code = val
}

// ProjectDeleteByIdNoContent is response for ProjectDeleteById operation.
type ProjectDeleteByIdNoContent struct{}

// Список проектов.
// Ref: #/components/schemas/ProjectGetAllResponse
type ProjectGetAllResponse struct {
	Data []Project `json:"data"`
	Meta OptMeta   `json:"meta"`
}

// GetData returns the value of Data.
func (s *ProjectGetAllResponse) GetData() []Project {
	return s.Data
}

// GetMeta returns the value of Meta.
func (s *ProjectGetAllResponse) GetMeta() OptMeta {
	return s.Meta
}

// SetData sets the value of Data.
func (s *ProjectGetAllResponse) SetData(val []Project) {
	s.Data = val
}

// SetMeta sets the value of Meta.
func (s *ProjectGetAllResponse) SetMeta(val OptMeta) {
	s.Meta = val
}

// Пользователь.
// Ref: #/components/schemas/User
type User struct {
	ID         uuid.UUID    `json:"id"`
	Email      string       `json:"email"`
	FirstName  string       `json:"firstName"`
	MiddleName OptNilString `json:"middleName"`
	LastName   string       `json:"lastName"`
	CreatedAt  time.Time    `json:"createdAt"`
	UpdatedAt  time.Time    `json:"updatedAt"`
}

// GetID returns the value of ID.
func (s *User) GetID() uuid.UUID {
	return s.ID
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
func (s *User) GetCreatedAt() time.Time {
	return s.CreatedAt
}

// GetUpdatedAt returns the value of UpdatedAt.
func (s *User) GetUpdatedAt() time.Time {
	return s.UpdatedAt
}

// SetID sets the value of ID.
func (s *User) SetID(val uuid.UUID) {
	s.ID = val
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
func (s *User) SetCreatedAt(val time.Time) {
	s.CreatedAt = val
}

// SetUpdatedAt sets the value of UpdatedAt.
func (s *User) SetUpdatedAt(val time.Time) {
	s.UpdatedAt = val
}

func (*User) userGetByIdRes() {}

// UserCreateCreated is response for UserCreate operation.
type UserCreateCreated struct{}

func (*UserCreateCreated) userCreateRes() {}

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
// Ref: #/components/schemas/UserListResponse
type UserListResponse struct {
	Data []User `json:"data"`
	Meta Meta   `json:"meta"`
}

// GetData returns the value of Data.
func (s *UserListResponse) GetData() []User {
	return s.Data
}

// GetMeta returns the value of Meta.
func (s *UserListResponse) GetMeta() Meta {
	return s.Meta
}

// SetData sets the value of Data.
func (s *UserListResponse) SetData(val []User) {
	s.Data = val
}

// SetMeta sets the value of Meta.
func (s *UserListResponse) SetMeta(val Meta) {
	s.Meta = val
}
