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
)

// AuthUserRole is an object representing the database table.
type AuthUserRole struct {
	AuthUserRoleID int `boil:"auth_user_role_id" json:"auth_user_role_id" toml:"auth_user_role_id" yaml:"auth_user_role_id"`
	AuthUserID     int `boil:"auth_user_id" json:"auth_user_id" toml:"auth_user_id" yaml:"auth_user_id"`
	AuthRoleID     int `boil:"auth_role_id" json:"auth_role_id" toml:"auth_role_id" yaml:"auth_role_id"`

	R *authUserRoleR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L authUserRoleL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// authUserRoleR is where relationships are stored.
type authUserRoleR struct {
	AuthUser *AuthUser
	AuthRole *AuthRole
}

// authUserRoleL is where Load methods for each relationship are stored.
type authUserRoleL struct{}

var (
	authUserRoleColumns               = []string{"auth_user_role_id", "auth_user_id", "auth_role_id"}
	authUserRoleColumnsWithoutDefault = []string{"auth_user_id", "auth_role_id"}
	authUserRoleColumnsWithDefault    = []string{"auth_user_role_id"}
	authUserRolePrimaryKeyColumns     = []string{"auth_user_role_id"}
)

type (
	// AuthUserRoleSlice is an alias for a slice of pointers to AuthUserRole.
	// This should generally be used opposed to []AuthUserRole.
	AuthUserRoleSlice []*AuthUserRole
	// AuthUserRoleHook is the signature for custom AuthUserRole hook methods
	AuthUserRoleHook func(boil.Executor, *AuthUserRole) error

	authUserRoleQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	authUserRoleType                 = reflect.TypeOf(&AuthUserRole{})
	authUserRoleMapping              = queries.MakeStructMapping(authUserRoleType)
	authUserRolePrimaryKeyMapping, _ = queries.BindMapping(authUserRoleType, authUserRoleMapping, authUserRolePrimaryKeyColumns)
	authUserRoleInsertCacheMut       sync.RWMutex
	authUserRoleInsertCache          = make(map[string]insertCache)
	authUserRoleUpdateCacheMut       sync.RWMutex
	authUserRoleUpdateCache          = make(map[string]updateCache)
	authUserRoleUpsertCacheMut       sync.RWMutex
	authUserRoleUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var authUserRoleBeforeInsertHooks []AuthUserRoleHook
var authUserRoleBeforeUpdateHooks []AuthUserRoleHook
var authUserRoleBeforeDeleteHooks []AuthUserRoleHook
var authUserRoleBeforeUpsertHooks []AuthUserRoleHook

var authUserRoleAfterInsertHooks []AuthUserRoleHook
var authUserRoleAfterSelectHooks []AuthUserRoleHook
var authUserRoleAfterUpdateHooks []AuthUserRoleHook
var authUserRoleAfterDeleteHooks []AuthUserRoleHook
var authUserRoleAfterUpsertHooks []AuthUserRoleHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AuthUserRole) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserRoleBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AuthUserRole) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserRoleBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AuthUserRole) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserRoleBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AuthUserRole) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserRoleBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AuthUserRole) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserRoleAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AuthUserRole) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserRoleAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AuthUserRole) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserRoleAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AuthUserRole) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserRoleAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AuthUserRole) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserRoleAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAuthUserRoleHook registers your hook function for all future operations.
func AddAuthUserRoleHook(hookPoint boil.HookPoint, authUserRoleHook AuthUserRoleHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		authUserRoleBeforeInsertHooks = append(authUserRoleBeforeInsertHooks, authUserRoleHook)
	case boil.BeforeUpdateHook:
		authUserRoleBeforeUpdateHooks = append(authUserRoleBeforeUpdateHooks, authUserRoleHook)
	case boil.BeforeDeleteHook:
		authUserRoleBeforeDeleteHooks = append(authUserRoleBeforeDeleteHooks, authUserRoleHook)
	case boil.BeforeUpsertHook:
		authUserRoleBeforeUpsertHooks = append(authUserRoleBeforeUpsertHooks, authUserRoleHook)
	case boil.AfterInsertHook:
		authUserRoleAfterInsertHooks = append(authUserRoleAfterInsertHooks, authUserRoleHook)
	case boil.AfterSelectHook:
		authUserRoleAfterSelectHooks = append(authUserRoleAfterSelectHooks, authUserRoleHook)
	case boil.AfterUpdateHook:
		authUserRoleAfterUpdateHooks = append(authUserRoleAfterUpdateHooks, authUserRoleHook)
	case boil.AfterDeleteHook:
		authUserRoleAfterDeleteHooks = append(authUserRoleAfterDeleteHooks, authUserRoleHook)
	case boil.AfterUpsertHook:
		authUserRoleAfterUpsertHooks = append(authUserRoleAfterUpsertHooks, authUserRoleHook)
	}
}

// OneP returns a single authUserRole record from the query, and panics on error.
func (q authUserRoleQuery) OneP() *AuthUserRole {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single authUserRole record from the query.
func (q authUserRoleQuery) One() (*AuthUserRole, error) {
	o := &AuthUserRole{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for auth_user_role")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all AuthUserRole records from the query, and panics on error.
func (q authUserRoleQuery) AllP() AuthUserRoleSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all AuthUserRole records from the query.
func (q authUserRoleQuery) All() (AuthUserRoleSlice, error) {
	var o AuthUserRoleSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AuthUserRole slice")
	}

	if len(authUserRoleAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all AuthUserRole records in the query, and panics on error.
func (q authUserRoleQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all AuthUserRole records in the query.
func (q authUserRoleQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count auth_user_role rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q authUserRoleQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q authUserRoleQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if auth_user_role exists")
	}

	return count > 0, nil
}

// AuthUserG pointed to by the foreign key.
func (o *AuthUserRole) AuthUserG(mods ...qm.QueryMod) authUserQuery {
	return o.AuthUser(boil.GetDB(), mods...)
}

// AuthUser pointed to by the foreign key.
func (o *AuthUserRole) AuthUser(exec boil.Executor, mods ...qm.QueryMod) authUserQuery {
	queryMods := []qm.QueryMod{
		qm.Where("auth_user_id=$1", o.AuthUserID),
	}

	queryMods = append(queryMods, mods...)

	query := AuthUsers(exec, queryMods...)
	queries.SetFrom(query.Query, "\"auth_user\"")

	return query
}

// AuthRoleG pointed to by the foreign key.
func (o *AuthUserRole) AuthRoleG(mods ...qm.QueryMod) authRoleQuery {
	return o.AuthRole(boil.GetDB(), mods...)
}

// AuthRole pointed to by the foreign key.
func (o *AuthUserRole) AuthRole(exec boil.Executor, mods ...qm.QueryMod) authRoleQuery {
	queryMods := []qm.QueryMod{
		qm.Where("auth_role_id=$1", o.AuthRoleID),
	}

	queryMods = append(queryMods, mods...)

	query := AuthRoles(exec, queryMods...)
	queries.SetFrom(query.Query, "\"auth_role\"")

	return query
}

// LoadAuthUser allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (authUserRoleL) LoadAuthUser(e boil.Executor, singular bool, maybeAuthUserRole interface{}) error {
	var slice []*AuthUserRole
	var object *AuthUserRole

	count := 1
	if singular {
		object = maybeAuthUserRole.(*AuthUserRole)
	} else {
		slice = *maybeAuthUserRole.(*AuthUserRoleSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &authUserRoleR{}
		args[0] = object.AuthUserID
	} else {
		for i, obj := range slice {
			obj.R = &authUserRoleR{}
			args[i] = obj.AuthUserID
		}
	}

	query := fmt.Sprintf(
		"select * from \"auth_user\" where \"auth_user_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load AuthUser")
	}
	defer results.Close()

	var resultSlice []*AuthUser
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice AuthUser")
	}

	if len(authUserRoleAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.AuthUser = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.AuthUserID == foreign.AuthUserID {
				local.R.AuthUser = foreign
				break
			}
		}
	}

	return nil
}

// LoadAuthRole allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (authUserRoleL) LoadAuthRole(e boil.Executor, singular bool, maybeAuthUserRole interface{}) error {
	var slice []*AuthUserRole
	var object *AuthUserRole

	count := 1
	if singular {
		object = maybeAuthUserRole.(*AuthUserRole)
	} else {
		slice = *maybeAuthUserRole.(*AuthUserRoleSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &authUserRoleR{}
		args[0] = object.AuthRoleID
	} else {
		for i, obj := range slice {
			obj.R = &authUserRoleR{}
			args[i] = obj.AuthRoleID
		}
	}

	query := fmt.Sprintf(
		"select * from \"auth_role\" where \"auth_role_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load AuthRole")
	}
	defer results.Close()

	var resultSlice []*AuthRole
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice AuthRole")
	}

	if len(authUserRoleAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.AuthRole = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.AuthRoleID == foreign.AuthRoleID {
				local.R.AuthRole = foreign
				break
			}
		}
	}

	return nil
}

// SetAuthUser of the auth_user_role to the related item.
// Sets o.R.AuthUser to related.
// Adds o to related.R.AuthUserRoles.
func (o *AuthUserRole) SetAuthUser(exec boil.Executor, insert bool, related *AuthUser) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"auth_user_role\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"auth_user_id"}),
		strmangle.WhereClause("\"", "\"", 2, authUserRolePrimaryKeyColumns),
	)
	values := []interface{}{related.AuthUserID, o.AuthUserRoleID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.AuthUserID = related.AuthUserID

	if o.R == nil {
		o.R = &authUserRoleR{
			AuthUser: related,
		}
	} else {
		o.R.AuthUser = related
	}

	if related.R == nil {
		related.R = &authUserR{
			AuthUserRoles: AuthUserRoleSlice{o},
		}
	} else {
		related.R.AuthUserRoles = append(related.R.AuthUserRoles, o)
	}

	return nil
}

// SetAuthRole of the auth_user_role to the related item.
// Sets o.R.AuthRole to related.
// Adds o to related.R.AuthUserRoles.
func (o *AuthUserRole) SetAuthRole(exec boil.Executor, insert bool, related *AuthRole) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"auth_user_role\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"auth_role_id"}),
		strmangle.WhereClause("\"", "\"", 2, authUserRolePrimaryKeyColumns),
	)
	values := []interface{}{related.AuthRoleID, o.AuthUserRoleID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.AuthRoleID = related.AuthRoleID

	if o.R == nil {
		o.R = &authUserRoleR{
			AuthRole: related,
		}
	} else {
		o.R.AuthRole = related
	}

	if related.R == nil {
		related.R = &authRoleR{
			AuthUserRoles: AuthUserRoleSlice{o},
		}
	} else {
		related.R.AuthUserRoles = append(related.R.AuthUserRoles, o)
	}

	return nil
}

// AuthUserRolesG retrieves all records.
func AuthUserRolesG(mods ...qm.QueryMod) authUserRoleQuery {
	return AuthUserRoles(boil.GetDB(), mods...)
}

// AuthUserRoles retrieves all the records using an executor.
func AuthUserRoles(exec boil.Executor, mods ...qm.QueryMod) authUserRoleQuery {
	mods = append(mods, qm.From("\"auth_user_role\""))
	return authUserRoleQuery{NewQuery(exec, mods...)}
}

// FindAuthUserRoleG retrieves a single record by ID.
func FindAuthUserRoleG(authUserRoleID int, selectCols ...string) (*AuthUserRole, error) {
	return FindAuthUserRole(boil.GetDB(), authUserRoleID, selectCols...)
}

// FindAuthUserRoleGP retrieves a single record by ID, and panics on error.
func FindAuthUserRoleGP(authUserRoleID int, selectCols ...string) *AuthUserRole {
	retobj, err := FindAuthUserRole(boil.GetDB(), authUserRoleID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindAuthUserRole retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAuthUserRole(exec boil.Executor, authUserRoleID int, selectCols ...string) (*AuthUserRole, error) {
	authUserRoleObj := &AuthUserRole{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"auth_user_role\" where \"auth_user_role_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, authUserRoleID)

	err := q.Bind(authUserRoleObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from auth_user_role")
	}

	return authUserRoleObj, nil
}

// FindAuthUserRoleP retrieves a single record by ID with an executor, and panics on error.
func FindAuthUserRoleP(exec boil.Executor, authUserRoleID int, selectCols ...string) *AuthUserRole {
	retobj, err := FindAuthUserRole(exec, authUserRoleID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *AuthUserRole) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *AuthUserRole) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *AuthUserRole) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *AuthUserRole) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no auth_user_role provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(authUserRoleColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	authUserRoleInsertCacheMut.RLock()
	cache, cached := authUserRoleInsertCache[key]
	authUserRoleInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			authUserRoleColumns,
			authUserRoleColumnsWithDefault,
			authUserRoleColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(authUserRoleType, authUserRoleMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(authUserRoleType, authUserRoleMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"auth_user_role\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into auth_user_role")
	}

	if !cached {
		authUserRoleInsertCacheMut.Lock()
		authUserRoleInsertCache[key] = cache
		authUserRoleInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single AuthUserRole record. See Update for
// whitelist behavior description.
func (o *AuthUserRole) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single AuthUserRole record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *AuthUserRole) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the AuthUserRole, and panics on error.
// See Update for whitelist behavior description.
func (o *AuthUserRole) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the AuthUserRole.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *AuthUserRole) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	authUserRoleUpdateCacheMut.RLock()
	cache, cached := authUserRoleUpdateCache[key]
	authUserRoleUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(authUserRoleColumns, authUserRolePrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update auth_user_role, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"auth_user_role\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, authUserRolePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(authUserRoleType, authUserRoleMapping, append(wl, authUserRolePrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update auth_user_role row")
	}

	if !cached {
		authUserRoleUpdateCacheMut.Lock()
		authUserRoleUpdateCache[key] = cache
		authUserRoleUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q authUserRoleQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q authUserRoleQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for auth_user_role")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AuthUserRoleSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o AuthUserRoleSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o AuthUserRoleSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AuthUserRoleSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authUserRolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"auth_user_role\" SET %s WHERE (\"auth_user_role_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(authUserRolePrimaryKeyColumns), len(colNames)+1, len(authUserRolePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in authUserRole slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *AuthUserRole) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *AuthUserRole) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *AuthUserRole) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *AuthUserRole) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no auth_user_role provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(authUserRoleColumnsWithDefault, o)

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

	authUserRoleUpsertCacheMut.RLock()
	cache, cached := authUserRoleUpsertCache[key]
	authUserRoleUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			authUserRoleColumns,
			authUserRoleColumnsWithDefault,
			authUserRoleColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			authUserRoleColumns,
			authUserRolePrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert auth_user_role, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(authUserRolePrimaryKeyColumns))
			copy(conflict, authUserRolePrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"auth_user_role\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(authUserRoleType, authUserRoleMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(authUserRoleType, authUserRoleMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for auth_user_role")
	}

	if !cached {
		authUserRoleUpsertCacheMut.Lock()
		authUserRoleUpsertCache[key] = cache
		authUserRoleUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single AuthUserRole record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *AuthUserRole) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single AuthUserRole record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *AuthUserRole) DeleteG() error {
	if o == nil {
		return errors.New("models: no AuthUserRole provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single AuthUserRole record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *AuthUserRole) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single AuthUserRole record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AuthUserRole) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no AuthUserRole provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), authUserRolePrimaryKeyMapping)
	sql := "DELETE FROM \"auth_user_role\" WHERE \"auth_user_role_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from auth_user_role")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q authUserRoleQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q authUserRoleQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no authUserRoleQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from auth_user_role")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o AuthUserRoleSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o AuthUserRoleSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no AuthUserRole slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o AuthUserRoleSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AuthUserRoleSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no AuthUserRole slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(authUserRoleBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authUserRolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"auth_user_role\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, authUserRolePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(authUserRolePrimaryKeyColumns), 1, len(authUserRolePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from authUserRole slice")
	}

	if len(authUserRoleAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *AuthUserRole) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *AuthUserRole) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *AuthUserRole) ReloadG() error {
	if o == nil {
		return errors.New("models: no AuthUserRole provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *AuthUserRole) Reload(exec boil.Executor) error {
	ret, err := FindAuthUserRole(exec, o.AuthUserRoleID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AuthUserRoleSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AuthUserRoleSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AuthUserRoleSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty AuthUserRoleSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AuthUserRoleSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	authUserRoles := AuthUserRoleSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authUserRolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"auth_user_role\".* FROM \"auth_user_role\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, authUserRolePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(authUserRolePrimaryKeyColumns), 1, len(authUserRolePrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&authUserRoles)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AuthUserRoleSlice")
	}

	*o = authUserRoles

	return nil
}

// AuthUserRoleExists checks if the AuthUserRole row exists.
func AuthUserRoleExists(exec boil.Executor, authUserRoleID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"auth_user_role\" where \"auth_user_role_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, authUserRoleID)
	}

	row := exec.QueryRow(sql, authUserRoleID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if auth_user_role exists")
	}

	return exists, nil
}

// AuthUserRoleExistsG checks if the AuthUserRole row exists.
func AuthUserRoleExistsG(authUserRoleID int) (bool, error) {
	return AuthUserRoleExists(boil.GetDB(), authUserRoleID)
}

// AuthUserRoleExistsGP checks if the AuthUserRole row exists. Panics on error.
func AuthUserRoleExistsGP(authUserRoleID int) bool {
	e, err := AuthUserRoleExists(boil.GetDB(), authUserRoleID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// AuthUserRoleExistsP checks if the AuthUserRole row exists. Panics on error.
func AuthUserRoleExistsP(exec boil.Executor, authUserRoleID int) bool {
	e, err := AuthUserRoleExists(exec, authUserRoleID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
