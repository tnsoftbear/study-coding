import React, { Component } from "react";

class Clock extends Component {
  constructor(props) {
    super(props);
    this.state = {
        date: Date
    }
    this.state.date = new Date().toLocaleTimeString()
  }

  componentDidMount() {
    this.timerId = setInterval(() => {
        this.setState({ date: new Date().toLocaleTimeString() })
    }, 1000)
  }

  componentWillUnmount() {
      clearInterval(this.timerId)
  }

  render() {
    const divStyle = {
      color: this.props.textColor ?? 'blue',
    };
    return (
      <div>
        <h1 style={divStyle}>{this.state.date}</h1>
      </div>
    );
  }
}

export default Clock;
