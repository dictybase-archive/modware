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

// AuthUserInfo is an object representing the database table.
type AuthUserInfo struct {
	AuthUserInfoID int         `boil:"auth_user_info_id" json:"auth_user_info_id" toml:"auth_user_info_id" yaml:"auth_user_info_id"`
	Organization   null.String `boil:"organization" json:"organization,omitempty" toml:"organization" yaml:"organization,omitempty"`
	GroupName      null.String `boil:"group_name" json:"group_name,omitempty" toml:"group_name" yaml:"group_name,omitempty"`
	FirstAddress   null.String `boil:"first_address" json:"first_address,omitempty" toml:"first_address" yaml:"first_address,omitempty"`
	SecondAddress  null.String `boil:"second_address" json:"second_address,omitempty" toml:"second_address" yaml:"second_address,omitempty"`
	City           null.String `boil:"city" json:"city,omitempty" toml:"city" yaml:"city,omitempty"`
	State          null.String `boil:"state" json:"state,omitempty" toml:"state" yaml:"state,omitempty"`
	Zipcode        null.String `boil:"zipcode" json:"zipcode,omitempty" toml:"zipcode" yaml:"zipcode,omitempty"`
	Country        null.String `boil:"country" json:"country,omitempty" toml:"country" yaml:"country,omitempty"`
	Phone          null.String `boil:"phone" json:"phone,omitempty" toml:"phone" yaml:"phone,omitempty"`
	AuthUserID     int         `boil:"auth_user_id" json:"auth_user_id" toml:"auth_user_id" yaml:"auth_user_id"`

	R *authUserInfoR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L authUserInfoL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// authUserInfoR is where relationships are stored.
type authUserInfoR struct {
	AuthUser *AuthUser
}

// authUserInfoL is where Load methods for each relationship are stored.
type authUserInfoL struct{}

var (
	authUserInfoColumns               = []string{"auth_user_info_id", "organization", "group_name", "first_address", "second_address", "city", "state", "zipcode", "country", "phone", "auth_user_id"}
	authUserInfoColumnsWithoutDefault = []string{"organization", "group_name", "first_address", "second_address", "city", "state", "zipcode", "country", "phone", "auth_user_id"}
	authUserInfoColumnsWithDefault    = []string{"auth_user_info_id"}
	authUserInfoPrimaryKeyColumns     = []string{"auth_user_info_id"}
)

type (
	// AuthUserInfoSlice is an alias for a slice of pointers to AuthUserInfo.
	// This should generally be used opposed to []AuthUserInfo.
	AuthUserInfoSlice []*AuthUserInfo
	// AuthUserInfoHook is the signature for custom AuthUserInfo hook methods
	AuthUserInfoHook func(boil.Executor, *AuthUserInfo) error

	authUserInfoQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	authUserInfoType                 = reflect.TypeOf(&AuthUserInfo{})
	authUserInfoMapping              = queries.MakeStructMapping(authUserInfoType)
	authUserInfoPrimaryKeyMapping, _ = queries.BindMapping(authUserInfoType, authUserInfoMapping, authUserInfoPrimaryKeyColumns)
	authUserInfoInsertCacheMut       sync.RWMutex
	authUserInfoInsertCache          = make(map[string]insertCache)
	authUserInfoUpdateCacheMut       sync.RWMutex
	authUserInfoUpdateCache          = make(map[string]updateCache)
	authUserInfoUpsertCacheMut       sync.RWMutex
	authUserInfoUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var authUserInfoBeforeInsertHooks []AuthUserInfoHook
var authUserInfoBeforeUpdateHooks []AuthUserInfoHook
var authUserInfoBeforeDeleteHooks []AuthUserInfoHook
var authUserInfoBeforeUpsertHooks []AuthUserInfoHook

var authUserInfoAfterInsertHooks []AuthUserInfoHook
var authUserInfoAfterSelectHooks []AuthUserInfoHook
var authUserInfoAfterUpdateHooks []AuthUserInfoHook
var authUserInfoAfterDeleteHooks []AuthUserInfoHook
var authUserInfoAfterUpsertHooks []AuthUserInfoHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AuthUserInfo) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserInfoBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AuthUserInfo) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserInfoBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AuthUserInfo) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserInfoBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AuthUserInfo) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserInfoBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AuthUserInfo) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserInfoAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AuthUserInfo) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserInfoAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AuthUserInfo) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserInfoAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AuthUserInfo) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserInfoAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AuthUserInfo) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range authUserInfoAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAuthUserInfoHook registers your hook function for all future operations.
func AddAuthUserInfoHook(hookPoint boil.HookPoint, authUserInfoHook AuthUserInfoHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		authUserInfoBeforeInsertHooks = append(authUserInfoBeforeInsertHooks, authUserInfoHook)
	case boil.BeforeUpdateHook:
		authUserInfoBeforeUpdateHooks = append(authUserInfoBeforeUpdateHooks, authUserInfoHook)
	case boil.BeforeDeleteHook:
		authUserInfoBeforeDeleteHooks = append(authUserInfoBeforeDeleteHooks, authUserInfoHook)
	case boil.BeforeUpsertHook:
		authUserInfoBeforeUpsertHooks = append(authUserInfoBeforeUpsertHooks, authUserInfoHook)
	case boil.AfterInsertHook:
		authUserInfoAfterInsertHooks = append(authUserInfoAfterInsertHooks, authUserInfoHook)
	case boil.AfterSelectHook:
		authUserInfoAfterSelectHooks = append(authUserInfoAfterSelectHooks, authUserInfoHook)
	case boil.AfterUpdateHook:
		authUserInfoAfterUpdateHooks = append(authUserInfoAfterUpdateHooks, authUserInfoHook)
	case boil.AfterDeleteHook:
		authUserInfoAfterDeleteHooks = append(authUserInfoAfterDeleteHooks, authUserInfoHook)
	case boil.AfterUpsertHook:
		authUserInfoAfterUpsertHooks = append(authUserInfoAfterUpsertHooks, authUserInfoHook)
	}
}

// OneP returns a single authUserInfo record from the query, and panics on error.
func (q authUserInfoQuery) OneP() *AuthUserInfo {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single authUserInfo record from the query.
func (q authUserInfoQuery) One() (*AuthUserInfo, error) {
	o := &AuthUserInfo{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for auth_user_info")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all AuthUserInfo records from the query, and panics on error.
func (q authUserInfoQuery) AllP() AuthUserInfoSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all AuthUserInfo records from the query.
func (q authUserInfoQuery) All() (AuthUserInfoSlice, error) {
	var o AuthUserInfoSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AuthUserInfo slice")
	}

	if len(authUserInfoAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all AuthUserInfo records in the query, and panics on error.
func (q authUserInfoQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all AuthUserInfo records in the query.
func (q authUserInfoQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count auth_user_info rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q authUserInfoQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q authUserInfoQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if auth_user_info exists")
	}

	return count > 0, nil
}

// AuthUserG pointed to by the foreign key.
func (o *AuthUserInfo) AuthUserG(mods ...qm.QueryMod) authUserQuery {
	return o.AuthUser(boil.GetDB(), mods...)
}

// AuthUser pointed to by the foreign key.
func (o *AuthUserInfo) AuthUser(exec boil.Executor, mods ...qm.QueryMod) authUserQuery {
	queryMods := []qm.QueryMod{
		qm.Where("auth_user_id=$1", o.AuthUserID),
	}

	queryMods = append(queryMods, mods...)

	query := AuthUsers(exec, queryMods...)
	queries.SetFrom(query.Query, "\"auth_user\"")

	return query
}

// LoadAuthUser allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (authUserInfoL) LoadAuthUser(e boil.Executor, singular bool, maybeAuthUserInfo interface{}) error {
	var slice []*AuthUserInfo
	var object *AuthUserInfo

	count := 1
	if singular {
		object = maybeAuthUserInfo.(*AuthUserInfo)
	} else {
		slice = *maybeAuthUserInfo.(*AuthUserInfoSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &authUserInfoR{}
		args[0] = object.AuthUserID
	} else {
		for i, obj := range slice {
			obj.R = &authUserInfoR{}
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

	if len(authUserInfoAfterSelectHooks) != 0 {
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

// SetAuthUser of the auth_user_info to the related item.
// Sets o.R.AuthUser to related.
// Adds o to related.R.AuthUserInfos.
func (o *AuthUserInfo) SetAuthUser(exec boil.Executor, insert bool, related *AuthUser) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"auth_user_info\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"auth_user_id"}),
		strmangle.WhereClause("\"", "\"", 2, authUserInfoPrimaryKeyColumns),
	)
	values := []interface{}{related.AuthUserID, o.AuthUserInfoID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.AuthUserID = related.AuthUserID

	if o.R == nil {
		o.R = &authUserInfoR{
			AuthUser: related,
		}
	} else {
		o.R.AuthUser = related
	}

	if related.R == nil {
		related.R = &authUserR{
			AuthUserInfos: AuthUserInfoSlice{o},
		}
	} else {
		related.R.AuthUserInfos = append(related.R.AuthUserInfos, o)
	}

	return nil
}

// AuthUserInfosG retrieves all records.
func AuthUserInfosG(mods ...qm.QueryMod) authUserInfoQuery {
	return AuthUserInfos(boil.GetDB(), mods...)
}

// AuthUserInfos retrieves all the records using an executor.
func AuthUserInfos(exec boil.Executor, mods ...qm.QueryMod) authUserInfoQuery {
	mods = append(mods, qm.From("\"auth_user_info\""))
	return authUserInfoQuery{NewQuery(exec, mods...)}
}

// FindAuthUserInfoG retrieves a single record by ID.
func FindAuthUserInfoG(authUserInfoID int, selectCols ...string) (*AuthUserInfo, error) {
	return FindAuthUserInfo(boil.GetDB(), authUserInfoID, selectCols...)
}

// FindAuthUserInfoGP retrieves a single record by ID, and panics on error.
func FindAuthUserInfoGP(authUserInfoID int, selectCols ...string) *AuthUserInfo {
	retobj, err := FindAuthUserInfo(boil.GetDB(), authUserInfoID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindAuthUserInfo retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAuthUserInfo(exec boil.Executor, authUserInfoID int, selectCols ...string) (*AuthUserInfo, error) {
	authUserInfoObj := &AuthUserInfo{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"auth_user_info\" where \"auth_user_info_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, authUserInfoID)

	err := q.Bind(authUserInfoObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from auth_user_info")
	}

	return authUserInfoObj, nil
}

// FindAuthUserInfoP retrieves a single record by ID with an executor, and panics on error.
func FindAuthUserInfoP(exec boil.Executor, authUserInfoID int, selectCols ...string) *AuthUserInfo {
	retobj, err := FindAuthUserInfo(exec, authUserInfoID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *AuthUserInfo) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *AuthUserInfo) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *AuthUserInfo) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *AuthUserInfo) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no auth_user_info provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(authUserInfoColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	authUserInfoInsertCacheMut.RLock()
	cache, cached := authUserInfoInsertCache[key]
	authUserInfoInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			authUserInfoColumns,
			authUserInfoColumnsWithDefault,
			authUserInfoColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(authUserInfoType, authUserInfoMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(authUserInfoType, authUserInfoMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"auth_user_info\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "models: unable to insert into auth_user_info")
	}

	if !cached {
		authUserInfoInsertCacheMut.Lock()
		authUserInfoInsertCache[key] = cache
		authUserInfoInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single AuthUserInfo record. See Update for
// whitelist behavior description.
func (o *AuthUserInfo) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single AuthUserInfo record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *AuthUserInfo) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the AuthUserInfo, and panics on error.
// See Update for whitelist behavior description.
func (o *AuthUserInfo) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the AuthUserInfo.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *AuthUserInfo) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	authUserInfoUpdateCacheMut.RLock()
	cache, cached := authUserInfoUpdateCache[key]
	authUserInfoUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(authUserInfoColumns, authUserInfoPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update auth_user_info, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"auth_user_info\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, authUserInfoPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(authUserInfoType, authUserInfoMapping, append(wl, authUserInfoPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update auth_user_info row")
	}

	if !cached {
		authUserInfoUpdateCacheMut.Lock()
		authUserInfoUpdateCache[key] = cache
		authUserInfoUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q authUserInfoQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q authUserInfoQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for auth_user_info")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AuthUserInfoSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o AuthUserInfoSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o AuthUserInfoSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AuthUserInfoSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authUserInfoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"auth_user_info\" SET %s WHERE (\"auth_user_info_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(authUserInfoPrimaryKeyColumns), len(colNames)+1, len(authUserInfoPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in authUserInfo slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *AuthUserInfo) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *AuthUserInfo) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *AuthUserInfo) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *AuthUserInfo) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no auth_user_info provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(authUserInfoColumnsWithDefault, o)

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

	authUserInfoUpsertCacheMut.RLock()
	cache, cached := authUserInfoUpsertCache[key]
	authUserInfoUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			authUserInfoColumns,
			authUserInfoColumnsWithDefault,
			authUserInfoColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			authUserInfoColumns,
			authUserInfoPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert auth_user_info, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(authUserInfoPrimaryKeyColumns))
			copy(conflict, authUserInfoPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"auth_user_info\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(authUserInfoType, authUserInfoMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(authUserInfoType, authUserInfoMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for auth_user_info")
	}

	if !cached {
		authUserInfoUpsertCacheMut.Lock()
		authUserInfoUpsertCache[key] = cache
		authUserInfoUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single AuthUserInfo record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *AuthUserInfo) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single AuthUserInfo record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *AuthUserInfo) DeleteG() error {
	if o == nil {
		return errors.New("models: no AuthUserInfo provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single AuthUserInfo record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *AuthUserInfo) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single AuthUserInfo record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AuthUserInfo) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no AuthUserInfo provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), authUserInfoPrimaryKeyMapping)
	sql := "DELETE FROM \"auth_user_info\" WHERE \"auth_user_info_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from auth_user_info")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q authUserInfoQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q authUserInfoQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no authUserInfoQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from auth_user_info")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o AuthUserInfoSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o AuthUserInfoSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no AuthUserInfo slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o AuthUserInfoSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AuthUserInfoSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no AuthUserInfo slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(authUserInfoBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authUserInfoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"auth_user_info\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, authUserInfoPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(authUserInfoPrimaryKeyColumns), 1, len(authUserInfoPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from authUserInfo slice")
	}

	if len(authUserInfoAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *AuthUserInfo) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *AuthUserInfo) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *AuthUserInfo) ReloadG() error {
	if o == nil {
		return errors.New("models: no AuthUserInfo provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *AuthUserInfo) Reload(exec boil.Executor) error {
	ret, err := FindAuthUserInfo(exec, o.AuthUserInfoID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AuthUserInfoSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AuthUserInfoSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AuthUserInfoSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty AuthUserInfoSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AuthUserInfoSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	authUserInfos := AuthUserInfoSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authUserInfoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"auth_user_info\".* FROM \"auth_user_info\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, authUserInfoPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(authUserInfoPrimaryKeyColumns), 1, len(authUserInfoPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&authUserInfos)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AuthUserInfoSlice")
	}

	*o = authUserInfos

	return nil
}

// AuthUserInfoExists checks if the AuthUserInfo row exists.
func AuthUserInfoExists(exec boil.Executor, authUserInfoID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"auth_user_info\" where \"auth_user_info_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, authUserInfoID)
	}

	row := exec.QueryRow(sql, authUserInfoID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if auth_user_info exists")
	}

	return exists, nil
}

// AuthUserInfoExistsG checks if the AuthUserInfo row exists.
func AuthUserInfoExistsG(authUserInfoID int) (bool, error) {
	return AuthUserInfoExists(boil.GetDB(), authUserInfoID)
}

// AuthUserInfoExistsGP checks if the AuthUserInfo row exists. Panics on error.
func AuthUserInfoExistsGP(authUserInfoID int) bool {
	e, err := AuthUserInfoExists(boil.GetDB(), authUserInfoID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// AuthUserInfoExistsP checks if the AuthUserInfo row exists. Panics on error.
func AuthUserInfoExistsP(exec boil.Executor, authUserInfoID int) bool {
	e, err := AuthUserInfoExists(exec, authUserInfoID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
