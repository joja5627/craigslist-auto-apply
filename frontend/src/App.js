import React, { useContext, Component, useState, useEffect } from 'react';
import './App.css';
import { Container, Grid } from 'semantic-ui-react';
import _ from 'lodash';
import { ProgressBarContainer } from './ProgressBar';
import { ListingsWrapper } from './ListingsWrapper';

import { WebSocketContext } from './WebSocketContextProvider';
import { Controller } from './Controller';

const App = () => {
  const { socket } = useContext(WebSocketContext);
  const { scrapePerc, setScrapePerc } = useState(0);
  useEffect(() => {
    socket.on('listingPercentComplete', message => {
      console.log(message);
      setScrapePerc(message);
    });
  });

  // connected: false,
  // responses: [],
  // eventTarget: null,
  // webSocket: null,
  // linkViewLimit: 300,
  // linkCount: 0,
  // emailAll: false,
  // selectedListing: null,
  // sendEmailStatus: null,
  // emailResponses: [],
  // sendingEmail: false,
  // emailQueue: new Queue(),
  // clListingScrapePercentage: 0

  //   const onInputChange = event => {
  //     this.setState({ eventTarget: event.target.value.toLowerCase() });
  //   };

  //   const filterFunction = link => {
  //     return link.toLowerCase().search(this.state.eventTarget) !== -1;
  //   };

  //   const filterList = listing => {
  //     if (this.state.eventTarget) {
  //       return listing.filter(link => this.filterFunction(link));
  //     } else {
  //       return listing;
  //     }
  //   };
  //   const onClickRemove = clickedListing => {
  //     let { listings } = this.state;
  //     _.remove(listings, function(listing) {
  //       return listing.id === clickedListing.id;
  //     });
  //     this.setState({ listings: listings });
  //   };
  //   const onClickRemoveEmailResponse = emailResponse => {
  //     let { responses } = this.state;
  //     _.remove(responses, function(response) {
  //       return emailResponse.id === response.id;
  //     });
  //     this.setState({ emailResponses: responses });
  //   };
  //   const addEmailQueue = clickedListing => {
  //     let { emailQueue } = this.state;
  //     emailQueue.push(clickedListing);
  //     this.setState({
  //       emailQueue: emailQueue
  //     });
  //     this.onClickRemove(clickedListing);
  //   };

  return (
    <div className="App">
      <ProgressBarContainer percentRange={scrapePerc} />
      <Container>
        <h1 className="ui header fontWeight100">Craigslist Web Scraper</h1>

        <Grid.Row>
          <Grid.Column width={4}></Grid.Column>
            <Controller></Controller>
          <Grid.Column width={12}>
            <ListingsWrapper />
          </Grid.Column>
          {/* <Grid.Column width={6}>
                {this.state.emailResponses.length > 0 && (
                  <HttpResponseList
                    onClickRemove={listing =>
                      this.onClickRemoveEmailResponse(listing)
                    }
                    responses={this.state.emailResponses}
                  />
                )}
              </Grid.Column> */}
        </Grid.Row>
      </Container>
    </div>
  );
};
export default App;
