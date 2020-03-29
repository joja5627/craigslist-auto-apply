
import React, { createContext, useState } from 'react'

export const ListingContext = createContext();

export const ListingContextProvider = ({children}) =>{
  const [listings, setListings] = useState([]);

  function addListing(listing){
      console.log(listing)
    setListings([...listings, listing]);
    console.log(listings)
  }
  function listingCount(){
   return listings.length
  }

  return <ListingContext.Provider value={{addListing,listingCount,listings}}>
        {children}, 
  </ListingContext.Provider>
}
