import React, { Component } from "react";
import * as Toggle from "./ToggleButtons.js";

class ToggleButtonsPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      enabledFlags: Array(4).fill(false),
    };
  }

  toggleFlag = buttonIndex => value => {
     
    const enabledFlags = Array(4).fill(false);
    enabledFlags[buttonIndex] = value;
    this.setState(
        { enabledFlags }
    );
  }

  render() {
    const toggleButtons = [
      <Toggle.ToggleButton1 isToggleOn={this.state.enabledFlags[0]} parentToggleFlag={this.toggleFlag(0)} />,
      <Toggle.ToggleButton2 isToggleOn={this.state.enabledFlags[1]} parentToggleFlag={this.toggleFlag(1)} />,
      <Toggle.ToggleButton3 isToggleOn={this.state.enabledFlags[2]} parentToggleFlag={this.toggleFlag(2)} />,
      <Toggle.ToggleButton4 isToggleOn={this.state.enabledFlags[3]} parentToggleFlag={this.toggleFlag(3)} />,
    ];
    const listToggleButtons = toggleButtons.map((tb, i) => <li key={i}>{tb}</li>);

    return (
      <div>
        <ol>{listToggleButtons}</ol>
      </div>
    );
  }
}

export default ToggleButtonsPage;
