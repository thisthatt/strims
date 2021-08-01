import React, { useContext, useEffect, useRef } from "react";
import { MdLoop } from "react-icons/md";
import useFullscreen from "use-fullscreen";

import useIdleTimeout from "../../hooks/useIdleTimeout";
import useMediaSource, { MediaSourceProps } from "../../hooks/useMediaSource";
import useVideo from "../../hooks/useVideo";
import { MainLayoutContext } from "../MainLayout";
import LogoButton from "./LogoButton";
import VideoControls from "./VideoControls";

type SwarmPlayerProps = Pick<MediaSourceProps, "networkKey" | "swarmUri" | "mimeType"> & {
  disableControls?: boolean;
  defaultAspectRatio?: string | number;
};

const SwarmPlayer: React.FC<SwarmPlayerProps> = ({
  networkKey,
  swarmUri,
  mimeType,
  disableControls = false,
  defaultAspectRatio = "16/9",
}) => {
  const rootRef = useRef();
  const [controlsHidden, renewControlsTimeout, clearControlsTimeout] = useIdleTimeout();
  const [isFullscreen, toggleFullscreen] = useFullscreen();
  const { theaterMode, toggleTheaterMode } = useContext(MainLayoutContext);
  const [videoState, videoProps, videoControls] = useVideo();
  const mediaSource = useMediaSource({ networkKey, swarmUri, mimeType, videoRef: videoProps.ref });

  useEffect(() => {
    console.log(">>>", videoState.error);
  }, [videoState.error]);

  useEffect(() => {
    if (!videoProps.ref.current || !mediaSource) {
      return;
    }

    const src = URL.createObjectURL(mediaSource);
    videoControls.setSrc(src);
    return () => URL.revokeObjectURL(src);
  }, [videoProps.ref, mediaSource]);

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

  const handleToggleFullscreen = () => toggleFullscreen(rootRef.current);

  useEffect(renewControlsTimeout, [videoState.volume]);

  let aspectRatio = defaultAspectRatio;
  if (videoState.videoWidth && videoState.videoHeight) {
    aspectRatio = `${videoState.videoWidth}/${videoState.videoHeight}`;
  }

  return (
    <div
      className="video_player"
      onMouseMove={renewControlsTimeout}
      onMouseLeave={clearControlsTimeout}
      onDoubleClick={handleToggleFullscreen}
      ref={rootRef}
      style={{ aspectRatio }}
    >
      <video onClick={(e) => e.preventDefault()} className="video_player__video" {...videoProps} />
      {waitingSpinner}
      <VideoControls
        videoState={videoState}
        videoControls={videoControls}
        visible={!controlsHidden && !disableControls}
        fullscreen={isFullscreen}
        toggleFullscreen={handleToggleFullscreen}
        theaterMode={theaterMode}
        toggleTheaterMode={toggleTheaterMode}
      />
    </div>
  );
};

export default SwarmPlayer;
