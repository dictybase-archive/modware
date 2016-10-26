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

// AuthUserProvider is an object representing the database table.
type AuthUserProvider struct {
	AuthUserProviderID int       `boil:"auth_user_provider_id" json:"auth_user_provider_id" toml:"auth_user_provider_id" yaml:"auth_user_provider_id"`
	Name               string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Email              string    `boil:"email" json:"email" toml:"email" yaml:"email"`
	AuthProviderID     int       `boil:"auth_provider_id" json:"auth_provider_id" toml:"auth_provider_id" yaml:"auth_provider_id"`
	CreatedAt          null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt          null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *authUserProviderR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L authUserProviderL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// authUserProviderR is where relationships are stored.
type authUserProviderR struct {
	AuthProvider *AuthProvider
}

// authUserProviderL is where Load methods for each relationship are stored.
type authUserProviderL struct{}

var (
	authUserProviderColumns               = []string{"auth_user_provider_id", "name", "email", "auth_provider_id", "created_at", "updated_at"}
	authUserProviderColumnsWithoutDefault = []string{"name", "email", "auth_provider_id"}
	authUserProviderColumnsWithDefault    = []string{"auth_user_provider_id", "created_at", "updated_at"}
	authUserProviderPrimaryKeyColumns     = []string{"auth_user_provider_id"}
)

type (
	// AuthUserProviderSlice is an alias for a slice of pointers to AuthUserProvider.
	// This should generally be used opposed to []AuthUserProvider.
	AuthUserProviderSlice []*AuthUserProvider
	// AuthUserProviderHook is the signature for custom AuthUserProvider hook methods
	AuthUserProviderHook func(boil.Executor, *AuthUserProvider) error

	authUserProviderQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	authUserProviderType                 = reflect.TypeOf(&AuthUserProvider{})
	authUserProviderMapping              = queries.MakeStructMapping(authUserProviderType)
	authUserProviderPrimaryKeyMapping, _ = queries.BindMapping(authUserProviderType, authUserProviderMapping, authUserProviderPrimaryKeyColumns)
	authUserProviderInsertCacheMut       sync.RWMutex
	authUserProviderInsertCache          = make(map[string]insertCache)
	authUserProviderUpdateCacheMut       sync.RWMutex
	authUserProviderUpdateCache          = make(map[string]updateCache)
	authUserProviderUpsertCacheMut       sync.RWMutex
	authUserProviderUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var authUserProviderBeforeInsertHooks []AuthUserProviderHook
var authUserProviderBeforeUpdateHooks []AuthUserProviderHook
var authUserProviderBeforeDeleteHooks []AuthUserProviderHook
var authUserProviderBeforeUpsertHooks []AuthUserProviderHook

var authUserProviderAfterInsertHooks []AuthUserProviderHook
var authUserProviderAfterSelectHooks []AuthUserProviderHook
var authUserProviderAfterUpdateHooks []AuthUserProviderHook
var authUserProviderAfterDeleteHooks []AuthUserProviderHook
var authUserProviderAfterUpsertHooks []AuthUserProviderHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AuthUserProvider) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserProviderBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AuthUserProvider) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserProviderBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AuthUserProvider) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserProviderBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AuthUserProvider) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserProviderBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AuthUserProvider) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserProviderAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AuthUserProvider) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserProviderAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AuthUserProvider) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserProviderAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AuthUserProvider) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserProviderAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AuthUserProvider) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserProviderAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAuthUserProviderHook registers your hook function for all future operations.
func AddAuthUserProviderHook(hookPoint boil.HookPoint, authUserProviderHook AuthUserProviderHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		authUserProviderBeforeInsertHooks = append(authUserProviderBeforeInsertHooks, authUserProviderHook)
	case boil.BeforeUpdateHook:
		authUserProviderBeforeUpdateHooks = append(authUserProviderBeforeUpdateHooks, authUserProviderHook)
	case boil.BeforeDeleteHook:
		authUserProviderBeforeDeleteHooks = append(authUserProviderBeforeDeleteHooks, authUserProviderHook)
	case boil.BeforeUpsertHook:
		authUserProviderBeforeUpsertHooks = append(authUserProviderBeforeUpsertHooks, authUserProviderHook)
	case boil.AfterInsertHook:
		authUserProviderAfterInsertHooks = append(authUserProviderAfterInsertHooks, authUserProviderHook)
	case boil.AfterSelectHook:
		authUserProviderAfterSelectHooks = append(authUserProviderAfterSelectHooks, authUserProviderHook)
	case boil.AfterUpdateHook:
		authUserProviderAfterUpdateHooks = append(authUserProviderAfterUpdateHooks, authUserProviderHook)
	case boil.AfterDeleteHook:
		authUserProviderAfterDeleteHooks = append(authUserProviderAfterDeleteHooks, authUserProviderHook)
	case boil.AfterUpsertHook:
		authUserProviderAfterUpsertHooks = append(authUserProviderAfterUpsertHooks, authUserProviderHook)
	}
}

// OneP returns a single authUserProvider record from the query, and panics on error.
func (q authUserProviderQuery) OneP() *AuthUserProvider {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single authUserProvider record from the query.
func (q authUserProviderQuery) One() (*AuthUserProvider, error) {
	o := &AuthUserProvider{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for auth_user_provider")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all AuthUserProvider records from the query, and panics on error.
func (q authUserProviderQuery) AllP() AuthUserProviderSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all AuthUserProvider records from the query.
func (q authUserProviderQuery) All() (AuthUserProviderSlice, error) {
	var o AuthUserProviderSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to AuthUserProvider slice")
	}

	if len(authUserProviderAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all AuthUserProvider records in the query, and panics on error.
func (q authUserProviderQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all AuthUserProvider records in the query.
func (q authUserProviderQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count auth_user_provider rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q authUserProviderQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q authUserProviderQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if auth_user_provider exists")
	}

	return count > 0, nil
}

// AuthProviderG pointed to by the foreign key.
func (o *AuthUserProvider) AuthProviderG(mods ...qm.QueryMod) authProviderQuery {
	return o.AuthProvider(boil.GetDB(), mods...)
}

// AuthProvider pointed to by the foreign key.
func (o *AuthUserProvider) AuthProvider(exec boil.Executor, mods ...qm.QueryMod) authProviderQuery {
	queryMods := []qm.QueryMod{
		qm.Where("auth_provider_id=$1", o.AuthProviderID),
	}

	queryMods = append(queryMods, mods...)

	query := AuthProviders(exec, queryMods...)
	queries.SetFrom(query.Query, "\"auth_provider\"")

	return query
}

// LoadAuthProvider allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (authUserProviderL) LoadAuthProvider(e boil.Executor, singular bool, maybeAuthUserProvider interface{}) error {
	var slice []*AuthUserProvider
	var object *AuthUserProvider

	count := 1
	if singular {
		object = maybeAuthUserProvider.(*AuthUserProvider)
	} else {
		slice = *maybeAuthUserProvider.(*AuthUserProviderSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &authUserProviderR{}
		args[0] = object.AuthProviderID
	} else {
		for i, obj := range slice {
			obj.R = &authUserProviderR{}
			args[i] = obj.AuthProviderID
		}
	}

	query := fmt.Sprintf(
		"select * from \"auth_provider\" where \"auth_provider_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load AuthProvider")
	}
	defer results.Close()

	var resultSlice []*AuthProvider
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice AuthProvider")
	}

	if len(authUserProviderAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.AuthProvider = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.AuthProviderID == foreign.AuthProviderID {
				local.R.AuthProvider = foreign
				break
			}
		}
	}

	return nil
}

// SetAuthProvider of the auth_user_provider to the related item.
// Sets o.R.AuthProvider to related.
// Adds o to related.R.AuthUserProviders.
func (o *AuthUserProvider) SetAuthProvider(exec boil.Executor, insert bool, related *AuthProvider) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"auth_user_provider\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"auth_provider_id"}),
		strmangle.WhereClause("\"", "\"", 2, authUserProviderPrimaryKeyColumns),
	)
	values := []interface{}{related.AuthProviderID, o.AuthUserProviderID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.AuthProviderID = related.AuthProviderID

	if o.R == nil {
		o.R = &authUserProviderR{
			AuthProvider: related,
		}
	} else {
		o.R.AuthProvider = related
	}

	if related.R == nil {
		related.R = &authProviderR{
			AuthUserProviders: AuthUserProviderSlice{o},
		}
	} else {
		related.R.AuthUserProviders = append(related.R.AuthUserProviders, o)
	}

	return nil
}

// AuthUserProvidersG retrieves all records.
func AuthUserProvidersG(mods ...qm.QueryMod) authUserProviderQuery {
	return AuthUserProviders(boil.GetDB(), mods...)
}

// AuthUserProviders retrieves all the records using an executor.
func AuthUserProviders(exec boil.Executor, mods ...qm.QueryMod) authUserProviderQuery {
	mods = append(mods, qm.From("\"auth_user_provider\""))
	return authUserProviderQuery{NewQuery(exec, mods...)}
}

// FindAuthUserProviderG retrieves a single record by ID.
func FindAuthUserProviderG(authUserProviderID int, selectCols ...string) (*AuthUserProvider, error) {
	return FindAuthUserProvider(boil.GetDB(), authUserProviderID, selectCols...)
}

// FindAuthUserProviderGP retrieves a single record by ID, and panics on error.
func FindAuthUserProviderGP(authUserProviderID int, selectCols ...string) *AuthUserProvider {
	retobj, err := FindAuthUserProvider(boil.GetDB(), authUserProviderID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindAuthUserProvider retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAuthUserProvider(exec boil.Executor, authUserProviderID int, selectCols ...string) (*AuthUserProvider, error) {
	authUserProviderObj := &AuthUserProvider{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"auth_user_provider\" where \"auth_user_provider_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, authUserProviderID)

	err := q.Bind(authUserProviderObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from auth_user_provider")
	}

	return authUserProviderObj, nil
}

// FindAuthUserProviderP retrieves a single record by ID with an executor, and panics on error.
func FindAuthUserProviderP(exec boil.Executor, authUserProviderID int, selectCols ...string) *AuthUserProvider {
	retobj, err := FindAuthUserProvider(exec, authUserProviderID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *AuthUserProvider) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *AuthUserProvider) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *AuthUserProvider) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *AuthUserProvider) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no auth_user_provider provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(authUserProviderColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	authUserProviderInsertCacheMut.RLock()
	cache, cached := authUserProviderInsertCache[key]
	authUserProviderInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			authUserProviderColumns,
			authUserProviderColumnsWithDefault,
			authUserProviderColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(authUserProviderType, authUserProviderMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(authUserProviderType, authUserProviderMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"auth_user_provider\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into auth_user_provider")
	}

	if !cached {
		authUserProviderInsertCacheMut.Lock()
		authUserProviderInsertCache[key] = cache
		authUserProviderInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single AuthUserProvider record. See Update for
// whitelist behavior description.
func (o *AuthUserProvider) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single AuthUserProvider record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *AuthUserProvider) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the AuthUserProvider, and panics on error.
// See Update for whitelist behavior description.
func (o *AuthUserProvider) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the AuthUserProvider.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *AuthUserProvider) Update(exec boil.Executor, whitelist ...string) error {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt.Time = currTime
	o.UpdatedAt.Valid = true

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	authUserProviderUpdateCacheMut.RLock()
	cache, cached := authUserProviderUpdateCache[key]
	authUserProviderUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(authUserProviderColumns, authUserProviderPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update auth_user_provider, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"auth_user_provider\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, authUserProviderPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(authUserProviderType, authUserProviderMapping, append(wl, authUserProviderPrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update auth_user_provider row")
	}

	if !cached {
		authUserProviderUpdateCacheMut.Lock()
		authUserProviderUpdateCache[key] = cache
		authUserProviderUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q authUserProviderQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q authUserProviderQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for auth_user_provider")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AuthUserProviderSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o AuthUserProviderSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o AuthUserProviderSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AuthUserProviderSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authUserProviderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"auth_user_provider\" SET %s WHERE (\"auth_user_provider_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(authUserProviderPrimaryKeyColumns), len(colNames)+1, len(authUserProviderPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in authUserProvider slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *AuthUserProvider) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *AuthUserProvider) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *AuthUserProvider) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *AuthUserProvider) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no auth_user_provider provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(authUserProviderColumnsWithDefault, o)

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

	authUserProviderUpsertCacheMut.RLock()
	cache, cached := authUserProviderUpsertCache[key]
	authUserProviderUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			authUserProviderColumns,
			authUserProviderColumnsWithDefault,
			authUserProviderColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			authUserProviderColumns,
			authUserProviderPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert auth_user_provider, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(authUserProviderPrimaryKeyColumns))
			copy(conflict, authUserProviderPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"auth_user_provider\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(authUserProviderType, authUserProviderMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(authUserProviderType, authUserProviderMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for auth_user_provider")
	}

	if !cached {
		authUserProviderUpsertCacheMut.Lock()
		authUserProviderUpsertCache[key] = cache
		authUserProviderUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single AuthUserProvider record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *AuthUserProvider) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single AuthUserProvider record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *AuthUserProvider) DeleteG() error {
	if o == nil {
		return errors.New("chado: no AuthUserProvider provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single AuthUserProvider record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *AuthUserProvider) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single AuthUserProvider record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AuthUserProvider) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no AuthUserProvider provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), authUserProviderPrimaryKeyMapping)
	sql := "DELETE FROM \"auth_user_provider\" WHERE \"auth_user_provider_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from auth_user_provider")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q authUserProviderQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q authUserProviderQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no authUserProviderQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from auth_user_provider")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o AuthUserProviderSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o AuthUserProviderSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no AuthUserProvider slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o AuthUserProviderSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AuthUserProviderSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no AuthUserProvider slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(authUserProviderBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authUserProviderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"auth_user_provider\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, authUserProviderPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(authUserProviderPrimaryKeyColumns), 1, len(authUserProviderPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from authUserProvider slice")
	}

	if len(authUserProviderAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *AuthUserProvider) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *AuthUserProvider) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *AuthUserProvider) ReloadG() error {
	if o == nil {
		return errors.New("chado: no AuthUserProvider provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *AuthUserProvider) Reload(exec boil.Executor) error {
	ret, err := FindAuthUserProvider(exec, o.AuthUserProviderID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AuthUserProviderSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AuthUserProviderSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AuthUserProviderSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty AuthUserProviderSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AuthUserProviderSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	authUserProviders := AuthUserProviderSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authUserProviderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"auth_user_provider\".* FROM \"auth_user_provider\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, authUserProviderPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(authUserProviderPrimaryKeyColumns), 1, len(authUserProviderPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&authUserProviders)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in AuthUserProviderSlice")
	}

	*o = authUserProviders

	return nil
}

// AuthUserProviderExists checks if the AuthUserProvider row exists.
func AuthUserProviderExists(exec boil.Executor, authUserProviderID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"auth_user_provider\" where \"auth_user_provider_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, authUserProviderID)
	}

	row := exec.QueryRow(sql, authUserProviderID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if auth_user_provider exists")
	}

	return exists, nil
}

// AuthUserProviderExistsG checks if the AuthUserProvider row exists.
func AuthUserProviderExistsG(authUserProviderID int) (bool, error) {
	return AuthUserProviderExists(boil.GetDB(), authUserProviderID)
}

// AuthUserProviderExistsGP checks if the AuthUserProvider row exists. Panics on error.
func AuthUserProviderExistsGP(authUserProviderID int) bool {
	e, err := AuthUserProviderExists(boil.GetDB(), authUserProviderID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// AuthUserProviderExistsP checks if the AuthUserProvider row exists. Panics on error.
func AuthUserProviderExistsP(exec boil.Executor, authUserProviderID int) bool {
	e, err := AuthUserProviderExists(exec, authUserProviderID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
