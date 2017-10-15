import React, { Component } from 'react';
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

const client = new ApolloClient({
  networkInterface: createNetworkInterface({
    uri: 'http://localhost:3030/graphql'
  })
});

class App extends Component {
  render() {
    return (
      <ApolloProvider client={client}>
        <div className="App">
          <header className="App-header">
            <img src={logo} className="App-logo" alt="logo" />
            <h1 className="App-title">Welcome to Smart Chat</h1>
          </header>
          <ChannelsListWithData />
        </div>
      </ApolloProvider>
    );
  }
}

export default App;
