import { ApplicationState, ApplicationStateInterface } from "./ButtonState";
import { ActionType } from "./reducer/action-type";

export interface UndoHistoryStateInterface {
  appState: ApplicationStateInterface;
  undoHistory: Array<ApplicationStateInterface>;
  undoActions: Array<ActionType>;
}

export class UndoHistoryState implements UndoHistoryStateInterface {
  public appState: ApplicationStateInterface;
  public undoHistory: Array<ApplicationStateInterface>;
  public undoActions: Array<ActionType>;
  constructor(initialCount: number = 0) {
    this.appState = new ApplicationState(initialCount);
    this.undoHistory = [];
    this.undoActions = [];
  }
}
