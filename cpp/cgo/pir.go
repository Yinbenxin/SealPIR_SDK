/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (https://www.swig.org).
 * Version 4.1.1
 *
 * Do not make changes to this file unless you know what you are doing - modify
 * the SWIG interface file instead.
 * ----------------------------------------------------------------------------- */

// source: /Users/admin/git/Pir-SDK/cpp/swig-go/pir.i

package pir

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stddef.h>
#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


typedef long long swig_type_1;
typedef long long swig_type_2;
typedef long long swig_type_3;
typedef long long swig_type_4;
typedef long long swig_type_5;
typedef long long swig_type_6;
typedef long long swig_type_7;
typedef long long swig_type_8;
extern void _wrap_Swig_free_pir_c6eb00fb9a209540(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_pir_c6eb00fb9a209540(swig_intgo arg1);
extern uintptr_t _wrap_new_Faker_pir_c6eb00fb9a209540(void);
extern uintptr_t _wrap_Faker_GenerateCNIDCardNumber__SWIG_0_pir_c6eb00fb9a209540(uintptr_t arg1, swig_intgo arg2);
extern uintptr_t _wrap_Faker_GenerateCNIDCardNumber__SWIG_1_pir_c6eb00fb9a209540(uintptr_t arg1);
extern uintptr_t _wrap_Faker_GenerateTWCompatriotNumber__SWIG_0_pir_c6eb00fb9a209540(uintptr_t arg1, swig_intgo arg2);
extern uintptr_t _wrap_Faker_GenerateTWCompatriotNumber__SWIG_1_pir_c6eb00fb9a209540(uintptr_t arg1);
extern uintptr_t _wrap_Faker_GenerateTWIDCardNumber__SWIG_0_pir_c6eb00fb9a209540(uintptr_t arg1, swig_intgo arg2);
extern uintptr_t _wrap_Faker_GenerateTWIDCardNumber__SWIG_1_pir_c6eb00fb9a209540(uintptr_t arg1);
extern uintptr_t _wrap_Faker_GenerateTW2CNIDCardNumber__SWIG_0_pir_c6eb00fb9a209540(uintptr_t arg1, swig_intgo arg2);
extern uintptr_t _wrap_Faker_GenerateTW2CNIDCardNumber__SWIG_1_pir_c6eb00fb9a209540(uintptr_t arg1);
extern uintptr_t _wrap_Faker_GeneratePhoneNumber__SWIG_0_pir_c6eb00fb9a209540(uintptr_t arg1, swig_intgo arg2);
extern uintptr_t _wrap_Faker_GeneratePhoneNumber__SWIG_1_pir_c6eb00fb9a209540(uintptr_t arg1);
extern uintptr_t _wrap_Faker_GenerateMACAddress__SWIG_0_pir_c6eb00fb9a209540(uintptr_t arg1, swig_intgo arg2);
extern uintptr_t _wrap_Faker_GenerateMACAddress__SWIG_1_pir_c6eb00fb9a209540(uintptr_t arg1);
extern uintptr_t _wrap_Faker_GenerateIMEI__SWIG_0_pir_c6eb00fb9a209540(uintptr_t arg1, swig_intgo arg2);
extern uintptr_t _wrap_Faker_GenerateIMEI__SWIG_1_pir_c6eb00fb9a209540(uintptr_t arg1);
extern uintptr_t _wrap_Faker_GenerateIMSI__SWIG_0_pir_c6eb00fb9a209540(uintptr_t arg1, swig_intgo arg2);
extern uintptr_t _wrap_Faker_GenerateIMSI__SWIG_1_pir_c6eb00fb9a209540(uintptr_t arg1);
extern uintptr_t _wrap_Faker_GenerateEmail__SWIG_0_pir_c6eb00fb9a209540(uintptr_t arg1, swig_intgo arg2);
extern uintptr_t _wrap_Faker_GenerateEmail__SWIG_1_pir_c6eb00fb9a209540(uintptr_t arg1);
extern uintptr_t _wrap_Faker_GeneratePassportNumber__SWIG_0_pir_c6eb00fb9a209540(uintptr_t arg1, swig_intgo arg2);
extern uintptr_t _wrap_Faker_GeneratePassportNumber__SWIG_1_pir_c6eb00fb9a209540(uintptr_t arg1);
extern void _wrap_delete_Faker_pir_c6eb00fb9a209540(uintptr_t arg1);
extern uintptr_t _wrap_new_PirClient_pir_c6eb00fb9a209540(swig_intgo arg1, uintptr_t arg2, _Bool arg3, swig_voidp arg4);
extern void _wrap_delete_PirClient_pir_c6eb00fb9a209540(uintptr_t arg1);
extern swig_intgo _wrap_PirClient_generate_data_query_pir_c6eb00fb9a209540(uintptr_t arg1, uintptr_t arg2, swig_intgo arg3, uintptr_t arg4);
extern swig_intgo _wrap_PirClient_decrypt_data_reply_pir_c6eb00fb9a209540(uintptr_t arg1, uintptr_t arg2, uintptr_t arg3);
extern swig_intgo _wrap_PirClient_generate_data_query_with_count_pir_c6eb00fb9a209540(uintptr_t arg1, uintptr_t arg2, swig_intgo arg3, uintptr_t arg4);
extern swig_intgo _wrap_PirClient_get_retrieved_indexes_and_ciphertext_pir_c6eb00fb9a209540(uintptr_t arg1, uintptr_t arg2, uintptr_t arg3, swig_intgo arg4, uintptr_t arg5, uintptr_t arg6);
extern swig_intgo _wrap_PirClient_generate_key_query_of_retrieved_indexes_pir_c6eb00fb9a209540(uintptr_t arg1, uintptr_t arg2, uintptr_t arg3);
extern swig_intgo _wrap_PirClient_decrypt_retrieved_ciphertext_pir_c6eb00fb9a209540(uintptr_t arg1, uintptr_t arg2, uintptr_t arg3, uintptr_t arg4);
extern swig_intgo _wrap_PirClient_set_key_pir_c6eb00fb9a209540(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_PirClient_get_private_key_pir_c6eb00fb9a209540(uintptr_t arg1);
extern uintptr_t _wrap_PirClient_get_public_key_pir_c6eb00fb9a209540(uintptr_t arg1);
extern uintptr_t _wrap_new_PirServer_pir_c6eb00fb9a209540(swig_intgo arg1, uintptr_t arg2, _Bool arg3, swig_voidp arg4);
extern void _wrap_delete_PirServer_pir_c6eb00fb9a209540(uintptr_t arg1);
extern swig_intgo _wrap_PirServer_generate_data_reply_pir_c6eb00fb9a209540(uintptr_t arg1, uintptr_t arg2, uintptr_t arg3, uintptr_t arg4, uintptr_t arg5);
extern swig_intgo _wrap_PirServer_generate_checked_data_reply_pir_c6eb00fb9a209540(uintptr_t arg1, uintptr_t arg2, uintptr_t arg3, uintptr_t arg4, uintptr_t arg5);
extern swig_intgo _wrap_PirServer_generate_checked_key_reply_pir_c6eb00fb9a209540(uintptr_t arg1, uintptr_t arg2, uintptr_t arg3, swig_voidp arg4);
extern swig_intgo _wrap_PirServer_set_key_pir_c6eb00fb9a209540(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_PirServer_get_private_key_pir_c6eb00fb9a209540(uintptr_t arg1);
extern uintptr_t _wrap_PirServer_get_public_key_pir_c6eb00fb9a209540(uintptr_t arg1);
extern uintptr_t _wrap_new_VectorUInt8__SWIG_0_pir_c6eb00fb9a209540(void);
extern uintptr_t _wrap_new_VectorUInt8__SWIG_1_pir_c6eb00fb9a209540(swig_type_1 arg1);
extern uintptr_t _wrap_new_VectorUInt8__SWIG_2_pir_c6eb00fb9a209540(uintptr_t arg1);
extern swig_type_2 _wrap_VectorUInt8_size_pir_c6eb00fb9a209540(uintptr_t arg1);
extern swig_type_3 _wrap_VectorUInt8_capacity_pir_c6eb00fb9a209540(uintptr_t arg1);
extern void _wrap_VectorUInt8_reserve_pir_c6eb00fb9a209540(uintptr_t arg1, swig_type_4 arg2);
extern _Bool _wrap_VectorUInt8_isEmpty_pir_c6eb00fb9a209540(uintptr_t arg1);
extern void _wrap_VectorUInt8_clear_pir_c6eb00fb9a209540(uintptr_t arg1);
extern void _wrap_VectorUInt8_add_pir_c6eb00fb9a209540(uintptr_t arg1, char arg2);
extern char _wrap_VectorUInt8_get_pir_c6eb00fb9a209540(uintptr_t arg1, swig_intgo arg2);
extern void _wrap_VectorUInt8_set_pir_c6eb00fb9a209540(uintptr_t arg1, swig_intgo arg2, char arg3);
extern void _wrap_delete_VectorUInt8_pir_c6eb00fb9a209540(uintptr_t arg1);
extern uintptr_t _wrap_new_ArrayofVectorUInt8__SWIG_0_pir_c6eb00fb9a209540(void);
extern uintptr_t _wrap_new_ArrayofVectorUInt8__SWIG_1_pir_c6eb00fb9a209540(swig_type_5 arg1);
extern uintptr_t _wrap_new_ArrayofVectorUInt8__SWIG_2_pir_c6eb00fb9a209540(uintptr_t arg1);
extern swig_type_6 _wrap_ArrayofVectorUInt8_size_pir_c6eb00fb9a209540(uintptr_t arg1);
extern swig_type_7 _wrap_ArrayofVectorUInt8_capacity_pir_c6eb00fb9a209540(uintptr_t arg1);
extern void _wrap_ArrayofVectorUInt8_reserve_pir_c6eb00fb9a209540(uintptr_t arg1, swig_type_8 arg2);
extern _Bool _wrap_ArrayofVectorUInt8_isEmpty_pir_c6eb00fb9a209540(uintptr_t arg1);
extern void _wrap_ArrayofVectorUInt8_clear_pir_c6eb00fb9a209540(uintptr_t arg1);
extern void _wrap_ArrayofVectorUInt8_add_pir_c6eb00fb9a209540(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_ArrayofVectorUInt8_get_pir_c6eb00fb9a209540(uintptr_t arg1, swig_intgo arg2);
extern void _wrap_ArrayofVectorUInt8_set_pir_c6eb00fb9a209540(uintptr_t arg1, swig_intgo arg2, uintptr_t arg3);
extern void _wrap_delete_ArrayofVectorUInt8_pir_c6eb00fb9a209540(uintptr_t arg1);
#undef intgo
*/
import "C"

import "syscall"
import "unsafe"
import "sync"


type _ syscall.Sockaddr




type _ unsafe.Pointer



var Swig_escape_always_false bool
var Swig_escape_val interface{}


type _swig_fnptr *byte
type _swig_memberptr *byte


func getSwigcptr(v interface { Swigcptr() uintptr }) uintptr {
	if v == nil {
		return 0
	}
	return v.Swigcptr()
}


type _ sync.Mutex

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_pir_c6eb00fb9a209540(C.swig_intgo(_swig_i_0)))
	return swig_r
}

type SwigcptrFaker uintptr

func (p SwigcptrFaker) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrFaker) SwigIsFaker() {
}

func NewFaker() (_swig_ret Faker) {
	var swig_r Faker
	swig_r = (Faker)(SwigcptrFaker(C._wrap_new_Faker_pir_c6eb00fb9a209540()))
	return swig_r
}

func (arg1 SwigcptrFaker) GenerateCNIDCardNumber__SWIG_0(arg2 int) (_swig_ret ArrayofVectorUInt8) {
	var swig_r ArrayofVectorUInt8
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (ArrayofVectorUInt8)(SwigcptrArrayofVectorUInt8(C._wrap_Faker_GenerateCNIDCardNumber__SWIG_0_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))))
	return swig_r
}

func (arg1 SwigcptrFaker) GenerateCNIDCardNumber__SWIG_1() (_swig_ret ArrayofVectorUInt8) {
	var swig_r ArrayofVectorUInt8
	_swig_i_0 := arg1
	swig_r = (ArrayofVectorUInt8)(SwigcptrArrayofVectorUInt8(C._wrap_Faker_GenerateCNIDCardNumber__SWIG_1_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (p SwigcptrFaker) GenerateCNIDCardNumber(a ...interface{}) ArrayofVectorUInt8 {
	argc := len(a)
	if argc == 0 {
		return p.GenerateCNIDCardNumber__SWIG_1()
	}
	if argc == 1 {
		return p.GenerateCNIDCardNumber__SWIG_0(a[0].(int))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrFaker) GenerateTWCompatriotNumber__SWIG_0(arg2 int) (_swig_ret ArrayofVectorUInt8) {
	var swig_r ArrayofVectorUInt8
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (ArrayofVectorUInt8)(SwigcptrArrayofVectorUInt8(C._wrap_Faker_GenerateTWCompatriotNumber__SWIG_0_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))))
	return swig_r
}

func (arg1 SwigcptrFaker) GenerateTWCompatriotNumber__SWIG_1() (_swig_ret ArrayofVectorUInt8) {
	var swig_r ArrayofVectorUInt8
	_swig_i_0 := arg1
	swig_r = (ArrayofVectorUInt8)(SwigcptrArrayofVectorUInt8(C._wrap_Faker_GenerateTWCompatriotNumber__SWIG_1_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (p SwigcptrFaker) GenerateTWCompatriotNumber(a ...interface{}) ArrayofVectorUInt8 {
	argc := len(a)
	if argc == 0 {
		return p.GenerateTWCompatriotNumber__SWIG_1()
	}
	if argc == 1 {
		return p.GenerateTWCompatriotNumber__SWIG_0(a[0].(int))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrFaker) GenerateTWIDCardNumber__SWIG_0(arg2 int) (_swig_ret ArrayofVectorUInt8) {
	var swig_r ArrayofVectorUInt8
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (ArrayofVectorUInt8)(SwigcptrArrayofVectorUInt8(C._wrap_Faker_GenerateTWIDCardNumber__SWIG_0_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))))
	return swig_r
}

func (arg1 SwigcptrFaker) GenerateTWIDCardNumber__SWIG_1() (_swig_ret ArrayofVectorUInt8) {
	var swig_r ArrayofVectorUInt8
	_swig_i_0 := arg1
	swig_r = (ArrayofVectorUInt8)(SwigcptrArrayofVectorUInt8(C._wrap_Faker_GenerateTWIDCardNumber__SWIG_1_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (p SwigcptrFaker) GenerateTWIDCardNumber(a ...interface{}) ArrayofVectorUInt8 {
	argc := len(a)
	if argc == 0 {
		return p.GenerateTWIDCardNumber__SWIG_1()
	}
	if argc == 1 {
		return p.GenerateTWIDCardNumber__SWIG_0(a[0].(int))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrFaker) GenerateTW2CNIDCardNumber__SWIG_0(arg2 int) (_swig_ret ArrayofVectorUInt8) {
	var swig_r ArrayofVectorUInt8
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (ArrayofVectorUInt8)(SwigcptrArrayofVectorUInt8(C._wrap_Faker_GenerateTW2CNIDCardNumber__SWIG_0_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))))
	return swig_r
}

func (arg1 SwigcptrFaker) GenerateTW2CNIDCardNumber__SWIG_1() (_swig_ret ArrayofVectorUInt8) {
	var swig_r ArrayofVectorUInt8
	_swig_i_0 := arg1
	swig_r = (ArrayofVectorUInt8)(SwigcptrArrayofVectorUInt8(C._wrap_Faker_GenerateTW2CNIDCardNumber__SWIG_1_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (p SwigcptrFaker) GenerateTW2CNIDCardNumber(a ...interface{}) ArrayofVectorUInt8 {
	argc := len(a)
	if argc == 0 {
		return p.GenerateTW2CNIDCardNumber__SWIG_1()
	}
	if argc == 1 {
		return p.GenerateTW2CNIDCardNumber__SWIG_0(a[0].(int))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrFaker) GeneratePhoneNumber__SWIG_0(arg2 int) (_swig_ret ArrayofVectorUInt8) {
	var swig_r ArrayofVectorUInt8
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (ArrayofVectorUInt8)(SwigcptrArrayofVectorUInt8(C._wrap_Faker_GeneratePhoneNumber__SWIG_0_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))))
	return swig_r
}

func (arg1 SwigcptrFaker) GeneratePhoneNumber__SWIG_1() (_swig_ret ArrayofVectorUInt8) {
	var swig_r ArrayofVectorUInt8
	_swig_i_0 := arg1
	swig_r = (ArrayofVectorUInt8)(SwigcptrArrayofVectorUInt8(C._wrap_Faker_GeneratePhoneNumber__SWIG_1_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (p SwigcptrFaker) GeneratePhoneNumber(a ...interface{}) ArrayofVectorUInt8 {
	argc := len(a)
	if argc == 0 {
		return p.GeneratePhoneNumber__SWIG_1()
	}
	if argc == 1 {
		return p.GeneratePhoneNumber__SWIG_0(a[0].(int))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrFaker) GenerateMACAddress__SWIG_0(arg2 int) (_swig_ret ArrayofVectorUInt8) {
	var swig_r ArrayofVectorUInt8
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (ArrayofVectorUInt8)(SwigcptrArrayofVectorUInt8(C._wrap_Faker_GenerateMACAddress__SWIG_0_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))))
	return swig_r
}

func (arg1 SwigcptrFaker) GenerateMACAddress__SWIG_1() (_swig_ret ArrayofVectorUInt8) {
	var swig_r ArrayofVectorUInt8
	_swig_i_0 := arg1
	swig_r = (ArrayofVectorUInt8)(SwigcptrArrayofVectorUInt8(C._wrap_Faker_GenerateMACAddress__SWIG_1_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (p SwigcptrFaker) GenerateMACAddress(a ...interface{}) ArrayofVectorUInt8 {
	argc := len(a)
	if argc == 0 {
		return p.GenerateMACAddress__SWIG_1()
	}
	if argc == 1 {
		return p.GenerateMACAddress__SWIG_0(a[0].(int))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrFaker) GenerateIMEI__SWIG_0(arg2 int) (_swig_ret ArrayofVectorUInt8) {
	var swig_r ArrayofVectorUInt8
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (ArrayofVectorUInt8)(SwigcptrArrayofVectorUInt8(C._wrap_Faker_GenerateIMEI__SWIG_0_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))))
	return swig_r
}

func (arg1 SwigcptrFaker) GenerateIMEI__SWIG_1() (_swig_ret ArrayofVectorUInt8) {
	var swig_r ArrayofVectorUInt8
	_swig_i_0 := arg1
	swig_r = (ArrayofVectorUInt8)(SwigcptrArrayofVectorUInt8(C._wrap_Faker_GenerateIMEI__SWIG_1_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (p SwigcptrFaker) GenerateIMEI(a ...interface{}) ArrayofVectorUInt8 {
	argc := len(a)
	if argc == 0 {
		return p.GenerateIMEI__SWIG_1()
	}
	if argc == 1 {
		return p.GenerateIMEI__SWIG_0(a[0].(int))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrFaker) GenerateIMSI__SWIG_0(arg2 int) (_swig_ret ArrayofVectorUInt8) {
	var swig_r ArrayofVectorUInt8
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (ArrayofVectorUInt8)(SwigcptrArrayofVectorUInt8(C._wrap_Faker_GenerateIMSI__SWIG_0_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))))
	return swig_r
}

func (arg1 SwigcptrFaker) GenerateIMSI__SWIG_1() (_swig_ret ArrayofVectorUInt8) {
	var swig_r ArrayofVectorUInt8
	_swig_i_0 := arg1
	swig_r = (ArrayofVectorUInt8)(SwigcptrArrayofVectorUInt8(C._wrap_Faker_GenerateIMSI__SWIG_1_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (p SwigcptrFaker) GenerateIMSI(a ...interface{}) ArrayofVectorUInt8 {
	argc := len(a)
	if argc == 0 {
		return p.GenerateIMSI__SWIG_1()
	}
	if argc == 1 {
		return p.GenerateIMSI__SWIG_0(a[0].(int))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrFaker) GenerateEmail__SWIG_0(arg2 int) (_swig_ret ArrayofVectorUInt8) {
	var swig_r ArrayofVectorUInt8
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (ArrayofVectorUInt8)(SwigcptrArrayofVectorUInt8(C._wrap_Faker_GenerateEmail__SWIG_0_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))))
	return swig_r
}

func (arg1 SwigcptrFaker) GenerateEmail__SWIG_1() (_swig_ret ArrayofVectorUInt8) {
	var swig_r ArrayofVectorUInt8
	_swig_i_0 := arg1
	swig_r = (ArrayofVectorUInt8)(SwigcptrArrayofVectorUInt8(C._wrap_Faker_GenerateEmail__SWIG_1_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (p SwigcptrFaker) GenerateEmail(a ...interface{}) ArrayofVectorUInt8 {
	argc := len(a)
	if argc == 0 {
		return p.GenerateEmail__SWIG_1()
	}
	if argc == 1 {
		return p.GenerateEmail__SWIG_0(a[0].(int))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrFaker) GeneratePassportNumber__SWIG_0(arg2 int) (_swig_ret ArrayofVectorUInt8) {
	var swig_r ArrayofVectorUInt8
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (ArrayofVectorUInt8)(SwigcptrArrayofVectorUInt8(C._wrap_Faker_GeneratePassportNumber__SWIG_0_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))))
	return swig_r
}

func (arg1 SwigcptrFaker) GeneratePassportNumber__SWIG_1() (_swig_ret ArrayofVectorUInt8) {
	var swig_r ArrayofVectorUInt8
	_swig_i_0 := arg1
	swig_r = (ArrayofVectorUInt8)(SwigcptrArrayofVectorUInt8(C._wrap_Faker_GeneratePassportNumber__SWIG_1_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (p SwigcptrFaker) GeneratePassportNumber(a ...interface{}) ArrayofVectorUInt8 {
	argc := len(a)
	if argc == 0 {
		return p.GeneratePassportNumber__SWIG_1()
	}
	if argc == 1 {
		return p.GeneratePassportNumber__SWIG_0(a[0].(int))
	}
	panic("No match for overloaded function call")
}

func DeleteFaker(arg1 Faker) {
	_swig_i_0 := getSwigcptr(arg1)
	C._wrap_delete_Faker_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0))
}

type Faker interface {
	Swigcptr() uintptr
	SwigIsFaker()
	GenerateCNIDCardNumber(a ...interface{}) ArrayofVectorUInt8
	GenerateTWCompatriotNumber(a ...interface{}) ArrayofVectorUInt8
	GenerateTWIDCardNumber(a ...interface{}) ArrayofVectorUInt8
	GenerateTW2CNIDCardNumber(a ...interface{}) ArrayofVectorUInt8
	GeneratePhoneNumber(a ...interface{}) ArrayofVectorUInt8
	GenerateMACAddress(a ...interface{}) ArrayofVectorUInt8
	GenerateIMEI(a ...interface{}) ArrayofVectorUInt8
	GenerateIMSI(a ...interface{}) ArrayofVectorUInt8
	GenerateEmail(a ...interface{}) ArrayofVectorUInt8
	GeneratePassportNumber(a ...interface{}) ArrayofVectorUInt8
}

type SwigcptrPirClient uintptr

func (p SwigcptrPirClient) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrPirClient) SwigIsPirClient() {
}

func NewPirClient(arg1 int, arg2 VectorUInt8, arg3 bool, arg4 *int) (_swig_ret PirClient) {
	var swig_r PirClient
	_swig_i_0 := arg1
	_swig_i_1 := getSwigcptr(arg2)
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	swig_r = (PirClient)(SwigcptrPirClient(C._wrap_new_PirClient_pir_c6eb00fb9a209540(C.swig_intgo(_swig_i_0), C.uintptr_t(_swig_i_1), C._Bool(_swig_i_2), C.swig_voidp(_swig_i_3))))
	return swig_r
}

func DeletePirClient(arg1 PirClient) {
	_swig_i_0 := getSwigcptr(arg1)
	C._wrap_delete_PirClient_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrPirClient) Generate_data_query(arg2 ArrayofVectorUInt8, arg3 int, arg4 ArrayofVectorUInt8) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	_swig_i_1 := getSwigcptr(arg2)
	_swig_i_2 := arg3
	_swig_i_3 := getSwigcptr(arg4)
	swig_r = (int)(C._wrap_PirClient_generate_data_query_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1), C.swig_intgo(_swig_i_2), C.uintptr_t(_swig_i_3)))
	return swig_r
}

func (arg1 SwigcptrPirClient) Decrypt_data_reply(arg2 ArrayofVectorUInt8, arg3 VectorUInt8) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	_swig_i_1 := getSwigcptr(arg2)
	_swig_i_2 := getSwigcptr(arg3)
	swig_r = (int)(C._wrap_PirClient_decrypt_data_reply_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1), C.uintptr_t(_swig_i_2)))
	return swig_r
}

func (arg1 SwigcptrPirClient) Generate_data_query_with_count(arg2 ArrayofVectorUInt8, arg3 int, arg4 ArrayofVectorUInt8) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	_swig_i_1 := getSwigcptr(arg2)
	_swig_i_2 := arg3
	_swig_i_3 := getSwigcptr(arg4)
	swig_r = (int)(C._wrap_PirClient_generate_data_query_with_count_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1), C.swig_intgo(_swig_i_2), C.uintptr_t(_swig_i_3)))
	return swig_r
}

func (arg1 SwigcptrPirClient) Get_retrieved_indexes_and_ciphertext(arg2 ArrayofVectorUInt8, arg3 ArrayofVectorUInt8, arg4 int, arg5 VectorUInt8, arg6 ArrayofVectorUInt8) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	_swig_i_1 := getSwigcptr(arg2)
	_swig_i_2 := getSwigcptr(arg3)
	_swig_i_3 := arg4
	_swig_i_4 := getSwigcptr(arg5)
	_swig_i_5 := getSwigcptr(arg6)
	swig_r = (int)(C._wrap_PirClient_get_retrieved_indexes_and_ciphertext_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1), C.uintptr_t(_swig_i_2), C.swig_intgo(_swig_i_3), C.uintptr_t(_swig_i_4), C.uintptr_t(_swig_i_5)))
	return swig_r
}

func (arg1 SwigcptrPirClient) Generate_key_query_of_retrieved_indexes(arg2 ArrayofVectorUInt8, arg3 VectorUInt8) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	_swig_i_1 := getSwigcptr(arg2)
	_swig_i_2 := getSwigcptr(arg3)
	swig_r = (int)(C._wrap_PirClient_generate_key_query_of_retrieved_indexes_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1), C.uintptr_t(_swig_i_2)))
	return swig_r
}

func (arg1 SwigcptrPirClient) Decrypt_retrieved_ciphertext(arg2 ArrayofVectorUInt8, arg3 VectorUInt8, arg4 VectorUInt8) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	_swig_i_1 := getSwigcptr(arg2)
	_swig_i_2 := getSwigcptr(arg3)
	_swig_i_3 := getSwigcptr(arg4)
	swig_r = (int)(C._wrap_PirClient_decrypt_retrieved_ciphertext_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1), C.uintptr_t(_swig_i_2), C.uintptr_t(_swig_i_3)))
	return swig_r
}

func (arg1 SwigcptrPirClient) Set_key(arg2 VectorUInt8) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	_swig_i_1 := getSwigcptr(arg2)
	swig_r = (int)(C._wrap_PirClient_set_key_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1)))
	return swig_r
}

func (arg1 SwigcptrPirClient) Get_private_key() (_swig_ret VectorUInt8) {
	var swig_r VectorUInt8
	_swig_i_0 := arg1
	swig_r = (VectorUInt8)(SwigcptrVectorUInt8(C._wrap_PirClient_get_private_key_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (arg1 SwigcptrPirClient) Get_public_key() (_swig_ret VectorUInt8) {
	var swig_r VectorUInt8
	_swig_i_0 := arg1
	swig_r = (VectorUInt8)(SwigcptrVectorUInt8(C._wrap_PirClient_get_public_key_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0))))
	return swig_r
}

type PirClient interface {
	Swigcptr() uintptr
	SwigIsPirClient()
	Generate_data_query(arg2 ArrayofVectorUInt8, arg3 int, arg4 ArrayofVectorUInt8) (_swig_ret int)
	Decrypt_data_reply(arg2 ArrayofVectorUInt8, arg3 VectorUInt8) (_swig_ret int)
	Generate_data_query_with_count(arg2 ArrayofVectorUInt8, arg3 int, arg4 ArrayofVectorUInt8) (_swig_ret int)
	Get_retrieved_indexes_and_ciphertext(arg2 ArrayofVectorUInt8, arg3 ArrayofVectorUInt8, arg4 int, arg5 VectorUInt8, arg6 ArrayofVectorUInt8) (_swig_ret int)
	Generate_key_query_of_retrieved_indexes(arg2 ArrayofVectorUInt8, arg3 VectorUInt8) (_swig_ret int)
	Decrypt_retrieved_ciphertext(arg2 ArrayofVectorUInt8, arg3 VectorUInt8, arg4 VectorUInt8) (_swig_ret int)
	Set_key(arg2 VectorUInt8) (_swig_ret int)
	Get_private_key() (_swig_ret VectorUInt8)
	Get_public_key() (_swig_ret VectorUInt8)
}

type SwigcptrPirServer uintptr

func (p SwigcptrPirServer) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrPirServer) SwigIsPirServer() {
}

func NewPirServer(arg1 int, arg2 VectorUInt8, arg3 bool, arg4 *int) (_swig_ret PirServer) {
	var swig_r PirServer
	_swig_i_0 := arg1
	_swig_i_1 := getSwigcptr(arg2)
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	swig_r = (PirServer)(SwigcptrPirServer(C._wrap_new_PirServer_pir_c6eb00fb9a209540(C.swig_intgo(_swig_i_0), C.uintptr_t(_swig_i_1), C._Bool(_swig_i_2), C.swig_voidp(_swig_i_3))))
	return swig_r
}

func DeletePirServer(arg1 PirServer) {
	_swig_i_0 := getSwigcptr(arg1)
	C._wrap_delete_PirServer_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrPirServer) Generate_data_reply(arg2 ArrayofVectorUInt8, arg3 ArrayofVectorUInt8, arg4 ArrayofVectorUInt8, arg5 ArrayofVectorUInt8) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	_swig_i_1 := getSwigcptr(arg2)
	_swig_i_2 := getSwigcptr(arg3)
	_swig_i_3 := getSwigcptr(arg4)
	_swig_i_4 := getSwigcptr(arg5)
	swig_r = (int)(C._wrap_PirServer_generate_data_reply_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1), C.uintptr_t(_swig_i_2), C.uintptr_t(_swig_i_3), C.uintptr_t(_swig_i_4)))
	return swig_r
}

func (arg1 SwigcptrPirServer) Generate_checked_data_reply(arg2 ArrayofVectorUInt8, arg3 ArrayofVectorUInt8, arg4 VectorUInt8, arg5 ArrayofVectorUInt8) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	_swig_i_1 := getSwigcptr(arg2)
	_swig_i_2 := getSwigcptr(arg3)
	_swig_i_3 := getSwigcptr(arg4)
	_swig_i_4 := getSwigcptr(arg5)
	swig_r = (int)(C._wrap_PirServer_generate_checked_data_reply_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1), C.uintptr_t(_swig_i_2), C.uintptr_t(_swig_i_3), C.uintptr_t(_swig_i_4)))
	return swig_r
}

func (arg1 SwigcptrPirServer) Generate_checked_key_reply(arg2 VectorUInt8, arg3 VectorUInt8, arg4 *int) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	_swig_i_1 := getSwigcptr(arg2)
	_swig_i_2 := getSwigcptr(arg3)
	_swig_i_3 := arg4
	swig_r = (int)(C._wrap_PirServer_generate_checked_key_reply_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1), C.uintptr_t(_swig_i_2), C.swig_voidp(_swig_i_3)))
	return swig_r
}

func (arg1 SwigcptrPirServer) Set_key(arg2 VectorUInt8) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	_swig_i_1 := getSwigcptr(arg2)
	swig_r = (int)(C._wrap_PirServer_set_key_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1)))
	return swig_r
}

func (arg1 SwigcptrPirServer) Get_private_key() (_swig_ret VectorUInt8) {
	var swig_r VectorUInt8
	_swig_i_0 := arg1
	swig_r = (VectorUInt8)(SwigcptrVectorUInt8(C._wrap_PirServer_get_private_key_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (arg1 SwigcptrPirServer) Get_public_key() (_swig_ret VectorUInt8) {
	var swig_r VectorUInt8
	_swig_i_0 := arg1
	swig_r = (VectorUInt8)(SwigcptrVectorUInt8(C._wrap_PirServer_get_public_key_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0))))
	return swig_r
}

type PirServer interface {
	Swigcptr() uintptr
	SwigIsPirServer()
	Generate_data_reply(arg2 ArrayofVectorUInt8, arg3 ArrayofVectorUInt8, arg4 ArrayofVectorUInt8, arg5 ArrayofVectorUInt8) (_swig_ret int)
	Generate_checked_data_reply(arg2 ArrayofVectorUInt8, arg3 ArrayofVectorUInt8, arg4 VectorUInt8, arg5 ArrayofVectorUInt8) (_swig_ret int)
	Generate_checked_key_reply(arg2 VectorUInt8, arg3 VectorUInt8, arg4 *int) (_swig_ret int)
	Set_key(arg2 VectorUInt8) (_swig_ret int)
	Get_private_key() (_swig_ret VectorUInt8)
	Get_public_key() (_swig_ret VectorUInt8)
}

type SwigcptrVectorUInt8 uintptr

func (p SwigcptrVectorUInt8) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrVectorUInt8) SwigIsVectorUInt8() {
}

func NewVectorUInt8__SWIG_0() (_swig_ret VectorUInt8) {
	var swig_r VectorUInt8
	swig_r = (VectorUInt8)(SwigcptrVectorUInt8(C._wrap_new_VectorUInt8__SWIG_0_pir_c6eb00fb9a209540()))
	return swig_r
}

func NewVectorUInt8__SWIG_1(arg1 int64) (_swig_ret VectorUInt8) {
	var swig_r VectorUInt8
	_swig_i_0 := arg1
	swig_r = (VectorUInt8)(SwigcptrVectorUInt8(C._wrap_new_VectorUInt8__SWIG_1_pir_c6eb00fb9a209540(C.swig_type_1(_swig_i_0))))
	return swig_r
}

func NewVectorUInt8__SWIG_2(arg1 VectorUInt8) (_swig_ret VectorUInt8) {
	var swig_r VectorUInt8
	_swig_i_0 := getSwigcptr(arg1)
	swig_r = (VectorUInt8)(SwigcptrVectorUInt8(C._wrap_new_VectorUInt8__SWIG_2_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func NewVectorUInt8(a ...interface{}) VectorUInt8 {
	argc := len(a)
	if argc == 0 {
		return NewVectorUInt8__SWIG_0()
	}
	if argc == 1 {
		if _, ok := a[0].(int64); !ok {
			goto check_2
		}
		return NewVectorUInt8__SWIG_1(a[0].(int64))
	}
check_2:
	if argc == 1 {
		return NewVectorUInt8__SWIG_2(a[0].(VectorUInt8))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrVectorUInt8) Size() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_VectorUInt8_size_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrVectorUInt8) Capacity() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_VectorUInt8_capacity_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrVectorUInt8) Reserve(arg2 int64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_VectorUInt8_reserve_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0), C.swig_type_4(_swig_i_1))
}

func (arg1 SwigcptrVectorUInt8) IsEmpty() (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_VectorUInt8_isEmpty_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrVectorUInt8) Clear() {
	_swig_i_0 := arg1
	C._wrap_VectorUInt8_clear_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrVectorUInt8) Add(arg2 byte) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_VectorUInt8_add_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0), C.char(_swig_i_1))
}

func (arg1 SwigcptrVectorUInt8) Get(arg2 int) (_swig_ret byte) {
	var swig_r byte
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (byte)(C._wrap_VectorUInt8_get_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1)))
	return swig_r
}

func (arg1 SwigcptrVectorUInt8) Set(arg2 int, arg3 byte) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	C._wrap_VectorUInt8_set_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.char(_swig_i_2))
}

func DeleteVectorUInt8(arg1 VectorUInt8) {
	_swig_i_0 := getSwigcptr(arg1)
	C._wrap_delete_VectorUInt8_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0))
}

type VectorUInt8 interface {
	Swigcptr() uintptr
	SwigIsVectorUInt8()
	Size() (_swig_ret int64)
	Capacity() (_swig_ret int64)
	Reserve(arg2 int64)
	IsEmpty() (_swig_ret bool)
	Clear()
	Add(arg2 byte)
	Get(arg2 int) (_swig_ret byte)
	Set(arg2 int, arg3 byte)
}

type SwigcptrArrayofVectorUInt8 uintptr

func (p SwigcptrArrayofVectorUInt8) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrArrayofVectorUInt8) SwigIsArrayofVectorUInt8() {
}

func NewArrayofVectorUInt8__SWIG_0() (_swig_ret ArrayofVectorUInt8) {
	var swig_r ArrayofVectorUInt8
	swig_r = (ArrayofVectorUInt8)(SwigcptrArrayofVectorUInt8(C._wrap_new_ArrayofVectorUInt8__SWIG_0_pir_c6eb00fb9a209540()))
	return swig_r
}

func NewArrayofVectorUInt8__SWIG_1(arg1 int64) (_swig_ret ArrayofVectorUInt8) {
	var swig_r ArrayofVectorUInt8
	_swig_i_0 := arg1
	swig_r = (ArrayofVectorUInt8)(SwigcptrArrayofVectorUInt8(C._wrap_new_ArrayofVectorUInt8__SWIG_1_pir_c6eb00fb9a209540(C.swig_type_5(_swig_i_0))))
	return swig_r
}

func NewArrayofVectorUInt8__SWIG_2(arg1 ArrayofVectorUInt8) (_swig_ret ArrayofVectorUInt8) {
	var swig_r ArrayofVectorUInt8
	_swig_i_0 := getSwigcptr(arg1)
	swig_r = (ArrayofVectorUInt8)(SwigcptrArrayofVectorUInt8(C._wrap_new_ArrayofVectorUInt8__SWIG_2_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func NewArrayofVectorUInt8(a ...interface{}) ArrayofVectorUInt8 {
	argc := len(a)
	if argc == 0 {
		return NewArrayofVectorUInt8__SWIG_0()
	}
	if argc == 1 {
		if _, ok := a[0].(int64); !ok {
			goto check_2
		}
		return NewArrayofVectorUInt8__SWIG_1(a[0].(int64))
	}
check_2:
	if argc == 1 {
		return NewArrayofVectorUInt8__SWIG_2(a[0].(ArrayofVectorUInt8))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrArrayofVectorUInt8) Size() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_ArrayofVectorUInt8_size_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrArrayofVectorUInt8) Capacity() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_ArrayofVectorUInt8_capacity_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrArrayofVectorUInt8) Reserve(arg2 int64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_ArrayofVectorUInt8_reserve_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0), C.swig_type_8(_swig_i_1))
}

func (arg1 SwigcptrArrayofVectorUInt8) IsEmpty() (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_ArrayofVectorUInt8_isEmpty_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrArrayofVectorUInt8) Clear() {
	_swig_i_0 := arg1
	C._wrap_ArrayofVectorUInt8_clear_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrArrayofVectorUInt8) Add(arg2 VectorUInt8) {
	_swig_i_0 := arg1
	_swig_i_1 := getSwigcptr(arg2)
	C._wrap_ArrayofVectorUInt8_add_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func (arg1 SwigcptrArrayofVectorUInt8) Get(arg2 int) (_swig_ret VectorUInt8) {
	var swig_r VectorUInt8
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (VectorUInt8)(SwigcptrVectorUInt8(C._wrap_ArrayofVectorUInt8_get_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))))
	return swig_r
}

func (arg1 SwigcptrArrayofVectorUInt8) Set(arg2 int, arg3 VectorUInt8) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := getSwigcptr(arg3)
	C._wrap_ArrayofVectorUInt8_set_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.uintptr_t(_swig_i_2))
}

func DeleteArrayofVectorUInt8(arg1 ArrayofVectorUInt8) {
	_swig_i_0 := getSwigcptr(arg1)
	C._wrap_delete_ArrayofVectorUInt8_pir_c6eb00fb9a209540(C.uintptr_t(_swig_i_0))
}

type ArrayofVectorUInt8 interface {
	Swigcptr() uintptr
	SwigIsArrayofVectorUInt8()
	Size() (_swig_ret int64)
	Capacity() (_swig_ret int64)
	Reserve(arg2 int64)
	IsEmpty() (_swig_ret bool)
	Clear()
	Add(arg2 VectorUInt8)
	Get(arg2 int) (_swig_ret VectorUInt8)
	Set(arg2 int, arg3 VectorUInt8)
}


