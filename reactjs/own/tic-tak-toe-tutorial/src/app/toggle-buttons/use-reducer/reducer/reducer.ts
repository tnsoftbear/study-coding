import { ApplicationStateInterface, ButtonState } from "../ButtonState";
import { ActionType } from "./action-type";

export type ReducerType = (
  state: ApplicationStateInterface,
  action: ActionType
) => ApplicationStateInterface;

export function reducer(
  state: ApplicationStateInterface,
  action: ActionType
): ApplicationStateInterface {
  switch (action.type) {
    case "toggle": {
      const { buttonIndex, newIsToggle } = action;
      const buttonStates = state.buttonStates;
      buttonStates.map((buttonState) => {
        buttonState.isToggle = false
        return buttonState
      });
      buttonStates[buttonIndex].isToggle = newIsToggle;
      buttonStates[buttonIndex].toggledCount += +newIsToggle;
      return {
        buttonStates,
      };
    }

    case "drop-random": {
      const buttonStates = state.buttonStates;
      const { droppedIndex } = action;
      buttonStates[droppedIndex].toggledCount = 0;
      buttonStates[droppedIndex].isToggle = false;
      return {
        buttonStates,
      };
    }

    case "add-toggle-button": {
      const buttonStates = state.buttonStates;
      buttonStates[buttonStates.length] = new ButtonState();
      return {
        buttonStates,
      };
    }

    case "remove-toggle-button": {
      const buttonStates = state.buttonStates;
      const { removeIndex } = action;
      return {
        buttonStates: [
          ...buttonStates.slice(0, removeIndex),
          ...buttonStates.slice(removeIndex + 1),
        ],
      };
    }

    default: {
      const buttonStates = state.buttonStates;
      return {
        buttonStates,
      };
    }
  }
}
