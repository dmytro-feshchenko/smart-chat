import casual from 'casual';

const mocks = {
  String: () => 'It works!',
  Query: () => ({
  }),
  User: () => ({ name: () => casual.firstName }),
  Message: () => ({ text: () => casual.sentences(3) })
};

export default mocks;
