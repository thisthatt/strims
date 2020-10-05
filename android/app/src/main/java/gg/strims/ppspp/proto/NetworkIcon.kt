// Code generated by Wire protocol buffer compiler, do not edit.
// Source: NetworkIcon in profile.proto
package gg.strims.ppspp.proto

import com.squareup.wire.FieldEncoding
import com.squareup.wire.Message
import com.squareup.wire.ProtoAdapter
import com.squareup.wire.ProtoReader
import com.squareup.wire.ProtoWriter
import com.squareup.wire.Syntax.PROTO_3
import com.squareup.wire.WireField
import com.squareup.wire.internal.sanitize
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

class NetworkIcon(
  @field:WireField(
    tag = 1,
    adapter = "com.squareup.wire.ProtoAdapter#BYTES",
    label = WireField.Label.OMIT_IDENTITY
  )
  val data: ByteString = ByteString.EMPTY,
  @field:WireField(
    tag = 2,
    adapter = "com.squareup.wire.ProtoAdapter#STRING",
    label = WireField.Label.OMIT_IDENTITY
  )
  val type: String = "",
  unknownFields: ByteString = ByteString.EMPTY
) : Message<NetworkIcon, Nothing>(ADAPTER, unknownFields) {
  @Deprecated(
    message = "Shouldn't be used in Kotlin",
    level = DeprecationLevel.HIDDEN
  )
  override fun newBuilder(): Nothing = throw AssertionError()

  override fun equals(other: Any?): Boolean {
    if (other === this) return true
    if (other !is NetworkIcon) return false
    if (unknownFields != other.unknownFields) return false
    if (data != other.data) return false
    if (type != other.type) return false
    return true
  }

  override fun hashCode(): Int {
    var result = super.hashCode
    if (result == 0) {
      result = unknownFields.hashCode()
      result = result * 37 + data.hashCode()
      result = result * 37 + type.hashCode()
      super.hashCode = result
    }
    return result
  }

  override fun toString(): String {
    val result = mutableListOf<String>()
    result += """data=$data"""
    result += """type=${sanitize(type)}"""
    return result.joinToString(prefix = "NetworkIcon{", separator = ", ", postfix = "}")
  }

  fun copy(
    data: ByteString = this.data,
    type: String = this.type,
    unknownFields: ByteString = this.unknownFields
  ): NetworkIcon = NetworkIcon(data, type, unknownFields)

  companion object {
    @JvmField
    val ADAPTER: ProtoAdapter<NetworkIcon> = object : ProtoAdapter<NetworkIcon>(
      FieldEncoding.LENGTH_DELIMITED, 
      NetworkIcon::class, 
      "type.googleapis.com/NetworkIcon", 
      PROTO_3, 
      null
    ) {
      override fun encodedSize(value: NetworkIcon): Int {
        var size = value.unknownFields.size
        if (value.data != ByteString.EMPTY) size += ProtoAdapter.BYTES.encodedSizeWithTag(1,
            value.data)
        if (value.type != "") size += ProtoAdapter.STRING.encodedSizeWithTag(2, value.type)
        return size
      }

      override fun encode(writer: ProtoWriter, value: NetworkIcon) {
        if (value.data != ByteString.EMPTY) ProtoAdapter.BYTES.encodeWithTag(writer, 1, value.data)
        if (value.type != "") ProtoAdapter.STRING.encodeWithTag(writer, 2, value.type)
        writer.writeBytes(value.unknownFields)
      }

      override fun decode(reader: ProtoReader): NetworkIcon {
        var data: ByteString = ByteString.EMPTY
        var type: String = ""
        val unknownFields = reader.forEachTag { tag ->
          when (tag) {
            1 -> data = ProtoAdapter.BYTES.decode(reader)
            2 -> type = ProtoAdapter.STRING.decode(reader)
            else -> reader.readUnknownField(tag)
          }
        }
        return NetworkIcon(
          data = data,
          type = type,
          unknownFields = unknownFields
        )
      }

      override fun redact(value: NetworkIcon): NetworkIcon = value.copy(
        unknownFields = ByteString.EMPTY
      )
    }

    private const val serialVersionUID: Long = 0L
  }
}
