import React, { Component } from 'react';

class Chat extends Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div>
        <p></p>
        <div className="chat__messages-list">
          <p>No messages yet.</p>
        </div>
        <div className="chat__send-message-form">
          <div>
            <input type="text" placeholder="Username" />
          </div>
          <div>
            <textarea placeholder="Type your message..." />
          </div>
          <div>
            <button type="button">Submit</button>
          </div>
        </div>
      </div>
    );
  }
}
