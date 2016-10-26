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

// Analysisfeature is an object representing the database table.
type Analysisfeature struct {
	AnalysisfeatureID int          `boil:"analysisfeature_id" json:"analysisfeature_id" toml:"analysisfeature_id" yaml:"analysisfeature_id"`
	FeatureID         int          `boil:"feature_id" json:"feature_id" toml:"feature_id" yaml:"feature_id"`
	AnalysisID        int          `boil:"analysis_id" json:"analysis_id" toml:"analysis_id" yaml:"analysis_id"`
	Rawscore          null.Float64 `boil:"rawscore" json:"rawscore,omitempty" toml:"rawscore" yaml:"rawscore,omitempty"`
	Normscore         null.Float64 `boil:"normscore" json:"normscore,omitempty" toml:"normscore" yaml:"normscore,omitempty"`
	Significance      null.Float64 `boil:"significance" json:"significance,omitempty" toml:"significance" yaml:"significance,omitempty"`
	Identity          null.Float64 `boil:"identity" json:"identity,omitempty" toml:"identity" yaml:"identity,omitempty"`

	R *analysisfeatureR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L analysisfeatureL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// analysisfeatureR is where relationships are stored.
type analysisfeatureR struct {
	Analysi             *Analysi
	Feature             *Feature
	Analysisfeatureprop *Analysisfeatureprop
}

// analysisfeatureL is where Load methods for each relationship are stored.
type analysisfeatureL struct{}

var (
	analysisfeatureColumns               = []string{"analysisfeature_id", "feature_id", "analysis_id", "rawscore", "normscore", "significance", "identity"}
	analysisfeatureColumnsWithoutDefault = []string{"feature_id", "analysis_id", "rawscore", "normscore", "significance", "identity"}
	analysisfeatureColumnsWithDefault    = []string{"analysisfeature_id"}
	analysisfeaturePrimaryKeyColumns     = []string{"analysisfeature_id"}
)

type (
	// AnalysisfeatureSlice is an alias for a slice of pointers to Analysisfeature.
	// This should generally be used opposed to []Analysisfeature.
	AnalysisfeatureSlice []*Analysisfeature
	// AnalysisfeatureHook is the signature for custom Analysisfeature hook methods
	AnalysisfeatureHook func(boil.Executor, *Analysisfeature) error

	analysisfeatureQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	analysisfeatureType                 = reflect.TypeOf(&Analysisfeature{})
	analysisfeatureMapping              = queries.MakeStructMapping(analysisfeatureType)
	analysisfeaturePrimaryKeyMapping, _ = queries.BindMapping(analysisfeatureType, analysisfeatureMapping, analysisfeaturePrimaryKeyColumns)
	analysisfeatureInsertCacheMut       sync.RWMutex
	analysisfeatureInsertCache          = make(map[string]insertCache)
	analysisfeatureUpdateCacheMut       sync.RWMutex
	analysisfeatureUpdateCache          = make(map[string]updateCache)
	analysisfeatureUpsertCacheMut       sync.RWMutex
	analysisfeatureUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var analysisfeatureBeforeInsertHooks []AnalysisfeatureHook
var analysisfeatureBeforeUpdateHooks []AnalysisfeatureHook
var analysisfeatureBeforeDeleteHooks []AnalysisfeatureHook
var analysisfeatureBeforeUpsertHooks []AnalysisfeatureHook

var analysisfeatureAfterInsertHooks []AnalysisfeatureHook
var analysisfeatureAfterSelectHooks []AnalysisfeatureHook
var analysisfeatureAfterUpdateHooks []AnalysisfeatureHook
var analysisfeatureAfterDeleteHooks []AnalysisfeatureHook
var analysisfeatureAfterUpsertHooks []AnalysisfeatureHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Analysisfeature) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range analysisfeatureBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Analysisfeature) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range analysisfeatureBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Analysisfeature) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range analysisfeatureBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Analysisfeature) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range analysisfeatureBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Analysisfeature) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range analysisfeatureAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Analysisfeature) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range analysisfeatureAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Analysisfeature) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range analysisfeatureAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Analysisfeature) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range analysisfeatureAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Analysisfeature) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range analysisfeatureAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAnalysisfeatureHook registers your hook function for all future operations.
func AddAnalysisfeatureHook(hookPoint boil.HookPoint, analysisfeatureHook AnalysisfeatureHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		analysisfeatureBeforeInsertHooks = append(analysisfeatureBeforeInsertHooks, analysisfeatureHook)
	case boil.BeforeUpdateHook:
		analysisfeatureBeforeUpdateHooks = append(analysisfeatureBeforeUpdateHooks, analysisfeatureHook)
	case boil.BeforeDeleteHook:
		analysisfeatureBeforeDeleteHooks = append(analysisfeatureBeforeDeleteHooks, analysisfeatureHook)
	case boil.BeforeUpsertHook:
		analysisfeatureBeforeUpsertHooks = append(analysisfeatureBeforeUpsertHooks, analysisfeatureHook)
	case boil.AfterInsertHook:
		analysisfeatureAfterInsertHooks = append(analysisfeatureAfterInsertHooks, analysisfeatureHook)
	case boil.AfterSelectHook:
		analysisfeatureAfterSelectHooks = append(analysisfeatureAfterSelectHooks, analysisfeatureHook)
	case boil.AfterUpdateHook:
		analysisfeatureAfterUpdateHooks = append(analysisfeatureAfterUpdateHooks, analysisfeatureHook)
	case boil.AfterDeleteHook:
		analysisfeatureAfterDeleteHooks = append(analysisfeatureAfterDeleteHooks, analysisfeatureHook)
	case boil.AfterUpsertHook:
		analysisfeatureAfterUpsertHooks = append(analysisfeatureAfterUpsertHooks, analysisfeatureHook)
	}
}

// OneP returns a single analysisfeature record from the query, and panics on error.
func (q analysisfeatureQuery) OneP() *Analysisfeature {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single analysisfeature record from the query.
func (q analysisfeatureQuery) One() (*Analysisfeature, error) {
	o := &Analysisfeature{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for analysisfeature")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Analysisfeature records from the query, and panics on error.
func (q analysisfeatureQuery) AllP() AnalysisfeatureSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Analysisfeature records from the query.
func (q analysisfeatureQuery) All() (AnalysisfeatureSlice, error) {
	var o AnalysisfeatureSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to Analysisfeature slice")
	}

	if len(analysisfeatureAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Analysisfeature records in the query, and panics on error.
func (q analysisfeatureQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Analysisfeature records in the query.
func (q analysisfeatureQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count analysisfeature rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q analysisfeatureQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q analysisfeatureQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if analysisfeature exists")
	}

	return count > 0, nil
}

// AnalysiG pointed to by the foreign key.
func (o *Analysisfeature) AnalysiG(mods ...qm.QueryMod) analysiQuery {
	return o.Analysi(boil.GetDB(), mods...)
}

// Analysi pointed to by the foreign key.
func (o *Analysisfeature) Analysi(exec boil.Executor, mods ...qm.QueryMod) analysiQuery {
	queryMods := []qm.QueryMod{
		qm.Where("analysis_id=$1", o.AnalysisID),
	}

	queryMods = append(queryMods, mods...)

	query := Analyses(exec, queryMods...)
	queries.SetFrom(query.Query, "\"analysis\"")

	return query
}

// FeatureG pointed to by the foreign key.
func (o *Analysisfeature) FeatureG(mods ...qm.QueryMod) featureQuery {
	return o.Feature(boil.GetDB(), mods...)
}

// Feature pointed to by the foreign key.
func (o *Analysisfeature) Feature(exec boil.Executor, mods ...qm.QueryMod) featureQuery {
	queryMods := []qm.QueryMod{
		qm.Where("feature_id=$1", o.FeatureID),
	}

	queryMods = append(queryMods, mods...)

	query := Features(exec, queryMods...)
	queries.SetFrom(query.Query, "\"feature\"")

	return query
}

// AnalysisfeaturepropG pointed to by the foreign key.
func (o *Analysisfeature) AnalysisfeaturepropG(mods ...qm.QueryMod) analysisfeaturepropQuery {
	return o.Analysisfeatureprop(boil.GetDB(), mods...)
}

// Analysisfeatureprop pointed to by the foreign key.
func (o *Analysisfeature) Analysisfeatureprop(exec boil.Executor, mods ...qm.QueryMod) analysisfeaturepropQuery {
	queryMods := []qm.QueryMod{
		qm.Where("analysisfeature_id=$1", o.AnalysisfeatureID),
	}

	queryMods = append(queryMods, mods...)

	query := Analysisfeatureprops(exec, queryMods...)
	queries.SetFrom(query.Query, "\"analysisfeatureprop\"")

	return query
}

// LoadAnalysi allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (analysisfeatureL) LoadAnalysi(e boil.Executor, singular bool, maybeAnalysisfeature interface{}) error {
	var slice []*Analysisfeature
	var object *Analysisfeature

	count := 1
	if singular {
		object = maybeAnalysisfeature.(*Analysisfeature)
	} else {
		slice = *maybeAnalysisfeature.(*AnalysisfeatureSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &analysisfeatureR{}
		args[0] = object.AnalysisID
	} else {
		for i, obj := range slice {
			obj.R = &analysisfeatureR{}
			args[i] = obj.AnalysisID
		}
	}

	query := fmt.Sprintf(
		"select * from \"analysis\" where \"analysis_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Analysi")
	}
	defer results.Close()

	var resultSlice []*Analysi
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Analysi")
	}

	if len(analysisfeatureAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Analysi = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.AnalysisID == foreign.AnalysisID {
				local.R.Analysi = foreign
				break
			}
		}
	}

	return nil
}

// LoadFeature allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (analysisfeatureL) LoadFeature(e boil.Executor, singular bool, maybeAnalysisfeature interface{}) error {
	var slice []*Analysisfeature
	var object *Analysisfeature

	count := 1
	if singular {
		object = maybeAnalysisfeature.(*Analysisfeature)
	} else {
		slice = *maybeAnalysisfeature.(*AnalysisfeatureSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &analysisfeatureR{}
		args[0] = object.FeatureID
	} else {
		for i, obj := range slice {
			obj.R = &analysisfeatureR{}
			args[i] = obj.FeatureID
		}
	}

	query := fmt.Sprintf(
		"select * from \"feature\" where \"feature_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Feature")
	}
	defer results.Close()

	var resultSlice []*Feature
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Feature")
	}

	if len(analysisfeatureAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Feature = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.FeatureID == foreign.FeatureID {
				local.R.Feature = foreign
				break
			}
		}
	}

	return nil
}

// LoadAnalysisfeatureprop allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (analysisfeatureL) LoadAnalysisfeatureprop(e boil.Executor, singular bool, maybeAnalysisfeature interface{}) error {
	var slice []*Analysisfeature
	var object *Analysisfeature

	count := 1
	if singular {
		object = maybeAnalysisfeature.(*Analysisfeature)
	} else {
		slice = *maybeAnalysisfeature.(*AnalysisfeatureSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &analysisfeatureR{}
		args[0] = object.AnalysisfeatureID
	} else {
		for i, obj := range slice {
			obj.R = &analysisfeatureR{}
			args[i] = obj.AnalysisfeatureID
		}
	}

	query := fmt.Sprintf(
		"select * from \"analysisfeatureprop\" where \"analysisfeature_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Analysisfeatureprop")
	}
	defer results.Close()

	var resultSlice []*Analysisfeatureprop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Analysisfeatureprop")
	}

	if len(analysisfeatureAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Analysisfeatureprop = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.AnalysisfeatureID == foreign.AnalysisfeatureID {
				local.R.Analysisfeatureprop = foreign
				break
			}
		}
	}

	return nil
}

// SetAnalysi of the analysisfeature to the related item.
// Sets o.R.Analysi to related.
// Adds o to related.R.Analysisfeature.
func (o *Analysisfeature) SetAnalysi(exec boil.Executor, insert bool, related *Analysi) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"analysisfeature\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"analysis_id"}),
		strmangle.WhereClause("\"", "\"", 2, analysisfeaturePrimaryKeyColumns),
	)
	values := []interface{}{related.AnalysisID, o.AnalysisfeatureID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.AnalysisID = related.AnalysisID

	if o.R == nil {
		o.R = &analysisfeatureR{
			Analysi: related,
		}
	} else {
		o.R.Analysi = related
	}

	if related.R == nil {
		related.R = &analysiR{
			Analysisfeature: o,
		}
	} else {
		related.R.Analysisfeature = o
	}

	return nil
}

// SetFeature of the analysisfeature to the related item.
// Sets o.R.Feature to related.
// Adds o to related.R.Analysisfeature.
func (o *Analysisfeature) SetFeature(exec boil.Executor, insert bool, related *Feature) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"analysisfeature\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"feature_id"}),
		strmangle.WhereClause("\"", "\"", 2, analysisfeaturePrimaryKeyColumns),
	)
	values := []interface{}{related.FeatureID, o.AnalysisfeatureID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.FeatureID = related.FeatureID

	if o.R == nil {
		o.R = &analysisfeatureR{
			Feature: related,
		}
	} else {
		o.R.Feature = related
	}

	if related.R == nil {
		related.R = &featureR{
			Analysisfeature: o,
		}
	} else {
		related.R.Analysisfeature = o
	}

	return nil
}

// SetAnalysisfeatureprop of the analysisfeature to the related item.
// Sets o.R.Analysisfeatureprop to related.
// Adds o to related.R.Analysisfeature.
func (o *Analysisfeature) SetAnalysisfeatureprop(exec boil.Executor, insert bool, related *Analysisfeatureprop) error {
	var err error

	if insert {
		related.AnalysisfeatureID = o.AnalysisfeatureID

		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	} else {
		updateQuery := fmt.Sprintf(
			"UPDATE \"analysisfeatureprop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, []string{"analysisfeature_id"}),
			strmangle.WhereClause("\"", "\"", 2, analysisfeaturepropPrimaryKeyColumns),
		)
		values := []interface{}{o.AnalysisfeatureID, related.AnalysisfeaturepropID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, updateQuery)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		if _, err = exec.Exec(updateQuery, values...); err != nil {
			return errors.Wrap(err, "failed to update foreign table")
		}

		related.AnalysisfeatureID = o.AnalysisfeatureID

	}

	if o.R == nil {
		o.R = &analysisfeatureR{
			Analysisfeatureprop: related,
		}
	} else {
		o.R.Analysisfeatureprop = related
	}

	if related.R == nil {
		related.R = &analysisfeaturepropR{
			Analysisfeature: o,
		}
	} else {
		related.R.Analysisfeature = o
	}
	return nil
}

// AnalysisfeaturesG retrieves all records.
func AnalysisfeaturesG(mods ...qm.QueryMod) analysisfeatureQuery {
	return Analysisfeatures(boil.GetDB(), mods...)
}

// Analysisfeatures retrieves all the records using an executor.
func Analysisfeatures(exec boil.Executor, mods ...qm.QueryMod) analysisfeatureQuery {
	mods = append(mods, qm.From("\"analysisfeature\""))
	return analysisfeatureQuery{NewQuery(exec, mods...)}
}

// FindAnalysisfeatureG retrieves a single record by ID.
func FindAnalysisfeatureG(analysisfeatureID int, selectCols ...string) (*Analysisfeature, error) {
	return FindAnalysisfeature(boil.GetDB(), analysisfeatureID, selectCols...)
}

// FindAnalysisfeatureGP retrieves a single record by ID, and panics on error.
func FindAnalysisfeatureGP(analysisfeatureID int, selectCols ...string) *Analysisfeature {
	retobj, err := FindAnalysisfeature(boil.GetDB(), analysisfeatureID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindAnalysisfeature retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAnalysisfeature(exec boil.Executor, analysisfeatureID int, selectCols ...string) (*Analysisfeature, error) {
	analysisfeatureObj := &Analysisfeature{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"analysisfeature\" where \"analysisfeature_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, analysisfeatureID)

	err := q.Bind(analysisfeatureObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from analysisfeature")
	}

	return analysisfeatureObj, nil
}

// FindAnalysisfeatureP retrieves a single record by ID with an executor, and panics on error.
func FindAnalysisfeatureP(exec boil.Executor, analysisfeatureID int, selectCols ...string) *Analysisfeature {
	retobj, err := FindAnalysisfeature(exec, analysisfeatureID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Analysisfeature) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Analysisfeature) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Analysisfeature) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Analysisfeature) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no analysisfeature provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(analysisfeatureColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	analysisfeatureInsertCacheMut.RLock()
	cache, cached := analysisfeatureInsertCache[key]
	analysisfeatureInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			analysisfeatureColumns,
			analysisfeatureColumnsWithDefault,
			analysisfeatureColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(analysisfeatureType, analysisfeatureMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(analysisfeatureType, analysisfeatureMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"analysisfeature\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into analysisfeature")
	}

	if !cached {
		analysisfeatureInsertCacheMut.Lock()
		analysisfeatureInsertCache[key] = cache
		analysisfeatureInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Analysisfeature record. See Update for
// whitelist behavior description.
func (o *Analysisfeature) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Analysisfeature record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Analysisfeature) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Analysisfeature, and panics on error.
// See Update for whitelist behavior description.
func (o *Analysisfeature) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Analysisfeature.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Analysisfeature) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	analysisfeatureUpdateCacheMut.RLock()
	cache, cached := analysisfeatureUpdateCache[key]
	analysisfeatureUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(analysisfeatureColumns, analysisfeaturePrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update analysisfeature, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"analysisfeature\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, analysisfeaturePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(analysisfeatureType, analysisfeatureMapping, append(wl, analysisfeaturePrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update analysisfeature row")
	}

	if !cached {
		analysisfeatureUpdateCacheMut.Lock()
		analysisfeatureUpdateCache[key] = cache
		analysisfeatureUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q analysisfeatureQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q analysisfeatureQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for analysisfeature")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AnalysisfeatureSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o AnalysisfeatureSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o AnalysisfeatureSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AnalysisfeatureSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), analysisfeaturePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"analysisfeature\" SET %s WHERE (\"analysisfeature_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(analysisfeaturePrimaryKeyColumns), len(colNames)+1, len(analysisfeaturePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in analysisfeature slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Analysisfeature) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Analysisfeature) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Analysisfeature) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Analysisfeature) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no analysisfeature provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(analysisfeatureColumnsWithDefault, o)

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

	analysisfeatureUpsertCacheMut.RLock()
	cache, cached := analysisfeatureUpsertCache[key]
	analysisfeatureUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			analysisfeatureColumns,
			analysisfeatureColumnsWithDefault,
			analysisfeatureColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			analysisfeatureColumns,
			analysisfeaturePrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert analysisfeature, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(analysisfeaturePrimaryKeyColumns))
			copy(conflict, analysisfeaturePrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"analysisfeature\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(analysisfeatureType, analysisfeatureMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(analysisfeatureType, analysisfeatureMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for analysisfeature")
	}

	if !cached {
		analysisfeatureUpsertCacheMut.Lock()
		analysisfeatureUpsertCache[key] = cache
		analysisfeatureUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Analysisfeature record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Analysisfeature) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Analysisfeature record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Analysisfeature) DeleteG() error {
	if o == nil {
		return errors.New("chado: no Analysisfeature provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Analysisfeature record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Analysisfeature) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Analysisfeature record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Analysisfeature) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Analysisfeature provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), analysisfeaturePrimaryKeyMapping)
	sql := "DELETE FROM \"analysisfeature\" WHERE \"analysisfeature_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from analysisfeature")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q analysisfeatureQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q analysisfeatureQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no analysisfeatureQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from analysisfeature")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o AnalysisfeatureSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o AnalysisfeatureSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no Analysisfeature slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o AnalysisfeatureSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AnalysisfeatureSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no Analysisfeature slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(analysisfeatureBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), analysisfeaturePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"analysisfeature\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, analysisfeaturePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(analysisfeaturePrimaryKeyColumns), 1, len(analysisfeaturePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from analysisfeature slice")
	}

	if len(analysisfeatureAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Analysisfeature) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Analysisfeature) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Analysisfeature) ReloadG() error {
	if o == nil {
		return errors.New("chado: no Analysisfeature provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Analysisfeature) Reload(exec boil.Executor) error {
	ret, err := FindAnalysisfeature(exec, o.AnalysisfeatureID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AnalysisfeatureSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AnalysisfeatureSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AnalysisfeatureSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty AnalysisfeatureSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AnalysisfeatureSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	analysisfeatures := AnalysisfeatureSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), analysisfeaturePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"analysisfeature\".* FROM \"analysisfeature\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, analysisfeaturePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(analysisfeaturePrimaryKeyColumns), 1, len(analysisfeaturePrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&analysisfeatures)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in AnalysisfeatureSlice")
	}

	*o = analysisfeatures

	return nil
}

// AnalysisfeatureExists checks if the Analysisfeature row exists.
func AnalysisfeatureExists(exec boil.Executor, analysisfeatureID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"analysisfeature\" where \"analysisfeature_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, analysisfeatureID)
	}

	row := exec.QueryRow(sql, analysisfeatureID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if analysisfeature exists")
	}

	return exists, nil
}

// AnalysisfeatureExistsG checks if the Analysisfeature row exists.
func AnalysisfeatureExistsG(analysisfeatureID int) (bool, error) {
	return AnalysisfeatureExists(boil.GetDB(), analysisfeatureID)
}

// AnalysisfeatureExistsGP checks if the Analysisfeature row exists. Panics on error.
func AnalysisfeatureExistsGP(analysisfeatureID int) bool {
	e, err := AnalysisfeatureExists(boil.GetDB(), analysisfeatureID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// AnalysisfeatureExistsP checks if the Analysisfeature row exists. Panics on error.
func AnalysisfeatureExistsP(exec boil.Executor, analysisfeatureID int) bool {
	e, err := AnalysisfeatureExists(exec, analysisfeatureID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
