// Code generated by Wire protocol buffer compiler, do not edit.
// Source: Profile in profile.proto
package gg.strims.ppspp.proto

import com.squareup.wire.FieldEncoding
import com.squareup.wire.Message
import com.squareup.wire.ProtoAdapter
import com.squareup.wire.ProtoReader
import com.squareup.wire.ProtoWriter
import com.squareup.wire.Syntax.PROTO_3
import com.squareup.wire.WireField
import com.squareup.wire.internal.immutableCopyOf
import com.squareup.wire.internal.redactElements
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

class Profile(
  @field:WireField(
    tag = 1,
    adapter = "com.squareup.wire.ProtoAdapter#UINT64",
    label = WireField.Label.OMIT_IDENTITY
  )
  val id: Long = 0L,
  @field:WireField(
    tag = 2,
    adapter = "com.squareup.wire.ProtoAdapter#STRING",
    label = WireField.Label.OMIT_IDENTITY
  )
  val name: String = "",
  @field:WireField(
    tag = 3,
    adapter = "com.squareup.wire.ProtoAdapter#BYTES",
    label = WireField.Label.OMIT_IDENTITY
  )
  val secret: ByteString = ByteString.EMPTY,
  @field:WireField(
    tag = 4,
    adapter = "gg.strims.ppspp.proto.Key#ADAPTER",
    label = WireField.Label.OMIT_IDENTITY
  )
  val key: Key? = null,
  networks: List<Network> = emptyList(),
  network_memberships: List<NetworkMembership> = emptyList(),
  unknownFields: ByteString = ByteString.EMPTY
) : Message<Profile, Nothing>(ADAPTER, unknownFields) {
  @field:WireField(
    tag = 5,
    adapter = "gg.strims.ppspp.proto.Network#ADAPTER",
    label = WireField.Label.REPEATED
  )
  val networks: List<Network> = immutableCopyOf("networks", networks)

  @field:WireField(
    tag = 6,
    adapter = "gg.strims.ppspp.proto.NetworkMembership#ADAPTER",
    label = WireField.Label.REPEATED,
    jsonName = "networkMemberships"
  )
  val network_memberships: List<NetworkMembership> = immutableCopyOf("network_memberships",
      network_memberships)

  @Deprecated(
    message = "Shouldn't be used in Kotlin",
    level = DeprecationLevel.HIDDEN
  )
  override fun newBuilder(): Nothing = throw AssertionError()

  override fun equals(other: Any?): Boolean {
    if (other === this) return true
    if (other !is Profile) return false
    if (unknownFields != other.unknownFields) return false
    if (id != other.id) return false
    if (name != other.name) return false
    if (secret != other.secret) return false
    if (key != other.key) return false
    if (networks != other.networks) return false
    if (network_memberships != other.network_memberships) return false
    return true
  }

  override fun hashCode(): Int {
    var result = super.hashCode
    if (result == 0) {
      result = unknownFields.hashCode()
      result = result * 37 + id.hashCode()
      result = result * 37 + name.hashCode()
      result = result * 37 + secret.hashCode()
      result = result * 37 + key.hashCode()
      result = result * 37 + networks.hashCode()
      result = result * 37 + network_memberships.hashCode()
      super.hashCode = result
    }
    return result
  }

  override fun toString(): String {
    val result = mutableListOf<String>()
    result += """id=$id"""
    result += """name=${sanitize(name)}"""
    result += """secret=$secret"""
    if (key != null) result += """key=$key"""
    if (networks.isNotEmpty()) result += """networks=$networks"""
    if (network_memberships.isNotEmpty()) result += """network_memberships=$network_memberships"""
    return result.joinToString(prefix = "Profile{", separator = ", ", postfix = "}")
  }

  fun copy(
    id: Long = this.id,
    name: String = this.name,
    secret: ByteString = this.secret,
    key: Key? = this.key,
    networks: List<Network> = this.networks,
    network_memberships: List<NetworkMembership> = this.network_memberships,
    unknownFields: ByteString = this.unknownFields
  ): Profile = Profile(id, name, secret, key, networks, network_memberships, unknownFields)

  companion object {
    @JvmField
    val ADAPTER: ProtoAdapter<Profile> = object : ProtoAdapter<Profile>(
      FieldEncoding.LENGTH_DELIMITED, 
      Profile::class, 
      "type.googleapis.com/Profile", 
      PROTO_3, 
      null
    ) {
      override fun encodedSize(value: Profile): Int {
        var size = value.unknownFields.size
        if (value.id != 0L) size += ProtoAdapter.UINT64.encodedSizeWithTag(1, value.id)
        if (value.name != "") size += ProtoAdapter.STRING.encodedSizeWithTag(2, value.name)
        if (value.secret != ByteString.EMPTY) size += ProtoAdapter.BYTES.encodedSizeWithTag(3,
            value.secret)
        if (value.key != null) size += Key.ADAPTER.encodedSizeWithTag(4, value.key)
        size += Network.ADAPTER.asRepeated().encodedSizeWithTag(5, value.networks)
        size += NetworkMembership.ADAPTER.asRepeated().encodedSizeWithTag(6,
            value.network_memberships)
        return size
      }

      override fun encode(writer: ProtoWriter, value: Profile) {
        if (value.id != 0L) ProtoAdapter.UINT64.encodeWithTag(writer, 1, value.id)
        if (value.name != "") ProtoAdapter.STRING.encodeWithTag(writer, 2, value.name)
        if (value.secret != ByteString.EMPTY) ProtoAdapter.BYTES.encodeWithTag(writer, 3,
            value.secret)
        if (value.key != null) Key.ADAPTER.encodeWithTag(writer, 4, value.key)
        Network.ADAPTER.asRepeated().encodeWithTag(writer, 5, value.networks)
        NetworkMembership.ADAPTER.asRepeated().encodeWithTag(writer, 6, value.network_memberships)
        writer.writeBytes(value.unknownFields)
      }

      override fun decode(reader: ProtoReader): Profile {
        var id: Long = 0L
        var name: String = ""
        var secret: ByteString = ByteString.EMPTY
        var key: Key? = null
        val networks = mutableListOf<Network>()
        val network_memberships = mutableListOf<NetworkMembership>()
        val unknownFields = reader.forEachTag { tag ->
          when (tag) {
            1 -> id = ProtoAdapter.UINT64.decode(reader)
            2 -> name = ProtoAdapter.STRING.decode(reader)
            3 -> secret = ProtoAdapter.BYTES.decode(reader)
            4 -> key = Key.ADAPTER.decode(reader)
            5 -> networks.add(Network.ADAPTER.decode(reader))
            6 -> network_memberships.add(NetworkMembership.ADAPTER.decode(reader))
            else -> reader.readUnknownField(tag)
          }
        }
        return Profile(
          id = id,
          name = name,
          secret = secret,
          key = key,
          networks = networks,
          network_memberships = network_memberships,
          unknownFields = unknownFields
        )
      }

      override fun redact(value: Profile): Profile = value.copy(
        key = value.key?.let(Key.ADAPTER::redact),
        networks = value.networks.redactElements(Network.ADAPTER),
        network_memberships = value.network_memberships.redactElements(NetworkMembership.ADAPTER),
        unknownFields = ByteString.EMPTY
      )
    }

    private const val serialVersionUID: Long = 0L
  }
}
