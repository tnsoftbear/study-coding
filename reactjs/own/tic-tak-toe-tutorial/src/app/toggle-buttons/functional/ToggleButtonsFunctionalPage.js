import React, { Component } from "react";
import ToggleButtonBuilder from "./ToggleButtonBuilder.js";
import ButtonState from "./ButtonState.js";

export const ToggleButtonContext = React.createContext();

class ToggleButtonsFunctionalPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      buttonStates: Array.from({length:props.count}, () => (new ButtonState()))
    };
  }

  toggle = (buttonIndex) => (newIsToggle) => {
    const buttonStates = this.state.buttonStates;
    for (let i = 0; i < buttonStates.length; i++) {
      buttonStates[i].isToggle = false;
    }
    buttonStates[buttonIndex].isToggle = newIsToggle;
    buttonStates[buttonIndex].toggledCount += newIsToggle;
    this.setState({ enabledFlags: buttonStates });
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
        <ToggleButtonContext.Provider value={this.state.buttonStates}>
          <ol>{listToggleButtons}</ol>
        </ToggleButtonContext.Provider>
      </div>
    );
  }
}

export default ToggleButtonsFunctionalPage;
