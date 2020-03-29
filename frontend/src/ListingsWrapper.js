import React, { useEffect, useState ,useContext} from 'react';
import _ from 'lodash';
import { CSSTransition, TransitionGroup } from 'react-transition-group';
import {Listings} from "./Listings";
import {WebSocketContext} from './WebSocketContextProvider';
import {ListingContext} from './ListingContextProvider';


export function ListingsWrapper() {
  
  const {listings,addListing,listingCount} = useContext(ListingContext);
  const {socket} = useContext(WebSocketContext);


  useEffect(() => {
    socket.on('listings',message => {
        console.log(message)
        // listingsHandler(message)
      });
   })
//    function listingsHandler(event){
//     return event
//    }
//    function listingPercentCompleteHandler(event){
//     return event
//    }
 



  return (
    <div>
      {listings.map((listing, index) => (
        <CSSTransition
          key={`${index}-listing-key`}
          timeout={500}
          className="move"
        >
        <div className="padding40 ">
          <ul>
              <li>{listing}</li>
          </ul>
        </div>
        </CSSTransition>
      ))}
    </div>
  );
}
