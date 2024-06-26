/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (https://www.swig.org).
 * Version 4.1.1
 *
 * Do not make changes to this file unless you know what you are doing - modify
 * the SWIG interface file instead.
 * ----------------------------------------------------------------------------- */

package com.pir;

public class Faker {
  private transient long swigCPtr;
  protected transient boolean swigCMemOwn;

  protected Faker(long cPtr, boolean cMemoryOwn) {
    swigCMemOwn = cMemoryOwn;
    swigCPtr = cPtr;
  }

  protected static long getCPtr(Faker obj) {
    return (obj == null) ? 0 : obj.swigCPtr;
  }

  protected static long swigRelease(Faker obj) {
    long ptr = 0;
    if (obj != null) {
      if (!obj.swigCMemOwn)
        throw new RuntimeException("Cannot release ownership as memory is not owned");
      ptr = obj.swigCPtr;
      obj.swigCMemOwn = false;
      obj.delete();
    }
    return ptr;
  }

  @SuppressWarnings("deprecation")
  protected void finalize() {
    delete();
  }

  public synchronized void delete() {
    if (swigCPtr != 0) {
      if (swigCMemOwn) {
        swigCMemOwn = false;
        pirJNI.delete_Faker(swigCPtr);
      }
      swigCPtr = 0;
    }
  }

  public Faker() {
    this(pirJNI.new_Faker(), true);
  }

  public ArrayofVectorUInt8 GenerateCNIDCardNumber(int number) {
    return new ArrayofVectorUInt8(pirJNI.Faker_GenerateCNIDCardNumber__SWIG_0(swigCPtr, this, number), true);
  }

  public ArrayofVectorUInt8 GenerateCNIDCardNumber() {
    return new ArrayofVectorUInt8(pirJNI.Faker_GenerateCNIDCardNumber__SWIG_1(swigCPtr, this), true);
  }

  public ArrayofVectorUInt8 GenerateTWCompatriotNumber(int number) {
    return new ArrayofVectorUInt8(pirJNI.Faker_GenerateTWCompatriotNumber__SWIG_0(swigCPtr, this, number), true);
  }

  public ArrayofVectorUInt8 GenerateTWCompatriotNumber() {
    return new ArrayofVectorUInt8(pirJNI.Faker_GenerateTWCompatriotNumber__SWIG_1(swigCPtr, this), true);
  }

  public ArrayofVectorUInt8 GenerateTWIDCardNumber(int number) {
    return new ArrayofVectorUInt8(pirJNI.Faker_GenerateTWIDCardNumber__SWIG_0(swigCPtr, this, number), true);
  }

  public ArrayofVectorUInt8 GenerateTWIDCardNumber() {
    return new ArrayofVectorUInt8(pirJNI.Faker_GenerateTWIDCardNumber__SWIG_1(swigCPtr, this), true);
  }

  public ArrayofVectorUInt8 GenerateTW2CNIDCardNumber(int number) {
    return new ArrayofVectorUInt8(pirJNI.Faker_GenerateTW2CNIDCardNumber__SWIG_0(swigCPtr, this, number), true);
  }

  public ArrayofVectorUInt8 GenerateTW2CNIDCardNumber() {
    return new ArrayofVectorUInt8(pirJNI.Faker_GenerateTW2CNIDCardNumber__SWIG_1(swigCPtr, this), true);
  }

  public ArrayofVectorUInt8 GeneratePhoneNumber(int number) {
    return new ArrayofVectorUInt8(pirJNI.Faker_GeneratePhoneNumber__SWIG_0(swigCPtr, this, number), true);
  }

  public ArrayofVectorUInt8 GeneratePhoneNumber() {
    return new ArrayofVectorUInt8(pirJNI.Faker_GeneratePhoneNumber__SWIG_1(swigCPtr, this), true);
  }

  public ArrayofVectorUInt8 GenerateMACAddress(int number) {
    return new ArrayofVectorUInt8(pirJNI.Faker_GenerateMACAddress__SWIG_0(swigCPtr, this, number), true);
  }

  public ArrayofVectorUInt8 GenerateMACAddress() {
    return new ArrayofVectorUInt8(pirJNI.Faker_GenerateMACAddress__SWIG_1(swigCPtr, this), true);
  }

  public ArrayofVectorUInt8 GenerateIMEI(int number) {
    return new ArrayofVectorUInt8(pirJNI.Faker_GenerateIMEI__SWIG_0(swigCPtr, this, number), true);
  }

  public ArrayofVectorUInt8 GenerateIMEI() {
    return new ArrayofVectorUInt8(pirJNI.Faker_GenerateIMEI__SWIG_1(swigCPtr, this), true);
  }

  public ArrayofVectorUInt8 GenerateIMSI(int number) {
    return new ArrayofVectorUInt8(pirJNI.Faker_GenerateIMSI__SWIG_0(swigCPtr, this, number), true);
  }

  public ArrayofVectorUInt8 GenerateIMSI() {
    return new ArrayofVectorUInt8(pirJNI.Faker_GenerateIMSI__SWIG_1(swigCPtr, this), true);
  }

  public ArrayofVectorUInt8 GenerateEmail(int number) {
    return new ArrayofVectorUInt8(pirJNI.Faker_GenerateEmail__SWIG_0(swigCPtr, this, number), true);
  }

  public ArrayofVectorUInt8 GenerateEmail() {
    return new ArrayofVectorUInt8(pirJNI.Faker_GenerateEmail__SWIG_1(swigCPtr, this), true);
  }

  public ArrayofVectorUInt8 GeneratePassportNumber(int number) {
    return new ArrayofVectorUInt8(pirJNI.Faker_GeneratePassportNumber__SWIG_0(swigCPtr, this, number), true);
  }

  public ArrayofVectorUInt8 GeneratePassportNumber() {
    return new ArrayofVectorUInt8(pirJNI.Faker_GeneratePassportNumber__SWIG_1(swigCPtr, this), true);
  }

}
