// Code generated by Wire protocol buffer compiler, do not edit.
// Source: ChatServerEvent in chat.proto
package gg.strims.ppspp.proto

import com.squareup.wire.FieldEncoding
import com.squareup.wire.Message
import com.squareup.wire.ProtoAdapter
import com.squareup.wire.ProtoReader
import com.squareup.wire.ProtoWriter
import com.squareup.wire.Syntax
import com.squareup.wire.Syntax.PROTO_3
import com.squareup.wire.WireField
import com.squareup.wire.internal.countNonNull
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

class ChatServerEvent(
  @field:WireField(
    tag = 1,
    adapter = "gg.strims.ppspp.proto.ChatServerEvent${'$'}Open#ADAPTER"
  )
  val open: Open? = null,
  @field:WireField(
    tag = 2,
    adapter = "gg.strims.ppspp.proto.ChatServerEvent${'$'}Close#ADAPTER"
  )
  val close: Close? = null,
  unknownFields: ByteString = ByteString.EMPTY
) : Message<ChatServerEvent, Nothing>(ADAPTER, unknownFields) {
  init {
    require(countNonNull(open, close) <= 1) {
      "At most one of open, close may be non-null"
    }
  }

  @Deprecated(
    message = "Shouldn't be used in Kotlin",
    level = DeprecationLevel.HIDDEN
  )
  override fun newBuilder(): Nothing = throw AssertionError()

  override fun equals(other: Any?): Boolean {
    if (other === this) return true
    if (other !is ChatServerEvent) return false
    if (unknownFields != other.unknownFields) return false
    if (open != other.open) return false
    if (close != other.close) return false
    return true
  }

  override fun hashCode(): Int {
    var result = super.hashCode
    if (result == 0) {
      result = unknownFields.hashCode()
      result = result * 37 + open.hashCode()
      result = result * 37 + close.hashCode()
      super.hashCode = result
    }
    return result
  }

  override fun toString(): String {
    val result = mutableListOf<String>()
    if (open != null) result += """open=$open"""
    if (close != null) result += """close=$close"""
    return result.joinToString(prefix = "ChatServerEvent{", separator = ", ", postfix = "}")
  }

  fun copy(
    open: Open? = this.open,
    close: Close? = this.close,
    unknownFields: ByteString = this.unknownFields
  ): ChatServerEvent = ChatServerEvent(open, close, unknownFields)

  companion object {
    @JvmField
    val ADAPTER: ProtoAdapter<ChatServerEvent> = object : ProtoAdapter<ChatServerEvent>(
      FieldEncoding.LENGTH_DELIMITED, 
      ChatServerEvent::class, 
      "type.googleapis.com/ChatServerEvent", 
      PROTO_3, 
      null
    ) {
      override fun encodedSize(value: ChatServerEvent): Int {
        var size = value.unknownFields.size
        size += Open.ADAPTER.encodedSizeWithTag(1, value.open)
        size += Close.ADAPTER.encodedSizeWithTag(2, value.close)
        return size
      }

      override fun encode(writer: ProtoWriter, value: ChatServerEvent) {
        Open.ADAPTER.encodeWithTag(writer, 1, value.open)
        Close.ADAPTER.encodeWithTag(writer, 2, value.close)
        writer.writeBytes(value.unknownFields)
      }

      override fun decode(reader: ProtoReader): ChatServerEvent {
        var open: Open? = null
        var close: Close? = null
        val unknownFields = reader.forEachTag { tag ->
          when (tag) {
            1 -> open = Open.ADAPTER.decode(reader)
            2 -> close = Close.ADAPTER.decode(reader)
            else -> reader.readUnknownField(tag)
          }
        }
        return ChatServerEvent(
          open = open,
          close = close,
          unknownFields = unknownFields
        )
      }

      override fun redact(value: ChatServerEvent): ChatServerEvent = value.copy(
        open = value.open?.let(Open.ADAPTER::redact),
        close = value.close?.let(Close.ADAPTER::redact),
        unknownFields = ByteString.EMPTY
      )
    }

    private const val serialVersionUID: Long = 0L
  }

  class Open(
    @field:WireField(
      tag = 1,
      adapter = "com.squareup.wire.ProtoAdapter#UINT64",
      label = WireField.Label.OMIT_IDENTITY,
      jsonName = "serverId"
    )
    val server_id: Long = 0L,
    unknownFields: ByteString = ByteString.EMPTY
  ) : Message<Open, Nothing>(ADAPTER, unknownFields) {
    @Deprecated(
      message = "Shouldn't be used in Kotlin",
      level = DeprecationLevel.HIDDEN
    )
    override fun newBuilder(): Nothing = throw AssertionError()

    override fun equals(other: Any?): Boolean {
      if (other === this) return true
      if (other !is Open) return false
      if (unknownFields != other.unknownFields) return false
      if (server_id != other.server_id) return false
      return true
    }

    override fun hashCode(): Int {
      var result = super.hashCode
      if (result == 0) {
        result = unknownFields.hashCode()
        result = result * 37 + server_id.hashCode()
        super.hashCode = result
      }
      return result
    }

    override fun toString(): String {
      val result = mutableListOf<String>()
      result += """server_id=$server_id"""
      return result.joinToString(prefix = "Open{", separator = ", ", postfix = "}")
    }

    fun copy(server_id: Long = this.server_id, unknownFields: ByteString = this.unknownFields): Open
        = Open(server_id, unknownFields)

    companion object {
      @JvmField
      val ADAPTER: ProtoAdapter<Open> = object : ProtoAdapter<Open>(
        FieldEncoding.LENGTH_DELIMITED, 
        Open::class, 
        "type.googleapis.com/ChatServerEvent.Open", 
        PROTO_3, 
        null
      ) {
        override fun encodedSize(value: Open): Int {
          var size = value.unknownFields.size
          if (value.server_id != 0L) size += ProtoAdapter.UINT64.encodedSizeWithTag(1,
              value.server_id)
          return size
        }

        override fun encode(writer: ProtoWriter, value: Open) {
          if (value.server_id != 0L) ProtoAdapter.UINT64.encodeWithTag(writer, 1, value.server_id)
          writer.writeBytes(value.unknownFields)
        }

        override fun decode(reader: ProtoReader): Open {
          var server_id: Long = 0L
          val unknownFields = reader.forEachTag { tag ->
            when (tag) {
              1 -> server_id = ProtoAdapter.UINT64.decode(reader)
              else -> reader.readUnknownField(tag)
            }
          }
          return Open(
            server_id = server_id,
            unknownFields = unknownFields
          )
        }

        override fun redact(value: Open): Open = value.copy(
          unknownFields = ByteString.EMPTY
        )
      }

      private const val serialVersionUID: Long = 0L
    }
  }

  class Close(
    unknownFields: ByteString = ByteString.EMPTY
  ) : Message<Close, Nothing>(ADAPTER, unknownFields) {
    @Deprecated(
      message = "Shouldn't be used in Kotlin",
      level = DeprecationLevel.HIDDEN
    )
    override fun newBuilder(): Nothing = throw AssertionError()

    override fun equals(other: Any?): Boolean {
      if (other === this) return true
      if (other !is Close) return false
      if (unknownFields != other.unknownFields) return false
      return true
    }

    override fun hashCode(): Int = unknownFields.hashCode()

    override fun toString(): String = "Close{}"

    fun copy(unknownFields: ByteString = this.unknownFields): Close = Close(unknownFields)

    companion object {
      @JvmField
      val ADAPTER: ProtoAdapter<Close> = object : ProtoAdapter<Close>(
        FieldEncoding.LENGTH_DELIMITED, 
        Close::class, 
        "type.googleapis.com/ChatServerEvent.Close", 
        PROTO_3, 
        null
      ) {
        override fun encodedSize(value: Close): Int {
          var size = value.unknownFields.size
          return size
        }

        override fun encode(writer: ProtoWriter, value: Close) {
          writer.writeBytes(value.unknownFields)
        }

        override fun decode(reader: ProtoReader): Close {
          val unknownFields = reader.forEachTag(reader::readUnknownField)
          return Close(
            unknownFields = unknownFields
          )
        }

        override fun redact(value: Close): Close = value.copy(
          unknownFields = ByteString.EMPTY
        )
      }

      private const val serialVersionUID: Long = 0L
    }
  }
}
