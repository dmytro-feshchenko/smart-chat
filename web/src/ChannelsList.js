import React, { Component } from 'react';
import {
  ApolloClient,
  gql,
  graphql,
  ApolloProvider,
} from 'react-apollo';

const ChannelsList = ({ data: {loading, error, channels }}) => {
  if (loading) {
    return <p>Loading channels. Please, wait for a moment</p>
  }
  if (error) {
    return <p>{error.message}</p>
  }
  return (
    <ul>
      {channels.map(ch => <li key={ch.id}>{ch.name}</li>)}
    </ul>
  )
}

const channelsListQuery = gql`
   query ChannelsListQuery {
     channels {
       id
       name
     }
   }
 `;
const ChannelsListWithData = graphql(channelsListQuery)(ChannelsList);

export default ChannelsListWithData;
