package controller

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"reflect"
	"strconv"
)

var Validator *validator.Validate

func init() {
	Validator = validator.New()
}

type notPtrError struct {
	Object interface{}
}

type badArgsCount struct {
	structArgs int
	inputArgs  int
}

func (b badArgsCount) Error() string {
	return fmt.Sprintf("needed args count=%d, actual=%d", b.structArgs, b.inputArgs)
}

func (err notPtrError) Error() string {
	return fmt.Sprintf("%#v isn't ptr", err.Object)
}

func FillAndValidateStruct(object interface{}, args []string) error {
	err := CreateFromStringArgs(args, object)
	if err != nil {
		return err
	}
	err = Validator.Struct(args)
	if err != nil {
		return err
	}
	return nil
}

func CreateFromStringArgs(args []string, obj interface{}) error {
	val := reflect.ValueOf(obj)
	if val.Kind() != reflect.Ptr {
		logrus.Error(notPtrError{obj})
		return errors.New("Внутренняя ошибка бота.")
	}
	structVal := val.Elem()
	if structVal.NumField() > len(args) {
		logrus.Error(badArgsCount{
			structArgs: structVal.NumField(),
			inputArgs:  len(args),
		})
		return errors.New(fmt.Sprintf("Недостаточно аргументов, нужно %d, получено %d", structVal.NumField(), len(args)))
	}

	numFields := structVal.NumField()
	for i := 0; i < numFields; i++ {
		field := structVal.Field(i)
		if !field.CanSet() {
			logrus.Error(errors.New("невозможно поменять поле " + structVal.String()))
			return errors.New("Внутренняя ошибка бота.")
		}
		switch field.Kind() {
		case reflect.Int:
			intVal, err := strconv.Atoi(args[i])
			if err != nil {
				return errors.New(fmt.Sprintf("%s - не число.", args[i]))
			}
			field.SetInt(int64(intVal))
		case reflect.String:
			field.SetString(args[i])
		case reflect.Float32:
			floatVal, err := strconv.ParseFloat(args[i], 32)
			if err != nil {
				return errors.New(fmt.Sprintf("%s - не число.", args[i]))
			}
			field.SetFloat(floatVal)
		}
	}

	return nil
}
