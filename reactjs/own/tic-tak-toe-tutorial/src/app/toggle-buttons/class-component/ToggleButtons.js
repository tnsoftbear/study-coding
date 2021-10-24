// https://reactjs.org/docs/handling-events.html
import React, { Component } from "react";

export class ToggleButton1 extends Component {
  constructor(props) {
    super(props);

    // This binding is necessary to make `this` work in the callback
    this.handleClick = this.handleClick.bind(this);
  }

  handleClick() {
    this.props.parentToggleFlag(!this.props.isToggleOn);
  }

  render() {
    return (
      <button onClick={this.handleClick}>
        {this.props.isToggleOn ? "ON" : "OFF"}
      </button>
    );
  }
}

export class ToggleButton2 extends Component {

  handleClick() {
    this.props.parentToggleFlag(!this.props.isToggleOn);
  }

  render() {
    return (
      <button onClick={() => this.handleClick()}>
        {this.props.isToggleOn ? "ON" : "OFF"}
      </button>
    );
  }
}

export class ToggleButton3 extends Component {

  handleClick = () => {
    this.props.parentToggleFlag(!this.props.isToggleOn);
  }

  render() {
    return (
      <button onClick={this.handleClick}>
        {this.props.isToggleOn ? "ON" : "OFF"}
      </button>
    );
  }
}

export class ToggleButton4 extends Component {

  handleClick() {
    this.props.parentToggleFlag(!this.props.isToggleOn);
  }

  render() {
    return (
      <button onClick={this.handleClick.bind(this)}>
        {this.props.isToggleOn ? "ON" : "OFF"}
      </button>
    );
  }
}