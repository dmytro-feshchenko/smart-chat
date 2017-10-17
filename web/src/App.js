import React, { Component } from 'react';
import {
  BrowserRouter as Router,
  Route,
  Link
} from 'react-router-dom'
import {
  ApolloClient,
  gql,
  graphql,
  ApolloProvider,
  createNetworkInterface
} from 'react-apollo';
import logo from './logo.svg';
import './App.css';
import ChannelsListWithData from './ChannelsList';
import Chat from './Chat';

const client = new ApolloClient({
  networkInterface: createNetworkInterface({
    uri: 'http://localhost:3030/graphql'
  })
});

class App extends Component {
  constructor(props) {
    super(props);
  }
  render() {
    return (
      <ApolloProvider client={client}>
        <div className="App">
          <header className="App-header">
            <img src={logo} className="App-logo" alt="logo" />
            <h1 className="App-title">Welcome to Smart Chat</h1>
          </header>

          <Router>
            <Route exact path="/" component={ChannelsListWithData} />
            <Route path="/channels/:id" component={Chat} />
          </Router>
        </div>
      </ApolloProvider>
    );
  }
}

export default App;
