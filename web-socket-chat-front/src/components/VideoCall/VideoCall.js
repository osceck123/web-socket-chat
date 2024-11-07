// src/components/VideoCall.js
import React, { useRef, useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { initiateVideoCall } from "../../services/api"
import { addPeer, clearPeers, setLocalStream } from "../redux/videoSlice";

function VideoCall() {
  const dispatch = useDispatch();
  const localVideoRef = useRef(null);
  const peers = useSelector((state) => state.video.peers);

  useEffect(() => {
    const startCall = async () => {
      try {
        await initiateVideoCall();

        // Acceso a cámara y micrófono
        const stream = await navigator.mediaDevices.getUserMedia({
          video: true,
          audio: true,
        });

        localVideoRef.current.srcObject = stream;
        dispatch(setLocalStream(stream));

        // Configuración de WebRTC y gestión de peers
        peers.forEach((peer) => {
          const peerConnection = new RTCPeerConnection();
          stream.getTracks().forEach((track) => peerConnection.addTrack(track, stream));
          
          peerConnection.ontrack = (event) => {
            dispatch(addPeer({ id: peer.id, stream: event.streams[0] }));
          };
        });
      } catch (error) {
        console.error("Failed to start video call:", error);
      }
    };

    startCall();
    return () => {
      dispatch(clearPeers());
    };
  }, [dispatch]);

  return (
    <div>
      <h2>Video Call</h2>
      <video ref={localVideoRef} autoPlay playsInline muted />
      <div className="remote-videos">
        {peers.map((peer) => (
          <video key={peer.id} srcObject={peer.stream} autoPlay playsInline />
        ))}
      </div>
    </div>
  );
}

export default VideoCall;
