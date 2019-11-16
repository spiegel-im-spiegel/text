package ecode

import (
	"errors"
	"fmt"
	"testing"

	"github.com/spiegel-im-spiegel/errs"
)

func TestNumError(t *testing.T) {
	testCases := []struct {
		err error
		str string
	}{
		{err: ECode(0), str: "unknown error (0)"},
		{err: ErrNullPointer, str: "null reference instance"},
		{err: ErrNoImplement, str: "no implementation"},
		{err: ErrInvalidOption, str: "invalid option"},
		{err: ECode(4), str: "unknown error (4)"},
	}

	for _, tc := range testCases {
		errStr := tc.err.Error()
		if errStr != tc.str {
			t.Errorf("\"%v\" != \"%v\"", errStr, tc.str)
		}
		fmt.Printf("Info(TestNumError): %+v\n", tc.err)
	}
}

func TestNumErrorEquality(t *testing.T) {
	testCases := []struct {
		err1 error
		err2 error
		res  bool
	}{
		{err1: nil, err2: nil, res: true},
		{err1: nil, err2: ErrNullPointer, res: false},
		{err1: ErrNullPointer, err2: ErrNullPointer, res: true},
		{err1: ErrNullPointer, err2: errs.Wrap(ErrNullPointer, "wrapping error"), res: false},
		{err1: ErrNullPointer, err2: ECode(0), res: false},
		{err1: ErrNullPointer, err2: nil, res: false},
	}

	for _, tc := range testCases {
		res := errors.Is(tc.err1, tc.err2)
		if res != tc.res {
			t.Errorf("\"%v\" == \"%v\" ? %v, want %v", tc.err1, tc.err2, res, tc.res)
		}
	}
}

func TestWrapError(t *testing.T) {
	testCases := []struct {
		err error
		msg string
		str string
	}{
		{err: ErrNullPointer, msg: "wrapping error", str: "wrapping error: null reference instance"},
	}

	for _, tc := range testCases {
		we := errs.Wrap(tc.err, tc.msg)
		if we.Error() != tc.str {
			t.Errorf("wrapError.Error() == \"%v\", want \"%v\"", we.Error(), tc.str)
		}
		fmt.Printf("Info(TestWrapError): %+v\n", we)
	}
}

func TestWrapNilError(t *testing.T) {
	if we := errs.Wrap(nil, "null error"); we != nil {
		t.Errorf("Wrap(nil) == \"%v\", want nil.", we)
	}
}

func TestWrapfError(t *testing.T) {
	testCases := []struct {
		err error
		msg string
		str string
	}{
		{err: ErrNullPointer, msg: "wrapping error", str: "wrapping error: null reference instance"},
	}

	for _, tc := range testCases {
		we := errs.Wrap(tc.err, tc.msg)
		if we.Error() != tc.str {
			t.Errorf("wrapError.Error() == \"%v\", want \"%v\"", we.Error(), tc.str)
		}
		fmt.Printf("Info(TestWrapfError): %+v\n", we)
	}
}

func TestWrapfNilError(t *testing.T) {
	if we := errs.Wrap(nil, "null error"); we != nil {
		t.Errorf("Wrapf(nil) == \"%v\", want nil.", we)
	}
}

func TestWrapErrorEquality(t *testing.T) {
	testCases := []struct {
		err1 error
		err2 error
		res  bool
	}{
		{err1: nil, err2: errs.Wrap(ErrNullPointer, "wrapping error"), res: false},
		{err1: errs.Wrap(ErrNullPointer, "wrapping error"), err2: ErrNullPointer, res: true},
		{err1: errs.Wrap(ErrNullPointer, "wrapping error"), err2: errs.Wrap(ECode(0), "wrapping error"), res: false},
		{err1: errs.Wrap(ErrNullPointer, "wrapping error"), err2: ECode(0), res: false},
		{err1: errs.Wrap(ErrNullPointer, "wrapping error"), err2: nil, res: false},
	}

	for _, tc := range testCases {
		res := errors.Is(tc.err1, tc.err2)
		if res != tc.res {
			t.Errorf("\"%v\" == \"%v\" ? %v, want %v", tc.err1, tc.err2, res, tc.res)
		}
	}
}

/* MIT License
 *
 * Copyright 2019 Spiegel
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
