
import React, { useEffect, useState ,useContext} from 'react';
import _ from 'lodash';
import {
    Button,
    Card,
    Form,
    Icon,
  } from 'semantic-ui-react';
import {WebSocketContext} from './WebSocketContextProvider';
import { ListingContext } from './ListingContextProvider';

export function Controller() {
    const {connect,disconnect,connected} = useContext(WebSocketContext);
    const {listings} = useContext(ListingContext);
    return (
        <Card>
        <Form>
          <Card.Content>
          <i>{`link count: ${listings.length}`}</i>
            <Card.Meta>
              <Button.Group>
                <Button
                  onClick={() => connect()}
                  icon
                >
                  <Icon name="play" />
                </Button>
                <Button onClick={() => disconnect()} icon>
                  <Icon name="stop" />
                </Button>
                {/* <Button onClick={() => onControllerChange({ touchY: 1 })} icon>
                <Icon name="redo" />
              </Button> */}
              </Button.Group>
            </Card.Meta>
            <Card.Description className="p5">
              <Button.Group>
                {/* <Button onClick={() => onControllerChange({ touchY: 1 })} icon>
                <Icon name="clock" />
              </Button> */}

                {/* <Button
                  active={this.state.emailAll}
                  onClick={() =>
                    this.setState({ emailAll: !this.state.emailAll })
                  }
                  icon
                >
                  <Icon name="envelope" />
                </Button> */}
              </Button.Group>
              {/* <input
                className="m-t-10"
                type="text"
                onChange={this.onInputChange}
                name="filter-field"
                placeholder="filter results"
              /> */}
              {/* <p className="margin0 padding0" align="left">
                <Icon
                  name="close icon"
                  onClick={this.onClickRemove}
                ></Icon>
                <a href={`${this.props.item}`}> {this.props.item}</a>
              </p>
              <div className="ui action input">
                <input type="text" placeholder="Search...">
                  <button className="ui button">Search</button>
                </input>
              </div> */}
            </Card.Description>
          </Card.Content>
        </Form>
      </Card>
    );
  }