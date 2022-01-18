import React, { ComponentType } from "react";
import { BrowserRouter, MemoryRouter } from "react-router-dom";

import { APIDialer } from "../contexts/Session";
import { IS_PWA } from "../lib/userAgent";
import Provider from "../root/Provider";
import RootRouter from "../root/Router";

const Router: ComponentType = IS_PWA ? MemoryRouter : BrowserRouter;

interface AppProps {
  apiDialer: APIDialer;
}

const App: React.FC<AppProps> = ({ apiDialer }) => (
  <Router>
    <Provider apiDialer={apiDialer}>
      <RootRouter />
    </Provider>
  </Router>
);

export default App;
