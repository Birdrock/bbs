// Code generated by counterfeiter. DO NOT EDIT.
package migrationfakes

import (
	"database/sql"
	"sync"

	"code.cloudfoundry.org/bbs/encryption"
	"code.cloudfoundry.org/bbs/migration"
	"code.cloudfoundry.org/clock"
	lager "code.cloudfoundry.org/lager/v3"
)

type FakeMigration struct {
	SetClockStub        func(clock.Clock)
	setClockMutex       sync.RWMutex
	setClockArgsForCall []struct {
		arg1 clock.Clock
	}
	SetCryptorStub        func(encryption.Cryptor)
	setCryptorMutex       sync.RWMutex
	setCryptorArgsForCall []struct {
		arg1 encryption.Cryptor
	}
	SetDBFlavorStub        func(string)
	setDBFlavorMutex       sync.RWMutex
	setDBFlavorArgsForCall []struct {
		arg1 string
	}
	SetRawSQLDBStub        func(*sql.DB)
	setRawSQLDBMutex       sync.RWMutex
	setRawSQLDBArgsForCall []struct {
		arg1 *sql.DB
	}
	StringStub        func() string
	stringMutex       sync.RWMutex
	stringArgsForCall []struct {
	}
	stringReturns struct {
		result1 string
	}
	stringReturnsOnCall map[int]struct {
		result1 string
	}
	UpStub        func(lager.Logger) error
	upMutex       sync.RWMutex
	upArgsForCall []struct {
		arg1 lager.Logger
	}
	upReturns struct {
		result1 error
	}
	upReturnsOnCall map[int]struct {
		result1 error
	}
	VersionStub        func() int64
	versionMutex       sync.RWMutex
	versionArgsForCall []struct {
	}
	versionReturns struct {
		result1 int64
	}
	versionReturnsOnCall map[int]struct {
		result1 int64
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMigration) SetClock(arg1 clock.Clock) {
	fake.setClockMutex.Lock()
	fake.setClockArgsForCall = append(fake.setClockArgsForCall, struct {
		arg1 clock.Clock
	}{arg1})
	stub := fake.SetClockStub
	fake.recordInvocation("SetClock", []interface{}{arg1})
	fake.setClockMutex.Unlock()
	if stub != nil {
		fake.SetClockStub(arg1)
	}
}

func (fake *FakeMigration) SetClockCallCount() int {
	fake.setClockMutex.RLock()
	defer fake.setClockMutex.RUnlock()
	return len(fake.setClockArgsForCall)
}

func (fake *FakeMigration) SetClockCalls(stub func(clock.Clock)) {
	fake.setClockMutex.Lock()
	defer fake.setClockMutex.Unlock()
	fake.SetClockStub = stub
}

func (fake *FakeMigration) SetClockArgsForCall(i int) clock.Clock {
	fake.setClockMutex.RLock()
	defer fake.setClockMutex.RUnlock()
	argsForCall := fake.setClockArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMigration) SetCryptor(arg1 encryption.Cryptor) {
	fake.setCryptorMutex.Lock()
	fake.setCryptorArgsForCall = append(fake.setCryptorArgsForCall, struct {
		arg1 encryption.Cryptor
	}{arg1})
	stub := fake.SetCryptorStub
	fake.recordInvocation("SetCryptor", []interface{}{arg1})
	fake.setCryptorMutex.Unlock()
	if stub != nil {
		fake.SetCryptorStub(arg1)
	}
}

func (fake *FakeMigration) SetCryptorCallCount() int {
	fake.setCryptorMutex.RLock()
	defer fake.setCryptorMutex.RUnlock()
	return len(fake.setCryptorArgsForCall)
}

func (fake *FakeMigration) SetCryptorCalls(stub func(encryption.Cryptor)) {
	fake.setCryptorMutex.Lock()
	defer fake.setCryptorMutex.Unlock()
	fake.SetCryptorStub = stub
}

func (fake *FakeMigration) SetCryptorArgsForCall(i int) encryption.Cryptor {
	fake.setCryptorMutex.RLock()
	defer fake.setCryptorMutex.RUnlock()
	argsForCall := fake.setCryptorArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMigration) SetDBFlavor(arg1 string) {
	fake.setDBFlavorMutex.Lock()
	fake.setDBFlavorArgsForCall = append(fake.setDBFlavorArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.SetDBFlavorStub
	fake.recordInvocation("SetDBFlavor", []interface{}{arg1})
	fake.setDBFlavorMutex.Unlock()
	if stub != nil {
		fake.SetDBFlavorStub(arg1)
	}
}

func (fake *FakeMigration) SetDBFlavorCallCount() int {
	fake.setDBFlavorMutex.RLock()
	defer fake.setDBFlavorMutex.RUnlock()
	return len(fake.setDBFlavorArgsForCall)
}

func (fake *FakeMigration) SetDBFlavorCalls(stub func(string)) {
	fake.setDBFlavorMutex.Lock()
	defer fake.setDBFlavorMutex.Unlock()
	fake.SetDBFlavorStub = stub
}

func (fake *FakeMigration) SetDBFlavorArgsForCall(i int) string {
	fake.setDBFlavorMutex.RLock()
	defer fake.setDBFlavorMutex.RUnlock()
	argsForCall := fake.setDBFlavorArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMigration) SetRawSQLDB(arg1 *sql.DB) {
	fake.setRawSQLDBMutex.Lock()
	fake.setRawSQLDBArgsForCall = append(fake.setRawSQLDBArgsForCall, struct {
		arg1 *sql.DB
	}{arg1})
	stub := fake.SetRawSQLDBStub
	fake.recordInvocation("SetRawSQLDB", []interface{}{arg1})
	fake.setRawSQLDBMutex.Unlock()
	if stub != nil {
		fake.SetRawSQLDBStub(arg1)
	}
}

func (fake *FakeMigration) SetRawSQLDBCallCount() int {
	fake.setRawSQLDBMutex.RLock()
	defer fake.setRawSQLDBMutex.RUnlock()
	return len(fake.setRawSQLDBArgsForCall)
}

func (fake *FakeMigration) SetRawSQLDBCalls(stub func(*sql.DB)) {
	fake.setRawSQLDBMutex.Lock()
	defer fake.setRawSQLDBMutex.Unlock()
	fake.SetRawSQLDBStub = stub
}

func (fake *FakeMigration) SetRawSQLDBArgsForCall(i int) *sql.DB {
	fake.setRawSQLDBMutex.RLock()
	defer fake.setRawSQLDBMutex.RUnlock()
	argsForCall := fake.setRawSQLDBArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMigration) String() string {
	fake.stringMutex.Lock()
	ret, specificReturn := fake.stringReturnsOnCall[len(fake.stringArgsForCall)]
	fake.stringArgsForCall = append(fake.stringArgsForCall, struct {
	}{})
	stub := fake.StringStub
	fakeReturns := fake.stringReturns
	fake.recordInvocation("String", []interface{}{})
	fake.stringMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMigration) StringCallCount() int {
	fake.stringMutex.RLock()
	defer fake.stringMutex.RUnlock()
	return len(fake.stringArgsForCall)
}

func (fake *FakeMigration) StringCalls(stub func() string) {
	fake.stringMutex.Lock()
	defer fake.stringMutex.Unlock()
	fake.StringStub = stub
}

func (fake *FakeMigration) StringReturns(result1 string) {
	fake.stringMutex.Lock()
	defer fake.stringMutex.Unlock()
	fake.StringStub = nil
	fake.stringReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeMigration) StringReturnsOnCall(i int, result1 string) {
	fake.stringMutex.Lock()
	defer fake.stringMutex.Unlock()
	fake.StringStub = nil
	if fake.stringReturnsOnCall == nil {
		fake.stringReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.stringReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeMigration) Up(arg1 lager.Logger) error {
	fake.upMutex.Lock()
	ret, specificReturn := fake.upReturnsOnCall[len(fake.upArgsForCall)]
	fake.upArgsForCall = append(fake.upArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	stub := fake.UpStub
	fakeReturns := fake.upReturns
	fake.recordInvocation("Up", []interface{}{arg1})
	fake.upMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMigration) UpCallCount() int {
	fake.upMutex.RLock()
	defer fake.upMutex.RUnlock()
	return len(fake.upArgsForCall)
}

func (fake *FakeMigration) UpCalls(stub func(lager.Logger) error) {
	fake.upMutex.Lock()
	defer fake.upMutex.Unlock()
	fake.UpStub = stub
}

func (fake *FakeMigration) UpArgsForCall(i int) lager.Logger {
	fake.upMutex.RLock()
	defer fake.upMutex.RUnlock()
	argsForCall := fake.upArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMigration) UpReturns(result1 error) {
	fake.upMutex.Lock()
	defer fake.upMutex.Unlock()
	fake.UpStub = nil
	fake.upReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMigration) UpReturnsOnCall(i int, result1 error) {
	fake.upMutex.Lock()
	defer fake.upMutex.Unlock()
	fake.UpStub = nil
	if fake.upReturnsOnCall == nil {
		fake.upReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.upReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeMigration) Version() int64 {
	fake.versionMutex.Lock()
	ret, specificReturn := fake.versionReturnsOnCall[len(fake.versionArgsForCall)]
	fake.versionArgsForCall = append(fake.versionArgsForCall, struct {
	}{})
	stub := fake.VersionStub
	fakeReturns := fake.versionReturns
	fake.recordInvocation("Version", []interface{}{})
	fake.versionMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMigration) VersionCallCount() int {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	return len(fake.versionArgsForCall)
}

func (fake *FakeMigration) VersionCalls(stub func() int64) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = stub
}

func (fake *FakeMigration) VersionReturns(result1 int64) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = nil
	fake.versionReturns = struct {
		result1 int64
	}{result1}
}

func (fake *FakeMigration) VersionReturnsOnCall(i int, result1 int64) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = nil
	if fake.versionReturnsOnCall == nil {
		fake.versionReturnsOnCall = make(map[int]struct {
			result1 int64
		})
	}
	fake.versionReturnsOnCall[i] = struct {
		result1 int64
	}{result1}
}

func (fake *FakeMigration) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.setClockMutex.RLock()
	defer fake.setClockMutex.RUnlock()
	fake.setCryptorMutex.RLock()
	defer fake.setCryptorMutex.RUnlock()
	fake.setDBFlavorMutex.RLock()
	defer fake.setDBFlavorMutex.RUnlock()
	fake.setRawSQLDBMutex.RLock()
	defer fake.setRawSQLDBMutex.RUnlock()
	fake.stringMutex.RLock()
	defer fake.stringMutex.RUnlock()
	fake.upMutex.RLock()
	defer fake.upMutex.RUnlock()
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeMigration) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ migration.Migration = new(FakeMigration)
