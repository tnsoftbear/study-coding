import lodash from "lodash";
import { useReducer } from "react";
import { ApplicationStateInterface } from "../ButtonState";
import { UndoHistoryStateInterface } from "../UndoHistoryState";
import { ActionType } from "./action-type";
import { ReducerType } from "./reducer";
// import { ll } from "../../common/debug/Debug.js";

export const undo =
  (reducer: ReducerType) =>
  (state: UndoHistoryStateInterface, action: ActionType) => {
    let { undoHistory, undoActions, appState } = lodash.cloneDeep(state);
    switch (action.type) {
      case "undo": {
        if (undoActions.length) {
          undoActions.pop();
          appState = undoHistory.pop() as ApplicationStateInterface;
        }
        break;
      }
      case "redo": {
        if (undoActions.length) {
          const copyState = lodash.cloneDeep(appState);
          undoHistory = [...undoHistory, copyState];
          undoActions = [...undoActions, undoActions[undoActions.length - 1]];
          appState = reducer(appState, undoActions[undoActions.length - 1]);
        }
        break;
      }
      default: {
        const copyState = lodash.cloneDeep(appState);
        const copyAction = action;
        undoHistory = [...undoHistory, copyState];
        undoActions = [...undoActions, copyAction];
        appState = reducer(appState, action);
      }
    }
    return { appState, undoHistory, undoActions };
  };

export const useUndoReducer = (
  reducer: ReducerType,
  initialState: UndoHistoryStateInterface
) => useReducer(undo(reducer), initialState);
