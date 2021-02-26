/*
 * SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company and Gardener contributors
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */

package fieldpath

import (
	"reflect"
)

type value interface {
	Type() reflect.Type
	Interface() interface{}
	Kind() reflect.Kind
	Set(reflect.Value) value
	IsValid() bool
	Len() int
	IsNil() bool
	Elem() value

	Value() reflect.Value
}

type reflectValue reflect.Value

var _ value = reflectValue{}

func (this reflectValue) Type() reflect.Type {
	return reflect.Value(this).Type()
}
func (this reflectValue) Interface() interface{} {
	return reflect.Value(this).Interface()
}
func (this reflectValue) Kind() reflect.Kind {
	return reflect.Value(this).Kind()
}
func (this reflectValue) Set(v reflect.Value) value {
	reflect.Value(this).Set(v)
	return this
}
func (this reflectValue) IsValid() bool {
	return reflect.Value(this).IsValid()
}
func (this reflectValue) Len() int {
	return reflect.Value(this).Len()
}
func (this reflectValue) IsNil() bool {
	return reflect.Value(this).IsNil()
}
func (this reflectValue) Elem() value {
	return reflectValue(reflect.Value(this).Elem())
}
func (this reflectValue) Value() reflect.Value {
	return reflect.Value(this)
}

type mapEntry struct {
	host reflect.Value
	key  reflect.Value
	elem *reflect.Value // effective element for further processing
	name string
}

var _ value = mapEntry{}

func (this mapEntry) Type() reflect.Type {
	if this.elem == nil {
		return this.host.Type().Elem()
	}
	return this.elem.Type()
}
func (this mapEntry) Interface() interface{} {
	if this.elem == nil {
		return this.host.MapIndex(this.key).Interface()
	}
	return this.elem.Interface()
}
func (this mapEntry) Kind() reflect.Kind {
	return this.Type().Kind()
}
func (this mapEntry) Set(v reflect.Value) value {
	this.elem = &v
	this.host.SetMapIndex(this.key, v)
	return this
}
func (this mapEntry) IsValid() bool {
	return this.Value().IsValid()
}
func (this mapEntry) Len() int {
	return this.Value().Len()
}
func (this mapEntry) IsNil() bool {
	return this.Value().IsNil()
}
func (this mapEntry) Value() reflect.Value {
	if this.elem == nil {
		return this.host.MapIndex(this.key)
	}
	return *this.elem
}
func (this mapEntry) Elem() value {
	var e reflect.Value
	if this.elem == nil {
		e = this.host.MapIndex(this.key)
	} else {
		e = this.elem.Elem()
	}
	this.elem = &e
	return this
}
