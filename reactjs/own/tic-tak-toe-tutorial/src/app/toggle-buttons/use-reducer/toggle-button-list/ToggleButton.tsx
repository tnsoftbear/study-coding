import { useContext } from "react";
import { ToggleButtonContext } from "../ToggleButtonContext";
import { ButtonStateInterface } from "../ButtonState";
// import {ll} from "../../common/debug/Debug.js";

interface Props {
  idx: number;
}

const ToggleButton = (props: Props) => {
  
  const { appState, dispatch } = useContext(ToggleButtonContext);

  const handleClick = (buttonIndex: number, newIsToggle: boolean) => {
    dispatch({ type: "toggle", buttonIndex, newIsToggle });
  };

  const drawContent = (buttonState: ButtonStateInterface) => {
    let output = buttonState.isToggle ? "ON" : "OFF";
    if (buttonState.toggledCount > 0) {
      output += ` (${buttonState.toggledCount})`;
    }
    return output;
  };

  return (
    <button
      className="btn btn-outline-primary"
      onClick={() => handleClick(props.idx, !appState.buttonStates[props.idx].isToggle)}
    >
      {drawContent(appState.buttonStates[props.idx])}
    </button>
  );
};

export default ToggleButton;
