import { ToggleButtonContext } from "./ToggleButtonsFunctionalPage.js";

const ToggleButtonBuilder = () => {
  return function(props) {

    const handleClick = (isToggle) => {
      props.doToggle(!isToggle);
    }

    const drawContent = (buttonState) => {
      let output = buttonState.isToggle ? "ON" : "OFF";
      if (buttonState.toggledCount > 0) {
        output += ` (${buttonState.toggledCount})`;
      }
      return output;
    }

    return (
      <ToggleButtonContext.Consumer>
        {(buttonStates) => (
          <button
            onClick={() => handleClick(buttonStates[props.idx].isToggle)}
          >
            {drawContent(buttonStates[props.idx])}
          </button>
        )}
      </ToggleButtonContext.Consumer>
    );
  };
};

export default ToggleButtonBuilder;
