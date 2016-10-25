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

// AuthRolePermission is an object representing the database table.
type AuthRolePermission struct {
	AuthRolePermissionID int       `boil:"auth_role_permission_id" json:"auth_role_permission_id" toml:"auth_role_permission_id" yaml:"auth_role_permission_id"`
	AuthRoleID           int       `boil:"auth_role_id" json:"auth_role_id" toml:"auth_role_id" yaml:"auth_role_id"`
	AuthPermissionID     int       `boil:"auth_permission_id" json:"auth_permission_id" toml:"auth_permission_id" yaml:"auth_permission_id"`
	CreatedAt            null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt            null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *authRolePermissionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L authRolePermissionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// authRolePermissionR is where relationships are stored.
type authRolePermissionR struct {
	AuthRole       *AuthRole
	AuthPermission *AuthPermission
}

// authRolePermissionL is where Load methods for each relationship are stored.
type authRolePermissionL struct{}

var (
	authRolePermissionColumns               = []string{"auth_role_permission_id", "auth_role_id", "auth_permission_id", "created_at", "updated_at"}
	authRolePermissionColumnsWithoutDefault = []string{"auth_role_id", "auth_permission_id"}
	authRolePermissionColumnsWithDefault    = []string{"auth_role_permission_id", "created_at", "updated_at"}
	authRolePermissionPrimaryKeyColumns     = []string{"auth_role_permission_id"}
)

type (
	// AuthRolePermissionSlice is an alias for a slice of pointers to AuthRolePermission.
	// This should generally be used opposed to []AuthRolePermission.
	AuthRolePermissionSlice []*AuthRolePermission
	// AuthRolePermissionHook is the signature for custom AuthRolePermission hook methods
	AuthRolePermissionHook func(boil.Executor, *AuthRolePermission) error

	authRolePermissionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	authRolePermissionType                 = reflect.TypeOf(&AuthRolePermission{})
	authRolePermissionMapping              = queries.MakeStructMapping(authRolePermissionType)
	authRolePermissionPrimaryKeyMapping, _ = queries.BindMapping(authRolePermissionType, authRolePermissionMapping, authRolePermissionPrimaryKeyColumns)
	authRolePermissionInsertCacheMut       sync.RWMutex
	authRolePermissionInsertCache          = make(map[string]insertCache)
	authRolePermissionUpdateCacheMut       sync.RWMutex
	authRolePermissionUpdateCache          = make(map[string]updateCache)
	authRolePermissionUpsertCacheMut       sync.RWMutex
	authRolePermissionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var authRolePermissionBeforeInsertHooks []AuthRolePermissionHook
var authRolePermissionBeforeUpdateHooks []AuthRolePermissionHook
var authRolePermissionBeforeDeleteHooks []AuthRolePermissionHook
var authRolePermissionBeforeUpsertHooks []AuthRolePermissionHook

var authRolePermissionAfterInsertHooks []AuthRolePermissionHook
var authRolePermissionAfterSelectHooks []AuthRolePermissionHook
var authRolePermissionAfterUpdateHooks []AuthRolePermissionHook
var authRolePermissionAfterDeleteHooks []AuthRolePermissionHook
var authRolePermissionAfterUpsertHooks []AuthRolePermissionHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AuthRolePermission) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authRolePermissionBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AuthRolePermission) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range authRolePermissionBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AuthRolePermission) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range authRolePermissionBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AuthRolePermission) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authRolePermissionBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AuthRolePermission) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authRolePermissionAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AuthRolePermission) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range authRolePermissionAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AuthRolePermission) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range authRolePermissionAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AuthRolePermission) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range authRolePermissionAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AuthRolePermission) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authRolePermissionAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAuthRolePermissionHook registers your hook function for all future operations.
func AddAuthRolePermissionHook(hookPoint boil.HookPoint, authRolePermissionHook AuthRolePermissionHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		authRolePermissionBeforeInsertHooks = append(authRolePermissionBeforeInsertHooks, authRolePermissionHook)
	case boil.BeforeUpdateHook:
		authRolePermissionBeforeUpdateHooks = append(authRolePermissionBeforeUpdateHooks, authRolePermissionHook)
	case boil.BeforeDeleteHook:
		authRolePermissionBeforeDeleteHooks = append(authRolePermissionBeforeDeleteHooks, authRolePermissionHook)
	case boil.BeforeUpsertHook:
		authRolePermissionBeforeUpsertHooks = append(authRolePermissionBeforeUpsertHooks, authRolePermissionHook)
	case boil.AfterInsertHook:
		authRolePermissionAfterInsertHooks = append(authRolePermissionAfterInsertHooks, authRolePermissionHook)
	case boil.AfterSelectHook:
		authRolePermissionAfterSelectHooks = append(authRolePermissionAfterSelectHooks, authRolePermissionHook)
	case boil.AfterUpdateHook:
		authRolePermissionAfterUpdateHooks = append(authRolePermissionAfterUpdateHooks, authRolePermissionHook)
	case boil.AfterDeleteHook:
		authRolePermissionAfterDeleteHooks = append(authRolePermissionAfterDeleteHooks, authRolePermissionHook)
	case boil.AfterUpsertHook:
		authRolePermissionAfterUpsertHooks = append(authRolePermissionAfterUpsertHooks, authRolePermissionHook)
	}
}

// OneP returns a single authRolePermission record from the query, and panics on error.
func (q authRolePermissionQuery) OneP() *AuthRolePermission {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single authRolePermission record from the query.
func (q authRolePermissionQuery) One() (*AuthRolePermission, error) {
	o := &AuthRolePermission{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for auth_role_permission")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all AuthRolePermission records from the query, and panics on error.
func (q authRolePermissionQuery) AllP() AuthRolePermissionSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all AuthRolePermission records from the query.
func (q authRolePermissionQuery) All() (AuthRolePermissionSlice, error) {
	var o AuthRolePermissionSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AuthRolePermission slice")
	}

	if len(authRolePermissionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all AuthRolePermission records in the query, and panics on error.
func (q authRolePermissionQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all AuthRolePermission records in the query.
func (q authRolePermissionQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count auth_role_permission rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q authRolePermissionQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q authRolePermissionQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if auth_role_permission exists")
	}

	return count > 0, nil
}

// AuthRoleG pointed to by the foreign key.
func (o *AuthRolePermission) AuthRoleG(mods ...qm.QueryMod) authRoleQuery {
	return o.AuthRole(boil.GetDB(), mods...)
}

// AuthRole pointed to by the foreign key.
func (o *AuthRolePermission) AuthRole(exec boil.Executor, mods ...qm.QueryMod) authRoleQuery {
	queryMods := []qm.QueryMod{
		qm.Where("auth_role_id=$1", o.AuthRoleID),
	}

	queryMods = append(queryMods, mods...)

	query := AuthRoles(exec, queryMods...)
	queries.SetFrom(query.Query, "\"auth_role\"")

	return query
}

// AuthPermissionG pointed to by the foreign key.
func (o *AuthRolePermission) AuthPermissionG(mods ...qm.QueryMod) authPermissionQuery {
	return o.AuthPermission(boil.GetDB(), mods...)
}

// AuthPermission pointed to by the foreign key.
func (o *AuthRolePermission) AuthPermission(exec boil.Executor, mods ...qm.QueryMod) authPermissionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("auth_permission_id=$1", o.AuthPermissionID),
	}

	queryMods = append(queryMods, mods...)

	query := AuthPermissions(exec, queryMods...)
	queries.SetFrom(query.Query, "\"auth_permission\"")

	return query
}

// LoadAuthRole allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (authRolePermissionL) LoadAuthRole(e boil.Executor, singular bool, maybeAuthRolePermission interface{}) error {
	var slice []*AuthRolePermission
	var object *AuthRolePermission

	count := 1
	if singular {
		object = maybeAuthRolePermission.(*AuthRolePermission)
	} else {
		slice = *maybeAuthRolePermission.(*AuthRolePermissionSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &authRolePermissionR{}
		args[0] = object.AuthRoleID
	} else {
		for i, obj := range slice {
			obj.R = &authRolePermissionR{}
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

	if len(authRolePermissionAfterSelectHooks) != 0 {
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

// LoadAuthPermission allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (authRolePermissionL) LoadAuthPermission(e boil.Executor, singular bool, maybeAuthRolePermission interface{}) error {
	var slice []*AuthRolePermission
	var object *AuthRolePermission

	count := 1
	if singular {
		object = maybeAuthRolePermission.(*AuthRolePermission)
	} else {
		slice = *maybeAuthRolePermission.(*AuthRolePermissionSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &authRolePermissionR{}
		args[0] = object.AuthPermissionID
	} else {
		for i, obj := range slice {
			obj.R = &authRolePermissionR{}
			args[i] = obj.AuthPermissionID
		}
	}

	query := fmt.Sprintf(
		"select * from \"auth_permission\" where \"auth_permission_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load AuthPermission")
	}
	defer results.Close()

	var resultSlice []*AuthPermission
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice AuthPermission")
	}

	if len(authRolePermissionAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.AuthPermission = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.AuthPermissionID == foreign.AuthPermissionID {
				local.R.AuthPermission = foreign
				break
			}
		}
	}

	return nil
}

// SetAuthRole of the auth_role_permission to the related item.
// Sets o.R.AuthRole to related.
// Adds o to related.R.AuthRolePermission.
func (o *AuthRolePermission) SetAuthRole(exec boil.Executor, insert bool, related *AuthRole) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"auth_role_permission\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"auth_role_id"}),
		strmangle.WhereClause("\"", "\"", 2, authRolePermissionPrimaryKeyColumns),
	)
	values := []interface{}{related.AuthRoleID, o.AuthRolePermissionID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.AuthRoleID = related.AuthRoleID

	if o.R == nil {
		o.R = &authRolePermissionR{
			AuthRole: related,
		}
	} else {
		o.R.AuthRole = related
	}

	if related.R == nil {
		related.R = &authRoleR{
			AuthRolePermission: o,
		}
	} else {
		related.R.AuthRolePermission = o
	}

	return nil
}

// SetAuthPermission of the auth_role_permission to the related item.
// Sets o.R.AuthPermission to related.
// Adds o to related.R.AuthRolePermission.
func (o *AuthRolePermission) SetAuthPermission(exec boil.Executor, insert bool, related *AuthPermission) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"auth_role_permission\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"auth_permission_id"}),
		strmangle.WhereClause("\"", "\"", 2, authRolePermissionPrimaryKeyColumns),
	)
	values := []interface{}{related.AuthPermissionID, o.AuthRolePermissionID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.AuthPermissionID = related.AuthPermissionID

	if o.R == nil {
		o.R = &authRolePermissionR{
			AuthPermission: related,
		}
	} else {
		o.R.AuthPermission = related
	}

	if related.R == nil {
		related.R = &authPermissionR{
			AuthRolePermission: o,
		}
	} else {
		related.R.AuthRolePermission = o
	}

	return nil
}

// AuthRolePermissionsG retrieves all records.
func AuthRolePermissionsG(mods ...qm.QueryMod) authRolePermissionQuery {
	return AuthRolePermissions(boil.GetDB(), mods...)
}

// AuthRolePermissions retrieves all the records using an executor.
func AuthRolePermissions(exec boil.Executor, mods ...qm.QueryMod) authRolePermissionQuery {
	mods = append(mods, qm.From("\"auth_role_permission\""))
	return authRolePermissionQuery{NewQuery(exec, mods...)}
}

// FindAuthRolePermissionG retrieves a single record by ID.
func FindAuthRolePermissionG(authRolePermissionID int, selectCols ...string) (*AuthRolePermission, error) {
	return FindAuthRolePermission(boil.GetDB(), authRolePermissionID, selectCols...)
}

// FindAuthRolePermissionGP retrieves a single record by ID, and panics on error.
func FindAuthRolePermissionGP(authRolePermissionID int, selectCols ...string) *AuthRolePermission {
	retobj, err := FindAuthRolePermission(boil.GetDB(), authRolePermissionID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindAuthRolePermission retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAuthRolePermission(exec boil.Executor, authRolePermissionID int, selectCols ...string) (*AuthRolePermission, error) {
	authRolePermissionObj := &AuthRolePermission{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"auth_role_permission\" where \"auth_role_permission_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, authRolePermissionID)

	err := q.Bind(authRolePermissionObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from auth_role_permission")
	}

	return authRolePermissionObj, nil
}

// FindAuthRolePermissionP retrieves a single record by ID with an executor, and panics on error.
func FindAuthRolePermissionP(exec boil.Executor, authRolePermissionID int, selectCols ...string) *AuthRolePermission {
	retobj, err := FindAuthRolePermission(exec, authRolePermissionID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *AuthRolePermission) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *AuthRolePermission) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *AuthRolePermission) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *AuthRolePermission) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no auth_role_permission provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(authRolePermissionColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	authRolePermissionInsertCacheMut.RLock()
	cache, cached := authRolePermissionInsertCache[key]
	authRolePermissionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			authRolePermissionColumns,
			authRolePermissionColumnsWithDefault,
			authRolePermissionColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(authRolePermissionType, authRolePermissionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(authRolePermissionType, authRolePermissionMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"auth_role_permission\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into auth_role_permission")
	}

	if !cached {
		authRolePermissionInsertCacheMut.Lock()
		authRolePermissionInsertCache[key] = cache
		authRolePermissionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single AuthRolePermission record. See Update for
// whitelist behavior description.
func (o *AuthRolePermission) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single AuthRolePermission record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *AuthRolePermission) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the AuthRolePermission, and panics on error.
// See Update for whitelist behavior description.
func (o *AuthRolePermission) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the AuthRolePermission.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *AuthRolePermission) Update(exec boil.Executor, whitelist ...string) error {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt.Time = currTime
	o.UpdatedAt.Valid = true

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	authRolePermissionUpdateCacheMut.RLock()
	cache, cached := authRolePermissionUpdateCache[key]
	authRolePermissionUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(authRolePermissionColumns, authRolePermissionPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update auth_role_permission, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"auth_role_permission\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, authRolePermissionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(authRolePermissionType, authRolePermissionMapping, append(wl, authRolePermissionPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update auth_role_permission row")
	}

	if !cached {
		authRolePermissionUpdateCacheMut.Lock()
		authRolePermissionUpdateCache[key] = cache
		authRolePermissionUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q authRolePermissionQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q authRolePermissionQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for auth_role_permission")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AuthRolePermissionSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o AuthRolePermissionSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o AuthRolePermissionSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AuthRolePermissionSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authRolePermissionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"auth_role_permission\" SET %s WHERE (\"auth_role_permission_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(authRolePermissionPrimaryKeyColumns), len(colNames)+1, len(authRolePermissionPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in authRolePermission slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *AuthRolePermission) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *AuthRolePermission) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *AuthRolePermission) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *AuthRolePermission) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no auth_role_permission provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(authRolePermissionColumnsWithDefault, o)

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

	authRolePermissionUpsertCacheMut.RLock()
	cache, cached := authRolePermissionUpsertCache[key]
	authRolePermissionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			authRolePermissionColumns,
			authRolePermissionColumnsWithDefault,
			authRolePermissionColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			authRolePermissionColumns,
			authRolePermissionPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert auth_role_permission, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(authRolePermissionPrimaryKeyColumns))
			copy(conflict, authRolePermissionPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"auth_role_permission\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(authRolePermissionType, authRolePermissionMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(authRolePermissionType, authRolePermissionMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for auth_role_permission")
	}

	if !cached {
		authRolePermissionUpsertCacheMut.Lock()
		authRolePermissionUpsertCache[key] = cache
		authRolePermissionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single AuthRolePermission record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *AuthRolePermission) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single AuthRolePermission record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *AuthRolePermission) DeleteG() error {
	if o == nil {
		return errors.New("models: no AuthRolePermission provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single AuthRolePermission record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *AuthRolePermission) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single AuthRolePermission record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AuthRolePermission) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no AuthRolePermission provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), authRolePermissionPrimaryKeyMapping)
	sql := "DELETE FROM \"auth_role_permission\" WHERE \"auth_role_permission_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from auth_role_permission")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q authRolePermissionQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q authRolePermissionQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no authRolePermissionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from auth_role_permission")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o AuthRolePermissionSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o AuthRolePermissionSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no AuthRolePermission slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o AuthRolePermissionSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AuthRolePermissionSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no AuthRolePermission slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(authRolePermissionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authRolePermissionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"auth_role_permission\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, authRolePermissionPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(authRolePermissionPrimaryKeyColumns), 1, len(authRolePermissionPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from authRolePermission slice")
	}

	if len(authRolePermissionAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *AuthRolePermission) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *AuthRolePermission) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *AuthRolePermission) ReloadG() error {
	if o == nil {
		return errors.New("models: no AuthRolePermission provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *AuthRolePermission) Reload(exec boil.Executor) error {
	ret, err := FindAuthRolePermission(exec, o.AuthRolePermissionID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AuthRolePermissionSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AuthRolePermissionSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AuthRolePermissionSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty AuthRolePermissionSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AuthRolePermissionSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	authRolePermissions := AuthRolePermissionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authRolePermissionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"auth_role_permission\".* FROM \"auth_role_permission\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, authRolePermissionPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(authRolePermissionPrimaryKeyColumns), 1, len(authRolePermissionPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&authRolePermissions)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AuthRolePermissionSlice")
	}

	*o = authRolePermissions

	return nil
}

// AuthRolePermissionExists checks if the AuthRolePermission row exists.
func AuthRolePermissionExists(exec boil.Executor, authRolePermissionID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"auth_role_permission\" where \"auth_role_permission_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, authRolePermissionID)
	}

	row := exec.QueryRow(sql, authRolePermissionID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if auth_role_permission exists")
	}

	return exists, nil
}

// AuthRolePermissionExistsG checks if the AuthRolePermission row exists.
func AuthRolePermissionExistsG(authRolePermissionID int) (bool, error) {
	return AuthRolePermissionExists(boil.GetDB(), authRolePermissionID)
}

// AuthRolePermissionExistsGP checks if the AuthRolePermission row exists. Panics on error.
func AuthRolePermissionExistsGP(authRolePermissionID int) bool {
	e, err := AuthRolePermissionExists(boil.GetDB(), authRolePermissionID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// AuthRolePermissionExistsP checks if the AuthRolePermission row exists. Panics on error.
func AuthRolePermissionExistsP(exec boil.Executor, authRolePermissionID int) bool {
	e, err := AuthRolePermissionExists(exec, authRolePermissionID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
