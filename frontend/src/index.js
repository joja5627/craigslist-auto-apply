import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap/dist/js/bootstrap.bundle.min';
import 'semantic-ui-css/semantic.min.css';
import React from 'react';
import ReactDOM from 'react-dom';
import App from './App.js';
import registerServiceWorker from './registerServiceWorker';
import  {WebSocketContextProvider}  from './WebSocketContextProvider';
import  {ListingContextProvider}  from './ListingContextProvider';

ReactDOM.render(<WebSocketContextProvider>
    <ListingContextProvider>
      <App />
    </ListingContextProvider>
  </WebSocketContextProvider>,
  document.getElementById('root')
);
registerServiceWorker();
