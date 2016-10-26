package chado

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/queries"
	"github.com/vattle/sqlboiler/queries/qm"
	"github.com/vattle/sqlboiler/strmangle"
	"gopkg.in/nullbio/null.v5"
)

// AuthUser is an object representing the database table.
type AuthUser struct {
	AuthUserID int       `boil:"auth_user_id" json:"auth_user_id" toml:"auth_user_id" yaml:"auth_user_id"`
	FirstName  string    `boil:"first_name" json:"first_name" toml:"first_name" yaml:"first_name"`
	LastName   string    `boil:"last_name" json:"last_name" toml:"last_name" yaml:"last_name"`
	Email      string    `boil:"email" json:"email" toml:"email" yaml:"email"`
	IsActive   bool      `boil:"is_active" json:"is_active" toml:"is_active" yaml:"is_active"`
	CreatedAt  null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt  null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *authUserR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L authUserL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// authUserR is where relationships are stored.
type authUserR struct {
	UserStockOrders StockOrderSlice
	AuthUserInfos   AuthUserInfoSlice
	AuthUserRoles   AuthUserRoleSlice
}

// authUserL is where Load methods for each relationship are stored.
type authUserL struct{}

var (
	authUserColumns               = []string{"auth_user_id", "first_name", "last_name", "email", "is_active", "created_at", "updated_at"}
	authUserColumnsWithoutDefault = []string{"first_name", "last_name", "email"}
	authUserColumnsWithDefault    = []string{"auth_user_id", "is_active", "created_at", "updated_at"}
	authUserPrimaryKeyColumns     = []string{"auth_user_id"}
)

type (
	// AuthUserSlice is an alias for a slice of pointers to AuthUser.
	// This should generally be used opposed to []AuthUser.
	AuthUserSlice []*AuthUser
	// AuthUserHook is the signature for custom AuthUser hook methods
	AuthUserHook func(boil.Executor, *AuthUser) error

	authUserQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	authUserType                 = reflect.TypeOf(&AuthUser{})
	authUserMapping              = queries.MakeStructMapping(authUserType)
	authUserPrimaryKeyMapping, _ = queries.BindMapping(authUserType, authUserMapping, authUserPrimaryKeyColumns)
	authUserInsertCacheMut       sync.RWMutex
	authUserInsertCache          = make(map[string]insertCache)
	authUserUpdateCacheMut       sync.RWMutex
	authUserUpdateCache          = make(map[string]updateCache)
	authUserUpsertCacheMut       sync.RWMutex
	authUserUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var authUserBeforeInsertHooks []AuthUserHook
var authUserBeforeUpdateHooks []AuthUserHook
var authUserBeforeDeleteHooks []AuthUserHook
var authUserBeforeUpsertHooks []AuthUserHook

var authUserAfterInsertHooks []AuthUserHook
var authUserAfterSelectHooks []AuthUserHook
var authUserAfterUpdateHooks []AuthUserHook
var authUserAfterDeleteHooks []AuthUserHook
var authUserAfterUpsertHooks []AuthUserHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AuthUser) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AuthUser) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AuthUser) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AuthUser) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AuthUser) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AuthUser) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AuthUser) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AuthUser) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AuthUser) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAuthUserHook registers your hook function for all future operations.
func AddAuthUserHook(hookPoint boil.HookPoint, authUserHook AuthUserHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		authUserBeforeInsertHooks = append(authUserBeforeInsertHooks, authUserHook)
	case boil.BeforeUpdateHook:
		authUserBeforeUpdateHooks = append(authUserBeforeUpdateHooks, authUserHook)
	case boil.BeforeDeleteHook:
		authUserBeforeDeleteHooks = append(authUserBeforeDeleteHooks, authUserHook)
	case boil.BeforeUpsertHook:
		authUserBeforeUpsertHooks = append(authUserBeforeUpsertHooks, authUserHook)
	case boil.AfterInsertHook:
		authUserAfterInsertHooks = append(authUserAfterInsertHooks, authUserHook)
	case boil.AfterSelectHook:
		authUserAfterSelectHooks = append(authUserAfterSelectHooks, authUserHook)
	case boil.AfterUpdateHook:
		authUserAfterUpdateHooks = append(authUserAfterUpdateHooks, authUserHook)
	case boil.AfterDeleteHook:
		authUserAfterDeleteHooks = append(authUserAfterDeleteHooks, authUserHook)
	case boil.AfterUpsertHook:
		authUserAfterUpsertHooks = append(authUserAfterUpsertHooks, authUserHook)
	}
}

// OneP returns a single authUser record from the query, and panics on error.
func (q authUserQuery) OneP() *AuthUser {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single authUser record from the query.
func (q authUserQuery) One() (*AuthUser, error) {
	o := &AuthUser{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for auth_user")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all AuthUser records from the query, and panics on error.
func (q authUserQuery) AllP() AuthUserSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all AuthUser records from the query.
func (q authUserQuery) All() (AuthUserSlice, error) {
	var o AuthUserSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to AuthUser slice")
	}

	if len(authUserAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all AuthUser records in the query, and panics on error.
func (q authUserQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all AuthUser records in the query.
func (q authUserQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count auth_user rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q authUserQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q authUserQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if auth_user exists")
	}

	return count > 0, nil
}

// UserStockOrdersG retrieves all the stock_order's stock order via user_id column.
func (o *AuthUser) UserStockOrdersG(mods ...qm.QueryMod) stockOrderQuery {
	return o.UserStockOrders(boil.GetDB(), mods...)
}

// UserStockOrders retrieves all the stock_order's stock order with an executor via user_id column.
func (o *AuthUser) UserStockOrders(exec boil.Executor, mods ...qm.QueryMod) stockOrderQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"user_id\"=$1", o.AuthUserID),
	)

	query := StockOrders(exec, queryMods...)
	queries.SetFrom(query.Query, "\"stock_order\" as \"a\"")
	return query
}

// AuthUserInfosG retrieves all the auth_user_info's auth user info.
func (o *AuthUser) AuthUserInfosG(mods ...qm.QueryMod) authUserInfoQuery {
	return o.AuthUserInfos(boil.GetDB(), mods...)
}

// AuthUserInfos retrieves all the auth_user_info's auth user info with an executor.
func (o *AuthUser) AuthUserInfos(exec boil.Executor, mods ...qm.QueryMod) authUserInfoQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"auth_user_id\"=$1", o.AuthUserID),
	)

	query := AuthUserInfos(exec, queryMods...)
	queries.SetFrom(query.Query, "\"auth_user_info\" as \"a\"")
	return query
}

// AuthUserRolesG retrieves all the auth_user_role's auth user role.
func (o *AuthUser) AuthUserRolesG(mods ...qm.QueryMod) authUserRoleQuery {
	return o.AuthUserRoles(boil.GetDB(), mods...)
}

// AuthUserRoles retrieves all the auth_user_role's auth user role with an executor.
func (o *AuthUser) AuthUserRoles(exec boil.Executor, mods ...qm.QueryMod) authUserRoleQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"auth_user_id\"=$1", o.AuthUserID),
	)

	query := AuthUserRoles(exec, queryMods...)
	queries.SetFrom(query.Query, "\"auth_user_role\" as \"a\"")
	return query
}

// LoadUserStockOrders allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (authUserL) LoadUserStockOrders(e boil.Executor, singular bool, maybeAuthUser interface{}) error {
	var slice []*AuthUser
	var object *AuthUser

	count := 1
	if singular {
		object = maybeAuthUser.(*AuthUser)
	} else {
		slice = *maybeAuthUser.(*AuthUserSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &authUserR{}
		args[0] = object.AuthUserID
	} else {
		for i, obj := range slice {
			obj.R = &authUserR{}
			args[i] = obj.AuthUserID
		}
	}

	query := fmt.Sprintf(
		"select * from \"stock_order\" where \"user_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load stock_order")
	}
	defer results.Close()

	var resultSlice []*StockOrder
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice stock_order")
	}

	if len(stockOrderAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.UserStockOrders = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.AuthUserID == foreign.UserID {
				local.R.UserStockOrders = append(local.R.UserStockOrders, foreign)
				break
			}
		}
	}

	return nil
}

// LoadAuthUserInfos allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (authUserL) LoadAuthUserInfos(e boil.Executor, singular bool, maybeAuthUser interface{}) error {
	var slice []*AuthUser
	var object *AuthUser

	count := 1
	if singular {
		object = maybeAuthUser.(*AuthUser)
	} else {
		slice = *maybeAuthUser.(*AuthUserSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &authUserR{}
		args[0] = object.AuthUserID
	} else {
		for i, obj := range slice {
			obj.R = &authUserR{}
			args[i] = obj.AuthUserID
		}
	}

	query := fmt.Sprintf(
		"select * from \"auth_user_info\" where \"auth_user_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load auth_user_info")
	}
	defer results.Close()

	var resultSlice []*AuthUserInfo
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice auth_user_info")
	}

	if len(authUserInfoAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.AuthUserInfos = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.AuthUserID == foreign.AuthUserID {
				local.R.AuthUserInfos = append(local.R.AuthUserInfos, foreign)
				break
			}
		}
	}

	return nil
}

// LoadAuthUserRoles allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (authUserL) LoadAuthUserRoles(e boil.Executor, singular bool, maybeAuthUser interface{}) error {
	var slice []*AuthUser
	var object *AuthUser

	count := 1
	if singular {
		object = maybeAuthUser.(*AuthUser)
	} else {
		slice = *maybeAuthUser.(*AuthUserSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &authUserR{}
		args[0] = object.AuthUserID
	} else {
		for i, obj := range slice {
			obj.R = &authUserR{}
			args[i] = obj.AuthUserID
		}
	}

	query := fmt.Sprintf(
		"select * from \"auth_user_role\" where \"auth_user_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load auth_user_role")
	}
	defer results.Close()

	var resultSlice []*AuthUserRole
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice auth_user_role")
	}

	if len(authUserRoleAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.AuthUserRoles = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.AuthUserID == foreign.AuthUserID {
				local.R.AuthUserRoles = append(local.R.AuthUserRoles, foreign)
				break
			}
		}
	}

	return nil
}

// AddUserStockOrders adds the given related objects to the existing relationships
// of the auth_user, optionally inserting them as new records.
// Appends related to o.R.UserStockOrders.
// Sets related.R.User appropriately.
func (o *AuthUser) AddUserStockOrders(exec boil.Executor, insert bool, related ...*StockOrder) error {
	var err error
	for _, rel := range related {
		rel.UserID = o.AuthUserID
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			if err = rel.Update(exec, "user_id"); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}
		}
	}

	if o.R == nil {
		o.R = &authUserR{
			UserStockOrders: related,
		}
	} else {
		o.R.UserStockOrders = append(o.R.UserStockOrders, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &stockOrderR{
				User: o,
			}
		} else {
			rel.R.User = o
		}
	}
	return nil
}

// AddAuthUserInfos adds the given related objects to the existing relationships
// of the auth_user, optionally inserting them as new records.
// Appends related to o.R.AuthUserInfos.
// Sets related.R.AuthUser appropriately.
func (o *AuthUser) AddAuthUserInfos(exec boil.Executor, insert bool, related ...*AuthUserInfo) error {
	var err error
	for _, rel := range related {
		rel.AuthUserID = o.AuthUserID
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			if err = rel.Update(exec, "auth_user_id"); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}
		}
	}

	if o.R == nil {
		o.R = &authUserR{
			AuthUserInfos: related,
		}
	} else {
		o.R.AuthUserInfos = append(o.R.AuthUserInfos, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &authUserInfoR{
				AuthUser: o,
			}
		} else {
			rel.R.AuthUser = o
		}
	}
	return nil
}

// AddAuthUserRoles adds the given related objects to the existing relationships
// of the auth_user, optionally inserting them as new records.
// Appends related to o.R.AuthUserRoles.
// Sets related.R.AuthUser appropriately.
func (o *AuthUser) AddAuthUserRoles(exec boil.Executor, insert bool, related ...*AuthUserRole) error {
	var err error
	for _, rel := range related {
		rel.AuthUserID = o.AuthUserID
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			if err = rel.Update(exec, "auth_user_id"); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}
		}
	}

	if o.R == nil {
		o.R = &authUserR{
			AuthUserRoles: related,
		}
	} else {
		o.R.AuthUserRoles = append(o.R.AuthUserRoles, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &authUserRoleR{
				AuthUser: o,
			}
		} else {
			rel.R.AuthUser = o
		}
	}
	return nil
}

// AuthUsersG retrieves all records.
func AuthUsersG(mods ...qm.QueryMod) authUserQuery {
	return AuthUsers(boil.GetDB(), mods...)
}

// AuthUsers retrieves all the records using an executor.
func AuthUsers(exec boil.Executor, mods ...qm.QueryMod) authUserQuery {
	mods = append(mods, qm.From("\"auth_user\""))
	return authUserQuery{NewQuery(exec, mods...)}
}

// FindAuthUserG retrieves a single record by ID.
func FindAuthUserG(authUserID int, selectCols ...string) (*AuthUser, error) {
	return FindAuthUser(boil.GetDB(), authUserID, selectCols...)
}

// FindAuthUserGP retrieves a single record by ID, and panics on error.
func FindAuthUserGP(authUserID int, selectCols ...string) *AuthUser {
	retobj, err := FindAuthUser(boil.GetDB(), authUserID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindAuthUser retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAuthUser(exec boil.Executor, authUserID int, selectCols ...string) (*AuthUser, error) {
	authUserObj := &AuthUser{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"auth_user\" where \"auth_user_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, authUserID)

	err := q.Bind(authUserObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from auth_user")
	}

	return authUserObj, nil
}

// FindAuthUserP retrieves a single record by ID with an executor, and panics on error.
func FindAuthUserP(exec boil.Executor, authUserID int, selectCols ...string) *AuthUser {
	retobj, err := FindAuthUser(exec, authUserID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *AuthUser) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *AuthUser) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *AuthUser) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *AuthUser) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no auth_user provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.Time.IsZero() {
		o.CreatedAt.Time = currTime
		o.CreatedAt.Valid = true
	}
	if o.UpdatedAt.Time.IsZero() {
		o.UpdatedAt.Time = currTime
		o.UpdatedAt.Valid = true
	}

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(authUserColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	authUserInsertCacheMut.RLock()
	cache, cached := authUserInsertCache[key]
	authUserInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			authUserColumns,
			authUserColumnsWithDefault,
			authUserColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(authUserType, authUserMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(authUserType, authUserMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"auth_user\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

		if len(cache.retMapping) != 0 {
			cache.query += fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "chado: unable to insert into auth_user")
	}

	if !cached {
		authUserInsertCacheMut.Lock()
		authUserInsertCache[key] = cache
		authUserInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single AuthUser record. See Update for
// whitelist behavior description.
func (o *AuthUser) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single AuthUser record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *AuthUser) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the AuthUser, and panics on error.
// See Update for whitelist behavior description.
func (o *AuthUser) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the AuthUser.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *AuthUser) Update(exec boil.Executor, whitelist ...string) error {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt.Time = currTime
	o.UpdatedAt.Valid = true

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	authUserUpdateCacheMut.RLock()
	cache, cached := authUserUpdateCache[key]
	authUserUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(authUserColumns, authUserPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update auth_user, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"auth_user\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, authUserPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(authUserType, authUserMapping, append(wl, authUserPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update auth_user row")
	}

	if !cached {
		authUserUpdateCacheMut.Lock()
		authUserUpdateCache[key] = cache
		authUserUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q authUserQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q authUserQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for auth_user")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AuthUserSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o AuthUserSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o AuthUserSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AuthUserSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("chado: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"auth_user\" SET %s WHERE (\"auth_user_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(authUserPrimaryKeyColumns), len(colNames)+1, len(authUserPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in authUser slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *AuthUser) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *AuthUser) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *AuthUser) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *AuthUser) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no auth_user provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.Time.IsZero() {
		o.CreatedAt.Time = currTime
		o.CreatedAt.Valid = true
	}
	o.UpdatedAt.Time = currTime
	o.UpdatedAt.Valid = true

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(authUserColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	authUserUpsertCacheMut.RLock()
	cache, cached := authUserUpsertCache[key]
	authUserUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			authUserColumns,
			authUserColumnsWithDefault,
			authUserColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			authUserColumns,
			authUserPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert auth_user, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(authUserPrimaryKeyColumns))
			copy(conflict, authUserPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"auth_user\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(authUserType, authUserMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(authUserType, authUserMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "chado: unable to upsert for auth_user")
	}

	if !cached {
		authUserUpsertCacheMut.Lock()
		authUserUpsertCache[key] = cache
		authUserUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single AuthUser record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *AuthUser) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single AuthUser record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *AuthUser) DeleteG() error {
	if o == nil {
		return errors.New("chado: no AuthUser provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single AuthUser record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *AuthUser) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single AuthUser record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AuthUser) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no AuthUser provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), authUserPrimaryKeyMapping)
	sql := "DELETE FROM \"auth_user\" WHERE \"auth_user_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from auth_user")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q authUserQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q authUserQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no authUserQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from auth_user")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o AuthUserSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o AuthUserSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no AuthUser slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o AuthUserSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AuthUserSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no AuthUser slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(authUserBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"auth_user\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, authUserPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(authUserPrimaryKeyColumns), 1, len(authUserPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from authUser slice")
	}

	if len(authUserAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *AuthUser) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *AuthUser) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *AuthUser) ReloadG() error {
	if o == nil {
		return errors.New("chado: no AuthUser provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *AuthUser) Reload(exec boil.Executor) error {
	ret, err := FindAuthUser(exec, o.AuthUserID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AuthUserSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AuthUserSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AuthUserSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty AuthUserSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AuthUserSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	authUsers := AuthUserSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"auth_user\".* FROM \"auth_user\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, authUserPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(authUserPrimaryKeyColumns), 1, len(authUserPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&authUsers)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in AuthUserSlice")
	}

	*o = authUsers

	return nil
}

// AuthUserExists checks if the AuthUser row exists.
func AuthUserExists(exec boil.Executor, authUserID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"auth_user\" where \"auth_user_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, authUserID)
	}

	row := exec.QueryRow(sql, authUserID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if auth_user exists")
	}

	return exists, nil
}

// AuthUserExistsG checks if the AuthUser row exists.
func AuthUserExistsG(authUserID int) (bool, error) {
	return AuthUserExists(boil.GetDB(), authUserID)
}

// AuthUserExistsGP checks if the AuthUser row exists. Panics on error.
func AuthUserExistsGP(authUserID int) bool {
	e, err := AuthUserExists(boil.GetDB(), authUserID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// AuthUserExistsP checks if the AuthUser row exists. Panics on error.
func AuthUserExistsP(exec boil.Executor, authUserID int) bool {
	e, err := AuthUserExists(exec, authUserID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
