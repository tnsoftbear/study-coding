import { useState } from "react";
import { ButtonState, ButtonStateInterface } from "./ButtonState";
import { ToggleButtonContext } from "./ToggleButtonContext";
import ClickedButtonInfo from "./footer/ClickedButtonInfo";
import DropRandomButton from "./drop-random/DropRandomButton";
import ToggleButtonList from "./toggle-button-list/ToggleButtonList";
// import { ll } from "../../common/debug/Debug.js";

interface Props {
  count: number;
}

/**
 * @param props
 * @returns
 */
const ToggleButtonsFunctionalWithHooksAndTsPage = (props: Props) => {
  const [buttonStates, setButtonStates] = useState<Array<ButtonStateInterface>>(
    () => Array.from({ length: props.count }, () => new ButtonState())
  );

  const [clickedButtonIndex, setClickedButtonIndex] = useState<number | null>(
    null
  );

  return (
    <div className="container">
      <h1>Toggle buttons functional with hook and typescript</h1>
      <div className="row">
        <div className="col">
          <ToggleButtonContext.Provider
            value={{ buttonStates, setButtonStates, setClickedButtonIndex }}
          >
            <ToggleButtonList count={props.count} />
          </ToggleButtonContext.Provider>
        </div>
        <div className="col">
          <DropRandomButton
            buttonStates={buttonStates}
            setButtonStates={setButtonStates}
          />
        </div>
      </div>
      <div className="row">
        <ClickedButtonInfo clickedButtonIndex={clickedButtonIndex} />
      </div>
    </div>
  );
};

export default ToggleButtonsFunctionalWithHooksAndTsPage;
