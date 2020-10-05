// Code generated by Wire protocol buffer compiler, do not edit.
// Source: DeleteChatServerRequest in chat.proto
package gg.strims.ppspp.proto

import com.squareup.wire.FieldEncoding
import com.squareup.wire.Message
import com.squareup.wire.ProtoAdapter
import com.squareup.wire.ProtoReader
import com.squareup.wire.ProtoWriter
import com.squareup.wire.Syntax.PROTO_3
import com.squareup.wire.WireField
import kotlin.Any
import kotlin.AssertionError
import kotlin.Boolean
import kotlin.Deprecated
import kotlin.DeprecationLevel
import kotlin.Int
import kotlin.Long
import kotlin.Nothing
import kotlin.String
import kotlin.hashCode
import kotlin.jvm.JvmField
import okio.ByteString

class DeleteChatServerRequest(
  @field:WireField(
    tag = 1,
    adapter = "com.squareup.wire.ProtoAdapter#UINT64",
    label = WireField.Label.OMIT_IDENTITY
  )
  val id: Long = 0L,
  unknownFields: ByteString = ByteString.EMPTY
) : Message<DeleteChatServerRequest, Nothing>(ADAPTER, unknownFields) {
  @Deprecated(
    message = "Shouldn't be used in Kotlin",
    level = DeprecationLevel.HIDDEN
  )
  override fun newBuilder(): Nothing = throw AssertionError()

  override fun equals(other: Any?): Boolean {
    if (other === this) return true
    if (other !is DeleteChatServerRequest) return false
    if (unknownFields != other.unknownFields) return false
    if (id != other.id) return false
    return true
  }

  override fun hashCode(): Int {
    var result = super.hashCode
    if (result == 0) {
      result = unknownFields.hashCode()
      result = result * 37 + id.hashCode()
      super.hashCode = result
    }
    return result
  }

  override fun toString(): String {
    val result = mutableListOf<String>()
    result += """id=$id"""
    return result.joinToString(prefix = "DeleteChatServerRequest{", separator = ", ", postfix = "}")
  }

  fun copy(id: Long = this.id, unknownFields: ByteString = this.unknownFields):
      DeleteChatServerRequest = DeleteChatServerRequest(id, unknownFields)

  companion object {
    @JvmField
    val ADAPTER: ProtoAdapter<DeleteChatServerRequest> = object :
        ProtoAdapter<DeleteChatServerRequest>(
      FieldEncoding.LENGTH_DELIMITED, 
      DeleteChatServerRequest::class, 
      "type.googleapis.com/DeleteChatServerRequest", 
      PROTO_3, 
      null
    ) {
      override fun encodedSize(value: DeleteChatServerRequest): Int {
        var size = value.unknownFields.size
        if (value.id != 0L) size += ProtoAdapter.UINT64.encodedSizeWithTag(1, value.id)
        return size
      }

      override fun encode(writer: ProtoWriter, value: DeleteChatServerRequest) {
        if (value.id != 0L) ProtoAdapter.UINT64.encodeWithTag(writer, 1, value.id)
        writer.writeBytes(value.unknownFields)
      }

      override fun decode(reader: ProtoReader): DeleteChatServerRequest {
        var id: Long = 0L
        val unknownFields = reader.forEachTag { tag ->
          when (tag) {
            1 -> id = ProtoAdapter.UINT64.decode(reader)
            else -> reader.readUnknownField(tag)
          }
        }
        return DeleteChatServerRequest(
          id = id,
          unknownFields = unknownFields
        )
      }

      override fun redact(value: DeleteChatServerRequest): DeleteChatServerRequest = value.copy(
        unknownFields = ByteString.EMPTY
      )
    }

    private const val serialVersionUID: Long = 0L
  }
}
