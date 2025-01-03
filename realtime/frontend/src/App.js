import React, { Component } from "react";
import "./App.css";
import { connect, sendMsg } from "./api";
import  Header from "./components/Header/Header.jsx";
import ChatHistory from "./components/ChatHistory/ChatHistory.jsx";

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      ChatHistory: []
    }
    // connect();
  }

  componentDidMount(){
    connect((msg) => {
      console.log("New Message")
      this.setState(prevState => ({
        ChatHistory: [...this.state.ChatHistory, msg]
      }))
      console.log(this.state)
    })
  }

  send() {
    console.log("hello");
    sendMsg("hello");
  }

  render() {
    return (
      <div className="App">
        <Header />
        <ChatHistory chatHistory={this.state.ChatHistory} />
        <button onClick={this.send}>Hit</button>
      </div>
    );
  }
}

export default App;