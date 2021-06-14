import { EventEmitter } from "events";

declare function postMessage<T>(message: T, targetOrigin?: string, transfer?: Transferable[]): void;
declare function postMessage<T>(message: T, transfer?: Transferable[]): void;

interface GenericMessagePort<TIn, TOut> extends MessagePort {
  onmessage: ((this: GenericMessagePort<TIn, TOut>, ev: MessageEvent<TIn>) => void) | null;
  onmessageerror: ((this: GenericMessagePort<TIn, TOut>, ev: MessageEvent<TIn>) => void) | null;
  postMessage(message: TOut, transfer: Transferable[]): void;
  postMessage(message: TOut, options?: PostMessageOptions): void;
}

type GenericChannel<TIn, TOut> = {
  port1: GenericMessagePort<TIn, TOut>;
  port2: GenericMessagePort<TOut, TIn>;
};

const enum EventType {
  OPEN_BUS = 0,
  OPEN_WEBRTC = 1,
  OPEN_DATA_CHANNEL = 2,
  DATA = 3,
  ERROR = 4,
  CLOSE = 5,
  ICE_CANDIDATE = 6,
  CONNECTION_STATE_CHANGE = 7,
  ICE_GATHERING_STATE_CHANGE = 8,
  SIGNALING_STATE_CHANGE = 9,
  CREATE_OFFER = 10,
  CREATE_ANSWER = 11,
  ADD_ICE_CANDIDATE = 12,
  SET_LOCAL_DESCRIPTION = 13,
  SET_REMOTE_DESCRIPTION = 14,
  DATA_CHANNEL = 15,
  DATA_CHANNEL_DATA = 16,
  DATA_CHANNEL_OPEN = 17,
  DATA_CHANNEL_CLOSE = 18,
  CREATE_DATA_CHANNEL = 19,
  OPEN_WORKER = 20,
}

type WorkerEvent =
  | {
      type: EventType.OPEN_WEBRTC;
      port: WebRTCMessagePort;
    }
  | {
      type: EventType.OPEN_DATA_CHANNEL;
      port: GenericMessagePort<void, WebRTCDataChannelMessagePort>;
      id: number;
    }
  | {
      type: EventType.OPEN_WORKER;
      port: WorkerMessagePort;
      service: string;
    }
  | {
      type: EventType.OPEN_BUS;
      port: BusMessagePort;
      label: string;
    };

type WebRTCWorkerEvent =
  | {
      type: EventType.CREATE_OFFER;
    }
  | {
      type: EventType.CREATE_ANSWER;
    }
  | {
      type: EventType.CREATE_DATA_CHANNEL;
      label: string;
    }
  | {
      type: EventType.ADD_ICE_CANDIDATE;
      candidate: string;
    }
  | {
      type: EventType.SET_LOCAL_DESCRIPTION;
      description: string;
    }
  | {
      type: EventType.SET_REMOTE_DESCRIPTION;
      description: string;
    }
  | {
      type: EventType.CLOSE;
    };

type WebRTCWindowEvent =
  | {
      type: EventType.ICE_CANDIDATE;
      candidate: string;
    }
  | {
      type: EventType.CONNECTION_STATE_CHANGE;
      state: string;
    }
  | {
      type: EventType.ICE_GATHERING_STATE_CHANGE;
      state: string;
    }
  | {
      type: EventType.SIGNALING_STATE_CHANGE;
      state: string;
    }
  | {
      type: EventType.CREATE_OFFER;
      error?: string;
      description?: string;
    }
  | {
      type: EventType.CREATE_ANSWER;
      error?: string;
      description?: string;
    }
  | {
      type: EventType.DATA_CHANNEL;
      id: number;
      label: string;
    };

type WebRTCMessagePort = GenericMessagePort<WebRTCWorkerEvent, WebRTCWindowEvent>;

type WebRTCDataChannelWorkerEvent =
  | {
      type: EventType.DATA_CHANNEL_DATA;
      data: ArrayBuffer;
    }
  | {
      type: EventType.DATA_CHANNEL_CLOSE;
    };

type WebRTCDataChannelWindowEvent =
  | {
      type: EventType.DATA_CHANNEL_DATA;
      data: ArrayBuffer;
      timestamp: number;
    }
  | {
      type: EventType.DATA_CHANNEL_OPEN;
    }
  | {
      type: EventType.DATA_CHANNEL_CLOSE;
    };

type WebRTCDataChannelMessagePort = GenericMessagePort<
  WebRTCDataChannelWindowEvent,
  WebRTCDataChannelWorkerEvent
>;

type WorkerWorkerEvent = {
  type: EventType.CLOSE;
};

type WorkerWindowEvent = {
  port: BusMessagePort;
};

type WorkerMessagePort = GenericMessagePort<WorkerWorkerEvent, WorkerWindowEvent>;

type BusEvent =
  | {
      type: EventType.DATA;
      data: ArrayBuffer;
      timestamp: number;
    }
  | {
      type: EventType.CLOSE;
    };

type BusMessagePort = GenericMessagePort<BusEvent, BusEvent>;

interface PortState {
  port: WebRTCDataChannelMessagePort;
  open: () => void;
}

export class WindowBridge extends EventEmitter {
  private nextDataChannelPortId = 0;
  private dataChannelPorts = new Map<number, PortState>();

  constructor(private workerConstructor: new () => Worker) {
    super();
    this.createWorker("default");
  }

  private createWorker(service: string, ...args: any[]): Worker {
    const worker = new this.workerConstructor();
    worker.onmessage = ({ data }: MessageEvent<WorkerEvent>) => {
      switch (data.type) {
        case EventType.OPEN_WEBRTC:
          this.openWebRTC(data.port);
          break;
        case EventType.OPEN_DATA_CHANNEL:
          this.openDataChannel(data.port, data.id);
          break;
        case EventType.OPEN_WORKER:
          this.openWorker(data.port, data.service);
          break;
        case EventType.OPEN_BUS:
          this.openBus(data.port, data.label);
          break;
      }
    };
    worker.postMessage({
      service,
      baseURI: location.origin,
      args,
    });
    return worker;
  }

  private openWebRTC(port: WebRTCMessagePort) {
    // TODO: specify ice servers...
    const peerConnection = new RTCPeerConnection({
      iceServers: [
        {
          urls: ["stun:stun.l.google.com:19302"],
        },
      ],
    });

    const dataChannels: { id: number; port: MessagePort; dataChannel: RTCDataChannel }[] = [];
    const handleDataChannel = (dataChannel: RTCDataChannel) => {
      let open: (...arg: any[]) => void;
      const ready = new Promise((resolve) => (open = resolve));

      const { port1, port2 } = new MessageChannel() as GenericChannel<
        WebRTCDataChannelWorkerEvent,
        WebRTCDataChannelWindowEvent
      >;
      const portId = this.nextDataChannelPortId++;
      this.dataChannelPorts.set(portId, { port: port2, open });

      port.postMessage({
        type: EventType.DATA_CHANNEL,
        id: portId,
        label: dataChannel.label,
      });

      dataChannel.binaryType = "arraybuffer";

      dataChannel.onmessage = (e: MessageEvent<ArrayBuffer>) =>
        ready.then(() =>
          port1.postMessage(
            {
              type: EventType.DATA_CHANNEL_DATA,
              timestamp: Date.now(),
              data: e.data,
            },
            [e.data]
          )
        );

      dataChannel.onopen = () =>
        ready.then(() =>
          port1.postMessage({
            type: EventType.DATA_CHANNEL_OPEN,
          })
        );

      dataChannel.onclose = onclose;

      port1.onmessage = ({ data }: MessageEvent<WebRTCDataChannelWorkerEvent>) =>
        ready.then(() => {
          // console.log("window data channel event", data);
          switch (data.type) {
            case EventType.DATA_CHANNEL_DATA:
              dataChannel.send(data.data);
              break;
            case EventType.DATA_CHANNEL_CLOSE:
              dataChannel.close();
              break;
          }
        });

      dataChannels.push({
        port: port1,
        id: portId,
        dataChannel,
      });
    };

    peerConnection.ondatachannel = ({ channel }: RTCDataChannelEvent) => handleDataChannel(channel);

    const onclose = () => {
      dataChannels.forEach(({ id, port, dataChannel }) => {
        port.postMessage({ type: EventType.DATA_CHANNEL_CLOSE });
        this.dataChannelPorts.delete(id);

        dataChannel.onmessage = null;
        dataChannel.onopen = null;
        dataChannel.onclose = null;
        dataChannel.close();

        port.onmessage = null;
        port.close();
      });

      peerConnection.ondatachannel = null;
      peerConnection.onicecandidate = null;
      peerConnection.onconnectionstatechange = null;
      peerConnection.onicegatheringstatechange = null;
      peerConnection.onsignalingstatechange = null;
      peerConnection.close();

      port.onmessage = null;
      port.close();
    };

    peerConnection.onicecandidate = (e: RTCPeerConnectionIceEvent) =>
      port.postMessage({
        type: EventType.ICE_CANDIDATE,
        candidate: JSON.stringify(e.candidate),
      });

    peerConnection.onconnectionstatechange = () => {
      const state = peerConnection.iceConnectionState;
      port.postMessage({
        type: EventType.CONNECTION_STATE_CHANGE,
        state,
      });

      if (state === "failed" || state === "disconnected" || state === "closed") {
        onclose();
      }
    };

    peerConnection.onicegatheringstatechange = () =>
      port.postMessage({
        type: EventType.ICE_GATHERING_STATE_CHANGE,
        state: peerConnection.iceGatheringState,
      });

    peerConnection.onsignalingstatechange = () =>
      port.postMessage({
        type: EventType.SIGNALING_STATE_CHANGE,
        state: peerConnection.signalingState,
      });

    port.onmessage = ({ data }: MessageEvent<WebRTCWorkerEvent>) => {
      // console.log("window event", eventTypeNames[data.type], data);
      switch (data.type) {
        case EventType.CREATE_OFFER:
          peerConnection
            .createOffer()
            .then((description) =>
              port.postMessage({
                type: EventType.CREATE_OFFER,
                description: JSON.stringify(description),
              })
            )
            .catch((error) => {
              port.postMessage({
                type: EventType.CREATE_OFFER,
                error: String(error),
              });
              onclose();
            });
          break;
        case EventType.CREATE_ANSWER:
          peerConnection
            .createAnswer()
            .then((description) =>
              port.postMessage({
                type: EventType.CREATE_ANSWER,
                description: JSON.stringify(description),
              })
            )
            .catch((error) => {
              port.postMessage({
                type: EventType.CREATE_ANSWER,
                error: String(error),
              });
              onclose();
            });
          break;
        case EventType.CREATE_DATA_CHANNEL:
          handleDataChannel(peerConnection.createDataChannel(data.label));
          break;
        case EventType.ADD_ICE_CANDIDATE:
          void peerConnection.addIceCandidate(new RTCIceCandidate(JSON.parse(data.candidate)));
          break;
        case EventType.SET_LOCAL_DESCRIPTION:
          void peerConnection.setLocalDescription(
            new RTCSessionDescription(JSON.parse(data.description))
          );
          break;
        case EventType.SET_REMOTE_DESCRIPTION:
          void peerConnection.setRemoteDescription(
            new RTCSessionDescription(JSON.parse(data.description))
          );
          break;
        case EventType.CLOSE:
          peerConnection.close();
          onclose();
          break;
      }
    };
  }

  private openDataChannel(
    port: GenericMessagePort<void, WebRTCDataChannelMessagePort>,
    id: number
  ) {
    const portState = this.dataChannelPorts.get(id);
    if (portState === undefined) {
      port.postMessage(undefined);
      return;
    }
    this.dataChannelPorts.delete(id);

    port.postMessage(portState.port, [portState.port]);
    setTimeout(() => portState.open(), 100);
  }

  private openWorker(port: WorkerMessagePort, service: string) {
    const worker = this.createWorker(service);
    this.once(`busport:${service}`, (p: MessagePort) => port.postMessage({ port: p }, [p]));

    port.onmessage = ({ data }: MessageEvent<BusEvent>) => {
      switch (data.type) {
        case EventType.CLOSE:
          worker.terminate();
          break;
      }
    };
  }

  private openBus(port: MessagePort, label: string) {
    this.emit(`busport:${label}`, port);
    this.emit(`busopen:${label}`, new Bus(port, label));
  }
}

export class Bus extends EventEmitter {
  constructor(public port: BusMessagePort, public label: string) {
    super();

    port.onmessage = ({ data }: MessageEvent<BusEvent>) => {
      switch (data.type) {
        case EventType.DATA:
          this.emit("data", data.data);
          break;
        case EventType.CLOSE:
          this.emit("close");
          break;
      }
    };
  }

  public write(data: ArrayBuffer | Uint8Array): void {
    const buffer = data instanceof ArrayBuffer ? data : data.buffer;

    this.port.postMessage(
      {
        type: EventType.DATA,
        data: buffer,
        timestamp: Date.now(),
      },
      [buffer]
    );
  }

  public close(): void {
    this.port.postMessage({ type: EventType.CLOSE });
    this.port.close();
  }
}

export interface WebSocketGoProxy {
  ondata(data: Uint8Array, n: number, timestamp: number): void;
  onopen(): void;
  onclose(): void;
  onerror(message: string): void;
}

interface WebSocketProxy {
  write: (data: Uint8Array) => void;
  close: () => void;
}

export interface WebRTCGoProxy {
  onicecandidate(candidate: string): void;
  onconnectionstatechange(state: string): void;
  onicegatheringstatechange(state: string): void;
  onsignalingstatechange(state: string): void;
  oncreateoffer(error: string, description: string): void;
  oncreateanswer(error: string, description: string): void;
  ondatachannel(id: number, label: string): void;
}

interface WebRTCProxy {
  createOffer: () => void;
  createAnswer: () => void;
  createDataChannel: (label: string) => void;
  addIceCandidate: (candidate: string) => void;
  setLocalDescription: (description: string) => void;
  setRemoteDescription: (description: string) => void;
  close: () => void;
}

export interface DataChannelGoProxy {
  onerror(message: string): void;
  ondata(data: Uint8Array, n: number, timestamp: number): void;
  onclose(): void;
  onopen(): void;
}

interface DataChannelProxy {
  write: ({ buffer }: Uint8Array) => void;
  close: () => void;
}

export interface ServiceGoProxy {
  openBus(any): DataChannelGoProxy;
}

interface KVStoreProxy {
  put: (key: string, value: Uint8Array, done: (error: string | null) => void) => void;
  delete: (key: string, done: (error: string | null) => void) => void;
  get: (key: string, done: (error: string | null, value?: Uint8Array) => void) => void;
  scanPrefix: (prefix: string, done: (error: string | null, value?: any[]) => void) => void;
  rollback: (done: (error: string | null) => void) => void;
  commit: (done: (error: string | null) => void) => void;
}

interface DataChannelProxy {
  write: ({ buffer }: Uint8Array) => void;
  close: () => void;
}

export class WorkerBridge {
  public openWebSocket(uri: string, proxy: WebSocketGoProxy): WebSocketProxy {
    const ws = new WebSocket(uri);

    const onclose = () => {
      ws.onopen = null;
      ws.onclose = null;
      ws.onerror = null;
      ws.onmessage = null;
    };

    ws.binaryType = "arraybuffer";
    ws.onopen = () => proxy.onopen();
    ws.onclose = () => {
      onclose();
      proxy.onclose();
    };
    ws.onerror = (e: ErrorEvent) => proxy.onerror(String(e.message || "unknown websocket error"));
    ws.onmessage = ({ data }: MessageEvent<ArrayBuffer>) =>
      proxy.ondata(new Uint8Array(data), data.byteLength, Date.now());

    return {
      write: (data: Uint8Array) => ws.send(data),
      close: () => {
        onclose();
        ws.close();
      },
    };
  }

  public openWebRTC(proxy: WebRTCGoProxy): WebRTCProxy {
    const { port1, port2 } = new MessageChannel() as GenericChannel<
      WebRTCWindowEvent,
      WebRTCWorkerEvent
    >;
    port1.onmessage = ({ data }) => {
      // console.log("worker event", eventTypeNames[data.type], data);
      switch (data.type) {
        case EventType.ICE_CANDIDATE:
          proxy.onicecandidate(data.candidate);
          break;
        case EventType.CONNECTION_STATE_CHANGE:
          proxy.onconnectionstatechange(data.state);
          break;
        case EventType.ICE_GATHERING_STATE_CHANGE:
          proxy.onicegatheringstatechange(data.state);
          break;
        case EventType.SIGNALING_STATE_CHANGE:
          proxy.onsignalingstatechange(data.state);
          break;
        case EventType.CREATE_OFFER:
          proxy.oncreateoffer(data.error, data.description);
          break;
        case EventType.CREATE_ANSWER:
          proxy.oncreateanswer(data.error, data.description);
          break;
        case EventType.DATA_CHANNEL:
          proxy.ondatachannel(data.id, data.label);
          break;
      }
    };

    postMessage<WorkerEvent>(
      {
        type: EventType.OPEN_WEBRTC,
        port: port2,
      },
      [port2]
    );

    return {
      createOffer: () =>
        port1.postMessage({
          type: EventType.CREATE_OFFER,
        }),
      createAnswer: () =>
        port1.postMessage({
          type: EventType.CREATE_ANSWER,
        }),
      createDataChannel: (label: string) =>
        port1.postMessage({
          type: EventType.CREATE_DATA_CHANNEL,
          label,
        }),
      addIceCandidate: (candidate: string) =>
        port1.postMessage({
          type: EventType.ADD_ICE_CANDIDATE,
          candidate,
        }),
      setLocalDescription: (description: string) =>
        port1.postMessage({
          type: EventType.SET_LOCAL_DESCRIPTION,
          description,
        }),
      setRemoteDescription: (description: string) =>
        port1.postMessage({
          type: EventType.SET_REMOTE_DESCRIPTION,
          description,
        }),
      close: () => {
        port1.postMessage({
          type: EventType.CLOSE,
        });
        port1.close();
      },
    };
  }

  // openDataChannel opens a data channel created by a call to openWebRTC. This
  // allows multiple workers to share an RTCPeerConnection.
  public openDataChannel(id: number, proxy: DataChannelGoProxy): DataChannelProxy {
    const { port1, port2 } = new MessageChannel() as GenericChannel<
      WebRTCDataChannelMessagePort,
      void
    >;

    const ready = new Promise<WebRTCDataChannelMessagePort>((resolve, reject) => {
      port1.onmessage = ({ data: port }) => {
        port1.close();

        if (port === undefined) {
          const err = new Error("data channel invalid or in use");
          proxy.onerror(err.message);
          reject(err);
          return;
        }
        resolve(port);

        port.onmessage = ({ data }) => {
          // console.log("worker data channel event", data);
          switch (data.type) {
            case EventType.DATA_CHANNEL_DATA:
              proxy.ondata(new Uint8Array(data.data), data.data.byteLength, data.timestamp);
              break;
            case EventType.DATA_CHANNEL_OPEN:
              proxy.onopen();
              break;
            case EventType.DATA_CHANNEL_CLOSE:
              proxy.onclose();
              port.close();
              break;
          }
        };
      };
    });

    postMessage<WorkerEvent>(
      {
        type: EventType.OPEN_DATA_CHANNEL,
        port: port2,
        id,
      },
      [port2]
    );

    return {
      write: ({ buffer }: Uint8Array) => {
        void ready.then((port) =>
          port.postMessage(
            {
              type: EventType.DATA_CHANNEL_DATA,
              data: buffer,
            },
            [buffer]
          )
        );
      },
      close: () => {
        void ready.then((port) => {
          port.postMessage({
            type: EventType.DATA_CHANNEL_CLOSE,
          });
          port.close();
        });
      },
    };
  }

  public openWorker(service: string, proxy: ServiceGoProxy): void {
    const { port1, port2 } = new MessageChannel() as GenericChannel<
      WorkerWindowEvent,
      WorkerWorkerEvent
    >;

    port1.onmessage = ({ data: { port } }) => {
      const bus = proxy.openBus({
        write: ({ buffer }: Uint8Array) =>
          port.postMessage(
            {
              type: EventType.DATA,
              data: buffer,
              timestamp: Date.now(),
            },
            [buffer]
          ),
        close: () => {
          port.postMessage({ type: EventType.CLOSE });
          port1.postMessage({ type: EventType.CLOSE });
        },
      });

      port.onmessage = ({ data }) => {
        switch (data.type) {
          case EventType.DATA:
            bus.ondata(new Uint8Array(data.data), data.data.byteLength, data.timestamp);
            break;
          case EventType.CLOSE:
            bus.onclose();
            break;
        }
      };
    };

    postMessage<WorkerEvent>(
      {
        type: EventType.OPEN_WORKER,
        port: port2,
        service,
      },
      [port2]
    );
  }

  public openBus(label: string, proxy: DataChannelGoProxy): DataChannelProxy {
    const { port1, port2 } = new MessageChannel() as GenericChannel<BusEvent, BusEvent>;

    port1.onmessage = ({ data }: MessageEvent<BusEvent>) => {
      switch (data.type) {
        case EventType.DATA:
          proxy.ondata(new Uint8Array(data.data), data.data.byteLength, data.timestamp);
          break;
        case EventType.CLOSE:
          proxy.onclose();
          break;
      }
    };

    postMessage<WorkerEvent>(
      {
        type: EventType.OPEN_BUS,
        port: port2,
        label,
      },
      [port2]
    );

    return {
      write: ({ buffer }: Uint8Array) =>
        port1.postMessage(
          {
            type: EventType.DATA,
            data: buffer,
            timestamp: Date.now(),
          },
          [buffer]
        ),
      close: () => port1.postMessage({ type: EventType.CLOSE }),
    };
  }

  public openKVStore(name: string, createTable: boolean, readOnly: boolean): KVStoreProxy {
    const openReq = indexedDB.open(name);
    openReq.onupgradeneeded = () => {
      if (!createTable) {
        openReq.transaction.abort();
        return;
      }
      openReq.result.createObjectStore("data");
    };

    const txReady: Promise<IDBTransaction> = new Promise((resolve, reject) => {
      const mode = readOnly ? "readonly" : "readwrite";
      openReq.onsuccess = () => resolve(openReq.result.transaction(["data"], mode));
      openReq.onerror = reject;
      openReq.onblocked = reject;
    });

    const transact = <T>(
      operator: (db: IDBObjectStore) => IDBRequest<T>,
      done: (error: string | null, value?: T) => void
    ) => {
      txReady
        .then((tx: IDBTransaction) => {
          const req = operator(tx.objectStore("data"));
          req.onsuccess = () => done(null, req.result);
          req.onerror = (e) => done(String(e));
        })
        .catch((e: Error) => done(String(e.message || "unknown storage error")));
    };

    return {
      put: (key: string, value: Uint8Array, done: (error: string | null) => void) => {
        transact((s) => s.put(value, key), done);
      },
      delete: (key: string, done: (error: string | null) => void) => {
        transact((s) => s.delete(key), done);
      },
      get: (key: string, done: (error: string | null, value?: Uint8Array) => void) => {
        transact((s) => s.get(key), done);
      },
      scanPrefix: (prefix: string, done: (error: string | null, value?: any[]) => void) => {
        const range = IDBKeyRange.bound(prefix, prefix + "\uffff", false, false);
        transact((s) => s.getAll(range), done);
      },
      rollback: (done: (error: string | null) => void) => {
        txReady
          .then((tx: IDBTransaction) => {
            tx.onabort = () => done(null);
            tx.onerror = (e) => done(String(e));
            tx.abort();
          })
          .catch((e) => done(String(e)));
      },
      commit: (done: (error: string | null) => void) => {
        txReady
          .then((tx: IDBTransaction) => {
            tx.oncomplete = () => done(null);
            tx.onerror = (e) => done(String(e));
            // eslint-disable-next-line
            (tx as any).commit?.();
          })
          .catch((e) => done(String(e)));
      },
    };
  }

  public deleteKVStore(name: string, done: (err: string | null) => void): void {
    const res = indexedDB.deleteDatabase(name);
    res.onsuccess = () => done(null);
    res.onerror = (e) => done(String(e));
  }

  public syncLogs(data: Uint8Array): void {
    const { level, caller, msg, ...args } = JSON.parse(new TextDecoder().decode(data)) as {
      [key: string]: string;
    };
    // eslint-disable-next-line
    console.log(
      level.toUpperCase().padEnd(10),
      caller.padEnd(30),
      msg,
      Object.keys(args).length === 0 ? "" : args
    );
  }
}
