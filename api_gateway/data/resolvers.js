const resolvers = {
  Query: {
    channels() {
      return [
        {
          id: 1,
          name: "Channel 1",
        },
        {
          id: 2,
          name: "Channel 2",
        }
      ]
    }
  },
  Channel: {
    messages(channel) {
      return [
        { id: 1, name: 'Message 1' },
        { id: 2, name: 'Message 2' },
        { id: 3, name: 'Message 3' },
      ]
    }
  },
  Message: {
    author(message) {
      return { id: 1, name: 'John Doe' };
    }
  }
}
