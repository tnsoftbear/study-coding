import { ToggleButtonContext } from "./ToggleButtonContext";
import DropRandomButton from "./drop-random/DropRandomButton";
import ToggleButtonList from "./toggle-button-list/ToggleButtonList";
// import { ll } from "../../common/debug/Debug.js";
import { reducer } from "./reducer/reducer";
import { useUndoReducer } from "./reducer/undo-reducer";
import AddToggleButton from "./add-toggle-button/AddToggleButton";
import RemoveToggleButton from "./remove-toggle-button/RemoveToggleButton";
import UndoButton from "./undo-redo/UndoButton";
import RedoButton from "./undo-redo/RedoButton";
import { UndoHistoryState } from "./UndoHistoryState";

interface Props {
  count: number;
}

/**
 * @param props
 * @returns
 */
const ToggleButtonsUseReducerPage = (props: Props) => {

  const [undoHistory, dispatch] = useUndoReducer(reducer, new UndoHistoryState(props.count));

  const { appState } = undoHistory;

  return (
    <div className="container">
      <h1>Toggle buttons with useReducer()</h1>
      <div className="row">
        <div className="col">
          <ToggleButtonContext.Provider
            value={{ appState, dispatch }}
          >
            <ToggleButtonList />
          </ToggleButtonContext.Provider>
        </div>
        <div className="col">
          <div>
            <DropRandomButton
              length={appState.buttonStates.length}
              dispatch={dispatch}
            />
          </div>
          <div>
            <AddToggleButton dispatch={dispatch} />
          </div>
          <div>
            <RemoveToggleButton
              length={appState.buttonStates.length}
              dispatch={dispatch}
            />
          </div>
          <div>
            <UndoButton dispatch={dispatch} />{" "}
            <RedoButton dispatch={dispatch} />
          </div>
        </div>
      </div>
    </div>
  );
};

export default ToggleButtonsUseReducerPage;
