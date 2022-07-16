// Copyright 2022 Strims contributors
// SPDX-License-Identifier: AGPL-3.0-only

import React, { useContext, useEffect } from "react";
import { useLocation, useParams } from "react-router-dom";

import { useLayout } from "../contexts/Layout";
import { PlayerContext, PlayerMode } from "../contexts/Player";
import useQuery from "../hooks/useQuery";

interface PlayerQueryParams {
  swarmUri: string;
  mimeType: string;
}

const Player: React.FC = () => {
  const params = useParams<"networkKey">();
  const location = useLocation();
  const query = useQuery<PlayerQueryParams>(location.search);
  const { toggleShowVideo, toggleOverlayOpen } = useLayout();

  const { setMode, setSource, setPath } = useContext(PlayerContext);
  useEffect(() => {
    toggleOverlayOpen(true);
    toggleShowVideo(true);
    setMode(PlayerMode.LARGE);
    setSource({
      type: "swarm",
      networkKey: params.networkKey,
      swarmUri: query.swarmUri,
      mimeType: query.mimeType,
    });
    setPath(location.pathname + location.search);
    return () => {
      toggleOverlayOpen(false);
      setMode(PlayerMode.PIP);
    };
  }, [params.networkKey, query.swarmUri, query.mimeType]);

  return null;
};

export default Player;
