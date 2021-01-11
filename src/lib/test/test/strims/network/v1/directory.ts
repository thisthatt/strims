import Reader from "../../../../pb/reader";
import Writer from "../../../../pb/writer";

import {
  Certificate as strims_type_Certificate,
  ICertificate as strims_type_ICertificate,
} from "../../type/certificate";

export interface IGetDirectoryEventsRequest {
  networkKey?: Uint8Array;
}

export class GetDirectoryEventsRequest {
  networkKey: Uint8Array = new Uint8Array();

  constructor(v?: IGetDirectoryEventsRequest) {
    this.networkKey = v?.networkKey || new Uint8Array();
  }

  static encode(m: GetDirectoryEventsRequest, w?: Writer): Writer {
    if (!w) w = new Writer(1024);
    if (m.networkKey) w.uint32(10).bytes(m.networkKey);
    return w;
  }

  static decode(r: Reader | Uint8Array, length?: number): GetDirectoryEventsRequest {
    r = r instanceof Reader ? r : new Reader(r);
    const end = length === undefined ? r.len : r.pos + length;
    const m = new GetDirectoryEventsRequest();
    while (r.pos < end) {
      const tag = r.uint32();
      switch (tag >> 3) {
        case 1:
        m.networkKey = r.bytes();
        break;
        default:
        r.skipType(tag & 7);
        break;
      }
    }
    return m;
  }
}

export interface ITestDirectoryPublishRequest {
  networkKey?: Uint8Array;
}

export class TestDirectoryPublishRequest {
  networkKey: Uint8Array = new Uint8Array();

  constructor(v?: ITestDirectoryPublishRequest) {
    this.networkKey = v?.networkKey || new Uint8Array();
  }

  static encode(m: TestDirectoryPublishRequest, w?: Writer): Writer {
    if (!w) w = new Writer(1024);
    if (m.networkKey) w.uint32(10).bytes(m.networkKey);
    return w;
  }

  static decode(r: Reader | Uint8Array, length?: number): TestDirectoryPublishRequest {
    r = r instanceof Reader ? r : new Reader(r);
    const end = length === undefined ? r.len : r.pos + length;
    const m = new TestDirectoryPublishRequest();
    while (r.pos < end) {
      const tag = r.uint32();
      switch (tag >> 3) {
        case 1:
        m.networkKey = r.bytes();
        break;
        default:
        r.skipType(tag & 7);
        break;
      }
    }
    return m;
  }
}

export interface ITestDirectoryPublishResponse {
}

export class TestDirectoryPublishResponse {

  constructor(v?: ITestDirectoryPublishResponse) {
    // noop
  }

  static encode(m: TestDirectoryPublishResponse, w?: Writer): Writer {
    if (!w) w = new Writer(1024);
    return w;
  }

  static decode(r: Reader | Uint8Array, length?: number): TestDirectoryPublishResponse {
    if (r instanceof Reader && length) r.skip(length);
    return new TestDirectoryPublishResponse();
  }
}

export interface IDirectoryListingSnippet {
  title?: string;
  description?: string;
  tags?: string[];
}

export class DirectoryListingSnippet {
  title: string = "";
  description: string = "";
  tags: string[] = [];

  constructor(v?: IDirectoryListingSnippet) {
    this.title = v?.title || "";
    this.description = v?.description || "";
    if (v?.tags) this.tags = v.tags;
  }

  static encode(m: DirectoryListingSnippet, w?: Writer): Writer {
    if (!w) w = new Writer(1024);
    if (m.title) w.uint32(10).string(m.title);
    if (m.description) w.uint32(18).string(m.description);
    for (const v of m.tags) w.uint32(26).string(v);
    return w;
  }

  static decode(r: Reader | Uint8Array, length?: number): DirectoryListingSnippet {
    r = r instanceof Reader ? r : new Reader(r);
    const end = length === undefined ? r.len : r.pos + length;
    const m = new DirectoryListingSnippet();
    while (r.pos < end) {
      const tag = r.uint32();
      switch (tag >> 3) {
        case 1:
        m.title = r.string();
        break;
        case 2:
        m.description = r.string();
        break;
        case 3:
        m.tags.push(r.string())
        break;
        default:
        r.skipType(tag & 7);
        break;
      }
    }
    return m;
  }
}

export interface IDirectoryListingMedia {
  startedAt?: bigint;
  mimeType?: string;
  bitrate?: number;
  swarmUri?: string;
}

export class DirectoryListingMedia {
  startedAt: bigint = BigInt(0);
  mimeType: string = "";
  bitrate: number = 0;
  swarmUri: string = "";

  constructor(v?: IDirectoryListingMedia) {
    this.startedAt = v?.startedAt || BigInt(0);
    this.mimeType = v?.mimeType || "";
    this.bitrate = v?.bitrate || 0;
    this.swarmUri = v?.swarmUri || "";
  }

  static encode(m: DirectoryListingMedia, w?: Writer): Writer {
    if (!w) w = new Writer(1024);
    if (m.startedAt) w.uint32(8).int64(m.startedAt);
    if (m.mimeType) w.uint32(18).string(m.mimeType);
    if (m.bitrate) w.uint32(24).uint32(m.bitrate);
    if (m.swarmUri) w.uint32(34).string(m.swarmUri);
    return w;
  }

  static decode(r: Reader | Uint8Array, length?: number): DirectoryListingMedia {
    r = r instanceof Reader ? r : new Reader(r);
    const end = length === undefined ? r.len : r.pos + length;
    const m = new DirectoryListingMedia();
    while (r.pos < end) {
      const tag = r.uint32();
      switch (tag >> 3) {
        case 1:
        m.startedAt = r.int64();
        break;
        case 2:
        m.mimeType = r.string();
        break;
        case 3:
        m.bitrate = r.uint32();
        break;
        case 4:
        m.swarmUri = r.string();
        break;
        default:
        r.skipType(tag & 7);
        break;
      }
    }
    return m;
  }
}

export interface IDirectoryListingService {
  type?: string;
}

export class DirectoryListingService {
  type: string = "";

  constructor(v?: IDirectoryListingService) {
    this.type = v?.type || "";
  }

  static encode(m: DirectoryListingService, w?: Writer): Writer {
    if (!w) w = new Writer(1024);
    if (m.type) w.uint32(10).string(m.type);
    return w;
  }

  static decode(r: Reader | Uint8Array, length?: number): DirectoryListingService {
    r = r instanceof Reader ? r : new Reader(r);
    const end = length === undefined ? r.len : r.pos + length;
    const m = new DirectoryListingService();
    while (r.pos < end) {
      const tag = r.uint32();
      switch (tag >> 3) {
        case 1:
        m.type = r.string();
        break;
        default:
        r.skipType(tag & 7);
        break;
      }
    }
    return m;
  }
}

export interface IDirectoryListing {
  creator?: strims_type_ICertificate | undefined;
  timestamp?: bigint;
  snippet?: IDirectoryListingSnippet | undefined;
  key?: Uint8Array;
  signature?: Uint8Array;
  content?: DirectoryListing.IContentOneOf
}

export class DirectoryListing {
  creator: strims_type_Certificate | undefined;
  timestamp: bigint = BigInt(0);
  snippet: DirectoryListingSnippet | undefined;
  key: Uint8Array = new Uint8Array();
  signature: Uint8Array = new Uint8Array();
  content: DirectoryListing.TContentOneOf;

  constructor(v?: IDirectoryListing) {
    this.creator = v?.creator && new strims_type_Certificate(v.creator);
    this.timestamp = v?.timestamp || BigInt(0);
    this.snippet = v?.snippet && new DirectoryListingSnippet(v.snippet);
    this.key = v?.key || new Uint8Array();
    this.signature = v?.signature || new Uint8Array();
    this.content = new DirectoryListing.ContentOneOf(v?.content);
  }

  static encode(m: DirectoryListing, w?: Writer): Writer {
    if (!w) w = new Writer(1024);
    if (m.creator) strims_type_Certificate.encode(m.creator, w.uint32(10).fork()).ldelim();
    if (m.timestamp) w.uint32(16).int64(m.timestamp);
    if (m.snippet) DirectoryListingSnippet.encode(m.snippet, w.uint32(26).fork()).ldelim();
    if (m.key) w.uint32(80010).bytes(m.key);
    if (m.signature) w.uint32(80018).bytes(m.signature);
    switch (m.content.case) {
      case DirectoryListing.ContentCase.MEDIA:
      DirectoryListingMedia.encode(m.content.media, w.uint32(8010).fork()).ldelim();
      break;
      case DirectoryListing.ContentCase.SERVICE:
      DirectoryListingService.encode(m.content.service, w.uint32(8018).fork()).ldelim();
      break;
    }
    return w;
  }

  static decode(r: Reader | Uint8Array, length?: number): DirectoryListing {
    r = r instanceof Reader ? r : new Reader(r);
    const end = length === undefined ? r.len : r.pos + length;
    const m = new DirectoryListing();
    while (r.pos < end) {
      const tag = r.uint32();
      switch (tag >> 3) {
        case 1:
        m.creator = strims_type_Certificate.decode(r, r.uint32());
        break;
        case 2:
        m.timestamp = r.int64();
        break;
        case 3:
        m.snippet = DirectoryListingSnippet.decode(r, r.uint32());
        break;
        case 1001:
        m.content = new DirectoryListing.ContentOneOf({ media: DirectoryListingMedia.decode(r, r.uint32()) });
        break;
        case 1002:
        m.content = new DirectoryListing.ContentOneOf({ service: DirectoryListingService.decode(r, r.uint32()) });
        break;
        case 10001:
        m.key = r.bytes();
        break;
        case 10002:
        m.signature = r.bytes();
        break;
        default:
        r.skipType(tag & 7);
        break;
      }
    }
    return m;
  }
}

export namespace DirectoryListing {
  export enum ContentCase {
    NOT_SET = 0,
    MEDIA = 1001,
    SERVICE = 1002,
  }

  export type IContentOneOf =
  { case?: ContentCase.NOT_SET }
  |{ case?: ContentCase.MEDIA, media: IDirectoryListingMedia }
  |{ case?: ContentCase.SERVICE, service: IDirectoryListingService }
  ;

  export type TContentOneOf = Readonly<
  { case: ContentCase.NOT_SET }
  |{ case: ContentCase.MEDIA, media: DirectoryListingMedia }
  |{ case: ContentCase.SERVICE, service: DirectoryListingService }
  >;

  class ContentOneOfImpl {
    media: DirectoryListingMedia;
    service: DirectoryListingService;
    case: ContentCase = ContentCase.NOT_SET;

    constructor(v?: IContentOneOf) {
      if (v && "media" in v) {
        this.case = ContentCase.MEDIA;
        this.media = new DirectoryListingMedia(v.media);
      } else
      if (v && "service" in v) {
        this.case = ContentCase.SERVICE;
        this.service = new DirectoryListingService(v.service);
      }
    }
  }

  export const ContentOneOf = ContentOneOfImpl as {
    new (): Readonly<{ case: ContentCase.NOT_SET }>;
    new <T extends IContentOneOf>(v: T): Readonly<
    T extends { media: IDirectoryListingMedia } ? { case: ContentCase.MEDIA, media: DirectoryListingMedia } :
    T extends { service: IDirectoryListingService } ? { case: ContentCase.SERVICE, service: DirectoryListingService } :
    never
    >;
  };

}

export interface IDirectoryEvent {
  body?: DirectoryEvent.IBodyOneOf
}

export class DirectoryEvent {
  body: DirectoryEvent.TBodyOneOf;

  constructor(v?: IDirectoryEvent) {
    this.body = new DirectoryEvent.BodyOneOf(v?.body);
  }

  static encode(m: DirectoryEvent, w?: Writer): Writer {
    if (!w) w = new Writer(1024);
    switch (m.body.case) {
      case DirectoryEvent.BodyCase.PUBLISH:
      DirectoryEvent.Publish.encode(m.body.publish, w.uint32(10).fork()).ldelim();
      break;
      case DirectoryEvent.BodyCase.UNPUBLISH:
      DirectoryEvent.Unpublish.encode(m.body.unpublish, w.uint32(18).fork()).ldelim();
      break;
      case DirectoryEvent.BodyCase.VIEWER_COUNT_CHANGE:
      DirectoryEvent.ViewerCountChange.encode(m.body.viewerCountChange, w.uint32(26).fork()).ldelim();
      break;
      case DirectoryEvent.BodyCase.VIEWER_STATE_CHANGE:
      DirectoryEvent.ViewerStateChange.encode(m.body.viewerStateChange, w.uint32(34).fork()).ldelim();
      break;
      case DirectoryEvent.BodyCase.PING:
      DirectoryEvent.Ping.encode(m.body.ping, w.uint32(42).fork()).ldelim();
      break;
      case DirectoryEvent.BodyCase.PADDING:
      DirectoryEvent.Padding.encode(m.body.padding, w.uint32(50).fork()).ldelim();
      break;
    }
    return w;
  }

  static decode(r: Reader | Uint8Array, length?: number): DirectoryEvent {
    r = r instanceof Reader ? r : new Reader(r);
    const end = length === undefined ? r.len : r.pos + length;
    const m = new DirectoryEvent();
    while (r.pos < end) {
      const tag = r.uint32();
      switch (tag >> 3) {
        case 1:
        m.body = new DirectoryEvent.BodyOneOf({ publish: DirectoryEvent.Publish.decode(r, r.uint32()) });
        break;
        case 2:
        m.body = new DirectoryEvent.BodyOneOf({ unpublish: DirectoryEvent.Unpublish.decode(r, r.uint32()) });
        break;
        case 3:
        m.body = new DirectoryEvent.BodyOneOf({ viewerCountChange: DirectoryEvent.ViewerCountChange.decode(r, r.uint32()) });
        break;
        case 4:
        m.body = new DirectoryEvent.BodyOneOf({ viewerStateChange: DirectoryEvent.ViewerStateChange.decode(r, r.uint32()) });
        break;
        case 5:
        m.body = new DirectoryEvent.BodyOneOf({ ping: DirectoryEvent.Ping.decode(r, r.uint32()) });
        break;
        case 6:
        m.body = new DirectoryEvent.BodyOneOf({ padding: DirectoryEvent.Padding.decode(r, r.uint32()) });
        break;
        default:
        r.skipType(tag & 7);
        break;
      }
    }
    return m;
  }
}

export namespace DirectoryEvent {
  export enum BodyCase {
    NOT_SET = 0,
    PUBLISH = 1,
    UNPUBLISH = 2,
    VIEWER_COUNT_CHANGE = 3,
    VIEWER_STATE_CHANGE = 4,
    PING = 5,
    PADDING = 6,
  }

  export type IBodyOneOf =
  { case?: BodyCase.NOT_SET }
  |{ case?: BodyCase.PUBLISH, publish: DirectoryEvent.IPublish }
  |{ case?: BodyCase.UNPUBLISH, unpublish: DirectoryEvent.IUnpublish }
  |{ case?: BodyCase.VIEWER_COUNT_CHANGE, viewerCountChange: DirectoryEvent.IViewerCountChange }
  |{ case?: BodyCase.VIEWER_STATE_CHANGE, viewerStateChange: DirectoryEvent.IViewerStateChange }
  |{ case?: BodyCase.PING, ping: DirectoryEvent.IPing }
  |{ case?: BodyCase.PADDING, padding: DirectoryEvent.IPadding }
  ;

  export type TBodyOneOf = Readonly<
  { case: BodyCase.NOT_SET }
  |{ case: BodyCase.PUBLISH, publish: DirectoryEvent.Publish }
  |{ case: BodyCase.UNPUBLISH, unpublish: DirectoryEvent.Unpublish }
  |{ case: BodyCase.VIEWER_COUNT_CHANGE, viewerCountChange: DirectoryEvent.ViewerCountChange }
  |{ case: BodyCase.VIEWER_STATE_CHANGE, viewerStateChange: DirectoryEvent.ViewerStateChange }
  |{ case: BodyCase.PING, ping: DirectoryEvent.Ping }
  |{ case: BodyCase.PADDING, padding: DirectoryEvent.Padding }
  >;

  class BodyOneOfImpl {
    publish: DirectoryEvent.Publish;
    unpublish: DirectoryEvent.Unpublish;
    viewerCountChange: DirectoryEvent.ViewerCountChange;
    viewerStateChange: DirectoryEvent.ViewerStateChange;
    ping: DirectoryEvent.Ping;
    padding: DirectoryEvent.Padding;
    case: BodyCase = BodyCase.NOT_SET;

    constructor(v?: IBodyOneOf) {
      if (v && "publish" in v) {
        this.case = BodyCase.PUBLISH;
        this.publish = new DirectoryEvent.Publish(v.publish);
      } else
      if (v && "unpublish" in v) {
        this.case = BodyCase.UNPUBLISH;
        this.unpublish = new DirectoryEvent.Unpublish(v.unpublish);
      } else
      if (v && "viewerCountChange" in v) {
        this.case = BodyCase.VIEWER_COUNT_CHANGE;
        this.viewerCountChange = new DirectoryEvent.ViewerCountChange(v.viewerCountChange);
      } else
      if (v && "viewerStateChange" in v) {
        this.case = BodyCase.VIEWER_STATE_CHANGE;
        this.viewerStateChange = new DirectoryEvent.ViewerStateChange(v.viewerStateChange);
      } else
      if (v && "ping" in v) {
        this.case = BodyCase.PING;
        this.ping = new DirectoryEvent.Ping(v.ping);
      } else
      if (v && "padding" in v) {
        this.case = BodyCase.PADDING;
        this.padding = new DirectoryEvent.Padding(v.padding);
      }
    }
  }

  export const BodyOneOf = BodyOneOfImpl as {
    new (): Readonly<{ case: BodyCase.NOT_SET }>;
    new <T extends IBodyOneOf>(v: T): Readonly<
    T extends { publish: DirectoryEvent.IPublish } ? { case: BodyCase.PUBLISH, publish: DirectoryEvent.Publish } :
    T extends { unpublish: DirectoryEvent.IUnpublish } ? { case: BodyCase.UNPUBLISH, unpublish: DirectoryEvent.Unpublish } :
    T extends { viewerCountChange: DirectoryEvent.IViewerCountChange } ? { case: BodyCase.VIEWER_COUNT_CHANGE, viewerCountChange: DirectoryEvent.ViewerCountChange } :
    T extends { viewerStateChange: DirectoryEvent.IViewerStateChange } ? { case: BodyCase.VIEWER_STATE_CHANGE, viewerStateChange: DirectoryEvent.ViewerStateChange } :
    T extends { ping: DirectoryEvent.IPing } ? { case: BodyCase.PING, ping: DirectoryEvent.Ping } :
    T extends { padding: DirectoryEvent.IPadding } ? { case: BodyCase.PADDING, padding: DirectoryEvent.Padding } :
    never
    >;
  };

  export interface IPublish {
    listing?: IDirectoryListing | undefined;
  }

  export class Publish {
    listing: DirectoryListing | undefined;

    constructor(v?: IPublish) {
      this.listing = v?.listing && new DirectoryListing(v.listing);
    }

    static encode(m: Publish, w?: Writer): Writer {
      if (!w) w = new Writer(1024);
      if (m.listing) DirectoryListing.encode(m.listing, w.uint32(10).fork()).ldelim();
      return w;
    }

    static decode(r: Reader | Uint8Array, length?: number): Publish {
      r = r instanceof Reader ? r : new Reader(r);
      const end = length === undefined ? r.len : r.pos + length;
      const m = new Publish();
      while (r.pos < end) {
        const tag = r.uint32();
        switch (tag >> 3) {
          case 1:
          m.listing = DirectoryListing.decode(r, r.uint32());
          break;
          default:
          r.skipType(tag & 7);
          break;
        }
      }
      return m;
    }
  }

  export interface IUnpublish {
    key?: Uint8Array;
  }

  export class Unpublish {
    key: Uint8Array = new Uint8Array();

    constructor(v?: IUnpublish) {
      this.key = v?.key || new Uint8Array();
    }

    static encode(m: Unpublish, w?: Writer): Writer {
      if (!w) w = new Writer(1024);
      if (m.key) w.uint32(10).bytes(m.key);
      return w;
    }

    static decode(r: Reader | Uint8Array, length?: number): Unpublish {
      r = r instanceof Reader ? r : new Reader(r);
      const end = length === undefined ? r.len : r.pos + length;
      const m = new Unpublish();
      while (r.pos < end) {
        const tag = r.uint32();
        switch (tag >> 3) {
          case 1:
          m.key = r.bytes();
          break;
          default:
          r.skipType(tag & 7);
          break;
        }
      }
      return m;
    }
  }

  export interface IViewerCountChange {
    key?: Uint8Array;
    count?: number;
  }

  export class ViewerCountChange {
    key: Uint8Array = new Uint8Array();
    count: number = 0;

    constructor(v?: IViewerCountChange) {
      this.key = v?.key || new Uint8Array();
      this.count = v?.count || 0;
    }

    static encode(m: ViewerCountChange, w?: Writer): Writer {
      if (!w) w = new Writer(1024);
      if (m.key) w.uint32(10).bytes(m.key);
      if (m.count) w.uint32(16).uint32(m.count);
      return w;
    }

    static decode(r: Reader | Uint8Array, length?: number): ViewerCountChange {
      r = r instanceof Reader ? r : new Reader(r);
      const end = length === undefined ? r.len : r.pos + length;
      const m = new ViewerCountChange();
      while (r.pos < end) {
        const tag = r.uint32();
        switch (tag >> 3) {
          case 1:
          m.key = r.bytes();
          break;
          case 2:
          m.count = r.uint32();
          break;
          default:
          r.skipType(tag & 7);
          break;
        }
      }
      return m;
    }
  }

  export interface IViewerStateChange {
    subject?: string;
    online?: boolean;
    viewingKeys?: Uint8Array[];
  }

  export class ViewerStateChange {
    subject: string = "";
    online: boolean = false;
    viewingKeys: Uint8Array[] = [];

    constructor(v?: IViewerStateChange) {
      this.subject = v?.subject || "";
      this.online = v?.online || false;
      if (v?.viewingKeys) this.viewingKeys = v.viewingKeys;
    }

    static encode(m: ViewerStateChange, w?: Writer): Writer {
      if (!w) w = new Writer(1024);
      if (m.subject) w.uint32(10).string(m.subject);
      if (m.online) w.uint32(16).bool(m.online);
      for (const v of m.viewingKeys) w.uint32(26).bytes(v);
      return w;
    }

    static decode(r: Reader | Uint8Array, length?: number): ViewerStateChange {
      r = r instanceof Reader ? r : new Reader(r);
      const end = length === undefined ? r.len : r.pos + length;
      const m = new ViewerStateChange();
      while (r.pos < end) {
        const tag = r.uint32();
        switch (tag >> 3) {
          case 1:
          m.subject = r.string();
          break;
          case 2:
          m.online = r.bool();
          break;
          case 3:
          m.viewingKeys.push(r.bytes())
          break;
          default:
          r.skipType(tag & 7);
          break;
        }
      }
      return m;
    }
  }

  export interface IPing {
    time?: bigint;
  }

  export class Ping {
    time: bigint = BigInt(0);

    constructor(v?: IPing) {
      this.time = v?.time || BigInt(0);
    }

    static encode(m: Ping, w?: Writer): Writer {
      if (!w) w = new Writer(1024);
      if (m.time) w.uint32(8).int64(m.time);
      return w;
    }

    static decode(r: Reader | Uint8Array, length?: number): Ping {
      r = r instanceof Reader ? r : new Reader(r);
      const end = length === undefined ? r.len : r.pos + length;
      const m = new Ping();
      while (r.pos < end) {
        const tag = r.uint32();
        switch (tag >> 3) {
          case 1:
          m.time = r.int64();
          break;
          default:
          r.skipType(tag & 7);
          break;
        }
      }
      return m;
    }
  }

  export interface IPadding {
    data?: Uint8Array;
  }

  export class Padding {
    data: Uint8Array = new Uint8Array();

    constructor(v?: IPadding) {
      this.data = v?.data || new Uint8Array();
    }

    static encode(m: Padding, w?: Writer): Writer {
      if (!w) w = new Writer(1024);
      if (m.data) w.uint32(10).bytes(m.data);
      return w;
    }

    static decode(r: Reader | Uint8Array, length?: number): Padding {
      r = r instanceof Reader ? r : new Reader(r);
      const end = length === undefined ? r.len : r.pos + length;
      const m = new Padding();
      while (r.pos < end) {
        const tag = r.uint32();
        switch (tag >> 3) {
          case 1:
          m.data = r.bytes();
          break;
          default:
          r.skipType(tag & 7);
          break;
        }
      }
      return m;
    }
  }

}

export interface IDirectoryPublishRequest {
  listing?: IDirectoryListing | undefined;
}

export class DirectoryPublishRequest {
  listing: DirectoryListing | undefined;

  constructor(v?: IDirectoryPublishRequest) {
    this.listing = v?.listing && new DirectoryListing(v.listing);
  }

  static encode(m: DirectoryPublishRequest, w?: Writer): Writer {
    if (!w) w = new Writer(1024);
    if (m.listing) DirectoryListing.encode(m.listing, w.uint32(10).fork()).ldelim();
    return w;
  }

  static decode(r: Reader | Uint8Array, length?: number): DirectoryPublishRequest {
    r = r instanceof Reader ? r : new Reader(r);
    const end = length === undefined ? r.len : r.pos + length;
    const m = new DirectoryPublishRequest();
    while (r.pos < end) {
      const tag = r.uint32();
      switch (tag >> 3) {
        case 1:
        m.listing = DirectoryListing.decode(r, r.uint32());
        break;
        default:
        r.skipType(tag & 7);
        break;
      }
    }
    return m;
  }
}

export interface IDirectoryPublishResponse {
}

export class DirectoryPublishResponse {

  constructor(v?: IDirectoryPublishResponse) {
    // noop
  }

  static encode(m: DirectoryPublishResponse, w?: Writer): Writer {
    if (!w) w = new Writer(1024);
    return w;
  }

  static decode(r: Reader | Uint8Array, length?: number): DirectoryPublishResponse {
    if (r instanceof Reader && length) r.skip(length);
    return new DirectoryPublishResponse();
  }
}

export interface IDirectoryUnpublishRequest {
  key?: Uint8Array;
}

export class DirectoryUnpublishRequest {
  key: Uint8Array = new Uint8Array();

  constructor(v?: IDirectoryUnpublishRequest) {
    this.key = v?.key || new Uint8Array();
  }

  static encode(m: DirectoryUnpublishRequest, w?: Writer): Writer {
    if (!w) w = new Writer(1024);
    if (m.key) w.uint32(10).bytes(m.key);
    return w;
  }

  static decode(r: Reader | Uint8Array, length?: number): DirectoryUnpublishRequest {
    r = r instanceof Reader ? r : new Reader(r);
    const end = length === undefined ? r.len : r.pos + length;
    const m = new DirectoryUnpublishRequest();
    while (r.pos < end) {
      const tag = r.uint32();
      switch (tag >> 3) {
        case 1:
        m.key = r.bytes();
        break;
        default:
        r.skipType(tag & 7);
        break;
      }
    }
    return m;
  }
}

export interface IDirectoryUnpublishResponse {
}

export class DirectoryUnpublishResponse {

  constructor(v?: IDirectoryUnpublishResponse) {
    // noop
  }

  static encode(m: DirectoryUnpublishResponse, w?: Writer): Writer {
    if (!w) w = new Writer(1024);
    return w;
  }

  static decode(r: Reader | Uint8Array, length?: number): DirectoryUnpublishResponse {
    if (r instanceof Reader && length) r.skip(length);
    return new DirectoryUnpublishResponse();
  }
}

export interface IDirectoryJoinRequest {
  key?: Uint8Array;
}

export class DirectoryJoinRequest {
  key: Uint8Array = new Uint8Array();

  constructor(v?: IDirectoryJoinRequest) {
    this.key = v?.key || new Uint8Array();
  }

  static encode(m: DirectoryJoinRequest, w?: Writer): Writer {
    if (!w) w = new Writer(1024);
    if (m.key) w.uint32(10).bytes(m.key);
    return w;
  }

  static decode(r: Reader | Uint8Array, length?: number): DirectoryJoinRequest {
    r = r instanceof Reader ? r : new Reader(r);
    const end = length === undefined ? r.len : r.pos + length;
    const m = new DirectoryJoinRequest();
    while (r.pos < end) {
      const tag = r.uint32();
      switch (tag >> 3) {
        case 1:
        m.key = r.bytes();
        break;
        default:
        r.skipType(tag & 7);
        break;
      }
    }
    return m;
  }
}

export interface IDirectoryJoinResponse {
}

export class DirectoryJoinResponse {

  constructor(v?: IDirectoryJoinResponse) {
    // noop
  }

  static encode(m: DirectoryJoinResponse, w?: Writer): Writer {
    if (!w) w = new Writer(1024);
    return w;
  }

  static decode(r: Reader | Uint8Array, length?: number): DirectoryJoinResponse {
    if (r instanceof Reader && length) r.skip(length);
    return new DirectoryJoinResponse();
  }
}

export interface IDirectoryPartRequest {
  key?: Uint8Array;
}

export class DirectoryPartRequest {
  key: Uint8Array = new Uint8Array();

  constructor(v?: IDirectoryPartRequest) {
    this.key = v?.key || new Uint8Array();
  }

  static encode(m: DirectoryPartRequest, w?: Writer): Writer {
    if (!w) w = new Writer(1024);
    if (m.key) w.uint32(10).bytes(m.key);
    return w;
  }

  static decode(r: Reader | Uint8Array, length?: number): DirectoryPartRequest {
    r = r instanceof Reader ? r : new Reader(r);
    const end = length === undefined ? r.len : r.pos + length;
    const m = new DirectoryPartRequest();
    while (r.pos < end) {
      const tag = r.uint32();
      switch (tag >> 3) {
        case 1:
        m.key = r.bytes();
        break;
        default:
        r.skipType(tag & 7);
        break;
      }
    }
    return m;
  }
}

export interface IDirectoryPartResponse {
}

export class DirectoryPartResponse {

  constructor(v?: IDirectoryPartResponse) {
    // noop
  }

  static encode(m: DirectoryPartResponse, w?: Writer): Writer {
    if (!w) w = new Writer(1024);
    return w;
  }

  static decode(r: Reader | Uint8Array, length?: number): DirectoryPartResponse {
    if (r instanceof Reader && length) r.skip(length);
    return new DirectoryPartResponse();
  }
}

export interface IDirectoryPingRequest {
}

export class DirectoryPingRequest {

  constructor(v?: IDirectoryPingRequest) {
    // noop
  }

  static encode(m: DirectoryPingRequest, w?: Writer): Writer {
    if (!w) w = new Writer(1024);
    return w;
  }

  static decode(r: Reader | Uint8Array, length?: number): DirectoryPingRequest {
    if (r instanceof Reader && length) r.skip(length);
    return new DirectoryPingRequest();
  }
}

export interface IDirectoryPingResponse {
}

export class DirectoryPingResponse {

  constructor(v?: IDirectoryPingResponse) {
    // noop
  }

  static encode(m: DirectoryPingResponse, w?: Writer): Writer {
    if (!w) w = new Writer(1024);
    return w;
  }

  static decode(r: Reader | Uint8Array, length?: number): DirectoryPingResponse {
    if (r instanceof Reader && length) r.skip(length);
    return new DirectoryPingResponse();
  }
}

export interface IDirectoryFrontendOpenRequest {
  networkKey?: Uint8Array;
}

export class DirectoryFrontendOpenRequest {
  networkKey: Uint8Array = new Uint8Array();

  constructor(v?: IDirectoryFrontendOpenRequest) {
    this.networkKey = v?.networkKey || new Uint8Array();
  }

  static encode(m: DirectoryFrontendOpenRequest, w?: Writer): Writer {
    if (!w) w = new Writer(1024);
    if (m.networkKey) w.uint32(10).bytes(m.networkKey);
    return w;
  }

  static decode(r: Reader | Uint8Array, length?: number): DirectoryFrontendOpenRequest {
    r = r instanceof Reader ? r : new Reader(r);
    const end = length === undefined ? r.len : r.pos + length;
    const m = new DirectoryFrontendOpenRequest();
    while (r.pos < end) {
      const tag = r.uint32();
      switch (tag >> 3) {
        case 1:
        m.networkKey = r.bytes();
        break;
        default:
        r.skipType(tag & 7);
        break;
      }
    }
    return m;
  }
}

export interface IDirectoryFrontendOpenResponse {
  event?: IDirectoryEvent | undefined;
}

export class DirectoryFrontendOpenResponse {
  event: DirectoryEvent | undefined;

  constructor(v?: IDirectoryFrontendOpenResponse) {
    this.event = v?.event && new DirectoryEvent(v.event);
  }

  static encode(m: DirectoryFrontendOpenResponse, w?: Writer): Writer {
    if (!w) w = new Writer(1024);
    if (m.event) DirectoryEvent.encode(m.event, w.uint32(10).fork()).ldelim();
    return w;
  }

  static decode(r: Reader | Uint8Array, length?: number): DirectoryFrontendOpenResponse {
    r = r instanceof Reader ? r : new Reader(r);
    const end = length === undefined ? r.len : r.pos + length;
    const m = new DirectoryFrontendOpenResponse();
    while (r.pos < end) {
      const tag = r.uint32();
      switch (tag >> 3) {
        case 1:
        m.event = DirectoryEvent.decode(r, r.uint32());
        break;
        default:
        r.skipType(tag & 7);
        break;
      }
    }
    return m;
  }
}

export interface IDirectoryFrontendTestRequest {
  networkKey?: Uint8Array;
}

export class DirectoryFrontendTestRequest {
  networkKey: Uint8Array = new Uint8Array();

  constructor(v?: IDirectoryFrontendTestRequest) {
    this.networkKey = v?.networkKey || new Uint8Array();
  }

  static encode(m: DirectoryFrontendTestRequest, w?: Writer): Writer {
    if (!w) w = new Writer(1024);
    if (m.networkKey) w.uint32(10).bytes(m.networkKey);
    return w;
  }

  static decode(r: Reader | Uint8Array, length?: number): DirectoryFrontendTestRequest {
    r = r instanceof Reader ? r : new Reader(r);
    const end = length === undefined ? r.len : r.pos + length;
    const m = new DirectoryFrontendTestRequest();
    while (r.pos < end) {
      const tag = r.uint32();
      switch (tag >> 3) {
        case 1:
        m.networkKey = r.bytes();
        break;
        default:
        r.skipType(tag & 7);
        break;
      }
    }
    return m;
  }
}

export interface IDirectoryFrontendTestResponse {
}

export class DirectoryFrontendTestResponse {

  constructor(v?: IDirectoryFrontendTestResponse) {
    // noop
  }

  static encode(m: DirectoryFrontendTestResponse, w?: Writer): Writer {
    if (!w) w = new Writer(1024);
    return w;
  }

  static decode(r: Reader | Uint8Array, length?: number): DirectoryFrontendTestResponse {
    if (r instanceof Reader && length) r.skip(length);
    return new DirectoryFrontendTestResponse();
  }
}

