/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (https://www.swig.org).
 * Version 4.1.1
 *
 * Do not make changes to this file unless you know what you are doing - modify
 * the SWIG interface file instead.
 * ----------------------------------------------------------------------------- */

package com.pir;

public class PirServer {
  private transient long swigCPtr;
  protected transient boolean swigCMemOwn;

  protected PirServer(long cPtr, boolean cMemoryOwn) {
    swigCMemOwn = cMemoryOwn;
    swigCPtr = cPtr;
  }

  protected static long getCPtr(PirServer obj) {
    return (obj == null) ? 0 : obj.swigCPtr;
  }

  protected static long swigRelease(PirServer obj) {
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
        pirJNI.delete_PirServer(swigCPtr);
      }
      swigCPtr = 0;
    }
  }

  public PirServer(int anonymity_number, VectorUInt8 key, boolean is_checked_count, SWIGTYPE_p_int error_code) {
    this(pirJNI.new_PirServer(anonymity_number, VectorUInt8.getCPtr(key), key, is_checked_count, SWIGTYPE_p_int.getCPtr(error_code)), true);
  }

  public int generate_data_reply(ArrayofVectorUInt8 data, ArrayofVectorUInt8 ID, ArrayofVectorUInt8 data_query, ArrayofVectorUInt8 data_reply) {
    return pirJNI.PirServer_generate_data_reply(swigCPtr, this, ArrayofVectorUInt8.getCPtr(data), data, ArrayofVectorUInt8.getCPtr(ID), ID, ArrayofVectorUInt8.getCPtr(data_query), data_query, ArrayofVectorUInt8.getCPtr(data_reply), data_reply);
  }

  public int generate_checked_data_reply(ArrayofVectorUInt8 checked_bits, ArrayofVectorUInt8 indexes, VectorUInt8 data_query, ArrayofVectorUInt8 data_reply) {
    return pirJNI.PirServer_generate_checked_data_reply(swigCPtr, this, ArrayofVectorUInt8.getCPtr(checked_bits), checked_bits, ArrayofVectorUInt8.getCPtr(indexes), indexes, VectorUInt8.getCPtr(data_query), data_query, ArrayofVectorUInt8.getCPtr(data_reply), data_reply);
  }

  public int generate_checked_key_reply(VectorUInt8 retrieved_key_query, VectorUInt8 retrieved_key_reply, SWIGTYPE_p_int checked_number) {
    return pirJNI.PirServer_generate_checked_key_reply(swigCPtr, this, VectorUInt8.getCPtr(retrieved_key_query), retrieved_key_query, VectorUInt8.getCPtr(retrieved_key_reply), retrieved_key_reply, SWIGTYPE_p_int.getCPtr(checked_number));
  }

  public int set_key(VectorUInt8 private_key) {
    return pirJNI.PirServer_set_key(swigCPtr, this, VectorUInt8.getCPtr(private_key), private_key);
  }

  public VectorUInt8 get_private_key() {
    return new VectorUInt8(pirJNI.PirServer_get_private_key(swigCPtr, this), true);
  }

  public VectorUInt8 get_public_key() {
    return new VectorUInt8(pirJNI.PirServer_get_public_key(swigCPtr, this), true);
  }

}
