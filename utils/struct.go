package utils

import (
	"errors"
	"reflect"
)

// verify 简单的类型判断
func verify(dst, src any) (srcT, dstT reflect.Type, srcV, dstV reflect.Value, err error) {
	srcT, srcV = reflect.TypeOf(src), reflect.ValueOf(src)
	if srcT.Kind() == reflect.Ptr {
		srcT, srcV = srcT.Elem(), srcV.Elem()
	}
	if srcT.Kind() != reflect.Struct {
		err = errors.New("仅支持 Struct 进行合并")
		return
	}
	dstT, dstV = reflect.TypeOf(dst), reflect.ValueOf(dst)
	if dstT.Kind() != reflect.Ptr || dstT.Elem().Kind() != reflect.Struct {
		err = errors.New("dst 必须为 Struct指针")
	} else {
		dstT, dstV = dstT.Elem(), dstV.Elem()
	}
	return
}

// Merge 合并两个结构体
// 危危危,反射很危险,多测试
func Merge(dst, src any) error {
	srcT, dstT, srcV, dstV, err := verify(dst, src)
	if err != nil {
		return err
	}
	if srcV.NumField() < dstV.NumField() {
		for i := 0; i < srcV.NumField(); i++ {
			curT, curV := srcT.Field(i), srcV.Field(i)
			f := dstV.FieldByName(curT.Name)
			if f.IsValid() && curV.Type() == f.Type() {
				f.Set(curV)
			}
		}
	} else {
		for i := 0; i < dstV.NumField(); i++ {
			curT, curV := dstT.Field(i), dstV.Field(i)
			f := srcV.FieldByName(curT.Name)
			if f.IsValid() && curV.Type() == f.Type() {
				curV.Set(f)
			}
		}
	}
	return nil
}
