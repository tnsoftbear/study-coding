import React, { Component } from "react";
import ToggleButtonBuilder from "./ToggleButtonBuilder.js";

export const ToggleButtonContext = React.createContext();

export default class ToggleButtonsContextPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      enabledFlags: Array(props.count).fill(false),
    };
  }

  toggle = (buttonIndex) => (value) => {
    const enabledFlags = Array(4).fill(false);
    enabledFlags[buttonIndex] = value;
    this.setState({ enabledFlags });
  };

  render() {
    const toggleButtons = [];
    for (let i = 0; i < this.props.count; i++) {
      const ToggleButton = ToggleButtonBuilder();
      toggleButtons[i] = <ToggleButton idx={i} doToggle={this.toggle(i)} />;
    }
    const listToggleButtons = toggleButtons.map((tb) => (
      <li key={tb.props.idx}>{tb}</li>
    ));

    return (
      <div>
        <ToggleButtonContext.Provider value={this.state.enabledFlags}>
          <ol>{listToggleButtons}</ol>
        </ToggleButtonContext.Provider>
      </div>
    );
  }
}
