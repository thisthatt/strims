// Code generated by Wire protocol buffer compiler, do not edit.
// Source: NickServToken in nickserv.proto
package gg.strims.ppspp.proto

import com.squareup.wire.FieldEncoding
import com.squareup.wire.Instant
import com.squareup.wire.Message
import com.squareup.wire.ProtoAdapter
import com.squareup.wire.ProtoReader
import com.squareup.wire.ProtoWriter
import com.squareup.wire.Syntax.PROTO_3
import com.squareup.wire.WireField
import com.squareup.wire.internal.immutableCopyOf
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
import kotlin.collections.List
import kotlin.hashCode
import kotlin.jvm.JvmField
import okio.ByteString

class NickServToken(
  @field:WireField(
    tag = 1,
    adapter = "com.squareup.wire.ProtoAdapter#BYTES",
    label = WireField.Label.OMIT_IDENTITY
  )
  val key: ByteString = ByteString.EMPTY,
  @field:WireField(
    tag = 2,
    adapter = "com.squareup.wire.ProtoAdapter#STRING",
    label = WireField.Label.OMIT_IDENTITY
  )
  val nick: String = "",
  @field:WireField(
    tag = 3,
    adapter = "com.squareup.wire.ProtoAdapter#INSTANT",
    label = WireField.Label.OMIT_IDENTITY,
    jsonName = "validUntil"
  )
  val valid_until: Instant? = null,
  @field:WireField(
    tag = 4,
    adapter = "com.squareup.wire.ProtoAdapter#BYTES",
    label = WireField.Label.OMIT_IDENTITY
  )
  val signature: ByteString = ByteString.EMPTY,
  roles: List<String> = emptyList(),
  unknownFields: ByteString = ByteString.EMPTY
) : Message<NickServToken, Nothing>(ADAPTER, unknownFields) {
  @field:WireField(
    tag = 5,
    adapter = "com.squareup.wire.ProtoAdapter#STRING",
    label = WireField.Label.REPEATED
  )
  val roles: List<String> = immutableCopyOf("roles", roles)

  @Deprecated(
    message = "Shouldn't be used in Kotlin",
    level = DeprecationLevel.HIDDEN
  )
  override fun newBuilder(): Nothing = throw AssertionError()

  override fun equals(other: Any?): Boolean {
    if (other === this) return true
    if (other !is NickServToken) return false
    if (unknownFields != other.unknownFields) return false
    if (key != other.key) return false
    if (nick != other.nick) return false
    if (valid_until != other.valid_until) return false
    if (signature != other.signature) return false
    if (roles != other.roles) return false
    return true
  }

  override fun hashCode(): Int {
    var result = super.hashCode
    if (result == 0) {
      result = unknownFields.hashCode()
      result = result * 37 + key.hashCode()
      result = result * 37 + nick.hashCode()
      result = result * 37 + valid_until.hashCode()
      result = result * 37 + signature.hashCode()
      result = result * 37 + roles.hashCode()
      super.hashCode = result
    }
    return result
  }

  override fun toString(): String {
    val result = mutableListOf<String>()
    result += """key=$key"""
    result += """nick=${sanitize(nick)}"""
    if (valid_until != null) result += """valid_until=$valid_until"""
    result += """signature=$signature"""
    if (roles.isNotEmpty()) result += """roles=${sanitize(roles)}"""
    return result.joinToString(prefix = "NickServToken{", separator = ", ", postfix = "}")
  }

  fun copy(
    key: ByteString = this.key,
    nick: String = this.nick,
    valid_until: Instant? = this.valid_until,
    signature: ByteString = this.signature,
    roles: List<String> = this.roles,
    unknownFields: ByteString = this.unknownFields
  ): NickServToken = NickServToken(key, nick, valid_until, signature, roles, unknownFields)

  companion object {
    @JvmField
    val ADAPTER: ProtoAdapter<NickServToken> = object : ProtoAdapter<NickServToken>(
      FieldEncoding.LENGTH_DELIMITED, 
      NickServToken::class, 
      "type.googleapis.com/NickServToken", 
      PROTO_3, 
      null
    ) {
      override fun encodedSize(value: NickServToken): Int {
        var size = value.unknownFields.size
        if (value.key != ByteString.EMPTY) size += ProtoAdapter.BYTES.encodedSizeWithTag(1,
            value.key)
        if (value.nick != "") size += ProtoAdapter.STRING.encodedSizeWithTag(2, value.nick)
        if (value.valid_until != null) size += ProtoAdapter.INSTANT.encodedSizeWithTag(3,
            value.valid_until)
        if (value.signature != ByteString.EMPTY) size += ProtoAdapter.BYTES.encodedSizeWithTag(4,
            value.signature)
        size += ProtoAdapter.STRING.asRepeated().encodedSizeWithTag(5, value.roles)
        return size
      }

      override fun encode(writer: ProtoWriter, value: NickServToken) {
        if (value.key != ByteString.EMPTY) ProtoAdapter.BYTES.encodeWithTag(writer, 1, value.key)
        if (value.nick != "") ProtoAdapter.STRING.encodeWithTag(writer, 2, value.nick)
        if (value.valid_until != null) ProtoAdapter.INSTANT.encodeWithTag(writer, 3,
            value.valid_until)
        if (value.signature != ByteString.EMPTY) ProtoAdapter.BYTES.encodeWithTag(writer, 4,
            value.signature)
        ProtoAdapter.STRING.asRepeated().encodeWithTag(writer, 5, value.roles)
        writer.writeBytes(value.unknownFields)
      }

      override fun decode(reader: ProtoReader): NickServToken {
        var key: ByteString = ByteString.EMPTY
        var nick: String = ""
        var valid_until: Instant? = null
        var signature: ByteString = ByteString.EMPTY
        val roles = mutableListOf<String>()
        val unknownFields = reader.forEachTag { tag ->
          when (tag) {
            1 -> key = ProtoAdapter.BYTES.decode(reader)
            2 -> nick = ProtoAdapter.STRING.decode(reader)
            3 -> valid_until = ProtoAdapter.INSTANT.decode(reader)
            4 -> signature = ProtoAdapter.BYTES.decode(reader)
            5 -> roles.add(ProtoAdapter.STRING.decode(reader))
            else -> reader.readUnknownField(tag)
          }
        }
        return NickServToken(
          key = key,
          nick = nick,
          valid_until = valid_until,
          signature = signature,
          roles = roles,
          unknownFields = unknownFields
        )
      }

      override fun redact(value: NickServToken): NickServToken = value.copy(
        valid_until = value.valid_until?.let(ProtoAdapter.INSTANT::redact),
        unknownFields = ByteString.EMPTY
      )
    }

    private const val serialVersionUID: Long = 0L
  }
}
