import React, { createContext, useEffect, useState } from 'react'
import useSocket from 'use-socket.io-client';

export const WebSocketContext = createContext();

export const WebSocketContextProvider = ({children}) =>{
  
  const sockerIOProps = {
    autoConnect: false,
    path: "/scrape_test",
    hostname: "localhost",
    secure: false,
    port: "8080"
  }
  const [socket] = useSocket(sockerIOProps);

  const [connected, setConnected] = useState(false)
  
  function connect(){
      console.log(socket)
      console.log("connect")
    setConnected(true)
    try {
        console.log("try connect")
        socket.connect()
       }catch(error){
        console.log(error)
       }
  }
  function disconnect(){
    console.log("disconnect")
    setConnected(false)
    try {
        socket.disconnect()
       }catch(error){
        console.log(error)
       }
   }
  
       


  return <WebSocketContext.Provider value={{socket,connected, connect,disconnect,connect}}>
        {children}, 
  </WebSocketContext.Provider>
}




// import { createContext, useEffect, useState,useRef } from 'react';
// import _ from 'lodash';







// export const useWebSocket = (props) => {

// //   const [websocketUrl, setWebsocketUrl] = useState(props.url);
// //   const [websocket, setWebsocket] = useState(WebSocket);
//   const [connected, setConnected] = useState(false);
//   const [gamepadConnected, setGamepadConnected] = useState(false);

//   const ws = useRef(null);

//   useEffect(() => {
//       if(props.connected === true){
//         ws.current = new WebSocket(props.url).readyState('CLOSED');
//         ws.onopen = event => this.onOpen(event);
//         ws.onmessage = event => this.onMessage(event);
//         ws.onclose = event => this.onClose(event);
//         ws.onerror = event => this.onError(event);
//         setConnected(props.connected)
//       }else {
//           console.log("closing")
//           ws.current.close();
//           setConnected(props.connected)
//       }
//   }, [props.connected]);




//   const onMessage = e => {
//     let event = JSON.parse(e.data);
//     if (event.messageType === 'state') {
//     } else if (event.messageType === 'listingPercentComplete') {
//       this.setState({ clListingScrapePercentage: Math.round(event.payload) });
//     } else if (event.messageType === 'listings') {
//       let newListing = event.payload;
//       props.handleMessage(newListing)
//     //   if (!_.find(this.state.listings, { url: newListing.url })) {
//     //     if (this.state.emailAll) {
//     //       this.sendEmail(newListing);
//     //     } else {
//     //       let currentLinkCount = this.state.linkCount;
//     //       this.setState({
//     //         listing: newListing,
//     //         linkCount: (currentLinkCount += 1)
//     //       });
//     //     }
//     //   }
//     }
//   };


  
//   // useEffect(() => {
//   //     if(!websocket) {
//   //         return;
//   //     }
//   //     const interval = setInterval(() => {
//   //         if(!websocket || websocket.readyState !== 1) {
//   //             setWebsocketConnected(false);
//   //             return;
//   //         }
//   //         setWebsocketConnected(true);
//   //         websocket.send(JSON.stringify(values));
//   //     }, 30);
//   //     return () => clearInterval(interval);
//   // }, [websocket]);

// }
