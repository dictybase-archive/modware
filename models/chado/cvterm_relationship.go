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

// CvtermRelationship is an object representing the database table.
type CvtermRelationship struct {
	CvtermRelationshipID int `boil:"cvterm_relationship_id" json:"cvterm_relationship_id" toml:"cvterm_relationship_id" yaml:"cvterm_relationship_id"`
	TypeID               int `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`
	SubjectID            int `boil:"subject_id" json:"subject_id" toml:"subject_id" yaml:"subject_id"`
	ObjectID             int `boil:"object_id" json:"object_id" toml:"object_id" yaml:"object_id"`

	R *cvtermRelationshipR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cvtermRelationshipL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// cvtermRelationshipR is where relationships are stored.
type cvtermRelationshipR struct {
	Object  *Cvterm
	Subject *Cvterm
	Type    *Cvterm
}

// cvtermRelationshipL is where Load methods for each relationship are stored.
type cvtermRelationshipL struct{}

var (
	cvtermRelationshipColumns               = []string{"cvterm_relationship_id", "type_id", "subject_id", "object_id"}
	cvtermRelationshipColumnsWithoutDefault = []string{"type_id", "subject_id", "object_id"}
	cvtermRelationshipColumnsWithDefault    = []string{"cvterm_relationship_id"}
	cvtermRelationshipPrimaryKeyColumns     = []string{"cvterm_relationship_id"}
)

type (
	// CvtermRelationshipSlice is an alias for a slice of pointers to CvtermRelationship.
	// This should generally be used opposed to []CvtermRelationship.
	CvtermRelationshipSlice []*CvtermRelationship
	// CvtermRelationshipHook is the signature for custom CvtermRelationship hook methods
	CvtermRelationshipHook func(boil.Executor, *CvtermRelationship) error

	cvtermRelationshipQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cvtermRelationshipType                 = reflect.TypeOf(&CvtermRelationship{})
	cvtermRelationshipMapping              = queries.MakeStructMapping(cvtermRelationshipType)
	cvtermRelationshipPrimaryKeyMapping, _ = queries.BindMapping(cvtermRelationshipType, cvtermRelationshipMapping, cvtermRelationshipPrimaryKeyColumns)
	cvtermRelationshipInsertCacheMut       sync.RWMutex
	cvtermRelationshipInsertCache          = make(map[string]insertCache)
	cvtermRelationshipUpdateCacheMut       sync.RWMutex
	cvtermRelationshipUpdateCache          = make(map[string]updateCache)
	cvtermRelationshipUpsertCacheMut       sync.RWMutex
	cvtermRelationshipUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var cvtermRelationshipBeforeInsertHooks []CvtermRelationshipHook
var cvtermRelationshipBeforeUpdateHooks []CvtermRelationshipHook
var cvtermRelationshipBeforeDeleteHooks []CvtermRelationshipHook
var cvtermRelationshipBeforeUpsertHooks []CvtermRelationshipHook

var cvtermRelationshipAfterInsertHooks []CvtermRelationshipHook
var cvtermRelationshipAfterSelectHooks []CvtermRelationshipHook
var cvtermRelationshipAfterUpdateHooks []CvtermRelationshipHook
var cvtermRelationshipAfterDeleteHooks []CvtermRelationshipHook
var cvtermRelationshipAfterUpsertHooks []CvtermRelationshipHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CvtermRelationship) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermRelationshipBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CvtermRelationship) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermRelationshipBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CvtermRelationship) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermRelationshipBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CvtermRelationship) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermRelationshipBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CvtermRelationship) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermRelationshipAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CvtermRelationship) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermRelationshipAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CvtermRelationship) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermRelationshipAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CvtermRelationship) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermRelationshipAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CvtermRelationship) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cvtermRelationshipAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCvtermRelationshipHook registers your hook function for all future operations.
func AddCvtermRelationshipHook(hookPoint boil.HookPoint, cvtermRelationshipHook CvtermRelationshipHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cvtermRelationshipBeforeInsertHooks = append(cvtermRelationshipBeforeInsertHooks, cvtermRelationshipHook)
	case boil.BeforeUpdateHook:
		cvtermRelationshipBeforeUpdateHooks = append(cvtermRelationshipBeforeUpdateHooks, cvtermRelationshipHook)
	case boil.BeforeDeleteHook:
		cvtermRelationshipBeforeDeleteHooks = append(cvtermRelationshipBeforeDeleteHooks, cvtermRelationshipHook)
	case boil.BeforeUpsertHook:
		cvtermRelationshipBeforeUpsertHooks = append(cvtermRelationshipBeforeUpsertHooks, cvtermRelationshipHook)
	case boil.AfterInsertHook:
		cvtermRelationshipAfterInsertHooks = append(cvtermRelationshipAfterInsertHooks, cvtermRelationshipHook)
	case boil.AfterSelectHook:
		cvtermRelationshipAfterSelectHooks = append(cvtermRelationshipAfterSelectHooks, cvtermRelationshipHook)
	case boil.AfterUpdateHook:
		cvtermRelationshipAfterUpdateHooks = append(cvtermRelationshipAfterUpdateHooks, cvtermRelationshipHook)
	case boil.AfterDeleteHook:
		cvtermRelationshipAfterDeleteHooks = append(cvtermRelationshipAfterDeleteHooks, cvtermRelationshipHook)
	case boil.AfterUpsertHook:
		cvtermRelationshipAfterUpsertHooks = append(cvtermRelationshipAfterUpsertHooks, cvtermRelationshipHook)
	}
}

// OneP returns a single cvtermRelationship record from the query, and panics on error.
func (q cvtermRelationshipQuery) OneP() *CvtermRelationship {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single cvtermRelationship record from the query.
func (q cvtermRelationshipQuery) One() (*CvtermRelationship, error) {
	o := &CvtermRelationship{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: failed to execute a one query for cvterm_relationship")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all CvtermRelationship records from the query, and panics on error.
func (q cvtermRelationshipQuery) AllP() CvtermRelationshipSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all CvtermRelationship records from the query.
func (q cvtermRelationshipQuery) All() (CvtermRelationshipSlice, error) {
	var o CvtermRelationshipSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "chado: failed to assign all query results to CvtermRelationship slice")
	}

	if len(cvtermRelationshipAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all CvtermRelationship records in the query, and panics on error.
func (q cvtermRelationshipQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all CvtermRelationship records in the query.
func (q cvtermRelationshipQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "chado: failed to count cvterm_relationship rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q cvtermRelationshipQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q cvtermRelationshipQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "chado: failed to check if cvterm_relationship exists")
	}

	return count > 0, nil
}

// ObjectG pointed to by the foreign key.
func (o *CvtermRelationship) ObjectG(mods ...qm.QueryMod) cvtermQuery {
	return o.Object(boil.GetDB(), mods...)
}

// Object pointed to by the foreign key.
func (o *CvtermRelationship) Object(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.ObjectID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// SubjectG pointed to by the foreign key.
func (o *CvtermRelationship) SubjectG(mods ...qm.QueryMod) cvtermQuery {
	return o.Subject(boil.GetDB(), mods...)
}

// Subject pointed to by the foreign key.
func (o *CvtermRelationship) Subject(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.SubjectID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// TypeG pointed to by the foreign key.
func (o *CvtermRelationship) TypeG(mods ...qm.QueryMod) cvtermQuery {
	return o.Type(boil.GetDB(), mods...)
}

// Type pointed to by the foreign key.
func (o *CvtermRelationship) Type(exec boil.Executor, mods ...qm.QueryMod) cvtermQuery {
	queryMods := []qm.QueryMod{
		qm.Where("cvterm_id=$1", o.TypeID),
	}

	queryMods = append(queryMods, mods...)

	query := Cvterms(exec, queryMods...)
	queries.SetFrom(query.Query, "\"cvterm\"")

	return query
}

// LoadObject allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermRelationshipL) LoadObject(e boil.Executor, singular bool, maybeCvtermRelationship interface{}) error {
	var slice []*CvtermRelationship
	var object *CvtermRelationship

	count := 1
	if singular {
		object = maybeCvtermRelationship.(*CvtermRelationship)
	} else {
		slice = *maybeCvtermRelationship.(*CvtermRelationshipSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermRelationshipR{}
		args[0] = object.ObjectID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermRelationshipR{}
			args[i] = obj.ObjectID
		}
	}

	query := fmt.Sprintf(
		"select * from \"cvterm\" where \"cvterm_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Cvterm")
	}
	defer results.Close()

	var resultSlice []*Cvterm
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Cvterm")
	}

	if len(cvtermRelationshipAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Object = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ObjectID == foreign.CvtermID {
				local.R.Object = foreign
				break
			}
		}
	}

	return nil
}

// LoadSubject allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermRelationshipL) LoadSubject(e boil.Executor, singular bool, maybeCvtermRelationship interface{}) error {
	var slice []*CvtermRelationship
	var object *CvtermRelationship

	count := 1
	if singular {
		object = maybeCvtermRelationship.(*CvtermRelationship)
	} else {
		slice = *maybeCvtermRelationship.(*CvtermRelationshipSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermRelationshipR{}
		args[0] = object.SubjectID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermRelationshipR{}
			args[i] = obj.SubjectID
		}
	}

	query := fmt.Sprintf(
		"select * from \"cvterm\" where \"cvterm_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Cvterm")
	}
	defer results.Close()

	var resultSlice []*Cvterm
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Cvterm")
	}

	if len(cvtermRelationshipAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Subject = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.SubjectID == foreign.CvtermID {
				local.R.Subject = foreign
				break
			}
		}
	}

	return nil
}

// LoadType allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (cvtermRelationshipL) LoadType(e boil.Executor, singular bool, maybeCvtermRelationship interface{}) error {
	var slice []*CvtermRelationship
	var object *CvtermRelationship

	count := 1
	if singular {
		object = maybeCvtermRelationship.(*CvtermRelationship)
	} else {
		slice = *maybeCvtermRelationship.(*CvtermRelationshipSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		object.R = &cvtermRelationshipR{}
		args[0] = object.TypeID
	} else {
		for i, obj := range slice {
			obj.R = &cvtermRelationshipR{}
			args[i] = obj.TypeID
		}
	}

	query := fmt.Sprintf(
		"select * from \"cvterm\" where \"cvterm_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Cvterm")
	}
	defer results.Close()

	var resultSlice []*Cvterm
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Cvterm")
	}

	if len(cvtermRelationshipAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Type = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.TypeID == foreign.CvtermID {
				local.R.Type = foreign
				break
			}
		}
	}

	return nil
}

// SetObject of the cvterm_relationship to the related item.
// Sets o.R.Object to related.
// Adds o to related.R.ObjectCvtermRelationship.
func (o *CvtermRelationship) SetObject(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"cvterm_relationship\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"object_id"}),
		strmangle.WhereClause("\"", "\"", 2, cvtermRelationshipPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.CvtermRelationshipID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ObjectID = related.CvtermID

	if o.R == nil {
		o.R = &cvtermRelationshipR{
			Object: related,
		}
	} else {
		o.R.Object = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			ObjectCvtermRelationship: o,
		}
	} else {
		related.R.ObjectCvtermRelationship = o
	}

	return nil
}

// SetSubject of the cvterm_relationship to the related item.
// Sets o.R.Subject to related.
// Adds o to related.R.SubjectCvtermRelationship.
func (o *CvtermRelationship) SetSubject(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"cvterm_relationship\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"subject_id"}),
		strmangle.WhereClause("\"", "\"", 2, cvtermRelationshipPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.CvtermRelationshipID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.SubjectID = related.CvtermID

	if o.R == nil {
		o.R = &cvtermRelationshipR{
			Subject: related,
		}
	} else {
		o.R.Subject = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			SubjectCvtermRelationship: o,
		}
	} else {
		related.R.SubjectCvtermRelationship = o
	}

	return nil
}

// SetType of the cvterm_relationship to the related item.
// Sets o.R.Type to related.
// Adds o to related.R.TypeCvtermRelationship.
func (o *CvtermRelationship) SetType(exec boil.Executor, insert bool, related *Cvterm) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"cvterm_relationship\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
		strmangle.WhereClause("\"", "\"", 2, cvtermRelationshipPrimaryKeyColumns),
	)
	values := []interface{}{related.CvtermID, o.CvtermRelationshipID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TypeID = related.CvtermID

	if o.R == nil {
		o.R = &cvtermRelationshipR{
			Type: related,
		}
	} else {
		o.R.Type = related
	}

	if related.R == nil {
		related.R = &cvtermR{
			TypeCvtermRelationship: o,
		}
	} else {
		related.R.TypeCvtermRelationship = o
	}

	return nil
}

// CvtermRelationshipsG retrieves all records.
func CvtermRelationshipsG(mods ...qm.QueryMod) cvtermRelationshipQuery {
	return CvtermRelationships(boil.GetDB(), mods...)
}

// CvtermRelationships retrieves all the records using an executor.
func CvtermRelationships(exec boil.Executor, mods ...qm.QueryMod) cvtermRelationshipQuery {
	mods = append(mods, qm.From("\"cvterm_relationship\""))
	return cvtermRelationshipQuery{NewQuery(exec, mods...)}
}

// FindCvtermRelationshipG retrieves a single record by ID.
func FindCvtermRelationshipG(cvtermRelationshipID int, selectCols ...string) (*CvtermRelationship, error) {
	return FindCvtermRelationship(boil.GetDB(), cvtermRelationshipID, selectCols...)
}

// FindCvtermRelationshipGP retrieves a single record by ID, and panics on error.
func FindCvtermRelationshipGP(cvtermRelationshipID int, selectCols ...string) *CvtermRelationship {
	retobj, err := FindCvtermRelationship(boil.GetDB(), cvtermRelationshipID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindCvtermRelationship retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCvtermRelationship(exec boil.Executor, cvtermRelationshipID int, selectCols ...string) (*CvtermRelationship, error) {
	cvtermRelationshipObj := &CvtermRelationship{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"cvterm_relationship\" where \"cvterm_relationship_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, cvtermRelationshipID)

	err := q.Bind(cvtermRelationshipObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "chado: unable to select from cvterm_relationship")
	}

	return cvtermRelationshipObj, nil
}

// FindCvtermRelationshipP retrieves a single record by ID with an executor, and panics on error.
func FindCvtermRelationshipP(exec boil.Executor, cvtermRelationshipID int, selectCols ...string) *CvtermRelationship {
	retobj, err := FindCvtermRelationship(exec, cvtermRelationshipID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *CvtermRelationship) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *CvtermRelationship) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *CvtermRelationship) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *CvtermRelationship) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no cvterm_relationship provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cvtermRelationshipColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	cvtermRelationshipInsertCacheMut.RLock()
	cache, cached := cvtermRelationshipInsertCache[key]
	cvtermRelationshipInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			cvtermRelationshipColumns,
			cvtermRelationshipColumnsWithDefault,
			cvtermRelationshipColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(cvtermRelationshipType, cvtermRelationshipMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cvtermRelationshipType, cvtermRelationshipMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO \"cvterm_relationship\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

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
		return errors.Wrap(err, "chado: unable to insert into cvterm_relationship")
	}

	if !cached {
		cvtermRelationshipInsertCacheMut.Lock()
		cvtermRelationshipInsertCache[key] = cache
		cvtermRelationshipInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single CvtermRelationship record. See Update for
// whitelist behavior description.
func (o *CvtermRelationship) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single CvtermRelationship record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *CvtermRelationship) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the CvtermRelationship, and panics on error.
// See Update for whitelist behavior description.
func (o *CvtermRelationship) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the CvtermRelationship.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *CvtermRelationship) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	cvtermRelationshipUpdateCacheMut.RLock()
	cache, cached := cvtermRelationshipUpdateCache[key]
	cvtermRelationshipUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(cvtermRelationshipColumns, cvtermRelationshipPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("chado: unable to update cvterm_relationship, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"cvterm_relationship\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, cvtermRelationshipPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cvtermRelationshipType, cvtermRelationshipMapping, append(wl, cvtermRelationshipPrimaryKeyColumns...))
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
		return errors.Wrap(err, "chado: unable to update cvterm_relationship row")
	}

	if !cached {
		cvtermRelationshipUpdateCacheMut.Lock()
		cvtermRelationshipUpdateCache[key] = cache
		cvtermRelationshipUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q cvtermRelationshipQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q cvtermRelationshipQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all for cvterm_relationship")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o CvtermRelationshipSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o CvtermRelationshipSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o CvtermRelationshipSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CvtermRelationshipSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cvtermRelationshipPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE \"cvterm_relationship\" SET %s WHERE (\"cvterm_relationship_id\") IN (%s)",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(cvtermRelationshipPrimaryKeyColumns), len(colNames)+1, len(cvtermRelationshipPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to update all in cvtermRelationship slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *CvtermRelationship) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *CvtermRelationship) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *CvtermRelationship) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *CvtermRelationship) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("chado: no cvterm_relationship provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(cvtermRelationshipColumnsWithDefault, o)

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

	cvtermRelationshipUpsertCacheMut.RLock()
	cache, cached := cvtermRelationshipUpsertCache[key]
	cvtermRelationshipUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			cvtermRelationshipColumns,
			cvtermRelationshipColumnsWithDefault,
			cvtermRelationshipColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			cvtermRelationshipColumns,
			cvtermRelationshipPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("chado: unable to upsert cvterm_relationship, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(cvtermRelationshipPrimaryKeyColumns))
			copy(conflict, cvtermRelationshipPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"cvterm_relationship\"", updateOnConflict, ret, update, conflict, whitelist)

		cache.valueMapping, err = queries.BindMapping(cvtermRelationshipType, cvtermRelationshipMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cvtermRelationshipType, cvtermRelationshipMapping, ret)
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
		return errors.Wrap(err, "chado: unable to upsert for cvterm_relationship")
	}

	if !cached {
		cvtermRelationshipUpsertCacheMut.Lock()
		cvtermRelationshipUpsertCache[key] = cache
		cvtermRelationshipUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single CvtermRelationship record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *CvtermRelationship) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single CvtermRelationship record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *CvtermRelationship) DeleteG() error {
	if o == nil {
		return errors.New("chado: no CvtermRelationship provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single CvtermRelationship record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *CvtermRelationship) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single CvtermRelationship record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CvtermRelationship) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no CvtermRelationship provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cvtermRelationshipPrimaryKeyMapping)
	sql := "DELETE FROM \"cvterm_relationship\" WHERE \"cvterm_relationship_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete from cvterm_relationship")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q cvtermRelationshipQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q cvtermRelationshipQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("chado: no cvtermRelationshipQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from cvterm_relationship")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o CvtermRelationshipSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o CvtermRelationshipSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("chado: no CvtermRelationship slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o CvtermRelationshipSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CvtermRelationshipSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("chado: no CvtermRelationship slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(cvtermRelationshipBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cvtermRelationshipPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM \"cvterm_relationship\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, cvtermRelationshipPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(cvtermRelationshipPrimaryKeyColumns), 1, len(cvtermRelationshipPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "chado: unable to delete all from cvtermRelationship slice")
	}

	if len(cvtermRelationshipAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *CvtermRelationship) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *CvtermRelationship) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *CvtermRelationship) ReloadG() error {
	if o == nil {
		return errors.New("chado: no CvtermRelationship provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *CvtermRelationship) Reload(exec boil.Executor) error {
	ret, err := FindCvtermRelationship(exec, o.CvtermRelationshipID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *CvtermRelationshipSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *CvtermRelationshipSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CvtermRelationshipSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("chado: empty CvtermRelationshipSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CvtermRelationshipSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	cvtermRelationships := CvtermRelationshipSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cvtermRelationshipPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT \"cvterm_relationship\".* FROM \"cvterm_relationship\" WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, cvtermRelationshipPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(cvtermRelationshipPrimaryKeyColumns), 1, len(cvtermRelationshipPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&cvtermRelationships)
	if err != nil {
		return errors.Wrap(err, "chado: unable to reload all in CvtermRelationshipSlice")
	}

	*o = cvtermRelationships

	return nil
}

// CvtermRelationshipExists checks if the CvtermRelationship row exists.
func CvtermRelationshipExists(exec boil.Executor, cvtermRelationshipID int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from \"cvterm_relationship\" where \"cvterm_relationship_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, cvtermRelationshipID)
	}

	row := exec.QueryRow(sql, cvtermRelationshipID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "chado: unable to check if cvterm_relationship exists")
	}

	return exists, nil
}

// CvtermRelationshipExistsG checks if the CvtermRelationship row exists.
func CvtermRelationshipExistsG(cvtermRelationshipID int) (bool, error) {
	return CvtermRelationshipExists(boil.GetDB(), cvtermRelationshipID)
}

// CvtermRelationshipExistsGP checks if the CvtermRelationship row exists. Panics on error.
func CvtermRelationshipExistsGP(cvtermRelationshipID int) bool {
	e, err := CvtermRelationshipExists(boil.GetDB(), cvtermRelationshipID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// CvtermRelationshipExistsP checks if the CvtermRelationship row exists. Panics on error.
func CvtermRelationshipExistsP(exec boil.Executor, cvtermRelationshipID int) bool {
	e, err := CvtermRelationshipExists(exec, cvtermRelationshipID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
