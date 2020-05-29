// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	internal "github.com/wal-g/wal-g/internal"
	archive "github.com/wal-g/wal-g/internal/databases/mongo/archive"

	io "io"

	mock "github.com/stretchr/testify/mock"

	models "github.com/wal-g/wal-g/internal/databases/mongo/models"
)

// Downloader is an autogenerated mock type for the Downloader type
type Downloader struct {
	mock.Mock
}

// BackupMeta provides a mock function with given fields: name
func (_m *Downloader) BackupMeta(name string) (archive.Backup, error) {
	ret := _m.Called(name)

	var r0 archive.Backup
	if rf, ok := ret.Get(0).(func(string) archive.Backup); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(archive.Backup)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DownloadOplogArchive provides a mock function with given fields: arch, writeCloser
func (_m *Downloader) DownloadOplogArchive(arch models.Archive, writeCloser io.WriteCloser) error {
	ret := _m.Called(arch, writeCloser)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Archive, io.WriteCloser) error); ok {
		r0 = rf(arch, writeCloser)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListBackupNames provides a mock function with given fields:
func (_m *Downloader) ListBackupNames() ([]internal.BackupTime, error) {
	ret := _m.Called()

	var r0 []internal.BackupTime
	if rf, ok := ret.Get(0).(func() []internal.BackupTime); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]internal.BackupTime)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListOplogArchives provides a mock function with given fields:
func (_m *Downloader) ListOplogArchives() ([]models.Archive, error) {
	ret := _m.Called()

	var r0 []models.Archive
	if rf, ok := ret.Get(0).(func() []models.Archive); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Archive)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadBackups provides a mock function with given fields: names
func (_m *Downloader) LoadBackups(names []string) ([]archive.Backup, error) {
	ret := _m.Called(names)

	var r0 []archive.Backup
	if rf, ok := ret.Get(0).(func([]string) []archive.Backup); ok {
		r0 = rf(names)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]archive.Backup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(names)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
