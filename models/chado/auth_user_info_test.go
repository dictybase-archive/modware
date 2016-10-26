package chado

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testAuthUserInfos(t *testing.T) {
	t.Parallel()

	query := AuthUserInfos(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testAuthUserInfosDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserInfo := &AuthUserInfo{}
	if err = randomize.Struct(seed, authUserInfo, authUserInfoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthUserInfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserInfo.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = authUserInfo.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthUserInfos(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthUserInfosQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserInfo := &AuthUserInfo{}
	if err = randomize.Struct(seed, authUserInfo, authUserInfoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthUserInfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserInfo.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = AuthUserInfos(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := AuthUserInfos(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthUserInfosSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserInfo := &AuthUserInfo{}
	if err = randomize.Struct(seed, authUserInfo, authUserInfoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthUserInfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserInfo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AuthUserInfoSlice{authUserInfo}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthUserInfos(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testAuthUserInfosExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserInfo := &AuthUserInfo{}
	if err = randomize.Struct(seed, authUserInfo, authUserInfoDBTypes, true, authUserInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserInfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserInfo.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := AuthUserInfoExists(tx, authUserInfo.AuthUserInfoID)
	if err != nil {
		t.Errorf("Unable to check if AuthUserInfo exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AuthUserInfoExistsG to return true, but got false.")
	}
}
func testAuthUserInfosFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserInfo := &AuthUserInfo{}
	if err = randomize.Struct(seed, authUserInfo, authUserInfoDBTypes, true, authUserInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserInfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserInfo.Insert(tx); err != nil {
		t.Error(err)
	}

	authUserInfoFound, err := FindAuthUserInfo(tx, authUserInfo.AuthUserInfoID)
	if err != nil {
		t.Error(err)
	}

	if authUserInfoFound == nil {
		t.Error("want a record, got nil")
	}
}
func testAuthUserInfosBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserInfo := &AuthUserInfo{}
	if err = randomize.Struct(seed, authUserInfo, authUserInfoDBTypes, true, authUserInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserInfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserInfo.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = AuthUserInfos(tx).Bind(authUserInfo); err != nil {
		t.Error(err)
	}
}

func testAuthUserInfosOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserInfo := &AuthUserInfo{}
	if err = randomize.Struct(seed, authUserInfo, authUserInfoDBTypes, true, authUserInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserInfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserInfo.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := AuthUserInfos(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAuthUserInfosAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserInfoOne := &AuthUserInfo{}
	authUserInfoTwo := &AuthUserInfo{}
	if err = randomize.Struct(seed, authUserInfoOne, authUserInfoDBTypes, false, authUserInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserInfo struct: %s", err)
	}
	if err = randomize.Struct(seed, authUserInfoTwo, authUserInfoDBTypes, false, authUserInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserInfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserInfoOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = authUserInfoTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := AuthUserInfos(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAuthUserInfosCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	authUserInfoOne := &AuthUserInfo{}
	authUserInfoTwo := &AuthUserInfo{}
	if err = randomize.Struct(seed, authUserInfoOne, authUserInfoDBTypes, false, authUserInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserInfo struct: %s", err)
	}
	if err = randomize.Struct(seed, authUserInfoTwo, authUserInfoDBTypes, false, authUserInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserInfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserInfoOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = authUserInfoTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthUserInfos(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func authUserInfoBeforeInsertHook(e boil.Executor, o *AuthUserInfo) error {
	*o = AuthUserInfo{}
	return nil
}

func authUserInfoAfterInsertHook(e boil.Executor, o *AuthUserInfo) error {
	*o = AuthUserInfo{}
	return nil
}

func authUserInfoAfterSelectHook(e boil.Executor, o *AuthUserInfo) error {
	*o = AuthUserInfo{}
	return nil
}

func authUserInfoBeforeUpdateHook(e boil.Executor, o *AuthUserInfo) error {
	*o = AuthUserInfo{}
	return nil
}

func authUserInfoAfterUpdateHook(e boil.Executor, o *AuthUserInfo) error {
	*o = AuthUserInfo{}
	return nil
}

func authUserInfoBeforeDeleteHook(e boil.Executor, o *AuthUserInfo) error {
	*o = AuthUserInfo{}
	return nil
}

func authUserInfoAfterDeleteHook(e boil.Executor, o *AuthUserInfo) error {
	*o = AuthUserInfo{}
	return nil
}

func authUserInfoBeforeUpsertHook(e boil.Executor, o *AuthUserInfo) error {
	*o = AuthUserInfo{}
	return nil
}

func authUserInfoAfterUpsertHook(e boil.Executor, o *AuthUserInfo) error {
	*o = AuthUserInfo{}
	return nil
}

func testAuthUserInfosHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &AuthUserInfo{}
	o := &AuthUserInfo{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, authUserInfoDBTypes, false); err != nil {
		t.Errorf("Unable to randomize AuthUserInfo object: %s", err)
	}

	AddAuthUserInfoHook(boil.BeforeInsertHook, authUserInfoBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	authUserInfoBeforeInsertHooks = []AuthUserInfoHook{}

	AddAuthUserInfoHook(boil.AfterInsertHook, authUserInfoAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	authUserInfoAfterInsertHooks = []AuthUserInfoHook{}

	AddAuthUserInfoHook(boil.AfterSelectHook, authUserInfoAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	authUserInfoAfterSelectHooks = []AuthUserInfoHook{}

	AddAuthUserInfoHook(boil.BeforeUpdateHook, authUserInfoBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	authUserInfoBeforeUpdateHooks = []AuthUserInfoHook{}

	AddAuthUserInfoHook(boil.AfterUpdateHook, authUserInfoAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	authUserInfoAfterUpdateHooks = []AuthUserInfoHook{}

	AddAuthUserInfoHook(boil.BeforeDeleteHook, authUserInfoBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	authUserInfoBeforeDeleteHooks = []AuthUserInfoHook{}

	AddAuthUserInfoHook(boil.AfterDeleteHook, authUserInfoAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	authUserInfoAfterDeleteHooks = []AuthUserInfoHook{}

	AddAuthUserInfoHook(boil.BeforeUpsertHook, authUserInfoBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	authUserInfoBeforeUpsertHooks = []AuthUserInfoHook{}

	AddAuthUserInfoHook(boil.AfterUpsertHook, authUserInfoAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	authUserInfoAfterUpsertHooks = []AuthUserInfoHook{}
}
func testAuthUserInfosInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserInfo := &AuthUserInfo{}
	if err = randomize.Struct(seed, authUserInfo, authUserInfoDBTypes, true, authUserInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserInfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserInfo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthUserInfos(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAuthUserInfosInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserInfo := &AuthUserInfo{}
	if err = randomize.Struct(seed, authUserInfo, authUserInfoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthUserInfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserInfo.Insert(tx, authUserInfoColumns...); err != nil {
		t.Error(err)
	}

	count, err := AuthUserInfos(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAuthUserInfoToOneAuthUserUsingAuthUser(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local AuthUserInfo
	var foreign AuthUser

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, authUserInfoDBTypes, true, authUserInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserInfo struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, authUserDBTypes, true, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.AuthUserID = foreign.AuthUserID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.AuthUser(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.AuthUserID != foreign.AuthUserID {
		t.Errorf("want: %v, got %v", foreign.AuthUserID, check.AuthUserID)
	}

	slice := AuthUserInfoSlice{&local}
	if err = local.L.LoadAuthUser(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.AuthUser == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.AuthUser = nil
	if err = local.L.LoadAuthUser(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.AuthUser == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testAuthUserInfoToOneSetOpAuthUserUsingAuthUser(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a AuthUserInfo
	var b, c AuthUser

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authUserInfoDBTypes, false, strmangle.SetComplement(authUserInfoPrimaryKeyColumns, authUserInfoColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, authUserDBTypes, false, strmangle.SetComplement(authUserPrimaryKeyColumns, authUserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, authUserDBTypes, false, strmangle.SetComplement(authUserPrimaryKeyColumns, authUserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*AuthUser{&b, &c} {
		err = a.SetAuthUser(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.AuthUser != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.AuthUserInfos[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.AuthUserID != x.AuthUserID {
			t.Error("foreign key was wrong value", a.AuthUserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.AuthUserID))
		reflect.Indirect(reflect.ValueOf(&a.AuthUserID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.AuthUserID != x.AuthUserID {
			t.Error("foreign key was wrong value", a.AuthUserID, x.AuthUserID)
		}
	}
}
func testAuthUserInfosReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserInfo := &AuthUserInfo{}
	if err = randomize.Struct(seed, authUserInfo, authUserInfoDBTypes, true, authUserInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserInfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserInfo.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = authUserInfo.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testAuthUserInfosReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserInfo := &AuthUserInfo{}
	if err = randomize.Struct(seed, authUserInfo, authUserInfoDBTypes, true, authUserInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserInfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserInfo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AuthUserInfoSlice{authUserInfo}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testAuthUserInfosSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserInfo := &AuthUserInfo{}
	if err = randomize.Struct(seed, authUserInfo, authUserInfoDBTypes, true, authUserInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserInfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserInfo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := AuthUserInfos(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	authUserInfoDBTypes = map[string]string{"AuthUserID": "integer", "AuthUserInfoID": "integer", "City": "character varying", "Country": "character varying", "FirstAddress": "text", "GroupName": "text", "Organization": "character varying", "Phone": "character varying", "SecondAddress": "text", "State": "character varying", "Zipcode": "character varying"}
	_                   = bytes.MinRead
)

func testAuthUserInfosUpdate(t *testing.T) {
	t.Parallel()

	if len(authUserInfoColumns) == len(authUserInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	authUserInfo := &AuthUserInfo{}
	if err = randomize.Struct(seed, authUserInfo, authUserInfoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthUserInfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserInfo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthUserInfos(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, authUserInfo, authUserInfoDBTypes, true, authUserInfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUserInfo struct: %s", err)
	}

	if err = authUserInfo.Update(tx); err != nil {
		t.Error(err)
	}
}

func testAuthUserInfosSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(authUserInfoColumns) == len(authUserInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	authUserInfo := &AuthUserInfo{}
	if err = randomize.Struct(seed, authUserInfo, authUserInfoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthUserInfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserInfo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthUserInfos(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, authUserInfo, authUserInfoDBTypes, true, authUserInfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AuthUserInfo struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(authUserInfoColumns, authUserInfoPrimaryKeyColumns) {
		fields = authUserInfoColumns
	} else {
		fields = strmangle.SetComplement(
			authUserInfoColumns,
			authUserInfoPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(authUserInfo))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := AuthUserInfoSlice{authUserInfo}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testAuthUserInfosUpsert(t *testing.T) {
	t.Parallel()

	if len(authUserInfoColumns) == len(authUserInfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	authUserInfo := AuthUserInfo{}
	if err = randomize.Struct(seed, &authUserInfo, authUserInfoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthUserInfo struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserInfo.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert AuthUserInfo: %s", err)
	}

	count, err := AuthUserInfos(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &authUserInfo, authUserInfoDBTypes, false, authUserInfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AuthUserInfo struct: %s", err)
	}

	if err = authUserInfo.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert AuthUserInfo: %s", err)
	}

	count, err = AuthUserInfos(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
