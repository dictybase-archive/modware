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
)

// AuthProvider is an object representing the database table.
type AuthProvider struct {
	AuthProviderID int    `boil:"auth_provider_id" json:"auth_provider_id" toml:"auth_provider_id" yaml:"auth_provider_id"`
	Name           string `boil:"name" json:"name" toml:"name" yaml:"name"`

	R *authProviderR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L authProviderL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// authProviderR is where relationships are stored.
type authProviderR struct {
	AuthUserProviders AuthUserProviderSlice
}

// authProviderL is where Load methods for each relationship are stored.
type authProviderL struct{}

var (
	authProviderColumns               = []string{"auth_provider_id", "name"}
	authProviderColumnsWithoutDefault = []string{"name"}
	authProviderColumnsWithDefault    = []string{"auth_provider_id"}
	authProviderPrimaryKeyColumns     = []string{"auth_provider_id"}
)

type (
	// AuthProviderSlice is an alias for a slice of pointers to AuthProvider.
	// This should generally be used opposed to []AuthProvider.
	AuthProviderSlice []*AuthProvider
	// AuthProviderHook is the signature for custom AuthProvider hook methods
	AuthProviderHook func(boil.Executor, *AuthProvider) error

	authProviderQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	authProviderType                 = reflect.TypeOf(&AuthProvider{})
	authProviderMapping              = queries.MakeStructMapping(authProviderType)
	authProviderPrimaryKeyMapping, _ = queries.BindMapping(authProviderType, authProviderMapping, authProviderPrimaryKeyColumns)
	authProviderInsertCacheMut       sync.RWMutex
	authProviderInsertCache          = make(map[string]insertCache)
	authProviderUpdateCacheMut       sync.RWMutex
	authProviderUpdateCache          = make(map[string]updateCache)
	authProviderUpsertCacheMut       sync.RWMutex
	authProviderUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var authProviderBeforeInsertHooks []AuthProviderHook
var authProviderBeforeUpdateHooks []AuthProviderHook
var authProviderBeforeDeleteHooks []AuthProviderHook
var authProviderBeforeUpsertHooks []AuthProviderHook

var authProviderAfterInsertHooks []AuthProviderHook
var authProviderAfterSelectHooks []AuthProviderHook
var authProviderAfterUpdateHooks []AuthProviderHook
var authProviderAfterDeleteHooks []AuthProviderHook
var authProviderAfterUpsertHooks []AuthProviderHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AuthProvider) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authProviderBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AuthProvider) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range authProviderBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AuthProvider) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range authProviderBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AuthProvider) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authProviderBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AuthProvider) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authProviderAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AuthProvider) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range authProviderAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AuthProvider) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range authProviderAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AuthProvider) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range authProviderAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AuthProvider) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authProviderAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAuthProviderHook registers your hook function for all future operations.
func AddAuthProviderHook(hookPoint boil.HookPoint, authProviderHook AuthProviderHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		authProviderBeforeInsertHooks = append(authProviderBeforeInsertHooks, authProviderHook)
	case boil.BeforeUpdateHook:
		authProviderBeforeUpdateHooks = append(authProviderBeforeUpdateHooks, authProviderHook)
	case boil.BeforeDeleteHook:
		authProviderBeforeDeleteHooks = append(authProviderBeforeDeleteHooks, authProviderHook)
	case boil.BeforeUpsertHook:
		authProviderBeforeUpsertHooks = append(authProviderBeforeUpsertHooks, authProviderHook)
	case boil.AfterInsertHook:
		authProviderAfterInsertHooks = append(authProviderAfterInsertHooks, authProviderHook)
	case boil.AfterSelectHook:
		authProviderAfterSelectHooks = append(authProviderAfterSelectHooks, authProviderHook)
	case boil.AfterUpdateHook:
		authProviderAfterUpdateHooks = append(authProviderAfterUpdateHooks, authProviderHook)
	case boil.AfterDeleteHook:
		authProviderAfterDeleteHooks = append(authProviderAfterDeleteHooks, authProviderHook)
	case boil.AfterUpsertHook:
		authProviderAfterUpsertHooks = append(authProviderAfterUpsertHooks, authProviderHook)
	}
}

// OneP returns a single authProvider record from the query, and panics on error.
func (q authProviderQuery) OneP() *AuthProvider {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single authProvider record from the query.
func (q authProviderQuery) One() (*AuthProvider, error) {
	o := &AuthProvider{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for auth_provider")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all AuthProvider records from the query, and panics on error.
func (q authProviderQuery) AllP() AuthProviderSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all AuthProvider records from the query.
func (q authProviderQuery) All() (AuthProviderSlice, error) {
	var o AuthProviderSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to AuthProvider slice")
	}

	if len(authProviderAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all AuthProvider records in the query, and panics on error.
func (q authProviderQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all AuthProvider records in the query.
func (q authProviderQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count auth_provider rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q authProviderQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q authProviderQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if auth_provider exists")
	}

	return count > 0, nil
}

// AuthUserProvidersG retrieves all the auth_user_provider's auth user provider.
func (o *AuthProvider) AuthUserProvidersG(mods ...qm.QueryMod) authUserProviderQuery {
	return o.AuthUserProviders(boil.GetDB(), mods...)
}

// AuthUserProviders retrieves all the auth_user_provider's auth user provider with an executor.
func (o *AuthProvider) AuthUserProviders(exec boil.Executor, mods ...qm.QueryMod) authUserProviderQuery {
	queryMods := []qm.QueryMod{
		qm.Select("\"a\".*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"a\".\"auth_provider_id\"=$1", o.AuthProviderID),
	)

	query := AuthUserProviders(exec, queryMods...)
	queries.SetFrom(query.Query, "\"auth_user_provider\" as \"a\"")
	return query
}

// LoadAuthUserProviders allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (authProviderL) LoadAuthUserProviders(e boil.Executor, singular bool, maybeAuthProvider interface{}) error {
	var slice []*AuthProvider
	var object *AuthProvider

	count := 1
	if singular {
		object = maybeAuthProvider.(*AuthProvider)
	} else {
		slice = *maybeAuthProvider.(*AuthProviderSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &authProviderR{}
		args[0] = object.AuthProviderID
	} else {
		for i, obj := range slice {
			obj.R = &authProviderR{}
			args[i] = obj.AuthProviderID
		}
	}

	query := fmt.Sprintf(
		"select * from \"auth_user_provider\" where \"auth_provider_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load auth_user_provider")
	}
	defer results.Close()

	var resultSlice []*AuthUserProvider
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice auth_user_provider")
	}

	if len(authUserProviderAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.AuthUserProviders = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.AuthProviderID == foreign.AuthProviderID {
				local.R.AuthUserProviders = append(local.R.AuthUserProviders, foreign)
				break
			}
		}
	}

	return nil
}

// AddAuthUserProviders adds the given related objects to the existing relationships
// of the auth_provider, optionally inserting them as new records.
// Appends related to o.R.AuthUserProviders.
// Sets related.R.AuthProvider appropriately.
func (o *AuthProvider) AddAuthUserProviders(exec boil.Executor, insert bool, related ...*AuthUserProvider) error {
	var err error
	for _, rel := range related {
		rel.AuthProviderID = o.AuthProviderID
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			if err = rel.Update(exec, "auth_provider_id"); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}
		}
	}

	if o.R == nil {
		o.R = &authProviderR{
			AuthUserProviders: related,
		}
	} else {
		o.R.AuthUserProviders = append(o.R.AuthUserProviders, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &authUserProviderR{
				AuthProvider: o,
			}
		} else {
			rel.R.AuthProvider = o
		}
	}
	return nil
}

// AuthProvidersG retrieves all records.
func AuthProvidersG(mods ...qm.QueryMod) authProviderQuery {
	return AuthProviders(boil.GetDB(), mods...)
}

// AuthProviders retrieves all the records using an executor.
func AuthProviders(exec boil.Executor, mods ...qm.QueryMod) authProviderQuery {
	mods = append(mods, qm.From("\"auth_provider\""))
	return authProviderQuery{NewQuery(exec, mods...)}
}

// FindAuthProviderG retrieves a single record by ID.
func FindAuthProviderG(authProviderID int, selectCols ...string) (*AuthProvider, error) {
	return FindAuthProvider(boil.GetDB(), authProviderID, selectCols...)
}

// FindAuthProviderGP retrieves a single record by ID, and panics on error.
func FindAuthProviderGP(authProviderID int, selectCols ...string) *AuthProvider {
	retobj, err := FindAuthProvider(boil.GetDB(), authProviderID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindAuthProvider retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAuthProvider(exec boil.Executor, authProviderID int, selectCols ...string) (*AuthProvider, error) {
	authProviderObj := &AuthProvider{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"auth_provider\" where \"auth_provider_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, authProviderID)

	err := q.Bind(authProviderObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from auth_provider")
	}

	return authProviderObj, nil
}

// FindAuthProviderP retrieves a single record by ID with an executor, and panics on error.
func FindAuthProviderP(exec boil.Executor, authProviderID int, selectCols ...string) *AuthProvider {
	retobj, err := FindAuthProvider(exec, authProviderID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *AuthProvider) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *AuthProvider) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *AuthProvider) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *AuthProvider) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no auth_provider provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(authProviderColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	authProviderInsertCacheMut.RLock()
	cache, cached := authProviderInsertCache[key]
	authProviderInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			authProviderColumns,
			authProviderColumnsWithDefault,
			authProviderColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(authProviderType, authProviderMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(authProviderType, authProviderMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"auth_provider\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into auth_provider")
	}

	if !cached {
		authProviderInsertCacheMut.Lock()
		authProviderInsertCache[key] = cache
		authProviderInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single AuthProvider record. See Update for
// whitelist behavior description.
func (o *AuthProvider) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single AuthProvider record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *AuthProvider) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the AuthProvider, and panics on error.
// See Update for whitelist behavior description.
func (o *AuthProvider) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the AuthProvider.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *AuthProvider) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	authProviderUpdateCacheMut.RLock()
	cache, cached := authProviderUpdateCache[key]
	authProviderUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(authProviderColumns, authProviderPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update auth_provider, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"auth_provider\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, authProviderPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(authProviderType, authProviderMapping, append(wl, authProviderPrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update auth_provider row")
	}

	if !cached {
		authProviderUpdateCacheMut.Lock()
		authProviderUpdateCache[key] = cache
		authProviderUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q authProviderQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q authProviderQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for auth_provider")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AuthProviderSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o AuthProviderSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o AuthProviderSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AuthProviderSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authProviderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"auth_provider\" SET %s WHERE (\"auth_provider_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(authProviderPrimaryKeyColumns), len(colNames)+1, len(authProviderPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in authProvider slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *AuthProvider) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *AuthProvider) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *AuthProvider) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *AuthProvider) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no auth_provider provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(authProviderColumnsWithDefault, o)

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

	authProviderUpsertCacheMut.RLock()
	cache, cached := authProviderUpsertCache[key]
	authProviderUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			authProviderColumns,
			authProviderColumnsWithDefault,
			authProviderColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			authProviderColumns,
			authProviderPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert auth_provider, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(authProviderPrimaryKeyColumns))
			copy(conflict, authProviderPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"auth_provider\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(authProviderType, authProviderMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(authProviderType, authProviderMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for auth_provider")
	}

	if !cached {
		authProviderUpsertCacheMut.Lock()
		authProviderUpsertCache[key] = cache
		authProviderUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single AuthProvider record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *AuthProvider) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single AuthProvider record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *AuthProvider) DeleteG() error {
	if o == nil {
		return errors.New("chado: no AuthProvider provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single AuthProvider record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *AuthProvider) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single AuthProvider record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AuthProvider) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no AuthProvider provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), authProviderPrimaryKeyMapping)
	sql := "DELETE FROM \"auth_provider\" WHERE \"auth_provider_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from auth_provider")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q authProviderQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q authProviderQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no authProviderQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from auth_provider")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o AuthProviderSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o AuthProviderSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no AuthProvider slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o AuthProviderSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AuthProviderSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no AuthProvider slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(authProviderBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authProviderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"auth_provider\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, authProviderPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(authProviderPrimaryKeyColumns), 1, len(authProviderPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from authProvider slice")
	}

	if len(authProviderAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *AuthProvider) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *AuthProvider) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *AuthProvider) ReloadG() error {
	if o == nil {
		return errors.New("chado: no AuthProvider provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *AuthProvider) Reload(exec boil.Executor) error {
	ret, err := FindAuthProvider(exec, o.AuthProviderID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AuthProviderSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AuthProviderSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AuthProviderSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty AuthProviderSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AuthProviderSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	authProviders := AuthProviderSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authProviderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"auth_provider\".* FROM \"auth_provider\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, authProviderPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(authProviderPrimaryKeyColumns), 1, len(authProviderPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&authProviders)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in AuthProviderSlice")
	}

	*o = authProviders

	return nil
}

// AuthProviderExists checks if the AuthProvider row exists.
func AuthProviderExists(exec boil.Executor, authProviderID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"auth_provider\" where \"auth_provider_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, authProviderID)
	}

	row := exec.QueryRow(sql, authProviderID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if auth_provider exists")
	}

	return exists, nil
}

// AuthProviderExistsG checks if the AuthProvider row exists.
func AuthProviderExistsG(authProviderID int) (bool, error) {
	return AuthProviderExists(boil.GetDB(), authProviderID)
}

// AuthProviderExistsGP checks if the AuthProvider row exists. Panics on error.
func AuthProviderExistsGP(authProviderID int) bool {
	e, err := AuthProviderExists(boil.GetDB(), authProviderID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// AuthProviderExistsP checks if the AuthProvider row exists. Panics on error.
func AuthProviderExistsP(exec boil.Executor, authProviderID int) bool {
	e, err := AuthProviderExists(exec, authProviderID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
