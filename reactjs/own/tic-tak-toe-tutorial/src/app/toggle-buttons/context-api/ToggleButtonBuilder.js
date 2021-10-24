import React, { Component } from "react";
import { ToggleButtonContext } from "./ToggleButtonsContextPage.js";

const ToggleButtonBuilder = () => {
  return class extends Component {
    constructor(props) {
      super(props);

      // This binding is necessary to make `this` work in the callback
      this.handleClick = this.handleClick.bind(this);
    }

    handleClick(isToggleOn) {
      this.props.doToggle(!isToggleOn);
    }

    render() {
      return (
        <ToggleButtonContext.Consumer>
          {(enabledFlags) => (
            <button
              onClick={() => this.handleClick(enabledFlags[this.props.idx])}
            >
              {enabledFlags[this.props.idx] ? "ON" : "OFF"}
            </button>
          )}
        </ToggleButtonContext.Consumer>
      );
    }
  };
};

export default ToggleButtonBuilder;
