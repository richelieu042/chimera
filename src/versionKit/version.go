package versionKit

import (
	"sort"

	"github.com/hashicorp/go-version"
	"github.com/richelieu042/chimera/v3/src/core/errorKit"
	"github.com/richelieu042/chimera/v3/src/core/strKit"
)

// NewVersion
/*
@param v	(1) 不能为""
			(2) e.g. 	""1.0.0-alpha.1+build.123""
						"1.0.0-alpha+001"
						"1.3.10+meta"
						"v1.3.10+meta"
*/
func NewVersion(v string) (*version.Version, error) {
	if err := strKit.AssertNotEmpty(v, "v"); err != nil {
		return nil, err
	}

	return version.NewVersion(v)
}

// CheckConstraint 判断范围...
/*
@param versionStr		e.g. "1.2"
@param constraintStr	e.g. ">= 1.0, < 1.4"
*/
func CheckConstraint(versionStr, constraintStr string) (bool, error) {
	v, err := NewVersion(versionStr)
	if err != nil {
		return false, err
	}
	constraint, err := version.NewConstraint(constraintStr)
	if err != nil {
		return false, err
	}

	return constraint.Check(v), nil
}

// Sort 排序.
func Sort(s []string) ([]*version.Version, error) {
	versions := make([]*version.Version, len(s))
	for i, ele := range s {
		v, err := version.NewVersion(ele)
		if err != nil {
			return nil, errorKit.Wrapf(err, "ele(index: %d, value: %s) of param s is invalid", i, ele)
		}
		versions[i] = v
	}

	// After this, the versions are properly sorted
	sort.Sort(version.Collection(versions))
	return versions, nil
}
