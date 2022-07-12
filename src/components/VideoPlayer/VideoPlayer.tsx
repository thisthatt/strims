// Copyright 2022 Strims contributors
// SPDX-License-Identifier: AGPL-3.0-only

import "./VideoPlayer.scss";

import React, { useEffect, useRef } from "react";
import { MdLoop } from "react-icons/md";
import useFullscreen from "use-fullscreen";

import { useLayout } from "../../contexts/Layout";
import useIdleTimeout from "../../hooks/useIdleTimeout";
import useMediaRelay from "../../hooks/useMediaRelay";
import useMediaSource, { MediaSourceProps } from "../../hooks/useMediaSource";
import useReady from "../../hooks/useReady";
import useVideo from "../../hooks/useVideo";
import { DEVICE_TYPE, DeviceType, OS } from "../../lib/userAgent";
import LogoButton from "./LogoButton";
import VideoControls from "./VideoControls";

interface SwarmPlayerProps extends Pick<MediaSourceProps, "networkKey" | "swarmUri" | "mimeType"> {
  defaultControlsVisible?: boolean;
  disableControls?: boolean;
  defaultAspectRatio?: string | number;
  handleClose: () => void;
}

const SwarmPlayer: React.FC<SwarmPlayerProps> = ({
  networkKey,
  swarmUri,
  mimeType,
  defaultControlsVisible,
  disableControls = false,
  defaultAspectRatio = "16/9",
  handleClose,
}) => {
  const rootRef = useRef();
  const [controlsHidden, renewControlsTimeout, clearControlsTimeout] = useIdleTimeout();
  const [isFullscreen, toggleFullscreen] = useFullscreen();
  const { theaterMode, toggleTheaterMode } = useLayout();
  const videoRef = useRef<HTMLVideoElement>();
  const [videoState, videoProps, videoControls] = useVideo(videoRef);

  if (window.MediaSource) {
    const mediaSource = useMediaSource({ networkKey, swarmUri, mimeType, videoRef });

    useReady(() => {
      const src = URL.createObjectURL(mediaSource);
      videoControls.setSrc(src);
      return () => URL.revokeObjectURL(src);
    }, [mediaSource]);
  } else {
    const src = useMediaRelay({ networkKey, swarmUri, mimeType, videoRef });
    useReady(() => videoControls.setSrc(src), [src]);

    // const client = useClient();
    // useEffect(() => {
    //   void client.hlsEgress
    //     .openStream({ swarmUri, networkKeys: [Base64.toUint8Array(networkKey)] })
    //     .then(({ playlistUrl }) => {
    //       videoRef.current.src = playlistUrl;
    //     });
    // }, [networkKey, swarmUri]);
  }

  useEffect(() => {
    console.log(">>>", videoState.error);
  }, [videoState.error]);

  const waitingSpinner =
    videoState.waiting && videoState.loaded ? (
      <div className="video_player__waiting_spinner">
        <MdLoop />
      </div>
    ) : (
      <LogoButton
        visible={!videoState.playing && !videoState.paused}
        onClick={videoControls.play}
        flicker={videoState.ended && !videoState.loaded}
        spin={videoState.waiting && videoState.loaded}
        disabled={videoState.waiting || !videoState.loaded}
        blur={true}
      />
    );

  const handleToggleFullscreen = () => {
    if (typeof videoRef.current.webkitSetPresentationMode === "function") {
      videoRef.current.webkitSetPresentationMode("fullscreen");
    } else {
      void toggleFullscreen(rootRef.current);
    }
  };

  useEffect(renewControlsTimeout, [videoState.volume]);

  let aspectRatio = defaultAspectRatio;
  if (videoState.videoWidth && videoState.videoHeight) {
    aspectRatio = `${videoState.videoWidth}/${videoState.videoHeight}`;
  }

  const controls =
    DEVICE_TYPE === DeviceType.Portable
      ? {
          showPlayButton: true,
          showVolume: true,
          showFullscreenButton: true,
          showCloseButton: true,
        }
      : {
          showPlayButton: true,
          showVolume: true,
          showPiPButton: true,
          showTheaterButton: true,
          showFullscreenButton: true,
        };

  return (
    <div
      className="video_player"
      onMouseMove={renewControlsTimeout}
      onMouseLeave={clearControlsTimeout}
      onDoubleClick={handleToggleFullscreen}
      ref={rootRef}
      style={{ aspectRatio }}
    >
      <video
        onClick={(e) => e.preventDefault()}
        className="video_player__video"
        autoPlay
        playsInline
        {...videoProps}
      />
      {waitingSpinner}
      <VideoControls
        videoState={videoState}
        videoControls={videoControls}
        visible={(!controlsHidden || defaultControlsVisible) && !disableControls}
        fullscreen={isFullscreen}
        toggleFullscreen={handleToggleFullscreen}
        theaterMode={theaterMode}
        toggleTheaterMode={toggleTheaterMode}
        close={handleClose}
        {...controls}
      />
    </div>
  );
};

export default SwarmPlayer;
