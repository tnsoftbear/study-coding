import { useContext } from "react";
import {
  ToggleButtonContext,
  ToggleButtonContextInterface,
} from "../ToggleButtonContext";
import { ButtonStateInterface } from "../ButtonState";
// import {ll} from "../../common/debug/Debug.js";

interface Props {
  idx: number;
}

const ToggleButton = (props: Props) => {
  // This type annotation is not necessary
  const { buttonStates, setButtonStates, setClickedButtonIndex } =
    useContext<ToggleButtonContextInterface>(ToggleButtonContext);

  const handleClick = (buttonIndex: number, newIsToggle: boolean) => {
    buttonStates.map((buttonState) => {
      buttonState.isToggle = false;
      return buttonState;
    });
    buttonStates[buttonIndex].isToggle = newIsToggle;
    buttonStates[buttonIndex].toggledCount += +newIsToggle;
    setButtonStates([...buttonStates]);
    setClickedButtonIndex(buttonIndex);
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
      onClick={() => handleClick(props.idx, !buttonStates[props.idx].isToggle)}
    >
      {drawContent(buttonStates[props.idx])}
    </button>
  );
};

export default ToggleButton;
