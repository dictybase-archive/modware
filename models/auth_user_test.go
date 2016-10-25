package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testAuthUsers(t *testing.T) {
	t.Parallel()

	query := AuthUsers(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testAuthUsersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUser := &AuthUser{}
	if err = randomize.Struct(seed, authUser, authUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUser.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = authUser.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthUsersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUser := &AuthUser{}
	if err = randomize.Struct(seed, authUser, authUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUser.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = AuthUsers(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := AuthUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthUsersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUser := &AuthUser{}
	if err = randomize.Struct(seed, authUser, authUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUser.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AuthUserSlice{authUser}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testAuthUsersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUser := &AuthUser{}
	if err = randomize.Struct(seed, authUser, authUserDBTypes, true, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUser.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := AuthUserExists(tx, authUser.AuthUserID)
	if err != nil {
		t.Errorf("Unable to check if AuthUser exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AuthUserExistsG to return true, but got false.")
	}
}
func testAuthUsersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUser := &AuthUser{}
	if err = randomize.Struct(seed, authUser, authUserDBTypes, true, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUser.Insert(tx); err != nil {
		t.Error(err)
	}

	authUserFound, err := FindAuthUser(tx, authUser.AuthUserID)
	if err != nil {
		t.Error(err)
	}

	if authUserFound == nil {
		t.Error("want a record, got nil")
	}
}
func testAuthUsersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUser := &AuthUser{}
	if err = randomize.Struct(seed, authUser, authUserDBTypes, true, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUser.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = AuthUsers(tx).Bind(authUser); err != nil {
		t.Error(err)
	}
}

func testAuthUsersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUser := &AuthUser{}
	if err = randomize.Struct(seed, authUser, authUserDBTypes, true, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUser.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := AuthUsers(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAuthUsersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUserOne := &AuthUser{}
	authUserTwo := &AuthUser{}
	if err = randomize.Struct(seed, authUserOne, authUserDBTypes, false, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}
	if err = randomize.Struct(seed, authUserTwo, authUserDBTypes, false, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = authUserTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := AuthUsers(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAuthUsersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	authUserOne := &AuthUser{}
	authUserTwo := &AuthUser{}
	if err = randomize.Struct(seed, authUserOne, authUserDBTypes, false, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}
	if err = randomize.Struct(seed, authUserTwo, authUserDBTypes, false, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUserOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = authUserTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func authUserBeforeInsertHook(e boil.Executor, o *AuthUser) error {
	*o = AuthUser{}
	return nil
}

func authUserAfterInsertHook(e boil.Executor, o *AuthUser) error {
	*o = AuthUser{}
	return nil
}

func authUserAfterSelectHook(e boil.Executor, o *AuthUser) error {
	*o = AuthUser{}
	return nil
}

func authUserBeforeUpdateHook(e boil.Executor, o *AuthUser) error {
	*o = AuthUser{}
	return nil
}

func authUserAfterUpdateHook(e boil.Executor, o *AuthUser) error {
	*o = AuthUser{}
	return nil
}

func authUserBeforeDeleteHook(e boil.Executor, o *AuthUser) error {
	*o = AuthUser{}
	return nil
}

func authUserAfterDeleteHook(e boil.Executor, o *AuthUser) error {
	*o = AuthUser{}
	return nil
}

func authUserBeforeUpsertHook(e boil.Executor, o *AuthUser) error {
	*o = AuthUser{}
	return nil
}

func authUserAfterUpsertHook(e boil.Executor, o *AuthUser) error {
	*o = AuthUser{}
	return nil
}

func testAuthUsersHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &AuthUser{}
	o := &AuthUser{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, authUserDBTypes, false); err != nil {
		t.Errorf("Unable to randomize AuthUser object: %s", err)
	}

	AddAuthUserHook(boil.BeforeInsertHook, authUserBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	authUserBeforeInsertHooks = []AuthUserHook{}

	AddAuthUserHook(boil.AfterInsertHook, authUserAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	authUserAfterInsertHooks = []AuthUserHook{}

	AddAuthUserHook(boil.AfterSelectHook, authUserAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	authUserAfterSelectHooks = []AuthUserHook{}

	AddAuthUserHook(boil.BeforeUpdateHook, authUserBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	authUserBeforeUpdateHooks = []AuthUserHook{}

	AddAuthUserHook(boil.AfterUpdateHook, authUserAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	authUserAfterUpdateHooks = []AuthUserHook{}

	AddAuthUserHook(boil.BeforeDeleteHook, authUserBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	authUserBeforeDeleteHooks = []AuthUserHook{}

	AddAuthUserHook(boil.AfterDeleteHook, authUserAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	authUserAfterDeleteHooks = []AuthUserHook{}

	AddAuthUserHook(boil.BeforeUpsertHook, authUserBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	authUserBeforeUpsertHooks = []AuthUserHook{}

	AddAuthUserHook(boil.AfterUpsertHook, authUserAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	authUserAfterUpsertHooks = []AuthUserHook{}
}
func testAuthUsersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUser := &AuthUser{}
	if err = randomize.Struct(seed, authUser, authUserDBTypes, true, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUser.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAuthUsersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUser := &AuthUser{}
	if err = randomize.Struct(seed, authUser, authUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUser.Insert(tx, authUserColumns...); err != nil {
		t.Error(err)
	}

	count, err := AuthUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAuthUserToManyUserStockOrders(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a AuthUser
	var b, c StockOrder

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authUserDBTypes, true, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, stockOrderDBTypes, false, stockOrderColumnsWithDefault...)
	randomize.Struct(seed, &c, stockOrderDBTypes, false, stockOrderColumnsWithDefault...)

	b.UserID = a.AuthUserID
	c.UserID = a.AuthUserID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	stockOrder, err := a.UserStockOrders(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range stockOrder {
		if v.UserID == b.UserID {
			bFound = true
		}
		if v.UserID == c.UserID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AuthUserSlice{&a}
	if err = a.L.LoadUserStockOrders(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserStockOrders); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.UserStockOrders = nil
	if err = a.L.LoadUserStockOrders(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserStockOrders); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", stockOrder)
	}
}

func testAuthUserToManyAuthUserInfos(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a AuthUser
	var b, c AuthUserInfo

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authUserDBTypes, true, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, authUserInfoDBTypes, false, authUserInfoColumnsWithDefault...)
	randomize.Struct(seed, &c, authUserInfoDBTypes, false, authUserInfoColumnsWithDefault...)

	b.AuthUserID = a.AuthUserID
	c.AuthUserID = a.AuthUserID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	authUserInfo, err := a.AuthUserInfos(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range authUserInfo {
		if v.AuthUserID == b.AuthUserID {
			bFound = true
		}
		if v.AuthUserID == c.AuthUserID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AuthUserSlice{&a}
	if err = a.L.LoadAuthUserInfos(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AuthUserInfos); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.AuthUserInfos = nil
	if err = a.L.LoadAuthUserInfos(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AuthUserInfos); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", authUserInfo)
	}
}

func testAuthUserToManyAuthUserRoles(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a AuthUser
	var b, c AuthUserRole

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authUserDBTypes, true, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, authUserRoleDBTypes, false, authUserRoleColumnsWithDefault...)
	randomize.Struct(seed, &c, authUserRoleDBTypes, false, authUserRoleColumnsWithDefault...)

	b.AuthUserID = a.AuthUserID
	c.AuthUserID = a.AuthUserID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	authUserRole, err := a.AuthUserRoles(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range authUserRole {
		if v.AuthUserID == b.AuthUserID {
			bFound = true
		}
		if v.AuthUserID == c.AuthUserID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AuthUserSlice{&a}
	if err = a.L.LoadAuthUserRoles(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AuthUserRoles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.AuthUserRoles = nil
	if err = a.L.LoadAuthUserRoles(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AuthUserRoles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", authUserRole)
	}
}

func testAuthUserToManyAddOpUserStockOrders(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a AuthUser
	var b, c, d, e StockOrder

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authUserDBTypes, false, strmangle.SetComplement(authUserPrimaryKeyColumns, authUserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*StockOrder{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, stockOrderDBTypes, false, strmangle.SetComplement(stockOrderPrimaryKeyColumns, stockOrderColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*StockOrder{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUserStockOrders(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.AuthUserID != first.UserID {
			t.Error("foreign key was wrong value", a.AuthUserID, first.UserID)
		}
		if a.AuthUserID != second.UserID {
			t.Error("foreign key was wrong value", a.AuthUserID, second.UserID)
		}

		if first.R.User != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.User != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.UserStockOrders[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.UserStockOrders[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.UserStockOrders(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testAuthUserToManyAddOpAuthUserInfos(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a AuthUser
	var b, c, d, e AuthUserInfo

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authUserDBTypes, false, strmangle.SetComplement(authUserPrimaryKeyColumns, authUserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*AuthUserInfo{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, authUserInfoDBTypes, false, strmangle.SetComplement(authUserInfoPrimaryKeyColumns, authUserInfoColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*AuthUserInfo{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddAuthUserInfos(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.AuthUserID != first.AuthUserID {
			t.Error("foreign key was wrong value", a.AuthUserID, first.AuthUserID)
		}
		if a.AuthUserID != second.AuthUserID {
			t.Error("foreign key was wrong value", a.AuthUserID, second.AuthUserID)
		}

		if first.R.AuthUser != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.AuthUser != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.AuthUserInfos[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.AuthUserInfos[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.AuthUserInfos(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testAuthUserToManyAddOpAuthUserRoles(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a AuthUser
	var b, c, d, e AuthUserRole

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authUserDBTypes, false, strmangle.SetComplement(authUserPrimaryKeyColumns, authUserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*AuthUserRole{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, authUserRoleDBTypes, false, strmangle.SetComplement(authUserRolePrimaryKeyColumns, authUserRoleColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*AuthUserRole{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddAuthUserRoles(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.AuthUserID != first.AuthUserID {
			t.Error("foreign key was wrong value", a.AuthUserID, first.AuthUserID)
		}
		if a.AuthUserID != second.AuthUserID {
			t.Error("foreign key was wrong value", a.AuthUserID, second.AuthUserID)
		}

		if first.R.AuthUser != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.AuthUser != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.AuthUserRoles[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.AuthUserRoles[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.AuthUserRoles(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testAuthUsersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUser := &AuthUser{}
	if err = randomize.Struct(seed, authUser, authUserDBTypes, true, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUser.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = authUser.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testAuthUsersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUser := &AuthUser{}
	if err = randomize.Struct(seed, authUser, authUserDBTypes, true, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUser.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AuthUserSlice{authUser}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testAuthUsersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authUser := &AuthUser{}
	if err = randomize.Struct(seed, authUser, authUserDBTypes, true, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUser.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := AuthUsers(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	authUserDBTypes = map[string]string{"AuthUserID": "integer", "CreatedAt": "timestamp with time zone", "Email": "USER-DEFINED", "FirstName": "text", "IsActive": "boolean", "LastName": "text", "UpdatedAt": "timestamp with time zone"}
	_               = bytes.MinRead
)

func testAuthUsersUpdate(t *testing.T) {
	t.Parallel()

	if len(authUserColumns) == len(authUserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	authUser := &AuthUser{}
	if err = randomize.Struct(seed, authUser, authUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUser.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, authUser, authUserDBTypes, true, authUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	if err = authUser.Update(tx); err != nil {
		t.Error(err)
	}
}

func testAuthUsersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(authUserColumns) == len(authUserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	authUser := &AuthUser{}
	if err = randomize.Struct(seed, authUser, authUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUser.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AuthUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, authUser, authUserDBTypes, true, authUserPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(authUserColumns, authUserPrimaryKeyColumns) {
		fields = authUserColumns
	} else {
		fields = strmangle.SetComplement(
			authUserColumns,
			authUserPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(authUser))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := AuthUserSlice{authUser}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testAuthUsersUpsert(t *testing.T) {
	t.Parallel()

	if len(authUserColumns) == len(authUserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	authUser := AuthUser{}
	if err = randomize.Struct(seed, &authUser, authUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = authUser.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert AuthUser: %s", err)
	}

	count, err := AuthUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &authUser, authUserDBTypes, false, authUserPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AuthUser struct: %s", err)
	}

	if err = authUser.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert AuthUser: %s", err)
	}

	count, err = AuthUsers(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
