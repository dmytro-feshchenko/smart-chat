import {
  makeExecutableSchema,
  addMockFunctionsToSchema,
} from 'graphql-tools';
import mocks from './mocks';
import resolvers from './resolvers';

const typeDefs = `
type Channel{
  id: ID!
  name: String,
  messages: [Message]
}
type User{
  id: Int,
  name: String,
  email: String
}
type Message{
  id: Int,
  text: String,
  owner: User
}
type Query{
  channels: [Channel]
}
`;

const schema = makeExecutableSchema({ typeDefs });

// addMockFunctionsToSchema({ schema, mocks });

export default schema;
