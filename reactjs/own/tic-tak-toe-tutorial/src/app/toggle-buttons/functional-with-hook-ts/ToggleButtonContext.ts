import React from "react";

import { ButtonStateInterface} from "./ButtonState";

// Interface is not necessary
export interface ToggleButtonContextInterface {
   buttonStates: Array<ButtonStateInterface>,
   // setButtonStates: Dispatch<Array<ButtonStateInterface>>
   setButtonStates: (buttonStates: Array<ButtonStateInterface>) => void
   setClickedButtonIndex: (buttonIndex: number) => void
}

export const ToggleButtonContext = React.createContext({
  buttonStates: new Array<ButtonStateInterface>(),
  setButtonStates: (buttonStates: Array<ButtonStateInterface>) => {},
  setClickedButtonIndex: (buttonIndex: number) => {}
});
