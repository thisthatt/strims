// Code generated by Wire protocol buffer compiler, do not edit.
// Source: DirectoryServerEvent in directory.proto
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

class DirectoryServerEvent(
  @field:WireField(
    tag = 1,
    adapter = "gg.strims.ppspp.proto.DirectoryServerEvent${'$'}Publish#ADAPTER"
  )
  val publish: Publish? = null,
  @field:WireField(
    tag = 2,
    adapter = "gg.strims.ppspp.proto.DirectoryServerEvent${'$'}Unpublish#ADAPTER"
  )
  val unpublish: Unpublish? = null,
  @field:WireField(
    tag = 3,
    adapter = "gg.strims.ppspp.proto.DirectoryServerEvent${'$'}ViewerChange#ADAPTER"
  )
  val open: ViewerChange? = null,
  @field:WireField(
    tag = 4,
    adapter = "gg.strims.ppspp.proto.DirectoryServerEvent${'$'}Ping#ADAPTER"
  )
  val ping: Ping? = null,
  unknownFields: ByteString = ByteString.EMPTY
) : Message<DirectoryServerEvent, Nothing>(ADAPTER, unknownFields) {
  init {
    require(countNonNull(publish, unpublish, open, ping) <= 1) {
      "At most one of publish, unpublish, open, ping may be non-null"
    }
  }

  @Deprecated(
    message = "Shouldn't be used in Kotlin",
    level = DeprecationLevel.HIDDEN
  )
  override fun newBuilder(): Nothing = throw AssertionError()

  override fun equals(other: Any?): Boolean {
    if (other === this) return true
    if (other !is DirectoryServerEvent) return false
    if (unknownFields != other.unknownFields) return false
    if (publish != other.publish) return false
    if (unpublish != other.unpublish) return false
    if (open != other.open) return false
    if (ping != other.ping) return false
    return true
  }

  override fun hashCode(): Int {
    var result = super.hashCode
    if (result == 0) {
      result = unknownFields.hashCode()
      result = result * 37 + publish.hashCode()
      result = result * 37 + unpublish.hashCode()
      result = result * 37 + open.hashCode()
      result = result * 37 + ping.hashCode()
      super.hashCode = result
    }
    return result
  }

  override fun toString(): String {
    val result = mutableListOf<String>()
    if (publish != null) result += """publish=$publish"""
    if (unpublish != null) result += """unpublish=$unpublish"""
    if (open != null) result += """open=$open"""
    if (ping != null) result += """ping=$ping"""
    return result.joinToString(prefix = "DirectoryServerEvent{", separator = ", ", postfix = "}")
  }

  fun copy(
    publish: Publish? = this.publish,
    unpublish: Unpublish? = this.unpublish,
    open: ViewerChange? = this.open,
    ping: Ping? = this.ping,
    unknownFields: ByteString = this.unknownFields
  ): DirectoryServerEvent = DirectoryServerEvent(publish, unpublish, open, ping, unknownFields)

  companion object {
    @JvmField
    val ADAPTER: ProtoAdapter<DirectoryServerEvent> = object : ProtoAdapter<DirectoryServerEvent>(
      FieldEncoding.LENGTH_DELIMITED, 
      DirectoryServerEvent::class, 
      "type.googleapis.com/DirectoryServerEvent", 
      PROTO_3, 
      null
    ) {
      override fun encodedSize(value: DirectoryServerEvent): Int {
        var size = value.unknownFields.size
        size += Publish.ADAPTER.encodedSizeWithTag(1, value.publish)
        size += Unpublish.ADAPTER.encodedSizeWithTag(2, value.unpublish)
        size += ViewerChange.ADAPTER.encodedSizeWithTag(3, value.open)
        size += Ping.ADAPTER.encodedSizeWithTag(4, value.ping)
        return size
      }

      override fun encode(writer: ProtoWriter, value: DirectoryServerEvent) {
        Publish.ADAPTER.encodeWithTag(writer, 1, value.publish)
        Unpublish.ADAPTER.encodeWithTag(writer, 2, value.unpublish)
        ViewerChange.ADAPTER.encodeWithTag(writer, 3, value.open)
        Ping.ADAPTER.encodeWithTag(writer, 4, value.ping)
        writer.writeBytes(value.unknownFields)
      }

      override fun decode(reader: ProtoReader): DirectoryServerEvent {
        var publish: Publish? = null
        var unpublish: Unpublish? = null
        var open: ViewerChange? = null
        var ping: Ping? = null
        val unknownFields = reader.forEachTag { tag ->
          when (tag) {
            1 -> publish = Publish.ADAPTER.decode(reader)
            2 -> unpublish = Unpublish.ADAPTER.decode(reader)
            3 -> open = ViewerChange.ADAPTER.decode(reader)
            4 -> ping = Ping.ADAPTER.decode(reader)
            else -> reader.readUnknownField(tag)
          }
        }
        return DirectoryServerEvent(
          publish = publish,
          unpublish = unpublish,
          open = open,
          ping = ping,
          unknownFields = unknownFields
        )
      }

      override fun redact(value: DirectoryServerEvent): DirectoryServerEvent = value.copy(
        publish = value.publish?.let(Publish.ADAPTER::redact),
        unpublish = value.unpublish?.let(Unpublish.ADAPTER::redact),
        open = value.open?.let(ViewerChange.ADAPTER::redact),
        ping = value.ping?.let(Ping.ADAPTER::redact),
        unknownFields = ByteString.EMPTY
      )
    }

    private const val serialVersionUID: Long = 0L
  }

  class Publish(
    @field:WireField(
      tag = 1,
      adapter = "gg.strims.ppspp.proto.DirectoryListing#ADAPTER",
      label = WireField.Label.OMIT_IDENTITY
    )
    val listing: DirectoryListing? = null,
    unknownFields: ByteString = ByteString.EMPTY
  ) : Message<Publish, Nothing>(ADAPTER, unknownFields) {
    @Deprecated(
      message = "Shouldn't be used in Kotlin",
      level = DeprecationLevel.HIDDEN
    )
    override fun newBuilder(): Nothing = throw AssertionError()

    override fun equals(other: Any?): Boolean {
      if (other === this) return true
      if (other !is Publish) return false
      if (unknownFields != other.unknownFields) return false
      if (listing != other.listing) return false
      return true
    }

    override fun hashCode(): Int {
      var result = super.hashCode
      if (result == 0) {
        result = unknownFields.hashCode()
        result = result * 37 + listing.hashCode()
        super.hashCode = result
      }
      return result
    }

    override fun toString(): String {
      val result = mutableListOf<String>()
      if (listing != null) result += """listing=$listing"""
      return result.joinToString(prefix = "Publish{", separator = ", ", postfix = "}")
    }

    fun copy(listing: DirectoryListing? = this.listing, unknownFields: ByteString =
        this.unknownFields): Publish = Publish(listing, unknownFields)

    companion object {
      @JvmField
      val ADAPTER: ProtoAdapter<Publish> = object : ProtoAdapter<Publish>(
        FieldEncoding.LENGTH_DELIMITED, 
        Publish::class, 
        "type.googleapis.com/DirectoryServerEvent.Publish", 
        PROTO_3, 
        null
      ) {
        override fun encodedSize(value: Publish): Int {
          var size = value.unknownFields.size
          if (value.listing != null) size += DirectoryListing.ADAPTER.encodedSizeWithTag(1,
              value.listing)
          return size
        }

        override fun encode(writer: ProtoWriter, value: Publish) {
          if (value.listing != null) DirectoryListing.ADAPTER.encodeWithTag(writer, 1,
              value.listing)
          writer.writeBytes(value.unknownFields)
        }

        override fun decode(reader: ProtoReader): Publish {
          var listing: DirectoryListing? = null
          val unknownFields = reader.forEachTag { tag ->
            when (tag) {
              1 -> listing = DirectoryListing.ADAPTER.decode(reader)
              else -> reader.readUnknownField(tag)
            }
          }
          return Publish(
            listing = listing,
            unknownFields = unknownFields
          )
        }

        override fun redact(value: Publish): Publish = value.copy(
          listing = value.listing?.let(DirectoryListing.ADAPTER::redact),
          unknownFields = ByteString.EMPTY
        )
      }

      private const val serialVersionUID: Long = 0L
    }
  }

  class Unpublish(
    @field:WireField(
      tag = 1,
      adapter = "com.squareup.wire.ProtoAdapter#BYTES",
      label = WireField.Label.OMIT_IDENTITY
    )
    val key: ByteString = ByteString.EMPTY,
    unknownFields: ByteString = ByteString.EMPTY
  ) : Message<Unpublish, Nothing>(ADAPTER, unknownFields) {
    @Deprecated(
      message = "Shouldn't be used in Kotlin",
      level = DeprecationLevel.HIDDEN
    )
    override fun newBuilder(): Nothing = throw AssertionError()

    override fun equals(other: Any?): Boolean {
      if (other === this) return true
      if (other !is Unpublish) return false
      if (unknownFields != other.unknownFields) return false
      if (key != other.key) return false
      return true
    }

    override fun hashCode(): Int {
      var result = super.hashCode
      if (result == 0) {
        result = unknownFields.hashCode()
        result = result * 37 + key.hashCode()
        super.hashCode = result
      }
      return result
    }

    override fun toString(): String {
      val result = mutableListOf<String>()
      result += """key=$key"""
      return result.joinToString(prefix = "Unpublish{", separator = ", ", postfix = "}")
    }

    fun copy(key: ByteString = this.key, unknownFields: ByteString = this.unknownFields): Unpublish
        = Unpublish(key, unknownFields)

    companion object {
      @JvmField
      val ADAPTER: ProtoAdapter<Unpublish> = object : ProtoAdapter<Unpublish>(
        FieldEncoding.LENGTH_DELIMITED, 
        Unpublish::class, 
        "type.googleapis.com/DirectoryServerEvent.Unpublish", 
        PROTO_3, 
        null
      ) {
        override fun encodedSize(value: Unpublish): Int {
          var size = value.unknownFields.size
          if (value.key != ByteString.EMPTY) size += ProtoAdapter.BYTES.encodedSizeWithTag(1,
              value.key)
          return size
        }

        override fun encode(writer: ProtoWriter, value: Unpublish) {
          if (value.key != ByteString.EMPTY) ProtoAdapter.BYTES.encodeWithTag(writer, 1, value.key)
          writer.writeBytes(value.unknownFields)
        }

        override fun decode(reader: ProtoReader): Unpublish {
          var key: ByteString = ByteString.EMPTY
          val unknownFields = reader.forEachTag { tag ->
            when (tag) {
              1 -> key = ProtoAdapter.BYTES.decode(reader)
              else -> reader.readUnknownField(tag)
            }
          }
          return Unpublish(
            key = key,
            unknownFields = unknownFields
          )
        }

        override fun redact(value: Unpublish): Unpublish = value.copy(
          unknownFields = ByteString.EMPTY
        )
      }

      private const val serialVersionUID: Long = 0L
    }
  }

  class ViewerChange(
    @field:WireField(
      tag = 1,
      adapter = "com.squareup.wire.ProtoAdapter#BYTES",
      label = WireField.Label.OMIT_IDENTITY
    )
    val key: ByteString = ByteString.EMPTY,
    @field:WireField(
      tag = 2,
      adapter = "com.squareup.wire.ProtoAdapter#UINT32",
      label = WireField.Label.OMIT_IDENTITY
    )
    val count: Int = 0,
    unknownFields: ByteString = ByteString.EMPTY
  ) : Message<ViewerChange, Nothing>(ADAPTER, unknownFields) {
    @Deprecated(
      message = "Shouldn't be used in Kotlin",
      level = DeprecationLevel.HIDDEN
    )
    override fun newBuilder(): Nothing = throw AssertionError()

    override fun equals(other: Any?): Boolean {
      if (other === this) return true
      if (other !is ViewerChange) return false
      if (unknownFields != other.unknownFields) return false
      if (key != other.key) return false
      if (count != other.count) return false
      return true
    }

    override fun hashCode(): Int {
      var result = super.hashCode
      if (result == 0) {
        result = unknownFields.hashCode()
        result = result * 37 + key.hashCode()
        result = result * 37 + count.hashCode()
        super.hashCode = result
      }
      return result
    }

    override fun toString(): String {
      val result = mutableListOf<String>()
      result += """key=$key"""
      result += """count=$count"""
      return result.joinToString(prefix = "ViewerChange{", separator = ", ", postfix = "}")
    }

    fun copy(
      key: ByteString = this.key,
      count: Int = this.count,
      unknownFields: ByteString = this.unknownFields
    ): ViewerChange = ViewerChange(key, count, unknownFields)

    companion object {
      @JvmField
      val ADAPTER: ProtoAdapter<ViewerChange> = object : ProtoAdapter<ViewerChange>(
        FieldEncoding.LENGTH_DELIMITED, 
        ViewerChange::class, 
        "type.googleapis.com/DirectoryServerEvent.ViewerChange", 
        PROTO_3, 
        null
      ) {
        override fun encodedSize(value: ViewerChange): Int {
          var size = value.unknownFields.size
          if (value.key != ByteString.EMPTY) size += ProtoAdapter.BYTES.encodedSizeWithTag(1,
              value.key)
          if (value.count != 0) size += ProtoAdapter.UINT32.encodedSizeWithTag(2, value.count)
          return size
        }

        override fun encode(writer: ProtoWriter, value: ViewerChange) {
          if (value.key != ByteString.EMPTY) ProtoAdapter.BYTES.encodeWithTag(writer, 1, value.key)
          if (value.count != 0) ProtoAdapter.UINT32.encodeWithTag(writer, 2, value.count)
          writer.writeBytes(value.unknownFields)
        }

        override fun decode(reader: ProtoReader): ViewerChange {
          var key: ByteString = ByteString.EMPTY
          var count: Int = 0
          val unknownFields = reader.forEachTag { tag ->
            when (tag) {
              1 -> key = ProtoAdapter.BYTES.decode(reader)
              2 -> count = ProtoAdapter.UINT32.decode(reader)
              else -> reader.readUnknownField(tag)
            }
          }
          return ViewerChange(
            key = key,
            count = count,
            unknownFields = unknownFields
          )
        }

        override fun redact(value: ViewerChange): ViewerChange = value.copy(
          unknownFields = ByteString.EMPTY
        )
      }

      private const val serialVersionUID: Long = 0L
    }
  }

  class Ping(
    @field:WireField(
      tag = 1,
      adapter = "com.squareup.wire.ProtoAdapter#INT64",
      label = WireField.Label.OMIT_IDENTITY
    )
    val time: Long = 0L,
    unknownFields: ByteString = ByteString.EMPTY
  ) : Message<Ping, Nothing>(ADAPTER, unknownFields) {
    @Deprecated(
      message = "Shouldn't be used in Kotlin",
      level = DeprecationLevel.HIDDEN
    )
    override fun newBuilder(): Nothing = throw AssertionError()

    override fun equals(other: Any?): Boolean {
      if (other === this) return true
      if (other !is Ping) return false
      if (unknownFields != other.unknownFields) return false
      if (time != other.time) return false
      return true
    }

    override fun hashCode(): Int {
      var result = super.hashCode
      if (result == 0) {
        result = unknownFields.hashCode()
        result = result * 37 + time.hashCode()
        super.hashCode = result
      }
      return result
    }

    override fun toString(): String {
      val result = mutableListOf<String>()
      result += """time=$time"""
      return result.joinToString(prefix = "Ping{", separator = ", ", postfix = "}")
    }

    fun copy(time: Long = this.time, unknownFields: ByteString = this.unknownFields): Ping =
        Ping(time, unknownFields)

    companion object {
      @JvmField
      val ADAPTER: ProtoAdapter<Ping> = object : ProtoAdapter<Ping>(
        FieldEncoding.LENGTH_DELIMITED, 
        Ping::class, 
        "type.googleapis.com/DirectoryServerEvent.Ping", 
        PROTO_3, 
        null
      ) {
        override fun encodedSize(value: Ping): Int {
          var size = value.unknownFields.size
          if (value.time != 0L) size += ProtoAdapter.INT64.encodedSizeWithTag(1, value.time)
          return size
        }

        override fun encode(writer: ProtoWriter, value: Ping) {
          if (value.time != 0L) ProtoAdapter.INT64.encodeWithTag(writer, 1, value.time)
          writer.writeBytes(value.unknownFields)
        }

        override fun decode(reader: ProtoReader): Ping {
          var time: Long = 0L
          val unknownFields = reader.forEachTag { tag ->
            when (tag) {
              1 -> time = ProtoAdapter.INT64.decode(reader)
              else -> reader.readUnknownField(tag)
            }
          }
          return Ping(
            time = time,
            unknownFields = unknownFields
          )
        }

        override fun redact(value: Ping): Ping = value.copy(
          unknownFields = ByteString.EMPTY
        )
      }

      private const val serialVersionUID: Long = 0L
    }
  }
}
