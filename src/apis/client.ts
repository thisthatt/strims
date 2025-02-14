// Copyright 2022 Strims contributors
// SPDX-License-Identifier: AGPL-3.0-only

import { Readable, Writable } from "stream";

import Host from "@memelabs/protobuf/lib/rpc/host";

import { AuthFrontendClient } from "./strims/auth/v1/auth_rpc";
import { AutoseedFrontendClient } from "./strims/autoseed/v1/autoseed_rpc";
import { ChatFrontendClient, ChatServerFrontendClient } from "./strims/chat/v1/chat_rpc";
import { DebugClient } from "./strims/debug/v1/debug_rpc";
import { DevToolsClient as DevToolsServiceClient } from "./strims/devtools/v1/devtools_rpc";
import { CapConnClient } from "./strims/devtools/v1/ppspp/capconn_rpc";
import { FundingClient as FundingServiceClient } from "./strims/funding/v1/funding_rpc";
import { BootstrapFrontendClient } from "./strims/network/v1/bootstrap/bootstrap_rpc";
import { DirectoryFrontendClient } from "./strims/network/v1/directory/directory_rpc";
import { InviteLinkClient } from "./strims/network/v1/invite/invite_rpc";
import { NetworkFrontendClient } from "./strims/network/v1/network_rpc";
import { NotificationFrontendClient } from "./strims/notification/v1/notification_rpc";
import { ProfileFrontendClient } from "./strims/profile/v1/profile_rpc";
import { CaptureClient as VideoCaptureClient } from "./strims/video/v1/capture_rpc";
import { VideoChannelFrontendClient } from "./strims/video/v1/channel_rpc";
import { EgressClient as VideoEgressClient } from "./strims/video/v1/egress_rpc";
import { HLSEgressClient } from "./strims/video/v1/hls_egress_rpc";
import { VideoIngressClient } from "./strims/video/v1/ingress_rpc";
import { VNICFrontendClient } from "./strims/vnic/v1/vnic_rpc";

export class FrontendClient {
  public auth: AuthFrontendClient;
  public autoseed: AutoseedFrontendClient;
  public bootstrap: BootstrapFrontendClient;
  public chat: ChatFrontendClient;
  public chatServer: ChatServerFrontendClient;
  public debug: DebugClient;
  public directory: DirectoryFrontendClient;
  public network: NetworkFrontendClient;
  public notification: NotificationFrontendClient;
  public profile: ProfileFrontendClient;
  public videoCapture: VideoCaptureClient;
  public videoChannel: VideoChannelFrontendClient;
  public videoEgress: VideoEgressClient;
  public videoIngress: VideoIngressClient;
  public hlsEgress: HLSEgressClient;
  public vnic: VNICFrontendClient;

  constructor(w: Writable, r: Readable) {
    const host = new Host(w, r);
    this.auth = new AuthFrontendClient(host);
    this.autoseed = new AutoseedFrontendClient(host);
    this.bootstrap = new BootstrapFrontendClient(host);
    this.chat = new ChatFrontendClient(host);
    this.chatServer = new ChatServerFrontendClient(host);
    this.debug = new DebugClient(host);
    this.directory = new DirectoryFrontendClient(host);
    this.network = new NetworkFrontendClient(host);
    this.notification = new NotificationFrontendClient(host);
    this.profile = new ProfileFrontendClient(host);
    this.videoCapture = new VideoCaptureClient(host);
    this.videoChannel = new VideoChannelFrontendClient(host);
    this.videoEgress = new VideoEgressClient(host);
    this.videoIngress = new VideoIngressClient(host);
    this.hlsEgress = new HLSEgressClient(host);
    this.vnic = new VNICFrontendClient(host);
  }
}

export class FundingClient {
  public funding: FundingServiceClient;

  constructor(w: Writable, r: Readable) {
    const host = new Host(w, r);
    this.funding = new FundingServiceClient(host);
  }
}

export class DevToolsClient {
  public devTools: DevToolsServiceClient;
  public ppsppCapConn: CapConnClient;

  constructor(w: Writable, r: Readable) {
    const host = new Host(w, r);
    this.devTools = new DevToolsServiceClient(host);
    this.ppsppCapConn = new CapConnClient(host);
  }
}

export class InviteClient {
  public inviteLink: InviteLinkClient;

  constructor(w: Writable, r: Readable) {
    const host = new Host(w, r);
    this.inviteLink = new InviteLinkClient(host);
  }
}

export type Client = FrontendClient | FundingClient | DevToolsClient | InviteClient;
