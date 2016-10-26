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

// AuthPermission is an object representing the database table.
type AuthPermission struct {
	AuthPermissionID int         `boil:"auth_permission_id" json:"auth_permission_id" toml:"auth_permission_id" yaml:"auth_permission_id"`
	Permission       string      `boil:"permission" json:"permission" toml:"permission" yaml:"permission"`
	Description      null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	CreatedAt        null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt        null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *authPermissionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L authPermissionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// authPermissionR is where relationships are stored.
type authPermissionR struct {
	AuthRolePermission *AuthRolePermission
}

// authPermissionL is where Load methods for each relationship are stored.
type authPermissionL struct{}

var (
	authPermissionColumns               = []string{"auth_permission_id", "permission", "description", "created_at", "updated_at"}
	authPermissionColumnsWithoutDefault = []string{"permission", "description"}
	authPermissionColumnsWithDefault    = []string{"auth_permission_id", "created_at", "updated_at"}
	authPermissionPrimaryKeyColumns     = []string{"auth_permission_id"}
)

type (
	// AuthPermissionSlice is an alias for a slice of pointers to AuthPermission.
	// This should generally be used opposed to []AuthPermission.
	AuthPermissionSlice []*AuthPermission
	// AuthPermissionHook is the signature for custom AuthPermission hook methods
	AuthPermissionHook func(boil.Executor, *AuthPermission) error

	authPermissionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	authPermissionType                 = reflect.TypeOf(&AuthPermission{})
	authPermissionMapping              = queries.MakeStructMapping(authPermissionType)
	authPermissionPrimaryKeyMapping, _ = queries.BindMapping(authPermissionType, authPermissionMapping, authPermissionPrimaryKeyColumns)
	authPermissionInsertCacheMut       sync.RWMutex
	authPermissionInsertCache          = make(map[string]insertCache)
	authPermissionUpdateCacheMut       sync.RWMutex
	authPermissionUpdateCache          = make(map[string]updateCache)
	authPermissionUpsertCacheMut       sync.RWMutex
	authPermissionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var authPermissionBeforeInsertHooks []AuthPermissionHook
var authPermissionBeforeUpdateHooks []AuthPermissionHook
var authPermissionBeforeDeleteHooks []AuthPermissionHook
var authPermissionBeforeUpsertHooks []AuthPermissionHook

var authPermissionAfterInsertHooks []AuthPermissionHook
var authPermissionAfterSelectHooks []AuthPermissionHook
var authPermissionAfterUpdateHooks []AuthPermissionHook
var authPermissionAfterDeleteHooks []AuthPermissionHook
var authPermissionAfterUpsertHooks []AuthPermissionHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AuthPermission) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authPermissionBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AuthPermission) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range authPermissionBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AuthPermission) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range authPermissionBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AuthPermission) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authPermissionBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AuthPermission) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authPermissionAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AuthPermission) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range authPermissionAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AuthPermission) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range authPermissionAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AuthPermission) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range authPermissionAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AuthPermission) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authPermissionAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAuthPermissionHook registers your hook function for all future operations.
func AddAuthPermissionHook(hookPoint boil.HookPoint, authPermissionHook AuthPermissionHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		authPermissionBeforeInsertHooks = append(authPermissionBeforeInsertHooks, authPermissionHook)
	case boil.BeforeUpdateHook:
		authPermissionBeforeUpdateHooks = append(authPermissionBeforeUpdateHooks, authPermissionHook)
	case boil.BeforeDeleteHook:
		authPermissionBeforeDeleteHooks = append(authPermissionBeforeDeleteHooks, authPermissionHook)
	case boil.BeforeUpsertHook:
		authPermissionBeforeUpsertHooks = append(authPermissionBeforeUpsertHooks, authPermissionHook)
	case boil.AfterInsertHook:
		authPermissionAfterInsertHooks = append(authPermissionAfterInsertHooks, authPermissionHook)
	case boil.AfterSelectHook:
		authPermissionAfterSelectHooks = append(authPermissionAfterSelectHooks, authPermissionHook)
	case boil.AfterUpdateHook:
		authPermissionAfterUpdateHooks = append(authPermissionAfterUpdateHooks, authPermissionHook)
	case boil.AfterDeleteHook:
		authPermissionAfterDeleteHooks = append(authPermissionAfterDeleteHooks, authPermissionHook)
	case boil.AfterUpsertHook:
		authPermissionAfterUpsertHooks = append(authPermissionAfterUpsertHooks, authPermissionHook)
	}
}

// OneP returns a single authPermission record from the query, and panics on error.
func (q authPermissionQuery) OneP() *AuthPermission {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single authPermission record from the query.
func (q authPermissionQuery) One() (*AuthPermission, error) {
	o := &AuthPermission{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for auth_permission")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all AuthPermission records from the query, and panics on error.
func (q authPermissionQuery) AllP() AuthPermissionSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all AuthPermission records from the query.
func (q authPermissionQuery) All() (AuthPermissionSlice, error) {
	var o AuthPermissionSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to AuthPermission slice")
	}

	if len(authPermissionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all AuthPermission records in the query, and panics on error.
func (q authPermissionQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all AuthPermission records in the query.
func (q authPermissionQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count auth_permission rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q authPermissionQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q authPermissionQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if auth_permission exists")
	}

	return count > 0, nil
}

// AuthRolePermissionG pointed to by the foreign key.
func (o *AuthPermission) AuthRolePermissionG(mods ...qm.QueryMod) authRolePermissionQuery {
	return o.AuthRolePermission(boil.GetDB(), mods...)
}

// AuthRolePermission pointed to by the foreign key.
func (o *AuthPermission) AuthRolePermission(exec boil.Executor, mods ...qm.QueryMod) authRolePermissionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("auth_permission_id=$1", o.AuthPermissionID),
	}

	queryMods = append(queryMods, mods...)

	query := AuthRolePermissions(exec, queryMods...)
	queries.SetFrom(query.Query, "\"auth_role_permission\"")

	return query
}

// LoadAuthRolePermission allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (authPermissionL) LoadAuthRolePermission(e boil.Executor, singular bool, maybeAuthPermission interface{}) error {
	var slice []*AuthPermission
	var object *AuthPermission

	count := 1
	if singular {
		object = maybeAuthPermission.(*AuthPermission)
	} else {
		slice = *maybeAuthPermission.(*AuthPermissionSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &authPermissionR{}
		args[0] = object.AuthPermissionID
	} else {
		for i, obj := range slice {
			obj.R = &authPermissionR{}
			args[i] = obj.AuthPermissionID
		}
	}

	query := fmt.Sprintf(
		"select * from \"auth_role_permission\" where \"auth_permission_id\" in (%s)",
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

	if len(authPermissionAfterSelectHooks) != 0 {
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
			if local.AuthPermissionID == foreign.AuthPermissionID {
				local.R.AuthRolePermission = foreign
				break
			}
		}
	}

	return nil
}

// SetAuthRolePermission of the auth_permission to the related item.
// Sets o.R.AuthRolePermission to related.
// Adds o to related.R.AuthPermission.
func (o *AuthPermission) SetAuthRolePermission(exec boil.Executor, insert bool, related *AuthRolePermission) error {
	var err error

	if insert {
		related.AuthPermissionID = o.AuthPermissionID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"auth_role_permission\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"auth_permission_id"}),
			strmangle.WhereClause("\"", "\"", 2, authRolePermissionPrimaryKeyColumns),
		)
		values := []interface{}{o.AuthPermissionID, related.AuthRolePermissionID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.AuthPermissionID = o.AuthPermissionID

	}

	if o.R == nil {
		o.R = &authPermissionR{
			AuthRolePermission: related,
		}
	} else {
		o.R.AuthRolePermission = related
	}

	if related.R == nil {
		related.R = &authRolePermissionR{
			AuthPermission: o,
		}
	} else {
		related.R.AuthPermission = o
	}
	return nil
}

// AuthPermissionsG retrieves all records.
func AuthPermissionsG(mods ...qm.QueryMod) authPermissionQuery {
	return AuthPermissions(boil.GetDB(), mods...)
}

// AuthPermissions retrieves all the records using an executor.
func AuthPermissions(exec boil.Executor, mods ...qm.QueryMod) authPermissionQuery {
	mods = append(mods, qm.From("\"auth_permission\""))
	return authPermissionQuery{NewQuery(exec, mods...)}
}

// FindAuthPermissionG retrieves a single record by ID.
func FindAuthPermissionG(authPermissionID int, selectCols ...string) (*AuthPermission, error) {
	return FindAuthPermission(boil.GetDB(), authPermissionID, selectCols...)
}

// FindAuthPermissionGP retrieves a single record by ID, and panics on error.
func FindAuthPermissionGP(authPermissionID int, selectCols ...string) *AuthPermission {
	retobj, err := FindAuthPermission(boil.GetDB(), authPermissionID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindAuthPermission retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAuthPermission(exec boil.Executor, authPermissionID int, selectCols ...string) (*AuthPermission, error) {
	authPermissionObj := &AuthPermission{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"auth_permission\" where \"auth_permission_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, authPermissionID)

	err := q.Bind(authPermissionObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from auth_permission")
	}

	return authPermissionObj, nil
}

// FindAuthPermissionP retrieves a single record by ID with an executor, and panics on error.
func FindAuthPermissionP(exec boil.Executor, authPermissionID int, selectCols ...string) *AuthPermission {
	retobj, err := FindAuthPermission(exec, authPermissionID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *AuthPermission) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *AuthPermission) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *AuthPermission) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *AuthPermission) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no auth_permission provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(authPermissionColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	authPermissionInsertCacheMut.RLock()
	cache, cached := authPermissionInsertCache[key]
	authPermissionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			authPermissionColumns,
			authPermissionColumnsWithDefault,
			authPermissionColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(authPermissionType, authPermissionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(authPermissionType, authPermissionMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"auth_permission\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into auth_permission")
	}

	if !cached {
		authPermissionInsertCacheMut.Lock()
		authPermissionInsertCache[key] = cache
		authPermissionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single AuthPermission record. See Update for
// whitelist behavior description.
func (o *AuthPermission) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single AuthPermission record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *AuthPermission) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the AuthPermission, and panics on error.
// See Update for whitelist behavior description.
func (o *AuthPermission) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the AuthPermission.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *AuthPermission) Update(exec boil.Executor, whitelist ...string) error {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt.Time = currTime
	o.UpdatedAt.Valid = true

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	authPermissionUpdateCacheMut.RLock()
	cache, cached := authPermissionUpdateCache[key]
	authPermissionUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(authPermissionColumns, authPermissionPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update auth_permission, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"auth_permission\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, authPermissionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(authPermissionType, authPermissionMapping, append(wl, authPermissionPrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update auth_permission row")
	}

	if !cached {
		authPermissionUpdateCacheMut.Lock()
		authPermissionUpdateCache[key] = cache
		authPermissionUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q authPermissionQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q authPermissionQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for auth_permission")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AuthPermissionSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o AuthPermissionSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o AuthPermissionSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AuthPermissionSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authPermissionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"auth_permission\" SET %s WHERE (\"auth_permission_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(authPermissionPrimaryKeyColumns), len(colNames)+1, len(authPermissionPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in authPermission slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *AuthPermission) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *AuthPermission) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *AuthPermission) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *AuthPermission) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no auth_permission provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(authPermissionColumnsWithDefault, o)

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

	authPermissionUpsertCacheMut.RLock()
	cache, cached := authPermissionUpsertCache[key]
	authPermissionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			authPermissionColumns,
			authPermissionColumnsWithDefault,
			authPermissionColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			authPermissionColumns,
			authPermissionPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert auth_permission, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(authPermissionPrimaryKeyColumns))
			copy(conflict, authPermissionPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"auth_permission\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(authPermissionType, authPermissionMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(authPermissionType, authPermissionMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for auth_permission")
	}

	if !cached {
		authPermissionUpsertCacheMut.Lock()
		authPermissionUpsertCache[key] = cache
		authPermissionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single AuthPermission record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *AuthPermission) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single AuthPermission record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *AuthPermission) DeleteG() error {
	if o == nil {
		return errors.New("chado: no AuthPermission provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single AuthPermission record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *AuthPermission) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single AuthPermission record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AuthPermission) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no AuthPermission provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), authPermissionPrimaryKeyMapping)
	sql := "DELETE FROM \"auth_permission\" WHERE \"auth_permission_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from auth_permission")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q authPermissionQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q authPermissionQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no authPermissionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from auth_permission")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o AuthPermissionSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o AuthPermissionSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no AuthPermission slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o AuthPermissionSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AuthPermissionSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no AuthPermission slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(authPermissionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authPermissionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"auth_permission\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, authPermissionPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(authPermissionPrimaryKeyColumns), 1, len(authPermissionPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from authPermission slice")
	}

	if len(authPermissionAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *AuthPermission) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *AuthPermission) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *AuthPermission) ReloadG() error {
	if o == nil {
		return errors.New("chado: no AuthPermission provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *AuthPermission) Reload(exec boil.Executor) error {
	ret, err := FindAuthPermission(exec, o.AuthPermissionID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AuthPermissionSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AuthPermissionSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AuthPermissionSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty AuthPermissionSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AuthPermissionSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	authPermissions := AuthPermissionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authPermissionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"auth_permission\".* FROM \"auth_permission\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, authPermissionPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(authPermissionPrimaryKeyColumns), 1, len(authPermissionPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&authPermissions)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in AuthPermissionSlice")
	}

	*o = authPermissions

	return nil
}

// AuthPermissionExists checks if the AuthPermission row exists.
func AuthPermissionExists(exec boil.Executor, authPermissionID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"auth_permission\" where \"auth_permission_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, authPermissionID)
	}

	row := exec.QueryRow(sql, authPermissionID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if auth_permission exists")
	}

	return exists, nil
}

// AuthPermissionExistsG checks if the AuthPermission row exists.
func AuthPermissionExistsG(authPermissionID int) (bool, error) {
	return AuthPermissionExists(boil.GetDB(), authPermissionID)
}

// AuthPermissionExistsGP checks if the AuthPermission row exists. Panics on error.
func AuthPermissionExistsGP(authPermissionID int) bool {
	e, err := AuthPermissionExists(boil.GetDB(), authPermissionID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// AuthPermissionExistsP checks if the AuthPermission row exists. Panics on error.
func AuthPermissionExistsP(exec boil.Executor, authPermissionID int) bool {
	e, err := AuthPermissionExists(exec, authPermissionID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
