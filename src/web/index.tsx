import "../styles/main.scss";
import "../lib/i18n";

import { Readable, Writable } from "stream";

import React from "react";
import ReactDOM from "react-dom";

import { APIDialer, ClientConstructor } from "../contexts/Session";
import { WindowBridge } from "../lib/bridge";
import { WSReadWriter } from "../lib/ws";
import App from "./App";
import Worker from "./svc.worker";

class WorkerConn {
  bridge: WindowBridge;
  bus: Promise<Readable & Writable>;

  constructor() {
    this.bridge = new WindowBridge(Worker);
    this.bus = new Promise<Readable & Writable>((resolve) => {
      this.bridge.once("busopen:default", (b: Readable & Writable) => resolve(b));
    });
  }

  async client<T>(C: ClientConstructor<T>): Promise<T> {
    const b = await this.bus;
    return new C(b, b);
  }

  close() {
    this.bridge.close();
  }
}

class WSConn extends WSReadWriter {
  client<T>(C: ClientConstructor<T>): Promise<T> {
    const ws = this as unknown as Readable & Writable;
    return Promise.resolve(new C(ws, ws));
  }
}

const apiDialer: APIDialer = {
  local: (): WorkerConn => new WorkerConn(),
  remote: (address: string): WSConn => new WSConn(address),
};

class Runner {
  root: HTMLDivElement;

  constructor() {
    this.root = document.createElement("div");
    this.root.setAttribute("id", "root");
    document.body.appendChild(this.root);
  }

  start() {
    ReactDOM.render(<App apiDialer={apiDialer} />, this.root);
  }

  stop() {
    ReactDOM.unmountComponentAtNode(this.root);
  }

  restart() {
    this.stop();
    this.start();
  }
}

const runner = new Runner();
runner.start();

if (import.meta.webpackHot) {
  import.meta.webpackHot.accept(
    ["../apis/client", "../lib/bridge", "../lib/ws", "./svc.worker"],
    () => runner.restart()
  );
}
