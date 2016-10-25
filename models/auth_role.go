package models

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

// AuthRole is an object representing the database table.
type AuthRole struct {
	AuthRoleID  int         `boil:"auth_role_id" json:"auth_role_id" toml:"auth_role_id" yaml:"auth_role_id"`
	Role        string      `boil:"role" json:"role" toml:"role" yaml:"role"`
	Description null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	CreatedAt   null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt   null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *authRoleR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L authRoleL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// authRoleR is where relationships are stored.
type authRoleR struct {
	AuthRolePermission *AuthRolePermission
	AuthUserRoles      AuthUserRoleSlice
}

// authRoleL is where Load methods for each relationship are stored.
type authRoleL struct{}

var (
	authRoleColumns               = []string{"auth_role_id", "role", "description", "created_at", "updated_at"}
	authRoleColumnsWithoutDefault = []string{"role", "description"}
	authRoleColumnsWithDefault    = []string{"auth_role_id", "created_at", "updated_at"}
	authRolePrimaryKeyColumns     = []string{"auth_role_id"}
)

type (
	// AuthRoleSlice is an alias for a slice of pointers to AuthRole.
	// This should generally be used opposed to []AuthRole.
	AuthRoleSlice []*AuthRole
	// AuthRoleHook is the signature for custom AuthRole hook methods
	AuthRoleHook func(boil.Executor, *AuthRole) error

	authRoleQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	authRoleType                 = reflect.TypeOf(&AuthRole{})
	authRoleMapping              = queries.MakeStructMapping(authRoleType)
	authRolePrimaryKeyMapping, _ = queries.BindMapping(authRoleType, authRoleMapping, authRolePrimaryKeyColumns)
	authRoleInsertCacheMut       sync.RWMutex
	authRoleInsertCache          = make(map[string]insertCache)
	authRoleUpdateCacheMut       sync.RWMutex
	authRoleUpdateCache          = make(map[string]updateCache)
	authRoleUpsertCacheMut       sync.RWMutex
	authRoleUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var authRoleBeforeInsertHooks []AuthRoleHook
var authRoleBeforeUpdateHooks []AuthRoleHook
var authRoleBeforeDeleteHooks []AuthRoleHook
var authRoleBeforeUpsertHooks []AuthRoleHook

var authRoleAfterInsertHooks []AuthRoleHook
var authRoleAfterSelectHooks []AuthRoleHook
var authRoleAfterUpdateHooks []AuthRoleHook
var authRoleAfterDeleteHooks []AuthRoleHook
var authRoleAfterUpsertHooks []AuthRoleHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AuthRole) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authRoleBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AuthRole) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range authRoleBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AuthRole) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range authRoleBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AuthRole) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authRoleBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AuthRole) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authRoleAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AuthRole) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range authRoleAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AuthRole) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range authRoleAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AuthRole) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range authRoleAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AuthRole) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authRoleAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAuthRoleHook registers your hook function for all future operations.
func AddAuthRoleHook(hookPoint boil.HookPoint, authRoleHook AuthRoleHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		authRoleBeforeInsertHooks = append(authRoleBeforeInsertHooks, authRoleHook)
	case boil.BeforeUpdateHook:
		authRoleBeforeUpdateHooks = append(authRoleBeforeUpdateHooks, authRoleHook)
	case boil.BeforeDeleteHook:
		authRoleBeforeDeleteHooks = append(authRoleBeforeDeleteHooks, authRoleHook)
	case boil.BeforeUpsertHook:
		authRoleBeforeUpsertHooks = append(authRoleBeforeUpsertHooks, authRoleHook)
	case boil.AfterInsertHook:
		authRoleAfterInsertHooks = append(authRoleAfterInsertHooks, authRoleHook)
	case boil.AfterSelectHook:
		authRoleAfterSelectHooks = append(authRoleAfterSelectHooks, authRoleHook)
	case boil.AfterUpdateHook:
		authRoleAfterUpdateHooks = append(authRoleAfterUpdateHooks, authRoleHook)
	case boil.AfterDeleteHook:
		authRoleAfterDeleteHooks = append(authRoleAfterDeleteHooks, authRoleHook)
	case boil.AfterUpsertHook:
		authRoleAfterUpsertHooks = append(authRoleAfterUpsertHooks, authRoleHook)
	}
}

// OneP returns a single authRole record from the query, and panics on error.
func (q authRoleQuery) OneP() *AuthRole {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single authRole record from the query.
func (q authRoleQuery) One() (*AuthRole, error) {
	o := &AuthRole{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for auth_role")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all AuthRole records from the query, and panics on error.
func (q authRoleQuery) AllP() AuthRoleSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all AuthRole records from the query.
func (q authRoleQuery) All() (AuthRoleSlice, error) {
	var o AuthRoleSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AuthRole slice")
	}

	if len(authRoleAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all AuthRole records in the query, and panics on error.
func (q authRoleQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all AuthRole records in the query.
func (q authRoleQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count auth_role rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q authRoleQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q authRoleQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if auth_role exists")
	}

	return count > 0, nil
}

// AuthRolePermissionG pointed to by the foreign key.
func (o *AuthRole) AuthRolePermissionG(mods ...qm.QueryMod) authRolePermissionQuery {
	return o.AuthRolePermission(boil.GetDB(), mods...)
}

// AuthRolePermission pointed to by the foreign key.
func (o *AuthRole) AuthRolePermission(exec boil.Executor, mods ...qm.QueryMod) authRolePermissionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("auth_role_id=$1", o.AuthRoleID),
	}

	queryMods = append(queryMods, mods...)

	query := AuthRolePermissions(exec, queryMods...)
	queries.SetFrom(query.Query, "\"auth_role_permission\"")

	return query
}

// AuthUserRolesG retrieves all the auth_user_role's auth user role.
func (o *AuthRole) AuthUserRolesG(mods ...qm.QueryMod) authUserRoleQuery {
	return o.AuthUserRoles(boil.GetDB(), mods...)
}

// AuthUserRoles retrieves all the auth_user_role's auth user role with an executor.
func (o *AuthRole) AuthUserRoles(exec boil.Executor, mods ...qm.QueryMod) authUserRoleQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"auth_role_id\"=$1", o.AuthRoleID),
	)

	query := AuthUserRoles(exec, queryMods...)
	queries.SetFrom(query.Query, "\"auth_user_role\" as \"a\"")
	return query
}

// LoadAuthRolePermission allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (authRoleL) LoadAuthRolePermission(e boil.Executor, singular bool, maybeAuthRole interface{}) error {
	var slice []*AuthRole
	var object *AuthRole

	count := 1
	if singular {
		object = maybeAuthRole.(*AuthRole)
	} else {
		slice = *maybeAuthRole.(*AuthRoleSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &authRoleR{}
		args[0] = object.AuthRoleID
	} else {
		for i, obj := range slice {
			obj.R = &authRoleR{}
			args[i] = obj.AuthRoleID
		}
	}

	query := fmt.Sprintf(
		"select * from \"auth_role_permission\" where \"auth_role_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load AuthRolePermission")
	}
	defer results.Close()

	var resultSlice []*AuthRolePermission
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice AuthRolePermission")
	}

	if len(authRoleAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.AuthRolePermission = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.AuthRoleID == foreign.AuthRoleID {
				local.R.AuthRolePermission = foreign
				break
			}
		}
	}

	return nil
}

// LoadAuthUserRoles allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (authRoleL) LoadAuthUserRoles(e boil.Executor, singular bool, maybeAuthRole interface{}) error {
	var slice []*AuthRole
	var object *AuthRole

	count := 1
	if singular {
		object = maybeAuthRole.(*AuthRole)
	} else {
		slice = *maybeAuthRole.(*AuthRoleSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &authRoleR{}
		args[0] = object.AuthRoleID
	} else {
		for i, obj := range slice {
			obj.R = &authRoleR{}
			args[i] = obj.AuthRoleID
		}
	}

	query := fmt.Sprintf(
		"select * from \"auth_user_role\" where \"auth_role_id\" in (%s)",
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
			if local.AuthRoleID == foreign.AuthRoleID {
				local.R.AuthUserRoles = append(local.R.AuthUserRoles, foreign)
				break
			}
		}
	}

	return nil
}

// SetAuthRolePermission of the auth_role to the related item.
// Sets o.R.AuthRolePermission to related.
// Adds o to related.R.AuthRole.
func (o *AuthRole) SetAuthRolePermission(exec boil.Executor, insert bool, related *AuthRolePermission) error {
	var err error

	if insert {
		related.AuthRoleID = o.AuthRoleID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"auth_role_permission\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"auth_role_id"}),
			strmangle.WhereClause("\"", "\"", 2, authRolePermissionPrimaryKeyColumns),
		)
		values := []interface{}{o.AuthRoleID, related.AuthRolePermissionID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.AuthRoleID = o.AuthRoleID

	}

	if o.R == nil {
		o.R = &authRoleR{
			AuthRolePermission: related,
		}
	} else {
		o.R.AuthRolePermission = related
	}

	if related.R == nil {
		related.R = &authRolePermissionR{
			AuthRole: o,
		}
	} else {
		related.R.AuthRole = o
	}
	return nil
}

// AddAuthUserRoles adds the given related objects to the existing relationships
// of the auth_role, optionally inserting them as new records.
// Appends related to o.R.AuthUserRoles.
// Sets related.R.AuthRole appropriately.
func (o *AuthRole) AddAuthUserRoles(exec boil.Executor, insert bool, related ...*AuthUserRole) error {
	var err error
	for _, rel := range related {
		rel.AuthRoleID = o.AuthRoleID
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			if err = rel.Update(exec, "auth_role_id"); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}
		}
	}

	if o.R == nil {
		o.R = &authRoleR{
			AuthUserRoles: related,
		}
	} else {
		o.R.AuthUserRoles = append(o.R.AuthUserRoles, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &authUserRoleR{
				AuthRole: o,
			}
		} else {
			rel.R.AuthRole = o
		}
	}
	return nil
}

// AuthRolesG retrieves all records.
func AuthRolesG(mods ...qm.QueryMod) authRoleQuery {
	return AuthRoles(boil.GetDB(), mods...)
}

// AuthRoles retrieves all the records using an executor.
func AuthRoles(exec boil.Executor, mods ...qm.QueryMod) authRoleQuery {
	mods = append(mods, qm.From("\"auth_role\""))
	return authRoleQuery{NewQuery(exec, mods...)}
}

// FindAuthRoleG retrieves a single record by ID.
func FindAuthRoleG(authRoleID int, selectCols ...string) (*AuthRole, error) {
	return FindAuthRole(boil.GetDB(), authRoleID, selectCols...)
}

// FindAuthRoleGP retrieves a single record by ID, and panics on error.
func FindAuthRoleGP(authRoleID int, selectCols ...string) *AuthRole {
	retobj, err := FindAuthRole(boil.GetDB(), authRoleID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindAuthRole retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAuthRole(exec boil.Executor, authRoleID int, selectCols ...string) (*AuthRole, error) {
	authRoleObj := &AuthRole{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"auth_role\" where \"auth_role_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, authRoleID)

	err := q.Bind(authRoleObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from auth_role")
	}

	return authRoleObj, nil
}

// FindAuthRoleP retrieves a single record by ID with an executor, and panics on error.
func FindAuthRoleP(exec boil.Executor, authRoleID int, selectCols ...string) *AuthRole {
	retobj, err := FindAuthRole(exec, authRoleID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *AuthRole) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *AuthRole) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *AuthRole) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *AuthRole) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no auth_role provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(authRoleColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	authRoleInsertCacheMut.RLock()
	cache, cached := authRoleInsertCache[key]
	authRoleInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			authRoleColumns,
			authRoleColumnsWithDefault,
			authRoleColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(authRoleType, authRoleMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(authRoleType, authRoleMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"auth_role\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into auth_role")
	}

	if !cached {
		authRoleInsertCacheMut.Lock()
		authRoleInsertCache[key] = cache
		authRoleInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single AuthRole record. See Update for
// whitelist behavior description.
func (o *AuthRole) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single AuthRole record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *AuthRole) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the AuthRole, and panics on error.
// See Update for whitelist behavior description.
func (o *AuthRole) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the AuthRole.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *AuthRole) Update(exec boil.Executor, whitelist ...string) error {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt.Time = currTime
	o.UpdatedAt.Valid = true

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	authRoleUpdateCacheMut.RLock()
	cache, cached := authRoleUpdateCache[key]
	authRoleUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(authRoleColumns, authRolePrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update auth_role, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"auth_role\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, authRolePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(authRoleType, authRoleMapping, append(wl, authRolePrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update auth_role row")
	}

	if !cached {
		authRoleUpdateCacheMut.Lock()
		authRoleUpdateCache[key] = cache
		authRoleUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q authRoleQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q authRoleQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for auth_role")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AuthRoleSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o AuthRoleSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o AuthRoleSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AuthRoleSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authRolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"auth_role\" SET %s WHERE (\"auth_role_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(authRolePrimaryKeyColumns), len(colNames)+1, len(authRolePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in authRole slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *AuthRole) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *AuthRole) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *AuthRole) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *AuthRole) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no auth_role provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(authRoleColumnsWithDefault, o)

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

	authRoleUpsertCacheMut.RLock()
	cache, cached := authRoleUpsertCache[key]
	authRoleUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			authRoleColumns,
			authRoleColumnsWithDefault,
			authRoleColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			authRoleColumns,
			authRolePrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert auth_role, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(authRolePrimaryKeyColumns))
			copy(conflict, authRolePrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"auth_role\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(authRoleType, authRoleMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(authRoleType, authRoleMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for auth_role")
	}

	if !cached {
		authRoleUpsertCacheMut.Lock()
		authRoleUpsertCache[key] = cache
		authRoleUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single AuthRole record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *AuthRole) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single AuthRole record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *AuthRole) DeleteG() error {
	if o == nil {
		return errors.New("models: no AuthRole provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single AuthRole record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *AuthRole) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single AuthRole record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AuthRole) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no AuthRole provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), authRolePrimaryKeyMapping)
	sql := "DELETE FROM \"auth_role\" WHERE \"auth_role_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from auth_role")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q authRoleQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q authRoleQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no authRoleQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from auth_role")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o AuthRoleSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o AuthRoleSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no AuthRole slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o AuthRoleSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AuthRoleSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no AuthRole slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(authRoleBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authRolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"auth_role\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, authRolePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(authRolePrimaryKeyColumns), 1, len(authRolePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from authRole slice")
	}

	if len(authRoleAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *AuthRole) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *AuthRole) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *AuthRole) ReloadG() error {
	if o == nil {
		return errors.New("models: no AuthRole provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *AuthRole) Reload(exec boil.Executor) error {
	ret, err := FindAuthRole(exec, o.AuthRoleID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AuthRoleSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AuthRoleSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AuthRoleSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty AuthRoleSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AuthRoleSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	authRoles := AuthRoleSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authRolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"auth_role\".* FROM \"auth_role\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, authRolePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(authRolePrimaryKeyColumns), 1, len(authRolePrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&authRoles)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AuthRoleSlice")
	}

	*o = authRoles

	return nil
}

// AuthRoleExists checks if the AuthRole row exists.
func AuthRoleExists(exec boil.Executor, authRoleID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"auth_role\" where \"auth_role_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, authRoleID)
	}

	row := exec.QueryRow(sql, authRoleID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if auth_role exists")
	}

	return exists, nil
}

// AuthRoleExistsG checks if the AuthRole row exists.
func AuthRoleExistsG(authRoleID int) (bool, error) {
	return AuthRoleExists(boil.GetDB(), authRoleID)
}

// AuthRoleExistsGP checks if the AuthRole row exists. Panics on error.
func AuthRoleExistsGP(authRoleID int) bool {
	e, err := AuthRoleExists(boil.GetDB(), authRoleID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// AuthRoleExistsP checks if the AuthRole row exists. Panics on error.
func AuthRoleExistsP(exec boil.Executor, authRoleID int) bool {
	e, err := AuthRoleExists(exec, authRoleID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
