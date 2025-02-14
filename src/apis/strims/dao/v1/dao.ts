import Reader from "@memelabs/protobuf/lib/pb/reader";
import Writer from "@memelabs/protobuf/lib/pb/writer";


export type ISecondaryIndexKey = {
  key?: Uint8Array;
  id?: bigint;
}

export class SecondaryIndexKey {
  key: Uint8Array;
  id: bigint;

  constructor(v?: ISecondaryIndexKey) {
    this.key = v?.key || new Uint8Array();
    this.id = v?.id || BigInt(0);
  }

  static encode(m: SecondaryIndexKey, w?: Writer): Writer {
    if (!w) w = new Writer();
    if (m.key.length) w.uint32(10).bytes(m.key);
    if (m.id) w.uint32(16).uint64(m.id);
    return w;
  }

  static decode(r: Reader | Uint8Array, length?: number): SecondaryIndexKey {
    r = r instanceof Reader ? r : new Reader(r);
    const end = length === undefined ? r.len : r.pos + length;
    const m = new SecondaryIndexKey();
    while (r.pos < end) {
      const tag = r.uint32();
      switch (tag >> 3) {
        case 1:
        m.key = r.bytes();
        break;
        case 2:
        m.id = r.uint64();
        break;
        default:
        r.skipType(tag & 7);
        break;
      }
    }
    return m;
  }
}

export type IMutex = {
  eol?: bigint;
  token?: Uint8Array;
}

export class Mutex {
  eol: bigint;
  token: Uint8Array;

  constructor(v?: IMutex) {
    this.eol = v?.eol || BigInt(0);
    this.token = v?.token || new Uint8Array();
  }

  static encode(m: Mutex, w?: Writer): Writer {
    if (!w) w = new Writer();
    if (m.eol) w.uint32(8).int64(m.eol);
    if (m.token.length) w.uint32(18).bytes(m.token);
    return w;
  }

  static decode(r: Reader | Uint8Array, length?: number): Mutex {
    r = r instanceof Reader ? r : new Reader(r);
    const end = length === undefined ? r.len : r.pos + length;
    const m = new Mutex();
    while (r.pos < end) {
      const tag = r.uint32();
      switch (tag >> 3) {
        case 1:
        m.eol = r.int64();
        break;
        case 2:
        m.token = r.bytes();
        break;
        default:
        r.skipType(tag & 7);
        break;
      }
    }
    return m;
  }
}

/* @internal */
export const strims_dao_v1_SecondaryIndexKey = SecondaryIndexKey;
/* @internal */
export type strims_dao_v1_SecondaryIndexKey = SecondaryIndexKey;
/* @internal */
export type strims_dao_v1_ISecondaryIndexKey = ISecondaryIndexKey;
/* @internal */
export const strims_dao_v1_Mutex = Mutex;
/* @internal */
export type strims_dao_v1_Mutex = Mutex;
/* @internal */
export type strims_dao_v1_IMutex = IMutex;
