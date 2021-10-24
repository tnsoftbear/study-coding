import React from "react";

import { ApplicationState, ApplicationStateInterface } from "./ButtonState";

// Interface is not necessary
export interface ToggleButtonContextInterface {
   appState: ApplicationStateInterface,
   dispatch: React.Dispatch<any>,
};

export const ToggleButtonContext = React.createContext({
  appState: new ApplicationState(),
  dispatch: (value: any) => {},
});
