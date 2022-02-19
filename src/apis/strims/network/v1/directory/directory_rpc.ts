import strims_rpc_Host, { UnaryCallOptions as strims_rpc_UnaryCallOptions } from "@memelabs/protobuf/lib/rpc/host";
import strims_rpc_Service from "@memelabs/protobuf/lib/rpc/service";
import { Call as strims_rpc_Call } from "@memelabs/protobuf/lib/apis/strims/rpc/rpc";
import { Readable as GenericReadable } from "@memelabs/protobuf/lib/rpc/stream";

import {
  IPublishRequest,
  PublishRequest,
  PublishResponse,
  IUnpublishRequest,
  UnpublishRequest,
  UnpublishResponse,
  IJoinRequest,
  JoinRequest,
  JoinResponse,
  IPartRequest,
  PartRequest,
  PartResponse,
  IPingRequest,
  PingRequest,
  PingResponse,
  IFrontendOpenRequest,
  FrontendOpenRequest,
  FrontendOpenResponse,
  IFrontendPublishRequest,
  FrontendPublishRequest,
  FrontendPublishResponse,
  IFrontendUnpublishRequest,
  FrontendUnpublishRequest,
  FrontendUnpublishResponse,
  IFrontendJoinRequest,
  FrontendJoinRequest,
  FrontendJoinResponse,
  IFrontendPartRequest,
  FrontendPartRequest,
  FrontendPartResponse,
  IFrontendTestRequest,
  FrontendTestRequest,
  FrontendTestResponse,
  IFrontendGetListingRecordRequest,
  FrontendGetListingRecordRequest,
  FrontendGetListingRecordResponse,
  IFrontendListListingRecordsRequest,
  FrontendListListingRecordsRequest,
  FrontendListListingRecordsResponse,
  IFrontendUpdateListingRecordRequest,
  FrontendUpdateListingRecordRequest,
  FrontendUpdateListingRecordResponse,
  ISnippetSubscribeRequest,
  SnippetSubscribeRequest,
  SnippetSubscribeResponse,
} from "./directory";

export interface DirectoryService {
  publish(req: PublishRequest, call: strims_rpc_Call): Promise<PublishResponse> | PublishResponse;
  unpublish(req: UnpublishRequest, call: strims_rpc_Call): Promise<UnpublishResponse> | UnpublishResponse;
  join(req: JoinRequest, call: strims_rpc_Call): Promise<JoinResponse> | JoinResponse;
  part(req: PartRequest, call: strims_rpc_Call): Promise<PartResponse> | PartResponse;
  ping(req: PingRequest, call: strims_rpc_Call): Promise<PingResponse> | PingResponse;
}

export class UnimplementedDirectoryService implements DirectoryService {
  publish(req: PublishRequest, call: strims_rpc_Call): Promise<PublishResponse> | PublishResponse { throw new Error("not implemented"); }
  unpublish(req: UnpublishRequest, call: strims_rpc_Call): Promise<UnpublishResponse> | UnpublishResponse { throw new Error("not implemented"); }
  join(req: JoinRequest, call: strims_rpc_Call): Promise<JoinResponse> | JoinResponse { throw new Error("not implemented"); }
  part(req: PartRequest, call: strims_rpc_Call): Promise<PartResponse> | PartResponse { throw new Error("not implemented"); }
  ping(req: PingRequest, call: strims_rpc_Call): Promise<PingResponse> | PingResponse { throw new Error("not implemented"); }
}

export const registerDirectoryService = (host: strims_rpc_Service, service: DirectoryService): void => {
  host.registerMethod<PublishRequest, PublishResponse>("strims.network.v1.directory.Directory.Publish", service.publish.bind(service), PublishRequest);
  host.registerMethod<UnpublishRequest, UnpublishResponse>("strims.network.v1.directory.Directory.Unpublish", service.unpublish.bind(service), UnpublishRequest);
  host.registerMethod<JoinRequest, JoinResponse>("strims.network.v1.directory.Directory.Join", service.join.bind(service), JoinRequest);
  host.registerMethod<PartRequest, PartResponse>("strims.network.v1.directory.Directory.Part", service.part.bind(service), PartRequest);
  host.registerMethod<PingRequest, PingResponse>("strims.network.v1.directory.Directory.Ping", service.ping.bind(service), PingRequest);
}

export class DirectoryClient {
  constructor(private readonly host: strims_rpc_Host) {}

  public publish(req?: IPublishRequest, opts?: strims_rpc_UnaryCallOptions): Promise<PublishResponse> {
    return this.host.expectOne(this.host.call("strims.network.v1.directory.Directory.Publish", new PublishRequest(req)), PublishResponse, opts);
  }

  public unpublish(req?: IUnpublishRequest, opts?: strims_rpc_UnaryCallOptions): Promise<UnpublishResponse> {
    return this.host.expectOne(this.host.call("strims.network.v1.directory.Directory.Unpublish", new UnpublishRequest(req)), UnpublishResponse, opts);
  }

  public join(req?: IJoinRequest, opts?: strims_rpc_UnaryCallOptions): Promise<JoinResponse> {
    return this.host.expectOne(this.host.call("strims.network.v1.directory.Directory.Join", new JoinRequest(req)), JoinResponse, opts);
  }

  public part(req?: IPartRequest, opts?: strims_rpc_UnaryCallOptions): Promise<PartResponse> {
    return this.host.expectOne(this.host.call("strims.network.v1.directory.Directory.Part", new PartRequest(req)), PartResponse, opts);
  }

  public ping(req?: IPingRequest, opts?: strims_rpc_UnaryCallOptions): Promise<PingResponse> {
    return this.host.expectOne(this.host.call("strims.network.v1.directory.Directory.Ping", new PingRequest(req)), PingResponse, opts);
  }
}

export interface DirectoryFrontendService {
  open(req: FrontendOpenRequest, call: strims_rpc_Call): GenericReadable<FrontendOpenResponse>;
  publish(req: FrontendPublishRequest, call: strims_rpc_Call): Promise<FrontendPublishResponse> | FrontendPublishResponse;
  unpublish(req: FrontendUnpublishRequest, call: strims_rpc_Call): Promise<FrontendUnpublishResponse> | FrontendUnpublishResponse;
  join(req: FrontendJoinRequest, call: strims_rpc_Call): Promise<FrontendJoinResponse> | FrontendJoinResponse;
  part(req: FrontendPartRequest, call: strims_rpc_Call): Promise<FrontendPartResponse> | FrontendPartResponse;
  test(req: FrontendTestRequest, call: strims_rpc_Call): Promise<FrontendTestResponse> | FrontendTestResponse;
  getListingRecord(req: FrontendGetListingRecordRequest, call: strims_rpc_Call): Promise<FrontendGetListingRecordResponse> | FrontendGetListingRecordResponse;
  listListingRecords(req: FrontendListListingRecordsRequest, call: strims_rpc_Call): Promise<FrontendListListingRecordsResponse> | FrontendListListingRecordsResponse;
  updateListingRecord(req: FrontendUpdateListingRecordRequest, call: strims_rpc_Call): Promise<FrontendUpdateListingRecordResponse> | FrontendUpdateListingRecordResponse;
}

export class UnimplementedDirectoryFrontendService implements DirectoryFrontendService {
  open(req: FrontendOpenRequest, call: strims_rpc_Call): GenericReadable<FrontendOpenResponse> { throw new Error("not implemented"); }
  publish(req: FrontendPublishRequest, call: strims_rpc_Call): Promise<FrontendPublishResponse> | FrontendPublishResponse { throw new Error("not implemented"); }
  unpublish(req: FrontendUnpublishRequest, call: strims_rpc_Call): Promise<FrontendUnpublishResponse> | FrontendUnpublishResponse { throw new Error("not implemented"); }
  join(req: FrontendJoinRequest, call: strims_rpc_Call): Promise<FrontendJoinResponse> | FrontendJoinResponse { throw new Error("not implemented"); }
  part(req: FrontendPartRequest, call: strims_rpc_Call): Promise<FrontendPartResponse> | FrontendPartResponse { throw new Error("not implemented"); }
  test(req: FrontendTestRequest, call: strims_rpc_Call): Promise<FrontendTestResponse> | FrontendTestResponse { throw new Error("not implemented"); }
  getListingRecord(req: FrontendGetListingRecordRequest, call: strims_rpc_Call): Promise<FrontendGetListingRecordResponse> | FrontendGetListingRecordResponse { throw new Error("not implemented"); }
  listListingRecords(req: FrontendListListingRecordsRequest, call: strims_rpc_Call): Promise<FrontendListListingRecordsResponse> | FrontendListListingRecordsResponse { throw new Error("not implemented"); }
  updateListingRecord(req: FrontendUpdateListingRecordRequest, call: strims_rpc_Call): Promise<FrontendUpdateListingRecordResponse> | FrontendUpdateListingRecordResponse { throw new Error("not implemented"); }
}

export const registerDirectoryFrontendService = (host: strims_rpc_Service, service: DirectoryFrontendService): void => {
  host.registerMethod<FrontendOpenRequest, FrontendOpenResponse>("strims.network.v1.directory.DirectoryFrontend.Open", service.open.bind(service), FrontendOpenRequest);
  host.registerMethod<FrontendPublishRequest, FrontendPublishResponse>("strims.network.v1.directory.DirectoryFrontend.Publish", service.publish.bind(service), FrontendPublishRequest);
  host.registerMethod<FrontendUnpublishRequest, FrontendUnpublishResponse>("strims.network.v1.directory.DirectoryFrontend.Unpublish", service.unpublish.bind(service), FrontendUnpublishRequest);
  host.registerMethod<FrontendJoinRequest, FrontendJoinResponse>("strims.network.v1.directory.DirectoryFrontend.Join", service.join.bind(service), FrontendJoinRequest);
  host.registerMethod<FrontendPartRequest, FrontendPartResponse>("strims.network.v1.directory.DirectoryFrontend.Part", service.part.bind(service), FrontendPartRequest);
  host.registerMethod<FrontendTestRequest, FrontendTestResponse>("strims.network.v1.directory.DirectoryFrontend.Test", service.test.bind(service), FrontendTestRequest);
  host.registerMethod<FrontendGetListingRecordRequest, FrontendGetListingRecordResponse>("strims.network.v1.directory.DirectoryFrontend.GetListingRecord", service.getListingRecord.bind(service), FrontendGetListingRecordRequest);
  host.registerMethod<FrontendListListingRecordsRequest, FrontendListListingRecordsResponse>("strims.network.v1.directory.DirectoryFrontend.ListListingRecords", service.listListingRecords.bind(service), FrontendListListingRecordsRequest);
  host.registerMethod<FrontendUpdateListingRecordRequest, FrontendUpdateListingRecordResponse>("strims.network.v1.directory.DirectoryFrontend.UpdateListingRecord", service.updateListingRecord.bind(service), FrontendUpdateListingRecordRequest);
}

export class DirectoryFrontendClient {
  constructor(private readonly host: strims_rpc_Host) {}

  public open(req?: IFrontendOpenRequest): GenericReadable<FrontendOpenResponse> {
    return this.host.expectMany(this.host.call("strims.network.v1.directory.DirectoryFrontend.Open", new FrontendOpenRequest(req)), FrontendOpenResponse);
  }

  public publish(req?: IFrontendPublishRequest, opts?: strims_rpc_UnaryCallOptions): Promise<FrontendPublishResponse> {
    return this.host.expectOne(this.host.call("strims.network.v1.directory.DirectoryFrontend.Publish", new FrontendPublishRequest(req)), FrontendPublishResponse, opts);
  }

  public unpublish(req?: IFrontendUnpublishRequest, opts?: strims_rpc_UnaryCallOptions): Promise<FrontendUnpublishResponse> {
    return this.host.expectOne(this.host.call("strims.network.v1.directory.DirectoryFrontend.Unpublish", new FrontendUnpublishRequest(req)), FrontendUnpublishResponse, opts);
  }

  public join(req?: IFrontendJoinRequest, opts?: strims_rpc_UnaryCallOptions): Promise<FrontendJoinResponse> {
    return this.host.expectOne(this.host.call("strims.network.v1.directory.DirectoryFrontend.Join", new FrontendJoinRequest(req)), FrontendJoinResponse, opts);
  }

  public part(req?: IFrontendPartRequest, opts?: strims_rpc_UnaryCallOptions): Promise<FrontendPartResponse> {
    return this.host.expectOne(this.host.call("strims.network.v1.directory.DirectoryFrontend.Part", new FrontendPartRequest(req)), FrontendPartResponse, opts);
  }

  public test(req?: IFrontendTestRequest, opts?: strims_rpc_UnaryCallOptions): Promise<FrontendTestResponse> {
    return this.host.expectOne(this.host.call("strims.network.v1.directory.DirectoryFrontend.Test", new FrontendTestRequest(req)), FrontendTestResponse, opts);
  }

  public getListingRecord(req?: IFrontendGetListingRecordRequest, opts?: strims_rpc_UnaryCallOptions): Promise<FrontendGetListingRecordResponse> {
    return this.host.expectOne(this.host.call("strims.network.v1.directory.DirectoryFrontend.GetListingRecord", new FrontendGetListingRecordRequest(req)), FrontendGetListingRecordResponse, opts);
  }

  public listListingRecords(req?: IFrontendListListingRecordsRequest, opts?: strims_rpc_UnaryCallOptions): Promise<FrontendListListingRecordsResponse> {
    return this.host.expectOne(this.host.call("strims.network.v1.directory.DirectoryFrontend.ListListingRecords", new FrontendListListingRecordsRequest(req)), FrontendListListingRecordsResponse, opts);
  }

  public updateListingRecord(req?: IFrontendUpdateListingRecordRequest, opts?: strims_rpc_UnaryCallOptions): Promise<FrontendUpdateListingRecordResponse> {
    return this.host.expectOne(this.host.call("strims.network.v1.directory.DirectoryFrontend.UpdateListingRecord", new FrontendUpdateListingRecordRequest(req)), FrontendUpdateListingRecordResponse, opts);
  }
}

export interface DirectorySnippetService {
  subscribe(req: SnippetSubscribeRequest, call: strims_rpc_Call): GenericReadable<SnippetSubscribeResponse>;
}

export class UnimplementedDirectorySnippetService implements DirectorySnippetService {
  subscribe(req: SnippetSubscribeRequest, call: strims_rpc_Call): GenericReadable<SnippetSubscribeResponse> { throw new Error("not implemented"); }
}

export const registerDirectorySnippetService = (host: strims_rpc_Service, service: DirectorySnippetService): void => {
  host.registerMethod<SnippetSubscribeRequest, SnippetSubscribeResponse>("strims.network.v1.directory.DirectorySnippet.Subscribe", service.subscribe.bind(service), SnippetSubscribeRequest);
}

export class DirectorySnippetClient {
  constructor(private readonly host: strims_rpc_Host) {}

  public subscribe(req?: ISnippetSubscribeRequest): GenericReadable<SnippetSubscribeResponse> {
    return this.host.expectMany(this.host.call("strims.network.v1.directory.DirectorySnippet.Subscribe", new SnippetSubscribeRequest(req)), SnippetSubscribeResponse);
  }
}

